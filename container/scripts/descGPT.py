
import openai
import re
import argparse
import logging
import os
import json
import ghidra_bridge

# Load Ghidra Bridge and make Ghidra namespace available
TIMEOUT = 1000
gb = ghidra_bridge.GhidraBridge(namespace=globals(), response_timeout=TIMEOUT)

def summarize_cpp_code(code):
    # Prepare the prompt for GPT-3.5
    prompt = "Code: " + code + "\n\nSummarize the functionality of the code:"

    # Set up OpenAI API credentials
    openai.api_key =   API_KEY # Replace with your OpenAI API keyPI

    # Generate the summary using GPT-3.5
    response = openai.Completion.create(
        engine='davinci',
        prompt=prompt,
        max_tokens=100,
        temperature=0.5,
        n=1,
        stop=None,
        timeout=10
    )

    # Extract the summary from the API response
    summary = response.choices[0].text.strip()

    return summary

# Example usage
cpp_code = """
#include <iostream>

int main() {
    std::cout << "Hello, World!" << std::endl;
    return 0;
}
"""

summary = summarize_cpp_code(cpp_code)
print("Code Summary:", summary)

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

    summary = summarize_cpp_code(function_body)
    j = {
        "name" : function_name,
        "parameters" : function_params,
        "body" : summary
    }

    functionsJSON["output"].append(j)

def extract_decomps():
    logging.info("Extracting decompiled functions...")
    decomp = ghidra.app.decompiler.DecompInterface()
    decomp.openProgram(currentProgram)
    functions = list(currentProgram.functionManager.getFunctions(True))
    failed_to_extract = []
    count = 0

    for function in functions:
        decomp_res = decomp.decompileFunction(function, TIMEOUT, monitor)

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

def main(function_name=None):
    """Main function."""
    program_info = get_program_info()

    # Default output directory to current directory + program name + _extraction
    extract_decomps() # Extract all
