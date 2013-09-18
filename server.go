package main

import (
	"flag"
	"fmt"
	"github.com/paulsmith/gogeos/geos"
	"net/http"
	"strconv"
)

var (
	port = flag.Int("port", 7979, "Server Port")
)

type OperationHandler func(*geos.Geometry, http.ResponseWriter, *http.Request)

func (o OperationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Extract the wkt/geojson
	r.ParseForm()
	wkt := r.PostFormValue("geom")
	fmt.Println(wkt)
	geom, err := geos.FromWKT(wkt)
	if err != nil {
		http.Error(w, "Could not parse geom", 400)
		return
	}
	o(geom, w, r)

}

func Buffer(g *geos.Geometry, w http.ResponseWriter, r *http.Request) {
	geom, err := g.Buffer(2)
	if err != nil {
		http.Error(w, "Could not buffer geom", 400)
		fmt.Println("Buffer", err)
		return
	}
	wkt, err := geom.ToWKT()
	if err != nil {
		http.Error(w, "Could not encode result", 400)
		fmt.Println("WKT", err)
		return
	}
	w.Write([]byte(wkt))
}

func main() {
	flag.Parse()
	p := strconv.Itoa(*port)

	http.Handle("/buffer", OperationHandler(Buffer))
	if err := http.ListenAndServe(":"+p, nil); err != nil {
		fmt.Println("Failed to start server: %v", err)
	}

}
