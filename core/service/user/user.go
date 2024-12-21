package usere

import (
	"be_test_linkque/core/entity"
	port "be_test_linkque/core/port/services/user"
	"be_test_linkque/core/port/utils"
)

type service struct {
	repo   port.Repository
	helper utils.Helper
	// middleware middleware.AuthJwt
}

func New(repo port.Repository, helper utils.Helper) *service {
	return &service{repo: repo,
		helper: helper,
	}
}

func (s *service) Create(req *entity.User) (*entity.User, error) {
	// check existing message
	// current, err := s.repo.GetByRequestId(req.RequestId)
	// if err != nil && err != constant.ErrDataNotFound {
	// 	return nil, err
	// }

	// if current.RequestId != "" {
	// 	return nil, constant.ErrMessageDeclinedAlreadyExists
	// }

	// pwd, err := s.helper.HashPassword(req.Password)
	// if err != nil {
	// 	return nil, err
	// }

	// req.Password = pwd

	return req, s.repo.Create(req)
}
