package test

import (
	"gin-scaffold/service"
	"testing"

	"github.com/gin-gonic/gin"
	_ "gorm.io/gorm"

	"gin-scaffold/controller"
)

func TestRunJobController_CreateJob(t *testing.T) {
	type fields struct {
		ApiVersion string
		Service    *service.RunJobService
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
			ctl := &controller.RunJobController{
				ApiVersion: tt.fields.ApiVersion,
				Service:    tt.fields.Service,
			}
			ctl.CreateJob(tt.args.c)
		})
	}
}

func TestRunJobController_GetAllJobs(t *testing.T) {
	type fields struct {
		ApiVersion string
		Service    *service.RunJobService
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
			ctl := &controller.RunJobController{
				ApiVersion: tt.fields.ApiVersion,
				Service:    tt.fields.Service,
			}
			ctl.GetAllJobs(tt.args.c)
		})
	}
}

func TestRunJobController_SearchJobByKeys(t *testing.T) {
	type fields struct {
		ApiVersion string
		Service    *service.RunJobService
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
			ctl := &controller.RunJobController{
				ApiVersion: tt.fields.ApiVersion,
				Service:    tt.fields.Service,
			}
			ctl.SearchJobByKeys(tt.args.c)
		})
	}
}

func TestRunJobController_GetJobByID(t *testing.T) {
	type fields struct {
		ApiVersion string
		Service    *service.RunJobService
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
			ctl := &controller.RunJobController{
				ApiVersion: tt.fields.ApiVersion,
				Service:    tt.fields.Service,
			}
			ctl.GetJobByID(tt.args.c)
		})
	}
}

func TestRunJobController_UpdateJob(t *testing.T) {
	type fields struct {
		ApiVersion string
		Service    *service.RunJobService
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
			ctl := &controller.RunJobController{
				ApiVersion: tt.fields.ApiVersion,
				Service:    tt.fields.Service,
			}
			ctl.UpdateJob(tt.args.c)
		})
	}
}

func TestRunJobController_DeleteJob(t *testing.T) {
	type fields struct {
		ApiVersion string
		Service    *service.RunJobService
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
			ctl := &controller.RunJobController{
				ApiVersion: tt.fields.ApiVersion,
				Service:    tt.fields.Service,
			}
			ctl.DeleteJob(tt.args.c)
		})
	}
}

func TestRunJobController_DataSyncJob(t *testing.T) {
	type fields struct {
		ApiVersion string
		Service    *service.RunJobService
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
			ctl := &controller.RunJobController{
				ApiVersion: tt.fields.ApiVersion,
				Service:    tt.fields.Service,
			}
			ctl.DataSyncJob(tt.args.c)
		})
	}
}

func TestRunJobController_SyncDatatoFile(t *testing.T) {
	type fields struct {
		ApiVersion string
		Service    *service.RunJobService
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := &controller.RunJobController{
				ApiVersion: tt.fields.ApiVersion,
				Service:    tt.fields.Service,
			}
			if _, err := ctl.SyncDatatoFile(); (err != nil) != tt.wantErr {
				t.Errorf("RunJobController.SyncDatatoFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
