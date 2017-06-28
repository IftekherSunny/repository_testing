package repository_testing

import "testing"

// Declaring testing required variables
var (
	id       interface{}
	name     interface{}
	username interface{}
)

// Mocking dbQueryBuilder
type MockDbQueryBuilder struct {
	t *testing.T
}

// Mocked QueryRow method
func (self *MockDbQueryBuilder) QueryRow(query string, args ...interface{}) {

	// assert
	if query != "select * from users where id = $1" {
		self.t.Error("Query string should be 'select * from users where id = $1'")
	}

	// assert
	if len(args) != 1 {
		self.t.Error("One argument should be passed for the query string")
	}
}

// Mocked Scan method
func (self *MockDbQueryBuilder) Scan(dest ...interface{}) error {

	// assert
	if len(dest) != 3 {
		self.t.Error("Three arguments must be passed")
	}

	// storing arguments memory address
	// for the next assertion
	id = dest[0]
	name = dest[1]
	username = dest[2]

	return nil
}

func TestGetUser(t *testing.T) {
	t.Run("It should return a user", func(t *testing.T) {

		// setup
		dbQueryBuilder = &MockDbQueryBuilder{t: t}

		ur := new(UserRepository)

		// act
		user := ur.GetUser(1)

		// assert
		if &user.Id != id {
			t.Error("Passing wrong memory address for the user's id")
		}

		if &user.Name != name {
			t.Error("Passing wrong memory address for the user's name")
		}

		if &user.Username != username {
			t.Error("Passing wrong memory address for the user's username")
		}
	})
}
