package lunchmodel

type MyType struct {
	MyProp      string
	myOtherProp string
}

// Be Explicit!!!!
type MyTypes []MyType

func (myType *MyType) MyFunc() {

	myType.myOtherProp = "who doesn't love side effects?!?!"

	myType.MyProp = "Anyone can get to me"
}

func (myType *MyType) GetOther() (val string) {

	return myType.myOtherProp
}
