package errors

import "google.golang.org/grpc/codes"

// NewAPIError создает APIError.
func NewAPIError(code codes.Code, message string, details []any, err error) *APIError {
	return &APIError{
		GrpcCode:     code,
		APIErrorCode: convertToCode(code),
		Message:      message,
		Details:      details,
		Err:          err,
	}
}

// NewInvalidArgumentAPIError создает APIError с GrpcCode=InvalidArgument, APIErrorCode=BAD_REQUEST и http статусом 400
func NewInvalidArgumentAPIError(message string, details []any, err error) *APIError {
	return NewAPIError(codes.InvalidArgument, message, details, err)
}

// NewUnauthenticatedAPIError создает APIError с GrpcCode=Unauthenticated, APIErrorCode=UNAUTHORIZED и http статусом 401
func NewUnauthenticatedAPIError(message string, details []any, err error) *APIError {
	return NewAPIError(codes.Unauthenticated, message, details, err)
}

// NewPermissionDeniedAPIError создает APIError с GrpcCode=PermissionDenied, APIErrorCode=FORBIDDEN и http статусом 403
func NewPermissionDeniedAPIError(message string, details []any, err error) *APIError {
	return NewAPIError(codes.PermissionDenied, message, details, err)
}

// NewNotFoundAPIError создает APIError с GrpcCode=NotFound, APIErrorCode=NOT_FOUND и http статусом 404
func NewNotFoundAPIError(message string, details []any, err error) *APIError {
	return NewAPIError(codes.NotFound, message, details, err)
}

// NewInternalAPIError создает APIError с GrpcCode=Internal, APIErrorCode=INTERNAL_SERVER_ERROR и http статусом 500
func NewInternalAPIError(message string, details []any, err error) *APIError {
	return NewAPIError(codes.Internal, message, details, err)
}

// NewUnimplementedAPIError создает APIError с GrpcCode=Unimplemented, APIErrorCode=NOT_IMPLEMENTED и http статусом 501
func NewUnimplementedAPIError(message string, details []any, err error) *APIError {
	return NewAPIError(codes.Unimplemented, message, details, err)
}

// convertToCode конвертируем grpc код в строковое представление APIErrorCode
func convertToCode(grpcCode codes.Code) string {
	apiErrorCode := GrpcToApiErrorCodeMap[grpcCode]
	if apiErrorCode == "" {
		apiErrorCode = "UNKNOWN_ERROR"
	}

	return apiErrorCode
}
