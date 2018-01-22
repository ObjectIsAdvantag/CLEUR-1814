package routes

import (
	"cleur-barcelona-2018/web/ciscolive.service.content/controllers"
	"github.com/buaazp/fasthttprouter"
)

// Init initializes routes for the fasthttprouter
func Init(r *fasthttprouter.Router) {
	r.GET("/catalog", controllers.SearchCatalog)
}
