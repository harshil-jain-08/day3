package service

import (
	"github.com/harshil-jain-08/day3/Models"
	"github.com/harshil-jain-08/day3/dto"
	"github.com/harshil-jain-08/day3/repo"
)

var Svc Service

type Service interface {
	CreateRecord(data *dto.Record) (model *Models.Record, err error)
	GetRecord(Record *[]Models.Record) (err error)
	UpdateRecord(record *Models.Record) (err error)
	GetRecordByRoll(record *[]Models.Record, roll string) (err error)
	GetRecordBySubject(record *[]Models.Record, subject string) (err error)
	GetRecordByRollAndSub(record *Models.Record, subject string, roll string) (err error)
	// UpdateRecord(record *Models.Record) (err error)
	DeleteRecord(record *Models.Record, roll string) (err error)
}

type service struct {
	repo repo.Repository
}

func (s *service) GetRecord(Record *[]Models.Record) (err error) {
	return s.repo.GetRecord(Record)
}

func (s *service) UpdateRecord(record *Models.Record) (err error) {
	return s.repo.SaveRecord(record)
}

func (s *service) GetRecordByRoll(record *[]Models.Record, roll string) (err error) {
	return s.repo.GetRecordByRoll(record, roll)
}

func (s *service) GetRecordBySubject(record *[]Models.Record, subject string) (err error) {
	return s.repo.GetRecordBySubject(record, subject)
}

func (s *service) GetRecordByRollAndSub(record *Models.Record, subject string, roll string) (err error) {
	return s.repo.GetRecordByRollAndSub(record, subject, roll)
}

func (s *service) DeleteRecord(record *Models.Record, roll string) (err error) {
	return s.repo.DeleteRecord(record, roll)
}

func NewService(repository repo.Repository) {
	Svc = &service{repo: repository}
}

func (s *service) CreateRecord(data *dto.Record) (model *Models.Record, err error) {
	model = &Models.Record{
		Roll:    data.Roll,
		Name:    data.Name,
		Subject: data.Subject,
		Marks:   data.Marks,
	}

	err = s.repo.SaveRecord(model)
	if err != nil {
		return nil, err
	}

	return model, err
}
