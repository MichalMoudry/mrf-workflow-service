package util

import (
	"testing"
	"workflow-service/transport/model/contracts"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_ContractValidation(t *testing.T) {
	type args struct {
		dto any
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test correct CreateAppRequest DTO validation",
			args: args{
				dto: contracts.CreateAppRequest{
					Name: "test_app_1",
				},
			},
			wantErr: false,
		},
		{
			name: "Test CreateAppRequest DTO with empty name",
			args: args{
				dto: contracts.CreateAppRequest{},
			},
			wantErr: true,
		},
		{
			name: "Test CreateAppRequest DTO with name overload",
			args: args{
				dto: contracts.CreateAppRequest{
					Name: uuid.NewString() + uuid.NewString() + uuid.NewString() + uuid.NewString() + uuid.NewString() + uuid.NewString(),
				},
			},
			wantErr: true,
		},
		{
			name: "Test correct CreateWorkflowRequest DTO validation",
			args: args{
				dto: contracts.CreateWorkflowRequest{
					Name:                  "test_workflow_1",
					AppId:                 uuid.NewString(),
					IsFullPageRecognition: "true",
					ExpectDifferentImages: "false",
					SkipImageEnhancement:  "false",
				},
			},
			wantErr: false,
		},
		{
			name: "Test empty CreateWorkflowRequest DTO validation",
			args: args{
				dto: contracts.CreateWorkflowRequest{},
			},
			wantErr: true,
		},
		{
			name: "Test incorrect AppId in CreateWorkflowRequest DTO",
			args: args{
				dto: contracts.CreateWorkflowRequest{
					Name:                  "test_workflow_1",
					AppId:                 "test_app_1",
					IsFullPageRecognition: "true",
					ExpectDifferentImages: "false",
					SkipImageEnhancement:  "false",
				},
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			isError := validate.Struct(test.args.dto) != nil
			assert.Equal(t, isError, test.wantErr)
		})
	}
}
