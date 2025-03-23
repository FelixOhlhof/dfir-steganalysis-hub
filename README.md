## DFIR Steg Hub

### Installation using Docker

1. docker-compose build
2. docker-compose up -d

### Configure the Workflow

1. Go to localhost (port 80 is default for the web server)
2. Edit the workflow in the editor or upload an existing yaml file
3. Test the workflow
   - Click choose file
   - Click execute

### Architecture

The Architecture consists out of clients, gateways, a webserver and multiple services. Clients can send requests via the grpc or the rest gateway to execute the whole workflow or a single service.

![alt text](image.png)
Clients can eather send a StegAnalysisRequest (see [steg_analysis.proto](common/proto/steg_analysis.proto) ) to the grpc gateway or a get request to the rest gateway.

### Using the Yara Client

In order to use Yara as a client you must create a custom module, follow the insturctions in the yara docs. Copy [stego.c](clients/Yara/stego.c) to the libyara project, register the class and compile the whole solution. In the future there will be precompiled binaries for demo purposes.

### Using the Velociraptor Client

Import the [Stego Artifact](clients/Velociraptor/Stego.yaml) and upload the [GRPC Connector](clients/Velociraptor/grpc_client.go) in the Tools section.
Collect the artefact on a single client or as a hunt and specify the requiered parameters such as gateway address and directory or file.

### Using the Autopsy Client

Install the [Steganalysis Plugin](clients/Autopsy/Steganalysis/) in Autopsy. Run the Plugin, the Results will be shown in the Blackboard.

### Implementing new Clients

#### GRPC Client

Protobuf files for python, cpp, go and java are already generated and can be copied to your client [common/pb/](common/pb/). Then you can send a StegAnalysisRequest to the grpc endpoint to execute the workflow.

#### HTTP Client

Send a GET request to the rest gateway:

- Address: /execute
- Body: form-data (needs a file key which contains the file to analyze)

### Implementing new Services

Import the proto files in your desiered language to your service. Implement the requierd server methods (Execute, GetStegServiceInfo). Register the service in the grpc gateway by adding the endpoint to the services env (if you are using docker you can add is directly in docker-compose.yaml). Restart the grpc server.
