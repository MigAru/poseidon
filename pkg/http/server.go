package http

type HttpServer interface {
	Run()
	Shutdown()
}
