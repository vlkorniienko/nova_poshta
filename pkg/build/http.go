package build

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

const readHeaderTimeout = 20 * time.Second

type HTTPConfig struct {
	Port               int    `config:"port"`
	AllowedOriginsCORS string `config:"allowed_origins_cors"`
}

func (c HTTPConfig) GetAllowedOriginsCORS() []string {
	// TODO: better way to serialize array from env variables
	return strings.Split(c.AllowedOriginsCORS, ";")
}

type HTTPHandlerPack interface {
	RegisterRoutes(r HTTPRouter) error
}

type HTTPRouter interface {
	NewRoute() *mux.Route
	HandleFunc(path string, handler func(w http.ResponseWriter, r *http.Request)) *mux.Route
	Handle(path string, handler http.Handler) *mux.Route
}

type HTTPServer struct {
	conf     HTTPConfig
	handlers []HTTPHandlerPack
	s        *http.Server
	r        *mux.Router
}

func RegisterHandlers(r HTTPRouter, handlers ...HTTPHandlerPack) error {
	for _, h := range handlers {
		if err := h.RegisterRoutes(r); err != nil {
			return fmt.Errorf("registering route: %w", err)
		}
	}

	return nil
}

func NewHTTPServer(conf HTTPConfig, handlerPacks ...HTTPHandlerPack) (*HTTPServer, error) {
	r := mux.NewRouter()

	if err := RegisterHandlers(r, handlerPacks...); err != nil {
		return nil, fmt.Errorf("registering handlers: %w", err)
	}

	h := handlers.CORS(
		handlers.AllowedOrigins(conf.GetAllowedOriginsCORS()),
		handlers.AllowedMethods([]string{http.MethodGet, http.MethodHead, http.MethodPost,
			http.MethodPut, http.MethodOptions, http.MethodDelete}),
	)(r)

	s := &HTTPServer{
		conf:     conf,
		handlers: handlerPacks,
		s: &http.Server{
			Addr:              fmt.Sprintf(":%v", conf.Port),
			Handler:           h,
			ReadHeaderTimeout: readHeaderTimeout,
		},
		r: r,
	}

	return s, nil
}

func (s *HTTPServer) Serve(_ context.Context) error {
	if err := logRoutes(s.r); err != nil {
		return fmt.Errorf("logging HTTP routes: %w", err)
	}

	log.Info().Strs("allowed_origins_cors", s.conf.GetAllowedOriginsCORS()).
		Msgf("Listening and serving on port = %v", s.conf.Port)

	if err := s.s.ListenAndServe(); err != nil {
		return fmt.Errorf("listening and serving: %w", err)
	}

	return nil
}


func logRoutes(r *mux.Router) error {
	log.Info().Msg("Available HTTP endpoints are below:")

	err := r.Walk(func(route *mux.Route, _ *mux.Router, _ []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			return fmt.Errorf("getting path templates: %w", err)
		}

		methods, err := route.GetMethods()
		if err != nil {
			log.Info().Msgf("NONE: %s", path)

			return nil // nolint:nilerr
		}

		log.Info().Msgf("%s: %s", strings.Join(methods, ", "), path)

		return nil
	})

	if err != nil {
		return fmt.Errorf("walking through routes: %w", err)
	}

	return nil
}
