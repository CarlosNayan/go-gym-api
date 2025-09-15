package utils

import (
	"api-gym-on-go/src/config/env"
	"fmt"
	"sort"
	"time"

	"github.com/gofiber/fiber/v2"
)

var methodMap = map[string]string{
	"GET":     "GET",
	"HEAD":    "HEAD",
	"POST":    "POST",
	"PUT":     "PUT",
	"DELETE":  "DELETE",
	"CONNECT": "CONNECT",
	"OPTIONS": "OPTIONS",
	"TRACE":   "TRACE",
	"PATCH":   "PATCH",
}

func RouteLogger(app *fiber.App, port int) {

	if env.ENVIRONMENT != "production" {
		var routes []*fiber.Route

		for _, routeList := range app.Stack() {
			for _, route := range routeList {
				if len(route.Handlers) > 0 && route.Path != "/" && route.Method != "HEAD" {
					routes = append(routes, route) // Adiciona o ponteiro para a rota
				}
			}
		}

		// Ordena as rotas pela parte da string após o primeiro "/"
		sort.Slice(routes, func(i, j int) bool {
			return routes[i].Path[1:] < routes[j].Path[1:]
		})

		// Exibe as rotas ordenadas
		for _, route := range routes {
			method := route.Method // Aqui usa diretamente a string do método
			fmt.Printf("\n\x1b[34m[Fiber]\x1b[0m - %s - \x1b[34m%s\x1b[0m - %s", time.Now().UTC().Format("02/01/2006 15:04:05"), route.Path, methodMap[method])
		}

		fmt.Printf("\n\x1b[34m[Fiber]\x1b[0m - %s - %v Routes Loaded. ", time.Now().UTC().Format("02/01/2006 15:04:05"),  len(routes))

		routes = nil
	}

	// Exibe mensagem de sucesso
	fmt.Printf("\n\n🚀 Servidor iniciado na porta %v \n\n", port)

}
