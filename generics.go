package main

import (
	"fmt"
	"lunchmodel"
)

func main() {

	myObjs := lunchmodel.MyTypes{
		{
			MyProp: "some intialized value 1",
		}, {
			MyProp: "some intialized value 2",
		},
	}

	for _, myObj := range myObjs {
		fmt.Println(myObj.MyProp)

		myObj.MyFunc()

		fmt.Println(myObj.MyProp)

		fmt.Println(myObj.GetOther())
	}
}
