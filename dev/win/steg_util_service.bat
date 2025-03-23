@echo off
set port=localhost:30124

cd /d ..\..\steg-services\util\
start cmd /k "color 0A & .venv\Scripts\python.exe server.py"