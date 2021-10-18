package requestresponse

import (
	"encoding/json"
	"fmt"
	"net/http"
	"uc_assignment/model"
)

func MakeResponse(w http.ResponseWriter, code int, resp interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	jsonResp, _ := json.Marshal(resp)
	fmt.Fprint(w, string(jsonResp))

}

type TripHistory struct {
	Ongoing   []*model.Trip
	Completed []*model.Trip
}
