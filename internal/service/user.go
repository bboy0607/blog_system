package service

type CreateUserRequest struct {
	Username  string `form:"username" binding:"required,min=6,max=100"`
	Password  string `form:"password" binding:"required,min=6,max=100"`
	Email     string `form:"email" binding:"required,max=100,email"`
	State     uint8  `form:"state,default=1" binding:"oneof=01"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
}

func (svc Service) CreateUser(param *CreateUserRequest) error {
	return svc.dao.CreateUser(param.Username, param.Password, param.Email, param.State, param.CreatedBy)
}
