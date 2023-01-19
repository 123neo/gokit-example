package transport

import (
	"context"
	"encoding/json"
	"gokit-example/pkg/gokit/endpoint"
	"gokit-example/pkg/gokit/service"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"go.uber.org/zap"
)

type errorer interface {
	error() error
}

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case service.ErrUserAlreadyExists:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func DecodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.CreateUserRequest

	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		log.Println("Error: " + e.Error())
		return nil, e
	}
	return req, nil

}

func EncodeCreateUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		log.Println("Response")
		log.Println(e.error())
		encodeErrorResponse(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func NewUserHTTPHandler(endpoints endpoint.UserSvcSet, logger *zap.Logger) http.Handler {
	m := http.NewServeMux()
	m.Handle("/createUser", httptransport.NewServer(
		endpoints.CreateUserEndpoint,
		DecodeCreateUserRequest,
		EncodeCreateUserResponse,
	))
	return m
}
