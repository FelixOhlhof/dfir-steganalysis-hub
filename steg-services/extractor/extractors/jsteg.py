import os
import tempfile
import subprocess
from typing import Optional
from .extractor import Extractor


class JstegExtractor(Extractor):
    """Extractor for data using jsteg."""

    def extract(self, input_file: bytes) -> Optional[bytes]:
        with tempfile.NamedTemporaryFile(delete=False, suffix=".jpg") as temp_input_file:
            temp_input_file.write(input_file)
            temp_input_file_path = temp_input_file.name
            temp_output_file_path = tempfile.NamedTemporaryFile(delete=False).name

        try:
            bin = os.path.join("jsteg", "jsteg-windows-amd64.exe") if os.name == 'nt' else os.path.join("jsteg", "jsteg-linux-amd64")
            command = [bin, "reveal", temp_input_file_path, temp_output_file_path]
            result = subprocess.run(command, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True)

            if result.returncode != 0:
                raise RuntimeError(f"Jsteg failed: {result.stderr}")

            with open(temp_output_file_path, "rb") as temp_output_file:
                output_bytes = temp_output_file.read()

            return output_bytes
        finally:
            try:
                if os.path.exists(temp_input_file_path):
                    os.remove(temp_input_file_path)
            except Exception as e: 
                print(e)  
            try:
                if os.path.exists(temp_output_file_path):
                    os.remove(temp_output_file_path) 
            except Exception as e: 
                print(e)  
        
