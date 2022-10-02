package service

import "context"

// ContextKey はContextに情報を保存するときのキーです
type ContextKey string

var (
	taskNameKey ContextKey = "taskName"
)

// SetTaskNameToContextは、taskの名前をContextにセットします
func SetTaskNameToContext(ctx context.Context, taskName string) context.Context {
	if taskName != "" {
		return context.WithValue(ctx, taskNameKey, taskName)
	}
	return ctx
}

// GetTaskNameFromContextは、Contextからtaskの名前を取得します
func GetTaskNameFromContext(ctx context.Context) (string, bool) {
	v := ctx.Value(taskNameKey)
	taskName, ok := v.(string)
	return taskName, ok
}
