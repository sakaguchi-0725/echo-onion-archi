package test

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/labstack/echo/v4"
)

func SetupRequest(e *echo.Echo, method, path, requestBody string) (*httptest.ResponseRecorder, *http.Request) {
	csrfReq := httptest.NewRequest(http.MethodGet, "/csrf-token", nil)
	csrfRec := httptest.NewRecorder()
	e.ServeHTTP(csrfRec, csrfReq)

	csrfToken := csrfRec.Result().Cookies()[0].Value

	req := httptest.NewRequest(method, path, strings.NewReader(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("X-CSRF-Token", csrfToken)
	req.AddCookie(&http.Cookie{Name: "csrf_token", Value: csrfToken})

	rec := httptest.NewRecorder()

	return rec, req
}
