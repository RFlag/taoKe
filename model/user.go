package model

import (
	"github.com/jmoiron/sqlx"
)

// 注册
func Register(db sqlx.Ext, mobile string) (int, error) {
	result, err := db.Exec(`INSERT INTO dataoke_user (create_at,active,phone) VALUES (now(),?,?)`,
		1, mobile)
	if err != nil {
		return 0, err
	}
	if rows, rowErr := result.RowsAffected(); rows != 1 || rowErr != nil {
		return 0, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastInsertId), nil
}

// 用户信息
func UserInfo()  {
	
}