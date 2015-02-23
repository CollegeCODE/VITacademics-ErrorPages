package base

import (
  "fmt"
  "net/http"
)

func MaintenanceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, string(GetStatus("maintenance")))
}

func UnknownHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, string(GetStatus("unknown")))
}
