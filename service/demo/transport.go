package demo

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

type CallRequest struct {
	Name string `json:"name"`
}

type CallResponse struct {
	Reply string `json:"reply"`
}

func decodeCallRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request CallRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeCallResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func NewHTTPServer(svc Service) http.Handler {
	callHandler := httptransport.NewServer(
		makeCallEndpoint(svc),
		decodeCallRequest,
		encodeCallResponse,
	)
	r := mux.NewRouter()
	r.Handle("/call", callHandler).Methods("POST")
	return r
}
