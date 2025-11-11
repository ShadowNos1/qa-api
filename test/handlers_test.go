package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ShadowNos1/qa-api/internal/app"
	"github.com/ShadowNos1/qa-api/internal/model"
)

// === Вопросы ===

func TestCreateAndListQuestions(t *testing.T) {
	svc := app.NewInMemoryService()
	h := app.NewHandler(svc)
	router := app.NewRouter(h)

	// Создание вопроса
	qBody := []byte(`{"text":"Что такое Go?"}`)
	req := httptest.NewRequest(http.MethodPost, "/questions/", bytes.NewBuffer(qBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Result().StatusCode != http.StatusCreated {
		t.Fatalf("ожидали 201 Created, получили %d", w.Result().StatusCode)
	}

	// Получение списка вопросов
	req2 := httptest.NewRequest(http.MethodGet, "/questions/", nil)
	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, req2)

	if w2.Result().StatusCode != http.StatusOK {
		t.Fatalf("ожидали 200 OK, получили %d", w2.Result().StatusCode)
	}

	var questions []model.Question
	if err := json.NewDecoder(w2.Body).Decode(&questions); err != nil {
		t.Fatalf("не удалось декодировать JSON: %v", err)
	}

	if len(questions) != 1 || questions[0].Text != "Что такое Go?" {
		t.Fatalf("вопрос не сохранён корректно")
	}
}

// === Ответы ===

func TestCreateAndGetAnswer(t *testing.T) {
	svc := app.NewInMemoryService()
	h := app.NewHandler(svc)
	router := app.NewRouter(h)

	// Создадим вопрос
	q := &model.Question{Text: "Вопрос для ответа?"}
	if err := svc.CreateQuestion(q); err != nil {
		t.Fatalf("не удалось создать вопрос: %v", err)
	}

	// Добавим ответ
	aBody := []byte(`{"user_id":"user123","text":"Это язык Go"}`)
	req := httptest.NewRequest(http.MethodPost, "/questions/1/answers/", bytes.NewBuffer(aBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Result().StatusCode != http.StatusCreated {
		t.Fatalf("ожидали 201 Created, получили %d", w.Result().StatusCode)
	}

	// Получим ответ
	req2 := httptest.NewRequest(http.MethodGet, "/answers/1", nil)
	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, req2)

	if w2.Result().StatusCode != http.StatusOK {
		t.Fatalf("ожидали 200 OK, получили %d", w2.Result().StatusCode)
	}

	var ans model.Answer
	if err := json.NewDecoder(w2.Body).Decode(&ans); err != nil {
		t.Fatalf("не удалось декодировать JSON: %v", err)
	}

	if ans.Text != "Это язык Go" || ans.UserID != "user123" {
		t.Fatalf("ответ не совпадает")
	}
}
