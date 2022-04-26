package api

// 用户
type ReqInUserInfo struct {
  Id   uint64 `json:"id"`
  Name string `json:"name"`
}