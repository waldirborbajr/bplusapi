package handler

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func RouteHandler() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Timeout(60 * time.Second))
	router.Use(changeMethod)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		// AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		// AllowCredentials: true,
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("B+ Server & API v.0.1.0"))
	})

	router.Route("/api", func(r chi.Router) {
		r.Get("/", apiRoute)
	})

	router.Route("/app", func(r chi.Router) {
		r.Get("/", appRoute)
	})

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", HandlerCustom)
	v1Router.Get("/errorz", HandlerError)

	router.Mount("/v1", v1Router)

	// router.Get("/", controllers.GetAllArticles)
	// router.Post("/upload", controllers.UploadHandler) // Add this
	// router.Get("/images/*", controllers.ServeImages)  // Add this
	// router.Route("/articles", func(r chi.Router) {
	// 	r.Get("/", controllers.NewArticle)
	// 	r.Post("/", controllers.CreateArticle)
	// 	r.Route("/{articleID}", func(r chi.Router) {
	// 		r.Use(controllers.ArticleCtx)
	// 		r.Get("/", controllers.GetArticle)       // GET /articles/1234
	// 		r.Put("/", controllers.UpdateArticle)    // PUT /articles/1234
	// 		r.Delete("/", controllers.DeleteArticle) // DELETE /articles/1234
	// 		r.Get("/edit", controllers.EditArticle)  // GET /articles/1234/edit
	// 	})
	// })

	return router
}

func appRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("app"))
}

func apiRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("api"))
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
