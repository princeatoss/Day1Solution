package entity

type Employee struct {
	Name     string
	Age      int
	Location Location
	Gender   string
}

func (employee Employee) GetCity() string {
	return employee.Location.GetCity()
}
