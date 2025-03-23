import steg_service_pb2 as _steg_service_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Status(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    SUCCESS: _ClassVar[Status]
    OUT_OF_CONDITION: _ClassVar[Status]
    GRPC_ERROR: _ClassVar[Status]
    EXCEEDED_TIMEOUT: _ClassVar[Status]
    FILE_TYPE_NOT_ALLOWED: _ClassVar[Status]
    EXCEEDED_FILESIZE: _ClassVar[Status]
    SERVICE_ERROR: _ClassVar[Status]
    CLIENT_ERROR: _ClassVar[Status]
SUCCESS: Status
OUT_OF_CONDITION: Status
GRPC_ERROR: Status
EXCEEDED_TIMEOUT: Status
FILE_TYPE_NOT_ALLOWED: Status
EXCEEDED_FILESIZE: Status
SERVICE_ERROR: Status
CLIENT_ERROR: Status

class StegAnalysisRequest(_message.Message):
    __slots__ = ("file", "params", "exec", "file_name")
    class ParamsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    FILE_FIELD_NUMBER: _ClassVar[int]
    PARAMS_FIELD_NUMBER: _ClassVar[int]
    EXEC_FIELD_NUMBER: _ClassVar[int]
    FILE_NAME_FIELD_NUMBER: _ClassVar[int]
    file: bytes
    params: _containers.ScalarMap[str, str]
    exec: str
    file_name: str
    def __init__(self, file: _Optional[bytes] = ..., params: _Optional[_Mapping[str, str]] = ..., exec: _Optional[str] = ..., file_name: _Optional[str] = ...) -> None: ...

class StegAnalysisResponse(_message.Message):
    __slots__ = ("task_results", "error", "duration_ms", "sha256")
    TASK_RESULTS_FIELD_NUMBER: _ClassVar[int]
    ERROR_FIELD_NUMBER: _ClassVar[int]
    DURATION_MS_FIELD_NUMBER: _ClassVar[int]
    SHA256_FIELD_NUMBER: _ClassVar[int]
    task_results: _containers.RepeatedCompositeFieldContainer[TaskResult]
    error: str
    duration_ms: int
    sha256: str
    def __init__(self, task_results: _Optional[_Iterable[_Union[TaskResult, _Mapping]]] = ..., error: _Optional[str] = ..., duration_ms: _Optional[int] = ..., sha256: _Optional[str] = ...) -> None: ...

class TaskResult(_message.Message):
    __slots__ = ("task_id", "service_name", "function_name", "service_response", "error", "status", "duration_ms")
    TASK_ID_FIELD_NUMBER: _ClassVar[int]
    SERVICE_NAME_FIELD_NUMBER: _ClassVar[int]
    FUNCTION_NAME_FIELD_NUMBER: _ClassVar[int]
    SERVICE_RESPONSE_FIELD_NUMBER: _ClassVar[int]
    ERROR_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    DURATION_MS_FIELD_NUMBER: _ClassVar[int]
    task_id: str
    service_name: str
    function_name: str
    service_response: _steg_service_pb2.StegServiceResponse
    error: str
    status: str
    duration_ms: int
    def __init__(self, task_id: _Optional[str] = ..., service_name: _Optional[str] = ..., function_name: _Optional[str] = ..., service_response: _Optional[_Union[_steg_service_pb2.StegServiceResponse, _Mapping]] = ..., error: _Optional[str] = ..., status: _Optional[str] = ..., duration_ms: _Optional[int] = ...) -> None: ...
