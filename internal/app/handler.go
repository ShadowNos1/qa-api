package app

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/ShadowNos1/qa-api/internal/model"
)

type Handler struct {
    s Service
}

func NewHandler(s Service) *Handler {
    return &Handler{s: s}
}

// ---------------- Questions ----------------

func (h *Handler) ListQuestions(w http.ResponseWriter, r *http.Request) {
    questions, err := h.s.ListQuestions()
    if err != nil {
        http.Error(w, "internal error", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(questions)
}

func (h *Handler) CreateQuestion(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Text string `json:"text"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "bad request", http.StatusBadRequest)
        return
    }
    q := &model.Question{Text: req.Text}
    if err := h.s.CreateQuestion(q); err != nil {
        http.Error(w, "internal error", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    _ = json.NewEncoder(w).Encode(q)
}

func (h *Handler) GetQuestion(w http.ResponseWriter, r *http.Request) {
    id, ok := parseIDFromPath(r.URL.Path, "/questions/")
    if !ok {
        http.Error(w, "bad id", http.StatusBadRequest)
        return
    }
    q, err := h.s.GetQuestion(id)
    if err != nil {
        if err == ErrNotFound {
            http.Error(w, "not found", http.StatusNotFound)
            return
        }
        http.Error(w, "internal error", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(q)
}

func (h *Handler) DeleteQuestion(w http.ResponseWriter, r *http.Request) {
    id, ok := parseIDFromPath(r.URL.Path, "/questions/")
    if !ok {
        http.Error(w, "bad id", http.StatusBadRequest)
        return
    }
    if err := h.s.DeleteQuestion(id); err != nil {
        if err == ErrNotFound {
            http.Error(w, "not found", http.StatusNotFound)
            return
        }
        http.Error(w, "internal error", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

// ---------------- Answers ----------------

func (h *Handler) CreateAnswer(w http.ResponseWriter, r *http.Request) {
    id, ok := parseIDFromPath(r.URL.Path, "/questions/")
    if !ok {
        http.Error(w, "bad question id", http.StatusBadRequest)
        return
    }
    var req struct {
        UserID string `json:"user_id"`
        Text   string `json:"text"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "bad request", http.StatusBadRequest)
        return
    }
    a := &model.Answer{
        QuestionID: uint(id),
        UserID:     req.UserID,
        Text:       req.Text,
    }
    if err := h.s.CreateAnswer(a); err != nil {
        if err == ErrNotFound {
            http.Error(w, "question not found", http.StatusNotFound)
            return
        }
        http.Error(w, "internal error", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    _ = json.NewEncoder(w).Encode(a)
}

func (h *Handler) GetAnswer(w http.ResponseWriter, r *http.Request) {
    id, ok := parseIDFromPath(r.URL.Path, "/answers/")
    if !ok {
        http.Error(w, "bad id", http.StatusBadRequest)
        return
    }
    a, err := h.s.GetAnswer(id)
    if err != nil {
        if err == ErrNotFound {
            http.Error(w, "not found", http.StatusNotFound)
            return
        }
        http.Error(w, "internal error", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(a)
}

func (h *Handler) DeleteAnswer(w http.ResponseWriter, r *http.Request) {
    id, ok := parseIDFromPath(r.URL.Path, "/answers/")
    if !ok {
        http.Error(w, "bad id", http.StatusBadRequest)
        return
    }
    if err := h.s.DeleteAnswer(id); err != nil {
        if err == ErrNotFound {
            http.Error(w, "not found", http.StatusNotFound)
            return
        }
        http.Error(w, "internal error", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

// ---------------- Utilities ----------------

func indexRune(s string, r rune) int {
    for i, c := range s {
        if c == r {
            return i
        }
    }
    return -1
}

func parseIDFromPath(path, prefix string) (uint, bool) {
    if len(path) <= len(prefix) {
        return 0, false
    }
    s := path[len(prefix):]
    if slash := indexRune(s, '/'); slash != -1 {
        s = s[:slash]
    }
    id64, err := strconv.ParseUint(s, 10, 64)
    if err != nil {
        return 0, false
    }
    return uint(id64), true
}
