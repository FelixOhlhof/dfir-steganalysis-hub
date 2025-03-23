import shutil
import os

def copy_and_replace(source_path, destination_path):
    try:
        if os.path.exists(destination_path):
            os.remove(destination_path)
        shutil.copy2(source_path, destination_path)
    except Exception as e:
        print(e)

files = [
    (os.path.join("..", "pb", "python", "steg_analysis_pb2.py"), os.path.join("..", "..", "clients", "Python", "steg_analysis_pb2.py")),
    (os.path.join("..", "pb", "python", "steg_analysis_pb2.pyi"), os.path.join("..", "..", "clients", "Python", "steg_analysis_pb2.pyi")),
    (os.path.join("..", "pb", "python", "steg_analysis_pb2_grpc.py"), os.path.join("..", "..", "clients", "Python", "steg_analysis_pb2_grpc.py")),
    (os.path.join("..", "pb", "python", "steg_service_pb2.py"), os.path.join("..", "..", "clients", "Python", "steg_service_pb2.py")),
    (os.path.join("..", "pb", "python", "steg_service_pb2.pyi"), os.path.join("..", "..", "clients", "Python", "steg_service_pb2.pyi")),
    (os.path.join("..", "pb", "python", "steg_service_pb2_grpc.py"), os.path.join("..", "..", "clients", "Python", "steg_service_pb2_grpc.py")),

    (os.path.join("..", "pb", "python", "steg_service_pb2.py"), os.path.join("..", "..", "steg-services", "aletheia", "steg_service_pb2.py")),
    (os.path.join("..", "pb", "python", "steg_service_pb2.pyi"), os.path.join("..", "..", "steg-services", "aletheia", "steg_service_pb2.pyi")),
    (os.path.join("..", "pb", "python", "steg_service_pb2_grpc.py"), os.path.join("..", "..", "steg-services", "aletheia", "steg_service_pb2_grpc.py")),

    (os.path.join("..", "pb", "python", "steg_service_pb2.py"), os.path.join("..", "..", "steg-services", "extractor", "steg_service_pb2.py")),
    (os.path.join("..", "pb", "python", "steg_service_pb2.pyi"), os.path.join("..", "..", "steg-services", "extractor", "steg_service_pb2.pyi")),
    (os.path.join("..", "pb", "python", "steg_service_pb2_grpc.py"), os.path.join("..", "..", "steg-services", "extractor", "steg_service_pb2_grpc.py")),

    (os.path.join("..", "pb", "python", "steg_service_pb2.py"), os.path.join("..", "..", "steg-services", "util", "steg_service_pb2.py")),
    (os.path.join("..", "pb", "python", "steg_service_pb2.pyi"), os.path.join("..", "..", "steg-services", "util", "steg_service_pb2.pyi")),
    (os.path.join("..", "pb", "python", "steg_service_pb2_grpc.py"), os.path.join("..", "..", "steg-services", "util", "steg_service_pb2_grpc.py")),

    (os.path.join("..", "pb", "go", "pb", "steg_service.pb.go"), os.path.join("..", "..", "steg-services", "vt", "pb", "steg_service.pb.go")),
    (os.path.join("..", "pb", "go", "pb", "steg_service_grpc.pb.go"), os.path.join("..", "..", "steg-services", "vt", "pb", "steg_service_grpc.pb.go")),

    (os.path.join("..", "pb", "go", "pb", "config.pb.go"), os.path.join("..", "..", "rest-gateway", "pb", "config.pb.go")),
    (os.path.join("..", "pb", "go", "pb", "config_grpc.pb.go"), os.path.join("..", "..", "rest-gateway", "pb", "config_grpc.pb.go")),
    (os.path.join("..", "pb", "go", "pb", "steg_analysis.pb.go"), os.path.join("..", "..", "rest-gateway", "pb", "steg_analysis.pb.go")),
    (os.path.join("..", "pb", "go", "pb", "steg_analysis_grpc.pb.go"), os.path.join("..", "..", "rest-gateway", "pb", "steg_analysis_grpc.pb.go")),
    (os.path.join("..", "pb", "go", "pb", "steg_service.pb.go"), os.path.join("..", "..", "rest-gateway", "pb", "steg_service.pb.go")),
    (os.path.join("..", "pb", "go", "pb", "steg_service_grpc.pb.go"), os.path.join("..", "..", "rest-gateway", "pb", "steg_service_grpc.pb.go")),

    (os.path.join("..", "pb", "go", "pb", "config.pb.go"), os.path.join("..", "..", "grpc-gateway", "pb", "config.pb.go")),
    (os.path.join("..", "pb", "go", "pb", "config_grpc.pb.go"), os.path.join("..", "..", "grpc-gateway", "pb", "config_grpc.pb.go")),
    (os.path.join("..", "pb", "go", "pb", "steg_analysis.pb.go"), os.path.join("..", "..", "grpc-gateway", "pb", "steg_analysis.pb.go")),
    (os.path.join("..", "pb", "go", "pb", "steg_analysis_grpc.pb.go"), os.path.join("..", "..", "grpc-gateway", "pb", "steg_analysis_grpc.pb.go")),
    (os.path.join("..", "pb", "go", "pb", "steg_service.pb.go"), os.path.join("..", "..", "grpc-gateway", "pb", "steg_service.pb.go")),
    (os.path.join("..", "pb", "go", "pb", "steg_service_grpc.pb.go"), os.path.join("..", "..", "grpc-gateway", "pb", "steg_service_grpc.pb.go")),

    (os.path.join("..", "pb", "go", "pb", "config.pb.go"), os.path.join("..", "..", "clients", "Demo", "pb", "config.pb.go")),
    (os.path.join("..", "pb", "go", "pb", "config_grpc.pb.go"), os.path.join("..", "..", "clients", "Demo", "pb", "config_grpc.pb.go")),
    (os.path.join("..", "pb", "go", "pb", "steg_analysis.pb.go"), os.path.join("..", "..", "clients", "Demo", "pb", "steg_analysis.pb.go")),
    (os.path.join("..", "pb", "go", "pb", "steg_analysis_grpc.pb.go"), os.path.join("..", "..", "clients", "Demo", "pb", "steg_analysis_grpc.pb.go")),
    (os.path.join("..", "pb", "go", "pb", "steg_service.pb.go"), os.path.join("..", "..", "clients", "Demo", "pb", "steg_service.pb.go")),
    (os.path.join("..", "pb", "go", "pb", "steg_service_grpc.pb.go"), os.path.join("..", "..", "clients", "Demo", "pb", "steg_service_grpc.pb.go")),

    (os.path.join("..", "pb", "go", "pb", "steg_analysis.pb.go"), os.path.join("..", "..", "clients", "Velociraptor", "pb", "steg_analysis.pb.go")),
    (os.path.join("..", "pb", "go", "pb", "steg_analysis_grpc.pb.go"), os.path.join("..", "..", "clients", "Velociraptor", "pb", "steg_analysis_grpc.pb.go")),
    (os.path.join("..", "pb", "go", "pb", "steg_service.pb.go"), os.path.join("..", "..", "clients", "Velociraptor", "pb", "steg_service.pb.go")),
    (os.path.join("..", "pb", "go", "pb", "steg_service_grpc.pb.go"), os.path.join("..", "..", "clients", "Velociraptor", "pb", "steg_service_grpc.pb.go")),

    (os.path.join("..", "pb", "java", "stego", "hub", "grpc", "wrapper", "StegAnalysis.java"), os.path.join("..", "..", "clients", "Autopsy", "grpc-wrapper", "lib", "src", "main", "java", "stego", "hub", "grpc", "wrapper", "StegAnalysis.java")),
    (os.path.join("..", "pb", "java", "stego", "hub", "grpc", "wrapper", "StegAnalysisServiceGrpc.java"), os.path.join("..", "..", "clients", "Autopsy", "grpc-wrapper", "lib", "src", "main", "java", "stego", "hub", "grpc", "wrapper", "StegAnalysisServiceGrpc.java")),
    (os.path.join("..", "pb", "java", "stego", "hub", "grpc", "wrapper", "StegServiceGrpc.java"), os.path.join("..", "..", "clients", "Autopsy", "grpc-wrapper", "lib", "src", "main", "java", "stego", "hub", "grpc", "wrapper", "StegServiceGrpc.java")),
    (os.path.join("..", "pb", "java", "stego", "hub", "grpc", "wrapper", "StegServiceOuterClass.java"), os.path.join("..", "..", "clients", "Autopsy", "grpc-wrapper", "lib", "src", "main", "java", "stego", "hub", "grpc", "wrapper", "StegServiceOuterClass.java")),
]

if __name__ == '__main__':
    for k, v in files:
        copy_and_replace(k, v)