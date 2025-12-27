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
		return 0, fmt.Errorf("[DATA][CreateBlogCreateBlog]: %w", err)
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
		results []blog.Blog
		result  blog.Blog
		err     error
	)
	sortType = strings.ToUpper(sortType)
	// defaultnya desc
	if sortType != "ASC" && sortType != "DESC" {
		sortType = "DESC"
	}
	tempQuery := strings.ReplaceAll(qGetAllBlog, "[SORTTYPE]", sortType)
	stmt, err := d.db.PreparexContext(ctx, tempQuery)
	if err != nil {
		return results, fmt.Errorf("[DATA][GetAllBlog]: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryxContext(ctx, limit, offset)
	if err != nil {
		return results, fmt.Errorf("[DATA][GetAllBlog]: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.StructScan(&result); err != nil {
			return results, fmt.Errorf("[DATA][GetAllBlog]: %w", err)
		}
		results = append(results, result)
	}

	return results, err
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
