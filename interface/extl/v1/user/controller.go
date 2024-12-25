package email

import (
	"be_test_linkque/core/entity"
	"be_test_linkque/core/port/services/user"
	"be_test_linkque/interface/extl/v1/common"
	"be_test_linkque/utils/validator"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service user.Service
}

func New(service user.Service) *Controller {
	return &Controller{service}
}

func (controller *Controller) Create(e echo.Context) error {
	request := new(Indentify)

	if err := e.Bind(&request); err != nil {
		return common.NewGeneralBadRequest(err.Error())
	}

	if err := validator.GetValidator().Struct(request); err != nil {
		return e.JSON(http.StatusBadRequest, &common.GeneralResponse{
			Status: common.GeneralStatus{
				Code:    "400",
				Message: "Save user failed",
				Error:   validator.GetGeneralValidatorMessage(err),
			},
		})
	}

	result, err := mapInputToStruct(request)
	if err != nil {
		return e.JSON(http.StatusBadRequest, &common.GeneralResponse{
			Status: common.GeneralStatus{
				Code:    "400",
				Message: "Save user failed",
				Error:   validator.GetGeneralValidatorMessage(err),
			},
		})
	}

	user, err := controller.service.Create(result)
	if err != nil {
		return common.NewGeneralUnprocessableEntity(err.Error())
	}

	return e.JSON(http.StatusCreated, &common.GeneralResponse{
		Status: common.GeneralStatus{
			Code:    "20100001",
			Message: "Save User success",
		},
		Data: mapResponsUser(user),
	})

}
func mapResponsUser(user *entity.User) *ResponseUser {
	res := new(ResponseUser)
	res.Name = user.Name
	res.Age = user.Age
	res.City = user.City

	res.Create_at = time.Now()
	return res
}
func mapInputToStruct(input *Indentify) (*entity.User, error) {
	user := new(entity.User)
	parts := strings.Fields(input.InputFormat)
	nameParts := []string{}
	age := ""
	cityParts := []string{}

	for i, part := range parts {
		if _, err := strconv.Atoi(part); err == nil {
			age = part

			// Check if the next part is TAHUN, THN, or TH and skip it
			if i+1 < len(parts) {
				nextPart := strings.ToUpper(parts[i+1])
				if nextPart == "TAHUN" || nextPart == "THN" || nextPart == "TH" {
					cityParts = parts[i+2:]
				} else {
					cityParts = parts[i+1:]
				}
			}

			break
		}
		nameParts = append(nameParts, part)
	}

	if age == "" || len(cityParts) == 0 {
		return user, fmt.Errorf("invalid input format")
	}

	name := strings.ToUpper(strings.Join(nameParts, " "))
	city := strings.ToUpper(strings.ToLower(strings.Join(cityParts, " ")))

	ageInt, err := strconv.Atoi(age)
	if err != nil {
		return user, fmt.Errorf("invalid age format")
	}
	user.Name = name
	user.Age = ageInt
	user.City = city
	return user, nil
}
