package lunchmodel

import (
	"io/ioutil"
	"strings"
)

type Car struct {
	Auto
	MyProp      string
	myOtherProp string
}

// Be Explicit!!!!
type Cars []Car

func (myCar *Car) MyFunc() {

	myCar.myOtherProp = getText()

	myCar.MyProp = "Anyone can get to me"
}

func (myCar *Car) GetOther() (val string) {

	return myCar.myOtherProp
}

func (t Car) Honk() (s string) {
	return "I am a car please move"
}

func GetMyCars() (myCars Cars) {

	return Cars{
		{
			MyProp: "some intialized value 1",
		}, {
			MyProp: "some intialized value 2",
		},
	}
}

var getText = func() (text string) {

	data, _ := ioutil.ReadFile("C:\\Projects\\golunch\\sometext.txt")

	return strings.TrimSpace(string(data))
}
