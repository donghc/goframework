package trace

// Response 响应信息
type Response struct {
	Header          interface{} `json:"header"`                      // Header 信息
	Body            interface{} `json:"body"`                        // Body 信息
	BusinessCode    int         `json:"business_code,omitempty"`     // 业务码
	BusinessCodeMsg string      `json:"business_code_msg,omitempty"` // 提示信息
	HttpCode        int         `json:"http_code"`                   // HTTP 状态码
	HttpCodeMsg     string      `json:"http_code_msg"`               // HTTP 状态码信息
	CostSeconds     float64     `json:"cost_seconds"`                // 执行时间(单位秒)
}
