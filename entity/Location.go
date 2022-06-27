package entity

type Location struct {
	Longitude int
	Latitude  int
	city      string
}

func (location Location) GetCity() string {
	return location.city
}
