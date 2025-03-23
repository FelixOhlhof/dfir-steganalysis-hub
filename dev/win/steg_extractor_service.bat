@echo off
set port=localhost:30126

cd /d ..\..\steg-services\extractor\
start cmd /k "color 0A & .venv\Scripts\python.exe server.py"