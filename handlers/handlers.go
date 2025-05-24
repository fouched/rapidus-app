package handlers

import (
	"github.com/fouched/rapidus"
	"myapp/data"
	"myapp/views"
	"net/http"
)

type Handlers struct {
	App    *rapidus.Rapidus
	Models data.Models
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	//defer h.App.LoadTime(time.Now())

	h.render(w, r, views.Home())
}
