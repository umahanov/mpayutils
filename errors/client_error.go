package errors

type ClientError struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Details []map[string]any `json:"details"`
}
