include "common.thrift" 

namespace go user

struct RegisterUserRequest{
    1: string username (api.form="username",api.vd="(len($) > 0 && len($) < 100)")
    2: string password (api.form="password", api.vd="(len($) >= 6)")
}

struct RegisterUserResponse{
    1: common.BaseResponse base
}

struct LoginUserRequest{
    1: string username (api.form="username", api.vd="(len($) > 0 && len($) < 100)")
    2: string password (api.form="password", api.vd="(len($) >= 6)")
    3: optional string code (api.form="code")
}

struct LoginUserResponse{
    1: common.BaseResponse base
    2: common.UserDataResponse data
}

struct InfoUserRequest{
    1: optional string user_id (api.query="user_id", api.vd="(len($) == 0) || (len($) > 0 && len($) < 100)")
}

struct InfoUserResponse{
    1: common.BaseResponse base
    2: common.UserDataResponse data
}

struct AvatarUploadUserRequest{
    // 1: optional file data (api.form="data")
}

struct AvatarUploadUserResponse{
    1: common.BaseResponse base
    2: common.UserDataResponse data
}

struct QrcodeMFAAuthRequest{
    // No parameters
    // 开启 MFA 时使用
}

struct QrcodeMFAAuthResponse{
    1: common.BaseResponse base
    2: common.QrcodeMFAAuthResponse data
}

struct BindMFAAuthRequest{
    1: optional string code (api.form="code", api.vd="len($) == 6")
    2: optional string secret (api.form="secret", api.vd="len($) == 0) || len($) > 0")
}

struct BindMFAAuthResponse{
    1: common.BaseResponse base
}

struct SearchImageRequest{
    // 1: optional file data (api.form="data")
}

struct SearchImageResponse{
    1: common.BaseResponse base
    2: string data // 搜索结果，为单个url
}

service UserService {
   RegisterUserResponse RegisterUser(1:RegisterUserRequest req)(api.post="/v1/user/register")
   LoginUserResponse LoginUser(1:LoginUserRequest req)(api.post="/v1/user/login")
   InfoUserResponse InfoUser(1:InfoUserRequest req)(api.get="/v1/user/info")
   AvatarUploadUserResponse AvatarUploadUser(1:AvatarUploadUserRequest req)(api.put="/v1/user/avatar/upload")
   QrcodeMFAAuthResponse QrcodeMFAAuth(1:QrcodeMFAAuthRequest req)(api.get="/v1/auth/mfa/qrcode")
   BindMFAAuthResponse BindMFAAuth(1:BindMFAAuthRequest req)(api.post="/v1/auth/mfa/bind")
   SearchImageResponse SearchImage(1:SearchImageRequest req)(api.post="/v1/user/image/search")
}