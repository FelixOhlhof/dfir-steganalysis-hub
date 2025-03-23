import tempfile
import subprocess
import json
import os

class BinwalkAnalyzer:
    def analyze(self, data: bytes):
        with tempfile.NamedTemporaryFile(delete=False) as tmp_file:
            tmp_file.write(data)
            tmp_file.flush()
            tmp_file_path = tmp_file.name  

        with tempfile.NamedTemporaryFile(delete=False, suffix=".json") as json_tmp_file:
            results_json_path = json_tmp_file.name

        try:
            subprocess.run(["binwalk", f"--log={results_json_path}", tmp_file_path], check=True)

            if os.path.exists(results_json_path):
                with open(results_json_path, "r") as json_file:
                    binwalk_results = json.load(json_file)
            else:
                binwalk_results = {"error": "Binwalk results file not found"}

        except subprocess.CalledProcessError as e:
            binwalk_results = {"error": f"Binwalk execution failed: {str(e)}"}

        finally:
            os.remove(tmp_file_path)
            if os.path.exists(results_json_path):
                os.remove(results_json_path)

        return binwalk_results

    def entropy(self, data: bytes):
        with tempfile.NamedTemporaryFile(delete=False) as tmp_file:
            tmp_file.write(data)
            tmp_file.flush()
            tmp_file_path = tmp_file.name

        with tempfile.NamedTemporaryFile(delete=False, suffix=".json") as json_tmp_file:
            results_json_path = json_tmp_file.name 
        try:
            subprocess.run(["binwalk", "--entropy", f"--log={results_json_path}", tmp_file_path], check=True)

            if os.path.exists(results_json_path):
                with open(results_json_path, "r") as json_file:
                    binwalk_results = json.load(json_file)
            else:
                binwalk_results = {"error": "Binwalk results file not found"}

        except subprocess.CalledProcessError as e:
            binwalk_results = {"error": f"Binwalk execution failed: {str(e)}"}

        finally:
            os.remove(tmp_file_path)
            if os.path.exists(results_json_path):
                os.remove(results_json_path)

        return binwalk_results
