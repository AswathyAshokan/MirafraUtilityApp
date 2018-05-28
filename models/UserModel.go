package  models

import (
	"log"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"fmt"
	"gopkg.in/mgo.v2/bson"

	"time"
	"os"
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
	uri := os.Getenv("MONGOLAB_URL")
	if uri == "" {
		fmt.Println("no connection string provided")
		os.Exit(1)
	}
	fmt.Println("jjj",uri)
	sess, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		fmt.Println("something happend")
		os.Exit(1)
	}
	defer sess.Close()
	hashedPassword, err := bcrypt.GenerateFromPassword(userDetails.Password, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return false
	}
	userDetails.Password = hashedPassword
	collection := sess.DB("mirafrautilityapp").C("user")

	fmt.Println("lllllllllllll",collection)
	// Insert
	if err := collection.Insert(userDetails); err != nil {
		return false
	}
	return true
}
//get birthday details

func GetBirthdayUsers () (bool,[]User){


	var BirthDay []User
	uri := os.Getenv("MONGOLAB_URL")
	if uri == "" {
		fmt.Println("no connection string provided")
		os.Exit(1)
	}

	sess, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		fmt.Println("something happend")
		os.Exit(1)
	}
	defer sess.Close()
	collection := sess.DB("mirafrautilityapp").C("user")
	err = collection.Find(nil).All(&BirthDay)
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
	uri := os.Getenv("MONGOLAB_URL")
	if uri == "" {
		fmt.Println("no connection string provided")
		os.Exit(1)
	}

	sess, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		fmt.Println("something happend")
		os.Exit(1)
	}
	defer sess.Close()
	collection := sess.DB("mirafrautilityapp").C("user")
	err = collection.Find(bson.M{"DateOfJoin": todayDate}).All(&NewJoiners)
	if err != nil {
		fmt.Println("error2",err)
		return  false,NewJoiners

	}
	return true,NewJoiners

}

