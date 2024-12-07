package server

import (
	"os"

	"github.com/bep/golibsass/libsass"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/pug/v2"
)

type Server struct {
	app *fiber.App
}

func New() *Server {
	engine := pug.New("./web/views", ".pug")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use("/style.css", func(c fiber.Ctx) error {
		mainScss := "./web/styles/main.scss"

		transpiler, err := libsass.New(libsass.Options{
			OutputStyle: libsass.CompressedStyle,
		})
		if err != nil {
			return err
		}

		scssContent, err := os.ReadFile(mainScss)
		if err != nil {
			return err
		}

		res, err := transpiler.Execute(string(scssContent))

		if err != nil {
			return err
		}

		c.Set("Content-Type", "text/css")
		return c.SendString(res.CSS)
	})

	return &Server{
		app: app,
	}
}

func (s *Server) SetupRoutes() {
	s.app.Get("/", func(c fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Ayana!",
		}, "layout")
	})
}

func (s *Server) Start(addr string) error {
	return s.app.Listen(addr)
}
