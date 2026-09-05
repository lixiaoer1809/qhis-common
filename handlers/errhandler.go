package handlers

import (
	"github.com/lixiaoer1809/qhis-common/result"
	"net/http"
)

// ErrHandler 错误返回
func ErrHandler(err error) (int, any) {
	return http.StatusOK, result.Fail(result.CodeSysError, err.Error())
}
