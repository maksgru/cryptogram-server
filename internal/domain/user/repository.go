package user

type Storage interface {
	Create(user *User)
	GetOne(id int) *User
	GetAll(offset, limit string) []*User
	Delete(id int)
}
