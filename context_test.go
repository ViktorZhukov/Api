package jsonrpc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
)

func TestRequestId(t *testing.T) {

	c := context.Background()
	id := "1"
	c = SetRequestId(c, id)
	var pick string
	require.NotPanics(t, func() {
		pick = RequestId(c)
	})
	require.Equal(t, id, pick)
}

func TestHeader(t *testing.T) {

	c := context.Background()
	headers := http.Header{
		"header": []string{
			"simple header 1",
			"simple header 2",
		},
	}
	c = SetHeaders(c, headers)
	var pick http.Header
	require.NotPanics(t, func() {
		pick = Headers(c)
	})
	require.Equal(t, headers, pick)
}


func TestResponseWriter(t *testing.T) {

	c := context.Background()
	w := httptest.NewRecorder()
	c = SetResponseWriter(c, w)
	var pick http.ResponseWriter
	require.NotPanics(t, func() {
		pick = ResponseWriter(c)
	})
	require.Equal(t, w, pick)
}


func TestCookies(t *testing.T) {

	c := context.Background()
	cookie := &http.Cookie{
		Name:"test",
	}

	getter := func(name string) (*http.Cookie, error){
		if name == "test" {
			return cookie, nil
		}
		return nil, http.ErrNoCookie
	}

	c = SetCookies(c, getter)
	pick, err := Cookies(c,"test")

	require.NoError(t, err)
	require.Equal(t, cookie, pick)
}