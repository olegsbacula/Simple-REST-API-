package repository

import "../models"

type UserRepo interface {
    Create(u *models.User) error
    GetByID(id int) (*models.User, error)
    Update(u *models.User) error
    Delete(id int) error
    List() ([]models.User, error)
}
