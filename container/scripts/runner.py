# USAGE: python runner.py [function_name] ..argv
# script that creates runnable file for whatever is selected

# 1. Create a main file [function_name]-main.c, with main method that only includes & calls [function_name]
# 2. Populate arguments values
# 3. gcc [function_name]-main.c -o [function_name]-main.o
# 4. ./[function_name]-main.o

import sys
import os
import subprocess
import json

MAIN_TEMPLATE = '''#include <iostream>
#include "{file_name}.h"

int main() {{
    std::cout << {function_name}({args}) << std::endl;
    return 0;
}} 
'''

def normalize_args(argv):
    # TODO
    pass

def runner(executable_path):

    # Run the GCC command using subprocess
    print(f"--- EXECUTING {executable_path} ---")
    try:
        output = subprocess.check_output(executable_path, stderr=subprocess.STDOUT)
        f = open("/container/output/output.json", "w")
        j = {
            "output" : output.decode()
        }
        f.write(json.dumps(j))
        f.close()
    except subprocess.CalledProcessError as e:
        print(f"--- EXECUTING FAILED ---")
        print("Compilation failed:", e.output.decode())
        exit(1)

def create_main(filename, args):
    print(f"--- CREATING MAIN {filename} ---")
    output_dir = "/container/output/"
    
    try:
        main_filename = f"{filename}-main.c"
        path = os.path.join(output_dir, main_filename)
        function_name = filename.split("@")[0]
        with open(path, "w") as f:
            f.write(MAIN_TEMPLATE.format(file_name = filename, function_name=function_name, args=','.join(args)))
        return path
    except Exception as e:
        print(f"--- CREATING MAIN FAILED ---")
        print(e)
        exit(1)
    

def compile(main_file):
    print(f"--- COMPILING {main_file} ---")
    # Define the GCC command to run
    output_name = main_file.split('.c')[0]+".o"
    gcc_command = ["g++", main_file, "-o", output_name]

    # Run the GCC command using subprocess
    try:
        output = subprocess.check_output(gcc_command, stderr=subprocess.STDOUT)
        if len(output.decode()) != 0:
            print(output.decode())
    except subprocess.CalledProcessError as e:
        print(f"--- COMPILING FAILED ---")
        print("Compilation failed:", e.output.decode())
        exit(1)

    return output_name

def read_json():
    f = open("/container/input/input.json", "r")
    data = json.loads(f.read())
    for func in data["functions"]:
        header = open("/container/output/" + func["name"] + ".h", "w")
        header.write(func["body"])
        header.close()
# filename arg must come in the format {function.name}@{function.getEntryPoint()}
def main(): 
    if len(sys.argv) < 2:
        print("No filename passed")
        return 1
    read_json()
    main_file = create_main(sys.argv[1], sys.argv[2:])
    executable_path = compile(main_file)
    runner(executable_path)

if __name__ == '__main__':
    main()


