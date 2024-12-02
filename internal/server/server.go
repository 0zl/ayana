package server

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/pug/v2"
)

type Server struct {
	app *fiber.App
}

func New() *Server {
	engine := pug.New("./views", ".pug")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(func(c fiber.Ctx) error {
		c.Set("Content-Security-Policy", "default-src 'self'")
		return c.Next()
	})

	return &Server{
		app: app,
	}
}

func (s *Server) SetupRoutes() {
	s.app.Get("/", func(c fiber.Ctx) error {
		return c.Render("pages/index", fiber.Map{
			"title": "Ayana!",
		})
	})
}

func (s *Server) Start(addr string) error {
	return s.app.Listen(addr)
}
