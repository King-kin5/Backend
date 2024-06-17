package handler
import(
	"github.com/labstack/echo/v4"
)
const (
	URLSignUp = "/signup"
	URLUser   = "/user"
)

func (h *Handler) RegisterRoutes(g *echo.Group) {
	g.GET("/", h.BaseRouter)

	//  routes
	g.POST(URLUser+URLSignUp, h.UserSignUp)  // /user/signup
	
}
