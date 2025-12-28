package blog

import (
	httpHelper "bythen-takehome/internal/delivery/http"
	"bythen-takehome/internal/entity/blog"
	"context"
	"fmt"
	"time"
)

func (s Service) CreateComment(ctx context.Context, req blog.Comments, _token string) (blog.Comments, error) {
	var (
		data blog.Comments
	)
	tokenDetail, err := s.GetJWTDetail(_token)
	if err != nil {
		return data, fmt.Errorf("[SERVICE][CreateComment]: %w", err)
	}
	if time.Now().Unix() > tokenDetail.ExpireIn {
		return data, httpHelper.ErrTokenExpired
	}

	req.AuthorID = tokenDetail.UserID
	commentID, err := s.data.CreateComment(ctx, req)
	if err != nil {
		return data, fmt.Errorf("[SERVICE][CreateComment]: %w", err)
	}

	data, err = s.data.GetCommentByID(ctx, commentID)
	if err != nil {
		return data, fmt.Errorf("[SERVICE][CreateComment]: %w", err)
	}

	return data, nil
}

func (s Service) GetAllCommentsByBlog(ctx context.Context, blogID int64, sortType string, page int, limit int, _token string) ([]blog.Comments, interface{}, error) {
	var (
		resp     = []blog.Comments{}
		metadata = make(map[string]int)
	)
	tokenDetail, err := s.GetJWTDetail(_token)
	if err != nil {
		return resp, metadata, fmt.Errorf("[SERVICE][GetAllCommentsByBlog]: %w", err)
	}

	if time.Now().Unix() > tokenDetail.ExpireIn {
		return resp, metadata, httpHelper.ErrTokenExpired
	}

	offset := (page - 1) * limit
	datas, err := s.data.GetAllCommentsByBlog(ctx, blogID, sortType, limit, offset)
	if err != nil {
		return resp, metadata, fmt.Errorf("[SERVICE][GetAllCommentsByBlog]: %w", err)
	}
	count, err := s.data.GetCountAllCommentsByBlog(ctx, blogID)
	if err != nil {
		return resp, metadata, fmt.Errorf("[SERVICE][GetAllCommentsByBlog]: %w", err)
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
