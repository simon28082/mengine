package errors

import (
	"net/http"
)

var (
	ErrDefaultBadRequest         = New(defaultId, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), http.StatusText(http.StatusBadRequest))
	ErrDefaultUnauthorized       = New(defaultId, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), http.StatusText(http.StatusUnauthorized))
	ErrDefaultForbidden          = New(defaultId, http.StatusForbidden, http.StatusText(http.StatusForbidden), http.StatusText(http.StatusForbidden))
	ErrDefaultNotFound           = New(defaultId, http.StatusNotFound, http.StatusText(http.StatusNotFound), http.StatusText(http.StatusNotFound))
	ErrDefaultNotAllowed         = New(defaultId, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed), http.StatusText(http.StatusMethodNotAllowed))
	ErrDefaultTimeout            = New(defaultId, http.StatusRequestTimeout, http.StatusText(http.StatusRequestTimeout), http.StatusText(http.StatusRequestTimeout))
	ErrDefaultConflict           = New(defaultId, http.StatusConflict, http.StatusText(http.StatusConflict), http.StatusText(http.StatusConflict))
	ErrDefaultServerError        = New(defaultId, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), http.StatusText(http.StatusInternalServerError))
	ErrDefaultBadGateway         = New(defaultId, http.StatusBadGateway, http.StatusText(http.StatusBadGateway), http.StatusText(http.StatusBadGateway))
	ErrDefaultServiceUnavailable = New(defaultId, http.StatusServiceUnavailable, http.StatusText(http.StatusServiceUnavailable), http.StatusText(http.StatusServiceUnavailable))
	ErrDefaultGatewayTimeout     = New(defaultId, http.StatusGatewayTimeout, http.StatusText(http.StatusGatewayTimeout), http.StatusText(http.StatusGatewayTimeout))
	ErrDefaultNotImplemented     = New(defaultId, http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented), http.StatusText(http.StatusNotImplemented))
)

var (
	ErrTypeInvalid = NewDefault(`invalid type`)
)
