package jsonrpc

import (
	"context"
	"net/http"

	"github.com/intel-go/fastjson"
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
	// add headers to base context
	c := r.Context()
	cont := context.WithValue(c, "headers", r.Header)

	resp := make([]*Response, len(rs))
	for i := range rs {
		resp[i] = mr.InvokeMethod(cont, rs[i])
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
	res.Result, res.Error = h.ServeJSONRPC(WithRequestID(c, r.ID), r.Params)
	if res.Error != nil {
		res.Result = nil
	}
	return res
}
