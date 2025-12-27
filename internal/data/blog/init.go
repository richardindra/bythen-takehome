package blog

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

type (
	Data struct {
		db   *sqlx.DB
		stmt *map[string]*sqlx.Stmt
	}

	statement struct {
		key   string
		query string
	}
)

const (
	/*--- USER ---*/
	createUser  = "CreateUser"
	qCreateUser = `INSERT INTO appdb.m_users
	(
		username, 
		name, 
		email, 
		password_hash, 
		status, 
		last_login_at, 
		created_at, 
		updated_at
	) 
	VALUES (?, ?, ?, ?, 'A', NOW(), NOW(), NOW())`

	checkUser  = "CheckUser"
	qCheckUser = `SELECT COUNT(id) FROM appdb.m_users
					WHERE username = ? OR email = ?`

	getUserByUsername  = "GetUserByUsername"
	qGetUserByUsername = `SELECT id, username, name, email, password_hash, status, last_login_at, created_at, updated_at
	FROM appdb.m_users 
	WHERE username = ?`

	updateLastLogin  = "UpdateLastLogin"
	qUpdateLastLogin = `UPDATE appdb.m_users
	SET last_login_at = NOW()
	WHERE username = ?`
)

var (
	selectStmt = []statement{
		{checkUser, qCheckUser},
		{getUserByUsername, qGetUserByUsername},
	}
	insertStmt = []statement{
		{createUser, qCreateUser},
	}
	updateStmt = []statement{
		{updateLastLogin, qUpdateLastLogin},
	}
	deleteStmt = []statement{}
)

// New ...
func New(db *sqlx.DB) *Data {
	var (
		stmts = make(map[string]*sqlx.Stmt)
	)

	d := &Data{
		db:   db,
		stmt: &stmts,
	}

	d.InitStmt()

	return d
}

func (d *Data) InitStmt() {
	var (
		err   error
		stmts = make(map[string]*sqlx.Stmt)
	)

	for _, v := range selectStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize select statement key %v, err : %v", v.key, err)
		}
	}

	for _, v := range insertStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize insert statement key %v, err : %v", v.key, err)
		}
	}

	for _, v := range updateStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize update statement key %v, err : %v", v.key, err)
		}
	}

	for _, v := range deleteStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize delete statement key %v, err : %v", v.key, err)
		}
	}

	*d.stmt = stmts
}
