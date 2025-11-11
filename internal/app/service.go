package app

import (
	"time"

	"github.com/ShadowNos1/qa-api/internal/model"
)

var ErrNotFound = model.ErrNotFound

type Service interface {
	// Questions
	ListQuestions() ([]*model.Question, error)
	CreateQuestion(q *model.Question) error
	GetQuestion(id uint) (*model.Question, error)
	DeleteQuestion(id uint) error

	// Answers
	CreateAnswer(a *model.Answer) error
	GetAnswer(id uint) (*model.Answer, error)
	DeleteAnswer(id uint) error
}

// InMemoryService хранит данные в памяти
type InMemoryService struct {
	questions []*model.Question
	answers   []*model.Answer
}

// Конструктор
func NewInMemoryService() *InMemoryService {
	return &InMemoryService{
		questions: []*model.Question{},
		answers:   []*model.Answer{},
	}
}

// ---------------- Questions ----------------

func (s *InMemoryService) ListQuestions() ([]*model.Question, error) {
	return s.questions, nil
}

func (s *InMemoryService) CreateQuestion(q *model.Question) error {
	q.ID = uint(len(s.questions) + 1)
	q.CreatedAt = time.Now()
	s.questions = append(s.questions, q)
	return nil
}

func (s *InMemoryService) GetQuestion(id uint) (*model.Question, error) {
	for _, q := range s.questions {
		if q.ID == id {
			return q, nil
		}
	}
	return nil, ErrNotFound
}

func (s *InMemoryService) DeleteQuestion(id uint) error {
	for i, q := range s.questions {
		if q.ID == id {
			s.questions = append(s.questions[:i], s.questions[i+1:]...)
			// каскадное удаление ответов
			var newAnswers []*model.Answer
			for _, a := range s.answers {
				if a.QuestionID != id {
					newAnswers = append(newAnswers, a)
				}
			}
			s.answers = newAnswers
			return nil
		}
	}
	return ErrNotFound
}

// ---------------- Answers ----------------

func (s *InMemoryService) CreateAnswer(a *model.Answer) error {
	_, err := s.GetQuestion(a.QuestionID)
	if err != nil {
		return ErrNotFound
	}
	a.ID = uint(len(s.answers) + 1)
	a.CreatedAt = time.Now()
	s.answers = append(s.answers, a)
	return nil
}

func (s *InMemoryService) GetAnswer(id uint) (*model.Answer, error) {
	for _, a := range s.answers {
		if a.ID == id {
			return a, nil
		}
	}
	return nil, ErrNotFound
}

func (s *InMemoryService) DeleteAnswer(id uint) error {
	for i, a := range s.answers {
		if a.ID == id {
			s.answers = append(s.answers[:i], s.answers[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}
