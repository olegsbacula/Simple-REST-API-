package main

import (
    "encoding/json"
    "errors"
    "log"
    "net/http"
    "strconv"
    "sync"
    "github.com/go-chi/chi"
    "Simple-REST-API-/internal/models"
)
type User struct {
    User_ID        int    `json:"user_id"`
    User_Name      string `json:"user_name"`
    User_Last_name string `json:"user_last_name"`
    User_Email     string `json:"user_email"`
    CreatedAt      string `json:"created_at"`
}

type UserService struct {
	mu     sync.Mutex
	store map[int]*User
	nextID int
}

func NewUserService() *UserService{
	return &UserService{
		store: make(map[int]*models.User),
		nextID:1,
	}
}

func (s *UserService) Create(u *models.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	u.User_ID=s.nextID
	s.nextID++

	s.store[u.User_ID]= u 
	return nil

}

func (s *UserService) List(u *models.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	users := make([]models.User, 0, len(s.store))
	
	for _, u:=range s.store{
		users = append (users, *u)
	}

	return users
}

func (s *UserService) Get(id int) (*models.User, error) {	
	s.mu.Lock()
	defer s.mu.Unlock()
	
	if u, ok := s.store[id]; ok {
        return u, nil
    }
	
	return nil, errors.New("user not found")
}

func (s *UserService) Update(u *models.User) error{	
	s.mu.Lock()
	defer s.mu.Unlock()
	
	if _, ok := s.store[u.User_ID]; ok {
        s.store[u.User_ID] = u 
		return nil
    }
	
	return errors.New("user not found")
}

func (s *UserService) Delete(id int) error {
    s.mu.Lock()
    defer s.mu.Unlock()

    if _, ok := s.store[id]; ok {
		delete(s.store,id)
	}
	return errors.New("user not found")
}