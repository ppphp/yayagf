package handlers

import (
	"net/http"
	"net/http/pprof"
)

var PProfHandlers = []Handler{
	{path:"/", handler:http.HandlerFunc(pprof.Index)},
	{path:"/cmdline", handler:http.HandlerFunc(pprof.Cmdline)},
	{path:"/profile", handler:http.HandlerFunc(pprof.Profile)},
	{path:"/symbol", handler:http.HandlerFunc(pprof.Symbol)},
	{path:"/trace", handler:http.HandlerFunc(pprof.Trace)},
	{path:"/allocs", handler:pprof.Handler("allocs")},
	{path:"/block", handler:pprof.Handler("block")},
	{path:"/goroutine", handler:pprof.Handler("goroutine")},
	{path:"/heap", handler:pprof.Handler("heap")},
	{path:"/mutex", handler:pprof.Handler("mutex")},
	{path:"/threadcreate", handler:pprof.Handler("threadcreate")},
}
