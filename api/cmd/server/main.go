package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"bitbucket.org/irenato/coverletter-hub/api/internal/auth"
	"bitbucket.org/irenato/coverletter-hub/api/internal/config"
	"bitbucket.org/irenato/coverletter-hub/api/internal/coverletter"
	"bitbucket.org/irenato/coverletter-hub/api/internal/database"
	"bitbucket.org/irenato/coverletter-hub/api/internal/llm"
	"bitbucket.org/irenato/coverletter-hub/api/internal/parser"
	"bitbucket.org/irenato/coverletter-hub/api/internal/profile"
	"bitbucket.org/irenato/coverletter-hub/api/internal/user"
	"bitbucket.org/irenato/coverletter-hub/api/internal/vacancy"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := config.Load()
	ctx := context.Background()

	pool, err := database.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}
	defer pool.Close()

	if err := database.Migrate(ctx, pool); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	userRepo := user.NewPostgresRepository(pool)
	profileRepo := profile.NewPostgresRepository(pool)
	docRepo := profile.NewPostgresDocRepository(pool)
	vacancyRepo := vacancy.NewPostgresRepository(pool)
	clRepo := coverletter.NewPostgresRepository(pool)

	linkedInClient := auth.NewLinkedInClient(cfg.LinkedInClientID, cfg.LinkedInClientSecret, cfg.LinkedInRedirectURI)
	llmClient := llm.NewClaudeClient(cfg.ClaudeAPIKey)
	cvParser := parser.New(llmClient)

	authHandler := auth.NewAuthHandler(userRepo, linkedInClient, cfg.JWTSecret, cfg.FrontendURL)
	profileService := profile.NewService(profileRepo, docRepo, cvParser)
	profileHandler := profile.NewHandler(profileService)
	vacancyHandler := vacancy.NewHandler(vacancyRepo)
	clService := coverletter.NewService(clRepo, vacancyRepo, profileRepo, llmClient)
	clHandler := coverletter.NewHandler(clService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/auth/linkedin", authHandler.HandleLinkedInRedirect)
		r.Get("/auth/linkedin/callback", authHandler.HandleLinkedInCallback)

		r.Group(func(r chi.Router) {
			r.Use(auth.JWTMiddleware(cfg.JWTSecret))

			r.Get("/auth/me", authHandler.HandleMe)

			r.Get("/profile", profileHandler.HandleGet)
			r.Put("/profile", profileHandler.HandleUpdate)
			r.Post("/profile/upload", profileHandler.HandleUpload)

			r.Get("/vacancies", vacancyHandler.HandleList)
			r.Post("/vacancies", vacancyHandler.HandleCreate)
			r.Get("/vacancies/{id}", vacancyHandler.HandleGet)
			r.Put("/vacancies/{id}", vacancyHandler.HandleUpdate)
			r.Delete("/vacancies/{id}", vacancyHandler.HandleDelete)
			r.Post("/vacancies/{id}/cover-letter", clHandler.HandleGenerate)

			r.Get("/cover-letters", clHandler.HandleList)
			r.Get("/cover-letters/{id}", clHandler.HandleGet)
			r.Put("/cover-letters/{id}", clHandler.HandleUpdateText)
			r.Patch("/cover-letters/{id}/status", clHandler.HandleUpdateStatus)
		})
	})

	addr := fmt.Sprintf(":%s", cfg.APIPort)
	log.Printf("starting server on %s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
