package api

import (
	"encoding/json"
	"msa-app/internal/models"
	"msa-app/pkg/handler"
	"net/http"
)

func EndPoint(w http.ResponseWriter, r *http.Request) {
	var routeHistory models.RoutingHistory

	routeHistory.History = append(routeHistory.History, myPort)
	postJsonData, postMarshalErr := json.Marshal(routeHistory)
	if handler.CheckHttpError(w, postMarshalErr, "MarshalErr request") {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(postJsonData)
}
