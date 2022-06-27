package rest

import "net/http"

func GetCompany(rw http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	rw.Write([]byte(name))
}
