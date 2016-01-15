package main

import (
	"fmt"
	"lunchmodel"
)

func main() {

	myCars := lunchmodel.GetMyCars()

	for _, myCar := range myCars {
		fmt.Println(myCar.MyProp)

		myCar.MyFunc()

		fmt.Println(myCar.MyProp)

		fmt.Println(myCar.GetOther())
	}

	myTrucks := lunchmodel.GetMyTrucks()

	autos := make([]lunchmodel.Auto, 3)
	autos[0], autos[1], autos[2] = myCars[0], myCars[1], myTrucks[0]

	for _, value := range autos {
		fmt.Println(value.Honk())
	}
}
