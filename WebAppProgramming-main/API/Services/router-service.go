package Services

import (
	"example.com/m/Controllers"
	"example.com/m/Middleware"
	"net/http"
)

type RouterService struct {
	mux *http.ServeMux
}

func NewRouterService() *RouterService {
	mux := http.NewServeMux()
	return &RouterService{mux: mux}
}

func (rs *RouterService) InitializeRoutes() {
	// Echo Controller Routes
	rs.mux.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			Controllers.GetEchoController(w, r)
		case http.MethodPost:
			Controllers.PostEchoController(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

func (rs *RouterService) GetHandler() http.Handler {
	return Middleware.LoggingMiddleware(rs.mux)
}
