package repository

import (
    "errors"
    "sync"
    "yourapp/internal/models"
)

type InMemoryUserRepo struct {
    mu     sync.Mutex
    store  map[int]*models.User
    nextID int
}