package main

import (
	"log"
	"os"

	"github.com/ObjectIsAdvantag/CLEUR-1814/5-webserver/routes"
	"github.com/ObjectIsAdvantag/CLEUR-1814/5-webserver/utils"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func main() {
	path := os.Getenv("GOPATH") + "/src/github.com/ObjectIsAdvantag/CLEUR-1814/5-webserver/constants/secrets.yml"
	utils.LoadYAML(path)

	r := fasthttprouter.New()
	routes.Init(r)

	host := os.Getenv("HOST")
	port := os.Getenv("HOST_PORT")

	log.Fatal(fasthttp.ListenAndServe(host+":"+port, r.Handler))

}
