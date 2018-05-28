package controllers

import (
	"github.com/AswathyAshokan/MirafraUtilityApp/models"
	"strconv"
	"time"
	"os"
	"log"
	"fmt"
	"io"
)

type PerformaneController struct {
	BaseController
}

func(c *PerformaneController)InsertPerformanceAward()bool{
	performance :=models.PerformanceModel{}
	r := c.Ctx.Request
	//w := c.Ctx.ResponseWriter
	if r.Method == "POST" {

		performance.EmpId = c.GetString("EmpId")
		performance.Name = c.GetString("Name")
		performance.Comment = c.GetString("Comment")

		msec := strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
		//creating a folder for uploading
		if _, err := os.Stat("./testUploadImage/"); os.IsNotExist(err) {

			os.Mkdir("./testUploadImage/", os.ModePerm)
		}
		file, header, err := r.FormFile("uploadfile")
		if err != nil {
			log.Println("uploading error", err)
			//http.Error(w,"error in uploading file",http.StatusInternalServerError)

		}else{
			f, err := os.OpenFile("./testUploadImage/"+msec+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println("image 4 error", err)

			}
			fmt.Println("jst")
			imagePath := "./testUploadImage" + msec + header.Filename

			io.Copy(f, file)
			defer file.Close()
			performance.Photo = imagePath
		}

		fmt.Println( "file  uploaded")
		dbStatus := performance.InsertAward()
		fmt.Println("llll",dbStatus)
		switch dbStatus {
		case true:
			return true
		case false:
			return false
		}
	}
return true
}


func (c *PerformaneController)DisplayPerformaceAward()[][]string{
	var performanceArray [][]string
	dbStatus,allPerformance :=models.GetPerformance()
	switch dbStatus{
	case true:
		for i:=0;i<len(allPerformance);i++{
			var performanceTempArray []string
			performanceTempArray =append(performanceTempArray,allPerformance[i].Name)
			performanceTempArray =append(performanceTempArray,allPerformance[i].EmpId)
			performanceTempArray =append(performanceTempArray,allPerformance[i].Comment)
			performanceTempArray =append(performanceTempArray,allPerformance[i].Photo)
			performanceArray =append(performanceArray,performanceTempArray)



		}
	case false:

	}
	return performanceArray
}