package viewmodel

import (
	"context"
	"encoding/json"

	"github.com/beeploop/our-srts/internal/pkg/contextkeys"
	"github.com/google/uuid"
)

type ToastType string

const (
	SUCCESS_TOAST ToastType = "success"
	ERROR_TOAST   ToastType = "error"
	INVALID_TOAST ToastType = "invalid"
)

type Toast struct {
	ID      string    `json:"id"`
	Type    ToastType `json:"type"`
	Title   string    `json:"title"`
	Message string    `json:"message"`
}

func NewSuccessToast(msg string) Toast {
	return Toast{
		ID:      uuid.New().String(),
		Type:    SUCCESS_TOAST,
		Title:   string(SUCCESS_TOAST),
		Message: msg,
	}
}

func NewErrorToast(msg string) Toast {
	return Toast{
		ID:      uuid.New().String(),
		Type:    ERROR_TOAST,
		Title:   string(ERROR_TOAST),
		Message: msg,
	}
}

func (t Toast) ToJson() string {
	b, err := json.Marshal(t)
	if err != nil {
		return ""
	}

	return string(b)
}

func ToastFromContext(ctx context.Context) string {
	if toast, ok := ctx.Value(contextkeys.ToastKey).(string); ok {
		return toast
	}
	return ""
}
