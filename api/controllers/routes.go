package controllers

import "github.com/lucassilveira96/healthcheck/api/middlewares"

func (s *Server) initializeRoutes() {

	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.GetAllDockers)).Methods("GET")
}
