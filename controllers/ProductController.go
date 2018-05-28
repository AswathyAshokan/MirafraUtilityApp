package controllers

import (
	"github.com/AswathyAshokan/MirafraUtilityApp/models"
	"strconv"
	"time"
	"os"
	"log"
	"fmt"
	"io"
)

type ProductController struct {
	BaseController
}

func (c *ProductController) InserProductDetails()bool{
	product :=models.Product{}
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	if r.Method == "POST" {
		product.ProductName = c.GetString("ProductName")
		product.ContactNo = c.GetString("ContactNo")
		product.Description = c.GetString("Description")
		product.Price = c.GetString("Price")
		product.ProductComment =c.GetString("ProductComment")
		product.ProductStatus = false
		//photoUploading

		msec := strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
		//creating a folder for uploading
		if _, err := os.Stat("./testUploadImage/"); os.IsNotExist(err) {

			os.Mkdir("./testUploadImage/", os.ModePerm)
		}
		file, header, err := r.FormFile("uploadfile")
		if err != nil {
			log.Println("uploading error", err)
			//http.Error(w, "error in uploading file", http.StatusInternalServerError)

		} else {
			f, err := os.OpenFile("./testUploadImage/"+msec+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println("image 4 error", err)

			}
			fmt.Println("jst")
			imagePath := "./testUploadImage" + msec + header.Filename

			io.Copy(f, file)
			defer file.Close()
			product.Photo = imagePath
			fmt.Fprintf(w, "file  uploaded")

		}

		dbStatus := product.InsertProductDetails()
		switch dbStatus {
		case true:
			return true
			fmt.Println("insert product")
		case false:
			fmt.Println("error in insertion")
		   return false
		}
		return true
	}
return true
}

func( c *ProductController) DisplayProductDetails()[][]string{
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
	return ProductArray
}

func(c *ProductController)UpdateProduct()bool{
	product :=models.Product{}

	productId:= c.Ctx.Input.Param(":productId")
	action :=c.Ctx.Input.Param(":action")
	product.ProductId=productId
	product.ProductComment =c.GetString("ProductComment")

	if action=="accept"{
		product.ProductStatus=true
	}else{
		product.ProductStatus=false
	}
	dbStatus :=product.ProductUpdate()
	switch dbStatus {
	case true:
		return true
	case false:
		return false

	}
	return true

}