package Controllers

import (
	"day3/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRecords(c *gin.Context) {
	var record []Models.Record
	err := Models.GetRecord(&record)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, record)
	}
}

func CreateRecord(c *gin.Context) {
	var record Models.Record
	c.BindJSON(&record)
	err := Models.CreateRecord(&record)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, record)
	}
}

func GetRecordbyRoll(c *gin.Context) {
	roll := c.Params.ByName("roll")
	var records []Models.Record
	err := Models.GetRecordByRoll(&records, roll)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, records)
	}
}

func GetRecordBySubject(c *gin.Context) {
	subject := c.Params.ByName("subject")
	var records []Models.Record
	err := Models.GetRecordBySubject(&records, subject)
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
	err := Models.GetRecordByRollAndSub(&record, subject, roll)
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
	err := Models.DeleteRecord(&user, roll)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + roll: "is deleted"})
	}
}
