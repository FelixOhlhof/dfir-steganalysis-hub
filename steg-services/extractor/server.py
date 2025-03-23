import os
import concurrent
import threading
import time
import grpc
import steg_service_pb2_grpc as pbgrpc
import steg_service_pb2 as pb

from google.protobuf.struct_pb2 import Struct, ListValue
from image_analyzer import ImageAnalyzer
from binwalk_analyzer import BinwalkAnalyzer
from extractors import *
import concurrent.futures
import sys
from pathlib import Path
#sys.path.append(str(Path(__file__).resolve().parent.parent))
from pyutils import util

class ExtractorService():
    def __init__(self, max_workers):
        self.executor = concurrent.futures.ThreadPoolExecutor(max_workers=max_workers)
        self.max_timeout = 5

        self.img_analizer = ImageAnalyzer()
        self.binwalk_analyzer = BinwalkAnalyzer()
        self.comment_extractor = CommentExtractor() 
        self.eoi_extractor = AppendToEoiExtractor() 
        self.outguess_extractor = OutguessExtractor() 
        self.steghide_extractor = SteghideExtractor() 
        self.lsbsteg_extractor = LSBSteganographyExtractor() 
        self.lsb_extractor = LsbExtractor()
        self.jsteg_extractor = JstegExtractor()

        self.service_info = pb.StegServiceInfo()
        self.service_info.name = "extractor"
        self.service_info.description = "Functions for extracting data from images."
        
        func_extract_exif = pb.StegServiceFunction(name="extract_exif", description="Extracts all exif fields.")
        func_extract_exif.return_fields.append(pb.StegServiceReturnFieldDefinition(name="exif_data", label="exif_data", type=pb.Type.DICT, description="Extracted exif data as dictionary."))
        self.service_info.functions.append(func_extract_exif)

        func_extract_comment = pb.StegServiceFunction(name="extract_comment", description="Extractor for data embedded into the comment section.")
        func_extract_comment.supported_file_types.append("png")
        func_extract_comment.supported_file_types.append("jpg")
        func_extract_comment.supported_file_types.append("tiff")
        func_extract_comment.return_fields.append(pb.StegServiceReturnFieldDefinition(name="data", label="data", type=pb.Type.BYTES, description="Extracted data."))
        self.service_info.functions.append(func_extract_comment)
        
        func_extract_eoi = pb.StegServiceFunction(name="extract_eoi", description="Extractor for data appended to the end-of-file marker.")
        func_extract_eoi.supported_file_types.append("jpg")
        func_extract_eoi.supported_file_types.append("png")
        func_extract_eoi.supported_file_types.append("gif")
        func_extract_eoi.return_fields.append(pb.StegServiceReturnFieldDefinition(name="data", label="data", type=pb.Type.BYTES, description="Extracted data."))
        self.service_info.functions.append(func_extract_eoi)
        
        func_extract_outguess = pb.StegServiceFunction(name="extract_outguess", description="Extract data using outguess.")
        func_extract_outguess.supported_file_types.append("jpg")
        func_extract_outguess.return_fields.append(pb.StegServiceReturnFieldDefinition(name="data", label="data", type=pb.Type.BYTES, description="Extracted data."))
        self.service_info.functions.append(func_extract_outguess)
        
        func_extract_steghide = pb.StegServiceFunction(name="extract_steghide", description="Extract data using steghide.")
        func_extract_steghide.supported_file_types.append("jpg")
        func_extract_steghide.return_fields.append(pb.StegServiceReturnFieldDefinition(name="data", label="data", type=pb.Type.BYTES, description="Extracted data."))
        self.service_info.functions.append(func_extract_steghide)
        
        func_extract_jsteg = pb.StegServiceFunction(name="extract_jsteg", description="Extract data using jsteg.")
        func_extract_jsteg.supported_file_types.append("jpg")
        func_extract_jsteg.return_fields.append(pb.StegServiceReturnFieldDefinition(name="data", label="data", type=pb.Type.BYTES, description="Extracted data."))
        self.service_info.functions.append(func_extract_jsteg)
        
        func_extract_lsbs = pb.StegServiceFunction(name="extract_lsbs", description="Extract LSBs.")
        func_extract_lsbs.supported_file_types.append("png")
        func_extract_lsbs.return_fields.append(pb.StegServiceReturnFieldDefinition(name="lsbs", label="lsbs", type=pb.Type.BYTES, description="Extracted lsbs."))
        self.service_info.functions.append(func_extract_lsbs)
                
        func_get_img_type = pb.StegServiceFunction(name="extract_img_type", description="Retrieves image type based on common magic bytes.")
        func_get_img_type.return_fields.append(pb.StegServiceReturnFieldDefinition(name="img_type", label="img_type", type=pb.Type.STRING, description="Image types:\nPNG\nJPG\nBMP\nGIF\nTIF\nSVG\nICO\nWEBP\nUNKNOWN"))
        self.service_info.functions.append(func_get_img_type)
        
        func_extract_strings = pb.StegServiceFunction(name="extract_strings", description="Executes string search.")
        func_extract_strings.return_fields.append(pb.StegServiceReturnFieldDefinition(name="strings", label="strings", type=pb.Type.LIST, description="List of extracted string."))
        func_extract_strings.parameter.append(pb.StegServiceParameterDefinition(name="min_length", default="6", description="Amount of minimal consecutively utf-8 characters.", optional=True, type=pb.Type.INT))
        func_extract_strings.parameter.append(pb.StegServiceParameterDefinition(name="unique", default="true", description="If true duplicate occurrences will be removed", optional=True, type=pb.Type.BOOL))
        self.service_info.functions.append(func_extract_strings)

        func_analyze = pb.StegServiceFunction(name="binwalk_analyze", description="Scans the file using binwalk.")
        func_analyze.return_fields.append(pb.StegServiceReturnFieldDefinition(name="result", label="result", type=pb.Type.DICT, description="Analysis result as dictionary."))
        self.service_info.functions.append(func_analyze)
        
        func_extract = pb.StegServiceFunction(name="binwalk_entropy", description="Calculates entropy using binwalk.")
        func_extract.return_fields.append(pb.StegServiceReturnFieldDefinition(name="result", label="result", type=pb.Type.DICT, description="Scan result as dictionary."))
        self.service_info.functions.append(func_extract)

        func_compare_exif = pb.StegServiceFunction(name="compare_exif", description="Compares exif fields of two images.")
        func_compare_exif.parameter.append(pb.StegServiceParameterDefinition(name="image_b", description="Image to compare with.", type=pb.Type.BYTES, optional=False))
        func_compare_exif.supported_file_types.append("png")
        func_compare_exif.supported_file_types.append("jpg")
        func_compare_exif.return_fields.append(pb.StegServiceReturnFieldDefinition(name="exif_diff", label="exif_diff", type=pb.Type.DICT, description="A dictionary with the EXIF field name as the key and the different values in the form “ExifData_a:ExifData_b” as the value."))
        self.service_info.functions.append(func_compare_exif)

    def Execute(self, request : pb.StegServiceRequest, context):
        print(f"Received request from {context.peer()} for function {request.function}")

        timeout = request.request_timeout_sec if request.request_timeout_sec != 0 else self.max_timeout
        result_container = {"response": None}
        stop_event = threading.Event() 

        def task():
            try:
                result_container["response"] = self._execute(request)
            except Exception as e:
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(str(e))
                result_container["response"] = pb.StegServiceResponse()
            finally:
                stop_event.set()

        task_thread = threading.Thread(target=task, daemon=True)
        task_thread.start()

        start_time = time.time()

        while time.time() - start_time < timeout:
            if not context.is_active():
                print("Client disconnected before completion.")
                stop_event.set()  
                return pb.StegServiceResponse()

            if result_container["response"] is not None:
                return result_container["response"]

            time.sleep(0.05)

        context.set_code(grpc.StatusCode.DEADLINE_EXCEEDED)
        context.set_details(f"Request timeout exceeded after {timeout} seconds")
        stop_event.set()  
        return pb.StegServiceResponse()

    def _execute(self, request : pb.StegServiceRequest):
        match request.function:
            case "extract_exif":
                return self.extract_exif(request)
            case "extract_img_type":
                return self.extract_img_type(request)
            case "extract_comment":
                data = self.comment_extractor.extract(request.file)
            case "extract_eoi":
                data = self.eoi_extractor.extract(request.file)
            case "extract_outguess":
                data = self.outguess_extractor.extract(request.file)
            case "extract_steghide":
                data = self.steghide_extractor.extract(request.file)
            case "extract_lsbs":
                data = self.lsb_extractor.extract(request.file)
            case "extract_jsteg":
                data = self.jsteg_extractor.extract(request.file)
            case "extract_strings":
                return self.extract_strings(request)
            case "binwalk_analyze":
                return self.binwalk_analyze(request)
            case "binwalk_entropy":
                return self.binwalk_entropy(request)
            case "compare_exif":
                return self.compare_exif(request)
            case _:
                raise Exception(f"Function {request.function} not implemented")
            
        response = pb.StegServiceResponse()
        if data != None:
            response.values["data"].binary_value = data
        return response
    
    def GetStegServiceInfo(self, request, context):
        print(f"recieved request from {context.peer()}")
        return self.service_info
    
    def extract_exif(self, request: pb.StegServiceRequest):
        exif_data = self.img_analizer.extract_exif(request.file)
        json_compatible_data = util.make_json_compatible(exif_data)

        response = pb.StegServiceResponse()
        struct_data = Struct()
        struct_data.update(json_compatible_data)
        response.values["exif_data"].structured_value.struct_value = struct_data
        return response
    
    def extract_img_type(self, request: pb.StegServiceRequest):
        img_type = self.img_analizer.get_image_type(request.file)

        response = pb.StegServiceResponse()
        response.values["img_type"].string_value = img_type.name
        return response

    def extract_strings(self, request: pb.StegServiceRequest):
        min_length = util.get_parameter(self.service_info, "min_length", request)
        unique = util.get_parameter(self.service_info, "unique", request)
        strings = self.img_analizer.extract_strings(request.file, min_length)
        if unique:
            strings = set(strings)

        listValue = ListValue()
        listValue.extend(strings)
        response = pb.StegServiceResponse()
        response.values["strings"].structured_value.list_value = listValue
        return response
    
    def binwalk_analyze(self, request: pb.StegServiceRequest):
        result = self.binwalk_analyzer.analyze(request.file)
        
        json_compatible_data = util.make_json_compatible(result)

        response = pb.StegServiceResponse()
        struct_data = Struct()
        struct_data.update(json_compatible_data[0])
        response.values["result"].structured_value.struct_value = struct_data
        return response
    
    def binwalk_entropy(self, request: pb.StegServiceRequest):
        result = self.binwalk_analyzer.entropy(request.file)
        
        json_compatible_data = util.make_json_compatible(result)

        response = pb.StegServiceResponse()
        struct_data = Struct()
        struct_data.update(json_compatible_data[0])
        response.values["result"].structured_value.struct_value = struct_data
        return response
    
    def compare_exif(self, request: pb.StegServiceRequest):
        image_b = util.get_parameter(self.service_info, "image_b", request)
        exif_data = self.img_analizer.compare_exif(request.file, image_b)
        json_compatible_data = util.make_json_compatible(exif_data)

        response = pb.StegServiceResponse()
        struct_data = Struct()
        struct_data.update(json_compatible_data)
        response.values["exif_diff"].structured_value.struct_value = struct_data
        return response
        

def serve():
    port = os.environ['port']
    if ":" not in port:
        port = f'[::]:{port}'
    max_workers = os.environ.get("max_workers")
    if max_workers == None:
        max_workers = 20
        
    
    MAX_MESSAGE_SIZE = 40 * 1024 * 1024
    
    extractor_svc = ExtractorService(max_workers)

    grpc_executor = concurrent.futures.ThreadPoolExecutor(max_workers=max_workers)
    server = grpc.server(
        grpc_executor,
        options=[
            ("grpc.max_send_message_length", MAX_MESSAGE_SIZE),
            ("grpc.max_receive_message_length", MAX_MESSAGE_SIZE)
        ],
    )

    pbgrpc.add_StegServiceServicer_to_server(extractor_svc, server)
    server.add_insecure_port(port)
    print(f"{extractor_svc.service_info.name} service started on port {port}")
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    serve()
