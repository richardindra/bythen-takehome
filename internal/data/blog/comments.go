package blog

import (
	httpHelper "bythen-takehome/internal/delivery/http"
	"bythen-takehome/internal/entity/blog"
	"context"
	"database/sql"
	"fmt"
	"strings"
)

func (d Data) CreateComment(ctx context.Context, blog blog.Comments) (int64, error) {
	var (
		commentID int64
	)
	res, err := (*d.stmt)[createComment].ExecContext(ctx,
		blog.BlogID,
		blog.AuthorID,
		blog.Content,
	)
	if err != nil {
		return 0, fmt.Errorf("[DATA][CreateComment]: %w", err)
	}
	commentID, _ = res.LastInsertId()

	return commentID, nil
}

func (d Data) GetCommentByID(ctx context.Context, id int64) (blog.Comments, error) {
	var (
		data blog.Comments
		err  error
	)

	err = (*d.stmt)[getCommentByID].QueryRowxContext(ctx, id).StructScan(&data)
	if err != nil && err != sql.ErrNoRows {
		return data, fmt.Errorf("[DATA][GetCommentByID]: %w", err)
	} else if err == sql.ErrNoRows {
		return blog.Comments{}, httpHelper.ErrDataNotFound
	}

	return data, err
}

func (d Data) GetAllCommentsByBlog(ctx context.Context, blogID int64, sortType string, limit int, offset int) ([]blog.Comments, error) {
	var (
		datas []blog.Comments
		data  blog.Comments
		err   error
	)
	sortType = strings.ToUpper(sortType)
	// defaultnya desc
	if sortType != "ASC" && sortType != "DESC" {
		sortType = "DESC"
	}
	tempQuery := strings.ReplaceAll(qGetAllCommentsByBlog, "[SORTTYPE]", sortType)
	stmt, err := d.db.PreparexContext(ctx, tempQuery)
	if err != nil {
		return datas, fmt.Errorf("[DATA][GetAllCommentsByBlog]: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryxContext(ctx, blogID, limit, offset)
	if err != nil {
		return datas, fmt.Errorf("[DATA][GetAllCommentsByBlog]: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.StructScan(&data); err != nil {
			return datas, fmt.Errorf("[DATA][GetAllCommentsByBlog]: %w", err)
		}
		datas = append(datas, data)
	}

	return datas, err
}

func (d Data) GetCountAllCommentsByBlog(ctx context.Context, blogID int64) (int, error) {
	var (
		total int
		err   error
	)

	err = (*d.stmt)[getCountAllCommentsByBlog].QueryRowxContext(ctx, blogID).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("[DATA][GetCountAllCommentsByBlog]: %w", err)
	}

	return total, err
}
