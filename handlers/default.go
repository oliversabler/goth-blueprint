package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"path/filepath"

	"github.com/a-h/templ"
	"github.com/oliversabler/goth-blueprint/components"
	"github.com/oliversabler/goth-blueprint/services"
)

type DefaultHandler struct {
	SiteName string
	log      *slog.Logger
}

func NewDefaultHandler(log *slog.Logger) *DefaultHandler {
	return &DefaultHandler{
		SiteName: "My site",
		log:      log,
	}
}

func (dh *DefaultHandler) Handle() {
	dh.handleAssets()
	dh.handleRoot()
	dh.handleBlog()
	dh.handleStaticPages()
}

func (dh *DefaultHandler) handleAssets() {
	absAssetsDir, err := filepath.Abs("assets/")
	if err != nil {
		dh.log.Error(
			fmt.Sprintf("Failed to get absolute filepath for assets"),
			slog.Any("error", err))
		return
	}

	http.Handle("/assets/",
		http.StripPrefix("/assets", http.FileServer(http.Dir(absAssetsDir))))
}

func (dh *DefaultHandler) handleRoot() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Redirect(w, r, "/404", http.StatusNotFound)
		} else {
			components.Home(dh.SiteName).Render(r.Context(), w)
		}
	})
}

func (dh *DefaultHandler) handleBlog() {
	posts := services.GetPosts()
	http.Handle("/blog/", templ.Handler(components.Blog(dh.SiteName, posts)))

	for _, post := range posts {
		http.Handle(post.URI, templ.Handler(components.Post(dh.SiteName, post)))
	}
}

func (dh *DefaultHandler) handleStaticPages() {
	http.Handle("/about/", templ.Handler(components.About(dh.SiteName)))
}
