package controllers

import (
	"cleur-barcelona-2018/web/ciscolive.service.content/utils"

	"cleur-barcelona-2018/web/ciscolive.service.content/models"
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
