// Code generated by goctl. DO NOT EDIT.
package types

type UserInfoReq struct {
	UserId int64 `json:"userId"`
}

type UserInfoResp struct {
	UserId   int64  `json:"userId"`
	NickName string `json:"nickname"`
}
