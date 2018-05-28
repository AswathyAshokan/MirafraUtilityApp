package models

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"os"
	"math/rand"
	"time"
)

type JobRequirement struct{

	JobId				string
	EmpId				string
	Designation			string
	JobLocation			string
	KeySkills			string
	JobDescription		string
	KeyResponsibility	string
	MinExperience		string
	MaxExperience		string



}

type JobCounter struct {
	JobId  string
	seq		int64
}




func (jobDetails *JobRequirement)InserInJobRequirement( job JobRequirement) (bool) {

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

	sess.SetSafe(&mgo.Safe{})
	var r *rand.Rand

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	const chars = "012345abcdefghijklmnopqrstuvwxyz6789"
	result := make([]byte, 8)
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}

	fmt.Println("gfghggghg", string(result))
	collection := sess.DB("mirafrautilityapp").C("jobRequirements")
	if collection !=nil{

		fmt.Println("error",collection)
	}
	job.JobId="Job"+ string(result)
	err = collection.Insert(job)
	if err !=nil {
		fmt.Println("error login ",err)
	}

return true
}
func DisplayJobDetails()(bool,[]JobRequirement){
	var Job []JobRequirement
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
	collection := sess.DB("mirafrautilityapp").C("jobRequirements")
	err = collection.Find(nil).All(&Job)
	if err != nil {
		fmt.Println("error2",err)
		return  false,Job

	}
	fmt.Println("job details struct",Job)
	return true,Job

}
