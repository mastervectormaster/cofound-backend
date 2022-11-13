package user

import (
	"github.com/mastervectormaster/cofound-backend/model"
)

type Store interface {
	GetByID(uint) (*model.User, error)
	GetByEmail(string) (*model.User, error)
	GetByUsername(string) (*model.User, error)
	Create(*model.User) error
	Update(*model.User) error
}
