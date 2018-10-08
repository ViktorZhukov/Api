package jsonrpc

import (
	"context"
	"net/http"
	"github.com/intel-go/fastjson"
	con "github.com/Aarabika/context"
)

// Handler links a method of JSON-RPC request.
type Handler interface {
	ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (result interface{}, err *Error)
}

// ServeHTTP provides basic JSON-RPC handling.
func (mr *MethodRepository) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	rs, batch, err := ParseRequest(r)
	if err != nil {
		SendResponse(w, []*Response{
			{
				Version: Version,
				Error:   err,
			},
		}, false)
		return
	}

	c := r.Context()
	// add Headers to base context
	c = SetHeaders(c, r.Header)
	// add ResponseWriter to base context
	c = SetResponseWriter(c, w)
	// add Cookies to base context
	cookieGetter := func(name string) (*http.Cookie, error) {
		return r.Cookie(name)
	}
	c = SetCookie(c, cookieGetter)

	resp := make([]*Response, len(rs))
	for i := range rs {
		resp[i] = mr.InvokeMethod(c, rs[i])
	}

	if err := SendResponse(w, resp, batch); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// InvokeMethod invokes JSON-RPC method.
func (mr *MethodRepository) InvokeMethod(c context.Context, r *Request) *Response {
	var h Handler
	res := NewResponse(r)
	h, res.Error = mr.TakeMethod(r)
	if res.Error != nil {
		return res
	}
	res.Result, res.Error = h.ServeJSONRPC(con.SetRequestId(c, r.ID), r.Params)
	if res.Error != nil {
		res.Result = nil
	}
	return res
}
