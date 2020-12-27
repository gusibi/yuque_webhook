package yuque

import "context"

// yuque webhook request 转发
type Notifer interface {
	Push(ctx context.Context, req *WebHookRequest) error
}

type WebHook struct {
	Hooks []Notifer
}

func NewWebHook() *WebHook {
	return &WebHook{}
}

func (h *WebHook) Register(ctx context.Context, hook Notifer) error {
	h.Hooks = append(h.Hooks, hook)
	return nil
}

func (h *WebHook) Notify(ctx context.Context, req *WebHookRequest) error {
	for _, hook := range h.Hooks {
		hook.Push(ctx, req)
	}
	return nil
}
