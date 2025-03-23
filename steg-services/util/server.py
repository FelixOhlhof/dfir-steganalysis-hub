import os
import grpc
import uuid
from google.protobuf.struct_pb2 import Struct
import steg_service_pb2_grpc as pbgrpc
import steg_service_pb2 as pb

from file_storage import FileStorage
from image_converter import ImageConverter
from file_stats import FileStats
from concurrent import futures

import sys
from pathlib import Path
#sys.path.append(str(Path(__file__).resolve().parent.parent))
from pyutils import util, file_utils



class UtilService():
    def __init__(self):
        self.file_handler = FileStorage()
        self.image_converter = ImageConverter()
        self.file_stats = FileStats()

        self.service_info = pb.StegServiceInfo()
        self.service_info.name = "util"
        self.service_info.description = "Collection of utility functions."
        
        func_save_file = pb.StegServiceFunction(name="save_file", description="Saves temporary file.")
        func_save_file.parameter.append(pb.StegServiceParameterDefinition(name="file_name", description="Temporary file name", type=pb.Type.STRING, optional=True))
        func_save_file.parameter.append(pb.StegServiceParameterDefinition(name="life_span", description="Temporary file life span in minutes. -1=unlimited", type=pb.Type.INT, optional=True, default="10"))
        func_save_file.parameter.append(pb.StegServiceParameterDefinition(name="file_limit", description="Will limit the maximal amount of file stats to return.", type=pb.Type.INT, optional=True, default="100"))
        func_save_file.return_fields.append(pb.StegServiceReturnFieldDefinition(name="file_name", label="file_name", type=pb.Type.STRING, description="Temporary file name which can be accessed."))
        func_save_file.return_fields.append(pb.StegServiceReturnFieldDefinition(name="expires_at", label="expires at", type=pb.Type.STRING, description="Date of expiration."))
        func_save_file.return_fields.append(pb.StegServiceReturnFieldDefinition(name="data_type", label="data_type", type=pb.Type.STRING, description="Data type."))
        func_save_file.return_fields.append(pb.StegServiceReturnFieldDefinition(name="sha256", label="sha256", type=pb.Type.STRING, description="Sha-256 Hash."))
        func_save_file.return_fields.append(pb.StegServiceReturnFieldDefinition(name="file_type", label="file_type", type=pb.Type.STRING, description="File type based on magic bytes.\nPNG\nJPEG\nPDF\nZIP\nGIF\nELF\nPE\nBMP\nGZIP\nOGG\nFLAC\nMP3\nICO\nCUR\nMP4\nMIDI\nWAV\nUnknown"))
        func_save_file.return_fields.append(pb.StegServiceReturnFieldDefinition(name="stats", label="stats", type=pb.Type.DICT, description="Info of current files."))
        self.service_info.functions.append(func_save_file)

        func_get_last_file = pb.StegServiceFunction(name="get_last_file", description="Gets the last saved temporary file.")
        func_get_last_file.parameter.append(pb.StegServiceParameterDefinition(name="delete_after", description="If true the file will be deleted.", type=pb.Type.BOOL, optional=True, default="False"))
        func_get_last_file.parameter.append(pb.StegServiceParameterDefinition(name="file_limit", description="Will limit the maximal amount of file stats to return.", type=pb.Type.INT, optional=True, default="100"))
        func_get_last_file.return_fields.append(pb.StegServiceReturnFieldDefinition(name="file_data", label="file_data", type=pb.Type.BYTES, description="File data."))
        func_get_last_file.return_fields.append(pb.StegServiceReturnFieldDefinition(name="stats", label="stats", type=pb.Type.DICT, description="Info of current files."))
        func_get_last_file.return_fields.append(pb.StegServiceReturnFieldDefinition(name="sha256", label="sha256", type=pb.Type.STRING, description="Sha-256 Hash."))
        func_get_last_file.return_fields.append(pb.StegServiceReturnFieldDefinition(name="file_type", label="file_type", type=pb.Type.STRING, description="File type based on magic bytes.\nPNG\nJPEG\nPDF\nZIP\nGIF\nELF\nPE\nBMP\nGZIP\nOGG\nFLAC\nMP3\nICO\nCUR\nMP4\nMIDI\nWAV\nUnknown"))
        func_get_last_file.file_optional = True
        self.service_info.functions.append(func_get_last_file)
        
        func_get_file = pb.StegServiceFunction(name="get_file", description="Retrieves temporary file.")
        func_get_file.parameter.append(pb.StegServiceParameterDefinition(name="file_name", description="File name to retrieve.", type=pb.Type.STRING, optional=False))
        func_get_file.parameter.append(pb.StegServiceParameterDefinition(name="delete_after", description="If true the file will be deleted.", type=pb.Type.BOOL, optional=True, default="False"))
        func_get_file.parameter.append(pb.StegServiceParameterDefinition(name="file_limit", description="Will limit the maximal amount of file stats to return.", type=pb.Type.INT, optional=True, default="100"))
        func_get_file.return_fields.append(pb.StegServiceReturnFieldDefinition(name="file_data", label="file_data", type=pb.Type.BYTES, description="File data."))
        func_get_file.return_fields.append(pb.StegServiceReturnFieldDefinition(name="stats", label="stats", type=pb.Type.DICT, description="Info of current files."))
        func_get_file.return_fields.append(pb.StegServiceReturnFieldDefinition(name="sha256", label="sha256", type=pb.Type.STRING, description="Sha-256 Hash."))
        func_get_file.return_fields.append(pb.StegServiceReturnFieldDefinition(name="file_type", label="file_type", type=pb.Type.STRING, description="File type based on magic bytes.\nPNG\nJPEG\nPDF\nZIP\nGIF\nELF\nPE\nBMP\nGZIP\nOGG\nFLAC\nMP3\nICO\nCUR\nMP4\nMIDI\nWAV\nUnknown"))
        func_get_file.file_optional = True
        self.service_info.functions.append(func_get_file)
        
        func_get_stats = pb.StegServiceFunction(name="get_stats", description="Retrieves current info of saved files.")
        func_get_stats.parameter.append(pb.StegServiceParameterDefinition(name="file_limit", description="Will limit the maximal amount of file stats to return.", type=pb.Type.INT, optional=True, default="100"))
        func_get_stats.return_fields.append(pb.StegServiceReturnFieldDefinition(name="stats", label="stats", type=pb.Type.DICT, description="Info of current files."))
        func_get_stats.file_optional = True
        self.service_info.functions.append(func_get_stats)
        
        func_compare_size = pb.StegServiceFunction(name="compare_size", description="Compares the size of two files and returns the difference in size in percent and bytes.")
        func_compare_size.parameter.append(pb.StegServiceParameterDefinition(name="file_b", description="File to compare with.", type=pb.Type.BYTES, optional=False))
        func_compare_size.return_fields.append(pb.StegServiceReturnFieldDefinition(name="size_a", label="size_a", type=pb.Type.INT, description="Size of file_a."))
        func_compare_size.return_fields.append(pb.StegServiceReturnFieldDefinition(name="size_b", label="size_b", type=pb.Type.INT, description="Size of file_a."))
        func_compare_size.return_fields.append(pb.StegServiceReturnFieldDefinition(name="diff_in_bytes", label="diff_in_bytes", type=pb.Type.INT, description="File size difference in bytes."))
        func_compare_size.return_fields.append(pb.StegServiceReturnFieldDefinition(name="diff_in_percentage", label="diff_in_percentage", type=pb.Type.FLOAT, description="File size difference in percentage."))
        self.service_info.functions.append(func_compare_size)
        
        func_get_file_type = pb.StegServiceFunction(name="get_file_type", description="Retrieves the type of a file based on known magic bytes.")
        func_get_file_type.return_fields.append(pb.StegServiceReturnFieldDefinition(name="file_type", label="file_type", type=pb.Type.STRING, description="File type based on magic bytes.\nPNG\nJPEG\nPDF\nZIP\nGIF\nELF\nPE\nBMP\nGZIP\nOGG\nFLAC\nMP3\nICO\nCUR\nMP4\nMIDI\nWAV\nUnknown"))
        self.service_info.functions.append(func_get_file_type)
        
        func_sha256 = pb.StegServiceFunction(name="sha256", description="Calculates the sha-256 hash value of a file.")
        func_sha256.return_fields.append(pb.StegServiceReturnFieldDefinition(name="sha256", label="sha256", type=pb.Type.STRING, description="Sha-256 Hash."))
        self.service_info.functions.append(func_sha256)
        
        func_png_to_jpg = pb.StegServiceFunction(name="png_to_jpg", description="Converts PNG to JPG.")
        func_png_to_jpg.return_fields.append(pb.StegServiceReturnFieldDefinition(name="jpg_img", label="jpg_img", type=pb.Type.BYTES, description="The JPG image."))
        self.service_info.functions.append(func_png_to_jpg)
        
        func_jpg_to_png = pb.StegServiceFunction(name="jpg_to_png", description="Converts JPG to PNG.")
        func_jpg_to_png.return_fields.append(pb.StegServiceReturnFieldDefinition(name="png_img", label="png_img", type=pb.Type.BYTES, description="The PNG image."))
        self.service_info.functions.append(func_jpg_to_png)
        
        func_clear_storage = pb.StegServiceFunction(name="clear_storage", description="Deletes all stores files from storage.")
        func_clear_storage.file_optional = True
        self.service_info.functions.append(func_clear_storage)
        
        func_nop = pb.StegServiceFunction(name="nop", description="No operation function.")
        func_nop.file_optional = True
        func_nop.is_nop = True
        self.service_info.functions.append(func_nop)
        

    def Execute(self, request : pb.StegServiceRequest, context):
        print(f"recieved request from {context.peer()}")

        match request.function:
            case "save_file":
                return self.save_file(request)
            case "get_last_file":
                return self.get_last_file(request)
            case "get_file":
                return self.get_file(request)
            case "get_stats":
                return self.get_stats(request)
            case "get_file_type":
                return self.get_file_type(request)
            case "png_to_jpg":
                return self.png_to_jpg(request)
            case "sha256":
                return self.get_sha256(request)
            case "jpg_to_png":
                return self.jpg_to_png(request)
            case "clear_storage":
                return self.clear_storage()
            case "compare_size":
                return self.compare_size()
            case "nop":
                return pb.StegServiceResponse()
            case _:
                raise Exception(f"Function {request.function} not implemented")
    
    def GetStegServiceInfo(self, request, context):
        print(f"recieved request from {context.peer()}")
        return self.service_info

    def save_file(self, request: pb.StegServiceRequest):
        file_name = util.get_parameter(self.service_info, "file_name", request)
        if not file_name:
            file_name = f"tmp_{uuid.uuid4().hex[:5]}"
        life_span = util.get_parameter(self.service_info, "life_span", request)
        result = self.file_handler.save_file(request.file, file_name, life_span)

        response = pb.StegServiceResponse()
        response.values["file_name"].string_value = file_name
        response.values["expires_at"].string_value = result["expires_at"]
        response.values["file_type"].string_value = result["file_type"]
        response.values["sha256"].string_value = result["sha256"]
        
        result = self.file_handler.get_stats(file_limit=util.get_parameter(self.service_info, "file_limit", request))
        json_compatible_data = util.make_json_compatible(result)
        
        struct_data = Struct()
        struct_data.update(json_compatible_data)
        response.values["stats"].structured_value.struct_value = struct_data
        
        return response
    
    def get_last_file(self, request: pb.StegServiceRequest):
        delete_after = util.get_parameter(self.service_info, "delete_after", request)
        result = self.file_handler.get_last_file(delete_after)
        
        if result:
            response = pb.StegServiceResponse()
            response.values["file_name"].string_value = result["file_name"]
            response.values["file_data"].binary_value = result["file_data"]
            response.values["file_type"].string_value = result["file_type"]
            response.values["expires_at"].string_value = result["expires_at"]
            response.values["sha256"].string_value = result["sha256"]

            result = self.file_handler.get_stats(file_limit=util.get_parameter(self.service_info, "file_limit", request))
            json_compatible_data = util.make_json_compatible(result)
            
            struct_data = Struct()
            struct_data.update(json_compatible_data)
            response.values["stats"].structured_value.struct_value = struct_data
            return response
        else:
            raise Exception("No file found")
    
    def get_file(self, request: pb.StegServiceRequest):
        file_name = util.get_parameter(self.service_info, "file_name", request)
        delete_after = util.get_parameter(self.service_info, "delete_after", request)
        result = self.file_handler.get_file(file_name, delete_after)
        
        if result:
            response = pb.StegServiceResponse()
            response.values["file_name"].string_value = result["file_name"]
            response.values["file_data"].string_value = result["file_data"]
            response.values["file_type"].string_value = result["file_type"]
            response.values["expires_at"].string_value = result["expires_at"]
            response.values["sha256"].string_value = result["sha256"]
            
            result = self.file_handler.get_stats(file_limit=util.get_parameter(self.service_info, "file_limit", request))
            json_compatible_data = util.make_json_compatible(result)
            
            struct_data = Struct()
            struct_data.update(json_compatible_data)
            response.values["stats"].structured_value.struct_value = struct_data
            return response
        else:
            raise Exception(f"File {file_name} not found")   
    
    def get_stats(self, request: pb.StegServiceRequest):
        response = pb.StegServiceResponse()
        result = self.file_handler.get_stats(file_limit=util.get_parameter(self.service_info, "file_limit", request))
        json_compatible_data = util.make_json_compatible(result)

        struct_data = Struct()
        struct_data.update(json_compatible_data)
        response.values["stats"].structured_value.struct_value = struct_data
        return response 
    
    def png_to_jpg(self, request: pb.StegServiceRequest):
        jpg_img = self.image_converter.convert_to_jpg(request.file)
        
        response = pb.StegServiceResponse()
        response.values["jpg_img"].binary_value = jpg_img
        return response 
    
    def get_file_type(self, request: pb.StegServiceRequest):
        file_type = file_utils.detect_file_type(request.file)

        response = pb.StegServiceResponse()
        response.values["file_type"].string_value = file_type
        return response 
    
    def jpg_to_png(self, request: pb.StegServiceRequest):
        png_img = self.image_converter.convert_to_png(request.file)
        
        response = pb.StegServiceResponse()
        response.values["png_img"].binary_value = png_img
        return response 
    
    def get_sha256(self, request: pb.StegServiceRequest):
        sha256 = file_utils.get_sha_256(request.file)
        
        response = pb.StegServiceResponse()
        response.values["sha256"].string_value = sha256
        return response 
    
    def clear_storage(self):
        self.file_handler.delete_all_files()
        return pb.StegServiceResponse()
    
    def compare_size(self, request: pb.StegServiceRequest):
        file_b = util.get_parameter(self.service_info, "file_b", request)
        result = self.file_stats.compare_bytes_size(request.file, file_b)
        
        response = pb.StegServiceResponse()
        response.values["size_a"].int_value = result["size_a"]
        response.values["size_b"].int_value = result["size_b"]
        response.values["diff_in_bytes"].int_value = result["diff_in_bytes"]
        response.values["diff_in_percentage"].float_value = result["diff_in_percentage"]
        return response 

def serve():
    port = os.environ['port']
    if ":" not in port:
        port = f'[::]:{port}'
    max_workers = os.environ.get("max_workers")
    if max_workers == None:
        max_workers = 10
        
    MAX_MESSAGE_SIZE = 40 * 1024 * 1024
    
    server = grpc.server(
        futures.ThreadPoolExecutor(max_workers=max_workers),                         
        options=[
        ("grpc.max_send_message_length", MAX_MESSAGE_SIZE),  
        ("grpc.max_receive_message_length", MAX_MESSAGE_SIZE) 
        ],)
    util_svc = UtilService()
    pbgrpc.add_StegServiceServicer_to_server(util_svc, server)
    server.add_insecure_port(port)
    print(f"{util_svc.service_info.name} service started on port {port}")
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    serve()
