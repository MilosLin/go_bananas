package protocol

/**
 * Error Code
 */
const (
	E_OK          = 0
	E_UNKNOWN     = 1
	E_WRONG_INPUT = 2
)

/**
 * api 回應格式
 */
type Response struct {
	Code    int8        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
