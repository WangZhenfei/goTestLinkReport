package routers

import (
	"github.com/astaxie/beego"
	"github.com/royburns/goTestLinkReport/apis"
	"github.com/royburns/goTestLinkReport/controllers"
)

func init() {

	// main
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/plan", &controllers.PlanController{})

	// beego.Router("/report", &controllers.ReportController{})
	beego.Router("/getlastexecution", &controllers.GetLastExecutionController{})
	beego.Router("/getsprintexecution", &controllers.GetSprintExecutionController{})
	beego.Router("/getreleaseplan", &controllers.ReleaseController{}, "get:GetReleasePlan")
	beego.Router("/getreleasereport", &controllers.ReleaseController{}, "get:GetReleaseReport")

	// Statistics
	// beego.Router("/statistics", &controllers.StatisticsController{})
	beego.Router("/statistics/sprint", &controllers.StatisticsController{})
	beego.Router("/api/statistics/sprint", &controllers.StatisticsController{}, "get:Sprint")

	// Test
	beego.Router("/test", &controllers.TestController{})

	// About
	beego.Router("/about", &controllers.AboutController{})

	// cmd
	// beego.Router("/cmd", &controllers.CmdController{})

	// api
	beego.Router("/api/getplan", &apis.ApiController{}, "get:GetPlan")
	beego.Router("/api/getlastexecution", &apis.ApiController{}, "get:GetLastExecution")
	beego.Router("/api/getsprintexecution", &apis.ApiController{}, "get:GetSprintExecution")
	beego.Router("/api/runcmd", &apis.ApiController{}, "get:RunCmd")

	// static file
	beego.SetStaticPath("/data", "data")
}
