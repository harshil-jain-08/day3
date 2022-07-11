package Models

import (
	"day3/Config"
	"fmt"
)

func GetRecord(Record *[]Record) (err error) {
	if err = Config.DB.Find(Record).Error; err != nil {
		return err
	}
	return nil
}

func CreateRecord(record *Record) (err error) {
	if err = Config.DB.Create(record).Error; err != nil {
		return err
	}
	return nil
}

func GetRecordByRoll(record *[]Record, roll string) (err error) {
	if err = Config.DB.Where("roll = ?", roll).Find(&record).Error; err != nil {
		return err
	}
	return nil
}

func GetRecordBySubject(record *[]Record, subject string) (err error) {
	if err = Config.DB.Where("subject = ?", subject).Find(&record).Error; err != nil {
		return err
	}
	return nil
}
func GetRecordByRollAndSub(record *Record, subject string, roll string) (err error) {
	if err = Config.DB.Where("subject = ? AND roll = ?", subject, roll).Find(&record).Error; err != nil {
		return err
	}
	return nil
}

func UpdateRecord(record *Record) (err error) {
	fmt.Println(record)
	Config.DB.Save(record)
	return nil
}
func DeleteRecord(record *Record, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(record)
	return nil
}
