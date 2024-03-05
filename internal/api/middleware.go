package api

import (
	"bytes"
	"encoding/json"
	"io"
	"msa-app/internal/models"
	"msa-app/pkg/handler"
	"net/http"
	"strconv"
)

func Middleware(w http.ResponseWriter, r *http.Request) {
	var routeingInfo models.RoutingInfo
	body, requestErr := io.ReadAll(r.Body)
	if handler.CheckHttpError(w, requestErr, "Body Err") {
		return
	}

	unmarshalErr := json.Unmarshal(body, &routeingInfo)
	if handler.CheckHttpError(w, unmarshalErr, "UnmarshalErr request") {
		return
	}

	if len(routeingInfo.Route) == 1 {
		body, _ := io.ReadAll(r.Body)
		resp, err := http.Post("http://localhost:"+strconv.Itoa(routeingInfo.Route[0])+"/api/v1/run", "application/json", bytes.NewBuffer(body))
		if handler.CheckHttpError(w, err, "Endpoint request failed") {
			return
		}

		defer resp.Body.Close()
	}
}
