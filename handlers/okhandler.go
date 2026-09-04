package handlers

import (
	"context"
	"github.com/lixiaoer1809/qhis-common/result"
)

// OkHandler 成功返回
func OkHandler(_ context.Context, v any) any {
	return result.Success(v)
}
