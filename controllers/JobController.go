package controllers

import (
	"MirafraUtilityApp/models"
	"fmt"
)

type JobController struct {

 	BaseController
}
func (c *JobController)JobPosting() bool{
	//request :=c.Ctx.Request
	//response :=c.Ctx.ResponseWriter
	r := c.Ctx.Request

	if r.Method == "POST" {
		jobPost := models.JobRequirement{}
		jobPost.Designation = c.GetString("Designation")
		jobPost.MinExperience = c.GetString("MinExperience")
		jobPost.MaxExperience = c.GetString("MaxExperience")
		jobPost.JobLocation = c.GetString("JobLocation")
		jobPost.KeySkills = c.GetString("KeySkills")
		jobPost.JobDescription = c.GetString("JobDescription")
		jobPost.KeyResponsibility = c.GetString("KeyResponsibility")
		dbStatus := jobPost.InserInJobRequirement(jobPost)
		switch dbStatus {
		case true:

			fmt.Println("insert")
			return true
		case false:
			fmt.Println("incorrect")
			return false

		}
		return true
	}
	return true
}
func( c *JobController) DisplayJobDetails()[][]string{
	dbStatus,jobDetails :=models.DisplayJobDetails()
	var JobArray [][]string
	switch dbStatus{
	case true:
		for i :=0;i<len(jobDetails);i++{
			var JobTempArray []string
			JobTempArray =append(JobTempArray,jobDetails[i].JobId)
			JobTempArray =append(JobTempArray,jobDetails[i].EmpId)
			JobTempArray =append(JobTempArray,jobDetails[i].KeyResponsibility)
			JobTempArray =append(JobTempArray,jobDetails[i].JobDescription)
			JobTempArray =append(JobTempArray,jobDetails[i].KeySkills)
			JobTempArray =append(JobTempArray,jobDetails[i].MinExperience)
			JobTempArray =append(JobTempArray,jobDetails[i].Designation)
			JobTempArray =append(JobTempArray,jobDetails[i].JobLocation)
			JobTempArray =append(JobTempArray,jobDetails[i].MaxExperience)
			JobArray =append(JobArray,JobTempArray)
			}
	case false:
	}
	fmt.Println("kkkkk",JobArray)
	return JobArray
}
