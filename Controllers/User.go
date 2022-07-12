package Controllers

import (
	"fmt"
	"net/http"

	"github.com/harshil-jain-08/day3/dto"
	"github.com/harshil-jain-08/day3/service"

	"github.com/gin-gonic/gin"
	"github.com/harshil-jain-08/day3/Models"
)

func GetRecords(c *gin.Context) {
	var record []Models.Record
	err := service.Svc.GetRecord(&record)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, record)
	}
}

func CreateRecord(c *gin.Context) {
	var record dto.Record
	c.BindJSON(&record)
	_, err := service.Svc.CreateRecord(&record)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, map[string]bool{"success": true})
	}
}

func GetRecordbyRoll(c *gin.Context) {
	roll := c.Params.ByName("roll")
	var records []Models.Record
	err := service.Svc.GetRecordByRoll(&records, roll)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, records)
	}
}

func GetRecordBySubject(c *gin.Context) {
	subject := c.Params.ByName("subject")
	var records []Models.Record
	err := service.Svc.GetRecordBySubject(&records, subject)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, records)
	}
}

func GetRecordByRollAndSub(c *gin.Context) {
	roll := c.Params.ByName("roll")
	subject := c.Params.ByName("subject")
	var record Models.Record
	err := service.Svc.GetRecordByRollAndSub(&record, subject, roll)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, record)
	}
}

/*func UpdateRecord(c *gin.Context) {
	var record Models.Record
	roll := c.Params.ByName("roll")
	subject := c.Params.ByName("subject")
	err := Models.GetRecordByRollAndSub(&record, subject, roll)
	if err != nil {
		c.JSON(http.StatusNotFound, record)
	}
	c.BindJSON(&record)
	err = Models.UpdateRecord(&record)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, record)
	}
}*/

func DeleteRecordbyRoll(c *gin.Context) {
	var user Models.Record
	roll := c.Params.ByName("roll")
	err := service.Svc.DeleteRecord(&user, roll)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + roll: "is deleted"})
	}
}
