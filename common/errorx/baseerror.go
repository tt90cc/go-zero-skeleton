package errorx

var ErrMSG = make(map[int]string)

const (
	ErrOK      = 0   // 成功
	ErrDEFAULT = 1   // 错误
	ErrPARAMS  = 499 // 参数错误
	ErrSYSTEM  = 500 // 系统内部错误
)

func init() {
	ErrMSG[ErrOK] = "请求成功"
	ErrMSG[ErrDEFAULT] = "系统未知错误"
	ErrMSG[ErrPARAMS] = "参数错误"
	ErrMSG[ErrSYSTEM] = "系统内部错误"
}

type CodeError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (r *CodeError) Error() string {
	return r.Message
}

func NewCodeError(code int) error {
	return &CodeError{
		Code:    code,
		Message: ErrMSG[code],
	}
}

type CodeErrorResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (r *CodeError) Info() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code:    r.Code,
		Message: r.Message,
		Data:    r.Data,
	}
}
