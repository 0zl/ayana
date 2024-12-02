package server

import (
	"os"

	"github.com/bep/godartsass"
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

	app.Use(func(c fiber.Ctx) error {
		c.Set("Content-Security-Policy", "default-src 'self'")
		return c.Next()
	})

	app.Use("/style.css", func(c fiber.Ctx) error {
		mainScss := "./web/styles/main.scss"

		transpiler, err := godartsass.Start(godartsass.Options{})
		if err != nil {
			return err
		}

		defer transpiler.Close()

		scssContent, err := os.ReadFile(mainScss)
		if err != nil {
			return err
		}

		res, err := transpiler.Execute(godartsass.Args{
			Source:       string(scssContent),
			IncludePaths: []string{"./web/styles"},
			OutputStyle:  godartsass.OutputStyleCompressed,
		})

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
		return c.Render("pages/index", fiber.Map{
			"title": "Ayana!",
		})
	})
}

func (s *Server) Start(addr string) error {
	return s.app.Listen(addr)
}
