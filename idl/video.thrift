include "common.thrift" 

namespace go video

struct FeedVideoRequest {
    1: optional string latest_time (api.query="latest_time", api.vd="(len($) == 0) || (len($) > 0 && len($) < 100)")
}

struct FeedVideoResponse {
    1: common.BaseResponse base
    2: common.VideoDataForPopularResponse data
}

struct PublishVideoRequest{   
    // 1: optional file data (api.form="data")
    2: optional string title (api.form="title", api.vd="(len($) == 0) || (len($) > 0 && len($) < 100)")
    3: optional string description (api.form="description", api.vd="((len($) == 0) || (len($) > 0 && len($) < 100))")
}

struct PublishVideoResponse{
    1: common.BaseResponse base
}

struct ListVideoRequest{
    1: string user_id (api.query="user_id", api.vd="(len($) > 0 && len($) < 100)")
    2: i64 page_num (api.query="page_num", api.vd="$>=0")
    3: i64 page_size (api.query="page_size", api.vd="$>=0")
}

struct ListVideoResponse{
    1: common.BaseResponse base
    2: common.VideoDataForListResponse data
}

struct PopularVideoRequest{
    1: optional i64 page_num (api.query="page_num", api.vd="(len($) == 0) || ($>=0)")
    2: optional i64 page_size (api.query="page_size", api.vd="(len($) == 0) || ($>=0)")
}

struct PopularVideoResponse{
    1: common.BaseResponse base
    2: common.VideoDataForPopularResponse data
}

struct SearchVideoRequest{
    1: string keyword (api.form="keyword", api.vd="(len($) >= 0 && len($) < 100)")
    2: i64 page_num (api.form="page_num", api.vd="$>=0")
    3: i64 page_size (api.form="page_size", api.vd="$>=0")
    4: optional i64 from_date (api.form="from_date", api.vd="(len($) == 0) || ($>=0)")
    5: optional i64 to_date (api.form="to_date", api.vd="(len($) == 0) || ($>=0)")
    6: optional string username (api.form="username", api.vd="(len($) == 0) || (len($) > 0 && len($) < 100)")
}

struct SearchVideoResponse{
    1: common.BaseResponse base
    2: common.VideoDataForListResponse data
}

service VideoService {
    FeedVideoResponse FeedVideo(1: FeedVideoRequest req) (api.get="/v1/video/feed/")
    PublishVideoResponse PublishVideo(1: PublishVideoRequest req) (api.post="/v1/video/publish/")
    ListVideoResponse ListVideo(1: ListVideoRequest req) (api.get="/v1/video/list/")
    PopularVideoResponse PopularVideo(1: PopularVideoRequest req) (api.get="/v1/video/popular/")
    SearchVideoResponse SearchVideo(1: SearchVideoRequest req) (api.post="/v1/video/search/")
}