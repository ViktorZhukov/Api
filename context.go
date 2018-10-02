package jsonrpc

import (
	"context"

	"net/http"
)

type requestId struct{}

type responseWriter struct{}

type headers struct{}

type cookies struct{}

type сookieGetter func(name string) (*http.Cookie, error)

// RequestId takes request id from context.
func RequestId(c context.Context) string {
	return c.Value(requestId{}).(string)
}

// WithRequestId adds request id to context.
func SetRequestId(c context.Context, id string) context.Context {
	return context.WithValue(c, requestId{}, id)
}

func Headers(c context.Context) http.Header  {
	return c.Value(headers{}).(http.Header)
}

func SetHeaders(c context.Context, h http.Header) context.Context  {
	return context.WithValue(c, headers{}, h)
}


func ResponseWriter(c context.Context) http.ResponseWriter  {
	return c.Value(responseWriter{}).(http.ResponseWriter)
}

func SetResponseWriter(c context.Context, writer http.ResponseWriter) context.Context  {
	return context.WithValue(c, responseWriter{}, writer)
}

func Cookie(c context.Context, name string) (*http.Cookie, error)  {
	return c.Value(cookies{}).(сookieGetter)(name)
}

func SetCookie(c context.Context, cookie сookieGetter) context.Context  {
	return context.WithValue(c, cookies{}, cookie)
}