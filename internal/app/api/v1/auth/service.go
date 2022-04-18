package auth

import "errors"

type IAuthService interface {
	GetUserByUserNameAndPassWord(userName, passWord string) (UserAccount, error)
	GetUserByID(userId int) (UserAccount, error)
	InsertUser(dto UserAccount) (UserAccount, error)
	GetUserInfoByID(UserAccountID int) (UserInfo, error)
}

type authService struct {
	authRepo IAuthRepository
}

func NewAuthService(repository IAuthRepository) *authService {
	return &authService{authRepo: repository}
}

func (s *authService) GetUserByUserNameAndPassWord(userName, passWord string) (UserAccount, error) {
	if userName == "" || passWord == "" {
		return UserAccount{}, errors.New("Username and password required")
	}
	return s.authRepo.GetUserByUserNameAndPassWord(userName, passWord)
}

func (s *authService) GetUserByID(userId int) (UserAccount, error) {
	return s.authRepo.GetUserByID(userId)
}

func (s *authService) InsertUser(dto UserAccount) (UserAccount, error) {
	return s.authRepo.InsertUser(dto)
}

func (s *authService) GetUserInfoByID(UserAccountID int) (UserInfo, error) {
	return s.authRepo.GetUserInfo(UserAccountID)
}
