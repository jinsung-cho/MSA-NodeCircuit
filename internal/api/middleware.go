package api

import (
	"bytes"
	"encoding/json"
	"io"
	"msa-app/internal/models"
	"msa-app/pkg/handler"
	"net/http"
)

func Middleware(w http.ResponseWriter, r *http.Request) {
	body, requestErr := io.ReadAll(r.Body)
	if handler.CheckHttpError(w, requestErr, "Body Err") {
		return
	}

	var routeingInfo models.RoutingInfo
	unmarshalErr := json.Unmarshal(body, &routeingInfo)
	if handler.CheckHttpError(w, unmarshalErr, "UnmarshalErr request") {
		return
	}

	var jsonData []byte
	var marshalErr error
	var url string
	if len(routeingInfo.Route) == 1 {
		url = "http://localhost:" + routeingInfo.Route[0] + "/api/end"
		routeingInfo = models.RoutingInfo{
			Route: []string{},
		}
		jsonData, marshalErr = json.Marshal(routeingInfo)
	} else {
		nextPort, remainRoutingInfo := routeingInfo.Route[0], routeingInfo.Route[1:]
		routeingInfo.Route = remainRoutingInfo
		url = "http://localhost:" + nextPort + "/api/mid"
		jsonData, marshalErr = json.Marshal(routeingInfo)
	}

	if handler.CheckHttpError(w, marshalErr, "MarshalErr request") {
		return
	}

	postR, postErr := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if handler.CheckHttpError(w, postErr, "post request failed") {
		return
	}
	defer postR.Body.Close()

	var routeHistory models.RoutingHistory
	postBody, postRequestErr := io.ReadAll(postR.Body)
	if handler.CheckHttpError(w, postRequestErr, "post Body Err") {
		return
	}

	postUnmarshalErr := json.Unmarshal(postBody, &routeHistory)
	if handler.CheckHttpError(w, postUnmarshalErr, "post MarshalErr request") {
		return
	}

	routeHistory.History = append(routeHistory.History, myPort)
	postJsonData, postMarshalErr := json.Marshal(routeHistory)
	if handler.CheckHttpError(w, postMarshalErr, "post MarshalErr request") {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(postJsonData)
}
