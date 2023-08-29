package service

type CreateUserRequest struct {
	Username  string `form:"username"`
	Password  string `form:"password"`
	Email     string `form:"email"`
	State     uint8  `form:"state"`
	CreatedBy string `form:"created_by"`
}

func (svc Service) CreateUser(param *CreateUserRequest) error {
	return svc.dao.CreateUser(param.Username, param.Password, param.Email, param.State, param.CreatedBy)
}
