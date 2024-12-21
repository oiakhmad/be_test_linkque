package user

import "be_test_linkque/core/entity"

type Service interface {
	Create(req *entity.User) (*entity.User, error)
}
