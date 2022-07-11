package Routes

import (
	"day3/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/user-api")
	{
		grp1.GET("user", Controllers.GetRecords)
		grp1.POST("user", Controllers.CreateRecord)
		grp1.GET("user/:roll", Controllers.GetRecordbyRoll)
		//grp1.GET("user/sub/:subject", Controllers.GetRecordBySubject)
		//grp1.GET("user/rollandsub", Controllers.GetRecordByRollAndSub)
		//grp1.PUT("user/:roll", Controllers.UpdateRecord)
		grp1.DELETE("user/:roll", Controllers.DeleteRecordbyRoll)
	}
	return r
}
