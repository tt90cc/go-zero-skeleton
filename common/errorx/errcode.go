package errorx

var message = make(map[uint32]string)

const (
	OK          = 0   // 成功
	ERR_DEFAULT = 1   // 错误
	ERR_PARAMS  = 499 // 参数错误
)

func init() {
	message[OK] = "请求成功"
	message[ERR_DEFAULT] = "系统未知错误"
	message[ERR_PARAMS] = "参数错误"
}
