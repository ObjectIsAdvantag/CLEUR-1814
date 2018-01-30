package routes

import (
	"github.com/ObjectIsAdvantag/CLEUR-1814/5-webserver/controllers"
	"github.com/buaazp/fasthttprouter"
)

// Init initializes routes for the fasthttprouter
func Init(r *fasthttprouter.Router) {
	r.GET("/catalog", controllers.SearchCatalog)
}
