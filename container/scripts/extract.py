#! /usr/bin/env python3
import re
import argparse
import logging
import os
import json
import ghidra_bridge

# Load Ghidra Bridge and make Ghidra namespace available
TIMEOUT = 1000
gb = ghidra_bridge.GhidraBridge(namespace=globals(), response_timeout=TIMEOUT)

functionsJSON = dict()
functionsJSON["output"] = []

def get_program_info():
    """Gather information for currentProgram in Ghidra."""
    logging.debug("Gathering program information...")
    program_info = {}
    program_info["program_name"] = currentProgram.getName()
    program_info["creation_date"] = gb.remote_eval("currentProgram.getCreationDate()")
    program_info["language_id"] = gb.remote_eval("currentProgram.getLanguageID()")
    program_info["compiler_spec_id"] = gb.remote_eval("currentProgram.getCompilerSpec().getCompilerSpecID()")
    
    #logging.info(f"Program Name: {program_info['program_name']}")
    #logging.info(f"Creation Date: {program_info['creation_date']}")
    #logging.info(f"Language ID: {program_info['language_id']}")
    #logging.info(f"Compiler Spec ID: {program_info['compiler_spec_id']}")

    return program_info

def matchregex(decomp_res, decomp_src, function):

    cout_pattern = r'(std::operator<<\(\(basic_ostream \*\)std::cout,(.*)\);)'
    cout_matches = re.findall(cout_pattern, decomp_src)

    # Replace the matched text with the new text
    for match in cout_matches:
        new_text = "std::cout << " + match[1] + ";"
        decomp_src = decomp_src.replace(match[0], new_text)

    # More regex simplifications will go here as we continue to research more ways to simplify Ghidra's outputs.

    return decomp_src


def getSubFunctionList(function, monitor):
    subFunctions = list(function.getCalledFunctions(monitor))
    nameList = []
    for subFunction in subFunctions:
        filename = f"{subFunction.name}@{subFunction.getEntryPoint()}.h"
        nameList.append(filename)

    return nameList

def write_function(function, subFunctionFilenames, decomp_res, decomp_src):
    decomp_src = matchregex(decomp_res, decomp_src, function)
    function_name = f"{function.name}@{function.getEntryPoint()}"
    function_params = [p.getFormalDataType().getDisplayName() for p in function.getParameters()]
    function_body = ""
    # write includes first
    for includeFilename in subFunctionFilenames:
        function_body += f"#include \"{includeFilename}\"\n"
    # write rest of file
    function_body += decomp_src
    j = {
        "name" : function_name,
        "parameters" : function_params,
        "body" : function_body
    }

    functionsJSON["output"].append(j)

def function_name_to_function(function_name):
    decomp = ghidra.app.decompiler.DecompInterface()
    decomp.openProgram(currentProgram)
    functions = list(currentProgram.functionManager.getFunctions(True))

    for function in functions:
        if function_name == f"{function.name}@{function.getEntryPoint()}":
            return function
    
    logging.error(f"Failed to find match for {function.name}")
    return None

def extract_lazy(entry_function):
    logging.info("Extracting decompiled functions...")
    decomp = ghidra.app.decompiler.DecompInterface()
    decomp.openProgram(currentProgram)

    failed_to_extract = set()
    count = 0
    functions_seen = set()

    def lazy_dfs(function):
        if f"{function.name}@{function.getEntryPoint()}" in functions_seen or f"{function.name}@{function.getEntryPoint()}" in failed_to_extract:
            return

        logging.debug(f"Decompiling {function.name}")
        decomp_res = decomp.decompileFunction(function, TIMEOUT, monitor)

        if decomp_res.isTimedOut():
            logging.warning("Timed out while attempting to decompile '{function.name}'")
        elif not decomp_res.decompileCompleted():
            logging.error(f"Failed to decompile {function.name}")
            logging.error("    Error: " + decomp_res.getErrorMessage())
            failed_to_extract.add(function.name)
            return

        printer = ghidra.app.decompiler.PrettyPrinter(function, decomp_res.getCCodeMarkup())
        decomp_src = printer.print(False).getC()

        # get functions called by this function
        subFunctionFilenames = getSubFunctionList(function, monitor)
        try: 
            write_function(function, subFunctionFilenames, decomp_res, decomp_src)
            functions_seen.add(f"{function.name}@{function.getEntryPoint()}")
            count += 1
        except Exception as e:
            logging.error(e)
            failed_to_extract.add("{function.name}@{function.getEntryPoint()}" )
            return

        subFunctions = list(function.getCalledFunctions(monitor))
        for subFunction in subFunctions:
            lazy_dfs(subFunction)

        if failed_to_extract:
            logging.warning("Failed to extract the following functions:\n\n  - " + "\n  - ".join(failed_to_extract))
        with open("/container/output/output.json", "w") as f:
            f.write(json.dumps(functionsJSON,indent=4))
            f.close()
        count = len(functions_seen)
        logging.info(f"Extracted {str(count)} out of {str(count + len(failed_to_extract))} functions")

    lazy_dfs(entry_function)


def extract_decomps():
    logging.info("Extracting decompiled functions...")
    decomp = ghidra.app.decompiler.DecompInterface()
    decomp.openProgram(currentProgram)
    functions = list(currentProgram.functionManager.getFunctions(True))
    failed_to_extract = []
    count = 0

    for function in functions:
        logging.debug(f"Decompiling {function.name}")
        decomp_res = decomp.decompileFunction(function, TIMEOUT, monitor)

        if decomp_res.isTimedOut():
            logging.warning("Timed out while attempting to decompile '{function.name}'")
        elif not decomp_res.decompileCompleted():
            logging.error(f"Failed to decompile {function.name}")
            logging.error("    Error: " + decomp_res.getErrorMessage())
            failed_to_extract.add(function.name)
            continue
        


        printer = ghidra.app.decompiler.PrettyPrinter(function, decomp_res.getCCodeMarkup())
        decomp_src = printer.print(False).getC()
        

        # get functions called by this function

        subFunctionFilenames = getSubFunctionList(function, monitor)

        try:
            write_function(function, subFunctionFilenames, decomp_res, decomp_src)
            count += 1
        except Exception as e:
            logging.error(e)
            failed_to_extract.add(function.name)
            continue
    with open("/container/output/output.json", "w") as f:
        f.write(json.dumps(functionsJSON,indent=4))
        f.close()
    logging.info(f"Extracted {str(count)} out of {str(len(functions))} functions")
    if failed_to_extract:
        logging.warning("Failed to extract the following functions:\n\n  - " + "\n  - ".join(failed_to_extract))

def main(function_name=None):
    """Main function."""
    program_info = get_program_info()

    # Default output directory to current directory + program name + _extraction
    if function_name:
        extract_lazy(function_name_to_function(function_name))
    else:
        extract_decomps() # Extract all




if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Extract ghidra decompilation output for currently loaded program.")
    parser.add_argument("-v", "--verbose", action="count", help="Display verbose logging output")
    parser.add_argument("-t", "--timeout", type=int, help="Custom timeout for individual function decompilation (default = 1000)")
    parser.add_argument("-f", "--function", type=str, help="The function name do be analyzed if one")
    args = parser.parse_args()

    if args.function:
        function = args.function
    else:
        function = None
    
    if args.verbose:
        logging.getLogger().setLevel(logging.DEBUG)
    else:
        logging.getLogger().setLevel(logging.INFO)

    if args.timeout:
        TIMEOUT = args.TIMEOUT

    main(function_name=function)