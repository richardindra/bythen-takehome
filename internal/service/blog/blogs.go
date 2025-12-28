package blog

import (
	httpHelper "bythen-takehome/internal/delivery/http"
	"bythen-takehome/internal/entity/blog"
	"context"
	"fmt"
	"time"
)

func (s Service) CreateBlog(ctx context.Context, req blog.Blog, _token string) (blog.Blog, error) {
	var (
		data blog.Blog
	)
	tokenDetail, err := s.GetJWTDetail(_token)
	if err != nil {
		return data, fmt.Errorf("[SERVICE][CreateBlog]: %w", err)
	}
	if time.Now().Unix() > tokenDetail.ExpireIn {
		return data, httpHelper.ErrTokenExpired
	}

	blogID, err := s.data.CreateBlog(ctx, req)
	if err != nil {
		return data, fmt.Errorf("[SERVICE][CreateBlog]: %w", err)
	}

	data, err = s.data.GetBlogByID(ctx, blogID)
	if err != nil {
		return data, fmt.Errorf("[SERVICE][CreateBlog]: %w", err)
	}

	return data, nil
}

func (s Service) GetBlogByID(ctx context.Context, id int64, _token string) (blog.Blog, error) {
	var (
		data blog.Blog
	)
	tokenDetail, err := s.GetJWTDetail(_token)
	if err != nil {
		return data, fmt.Errorf("[SERVICE][GetBlogByID]: %w", err)
	}

	if time.Now().Unix() > tokenDetail.ExpireIn {
		return data, httpHelper.ErrTokenExpired
	}

	data, err = s.data.GetBlogByID(ctx, id)
	if err != nil {
		return data, fmt.Errorf("[SERVICE][GetBlogByID]: %w", err)
	}

	// +1 view count kalau bukan author
	if data.AuthorID != tokenDetail.UserID {
		err = s.data.UpdateViewCount(ctx, id)
		if err != nil {
			return data, fmt.Errorf("[SERVICE][GetBlogByID]: %w", err)
		}

		data.ViewCount += 1
	}

	return data, nil

}

func (s Service) GetAllBlog(ctx context.Context, sortType string, page int, limit int, _token string) ([]blog.Blog, interface{}, error) {
	var (
		resp     = []blog.Blog{}
		metadata = make(map[string]int)
	)
	tokenDetail, err := s.GetJWTDetail(_token)
	if err != nil {
		return resp, metadata, fmt.Errorf("[SERVICE][GetAllBlog]: %w", err)
	}

	if time.Now().Unix() > tokenDetail.ExpireIn {
		return resp, metadata, httpHelper.ErrTokenExpired
	}

	offset := (page - 1) * limit
	datas, err := s.data.GetAllBlog(ctx, sortType, limit, offset)
	if err != nil {
		return resp, metadata, fmt.Errorf("[SERVICE][GetAllBlog]: %w", err)
	}
	count, err := s.data.GetCountAllBlog(ctx)
	if err != nil {
		return resp, metadata, fmt.Errorf("[SERVICE][GetAllBlog]: %w", err)
	}

	metadata["total_data"] = count
	if count == 0 {
		metadata["total_page"] = 0
		return resp, metadata, nil
	}
	resp = datas

	if count%limit == 0 {
		metadata["total_page"] = count / limit
	} else {
		metadata["total_page"] = count/limit + 1
	}

	return resp, metadata, nil
}

func (s Service) GetAllBlogByAuthor(ctx context.Context, authorID int64, sortType string, page int, limit int, _token string) ([]blog.Blog, interface{}, error) {
	var (
		resp     = []blog.Blog{}
		metadata = make(map[string]int)
	)
	tokenDetail, err := s.GetJWTDetail(_token)
	if err != nil {
		return resp, metadata, fmt.Errorf("[SERVICE][GetAllBlogByAuthor]: %w", err)
	}

	if time.Now().Unix() > tokenDetail.ExpireIn {
		return resp, metadata, httpHelper.ErrTokenExpired
	}

	offset := (page - 1) * limit
	datas, err := s.data.GetAllBlogByAuthor(ctx, authorID, sortType, limit, offset)
	if err != nil {
		return resp, metadata, fmt.Errorf("[SERVICE][GetAllBlogByAuthor]: %w", err)
	}
	count, err := s.data.GetCountAllBlogByAuthor(ctx, authorID)
	if err != nil {
		return resp, metadata, fmt.Errorf("[SERVICE][GetAllBlogByAuthor]: %w", err)
	}

	metadata["total_data"] = count
	if count == 0 {
		metadata["total_page"] = 0
		return resp, metadata, nil
	}
	resp = datas

	if count%limit == 0 {
		metadata["total_page"] = count / limit
	} else {
		metadata["total_page"] = count/limit + 1
	}

	return resp, metadata, nil
}

func (s Service) UpdatePost(ctx context.Context, id int64, body blog.Blog, _token string) (blog.Blog, error) {
	var (
		data blog.Blog
	)
	tokenDetail, err := s.GetJWTDetail(_token)
	if err != nil {
		return data, fmt.Errorf("[SERVICE][UpdatePost]: %w", err)
	}

	if time.Now().Unix() > tokenDetail.ExpireIn {
		return data, httpHelper.ErrTokenExpired
	}

	err = s.data.UpdatePost(ctx, id, body)
	if err != nil {
		return data, fmt.Errorf("[SERVICE][UpdatePost]: %w", err)
	}

	data, err = s.data.GetBlogByID(ctx, id)
	if err != nil {
		return data, fmt.Errorf("[SERVICE][UpdatePost]: %w", err)
	}

	return data, nil
}

func (s Service) DeletePost(ctx context.Context, id int64, _token string) error {
	tokenDetail, err := s.GetJWTDetail(_token)
	if err != nil {
		return fmt.Errorf("[SERVICE][DeletePost]: %w", err)
	}

	if time.Now().Unix() > tokenDetail.ExpireIn {
		return httpHelper.ErrTokenExpired
	}

	err = s.data.DeletePost(ctx, id)
	if err != nil {
		return fmt.Errorf("[SERVICE][DeletePost]: %w", err)
	}

	return nil
}
