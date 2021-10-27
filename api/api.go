package api

import (
	_ "image/png"
	"net/http"

	"github.com/allegro/bigcache"
	"github.com/dotkom/image-server/storage"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type API struct {
	storage storage.Storage
	db      *gorm.DB
	cache   *bigcache.BigCache
	router  *mux.Router
}

func New(storage storage.Storage, db *gorm.DB, router *mux.Router, cache *bigcache.BigCache) *API {
	api := &API{}
	api.storage = storage
	api.db = db
	api.router = router
	api.cache = cache
	api.setupRoutes()
	return api
}

func (api *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.router.ServeHTTP(w, r)
}