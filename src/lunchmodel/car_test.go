package lunchmodel

import "testing"

func TestSideEffect(t *testing.T) {
	myCar := Car{}

	//	getText = mockGetText

	myCar.MyFunc()

	//	if myObj.myOtherProp != "mock call" {

	if myCar.myOtherProp != "who doesn't love side effects?!?!" {
		t.Errorf("myOtherProp should be set", myCar)
	}
}

func mockGetText() (text string) {

	return "mock call"
}
