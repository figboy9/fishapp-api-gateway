# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [auth.proto](#auth.proto)
    - [CreateReq](#auth_grpc.CreateReq)
    - [ID](#auth_grpc.ID)
    - [LoginReq](#auth_grpc.LoginReq)
    - [TokenPair](#auth_grpc.TokenPair)
    - [UpdateReq](#auth_grpc.UpdateReq)
    - [User](#auth_grpc.User)
    - [UserWithToken](#auth_grpc.UserWithToken)
  
  
  
    - [AuthService](#auth_grpc.AuthService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="auth.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## auth.proto



<a name="auth_grpc.CreateReq"></a>

### CreateReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  |  |
| password | [string](#string) |  | 6文字以上72文字以下の英数字 |






<a name="auth_grpc.ID"></a>

### ID



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="auth_grpc.LoginReq"></a>

### LoginReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  |  |
| password | [string](#string) |  | 6文字以上72文字以下の英数字 |






<a name="auth_grpc.TokenPair"></a>

### TokenPair



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id_token | [string](#string) |  |  |
| refresh_token | [string](#string) |  |  |






<a name="auth_grpc.UpdateReq"></a>

### UpdateReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  |  |
| password | [string](#string) |  | 6文字以上72文字以下の英数字 |






<a name="auth_grpc.User"></a>

### User



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| email | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="auth_grpc.UserWithToken"></a>

### UserWithToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#auth_grpc.User) |  |  |
| token_pair | [TokenPair](#auth_grpc.TokenPair) |  |  |





 

 

 


<a name="auth_grpc.AuthService"></a>

### AuthService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateReq](#auth_grpc.CreateReq) | [UserWithToken](#auth_grpc.UserWithToken) |  |
| GetByID | [ID](#auth_grpc.ID) | [User](#auth_grpc.User) |  |
| Update | [UpdateReq](#auth_grpc.UpdateReq) | [User](#auth_grpc.User) |  |
| Delete | [.google.protobuf.Empty](#google.protobuf.Empty) | [.google.protobuf.BoolValue](#google.protobuf.BoolValue) |  |
| Login | [LoginReq](#auth_grpc.LoginReq) | [UserWithToken](#auth_grpc.UserWithToken) |  |
| Logout | [.google.protobuf.Empty](#google.protobuf.Empty) | [.google.protobuf.BoolValue](#google.protobuf.BoolValue) |  |
| RefreshIDToken | [.google.protobuf.Empty](#google.protobuf.Empty) | [TokenPair](#auth_grpc.TokenPair) |  |

 



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

