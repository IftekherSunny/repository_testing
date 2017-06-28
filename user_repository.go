package repository_testing

// UserRepository struct
type UserRepository struct{}

// GetUser returns a user by the given user id
func (self *UserRepository) GetUser(id int) *User {
	user := new(User)

	dbQueryBuilder.QueryRow("select * from users where id = $1", 1)
	dbQueryBuilder.Scan(&user.Id, &user.Name, &user.Username)

	return user
}
