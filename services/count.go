package services

import "P1/entity"

func Count(companies []entity.Company, name string, city string) int {

	no := 0

	for _, c := range companies {
		if c.Name == name {
			for _, e := range c.Employees {
				if e.Location.GetCity() == city {
					no++
				}
			}
		}
	}

	return no
}
