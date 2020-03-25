package collectors

import (
	"context"

	"github.com/taiyang-li/automi/api"
	autoctx "github.com/taiyang-li/automi/api/context"
	"github.com/taiyang-li/automi/util"
)

type NullCollector struct {
	input <-chan interface{}
	logf  api.LogFunc
}

func Null() *NullCollector {
	return new(NullCollector)
}

func (s *NullCollector) SetInput(in <-chan interface{}) {
	s.input = in
}

// Open opens the node to start collecting
func (s *NullCollector) Open(ctx context.Context) <-chan error {
	result := make(chan error)
	s.logf = autoctx.GetLogFunc(ctx)
	util.Logfn(s.logf, "Opening null collector")

	go func() {
		defer func() {
			util.Logfn(s.logf, "Closing null collector")
			close(result)
		}()

		for {
			select {
			case _, opened := <-s.input:
				if !opened {
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return result
}
