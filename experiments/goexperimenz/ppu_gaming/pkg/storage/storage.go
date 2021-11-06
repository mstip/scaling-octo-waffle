package storage

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name         string `gorm:";not null"`
	Email        string `gorm:";unique;not null"`
	PasswordHash string `gorm:";not null"`
}

type Storage struct {
	User *UserRepository
}

func (s *Storage) Close() error {
	if err := s.User.DB.Close(); err != nil {
		return err
	}
	return nil
}

func NewInMemStorage() (*Storage, error) {
	return NewStorage("sqlite3", "")
}

func NewStorage(database string, databaseConn string) (*Storage, error) {
	db, err := gorm.Open(database, databaseConn)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&User{})

	return &Storage{
		User: &UserRepository{DB: db},
	}, nil
}

type UserRepository struct {
	DB *gorm.DB
}

func (ur *UserRepository) Create(name string, email string, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	userModel := &User{Name: name, Email: email, PasswordHash: string(hash)}
	errs := ur.DB.Create(userModel).GetErrors()
	if errs != nil && len(errs) != 0 {
		return nil, errs[0]
	}
	return userModel, nil
}
