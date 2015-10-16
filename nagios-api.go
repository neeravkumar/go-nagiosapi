package main

import (
	"github.com/op/go-logging"
	"os"
	//	"strings"
	"encoding/json"
	"fmt"
	"github.com/efigence/go-nagios"
	"github.com/unrolled/render" // or "gopkg.in/unrolled/render.v1")
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
	"webapi"
)

var version string
var log = logging.MustGetLogger("main")
var stdout_log_format = logging.MustStringFormatter("%{color:bold}%{time:2006-01-02T15:04:05.9999Z-07:00}%{color:reset}%{color} [%{level:.1s}] %{color:reset}%{shortpkg}[%{longfunc}] %{message}")

func main() {
	stderrBackend := logging.NewLogBackend(os.Stderr, "", 0)
	stderrFormatter := logging.NewBackendFormatter(stderrBackend, stdout_log_format)
	logging.SetBackend(stderrFormatter)
	logging.SetFormatter(stdout_log_format)

	log.Info("Starting app")
	log.Debug("version: %s", version)
	log.Info("Starting web server on 8000")
	r := render.New(render.Options{
		IndentJSON: true,
	})
	file, _ := os.Open("t-data/status.dat.local")
	st := nagios.LoadStatus(file)
	file.Close()
	js, _ := json.Marshal(st)

	app := webapi.NewWebapp()
	app.NagiosStatus = st

	goji.Get("/", func(c web.C, w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "%s", js)
		//		r.JSON(w, http.StatusOK, st) // map[string]string{"welcome": "This is rendered JSON!"})
	})
	goji.Get("/host/:host", func(c web.C, w http.ResponseWriter, req *http.Request) {
		app.NagiosHost(c, w, req, st)
	})
	goji.Get("/update", func(c web.C, w http.ResponseWriter, req *http.Request) {
		file, _ := os.Open("t-data/status.dat.local")
		st.UpdateStatus(file)
		file.Close()
	})

	goji.Serve() // Defaults to ":8000".
	_ = r
}
