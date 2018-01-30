package controllers

import (
	"github.com/ObjectIsAdvantag/CLEUR-1814/5-webserver/models"
	"github.com/ObjectIsAdvantag/CLEUR-1814/5-webserver/utils"
	"github.com/valyala/fasthttp"
)

func RespondWith(ctx *fasthttp.RequestCtx, res utils.Result) {
	ctx.Response.Header.SetContentType("application/json")
	ctx.Response.Header.SetConnectionClose()
	ctx.Response.SetStatusCode(res.GetStatus())
	ctx.Response.SetBody(res.JSON())
}

func SearchCatalog(ctx *fasthttp.RequestCtx) {
	res := models.SearchCatalog(ctx)
	RespondWith(ctx, res)
}
