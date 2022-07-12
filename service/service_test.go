package service

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/harshil-jain-08/day3/Models"
	"github.com/harshil-jain-08/day3/dto"
	repomock "github.com/harshil-jain-08/day3/repo/mock"
	"github.com/stretchr/testify/assert"
)

func TestCreateRecord_Success(t *testing.T) {
	assertion := assert.New(t)

	ctrl := gomock.NewController(t)
	repo := repomock.NewMockRepository(ctrl)
	NewService(repo)

	data := dto.Record{
		Roll:    "01",
		Name:    "Arvind",
		Subject: "Science",
		Marks:   "24",
	}

	expectedModel := Models.Record{
		Roll:    data.Roll,
		Name:    data.Name,
		Subject: data.Subject,
		Marks:   data.Marks,
	}

	repo.EXPECT().SaveRecord(&expectedModel).Do(
		func(model *Models.Record) {
			model.ID = 1
			model.CreatedAt = time.Now()
			model.UpdatedAt = time.Now()
		}).Return(nil).Times(1)

	out, err := Svc.CreateRecord(&data)
	assertion.Nil(err)
	assertion.Equal(out.Roll, expectedModel.Roll)
	assertion.Equal(out.Name, expectedModel.Name)
	assertion.Equal(out.Subject, expectedModel.Subject)
	assertion.Equal(out.Marks, expectedModel.Marks)
	assertion.NotEmpty(out.ID)
	assertion.NotEmpty(out.CreatedAt)
	assertion.NotEmpty(out.UpdatedAt)
}
