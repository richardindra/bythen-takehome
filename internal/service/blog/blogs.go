package blog

import (
	"bythen-takehome/internal/entity/blog"
	"context"
	"fmt"
)

func (s Service) GetBlogByID(ctx context.Context, id int64) (blog.Blog, error) {
	data, err := s.data.GetBlogByID(ctx, id)
	if err != nil {
		return data, fmt.Errorf("[SERVICE][GetBlogByID]: %w", err)
	}

	return data, err
}
