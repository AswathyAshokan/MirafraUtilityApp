package controllers

import (
	"github.com/AswathyAshokan/MirafraUtilityApp/models"
	"fmt"
)

type EventController struct {
	BaseController
}
func (c *EventController)InsertEventDetails()bool{
	eventDetails :=models.EventModel{}
	r := c.Ctx.Request

	if r.Method == "POST" {
		eventDetails.EventName = c.GetString("eventName")
		eventDetails.Date = c.GetString("date")
		eventDetails.Location = c.GetString("location")
		eventDetails.Description = c.GetString("description")
		eventDetails.EventType = c.GetString("eventType")
		eventDetails.Comment = c.GetString("comment")
		eventDetails.Time = c.GetString("time")
		eventDetails.EventType = c.GetString("eventType")
		eventDetails.EventStatus = false
		dbStatus := eventDetails.InsertEvent()
		fmt.Println("jjjj",dbStatus)
		switch dbStatus {
		case true:
			fmt.Println("inserted")
			return true
		case false:
			fmt.Println("error")
			return false
		}
		return true
	}
	return true
}

func(c *EventController)DisplayEventDetails()[][]string{
	dbStatus,eventDetails :=models.DisplayEventDetails()
	var EventArray [][]string
	switch dbStatus{
	case true:
		for i:=0;i<len(eventDetails);i++{
			var EventTempArray []string
			EventTempArray=append(EventTempArray,eventDetails[i].EventType)
			EventTempArray=append(EventTempArray,eventDetails[i].Time)
			EventTempArray=append(EventTempArray,eventDetails[i].Comment)
			EventTempArray=append(EventTempArray,eventDetails[i].Description)
			EventTempArray=append(EventTempArray,eventDetails[i].Location)
			EventTempArray=append(EventTempArray,eventDetails[i].Date)
			EventTempArray=append(EventTempArray,eventDetails[i].CreatedBy)
			EventTempArray=append(EventTempArray,eventDetails[i].EmailId)
			EventTempArray=append(EventTempArray,eventDetails[i].ContactNo)
			EventTempArray=append(EventTempArray,eventDetails[i].EventId)

			EventArray =append(EventArray,EventTempArray)


		}
	}
  return EventArray
}
func(c *EventController)UpdateEvent()bool{
	eventDetails :=models.EventModel{}

	 eventId:= c.Ctx.Input.Param(":eventId")
	 action :=c.Ctx.Input.Param(":action")
	eventDetails.EventId=eventId
	eventDetails.EventComment =c.GetString("EventComment")

	if action=="accept"{
		eventDetails.EventStatus=true
	}else{
		eventDetails.EventStatus=false
	}
	dbStatus :=eventDetails.EventUpdate()
		switch dbStatus {
		case true:
			return true
		case false:
			return false

		}
	return true

}