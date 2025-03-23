import os
import tempfile
import subprocess
from typing import Optional
from .extractor import Extractor


class OutguessExtractor(Extractor):
    """Extractor for data using outguess."""

    def extract(self, input_file: bytes) -> Optional[bytes]:
        with tempfile.NamedTemporaryFile(delete=True, suffix=".jpg") as temp_input_file:
            temp_input_file.write(input_file)
            temp_input_file_path = temp_input_file.name
            temp_output_file_path = tempfile.NamedTemporaryFile(delete=True).name

        try:
            command = ["outguess", "-r", temp_input_file_path, temp_output_file_path]
            result = subprocess.run(command, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True)

            if result.returncode != 0:
                raise RuntimeError(f"Outguess failed: {result.stderr}")

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
