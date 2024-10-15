package letual

import (
	"context"
	"github.com/go-chi/cors"

	"log/slog"
	"net/http"

	"letual/internal/config"
	"letual/internal/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type ProductModule interface {
	GetProducts(ctx context.Context) ([]*models.Product, error)
	GetProduct(ctx context.Context, id int) (*models.Product, error)
	SaveProduct(ctx context.Context, product *models.Product) error
}

type UserModule interface {
	Login(ctx context.Context, email, password string) (string, error)
	Register(ctx context.Context, name, email, password string) error
}

type Server struct {
	ctx           context.Context
	logger        *slog.Logger
	cfg           *config.HTTPServer
	productModule ProductModule
	userModule    UserModule
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func OK() Response {
	return Response{
		Status: StatusOK,
	}
}

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

func NewServer(ctx context.Context, log *slog.Logger, cfg *config.HTTPServer, productModule ProductModule, userModule UserModule) *Server {
	return &Server{
		ctx:           ctx,
		logger:        log,
		cfg:           cfg,
		productModule: productModule,
		userModule:    userModule,
	}
}

func (s *Server) Run() error {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Route("/products", func(r chi.Router) {
		//r.Use(middleware.BasicAuth("url-shortener", map[string]string{
		//	s.cfg.User: s.cfg.Password,
		//}))

		r.Get("/", s.getProducts(s.ctx))
		r.Get("/{id}", s.getProduct(s.ctx))
		r.Post("/", s.saveProduct(s.ctx))
	})

	router.Route("/users", func(r chi.Router) {
		r.Post("/sign-in", s.signIn(s.ctx))
		r.Post("/sign-up", s.signUp(s.ctx))
	})

	srv := &http.Server{
		Addr:    s.cfg.Address,
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
