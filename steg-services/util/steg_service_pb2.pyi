from google.protobuf import empty_pb2 as _empty_pb2
from google.protobuf import struct_pb2 as _struct_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Type(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    STRING: _ClassVar[Type]
    INT: _ClassVar[Type]
    FLOAT: _ClassVar[Type]
    BOOL: _ClassVar[Type]
    DICT: _ClassVar[Type]
    LIST: _ClassVar[Type]
    BYTES: _ClassVar[Type]
STRING: Type
INT: Type
FLOAT: Type
BOOL: Type
DICT: Type
LIST: Type
BYTES: Type

class StegServiceRequest(_message.Message):
    __slots__ = ("file", "function", "params", "request_timeout_sec")
    class ParamsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: StegServiceRequestParameterValue
        def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[StegServiceRequestParameterValue, _Mapping]] = ...) -> None: ...
    FILE_FIELD_NUMBER: _ClassVar[int]
    FUNCTION_FIELD_NUMBER: _ClassVar[int]
    PARAMS_FIELD_NUMBER: _ClassVar[int]
    REQUEST_TIMEOUT_SEC_FIELD_NUMBER: _ClassVar[int]
    file: bytes
    function: str
    params: _containers.MessageMap[str, StegServiceRequestParameterValue]
    request_timeout_sec: int
    def __init__(self, file: _Optional[bytes] = ..., function: _Optional[str] = ..., params: _Optional[_Mapping[str, StegServiceRequestParameterValue]] = ..., request_timeout_sec: _Optional[int] = ...) -> None: ...

class StegServiceRequestParameterValue(_message.Message):
    __slots__ = ("string_value", "int_value", "float_value", "bool_value", "binary_value")
    STRING_VALUE_FIELD_NUMBER: _ClassVar[int]
    INT_VALUE_FIELD_NUMBER: _ClassVar[int]
    FLOAT_VALUE_FIELD_NUMBER: _ClassVar[int]
    BOOL_VALUE_FIELD_NUMBER: _ClassVar[int]
    BINARY_VALUE_FIELD_NUMBER: _ClassVar[int]
    string_value: str
    int_value: int
    float_value: float
    bool_value: bool
    binary_value: bytes
    def __init__(self, string_value: _Optional[str] = ..., int_value: _Optional[int] = ..., float_value: _Optional[float] = ..., bool_value: bool = ..., binary_value: _Optional[bytes] = ...) -> None: ...

class ResponseValue(_message.Message):
    __slots__ = ("string_value", "int_value", "float_value", "bool_value", "binary_value", "structured_value")
    STRING_VALUE_FIELD_NUMBER: _ClassVar[int]
    INT_VALUE_FIELD_NUMBER: _ClassVar[int]
    FLOAT_VALUE_FIELD_NUMBER: _ClassVar[int]
    BOOL_VALUE_FIELD_NUMBER: _ClassVar[int]
    BINARY_VALUE_FIELD_NUMBER: _ClassVar[int]
    STRUCTURED_VALUE_FIELD_NUMBER: _ClassVar[int]
    string_value: str
    int_value: int
    float_value: float
    bool_value: bool
    binary_value: bytes
    structured_value: _struct_pb2.Value
    def __init__(self, string_value: _Optional[str] = ..., int_value: _Optional[int] = ..., float_value: _Optional[float] = ..., bool_value: bool = ..., binary_value: _Optional[bytes] = ..., structured_value: _Optional[_Union[_struct_pb2.Value, _Mapping]] = ...) -> None: ...

class StegServiceResponse(_message.Message):
    __slots__ = ("values", "error")
    class ValuesEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: ResponseValue
        def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[ResponseValue, _Mapping]] = ...) -> None: ...
    VALUES_FIELD_NUMBER: _ClassVar[int]
    ERROR_FIELD_NUMBER: _ClassVar[int]
    values: _containers.MessageMap[str, ResponseValue]
    error: str
    def __init__(self, values: _Optional[_Mapping[str, ResponseValue]] = ..., error: _Optional[str] = ...) -> None: ...

class StegServiceReturnFieldDefinition(_message.Message):
    __slots__ = ("name", "label", "type", "description", "isIterable")
    NAME_FIELD_NUMBER: _ClassVar[int]
    LABEL_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    ISITERABLE_FIELD_NUMBER: _ClassVar[int]
    name: str
    label: str
    type: Type
    description: str
    isIterable: bool
    def __init__(self, name: _Optional[str] = ..., label: _Optional[str] = ..., type: _Optional[_Union[Type, str]] = ..., description: _Optional[str] = ..., isIterable: bool = ...) -> None: ...

class StegServiceParameterDefinition(_message.Message):
    __slots__ = ("name", "type", "default", "description", "optional")
    NAME_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    DEFAULT_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    OPTIONAL_FIELD_NUMBER: _ClassVar[int]
    name: str
    type: Type
    default: str
    description: str
    optional: bool
    def __init__(self, name: _Optional[str] = ..., type: _Optional[_Union[Type, str]] = ..., default: _Optional[str] = ..., description: _Optional[str] = ..., optional: bool = ...) -> None: ...

class StegServiceFunction(_message.Message):
    __slots__ = ("name", "description", "parameter", "return_fields", "supported_file_types", "file_optional", "is_nop")
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    PARAMETER_FIELD_NUMBER: _ClassVar[int]
    RETURN_FIELDS_FIELD_NUMBER: _ClassVar[int]
    SUPPORTED_FILE_TYPES_FIELD_NUMBER: _ClassVar[int]
    FILE_OPTIONAL_FIELD_NUMBER: _ClassVar[int]
    IS_NOP_FIELD_NUMBER: _ClassVar[int]
    name: str
    description: str
    parameter: _containers.RepeatedCompositeFieldContainer[StegServiceParameterDefinition]
    return_fields: _containers.RepeatedCompositeFieldContainer[StegServiceReturnFieldDefinition]
    supported_file_types: _containers.RepeatedScalarFieldContainer[str]
    file_optional: bool
    is_nop: bool
    def __init__(self, name: _Optional[str] = ..., description: _Optional[str] = ..., parameter: _Optional[_Iterable[_Union[StegServiceParameterDefinition, _Mapping]]] = ..., return_fields: _Optional[_Iterable[_Union[StegServiceReturnFieldDefinition, _Mapping]]] = ..., supported_file_types: _Optional[_Iterable[str]] = ..., file_optional: bool = ..., is_nop: bool = ...) -> None: ...

class StegServiceInfo(_message.Message):
    __slots__ = ("name", "description", "functions")
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    FUNCTIONS_FIELD_NUMBER: _ClassVar[int]
    name: str
    description: str
    functions: _containers.RepeatedCompositeFieldContainer[StegServiceFunction]
    def __init__(self, name: _Optional[str] = ..., description: _Optional[str] = ..., functions: _Optional[_Iterable[_Union[StegServiceFunction, _Mapping]]] = ...) -> None: ...
