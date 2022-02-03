package controllers

import "http_monitoring/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

 
	s.Router.HandleFunc("/urls", middlewares.SetMiddlewareJSON(s.Createurl)).Methods("POST")
	s.Router.HandleFunc("/urls", middlewares.SetMiddlewareJSON(s.Geturls)).Methods("GET")
	s.Router.HandleFunc("/urls/{id}", middlewares.SetMiddlewareJSON(s.Geturl)).Methods("GET")
	s.Router.HandleFunc("/urls/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.Updateurl))).Methods("PUT")
	s.Router.HandleFunc("/urls/{id}", middlewares.SetMiddlewareAuthentication(s.Deleteurl)).Methods("DELETE")
}