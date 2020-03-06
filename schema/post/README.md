# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [post.proto](#post.proto)
    - [ApplyPost](#post_grpc.ApplyPost)
    - [CreateApplyPostReq](#post_grpc.CreateApplyPostReq)
    - [CreateApplyPostRes](#post_grpc.CreateApplyPostRes)
    - [CreatePostReq](#post_grpc.CreatePostReq)
    - [CreatePostRes](#post_grpc.CreatePostRes)
    - [DeleteApplyPostReq](#post_grpc.DeleteApplyPostReq)
    - [DeleteApplyPostRes](#post_grpc.DeleteApplyPostRes)
    - [DeletePostReq](#post_grpc.DeletePostReq)
    - [DeletePostRes](#post_grpc.DeletePostRes)
    - [GetListPostsReq](#post_grpc.GetListPostsReq)
    - [GetListPostsRes](#post_grpc.GetListPostsRes)
    - [GetPostByIDReq](#post_grpc.GetPostByIDReq)
    - [GetPostByIDRes](#post_grpc.GetPostByIDRes)
    - [Post](#post_grpc.Post)
    - [UpdatePostReq](#post_grpc.UpdatePostReq)
    - [UpdatePostRes](#post_grpc.UpdatePostRes)
  
  
  
    - [PostService](#post_grpc.PostService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="post.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## post.proto



<a name="post_grpc.ApplyPost"></a>

### ApplyPost



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| post_id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="post_grpc.CreateApplyPostReq"></a>

### CreateApplyPostReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| post_id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |






<a name="post_grpc.CreateApplyPostRes"></a>

### CreateApplyPostRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| apply_post | [ApplyPost](#post_grpc.ApplyPost) |  |  |






<a name="post_grpc.CreatePostReq"></a>

### CreatePostReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  |  |
| content | [string](#string) |  |  |
| user_id | [int64](#int64) |  | トークンに含まれていたidを送る |






<a name="post_grpc.CreatePostRes"></a>

### CreatePostRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| post | [Post](#post_grpc.Post) |  |  |






<a name="post_grpc.DeleteApplyPostReq"></a>

### DeleteApplyPostReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |






<a name="post_grpc.DeleteApplyPostRes"></a>

### DeleteApplyPostRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |






<a name="post_grpc.DeletePostReq"></a>

### DeletePostReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |






<a name="post_grpc.DeletePostRes"></a>

### DeletePostRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |






<a name="post_grpc.GetListPostsReq"></a>

### GetListPostsReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| num | [int64](#int64) |  | created_at DESC |






<a name="post_grpc.GetListPostsRes"></a>

### GetListPostsRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| posts | [Post](#post_grpc.Post) | repeated |  |






<a name="post_grpc.GetPostByIDReq"></a>

### GetPostByIDReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="post_grpc.GetPostByIDRes"></a>

### GetPostByIDRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| post | [Post](#post_grpc.Post) |  |  |






<a name="post_grpc.Post"></a>

### Post



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| title | [string](#string) |  |  |
| content | [string](#string) |  |  |
| user_id | [int64](#int64) |  |  |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="post_grpc.UpdatePostReq"></a>

### UpdatePostReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| title | [string](#string) |  |  |
| content | [string](#string) |  |  |
| user_id | [int64](#int64) |  | トークンに含まれていたidを送る |






<a name="post_grpc.UpdatePostRes"></a>

### UpdatePostRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| post | [Post](#post_grpc.Post) |  |  |





 

 

 


<a name="post_grpc.PostService"></a>

### PostService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetPostByID | [GetPostByIDReq](#post_grpc.GetPostByIDReq) | [GetPostByIDRes](#post_grpc.GetPostByIDRes) | post_idで投稿を取得 |
| GetListPosts | [GetListPostsReq](#post_grpc.GetListPostsReq) | [GetListPostsRes](#post_grpc.GetListPostsRes) | 取得件数と日時を指定して投稿を複数取得 |
| CreatePost | [CreatePostReq](#post_grpc.CreatePostReq) | [CreatePostRes](#post_grpc.CreatePostRes) | 投稿を作成 |
| UpdatePost | [UpdatePostReq](#post_grpc.UpdatePostReq) | [UpdatePostRes](#post_grpc.UpdatePostRes) | 投稿を更新 |
| DeletePost | [DeletePostReq](#post_grpc.DeletePostReq) | [DeletePostRes](#post_grpc.DeletePostRes) | 投稿を削除 |
| CreateApplyPost | [CreateApplyPostReq](#post_grpc.CreateApplyPostReq) | [CreateApplyPostRes](#post_grpc.CreateApplyPostRes) |  |
| DeleteApplyPost | [DeleteApplyPostReq](#post_grpc.DeleteApplyPostReq) | [DeleteApplyPostRes](#post_grpc.DeleteApplyPostRes) |  |

 



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

