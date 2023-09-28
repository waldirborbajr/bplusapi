package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/waldirborbajr/bplusapi/internal/controllers"
)

func NewHandler() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(changeMethod)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Get("/", controllers.GetAllArticles)
	router.Post("/upload", controllers.UploadHandler) // Add this
	router.Get("/images/*", controllers.ServeImages)  // Add this
	router.Route("/articles", func(r chi.Router) {
		r.Get("/", controllers.NewArticle)
		r.Post("/", controllers.CreateArticle)
		r.Route("/{articleID}", func(r chi.Router) {
			r.Use(controllers.ArticleCtx)
			r.Get("/", controllers.GetArticle)       // GET /articles/1234
			r.Put("/", controllers.UpdateArticle)    // PUT /articles/1234
			r.Delete("/", controllers.DeleteArticle) // DELETE /articles/1234
			r.Get("/edit", controllers.EditArticle)  // GET /articles/1234/edit
		})
	})

	return router
}

func changeMethod(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			switch method := r.PostFormValue("_method"); method {
			case http.MethodPut:
				fallthrough
			case http.MethodPatch:
				fallthrough
			case http.MethodDelete:
				r.Method = method
			default:
			}
		}
		next.ServeHTTP(w, r)
	})
}
