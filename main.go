package main

import (
	"P1/entity"
	"P1/rest"
	"P1/services"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	companies := entity.InitCompanies()

	var name string
	fmt.Print("Insert company name: ")
	fmt.Scanln(&name)
	var city string
	fmt.Print("Insert employees' city: ")
	fmt.Scanln(&city)
	fmt.Println("There are:", services.Count(companies, name, city), "employees that work at", name, ", that live in", city)

	router := chi.NewRouter()
	router.Route("/v1", func(r chi.Router) {
		r.Get("/welcome", rest.Welcome)
		r.Get("/companies", rest.GetCompany)
		r.Post("/companies", rest.PostCompany)

	})
	http.ListenAndServe(":9090", router)
}
