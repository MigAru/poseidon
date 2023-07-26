package http

type Server interface {
	Run()
	Shutdown()
}
