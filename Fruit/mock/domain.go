package mock

import (
	"github.com/arijitnayak92/taskAfford/Fruit/model"
	"github.com/stretchr/testify/mock"
)

// Domain ...
type Domain struct {
	mock.Mock
}

// GetPostgresHealth ...
func (mock *Domain) GetPostgresHealth() bool {
	args := mock.Called()

	return args.Bool(0)
}

// GetMongoHealth ...
func (mock *Domain) GetMongoHealth() bool {
	args := mock.Called()

	return args.Bool(0)
}

// UserSignup ...
func (mock *Domain) UserSignup(email string, password string, firstname string, lastname string, role string, cartid string) (bool, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(bool), args.Error(1)
}

// GetUser ...
func (mock *Domain) GetUser(email string) (*model.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*model.User), args.Error(1)
}

// Login ...
func (mock *Domain) Login(email string, password string) (string, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(string), args.Error(1)
}
