package repo

import (
	"fmt"

	"github.com/harshil-jain-08/day3/Models"

	"github.com/harshil-jain-08/day3/Config"
)

type Repository interface {
	GetRecord(Record *[]Models.Record) (err error)
	SaveRecord(record *Models.Record) (err error)
	GetRecordByRoll(record *[]Models.Record, roll string) (err error)
	GetRecordBySubject(record *[]Models.Record, subject string) (err error)
	GetRecordByRollAndSub(record *Models.Record, subject string, roll string) (err error)
	// UpdateRecord(record *Models.Record) (err error)
	DeleteRecord(record *Models.Record, roll string) (err error)
}

func NewRepo() Repository {
	return &repo{}
}

type repo struct{}

func (r *repo) GetRecord(Record *[]Models.Record) (err error) {
	if err = Config.DB.Find(Record).Error; err != nil {
		return err
	}
	return nil
}

func (r *repo) SaveRecord(record *Models.Record) (err error) {
	if err = Config.DB.Save(record).Error; err != nil {
		return err
	}
	return nil
}

func (r *repo) GetRecordByRoll(record *[]Models.Record, roll string) (err error) {
	fmt.Printf("roll called for %s", roll)
	if err = Config.DB.Where("roll = ?", roll).Find(&record).Error; err != nil {
		return err
	}
	return nil
}

func (r *repo) GetRecordBySubject(record *[]Models.Record, subject string) (err error) {
	if err = Config.DB.Where("subject = ?", subject).Find(&record).Error; err != nil {
		return err
	}
	return nil
}

func (r *repo) GetRecordByRollAndSub(record *Models.Record, subject string, roll string) (err error) {
	if err = Config.DB.Where("subject = ? AND roll = ?", subject, roll).Find(&record).Error; err != nil {
		return err
	}
	return nil
}

func (r *repo) DeleteRecord(record *Models.Record, roll string) (err error) {
	Config.DB.Where("roll = ?", roll).Delete(record)
	return nil
}
