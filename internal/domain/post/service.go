package post

import "context"

type Service interface {
	GetPostById(ctx context.Context, id int) *Post
	GetAllPosts(ctx context.Context) []*Post
}
