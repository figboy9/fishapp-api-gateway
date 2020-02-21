# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [chat.proto](#chat.proto)
    - [AddMemberReq](#chat_grpc.AddMemberReq)
    - [CreateRoomReq](#chat_grpc.CreateRoomReq)
    - [Member](#chat_grpc.Member)
    - [Message](#chat_grpc.Message)
    - [Room](#chat_grpc.Room)
    - [SendMessageReq](#chat_grpc.SendMessageReq)
    - [StreamMessageReq](#chat_grpc.StreamMessageReq)
  
  
  
    - [ChatService](#chat_grpc.ChatService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="chat.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## chat.proto



<a name="chat_grpc.AddMemberReq"></a>

### AddMemberReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| room_id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |






<a name="chat_grpc.CreateRoomReq"></a>

### CreateRoomReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| post_id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |






<a name="chat_grpc.Member"></a>

### Member



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| room_id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="chat_grpc.Message"></a>

### Message



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| body | [string](#string) |  |  |
| room_id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="chat_grpc.Room"></a>

### Room



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| post_id | [int64](#int64) |  |  |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="chat_grpc.SendMessageReq"></a>

### SendMessageReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| body | [string](#string) |  |  |
| room_id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |






<a name="chat_grpc.StreamMessageReq"></a>

### StreamMessageReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| room_ids | [int64](#int64) | repeated |  |
| user_id | [int64](#int64) |  |  |





 

 

 


<a name="chat_grpc.ChatService"></a>

### ChatService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateChatRoom | [CreateRoomReq](#chat_grpc.CreateRoomReq) | [Room](#chat_grpc.Room) |  |
| AddMember | [AddMemberReq](#chat_grpc.AddMemberReq) | [Member](#chat_grpc.Member) |  |
| SendMessage | [SendMessageReq](#chat_grpc.SendMessageReq) | [Message](#chat_grpc.Message) |  |
| StreamMessage | [StreamMessageReq](#chat_grpc.StreamMessageReq) | [Message](#chat_grpc.Message) stream |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

