package test

import (
	"gin-scaffold/service"
	"testing"

	"gin-scaffold/controller"

	"github.com/gin-gonic/gin"
)

func TestUserController_CreateUser(t *testing.T) {
	type fields struct {
		Service *service.UserService
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &controller.UserController{
				Service: tt.fields.Service,
			}
			uc.CreateUser(tt.args.c)
		})
	}
}

func TestUserController_GetAllUsers(t *testing.T) {
	type fields struct {
		ApiVersion string
		Service    *service.UserService
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &controller.UserController{
				Service: tt.fields.Service,
			}
			uc.GetAllUsers(tt.args.c)
		})
	}
}

func TestUserController_SearchUsersByKeys(t *testing.T) {
	type fields struct {
		ApiVersion string
		Service    *service.UserService
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &controller.UserController{
				Service: tt.fields.Service,
			}
			uc.SearchUsersByKeys(tt.args.c)
		})
	}
}

func TestUserController_GetUserByID(t *testing.T) {
	type fields struct {
		ApiVersion string
		Service    *service.UserService
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &controller.UserController{
				Service: tt.fields.Service,
			}
			uc.GetUserByID()
		})
	}
}

func TestUserController_UpdateUser(t *testing.T) {
	type fields struct {
		ApiVersion string
		Service    *service.UserService
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &controller.UserController{
				Service: tt.fields.Service,
			}
			uc.UpdateUser(tt.args.c)
		})
	}
}

func TestUserController_DeleteUser(t *testing.T) {
	type fields struct {
		ApiVersion string
		Service    *service.UserService
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &controller.UserController{
				Service: tt.fields.Service,
			}
			uc.DeleteUser(tt.args.c)
		})
	}
}
