package lunchmodel

type Auto interface {
	Honk() string
}

type Truck struct {
	Auto
	Name string
}

type Trucks []Truck

func (t Truck) Honk() (s string) {
	return "I am a big truck. move"
}

func GetMyTrucks() (myAutos Trucks) {

	return Trucks{
		{
			Name: "Name1",
		},
	}
}
