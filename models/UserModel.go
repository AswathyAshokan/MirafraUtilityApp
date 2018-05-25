package  models

import (
	"log"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"fmt"
	"gopkg.in/mgo.v2/bson"

	"time"
)

type User struct{

	//personal details
	FirstName		string
	LastName		string
	PhoneNumber		string
	Address			string
	State			string
	City			string
	PinCode			string
	Gender			string

	//official details

	EmpId			string
	PositionType	string
	JobStatus		string
	Email			string
	Password		[]byte
	Dob				string
	DateOfJoin 		string
	LastCompanyLocation	string
	Manager		string
	Experience	string
	LastCompany	string
	UploadPhoto	string

}
func (userDetails *User)InserIntoUser( user User) (bool) {
	session,err:=mgo.Dial("127.0.0.1")

	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	hashedPassword, err := bcrypt.GenerateFromPassword(userDetails.Password, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return false
	}
	userDetails.Password = hashedPassword
	c := session.DB("MirafraUtility").C("user")

	// Insert
	if err := c.Insert(userDetails); err != nil {
		return false
	}
	return true
}
//get birthday details

func GetBirthdayUsers () (bool,[]User){


	var BirthDay []User
	session,err:=mgo.Dial("127.0.0.1")

	if err != nil {
		fmt.Println("error1",err)
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("MirafraUtility").C("user")
	err = c.Find(nil).All(&BirthDay)
	if err != nil {
		fmt.Println("error2",err)
		return  false,BirthDay

	}
	fmt.Println("birthday details struct",BirthDay)
	return true,BirthDay

}
//get NewJoiners details

func GetNewJoiners() (bool,[]User){
	currentTime := time.Now()
	todayDate :=currentTime.Format("02/01/2006")
	var NewJoiners []User
	session,err:=mgo.Dial("127.0.0.1")

	if err != nil {
		fmt.Println("error1",err)
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("MirafraUtility").C("user")
	err = c.Find(bson.M{"DateOfJoin": todayDate}).All(&NewJoiners)
	if err != nil {
		fmt.Println("error2",err)
		return  false,NewJoiners

	}
	return true,NewJoiners

}

