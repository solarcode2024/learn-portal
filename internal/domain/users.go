package domain

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

// ini untuk definisi service/fitur yang berkaitan dengan user
type UserService interface {
	Register(data *User) error

	GetAllUser() ([]User, error)

	GetUserProfile(id int) (*User, error)
}

// ini untuk definisi interaksi ke database
type UserRepository interface {
	// ini untuk insert user baru
	InsertUser(data *User) error

	// ini untuk ngambil data users return 0 => []User, 1 => error
	SelectAllUser() ([]User, error)

	// ini untuk ngambil data users berdasakan id
	SelectOneUserByID(id int) (*User, error)
}
