package routers
import (
	"MirafraUtilityApp/controllers"
	"github.com/astaxie/beego"
)

func init() {


	beego.Router("/login", &controllers.LoginControllers{}, "*:LoginChecking")
	beego.Router("/user", &controllers.UserControllers{}, "*:UserInsert")


	//event
	beego.Router("/insertEvent", &controllers.EventController{}, "*:InsertEventDetails")
	beego.Router("/displayEvent", &controllers.EventController{}, "*:DisplayEventDetails")
	beego.Router("/updateEvent/:eventId/:action", &controllers.EventController{}, "*:UpdateEvent")


	//job

	beego.Router("/jobPost", &controllers.JobController{}, "*:JobPosting")
	beego.Router("/jobDisplay", &controllers.JobController{}, "*:DisplayJobDetails")

	//dashboard loading

	beego.Router("/dashBoard", &controllers.DashBoardController{},"*:DashBoardDisplay")

	//product
	beego.Router("/productAdd", &controllers.ProductController{},"*:InserProductDetails")
	beego.Router("/updateProduct/:productId/:action", &controllers.ProductController{}, "*:UpdateProduct")

	//performance
	beego.Router("/performanceAdd", &controllers.PerformaneController{},"*:InsertPerformanceAward")

	//refer
	beego.Router("/referAdd", &controllers.JobReferController{},"*:InsertReferDetails")
	beego.Router("/referDisplay", &controllers.JobReferController{},"*:DisplayJobReferDetails")


}