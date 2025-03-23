@echo off
set port=localhost:30125
set "octave_path=C:\Program Files\GNU Octave\Octave-9.2.0\octave.vbs"

cd /d ..\..\steg-services\aletheia\
start cmd /k "color 0A & .venv\Scripts\python.exe server.py"