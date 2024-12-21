package user

import "be_test_linkque/core/entity"

type Repository interface {
	Create(req *entity.User) error
}
