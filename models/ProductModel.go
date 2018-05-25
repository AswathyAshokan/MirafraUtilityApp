package models

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"strconv"
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

	session,err:=mgo.Dial("127.0.0.1")

	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	productCounter :=ProductCounter{}
	productCounter.ProductId ="ProductId"
	productCounter.seq= 0
	db :=session.DB("MirafraUtility").C("productCounter")
	if db !=nil{

		fmt.Println("error",db)
	}
	err = db.Insert(productCounter)
	if err !=nil {
		fmt.Println("error login ",err)
	}

	//updating the sequence
	err =db.Find(nil).All(productCounter)
	counter :=productCounter.seq+1
	err = db.Update(bson.M{"ProductId":productCounter.ProductId }, bson.M{"$set": bson.M{"seq": counter}})


	product.ProductId="p00"+ strconv.FormatInt(counter, 10)

	c := session.DB("MirafraUtility").C("product")

	if err := c.Insert(product); err != nil {
		return false
	}
	return true
}
func DisplayProductDetails()(bool,[]Product){
	var Product []Product
	session,err:=mgo.Dial("127.0.0.1")

	if err != nil {
		fmt.Println("error1",err)
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("MirafraUtility").C("product")
	err = c.Find(bson.M{"ProductStatus": true}).All(&Product)
	if err != nil {
		fmt.Println("error2",err)
		return  false,Product

	}
	fmt.Println("product details struct",Product)
	return true,Product

}

func (product Product)ProductUpdate()bool{

	session,err:=mgo.Dial("127.0.0.1")

	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	db :=session.DB("MirafraUtility").C("product")
	if db !=nil{

		fmt.Println("error",db)
	}

	//updating the event

	err = db.Update(bson.M{"ProductId":product.ProductId }, bson.M{"$set": bson.M{"ProductStatus": product.ProductStatus,"ProductComment":product.ProductComment}})
	if err !=nil{
		return false
	}
	return true
}