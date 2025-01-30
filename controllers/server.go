package controllers

type Server struct{}

func NewServer() Server {
	return Server{}
}

type StrictServer struct{}

func NewStrictServer() StrictServer {
	return StrictServer{}
}
