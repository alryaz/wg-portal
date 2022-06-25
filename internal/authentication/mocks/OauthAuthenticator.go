// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	authentication "github.com/h44z/wg-portal/internal/authentication"

	mock "github.com/stretchr/testify/mock"

	oauth2 "golang.org/x/oauth2"
)

// OauthAuthenticator is an autogenerated mock type for the OauthAuthenticator type
type OauthAuthenticator struct {
	mock.Mock
}

// AuthCodeURL provides a mock function with given fields: state, opts
func (_m *OauthAuthenticator) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, state)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, ...oauth2.AuthCodeOption) string); ok {
		r0 = rf(state, opts...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Exchange provides a mock function with given fields: ctx, code, opts
func (_m *OauthAuthenticator) Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, code)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *oauth2.Token
	if rf, ok := ret.Get(0).(func(context.Context, string, ...oauth2.AuthCodeOption) *oauth2.Token); ok {
		r0 = rf(ctx, code, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oauth2.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...oauth2.AuthCodeOption) error); ok {
		r1 = rf(ctx, code, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetType provides a mock function with given fields:
func (_m *OauthAuthenticator) GetType() authentication.AuthenticatorType {
	ret := _m.Called()

	var r0 authentication.AuthenticatorType
	if rf, ok := ret.Get(0).(func() authentication.AuthenticatorType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(authentication.AuthenticatorType)
	}

	return r0
}

// GetUserInfo provides a mock function with given fields: ctx, token, nonce
func (_m *OauthAuthenticator) GetUserInfo(ctx context.Context, token *oauth2.Token, nonce string) (map[string]interface{}, error) {
	ret := _m.Called(ctx, token, nonce)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(context.Context, *oauth2.Token, string) map[string]interface{}); ok {
		r0 = rf(ctx, token, nonce)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *oauth2.Token, string) error); ok {
		r1 = rf(ctx, token, nonce)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseUserInfo provides a mock function with given fields: raw
func (_m *OauthAuthenticator) ParseUserInfo(raw map[string]interface{}) (*authentication.AuthenticatorUserInfo, error) {
	ret := _m.Called(raw)

	var r0 *authentication.AuthenticatorUserInfo
	if rf, ok := ret.Get(0).(func(map[string]interface{}) *authentication.AuthenticatorUserInfo); ok {
		r0 = rf(raw)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*authentication.AuthenticatorUserInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string]interface{}) error); ok {
		r1 = rf(raw)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegistrationEnabled provides a mock function with given fields:
func (_m *OauthAuthenticator) RegistrationEnabled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}