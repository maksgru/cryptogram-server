package post

type Storage interface {
	Create(post *Post)
	GetOne(id int) *Post
	GetAll(offset, limit string) []*Post
	Delete(id int)
}
