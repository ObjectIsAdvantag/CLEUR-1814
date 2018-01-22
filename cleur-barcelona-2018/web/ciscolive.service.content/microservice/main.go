package main

import (
	"log"
	"os"

	"cleur-barcelona-2018/web/ciscolive.service.content/routes"
	"cleur-barcelona-2018/web/ciscolive.service.content/utils"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func main() {
	path := os.Getenv("GOPATH") + "/src/cleur-barcelona-2018/web/ciscolive.service.content/constants/secrets.yml"
	utils.LoadYAML(path)

	r := fasthttprouter.New()
	routes.Init(r)

	host := os.Getenv("HOST")
	port := os.Getenv("HOST_PORT")

	log.Fatal(fasthttp.ListenAndServe(host+":"+port, r.Handler))

}
