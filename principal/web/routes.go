package main

import "net/http"

func (app *application) routes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/savedata", app.saveData)
	mux.HandleFunc("/delete", app.Delete)
	mux.HandleFunc("/editpage", app.editPage)
	mux.HandleFunc("/editchange", app.editChanger)
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return app.logRequest(secureHeaders(mux))
}
