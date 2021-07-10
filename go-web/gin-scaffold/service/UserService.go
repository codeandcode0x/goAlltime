package service

import (
	"gin-scaffold/model"
)

type UserService struct {
	DAO model.UserDAO
}

// get user service
func (s *UserService) getSvc() *UserService {
	var m model.BaseModel
	return &UserService{
		DAO: &model.User{BaseModel: m},
	}
}

// find all users
func (s *UserService) FindAllUsers() (interface{}, error) {
	var users []model.User
	err := s.getSvc().DAO.FindAll(&users)
	return users, err
}

// find all Jobs with key
func (s *UserService) FindUserByEmail(email string) (*model.User, error) {
	user := model.User{}
	return &user, s.getSvc().DAO.FindByKeys(&user, map[string]interface{}{"email": email})
}

// find user by id
func (s *UserService) FindUserById(id uint64) (*model.User, error) {
	user := model.User{}
	return &user, s.getSvc().DAO.FindByKeys(&user, map[string]interface{}{"id": id})
}

// find all Jobs with key
func (s *UserService) FindAllUsersWithKeys(keys map[string]interface{}) ([]model.User, error) {
	users := []model.User{}
	return users, s.getSvc().DAO.FindByKeys(&users, keys)
}

// find all Jobs with key
func (s *UserService) FindAllUserByJobId(keys map[string]interface{}) (model.User, error) {
	user := model.User{}
	return user, s.getSvc().DAO.FindByKeys(&user, keys)
}

// create user
func (s *UserService) CreateUser(user *model.User) error {
	return s.getSvc().DAO.Create(user)
}

// create user
func (s *UserService) GetUserByJobId(keys map[string]interface{}) (*model.User, error) {
	users := model.User{}
	return &users, s.getSvc().DAO.FindByKeys(&users, keys)
}

// update user
func (s *UserService) UpdateUser(id uint64, user *model.User) (int64, error) {
	rowsAffected, err := s.getSvc().DAO.Update(user, user.ID)
	return rowsAffected, err
}

// delete user
func (s *UserService) DeleteUser(id uint64) (int64, error) {
	return s.getSvc().DAO.Delete(&model.User{}, id)
}

// find all users
func (s *UserService) FindAllUserByPages(currentPage, pageSize int, totalRows *int64) ([]model.User, error) {
	users := []model.User{}
	err := s.getSvc().DAO.Count(&users, totalRows)
	if err != nil {
		return users, err
	}
	return users, s.getSvc().DAO.FindByPages(&users, currentPage, pageSize)
}

//search users by keys
func (s *UserService) FindAllUserByPagesWithKeys(keys, keyOpts map[string]interface{}, currentPage, pageSize int, totalRows *int64) (interface{}, error) {
	var users []model.User
	err := s.getSvc().DAO.CountWithKeys(&users, totalRows, keys, keyOpts)
	if err != nil {
		return users, err
	}

	return users, s.getSvc().DAO.FindByPagesWithKeys(&users, keys, currentPage, pageSize)
}

//search users by keys
func (s *UserService) SearchUserByPagesWithKeys(keys, keyOpts map[string]interface{}, currentPage, pageSize int, totalRows *int64) (interface{}, error) {
	var users []model.User
	err := s.getSvc().DAO.CountWithKeys(&users, totalRows, keys, keyOpts)
	if err != nil {
		return users, err
	}

	return users, s.getSvc().DAO.SearchByPagesWithKeys(&users, keys, keyOpts, currentPage, pageSize)
}
