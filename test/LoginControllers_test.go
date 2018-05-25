package test
import (
	"testing"
	"MirafraUtilityApp/models"
	"fmt"
	"gopkg.in/mgo.v2"
	"os"
)




func TestLoginChecking(t *testing.T) {
	//dbLogin := models.Login{}
	//var (
	//	err error
	//)
	//fmt.Println("test code is running")
	//// Get post collection connection
	//c := newPostCollection()
	//defer c.Close()




	session,err:=mgo.Dial("127.0.0.1")

	if err != nil {
		panic(err)
	}

	defer session.Close()
	//var isDropMe = true
	//if isDropMe {
	//	err = session.DB("MirafraUtility").DropDatabase()
	//	if err != nil {
	//		panic(err)
	//	}
	//}

	session.SetMode(mgo.Monotonic, true)

	db:=session.DB("MirafraUtility").C("login")
	if db !=nil{

		fmt.Println("error",db)
	}

	if _, err := os.Stat("./testUploadImage/"); os.IsNotExist(err) {

		os.Mkdir("./testUploadImage/",os.ModePerm)
	}
	db =session.DB("MirafraUtility").C("user")
	if db !=nil{

		fmt.Println("error",db)
	}
	user :=models.User{}
	user.PositionType ="Senior Software Engineer"
	user.DateOfJoin ="15/05/2018"
	user.Dob		="15/05/1993"
	user.Email		= "mansi.m@gmail.com"
	user.EmpId		="EMP1755"
	user.Experience	="3.4"
	user.LastCompany ="techgenis"
	user.LastCompanyLocation	="Banglore"
	user.Manager	="sundar"
	user.Password		=[]byte("mirafra")
	user.FirstName		="Manasi"
	user.LastName  ="Pandya"
	user.Address ="Nikarth"
	user.City ="Ahmedabad"
	user.State="Gujarat"
	user.PinCode ="688540"
	user.PhoneNumber ="7736265636"
	user.Gender="female"
	user.JobStatus="fullTime"
	user.UploadPhoto	="./testUploadImage/aswathy.jpg"
	dbStatus :=user.InserIntoUser(user)
	switch dbStatus{
	case true:
		fmt.Println("insert")
	case false:
		fmt.Println("incorrect")

	}



	//performane award
	//db = session.DB("MirafraUtility").C("performance")
	//if db !=nil{
	//
	//	fmt.Println("error",db)
	//}
	//performance :=models.PerformanceModel{}
	//performance.EmpId="EMP1753"
	//performance.Name ="Aswathy Ashok"
	//performance.Comment="has done excellent work in bringing up ..."
	//dbStatus :=performance.InsertAward()
	//switch dbStatus{
	//case true:
	//	fmt.Println("insert performace")
	//case false:
	//	fmt.Println("incorrect performance")
	//
	//}

	//product insertion

	db = session.DB("MirafraUtility").C("product")
	if db !=nil{

		fmt.Println("error",db)
	}
	product :=models.Product{}
	product.Photo="./testUploadImage/aswathy.jpg"
	product.ContactNo="7736265636"
	product.Description="well furnished room ....."
	product.Price="4500"
	product.ProductId="p1300"
	product.ProductStatus=true
	product.EmpId="EMP1753"
	product.ProductName="room"
	dbStatus =product.InsertProductDetails()
	switch dbStatus {
	case true:
		fmt.Println("insert product")
	case false:
		fmt.Println("error in insertion")
	}



	db =session.DB("MirafraUtility").C("event")
	if db !=nil{

		fmt.Println("error",db)
	}
	event :=models.EventModel{}
	event.ContactNo="7736265636"
	event.EmailId="parvathy@mirafra.com"
	event.CreatedBy="parvathy"
	event.Date="15/05/2018"
	event.Location="Banglore"
	event.Description="docker"
	event.Time="12:30 pm"
	event.EventType="tech"
	event.EventName="docker implementation"
	event.EmpId="EMP1753"
	event.EventStatus=true
	dbStatus =event.InsertEvent()
	switch dbStatus{
	case true:
		fmt.Println("insert")
	case false:
		fmt.Println("incorrect")

	}




	// set default mongodb ID  and created date


	//dbLogin ="aswathy"
	//dbLogin.Password ="mirafra"
	//dbLogin.AccountType ="employee"

	// Insert post to mongodb
	//err = c.Session.Insert(&dbLogin)
	//if err != nil {
	//	panic(err)
	//}
	//err = db.Insert(dbLogin)
	//if err !=nil {
	//	fmt.Println("error login ",err)
	//}
	//
	//dbStatus :=dbLogin.CheckLogin(dbLogin.UserId,dbLogin.Password)
	//switch dbStatus {
	//case true:
	//	fmt.Println("true")
	//case false:
	//	fmt.Println("false")
	//}
	

//create collection for user


	
	 //dbUser := models.User{}
	 //dbUser.Emp_id ="EMP1753"
	 //dbUser.Name ="aswathy"
	 //dbUser.Category="employee"
	 //dbUser.Email="aswathyashok85@gmail.com"
	 //dbUser.Password="mirafra"
	 //dbUser.Dob="01/06/1993"
	 //dbUser.DateOfJoin="11/04/2018"
	// err = dbu.Insert(dbUser)
	//if err !=nil {
	//	fmt.Println("error login ",err)
	//}

	//create collection for user






//create collection for job requirements
 db=session.DB("MirafraUtility").C("jobRequirements")
	if db !=nil{

		fmt.Println("error",db)
	}
//dbJobRequirement :=models.Job_requirement{}
//
//dbJobRequirement.User_id ="EMP1753"
//dbJobRequirement.JobType ="python"
//dbJobRequirement.Job_Id ="py001"
//dbJobRequirement.Job_Description ="python developer with 1 year experience"
//dbJobRequirement.Client_Name ="cisco"
//err = db.Insert(dbJobRequirement)
//	if err !=nil {
//		fmt.Println("error login ",err)
//	}



	//counter for jobId





	//create collection for refferals
	db=session.DB("MirafraUtility").C("refferals")
	if db !=nil{

		fmt.Println("error",db)
	}
	//dbrefferal :=models.Refferals{}
	//dbrefferal.Reffreral_id ="ref001"
	//dbrefferal.JobType ="python"
	//dbrefferal.User_id ="EMP1753"
	//dbrefferal.Candidate_Name ="manasi"
	//dbrefferal.Candidate_PhoneNumber =7736265636
	//dbrefferal.Location ="chennai"
	//dbrefferal.Experience ="1"
	//dbrefferal.Resume_Path ="kkkkk"
	//err = db.Insert(dbrefferal)
	//if err !=nil {
	//	fmt.Println("error login ",err)
	//}


	//create collection for events

	db=session.DB("MirafraUtility").C("events")
	if db !=nil{

		fmt.Println("error",db)
	}
	//dbEvent := models.Events{}
	//dbEvent.Event_id="evt001"
	//dbEvent.User_id="EMP1753"
	//dbEvent.EventTitle="Team Outing"
	//dbEvent.Created_By="aswathy"
	//dbEvent.Date="17/04/2018"
	//dbEvent.Venue="bangalore"
	//dbEvent.EventName="event"
	//dbEvent.Description="new"
	//dbEvent.Time="10:00"
	//dbEvent.Event_status=false
	//dbEvent.Comment="new"
	//err = db.Insert(dbEvent)
	//if err !=nil {
	//	fmt.Println("error login ",err)
	//}


	//create event registration
	db =session.DB("MirafraUtility").C("event_registraion")
	if db !=nil{

		fmt.Println("error",db)
	}
	//dbEventRegister :=models.EventRegister{}
	//dbEventRegister.Event_id ="evt001"
	//dbEventRegister.User_id ="EMP1753"
	//err = db.Insert(dbEventRegister)
	//if err !=nil {
	//	fmt.Println("error login ",err)
	//}


	







	 






}

//create post method



