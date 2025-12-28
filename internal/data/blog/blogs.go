package blog

import (
	httpHelper "bythen-takehome/internal/delivery/http"
	"bythen-takehome/internal/entity/blog"
	"context"
	"database/sql"
	"fmt"
	"strings"
)

func (d Data) CreateBlog(ctx context.Context, blog blog.Blog) (int64, error) {
	var (
		blogID int64
	)
	res, err := (*d.stmt)[createBlog].ExecContext(ctx,
		blog.Title,
		blog.Content,
		blog.AuthorID,
	)
	if err != nil {
		return 0, fmt.Errorf("[DATA][CreateBlog]: %w", err)
	}
	blogID, _ = res.LastInsertId()

	return blogID, nil
}

func (d Data) UpdateViewCount(ctx context.Context, id int64) error {
	_, err := (*d.stmt)[updateViewCount].ExecContext(ctx, id)
	if err != nil {
		return fmt.Errorf("[DATA][UpdateViewCount]: %w", err)
	}

	return nil
}

func (d Data) GetBlogByID(ctx context.Context, id int64) (blog.Blog, error) {
	var (
		data blog.Blog
		err  error
	)

	err = (*d.stmt)[getBlogByID].QueryRowxContext(ctx, id).StructScan(&data)
	if err != nil && err != sql.ErrNoRows {
		return data, fmt.Errorf("[DATA][GetBlogByID]: %w", err)
	} else if err == sql.ErrNoRows {
		return blog.Blog{}, httpHelper.ErrDataNotFound
	}

	return data, err
}

func (d Data) GetAllBlog(ctx context.Context, sortType string, limit int, offset int) ([]blog.Blog, error) {
	var (
		datas []blog.Blog
		data  blog.Blog
		err   error
	)
	sortType = strings.ToUpper(sortType)
	// defaultnya desc
	if sortType != "ASC" && sortType != "DESC" {
		sortType = "DESC"
	}
	tempQuery := strings.ReplaceAll(qGetAllBlog, "[SORTTYPE]", sortType)
	stmt, err := d.db.PreparexContext(ctx, tempQuery)
	if err != nil {
		return datas, fmt.Errorf("[DATA][GetAllBlog]: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryxContext(ctx, limit, offset)
	if err != nil {
		return datas, fmt.Errorf("[DATA][GetAllBlog]: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.StructScan(&data); err != nil {
			return datas, fmt.Errorf("[DATA][GetAllBlog]: %w", err)
		}
		datas = append(datas, data)
	}

	return datas, err
}

func (d Data) GetCountAllBlog(ctx context.Context) (int, error) {
	var (
		total int
		err   error
	)

	err = (*d.stmt)[getCountAllBlog].QueryRowxContext(ctx).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("[DATA][GetCountAllBlog]: %w", err)
	}

	return total, err
}

func (d Data) GetAllBlogByAuthor(ctx context.Context, author int64, sortType string, limit int, offset int) ([]blog.Blog, error) {
	var (
		datas []blog.Blog
		data  blog.Blog
		err   error
	)
	sortType = strings.ToUpper(sortType)
	// defaultnya desc
	if sortType != "ASC" && sortType != "DESC" {
		sortType = "DESC"
	}
	tempQuery := strings.ReplaceAll(qGetAllBlogByAuthor, "[SORTTYPE]", sortType)
	stmt, err := d.db.PreparexContext(ctx, tempQuery)
	if err != nil {
		return datas, fmt.Errorf("[DATA][GetAllBlogByAuthor]: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryxContext(ctx, author, limit, offset)
	if err != nil {
		return datas, fmt.Errorf("[DATA][GetAllBlogByAuthor]: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.StructScan(&data); err != nil {
			return datas, fmt.Errorf("[DATA][GetAllBlogByAuthor]: %w", err)
		}
		datas = append(datas, data)
	}

	return datas, err
}

func (d Data) GetCountAllBlogByAuthor(ctx context.Context, author int64) (int, error) {
	var (
		total int
		err   error
	)

	err = (*d.stmt)[getCountAllBlogByAuthor].QueryRowxContext(ctx, author).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("[DATA][GetCountAllBlogByAuthor]: %w", err)
	}

	return total, err
}

func (d Data) UpdatePost(ctx context.Context, id int64, body blog.Blog) error {
	var (
		err error
	)

	_, err = (*d.stmt)[updatePost].ExecContext(ctx, body.Title, body.Content, id)
	if err != nil {
		return fmt.Errorf("[DATA][UpdatePost]: %w", err)
	}

	return err
}

func (d Data) DeletePost(ctx context.Context, id int64) error {
	var (
		err error
	)

	_, err = (*d.stmt)[deletePost].ExecContext(ctx, id)
	if err != nil {
		return fmt.Errorf("[DATA][DeletePost]: %w", err)
	}

	return err
}
