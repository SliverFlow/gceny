package gceny

import (
	"fmt"
	"net/http"
)

type Engine struct {
}

func New() *Engine {
	return &Engine{}
}

// implements http.Handler ServeHTTP func
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request info")
}
