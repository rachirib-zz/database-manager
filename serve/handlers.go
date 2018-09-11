package serve

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Database Manager by Ac2 Team")
}

func DatabaseIndex(w http.ResponseWriter, r *http.Request) {
	databases := Databases{
		Database{Name: "Clair"},
		Database{Name: "Quay"},
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(databases); err != nil {
		panic(err)
	}
}

func DatabaseShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	databaseId := vars["databaseId"]
	fmt.Fprintln(w, "Database show: ", databaseId)
}
