package main

import "github.com/go-chi/chi"

func (apiCfg *apiConfig) Route (router *chi.Mux) {
	router.Get("/healthz", handlerReadiness)
	router.Get("/err", handlerErr)
	router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	router.Post("/users", apiCfg.handlerCreateUser)
	router.Get("/feeds", apiCfg.handlerGetFeeds)
	router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	router.Get("/feeds_follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollow))
	router.Post("/feeds_follows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
	router.Delete("/feeds_follows/{feedFollowId}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollow))
	router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handlerGetUserFeedPosts))
}
