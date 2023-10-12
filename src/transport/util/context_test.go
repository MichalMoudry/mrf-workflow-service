package util

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithUserId(t *testing.T) {
	type args struct {
		ctx    context.Context
		userId string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test adding a user ID to a context",
			args: args{
				ctx:    context.Background(),
				userId: "test_id",
			},
		},
	}
	for _, tt := range tests {
		ctx := WithUserId(tt.args.ctx, tt.args.userId)
		uid, ok := GetUserIdFromCtx(ctx)

		assert.Equal(t, tt.args.userId, uid)
		assert.True(t, ok)
		assert.NotNil(t, ctx)
	}
}
