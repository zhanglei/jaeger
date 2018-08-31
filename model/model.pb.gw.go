// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: model.proto

/*
Package model is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package model

import (
	"io"
	"net/http"

	"context"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_CollectorServiceV2_PostSpans_0(ctx context.Context, marshaler runtime.Marshaler, client CollectorServiceV2Client, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq PostSpansRequest
	var metadata runtime.ServerMetadata

	if req.ContentLength > 0 {
		if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
			return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
		}
	}

	msg, err := client.PostSpans(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_QueryServiceV2_GetTrace_0(ctx context.Context, marshaler runtime.Marshaler, client QueryServiceV2Client, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetTraceRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "id")
	}

	protoReq.Id, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", err)
	}

	msg, err := client.GetTrace(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterCollectorServiceV2HandlerFromEndpoint is same as RegisterCollectorServiceV2Handler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterCollectorServiceV2HandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterCollectorServiceV2Handler(ctx, mux, conn)
}

// RegisterCollectorServiceV2Handler registers the http handlers for service CollectorServiceV2 to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterCollectorServiceV2Handler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterCollectorServiceV2HandlerClient(ctx, mux, NewCollectorServiceV2Client(conn))
}

// RegisterCollectorServiceV2Handler registers the http handlers for service CollectorServiceV2 to "mux".
// The handlers forward requests to the grpc endpoint over the given implementation of "CollectorServiceV2Client".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "CollectorServiceV2Client"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "CollectorServiceV2Client" to call the correct interceptors.
func RegisterCollectorServiceV2HandlerClient(ctx context.Context, mux *runtime.ServeMux, client CollectorServiceV2Client) error {

	mux.Handle("POST", pattern_CollectorServiceV2_PostSpans_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CollectorServiceV2_PostSpans_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CollectorServiceV2_PostSpans_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_CollectorServiceV2_PostSpans_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"api", "v2", "spans"}, ""))
)

var (
	forward_CollectorServiceV2_PostSpans_0 = runtime.ForwardResponseMessage
)

// RegisterQueryServiceV2HandlerFromEndpoint is same as RegisterQueryServiceV2Handler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterQueryServiceV2HandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterQueryServiceV2Handler(ctx, mux, conn)
}

// RegisterQueryServiceV2Handler registers the http handlers for service QueryServiceV2 to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterQueryServiceV2Handler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterQueryServiceV2HandlerClient(ctx, mux, NewQueryServiceV2Client(conn))
}

// RegisterQueryServiceV2Handler registers the http handlers for service QueryServiceV2 to "mux".
// The handlers forward requests to the grpc endpoint over the given implementation of "QueryServiceV2Client".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "QueryServiceV2Client"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "QueryServiceV2Client" to call the correct interceptors.
func RegisterQueryServiceV2HandlerClient(ctx context.Context, mux *runtime.ServeMux, client QueryServiceV2Client) error {

	mux.Handle("GET", pattern_QueryServiceV2_GetTrace_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_QueryServiceV2_GetTrace_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_QueryServiceV2_GetTrace_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_QueryServiceV2_GetTrace_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3}, []string{"api", "v2", "traces", "id"}, ""))
)

var (
	forward_QueryServiceV2_GetTrace_0 = runtime.ForwardResponseMessage
)
