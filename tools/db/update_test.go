package db

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nmarsollier/cartgo/tools/errs"
	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDb := NewMockDBConn(ctrl)
	mockDb.EXPECT().Exec("query", "arg1", "arg2").Return(nil).Times(1)

	update := GetUpdate(mockDb)

	err := update.Exec("query", "arg1", "arg2")
	assert.NoError(t, err)
}

func TestUpdate_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDb := NewMockDBConn(ctrl)
	mockDb.EXPECT().Exec("query", "arg1", "arg2").Return(errs.NotFound).Times(1)

	update := GetUpdate(mockDb)

	err := update.Exec("query", "arg1", "arg2")
	assert.Error(t, err)
	assert.Equal(t, errs.NotFound, err)
}
