package trace

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"sync"
)

const Header = "TRACE-ID"

var _ T = (*Trace)(nil)

type T interface {
	//强制  接口中所有方法只能在本包中去实现，在其他包中不允许去实现。因为接口中有小写方法，所以在其他包无法去实现
	i()
	ID() string
	WithRequest(req *Request) *Trace
	WithResponse(resp *Response) *Trace
	AppendDialog(dialog *Dialog) *Trace
	AppendSQL(sql *SQL) *Trace
	AppendRedis(redis *Redis) *Trace
}

// Trace 记录的参数
type Trace struct {
	mux                sync.Mutex
	Identifier         string    `json:"trace_id"`             // 链路ID
	Request            *Request  `json:"request"`              // 请求信息
	Response           *Response `json:"response"`             // 返回信息
	ThirdPartyRequests []*Dialog `json:"third_party_requests"` // 调用第三方接口的信息
	Debugs             []*Debug  `json:"debugs"`               // 调试信息
	SQLs               []*SQL    `json:"sqls"`                 // 执行的 SQL 信息
	Redis              []*Redis  `json:"redis"`                // 执行的 Redis 信息
	Success            bool      `json:"success"`              // 请求结果 true or false
	CostSeconds        float64   `json:"cost_seconds"`         // 执行时长(单位秒)
}

func New(id string) *Trace {
	if id == "" {
		buf := make([]byte, 10)
		io.ReadFull(rand.Reader, buf)
		id = hex.EncodeToString(buf)
	}
	return &Trace{
		Identifier: id,
	}
}

func (t *Trace) i() {}

// ID 唯一标识符
func (t *Trace) ID() string {
	return t.Identifier
}

// WithRequest 设置request
func (t *Trace) WithRequest(req *Request) *Trace {
	t.Request = req
	return t
}

// WithResponse 设置response
func (t *Trace) WithResponse(resp *Response) *Trace {
	t.Response = resp
	return t
}

// AppendDialog 安全的追加内部调用过程dialog
func (t *Trace) AppendDialog(dialog *Dialog) *Trace {
	if dialog == nil {
		return t
	}
	t.mux.Lock()
	defer t.mux.Unlock()

	t.ThirdPartyRequests = append(t.ThirdPartyRequests, dialog)
	return t
}

// AppendDebug 追加 debug
func (t *Trace) AppendDebug(debug *Debug) *Trace {
	if debug == nil {
		return t
	}

	t.mux.Lock()
	defer t.mux.Unlock()

	t.Debugs = append(t.Debugs, debug)
	return t
}

// AppendSQL 追加sql
func (t *Trace) AppendSQL(sql *SQL) *Trace {
	if sql == nil {
		return t
	}
	t.mux.Lock()
	defer t.mux.Unlock()

	t.SQLs = append(t.SQLs, sql)
	return t
}

func (t *Trace) AppendRedis(redis *Redis) *Trace {
	if redis == nil {
		return t
	}
	t.mux.Lock()
	t.mux.Unlock()

	t.Redis = append(t.Redis, redis)
	return t
}
