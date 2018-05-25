
package controllers

import (
	"MirafraUtilityApp/models"
	"strconv"
	"time"
	"os"
	"log"
	"fmt"
	"io"
)

type UserControllers struct {
	BaseController
}

func (c *UserControllers)UserInsert() bool{
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	user :=models.User{}
	if r.Method == "POST" {
		user.PositionType = c.GetString("PositionType")
		user.DateOfJoin = c.GetString("DateOfJoin")
		user.Dob = c.GetString("Dob")
		user.Email = c.GetString("Email")
		user.EmpId = c.GetString("EmpId")
		user.Experience = c.GetString("Experience")
		user.LastCompany = c.GetString("LastCompany")
		user.LastCompanyLocation = c.GetString("LastCompanyLocation")
		user.Manager = c.GetString("Manager")
		user.Password = []byte("mirafra")
		user.FirstName = c.GetString("FirstName")
		user.LastName = c.GetString("LastName")
		user.Address = c.GetString("Address")
		user.City = c.GetString("City")
		user.State = c.GetString("State")
		user.PinCode = c.GetString("PinCode")
		user.PhoneNumber = c.GetString("PhoneNumber")
		user.Gender = c.GetString("Gender")
		user.JobStatus = c.GetString("JobStatus")

		//image upload code
		msec := strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
		//creating a folder for uploading
		if _, err := os.Stat("./testUploadImage/"); os.IsNotExist(err) {

			os.Mkdir("./testUploadImage/", os.ModePerm)
		}
		file, header, err := r.FormFile("uploadfile")
		if err != nil {
			log.Println("uploading error", err)
			//http.Error(w,"error in uploading file",http.StatusInternalServerError)

		}
		f, err := os.OpenFile("./testUploadImage/"+msec+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("image 4 error", err)

		}
		fmt.Println("jst")
		imagePath := "./testUploadImage" + msec + header.Filename

		io.Copy(f, file)
		defer file.Close()
		user.UploadPhoto = imagePath
		fmt.Fprintf(w, "file  uploaded")

		//insertion of data
		dbStatus := user.InserIntoUser(user)
		switch dbStatus {
		case true:

			return true
			fmt.Println("insert")
		case false:
			return false
			fmt.Println("incorrect")

		}
	}
	return true

}

