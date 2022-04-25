package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserNicknameModel = (*customUserNicknameModel)(nil)

type (
	// UserNicknameModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserNicknameModel.
	UserNicknameModel interface {
		userNicknameModel
	}

	customUserNicknameModel struct {
		*defaultUserNicknameModel
	}
)

// NewUserNicknameModel returns a model for the database table.
func NewUserNicknameModel(conn sqlx.SqlConn) UserNicknameModel {
	return &customUserNicknameModel{
		defaultUserNicknameModel: newUserNicknameModel(conn),
	}
}
