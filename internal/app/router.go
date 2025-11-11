package app

import (
    "net/http"
    "strings"
)

func NewRouter(h *Handler) http.Handler {
    mux := http.NewServeMux()

    mux.HandleFunc("/questions/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path == "/questions/" {
            if r.Method == http.MethodGet {
                h.ListQuestions(w, r)
                return
            }
            if r.Method == http.MethodPost {
                h.CreateQuestion(w, r)
                return
            }
        }

        if r.Method == http.MethodGet {
            h.GetQuestion(w, r)
            return
        }
        if r.Method == http.MethodDelete {
            h.DeleteQuestion(w, r)
            return
        }

        if r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, "/answers/") {
            h.CreateAnswer(w, r)
            return
        }

        http.NotFound(w, r)
    })

    mux.HandleFunc("/answers/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            h.GetAnswer(w, r)
            return
        }
        if r.Method == http.MethodDelete {
            h.DeleteAnswer(w, r)
            return
        }
        http.NotFound(w, r)
    })

    return mux
}
