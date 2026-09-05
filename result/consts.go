package result

const (
	// CodeSuccess ---------------------- 成功码 ----------------------
	CodeSuccess = 0 // 请求成功

	// CodeSysError ---------------------- 系统通用错误 1~99 ----------------------
	CodeSysError       = 1 // 系统内部异常
	CodeParamInvalid   = 2 // 请求参数非法
	CodeNotFound       = 3 // 数据不存在
	CodeDuplicate      = 4 // 数据重复
	CodeForbidden      = 5 // 禁止访问
	CodeTimeout        = 6 // 请求超时
	CodeDbError        = 7 // 数据库异常
	CodeCacheError     = 8 // 缓存异常
	CodeFileUploadFail = 9 // 文件上传失败

	// CodeTokenEmpty ---------------------- 登录鉴权 1000‑1999 ----------------------
	CodeTokenEmpty     = 1001 // Token为空
	CodeTokenExpire    = 1002 // Token已过期
	CodeTokenInvalid   = 1003 // Token非法
	CodeUserNotExist   = 1004 // 用户不存在
	CodePasswordErr    = 1005 // 密码错误
	CodeAccountDisable = 1006 // 账号已禁用
	CodeNoPermission   = 1007 // 无操作权限

	// CodePatientNotFound ---------------------- 患者档案 2000‑2999 ----------------------
	CodePatientNotFound = 2001 // 患者档案不存在
	CodeIdCardRepeat    = 2002 // 身份证号重复
	CodePatientDisabled = 2003 // 患者档案已停用

	// CodeDiagnoseEmpty ---------------------- 中医问诊辨证 3000‑3999 ----------------------
	CodeDiagnoseEmpty    = 3001 // 问诊记录为空
	CodeSymptomParseFail = 3002 // 症状解析失败
	CodeTCMSyndromeFail  = 3003 // AI辨证失败
	CodeDiagnoseClosed   = 3004 // 问诊已结束不可修改

	// CodePrescriptionNotFound ---------------------- 处方方剂药材 4000‑4999 ----------------------
	CodePrescriptionNotFound = 4001 // 处方不存在
	CodeHerbEmpty            = 4002 // 药材列表不能为空
	CodeHerbNotExist         = 4003 // 药材库不存在
	CodePrescriptionAuditNo  = 4004 // 处方未通过审核
	CodeFormulaRepeat        = 4005 // 方剂重复保存

	// CodeReportNotFound ---------------------- 健康报告/评估 5000‑5999 ----------------------
	CodeReportNotFound = 5001 // 健康报告不存在
	CodeEvaluateFail   = 5002 // 体质评估计算失败

	// CodeOrderNotFound ---------------------- 订单支付 6000‑6999 ----------------------
	CodeOrderNotFound = 6001 // 订单不存在
	CodePayFail       = 6002 // 支付失败
	CodeOrderExpire   = 6003 // 订单已过期
)

// MsgMap 业务码与提示文字映射（全局统一文案）
var MsgMap = map[int]string{
	CodeSuccess:        "操作成功",
	CodeSysError:       "服务端出错了,请联系管理员",
	CodeParamInvalid:   "请求参数错误",
	CodeNotFound:       "数据不存在",
	CodeDuplicate:      "数据已存在",
	CodeForbidden:      "访问被拒绝",
	CodeTimeout:        "请求超时",
	CodeDbError:        "数据库异常",
	CodeCacheError:     "缓存服务异常",
	CodeFileUploadFail: "文件上传失败",

	CodeTokenEmpty:     "身份凭证为空，请登录",
	CodeTokenExpire:    "登录已过期，请重新登录",
	CodeTokenInvalid:   "非法的身份凭证",
	CodeUserNotExist:   "用户不存在",
	CodePasswordErr:    "账号或密码错误",
	CodeAccountDisable: "账号已被禁用",
	CodeNoPermission:   "暂无操作权限",

	CodePatientNotFound: "患者档案不存在",
	CodeIdCardRepeat:    "身份证已注册",
	CodePatientDisabled: "该患者档案已停用",

	CodeDiagnoseEmpty:    "暂无问诊数据",
	CodeSymptomParseFail: "症状识别解析失败",
	CodeTCMSyndromeFail:  "AI辨证服务异常",
	CodeDiagnoseClosed:   "问诊已归档，不可编辑",

	CodePrescriptionNotFound: "处方记录不存在",
	CodeHerbEmpty:            "处方药材不能为空",
	CodeHerbNotExist:         "药材不存在于药材库",
	CodePrescriptionAuditNo:  "处方尚未审核通过",
	CodeFormulaRepeat:        "方剂名称重复",

	CodeReportNotFound: "健康评估报告不存在",
	CodeEvaluateFail:   "体质评估计算异常",

	CodeOrderNotFound: "订单不存在",
	CodePayFail:       "支付处理失败",
	CodeOrderExpire:   "订单已失效",
}
