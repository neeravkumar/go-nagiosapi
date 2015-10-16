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
	w.Render.JSON(wr, http.StatusOK, status.Host[c.URLParams["host"]])
	status.RUnlock()
}
