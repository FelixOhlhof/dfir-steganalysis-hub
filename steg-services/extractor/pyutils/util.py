import steg_service_pb2 as pb



def get_parameter(steg_service_info: pb.StegServiceInfo, param_name: str, request: pb.StegServiceRequest):
    try:
        params = next(f.parameter for f in steg_service_info.functions if f.name == request.function)
    except StopIteration:
        raise Exception(f"no function '{request.function}' found in StegServiceInfo")

    try:
        param: pb.StegServiceParameterDefinition = next(p for p in params if p.name == param_name)
    except StopIteration:
        raise Exception(f"no parameter '{param_name}' found for function '{request.function}'")

    raw_param = request.params.get(param_name)

    if raw_param is None:
        if param.default:
            match param.type:
                case pb.Type.STRING:
                    return str(param.default)
                case pb.Type.INT:
                    return int(param.default)
                case pb.Type.BOOL:
                    return param.default.lower() == "true"
                case pb.Type.FLOAT:
                    return float(param.default)
                case _:
                    raise Exception(f"{param.type} not a valid type for default parameter")
        if param.optional:
            return None
        raise Exception(f"no value for mandatory parameter '{param_name}' specified")

    req_param = None

    match param.type:
        case pb.Type.INT:
            req_param = raw_param.int_value or int(raw_param.string_value or 0)
        case pb.Type.STRING:
            req_param = raw_param.string_value
        case pb.Type.BOOL:
            req_param = raw_param.bool_value or (raw_param.string_value.lower() == "true" if raw_param.string_value else False)
        case pb.Type.FLOAT:
            try:
                req_param = float(raw_param.float_value or raw_param.string_value)
            except ValueError:
                raise Exception(f"invalid float value for parameter '{param_name}'")
        case pb.Type.BYTES:
            req_param = raw_param.binary_value or raw_param.string_value.encode()
        case _:
            raise Exception(f"type of parameter '{param_name}' not supported")

    if req_param is None and not param.optional:
        raise Exception(f"no value for mandatory parameter '{param_name}' specified")

    return req_param

def make_json_compatible(data):
    if isinstance(data, dict):
        return {key: make_json_compatible(value) for key, value in data.items()}
    elif isinstance(data, list):
        return [make_json_compatible(item) for item in data]
    elif isinstance(data, (str, int, float, bool)) or data is None:
        return data
    else:
        return str(data)
    
def parse_dict(data_dict : dict) -> pb.StegServiceResponse:
    response = pb.StegServiceResponse()

    for key, value in data_dict.items():        
        if isinstance(value, str):
            response.values[key].string_value = value
        elif isinstance(value, bool):
            response.values[key].bool_value = value
        elif isinstance(value, int):
            response.values[key].int_value = value
        elif isinstance(value, float):
            response.values[key].float_value = value
        elif isinstance(value, bytes):
            response.values[key].binary_value = value
        else:
            return pb.StegServiceResponse(error=f"Unsupported type: {type(value)} for key: {key}")
    return response