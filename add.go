package main

import (
	"log"
	"strconv"

	"github.com/kataras/iris/v12"
)

// post >> /customer/
func (r *Repository) addCustomerData(ctx iris.Context) {
	inputName := ctx.PostValue("name")
	inputAge, _ := strconv.ParseUint(ctx.PostValue("age"), 10, 64)
	inputEmail := ctx.PostValue("email")

	customer := Customer{
		Name:  inputName,
		Age:   inputAge,
		Email: inputEmail,
	}

	err := r.DB.Create(&customer).Error
	if err != nil {
		log.Print("System:", err)
	} else {
		ctx.JSON(iris.Map{"message": "new customer added successfully"})
		log.Print("System: new customer added successfully")
	}
}
