from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Query(_message.Message):
    __slots__ = ("unoms", "startDate", "endDate", "threshold")
    UNOMS_FIELD_NUMBER: _ClassVar[int]
    STARTDATE_FIELD_NUMBER: _ClassVar[int]
    ENDDATE_FIELD_NUMBER: _ClassVar[int]
    THRESHOLD_FIELD_NUMBER: _ClassVar[int]
    unoms: _containers.RepeatedScalarFieldContainer[int]
    startDate: _timestamp_pb2.Timestamp
    endDate: _timestamp_pb2.Timestamp
    threshold: float
    def __init__(self, unoms: _Optional[_Iterable[int]] = ..., startDate: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., endDate: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., threshold: _Optional[float] = ...) -> None: ...

class Response(_message.Message):
    __slots__ = ("predictions",)
    PREDICTIONS_FIELD_NUMBER: _ClassVar[int]
    predictions: _containers.RepeatedCompositeFieldContainer[Prediction]
    def __init__(self, predictions: _Optional[_Iterable[_Union[Prediction, _Mapping]]] = ...) -> None: ...

class Prediction(_message.Message):
    __slots__ = ("unom", "date", "p1", "p2", "t1", "t2", "no", "noHeating", "leak", "strongLeak", "tempLow", "tempLowCommon", "leakSystem")
    UNOM_FIELD_NUMBER: _ClassVar[int]
    DATE_FIELD_NUMBER: _ClassVar[int]
    P1_FIELD_NUMBER: _ClassVar[int]
    P2_FIELD_NUMBER: _ClassVar[int]
    T1_FIELD_NUMBER: _ClassVar[int]
    T2_FIELD_NUMBER: _ClassVar[int]
    NO_FIELD_NUMBER: _ClassVar[int]
    NOHEATING_FIELD_NUMBER: _ClassVar[int]
    LEAK_FIELD_NUMBER: _ClassVar[int]
    STRONGLEAK_FIELD_NUMBER: _ClassVar[int]
    TEMPLOW_FIELD_NUMBER: _ClassVar[int]
    TEMPLOWCOMMON_FIELD_NUMBER: _ClassVar[int]
    LEAKSYSTEM_FIELD_NUMBER: _ClassVar[int]
    unom: int
    date: _timestamp_pb2.Timestamp
    p1: float
    p2: float
    t1: float
    t2: float
    no: float
    noHeating: float
    leak: float
    strongLeak: float
    tempLow: float
    tempLowCommon: float
    leakSystem: float
    def __init__(self, unom: _Optional[int] = ..., date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., p1: _Optional[float] = ..., p2: _Optional[float] = ..., t1: _Optional[float] = ..., t2: _Optional[float] = ..., no: _Optional[float] = ..., noHeating: _Optional[float] = ..., leak: _Optional[float] = ..., strongLeak: _Optional[float] = ..., tempLow: _Optional[float] = ..., tempLowCommon: _Optional[float] = ..., leakSystem: _Optional[float] = ...) -> None: ...
