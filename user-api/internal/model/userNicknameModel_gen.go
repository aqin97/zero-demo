// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userNicknameFieldNames          = builder.RawFieldNames(&UserNickname{})
	userNicknameRows                = strings.Join(userNicknameFieldNames, ",")
	userNicknameRowsExpectAutoSet   = strings.Join(stringx.Remove(userNicknameFieldNames, "`create_time`", "`update_time`"), ",")
	userNicknameRowsWithPlaceHolder = strings.Join(stringx.Remove(userNicknameFieldNames, "`userId`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	userNicknameModel interface {
		Insert(ctx context.Context, data *UserNickname) (sql.Result, error)
		FindOne(ctx context.Context, userId int64) (*UserNickname, error)
		Update(ctx context.Context, data *UserNickname) error
		Delete(ctx context.Context, userId int64) error
	}

	defaultUserNicknameModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UserNickname struct {
		UserId   int64  `db:"userId"`
		Nickname string `db:"nickname"`
	}
)

func newUserNicknameModel(conn sqlx.SqlConn) *defaultUserNicknameModel {
	return &defaultUserNicknameModel{
		conn:  conn,
		table: "`user_nickname`",
	}
}

func (m *defaultUserNicknameModel) Insert(ctx context.Context, data *UserNickname) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, userNicknameRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.Nickname)
	return ret, err
}

func (m *defaultUserNicknameModel) FindOne(ctx context.Context, userId int64) (*UserNickname, error) {
	query := fmt.Sprintf("select %s from %s where `userId` = ? limit 1", userNicknameRows, m.table)
	var resp UserNickname
	err := m.conn.QueryRowCtx(ctx, &resp, query, userId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserNicknameModel) Update(ctx context.Context, data *UserNickname) error {
	query := fmt.Sprintf("update %s set %s where `userId` = ?", m.table, userNicknameRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Nickname, data.UserId)
	return err
}

func (m *defaultUserNicknameModel) Delete(ctx context.Context, userId int64) error {
	query := fmt.Sprintf("delete from %s where `userId` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, userId)
	return err
}

func (m *defaultUserNicknameModel) tableName() string {
	return m.table
}