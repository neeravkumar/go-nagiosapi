package webapi

import (
	"github.com/efigence/go-nagios"
	"github.com/unrolled/render" // or "gopkg.in/unrolled/render.v1")
	"net/http"
	//	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

var renderer = render.New(render.Options{
	IndentJSON: true,
})

var json404 = map[string]string{
	"msg": "not found",
}

type Webapp struct {
	Render       *render.Render
	NagiosStatus nagios.Status
}

func NewWebapp() Webapp {
	var w Webapp
	w.Render = render.New(render.Options{
		IndentJSON: true,
	})
	return w
}

func (w *Webapp) NagiosHost(c web.C, wr http.ResponseWriter, r *http.Request, status nagios.Status) {
	status.RLock()
	if val, ok := status.Host[c.URLParams["host"]]; ok {
		w.Render.JSON(wr, http.StatusOK, val)
	} else {
		w.Render.JSON(wr, http.StatusNotFound, json404)
	}
	status.RUnlock()
}

// return all services belonging to host
func (w *Webapp) NagiosHostServices(c web.C, wr http.ResponseWriter, r *http.Request, status nagios.Status) {
	status.RLock()
	if val, ok := status.Service[c.URLParams["host"]]; ok {
		w.Render.JSON(wr, http.StatusOK, val)
	} else {
		w.Render.JSON(wr, http.StatusNotFound, json404)
	}
	status.RUnlock()
}

// return single service status
func (w *Webapp) NagiosService(c web.C, wr http.ResponseWriter, r *http.Request, status nagios.Status) {
	status.RLock()
	if val, ok := status.Service[c.URLParams["host"]][c.URLParams["service"]]; ok {
		w.Render.JSON(wr, http.StatusOK, val)
	} else {
		w.Render.JSON(wr, http.StatusNotFound, json404)
	}
	status.RUnlock()
}
