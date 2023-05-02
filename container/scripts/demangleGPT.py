import openai
import re
import argparse
import logging
import os
import json
import ghidra_bridge
from dotenv import load_dotenv

# Load Ghidra Bridge and make Ghidra namespace available
TIMEOUT = 1000
gb = ghidra_bridge.GhidraBridge(namespace=globals(), response_timeout=TIMEOUT)
functionsJSON = dict()
functionsJSON["output"] = []
load_dotenv()

def demangle_cpp_code(code):
    # Prepare the prompt for GPT-3.5
    prompt = "Given the following code: " + code + "\n\nRewrite the code, but rename function parameters and variables to better describe their purposes. Don't rename functions or #include statements. Just respond with the rewritten code and nothing else."

    # Set up OpenAI API credentials
    openai.api_key = os.getenv("WGHIDRA_SECRET_KEY") # Replace with your OpenAI API keyPI

    completion = openai.ChatCompletion.create(
    model="gpt-3.5-turbo",
    messages=[
        {"role": "user", "content": prompt}
    ]
    )

    # Extract the summary from the API response
    summary = completion.choices[0].message.content

    return summary

def getSubFunctionList(function, monitor):
    subFunctions = list(function.getCalledFunctions(monitor))
    nameList = []
    for subFunction in subFunctions:
        filename = f"{subFunction.name}@{subFunction.getEntryPoint()}.h"
        nameList.append(filename)

    return nameList

def write_function(function, subFunctionFilenames, decomp_res, decomp_src):
    function_name = f"{function.name}@{function.getEntryPoint()}"
    function_body = ""
    # write includes first
    for includeFilename in subFunctionFilenames:
        function_body += f"#include \"{includeFilename}\"\n"
    # write rest of file
    function_body += decomp_src

    demangled_body = demangle_cpp_code(function_body)
    j = {
        "name" : function_name,
        "body" : demangled_body
    }

    functionsJSON["output"].append(j)

def extract_decomps():
    logging.info("Extracting decompiled functions...")
    decomp = ghidra.app.decompiler.DecompInterface()
    decomp.openProgram(currentProgram)
    functions = list(currentProgram.functionManager.getFunctions(True))
    failed_to_extract = set()
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
            failed_to_extract.append(function.name)
            continue
    with open("/container/output/output.json", "w") as f:
        f.write(json.dumps(functionsJSON,indent=4))
        f.close()
    logging.info(f"Extracted {str(count)} out of {str(len(functions))} functions")
    if failed_to_extract:
        logging.warning("Failed to extract the following functions:\n\n  - " + "\n  - ".join(failed_to_extract))


extract_decomps() # Extract and demangle all
