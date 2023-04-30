#!/bin/bash

python -m ghidra_bridge.install_server ./scripts
ghidra-analyzeHeadless ./ghidra ghidra_project -import ./input/$1 -scriptPath ./scripts -postscript ghidra_bridge_server.py
