package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log" // https://pkg.go.dev/log
	"net/http"

	_ "github.com/lib/pq"
)

/**
 * \brief
 */
type OrderPlaceHandler struct {
	dbconn *DatabaseConnector
}

type OrderPlaceResponse struct {
	StatusCode ErrorCode `json:"statusCode"`
	Message    string    `json:"message"`
}

func (r *OrderPlaceResponse) write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	if bytes, err := json.Marshal(r); err != nil {
		// @NOTE Should NEVER happen !!!
		io.WriteString(w, fmt.Sprintf("{\"statusCode\":%d,\"message\":\"%s\"}", r.StatusCode, r.Message))
	} else {
		log.Println(string(bytes))
		io.Writer.Write(w, bytes)
	}
}

func OrderPlaceResponse_Success() OrderPlaceResponse {
	return OrderPlaceResponse{
		StatusCode: SUCCEEDED,
		Message:    "OK",
	}
}
func OrderPlaceResponse_Error(err error) OrderPlaceResponse {
	return OrderPlaceResponse{
		StatusCode: FAILED,
		Message:    err.Error(),
	}
}

func (h OrderPlaceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("HTTP Request `OrderPlace`")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var resp OrderPlaceResponse
	defer resp.write(w)

	if auth, authFound := r.Header["Authorization"]; authFound {
		log.Println("Authorization for", auth)
		resp = OrderPlaceResponse_Success()
	} else {
		resp = OrderPlaceResponse_Error(fmt.Errorf("Auth failed"))
	}
}
