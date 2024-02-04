from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class InitDBRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class InitDBResponse(_message.Message):
    __slots__ = ("initialized",)
    INITIALIZED_FIELD_NUMBER: _ClassVar[int]
    initialized: bool
    def __init__(self, initialized: bool = ...) -> None: ...
