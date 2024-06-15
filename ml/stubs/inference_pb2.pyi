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
    __slots__ = ("unom", "prediction_id")
    UNOM_FIELD_NUMBER: _ClassVar[int]
    PREDICTION_ID_FIELD_NUMBER: _ClassVar[int]
    unom: int
    prediction_id: int
    def __init__(self, unom: _Optional[int] = ..., prediction_id: _Optional[int] = ...) -> None: ...
