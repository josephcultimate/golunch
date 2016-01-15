package main

import (
	"fmt"
	"lunchmodel"
)

func main_generic() {

	myObjs := lunchmodel.GetMyTypes()

	for _, myObj := range myObjs {
		fmt.Println(myObj.MyProp)

		myObj.MyFunc()

		fmt.Println(myObj.MyProp)

		fmt.Println(myObj.GetOther())
	}
}
