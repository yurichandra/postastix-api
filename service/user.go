package service

import (
	"errors"

	"github.com/dewadg/postastix-api/db"
	"github.com/dewadg/postastix-api/repository"
	"golang.org/x/crypto/bcrypt"
)

// UserService represents a service of User.
type UserService struct {
	repo repository.UserRepositoryContract
}

// Get returns available users.
func (s *UserService) Get() []*db.User {
	return s.repo.Get()
}

func (s *UserService) isUsernameUnique(name string, exceptID uint) bool {
	foundUser := &db.User{}
	query := db.Get().Where("name = ?", name)

	if exceptID != 0 {
		query = query.Where("ID != ?", exceptID)
	}

	query.First(foundUser)

	return foundUser.ID == 0
}

func (s *UserService) generatePassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", errors.New("Failed to generate password hash")
	}
	return string(hashed), nil
}

// Create stores new user and returns it.
func (s *UserService) Create(name string, fullName string, password string) (*db.User, error) {
	if !s.isUsernameUnique(name, 0) {
		return nil, errors.New("Username already taken")
	}

	hashedPassword, err := s.generatePassword(password)
	if err != nil {
		return nil, err
	}

	newUser := &db.User{
		Name:     name,
		FullName: fullName,
		Password: hashedPassword,
	}
	s.repo.Push(newUser)

	return newUser, nil
}

// Find returns user by ID.
func (s *UserService) Find(id uint) (*db.User, error) {
	user := s.repo.Find(id)

	if user.ID == 0 {
		return nil, errors.New("User not found")
	}
	return user, nil
}

// Update updates a user and returns it.
func (s *UserService) Update(id uint, name string, fullName string) (*db.User, error) {
	user := s.repo.Find(id)

	if user.ID == 0 {
		return nil, errors.New("User not found")
	}

	if !s.isUsernameUnique(name, id) {
		return nil, errors.New("Username already taken")
	}

	user.Name = name
	user.FullName = fullName
	db.Get().Save(user)

	return user, nil
}

// Delete deletes a user.
func (s *UserService) Delete(id uint) error {
	if s.repo.Find(id).ID == 0 {
		return errors.New("User not found")
	}

	s.repo.Delete(id)
	return nil
}

// ChangePassword changes a user password.
func (s *UserService) ChangePassword(id uint, oldPassword string, newPassword string, confirmNewPassword string) error {
	user := s.repo.Find(id)
	if user.ID == 0 {
		return errors.New("User not found")
	}

	passwordCheck := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
	if passwordCheck != nil {
		return errors.New("Wrong old password")
	}

	if newPassword != confirmNewPassword {
		return errors.New("New password not match")
	}

	hashedPassword, err := s.generatePassword(newPassword)
	if err != nil {
		return errors.New("Failed to generate password hash")
	}

	user.Password = hashedPassword
	db.Get().Save(user)

	return nil
}
