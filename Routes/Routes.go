package Routes

import (
	"github.com/harshil-jain-08/day3/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/user-api")
	{
		grp1.GET("user", Controllers.GetRecords)
		grp1.POST("user", Controllers.CreateRecord)
		grp1.GET("roll/:roll", Controllers.GetRecordbyRoll)
		grp1.GET("sub/:subject", Controllers.GetRecordBySubject)
		grp1.GET("rollandsub/:roll/:subject", Controllers.GetRecordByRollAndSub)
		// grp1.GET("roll/:roll/sub/:subject", Controllers.GetRecordByRollAndSub)
		// grp1.PUT("update/:roll/:subject", Controllers.UpdateRecord)
		grp1.DELETE("user/:roll", Controllers.DeleteRecordbyRoll)
	}
	return r
}
