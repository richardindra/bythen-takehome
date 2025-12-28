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
	qGetBlogByID = `SELECT bp.id, bp.title, bp.content, bp.author_id, bp.status, bp.view_count, bp.created_at, bp.updated_at,
	IFNULL(usr.username, '') AS username, 
	IFNULL(usr.name, '') AS name
	FROM blog.m_blog_posts bp
	LEFT JOIN blog.m_users usr
	ON bp.author_id = usr.id
	WHERE bp.id = ?`

	getAllBlog  = "GetAllBlog"
	qGetAllBlog = `SELECT bp.id, bp.title, bp.content, bp.author_id, bp.status, bp.view_count, bp.created_at, bp.updated_at,
	IFNULL(usr.username, '') AS username, 
	IFNULL(usr.name, '') AS name
	FROM blog.m_blog_posts bp
	LEFT JOIN blog.m_users usr
	ON bp.author_id = usr.id
	ORDER BY bp.id [SORTTYPE]
	LIMIT ? OFFSET ?`

	getCountAllBlog  = "GetCountAllBlog"
	qGetCountAllBlog = `SELECT COUNT(*) AS Count FROM blog.m_blog_posts`

	getAllBlogByAuthor  = "GetAllBlogByAuthor"
	qGetAllBlogByAuthor = `SELECT bp.id, bp.title, bp.content, bp.author_id, bp.status, bp.view_count, bp.created_at, bp.updated_at,
	IFNULL(usr.username, '') AS username, 
	IFNULL(usr.name, '') AS name
	FROM blog.m_blog_posts bp
	LEFT JOIN blog.m_users usr
	ON bp.author_id = usr.id
	WHERE bp.author_id = ?
	ORDER BY bp.id [SORTTYPE]
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

	/*--- COMMENTS ---*/
	createComment  = "CreateComment"
	qCreateComment = `INSERT INTO blog.m_comments
	(
		blog_id,
		author_id,
		content,
		created_at,
		updated_at
	) 
	VALUES (?, ?, ?, NOW(), NOW())`

	getCommentByID  = "GetCommentByID"
	qGetCommentByID = `SELECT cmt.id, cmt.blog_id, cmt.author_id, cmt.content, cmt.created_at, cmt.updated_at,
	IFNULL(usr.username, '') AS username, 
	IFNULL(usr.name, '') AS name
	FROM blog.m_comments cmt
	LEFT JOIN blog.m_users usr
	ON cmt.author_id = usr.id
	WHERE cmt.id = ?`

	getAllCommentsByBlog  = "GetAllCommentsByBlog"
	qGetAllCommentsByBlog = `SELECT cmt.id, cmt.blog_id, cmt.author_id, cmt.content, cmt.created_at, cmt.updated_at,
	IFNULL(usr.username, '') AS username, 
	IFNULL(usr.name, '') AS name
	FROM blog.m_comments cmt
	LEFT JOIN blog.m_users usr
	ON cmt.author_id = usr.id
	WHERE cmt.blog_id = ?
	ORDER BY cmt.id [SORTTYPE]
	LIMIT ? OFFSET ?`

	getCountAllCommentsByBlog  = "GetCountAllCommentsByBlog"
	qGetCountAllCommentsByBlog = `SELECT COUNT(*) AS Count FROM blog.m_comments
	WHERE blog_id = ?`
)

var (
	selectStmt = []statement{
		/*--- BLOG ---*/
		{getBlogByID, qGetBlogByID},
		{getCountAllBlog, qGetCountAllBlog},
		{getCountAllBlogByAuthor, qGetCountAllBlogByAuthor},

		/*--- COMMENTS ---*/
		{getCommentByID, qGetCommentByID},
		{getCountAllCommentsByBlog, qGetCountAllCommentsByBlog},
	}
	insertStmt = []statement{
		/*--- BLOG ---*/
		{createBlog, qCreateBlog},
		
		/*--- COMMENTS ---*/
		{createComment, qCreateComment},
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
