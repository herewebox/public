from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class regionBase(_message.Message):
    __slots__ = ("id", "name")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ...) -> None: ...

class ListRegionRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ListRegionResponse(_message.Message):
    __slots__ = ("regions",)
    REGIONS_FIELD_NUMBER: _ClassVar[int]
    regions: _containers.RepeatedCompositeFieldContainer[regionBase]
    def __init__(self, regions: _Optional[_Iterable[_Union[regionBase, _Mapping]]] = ...) -> None: ...

class GetRegionRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class GetRegionResponse(_message.Message):
    __slots__ = ("region",)
    REGION_FIELD_NUMBER: _ClassVar[int]
    region: regionBase
    def __init__(self, region: _Optional[_Union[regionBase, _Mapping]] = ...) -> None: ...
