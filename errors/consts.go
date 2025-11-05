package errors

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

// GrpcToHttpMap таблица маппинга кодов grpc и http статусов
var GrpcToHttpMap = map[codes.Code]int{
	codes.Canceled:           499, // Client Closed Request
	codes.Unknown:            http.StatusInternalServerError,
	codes.InvalidArgument:    http.StatusBadRequest,
	codes.DeadlineExceeded:   http.StatusGatewayTimeout,
	codes.NotFound:           http.StatusNotFound,
	codes.AlreadyExists:      http.StatusConflict,
	codes.PermissionDenied:   http.StatusForbidden,
	codes.ResourceExhausted:  http.StatusTooManyRequests,
	codes.FailedPrecondition: http.StatusBadRequest,
	codes.Aborted:            http.StatusConflict,
	codes.OutOfRange:         http.StatusBadRequest,
	codes.Unimplemented:      http.StatusNotImplemented,
	codes.Unavailable:        http.StatusServiceUnavailable,
	codes.DataLoss:           http.StatusInternalServerError,
	codes.Unauthenticated:    http.StatusUnauthorized,
}

// GrpcToApiErrorCodeMap Таблица маппинга кодов grpc и http статусов
var GrpcToApiErrorCodeMap = map[codes.Code]string{
	codes.Canceled:           "CLIENT_CLOSED_REQUEST",
	codes.Unknown:            "INTERNAL_SERVER_ERROR",
	codes.InvalidArgument:    "BAD_REQUEST",
	codes.DeadlineExceeded:   "GATEWAY_TIMEOUT",
	codes.NotFound:           "NOT_FOUND",
	codes.AlreadyExists:      "STATUS_CONFLICT",
	codes.PermissionDenied:   "FORBIDDEN",
	codes.ResourceExhausted:  "TOO_MANY_REQUESTS",
	codes.FailedPrecondition: "BAD_REQUEST",
	codes.Aborted:            "CONFLICT",
	codes.OutOfRange:         "BAD_REQUEST",
	codes.Unimplemented:      "NOT_IMPLEMENTED",
	codes.Internal:           "INTERNAL_SERVER_ERROR",
	codes.Unavailable:        "SERVICE_UNAVAILABLE",
	codes.DataLoss:           "INTERNAL_SERVER_ERROR",
	codes.Unauthenticated:    "UNAUTHENTICATED",
}
