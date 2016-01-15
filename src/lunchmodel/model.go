package lunchmodel

import (
	"io/ioutil"
	"strings"
)

type MyType struct {
	MyProp      string
	myOtherProp string
}

// Be Explicit!!!!
type MyTypes []MyType

func (myType *MyType) MyFunc() {

	myType.myOtherProp = getText()

	myType.MyProp = "Anyone can get to me"
}

func (myType *MyType) GetOther() (val string) {

	return myType.myOtherProp
}

func GetMyTypes() (myTypes MyTypes) {

	return MyTypes{
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
