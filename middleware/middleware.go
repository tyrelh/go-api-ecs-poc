package middleware

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"time"

	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
)

// this middleware logs the HTTP method, request path and the time taken to process the request
func LogEndpoint(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

// this middleware logs the request and response
func LogRequestResponse(f strictnethttp.StrictHTTPHandlerFunc, operationID string) strictnethttp.StrictHTTPHandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (response interface{}, err error) {
		if request != nil || !reflect.DeepEqual(request, reflect.Zero(reflect.TypeOf(request)).Interface()) {
			log.Println("Operation ID: " + operationID)
			jsonResponse, jsonErr := json.Marshal(request)
			if jsonErr != nil {
				http.Error(w, jsonErr.Error(), http.StatusInternalServerError)
				return nil, jsonErr
			}
			log.Println("Request: " + string(jsonResponse))
		}

		serverResponse, err := f(ctx, w, r, request)

		if err == nil {
			// w.Header().Set("Content-Type", "application/json")
			jsonResponse, jsonErr := json.Marshal(serverResponse)
			if jsonErr != nil {
				http.Error(w, jsonErr.Error(), http.StatusInternalServerError)
				return nil, jsonErr
			}
			log.Println("Origin Response: " + string(jsonResponse))
		}

		return serverResponse, err
	}
}

//hello
