package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestIndexhandler(t *testing.T) {
	e := echo.New()

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	expected := `{"status": "OK", "message": "Hello World!"}`
	if assert.NoError(t, indexHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, expected, rec.Body.String())
	}
}
