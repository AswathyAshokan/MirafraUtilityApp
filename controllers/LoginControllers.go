

package controllers
import (

	"fmt"
	"MirafraUtilityApp/models"

)

type LoginControllers struct {
	BaseController
}

func (c *LoginControllers)LoginChecking() (bool,string){
	login := models.Login{}
	login.Email = c.GetString("email")
	login.Password = []byte(c.GetString("password"))
	fmt.Println("insideeeee")

	w := c.Ctx.ResponseWriter

	dbStatus,positionType := login.CheckLogin()
	 switch dbStatus {
	 case true :
	 	return true,positionType
		 fmt.Fprintf(w,"logged .......")
	 case false:
	 	return false,positionType
		 fmt.Fprintf(w,"incorrect username or password")

	 }
	return true,positionType
}