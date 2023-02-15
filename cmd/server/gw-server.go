package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
	"sever-codebase/application/models"
	"sever-codebase/proto/server-proto/pb"
)

type ErrorResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func StartGrpcGatewayServer(cfg *models.Config) {
	// Custom response When error
	runtime.HTTPError = customHTTPError

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}))
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Register
	err := pb.RegisterBaseServerHandlerFromEndpoint(context.Background(), mux, fmt.Sprintf(":%d", cfg.Server.HttpPort), opts)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	muxHttp := http.NewServeMux()
	muxHttp.Handle("/", allowCORS(mux))

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.HttpPort), muxHttp))
}

func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		origin := req.Header.Get("Origin")
		if &origin != nil && origin != "" {
			writer.Header().Set("Access-Control-Allow-Origin", "*")
			writer.Header().Set("Access-Control-Allow-Credentials", "true")
			writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		}
		h.ServeHTTP(writer, req)
	})
}

func customHTTPError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	const fallback = `{"error": "failed to marshal error message"}`

	w.Header().Set("Content-type", marshaler.ContentType())
	w.WriteHeader(httpStatusFromCode(status.Code(err)))

	jErr := json.NewEncoder(w).Encode(
		ErrorResponse{
			Code:    int(status.Code(err)),
			Message: status.Convert(err).Message(),
			Data:    nil,
		})

	if jErr != nil {
		w.Write([]byte(fallback))
	}
}

func httpStatusFromCode(code codes.Code) int {
	return 0
}
