// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// Admin endpoints
//
// Command:
// $ goa gen github.com/tonouchi510/goa2-sample/design

package admin

import (
	"context"

	goa "goa.design/goa"
)

// Endpoints wraps the "Admin" service endpoints.
type Endpoints struct {
	UserNumber      goa.Endpoint
	AdminListUser   goa.Endpoint
	AdminGetUser    goa.Endpoint
	AdminCreateUser goa.Endpoint
	AdminUpdateUser goa.Endpoint
	AdminDeleteUser goa.Endpoint
}

// NewEndpoints wraps the methods of the "Admin" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		UserNumber:      NewUserNumberEndpoint(s),
		AdminListUser:   NewAdminListUserEndpoint(s),
		AdminGetUser:    NewAdminGetUserEndpoint(s),
		AdminCreateUser: NewAdminCreateUserEndpoint(s),
		AdminUpdateUser: NewAdminUpdateUserEndpoint(s),
		AdminDeleteUser: NewAdminDeleteUserEndpoint(s),
	}
}

// Use applies the given middleware to all the "Admin" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.UserNumber = m(e.UserNumber)
	e.AdminListUser = m(e.AdminListUser)
	e.AdminGetUser = m(e.AdminGetUser)
	e.AdminCreateUser = m(e.AdminCreateUser)
	e.AdminUpdateUser = m(e.AdminUpdateUser)
	e.AdminDeleteUser = m(e.AdminDeleteUser)
}

// NewUserNumberEndpoint returns an endpoint function that calls the method
// "user_number" of service "Admin".
func NewUserNumberEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		res, err := s.UserNumber(ctx)
		if err != nil {
			return nil, err
		}
		vres := NewViewedGoa2SampleAdminUserNumber(res, "default")
		return vres, nil
	}
}

// NewAdminListUserEndpoint returns an endpoint function that calls the method
// "admin list user" of service "Admin".
func NewAdminListUserEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		res, err := s.AdminListUser(ctx)
		if err != nil {
			return nil, err
		}
		vres := NewViewedGoa2SampleAdminUserCollection(res, "default")
		return vres, nil
	}
}

// NewAdminGetUserEndpoint returns an endpoint function that calls the method
// "admin get user" of service "Admin".
func NewAdminGetUserEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetUserPayload)
		res, err := s.AdminGetUser(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedGoa2SampleAdminUser(res, "default")
		return vres, nil
	}
}

// NewAdminCreateUserEndpoint returns an endpoint function that calls the
// method "admin create user" of service "Admin".
func NewAdminCreateUserEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CreateUserPayload)
		return s.AdminCreateUser(ctx, p)
	}
}

// NewAdminUpdateUserEndpoint returns an endpoint function that calls the
// method "admin update user" of service "Admin".
func NewAdminUpdateUserEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdateUserPayload)
		res, err := s.AdminUpdateUser(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedGoa2SampleAdminUser(res, "default")
		return vres, nil
	}
}

// NewAdminDeleteUserEndpoint returns an endpoint function that calls the
// method "admin delete user" of service "Admin".
func NewAdminDeleteUserEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeleteUserPayload)
		return nil, s.AdminDeleteUser(ctx, p)
	}
}