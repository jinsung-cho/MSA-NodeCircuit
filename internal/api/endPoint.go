package api

import (
	"encoding/json"
	"io"
	"msa-app/internal/models"
	"msa-app/pkg/handler"
	"net/http"
)

func EndPoint(w http.ResponseWriter, r *http.Request) {
	var routeingInfo models.RoutingInfo
	body, requestErr := io.ReadAll(r.Body)
	if handler.CheckHttpError(w, requestErr, "Body Err") {
		return
	}

	unmarshalErr := json.Unmarshal(body, &routeingInfo)
	if handler.CheckHttpError(w, unmarshalErr, "UnmarshalErr request") {
		return
	}

}
