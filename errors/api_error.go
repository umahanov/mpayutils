package errors

import (
	"encoding/json"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// APIError кастомная ошибка для передачи клиенту.
type APIError struct {
	GrpcCode     codes.Code `json:"-"`                 // grpc код ошибки. Не передается клиенту.
	APIErrorCode string     `json:"code"`              // Строковое представление кода ошибки для передачи клиенту
	Message      string     `json:"message"`           // Человекочитаемое сообщение об ошибке.
	Details      []any      `json:"details,omitempty"` // Дополнительная информация об ошибке.
	Err          error      `json:"-"`                 // Исходная ошибка. Не передается клиенту.
}

// GRPCStatus реализует интерфейс GRPCStatus чтобы можно было пробрасывать в grpc-gw code и message кастомной ошибки
func (e *APIError) GRPCStatus() *status.Status {
	return status.New(e.GrpcCode, e.Message)
}

// Error реализует error интерфейс для APIError.
func (e *APIError) Error() string {
	errMsg := e.Message
	if e.Err != nil {
		errMsg = e.Err.Error()
	}

	errJson, _ := json.Marshal(map[string]any{
		"code":    e.APIErrorCode,
		"message": e.Message,
		"details": e.Details,
		"error":   errMsg,
	})

	return string(errJson)
}
