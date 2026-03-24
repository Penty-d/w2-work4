namespace go common

// BaseResponse
struct BaseResponse{
    1: string code
    2: string msg
}

// User

struct User{
    1: string id
    2: string username
    3: string password
    4: string avatar_url
    5: string create_at
    6: string update_at
    7: string delete_at
    8: string mfa_secret // MFA 密钥，如果开启 MFA 则不为空
}

struct UserDataResponse{
    1: string id
    2: string username
    3: string avatar_url
    4: string create_at
    5: string update_at
    6: string delete_at
}

// MFA
struct QrcodeMFAAuthResponse{
    1: string secret
    2: string qrcode
}

// image

struct Image{
    1: string id // 唯一标识符
    2: string user_id // 上传用户的ID
    3: string url // 图片的URL地址
    4: string original_filename // 原始文件名
    5: string filepath // 文件存储路径
    6: i64 filesize // 文件大小
    7: string mime_type // 文件的MIME类型
    8: string created_at // 创建时间
    9: string updated_at // 更新时间
    10: string deleted_at // 软删除标记
}

// Video
struct VideoItems{
    1: string id
    2: string user_id
    3: string video_url
    4: string cover_url
    5: string title
    6: string description
    7: i64 visit_count
    8: i64 like_count
    9: i64 comment_count
    10: string create_at
    11: string update_at
    12: string delete_at
}

struct VideoDataForPopularResponse{
    1: list<VideoItems> items
}

struct VideoDataForListResponse{
    1: list<VideoItems> items
    2: i64 total
}

// interact
struct LikeItems{
    1: string id // 这里指的是点赞记录id
    2: string user_id // 这里指的是点赞用户
    
    3: string LikeableID   // 被点赞对象视频ID或评论ID
	4: string LikeableType   // 被点赞对象的类型 "video" 或 "comment"

    5: string create_at
    6: string update_at
    7: string delete_at
}

struct LikeVideoDTO { 
	1: string ID           
	2: string UserID       
	3: string VideoURL     
	4: string CoverURL     
	5: string Title        
	6: string Description
	7: i64 VisitCount
	8: i64 LikeCount   
	9: i64 CommentCount  
	10: string CreatedAt
	11: string UpdatedAt
	12: string DeletedAt   
}

struct LikeListResponse{
    1: list<LikeVideoDTO> items
}

struct CommentItems{
    1: string id
    2: string user_id
    3: string video_id
    4: string parent_id
    5: i64 like_count
    6: i64 child_count
    7: string content 
    8: string created_at
    9: string updated_at
    10: string deleted_at
}

struct CommentDataForListResponse{
    1: list<CommentItems> items
}

// social

struct SocialItems{
    1: string id // 这里指的是关注关系id
    2: string follower_id // 关注者的用户ID
    3: string followed_id // 被关注者的用户ID
    4: string created_at // 创建时间
    5: string updated_at // 更新时间
    6: string deleted_at // 软删除标记
}

struct SocialDTO{
    1: string id
    2: string username
    3: string avatar_url
}

struct SocialDataForListResponse{
    1: list<SocialDTO> items
    2: i64 total
}