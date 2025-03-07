//go:build go1.22

// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/oapi-codegen/runtime"
	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
)

// RewardRequest defines model for RewardRequest.
type RewardRequest struct {
	Brand        *string  `json:"brand,omitempty"`
	Currency     *string  `json:"currency,omitempty"`
	Denomination *float32 `json:"denomination,omitempty"`
}

// RewardResponse defines model for RewardResponse.
type RewardResponse struct {
	Brand        *string  `json:"brand,omitempty"`
	Currency     *string  `json:"currency,omitempty"`
	Denomination *float32 `json:"denomination,omitempty"`
	Id           *int     `json:"id,omitempty"`
}

// PostGoRewardJSONRequestBody defines body for PostGoReward for application/json ContentType.
type PostGoRewardJSONRequestBody = RewardRequest

// PutGoRewardIdJSONRequestBody defines body for PutGoRewardId for application/json ContentType.
type PutGoRewardIdJSONRequestBody = RewardRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get a list of active rewards for your account
	// (GET /go/reward)
	GetGoReward(w http.ResponseWriter, r *http.Request)
	// Create a reward
	// (POST /go/reward)
	PostGoReward(w http.ResponseWriter, r *http.Request)
	// Delete a Reward by ID
	// (DELETE /go/reward/{id})
	DeleteGoRewardId(w http.ResponseWriter, r *http.Request, id int)
	// Get a Reward by ID
	// (GET /go/reward/{id})
	GetGoRewardId(w http.ResponseWriter, r *http.Request, id int)
	// Update a Reward by ID
	// (PUT /go/reward/{id})
	PutGoRewardId(w http.ResponseWriter, r *http.Request, id int)
	// Get the health of the API
	// (GET /go/system/health)
	GetGoSystemHealth(w http.ResponseWriter, r *http.Request)
	// Get the version of the API
	// (GET /go/system/version)
	GetGoSystemVersion(w http.ResponseWriter, r *http.Request)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetGoReward operation middleware
func (siw *ServerInterfaceWrapper) GetGoReward(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetGoReward(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// PostGoReward operation middleware
func (siw *ServerInterfaceWrapper) PostGoReward(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostGoReward(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// DeleteGoRewardId operation middleware
func (siw *ServerInterfaceWrapper) DeleteGoRewardId(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", r.PathValue("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteGoRewardId(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetGoRewardId operation middleware
func (siw *ServerInterfaceWrapper) GetGoRewardId(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", r.PathValue("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetGoRewardId(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// PutGoRewardId operation middleware
func (siw *ServerInterfaceWrapper) PutGoRewardId(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", r.PathValue("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutGoRewardId(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetGoSystemHealth operation middleware
func (siw *ServerInterfaceWrapper) GetGoSystemHealth(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetGoSystemHealth(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetGoSystemVersion operation middleware
func (siw *ServerInterfaceWrapper) GetGoSystemVersion(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetGoSystemVersion(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{})
}

// ServeMux is an abstraction of http.ServeMux.
type ServeMux interface {
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type StdHTTPServerOptions struct {
	BaseURL          string
	BaseRouter       ServeMux
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, m ServeMux) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseRouter: m,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, m ServeMux, baseURL string) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseURL:    baseURL,
		BaseRouter: m,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options StdHTTPServerOptions) http.Handler {
	m := options.BaseRouter

	if m == nil {
		m = http.NewServeMux()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	m.HandleFunc("GET "+options.BaseURL+"/go/reward", wrapper.GetGoReward)
	m.HandleFunc("POST "+options.BaseURL+"/go/reward", wrapper.PostGoReward)
	m.HandleFunc("DELETE "+options.BaseURL+"/go/reward/{id}", wrapper.DeleteGoRewardId)
	m.HandleFunc("GET "+options.BaseURL+"/go/reward/{id}", wrapper.GetGoRewardId)
	m.HandleFunc("PUT "+options.BaseURL+"/go/reward/{id}", wrapper.PutGoRewardId)
	m.HandleFunc("GET "+options.BaseURL+"/go/system/health", wrapper.GetGoSystemHealth)
	m.HandleFunc("GET "+options.BaseURL+"/go/system/version", wrapper.GetGoSystemVersion)

	return m
}

type GetGoRewardRequestObject struct {
}

type GetGoRewardResponseObject interface {
	VisitGetGoRewardResponse(w http.ResponseWriter) error
}

type GetGoReward200JSONResponse struct {
	Rewards *[]RewardResponse `json:"rewards,omitempty"`
}

func (response GetGoReward200JSONResponse) VisitGetGoRewardResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostGoRewardRequestObject struct {
	Body *PostGoRewardJSONRequestBody
}

type PostGoRewardResponseObject interface {
	VisitPostGoRewardResponse(w http.ResponseWriter) error
}

type PostGoReward201JSONResponse RewardResponse

func (response PostGoReward201JSONResponse) VisitPostGoRewardResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type DeleteGoRewardIdRequestObject struct {
	Id int `json:"id"`
}

type DeleteGoRewardIdResponseObject interface {
	VisitDeleteGoRewardIdResponse(w http.ResponseWriter) error
}

type DeleteGoRewardId204Response struct {
}

func (response DeleteGoRewardId204Response) VisitDeleteGoRewardIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type DeleteGoRewardId404Response struct {
}

func (response DeleteGoRewardId404Response) VisitDeleteGoRewardIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetGoRewardIdRequestObject struct {
	Id int `json:"id"`
}

type GetGoRewardIdResponseObject interface {
	VisitGetGoRewardIdResponse(w http.ResponseWriter) error
}

type GetGoRewardId200JSONResponse RewardResponse

func (response GetGoRewardId200JSONResponse) VisitGetGoRewardIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PutGoRewardIdRequestObject struct {
	Id   int `json:"id"`
	Body *PutGoRewardIdJSONRequestBody
}

type PutGoRewardIdResponseObject interface {
	VisitPutGoRewardIdResponse(w http.ResponseWriter) error
}

type PutGoRewardId200JSONResponse RewardRequest

func (response PutGoRewardId200JSONResponse) VisitPutGoRewardIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetGoSystemHealthRequestObject struct {
}

type GetGoSystemHealthResponseObject interface {
	VisitGetGoSystemHealthResponse(w http.ResponseWriter) error
}

type GetGoSystemHealth200JSONResponse struct {
	Region *string `json:"region,omitempty"`
	Status *string `json:"status,omitempty"`
}

func (response GetGoSystemHealth200JSONResponse) VisitGetGoSystemHealthResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetGoSystemVersionRequestObject struct {
}

type GetGoSystemVersionResponseObject interface {
	VisitGetGoSystemVersionResponse(w http.ResponseWriter) error
}

type GetGoSystemVersion200JSONResponse struct {
	Version *string `json:"version,omitempty"`
}

func (response GetGoSystemVersion200JSONResponse) VisitGetGoSystemVersionResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Get a list of active rewards for your account
	// (GET /go/reward)
	GetGoReward(ctx context.Context, request GetGoRewardRequestObject) (GetGoRewardResponseObject, error)
	// Create a reward
	// (POST /go/reward)
	PostGoReward(ctx context.Context, request PostGoRewardRequestObject) (PostGoRewardResponseObject, error)
	// Delete a Reward by ID
	// (DELETE /go/reward/{id})
	DeleteGoRewardId(ctx context.Context, request DeleteGoRewardIdRequestObject) (DeleteGoRewardIdResponseObject, error)
	// Get a Reward by ID
	// (GET /go/reward/{id})
	GetGoRewardId(ctx context.Context, request GetGoRewardIdRequestObject) (GetGoRewardIdResponseObject, error)
	// Update a Reward by ID
	// (PUT /go/reward/{id})
	PutGoRewardId(ctx context.Context, request PutGoRewardIdRequestObject) (PutGoRewardIdResponseObject, error)
	// Get the health of the API
	// (GET /go/system/health)
	GetGoSystemHealth(ctx context.Context, request GetGoSystemHealthRequestObject) (GetGoSystemHealthResponseObject, error)
	// Get the version of the API
	// (GET /go/system/version)
	GetGoSystemVersion(ctx context.Context, request GetGoSystemVersionRequestObject) (GetGoSystemVersionResponseObject, error)
}

type StrictHandlerFunc = strictnethttp.StrictHTTPHandlerFunc
type StrictMiddlewareFunc = strictnethttp.StrictHTTPMiddlewareFunc

type StrictHTTPServerOptions struct {
	RequestErrorHandlerFunc  func(w http.ResponseWriter, r *http.Request, err error)
	ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}}
}

func NewStrictHandlerWithOptions(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: options}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
	options     StrictHTTPServerOptions
}

// GetGoReward operation middleware
func (sh *strictHandler) GetGoReward(w http.ResponseWriter, r *http.Request) {
	var request GetGoRewardRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetGoReward(ctx, request.(GetGoRewardRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetGoReward")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetGoRewardResponseObject); ok {
		if err := validResponse.VisitGetGoRewardResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostGoReward operation middleware
func (sh *strictHandler) PostGoReward(w http.ResponseWriter, r *http.Request) {
	var request PostGoRewardRequestObject

	var body PostGoRewardJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostGoReward(ctx, request.(PostGoRewardRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostGoReward")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostGoRewardResponseObject); ok {
		if err := validResponse.VisitPostGoRewardResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// DeleteGoRewardId operation middleware
func (sh *strictHandler) DeleteGoRewardId(w http.ResponseWriter, r *http.Request, id int) {
	var request DeleteGoRewardIdRequestObject

	request.Id = id

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteGoRewardId(ctx, request.(DeleteGoRewardIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteGoRewardId")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(DeleteGoRewardIdResponseObject); ok {
		if err := validResponse.VisitDeleteGoRewardIdResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetGoRewardId operation middleware
func (sh *strictHandler) GetGoRewardId(w http.ResponseWriter, r *http.Request, id int) {
	var request GetGoRewardIdRequestObject

	request.Id = id

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetGoRewardId(ctx, request.(GetGoRewardIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetGoRewardId")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetGoRewardIdResponseObject); ok {
		if err := validResponse.VisitGetGoRewardIdResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PutGoRewardId operation middleware
func (sh *strictHandler) PutGoRewardId(w http.ResponseWriter, r *http.Request, id int) {
	var request PutGoRewardIdRequestObject

	request.Id = id

	var body PutGoRewardIdJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PutGoRewardId(ctx, request.(PutGoRewardIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PutGoRewardId")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PutGoRewardIdResponseObject); ok {
		if err := validResponse.VisitPutGoRewardIdResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetGoSystemHealth operation middleware
func (sh *strictHandler) GetGoSystemHealth(w http.ResponseWriter, r *http.Request) {
	var request GetGoSystemHealthRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetGoSystemHealth(ctx, request.(GetGoSystemHealthRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetGoSystemHealth")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetGoSystemHealthResponseObject); ok {
		if err := validResponse.VisitGetGoSystemHealthResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetGoSystemVersion operation middleware
func (sh *strictHandler) GetGoSystemVersion(w http.ResponseWriter, r *http.Request) {
	var request GetGoSystemVersionRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetGoSystemVersion(ctx, request.(GetGoSystemVersionRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetGoSystemVersion")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetGoSystemVersionResponseObject); ok {
		if err := validResponse.VisitGetGoSystemVersionResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}
