package protocol

/**
 * Error Code
 */
const (
	// ErrCodeOK : Success Response
	ErrCodeOK = 0

	// ErrCodeUnknown : Unhandle Error
	ErrCodeUnknown = 1

	// ErrCodeWrongInput : Wrong Inout
	ErrCodeWrongInput = 2
)

// Response api 回應格式
type Response struct {
	Code    int8        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
