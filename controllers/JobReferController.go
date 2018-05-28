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

type JobReferController struct {
	BaseController
}
func(c *JobReferController)InsertReferDetails()bool{
	jobRefer :=models.JobRefer{}
	r := c.Ctx.Request
	//w := c.Ctx.ResponseWriter


	if r.Method == "POST" {
		jobRefer.JobId = c.GetString("JobId")
		jobRefer.Location = c.GetString("Location")
		jobRefer.CandidateName = c.GetString("CandidateName")
		jobRefer.Experience = c.GetString("Experience")
		jobRefer.JobTitle = c.GetString("JobTitle")
		jobRefer.ReferStatus = false
		msec := strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
		//creating a folder for uploading
		if _, err := os.Stat("./testUploadImage/"); os.IsNotExist(err) {

			os.Mkdir("./testUploadImage/", os.ModePerm)
		}
		file, header, err := r.FormFile("uploadfile")
		if err != nil {
			log.Println("uploading error", err)
		}else{

			f, err := os.OpenFile("./testUploadImage/"+msec+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println("image 4 error", err)
				return true
			}
			fmt.Println("jst")
			imagePath := "./testUploadImage" + msec + header.Filename

			io.Copy(f, file)
			defer file.Close()
			jobRefer.Resume = imagePath
		}

		//fmt.Fprintf(w, "file  uploaded")
		dbStatus := jobRefer.InsertJobRrfer()
		switch  dbStatus {
		case true:
			return  true
		case false:
			return false

		}
	}
return true
}
func( c *JobReferController) DisplayJobReferDetails()[][]string{
	dbStatus,jobReferDetails :=models.DisplayJobReferDetails()
	var JobReferArray [][]string
	switch dbStatus{
	case true:
		for i :=0;i<len(jobReferDetails);i++{
			var JobReferTempArray []string
			JobReferTempArray =append(JobReferTempArray,jobReferDetails[i].JobId)
			JobReferTempArray =append(JobReferTempArray,jobReferDetails[i].Resume)
			JobReferTempArray =append(JobReferTempArray,jobReferDetails[i].JobTitle)
			JobReferTempArray =append(JobReferTempArray,jobReferDetails[i].Experience)
			JobReferTempArray =append(JobReferTempArray,jobReferDetails[i].CandidateName)
			JobReferTempArray =append(JobReferTempArray,jobReferDetails[i].Location)
			JobReferTempArray =append(JobReferTempArray,jobReferDetails[i].CandidatePhoneNumber)
			JobReferArray =append(JobReferArray,JobReferTempArray)
		}
	case false:
	}
	fmt.Println("lllll",JobReferArray)
	return JobReferArray
}