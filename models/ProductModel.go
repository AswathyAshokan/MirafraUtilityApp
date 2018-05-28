package models

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"os"
	"time"
	"math/rand"
)

type Product struct {
	EmpId			string
	ProductName		string
	Description		string
	Price			string
	ContactNo		string
	Photo			string
	ProductStatus	bool
	ProductId		string
	ProductComment	string
}
type ProductCounter struct {
	ProductId  string
	seq		int64
}
func(product Product)InsertProductDetails()bool{

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
	//productCounter :=ProductCounter{}
	//productCounter.ProductId ="ProductId"
	//productCounter.seq= 0
	//collection := sess.DB("mirafrautilityapp").C("productCounter")
	//if collection !=nil{

		//fmt.Println("error",collection)
	//}
	//err = collection.Insert(productCounter)
	//if err !=nil {
		//fmt.Println("error login ",err)
	//}

	//updating the sequence
	//err =collection.Find(nil).All(productCounter)
	//counter :=productCounter.seq+1
	//err = collection.Update(bson.M{"ProductId":productCounter.ProductId }, bson.M{"$set": bson.M{"seq": counter}})

	var r *rand.Rand

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	const chars = "012345abcdefghijklmnopqrstuvwxyz6789"
	result := make([]byte, 8)
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}

	fmt.Println("gfghggghg", string(result))
	product.ProductId="P"+ string(result)

	collection := sess.DB("mirafrautilityapp").C("product")

	if err := collection.Insert(product); err != nil {
		return false
	}
	return true
}
func DisplayProductDetails()(bool,[]Product){
	var Product []Product
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
	collection := sess.DB("mirafrautilityapp").C("product")
	err = collection.Find(bson.M{"productstatus":false}).All(&Product)
	if err != nil {
		fmt.Println("error2",err)
		return  false,Product

	}
	fmt.Println("product details struct",Product)
	return true,Product

}

func (product Product)ProductUpdate()bool{

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

	collection := sess.DB("mirafrautilityapp").C("product")
	if collection !=nil{

		fmt.Println("error",collection)
	}

	//updating the event

	err = collection.Update(bson.M{"ProductId":product.ProductId }, bson.M{"$set": bson.M{"ProductStatus": product.ProductStatus,"ProductComment":product.ProductComment}})
	if err !=nil{
		return false
	}
	return true
}