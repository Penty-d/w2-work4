include "common.thrift" 

namespace go interact

// 点赞操作
// 点赞是1，取消点赞是2

enum ActionTypeLike{
    LIKE = 1,
    UNLIKE = 2,
}

struct ActionLikeRequest{
    1: optional string video_id (api.form="video_id", api.vd="(len($)==0 || len($) > 0 && len($) < 100)")
    2: optional string comment_id (api.form="comment_id", api.vd="(len($)==0) || (len($) > 0 && len($) < 100)")
    3: optional ActionTypeLike action_type (api.form="action_type", api.vd="$ in [1, 2]")
}

struct ActionLikeResponse{
    1: common.BaseResponse base
}

struct ListLikeRequest{
    1: optional string user_id (api.query="user_id", api.vd="(len($)==0) || (len($) > 0 && len($) < 100)")
    2: optional i64 page_size (api.query="page_size", api.vd="(len($) == 0) || ( $ > 0 && $ < 100 )")
    3: optional i64 page_num (api.query="page_num", api.vd="(len($) == 0) || ( $ > 0 )")
}

struct ListLikeResponse{
    1: common.BaseResponse base
    2: common.LikeListResponse data
}

// 评论操作

struct PublishCommentRequest{
    1: optional string video_id (api.form="video_id", api.vd="(len($)==0) || (len($) > 0 && len($) < 100)")
    2: optional string comment_id (api.form="comment_id", api.vd="(len($)==0) || (len($) > 0 && len($) < 100)")
    3: optional string content (api.form="content", api.vd="(len($)==0) || (len($) > 0 && len($) < 100)")
}

struct PublishCommentResponse{
    1: common.BaseResponse base
}

struct ListCommentRequest{
    1: optional string video_id (api.query="video_id", api.vd="(len($)==0) || (len($) > 0 && len($) < 100)")
    2: optional string comment_id (api.query="comment_id", api.vd="(len($)==0) || (len($) > 0 && len($) < 100)")
    3: optional i64 page_size (api.query="page_size", api.vd="(len($) == 0) || ( $ > 0 && $ < 100 )")
    4: optional i64 page_num (api.query="page_num", api.vd="(len($) == 0) || ( $ > 0 )")
}

struct ListCommentResponse{
    1: common.BaseResponse base
    2: common.CommentDataForListResponse data
}

struct DeleteCommentRequest{
    1: optional string video_id (api.form="video_id", api.vd="(len($)==0) || (len($) > 0 && len($) < 100)")
    2: optional string comment_id (api.form="comment_id", api.vd="(len($)==0) || (len($) > 0 && len($) < 100)")
}

struct DeleteCommentResponse{
    1: common.BaseResponse base
}

service InteractService {
   ActionLikeResponse ActionLike(1:ActionLikeRequest req)(api.post="/v1/like/action")
   ListLikeResponse ListLike(1:ListLikeRequest req)(api.get="/v1/like/list")
   PublishCommentResponse PublishComment(1:PublishCommentRequest req)(api.post="/v1/comment/publish")
   ListCommentResponse ListComment(1:ListCommentRequest req)(api.get="/v1/comment/list")
   DeleteCommentResponse DeleteComment(1:DeleteCommentRequest req)(api.delete="/v1/comment/delete")
}