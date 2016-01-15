package lunchmodel

import "testing"

func TestSideEffect(t *testing.T) {
	myObj := MyType{}

	//	getText = mockGetText

	myObj.MyFunc()

	//	if myObj.myOtherProp != "mock call" {

	if myObj.myOtherProp != "who doesn't love side effects?!?!" {
		t.Errorf("myOtherProp should be set", myObj)
	}
}

func mockGetText() (text string) {

	return "mock call"
}
