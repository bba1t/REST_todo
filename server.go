package todo

import (
	"context"
	"net/http"
	"time"
)

// Эта структура является небольшой абстракцией над структурой Server из пакета http

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		// Даю значения полям из &http.Server, которые передадутся моему полю s.httpServer
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,          // 1 MB
		ReadTimeout:    10 * time.Second, // 10 sec
		WriteTimeout:   10 * time.Second,
	}

	// Этот метод под капотом запускает бесконечный цикл for, который слушает все входящие запросы
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
