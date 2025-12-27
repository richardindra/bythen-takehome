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
	/*--- BLOG ---*/
	createBlog  = "CreateBlog"
	qCreateBlog = `INSERT INTO blog.m_blog_posts
	(
		title, 
		content, 
		author_id, 
		status, 
		created_at, 
		updated_at
	) 
	VALUES (?, ?, ?, 'Y', NOW(), NOW())`

	updateViewCount  = "UpdateViewCount"
	qUpdateViewCount = `UPDATE blog.m_blog_posts
	SET view_count = view_count + 1
	WHERE id = ?`

	getBlogByID  = "GetBlogByID"
	qGetBlogByID = `SELECT id, title, content, author_id, status, view_count, created_at, updated_at
	FROM blog.m_blog_posts
	WHERE id = ?`

	getAllBlog  = "GetAllBlog"
	qGetAllBlog = `SELECT id, title, content, author_id, status, view_count, created_at, updated_at
	FROM blog.m_blog_posts
	ORDER BY id [SORTTYPE]
	LIMIT ? OFFSET ?`

	getCountAllBlog  = "GetCountAllBlog"
	qGetCountAllBlog = `SELECT COUNT(*) AS Count FROM blog.m_blog_posts`

	getAllBlogByAuthor  = "GetAllBlogByAuthor"
	qGetAllBlogByAuthor = `SELECT id, title, content, author_id, status, view_count, created_at, updated_at
	FROM blog.m_blog_posts
	WHERE author_id = ?
	ORDER BY id [SORTTYPE]
	LIMIT ? OFFSET ?`

	getCountAllBlogByAuthor  = "GetCountAllBlogByAuthor"
	qGetCountAllBlogByAuthor = `SELECT COUNT(*) AS Count FROM blog.m_blog_posts
	WHERE author_id = ?`

	updatePost  = "UpdatePost"
	qUpdatePost = `UPDATE blog.m_blog_posts
	SET title = ?, content = ?, updated_at = NOW()
	WHERE id = ?`

	deletePost  = "DeletePost"
	qDeletePost = `DELETE FROM blog.m_blog_posts
	WHERE id = ?`
)

var (
	selectStmt = []statement{
		/*--- BLOG ---*/
		{getBlogByID, qGetBlogByID},
		{getCountAllBlog, qGetCountAllBlog},
		{getCountAllBlogByAuthor, qGetCountAllBlogByAuthor},
	}
	insertStmt = []statement{
		/*--- BLOG ---*/
		{createBlog, qCreateBlog},
	}
	updateStmt = []statement{
		/*--- BLOG ---*/
		{updateViewCount, qUpdateViewCount},
		{updatePost, qUpdatePost},
	}
	deleteStmt = []statement{
		/*--- BLOG ---*/
		{deletePost, qDeletePost},
	}
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
