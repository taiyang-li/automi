package util

import (
	"github.com/taiyang-li/automi/api"
)

// Logfn if f is not nil, invokes f(msg)
func Logfn(f api.LogFunc, msg interface{}) {
	if f != nil {
		f(msg)
	}
}
