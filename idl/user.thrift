namespace go user

struct User {
  1: i64 id,
  2: string name,
  3: i64 follow_count,
  4: i64 follower_count,
  5: bool is_follow,
  6: string avatar,
  7: string background_image,
  8: string signature,
  9: i64 total_favorited,
  10: i64 work_count,
  11: i64 favorite_count
}

struct douyin_user_register_request {
  1: string username; // 注册用户名，最长32个字符
  2: string password; // 密码，最长32个字符
}

struct douyin_user_register_response {
  1: i32 status_code; // 状态码，0-成功，其他值-失败
  2: string status_msg; // 返回状态描述
  3: i64 user_id; // 用户id
  4: string token; // 用户鉴权token
}

struct douyin_user_login_request {
  1: string username; // 登录用户名
  2: string password; // 登录密码
}

struct douyin_user_login_response {
  1: i32 status_code; // 状态码，0-成功，其他值-失败
  2: string status_msg; // 返回状态描述
  3: i64 user_id; // 用户id
  4: string token; // 用户鉴权token
}

struct douyin_user_request {
  1: i64 user_id; // 用户id
  2: string token; // 用户鉴权token
}

struct douyin_user_response {
  1: i32 status_code // 状态码，0-成功，其他值-失败
  2: string status_msg; // 返回状态描述
  3: User user; // 用户信息
}

service UserService {
    douyin_user_register_response Register(1: douyin_user_register_request request)
    douyin_user_login_response Login(1: douyin_user_login_request request)
    douyin_user_response GetUserById(1: douyin_user_request request)
}
