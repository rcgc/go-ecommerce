package user

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"

	"github.com/rcgc/go-ecommerce/domain/user"
	storageUser "github.com/rcgc/go-ecommerce/infraestructure/postgres/user"
)

func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool){
	h := buildHandler(dbPool)

	adminRoutes(e, h)
	publicRoutes(e, h)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	storage := storageUser.New(dbPool)
	useCase := user.New(storage)

	return newHandler(useCase)
}

// rutas publicas, rutas privadas y rutas administrativas
func adminRoutes(e *echo.Echo, h handler) {
	g := e.Group("/api/v1/admin/users")

	g.GET("", h.GetAll)
}

func publicRoutes(e *echo.Echo, h handler) {
	g := e.Group("/api/v1/public/users")

	g.POST("", h.Create)
}
