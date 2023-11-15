package service

import (
	"encoding/json"
	"testing"
	"workflow-service/service/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_WorkflowDataMarshalling(t *testing.T) {
	type args struct {
		dto model.WorkflowData
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test basic workflow data marshalling",
			args: args{
				dto: model.WorkflowData{
					Id:                    uuid.New(),
					IsFullPageRecognition: true,
					SkipImageEnhancement:  false,
					ExpectDifferentImages: false,
				},
			},
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data, err := json.Marshal(test.args.dto)
			content := string(data[:])
			assert.Equal(t, err != nil, test.wantErr)
			assert.Greater(t, len(content), 0)
		})
	}
}
