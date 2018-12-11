package autotrace

import (
	"net"
	"net/http"

	"go.elastic.co/apm/module/apmhttp"
)

func ptrTestHookServerServe() *func(*http.Server, net.Listener)

func init() {
	*ptrTestHookServerServe() = func(srv *http.Server, lis net.Listener) {
		handler := srv.Handler
		if handler == nil {
			handler = http.DefaultServeMux
		}
		// TODO(axw) this will create duplicate transactions if the
		// handler is already instrumented. We should consider changing
		// the instrumentation modules to exit early if the context
		// already contains a transaction.
		srv.Handler = apmhttp.Wrap(handler)
	}
}
