// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// Admin HTTP server
//
// Command:
// $ goa gen github.com/tonouchi510/goa2-sample/design

package server

import (
	"context"
	"net/http"
	"regexp"

	admin "github.com/tonouchi510/goa2-sample/gen/admin"
	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
	"goa.design/plugins/cors"
)

// Server lists the Admin service endpoint HTTP handlers.
type Server struct {
	Mounts          []*MountPoint
	UserNumber      http.Handler
	AdminListUser   http.Handler
	AdminGetUser    http.Handler
	AdminCreateUser http.Handler
	AdminUpdateUser http.Handler
	AdminDeleteUser http.Handler
	CORS            http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the Admin service endpoints.
func New(
	e *admin.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"UserNumber", "GET", "/api/v1/admin/user_number"},
			{"AdminListUser", "GET", "/api/v1/admin/users"},
			{"AdminGetUser", "GET", "/api/v1/admin/users/{id}"},
			{"AdminCreateUser", "POST", "/api/v1/admin/users"},
			{"AdminUpdateUser", "PUT", "/api/v1/admin/users/{id}"},
			{"AdminDeleteUser", "DELETE", "/api/v1/admin/users/{id}"},
			{"CORS", "OPTIONS", "/api/v1/admin/user_number"},
			{"CORS", "OPTIONS", "/api/v1/admin/users"},
			{"CORS", "OPTIONS", "/api/v1/admin/users/{id}"},
		},
		UserNumber:      NewUserNumberHandler(e.UserNumber, mux, dec, enc, eh),
		AdminListUser:   NewAdminListUserHandler(e.AdminListUser, mux, dec, enc, eh),
		AdminGetUser:    NewAdminGetUserHandler(e.AdminGetUser, mux, dec, enc, eh),
		AdminCreateUser: NewAdminCreateUserHandler(e.AdminCreateUser, mux, dec, enc, eh),
		AdminUpdateUser: NewAdminUpdateUserHandler(e.AdminUpdateUser, mux, dec, enc, eh),
		AdminDeleteUser: NewAdminDeleteUserHandler(e.AdminDeleteUser, mux, dec, enc, eh),
		CORS:            NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "Admin" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.UserNumber = m(s.UserNumber)
	s.AdminListUser = m(s.AdminListUser)
	s.AdminGetUser = m(s.AdminGetUser)
	s.AdminCreateUser = m(s.AdminCreateUser)
	s.AdminUpdateUser = m(s.AdminUpdateUser)
	s.AdminDeleteUser = m(s.AdminDeleteUser)
	s.CORS = m(s.CORS)
}

// Mount configures the mux to serve the Admin endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountUserNumberHandler(mux, h.UserNumber)
	MountAdminListUserHandler(mux, h.AdminListUser)
	MountAdminGetUserHandler(mux, h.AdminGetUser)
	MountAdminCreateUserHandler(mux, h.AdminCreateUser)
	MountAdminUpdateUserHandler(mux, h.AdminUpdateUser)
	MountAdminDeleteUserHandler(mux, h.AdminDeleteUser)
	MountCORSHandler(mux, h.CORS)
}

// MountUserNumberHandler configures the mux to serve the "Admin" service
// "user_number" endpoint.
func MountUserNumberHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleAdminOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/api/v1/admin/user_number", f)
}

// NewUserNumberHandler creates a HTTP handler which loads the HTTP request and
// calls the "Admin" service "user_number" endpoint.
func NewUserNumberHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		encodeResponse = EncodeUserNumberResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "user_number")
		ctx = context.WithValue(ctx, goa.ServiceKey, "Admin")

		res, err := endpoint(ctx, nil)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountAdminListUserHandler configures the mux to serve the "Admin" service
// "admin list user" endpoint.
func MountAdminListUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleAdminOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/api/v1/admin/users", f)
}

// NewAdminListUserHandler creates a HTTP handler which loads the HTTP request
// and calls the "Admin" service "admin list user" endpoint.
func NewAdminListUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		encodeResponse = EncodeAdminListUserResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "admin list user")
		ctx = context.WithValue(ctx, goa.ServiceKey, "Admin")

		res, err := endpoint(ctx, nil)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountAdminGetUserHandler configures the mux to serve the "Admin" service
// "admin get user" endpoint.
func MountAdminGetUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleAdminOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/api/v1/admin/users/{id}", f)
}

// NewAdminGetUserHandler creates a HTTP handler which loads the HTTP request
// and calls the "Admin" service "admin get user" endpoint.
func NewAdminGetUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeAdminGetUserRequest(mux, dec)
		encodeResponse = EncodeAdminGetUserResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "admin get user")
		ctx = context.WithValue(ctx, goa.ServiceKey, "Admin")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountAdminCreateUserHandler configures the mux to serve the "Admin" service
// "admin create user" endpoint.
func MountAdminCreateUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleAdminOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/api/v1/admin/users", f)
}

// NewAdminCreateUserHandler creates a HTTP handler which loads the HTTP
// request and calls the "Admin" service "admin create user" endpoint.
func NewAdminCreateUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeAdminCreateUserRequest(mux, dec)
		encodeResponse = EncodeAdminCreateUserResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "admin create user")
		ctx = context.WithValue(ctx, goa.ServiceKey, "Admin")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountAdminUpdateUserHandler configures the mux to serve the "Admin" service
// "admin update user" endpoint.
func MountAdminUpdateUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleAdminOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PUT", "/api/v1/admin/users/{id}", f)
}

// NewAdminUpdateUserHandler creates a HTTP handler which loads the HTTP
// request and calls the "Admin" service "admin update user" endpoint.
func NewAdminUpdateUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeAdminUpdateUserRequest(mux, dec)
		encodeResponse = EncodeAdminUpdateUserResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "admin update user")
		ctx = context.WithValue(ctx, goa.ServiceKey, "Admin")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountAdminDeleteUserHandler configures the mux to serve the "Admin" service
// "admin delete user" endpoint.
func MountAdminDeleteUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleAdminOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("DELETE", "/api/v1/admin/users/{id}", f)
}

// NewAdminDeleteUserHandler creates a HTTP handler which loads the HTTP
// request and calls the "Admin" service "admin delete user" endpoint.
func NewAdminDeleteUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeAdminDeleteUserRequest(mux, dec)
		encodeResponse = EncodeAdminDeleteUserResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "admin delete user")
		ctx = context.WithValue(ctx, goa.ServiceKey, "Admin")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service Admin.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = handleAdminOrigin(h)
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("OPTIONS", "/api/v1/admin/user_number", f)
	mux.Handle("OPTIONS", "/api/v1/admin/users", f)
	mux.Handle("OPTIONS", "/api/v1/admin/users/{id}", f)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// handleAdminOrigin applies the CORS response headers corresponding to the
// origin for the service Admin.
func handleAdminOrigin(h http.Handler) http.Handler {
	spec0 := regexp.MustCompile(".*localhost.*")
	origHndlr := h.(http.HandlerFunc)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			origHndlr(w, r)
			return
		}
		if cors.MatchOriginRegexp(origin, spec0) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Max-Age", "600")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "authorization, content-type")
			}
			origHndlr(w, r)
			return
		}
		origHndlr(w, r)
		return
	})
}
