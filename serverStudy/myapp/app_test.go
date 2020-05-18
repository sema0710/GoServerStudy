package myapp

import (
	// "io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIndexPathHandler(t *testing.T) {

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(t,http.StatusOK, res.Code)
	assert.Equal(t,"Hello World", "")
}
