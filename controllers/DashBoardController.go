package controllers

import (
	"time"
	"fmt"
	"MirafraUtilityApp/models"
	"strings"
)


type  DashBoardController struct {
	BaseController
}
func (c *DashBoardController) DashBoardDisplay() ([][]string,[][]string,[][]string,[][]string) {

	//r := c.Ctx.Request
	//w :=c.Ctx.ResponseWriter

	currentTime := time.Now()

	today :=currentTime.Format("02/01/")
	fmt.Println("todays date",today)


	//fetching birthday details
	var BirthdayArray [][]string
	dbStatus,BirthdayDetails :=models.GetBirthdayUsers()
	switch dbStatus {
	case true:

		for i:=0;i<len(BirthdayDetails);i++{
			if (strings.Contains(BirthdayDetails[i].Dob,today )) {
				var BirthdayTempArray []string
				BirthdayTempArray=append(BirthdayTempArray,BirthdayDetails[i].FirstName+ " "+BirthdayDetails[i].LastName)
				BirthdayTempArray=append(BirthdayTempArray,BirthdayDetails[i].Email)
				BirthdayTempArray=append(BirthdayTempArray,BirthdayDetails[i].PositionType)
				BirthdayTempArray=append(BirthdayTempArray,BirthdayDetails[i].EmpId)
				BirthdayTempArray=append(BirthdayTempArray,BirthdayDetails[i].UploadPhoto)
				BirthdayArray =append(BirthdayArray,BirthdayTempArray)
			}
		}
		fmt.Println("details",BirthdayDetails)
	case false:
		fmt.Println("no details")

	}


	//fetching newJoiners details
	var NewJoinersArray [][]string
	dbStatus,NewJoinersDetails :=models.GetNewJoiners()
	switch dbStatus{
	case true:
		for i:=0;i<len(NewJoinersDetails);i++{
			var NewJoinersTempArray []string
			NewJoinersTempArray=append(NewJoinersTempArray,NewJoinersDetails[i].FirstName+" "+NewJoinersDetails[i].LastName)
			NewJoinersTempArray=append(NewJoinersTempArray,NewJoinersDetails[i].PositionType)
			NewJoinersTempArray=append(NewJoinersTempArray,NewJoinersDetails[i].Experience)
			NewJoinersTempArray=append(NewJoinersTempArray,NewJoinersDetails[i].LastCompany)
			NewJoinersTempArray=append(NewJoinersTempArray,NewJoinersDetails[i].LastCompanyLocation)
			NewJoinersTempArray=append(NewJoinersTempArray,NewJoinersDetails[i].City)
			NewJoinersTempArray=append(NewJoinersTempArray,NewJoinersDetails[i].State)
			NewJoinersTempArray=append(NewJoinersTempArray,NewJoinersDetails[i].Email)
			NewJoinersTempArray=append(NewJoinersTempArray,NewJoinersDetails[i].EmpId)
			NewJoinersArray =append(NewJoinersArray,NewJoinersTempArray)

		}
	}

	//performance

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

	//product

	dbStatus,productDetails :=models.DisplayProductDetails()
	var ProductArray [][]string
	switch dbStatus{
	case true:
		for i :=0;i<len(productDetails);i++{
			var ProductTempArray []string
			ProductTempArray =append(ProductTempArray,productDetails[i].ProductName)
			ProductTempArray =append(ProductTempArray,productDetails[i].ProductId)
			ProductTempArray =append(ProductTempArray,productDetails[i].Price)
			ProductTempArray =append(ProductTempArray,productDetails[i].Description)
			ProductTempArray =append(ProductTempArray,productDetails[i].ContactNo)
			ProductTempArray =append(ProductTempArray,productDetails[i].Photo)
			ProductArray =append(ProductArray,ProductTempArray)




		}
	case false:
	}


	return BirthdayArray,NewJoinersArray,performanceArray,ProductArray

}