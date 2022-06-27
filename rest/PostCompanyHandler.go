package rest

import (
	"P1/entity"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func PostCompany(rw http.ResponseWriter, r *http.Request) {
	reqBody := r.Body

	bodyBytes, err := ioutil.ReadAll(reqBody)
	if err != nil {
		fmt.Printf("Error ocurrend")
	}

	var company = entity.Company{}
	err = json.Unmarshal(bodyBytes, &company)
	if err != nil {
		fmt.Printf("Unmarshall error")
	}

	fmt.Println(company)
	rw.Write([]byte(bodyBytes))
}
