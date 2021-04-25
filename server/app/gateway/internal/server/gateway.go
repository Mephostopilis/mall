package server

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/encoding/json"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http/binding"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

const (
	// SupportPackageIsVersion1 These constants should not be referenced from any other code.
	SupportPackageIsVersion1 = true

	baseContentType = "application"
)

var (
	acceptHeader      = http.CanonicalHeaderKey("Accept")
	contentTypeHeader = http.CanonicalHeaderKey("Content-Type")
)

func newGateway() *gwruntime.ServeMux {
	mux := gwruntime.NewServeMux(
		gwruntime.WithErrorHandler(handleGatewayError),
	)
	return mux
}

// decodeRequest decodes the request body to object.
func decodeRequest(req *http.Request, v interface{}) error {
	subtype := contentSubtype(req.Header.Get(contentTypeHeader))
	if codec := encoding.GetCodec(subtype); codec != nil {
		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return err
		}
		return codec.Unmarshal(data, v)
	}
	return binding.BindForm(req, v)
}

// encodeResponse encodes the object to the HTTP response.
func encodeResponse(w http.ResponseWriter, r *http.Request, v interface{}) error {
	codec := codecForRequest(r)
	data, err := codec.Marshal(v)
	if err != nil {
		return err
	}
	w.Header().Set(contentTypeHeader, contentType(codec.Name()))
	_, _ = w.Write(data)
	return nil
}

// encodeError encodes the error to the HTTP response.
func encodeError(w http.ResponseWriter, r *http.Request, err error) {
	se, ok := errors.FromError(err)
	if !ok {
		se = &errors.StatusError{
			Code:    2,
			Message: err.Error(),
		}
	}
	codec := codecForRequest(r)
	data, _ := codec.Marshal(se)
	w.Header().Set(contentTypeHeader, contentType(codec.Name()))
	w.WriteHeader(se.HTTPStatus())
	_, _ = w.Write(data)
}

// codecForRequest get encoding.Codec via http.Request
func codecForRequest(r *http.Request) encoding.Codec {
	var codec encoding.Codec
	for _, accept := range r.Header[acceptHeader] {
		if codec = encoding.GetCodec(contentSubtype(accept)); codec != nil {
			break
		}
	}
	if codec == nil {
		codec = encoding.GetCodec(json.Name)
	}
	return codec
}

func contentType(subtype string) string {
	return strings.Join([]string{baseContentType, subtype}, "/")
}

func contentSubtype(contentType string) string {
	if contentType == baseContentType {
		return ""
	}
	if !strings.HasPrefix(contentType, baseContentType) {
		return ""
	}
	switch contentType[len(baseContentType)] {
	case '/', ';':
		if i := strings.Index(contentType, ";"); i != -1 {
			return contentType[len(baseContentType)+1 : i]
		}
		return contentType[len(baseContentType)+1:]
	default:
		return ""
	}
}

func handleGatewayError(ctx context.Context, mux *gwruntime.ServeMux, mar gwruntime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	encodeError(w, r, err)
}
