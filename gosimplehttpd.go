package gosimplehttpd

import (
	"fmt"
	"net/http"
)

// Handler is struct for handle http
type FileHttpHandler struct {
	Directory, Port string
}

func (handler *FileHttpHandler) Awake() {
	fileServer := http.StripPrefix("/", http.FileServer(http.Dir(handler.Directory)))

	http.ListenAndServe(fmt.Sprintf(`:%s`, handler.Port), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fileServer.ServeHTTP(w, r)
	}))

}
