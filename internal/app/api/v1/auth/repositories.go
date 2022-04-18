package auth

import (
	"database/sql"

	"github.com/go-pg/pg/v10"
)

type IAuthRepository interface {
	GetUserByUserNameAndPassWord(userName, passWord string) (UserAccount, error)
	GetUserByID(userId int) (UserAccount, error)
	InsertUser(dto UserAccount) (UserAccount, error)
	GetListUsers() ([]UserAccount, error)
	GetUserInfo(userAccountID int) (UserInfo, error)
	GetListUserInfoByUserAccountIDs(userAccountIDs []int) ([]UserInfo, error)
}

type authRepository struct {
	db   *sql.DB
	dbpg *pg.DB
}

func NewAuthRepository(db *sql.DB, dbpg *pg.DB) *authRepository {
	return &authRepository{db: db, dbpg: dbpg}
}

func (a *authRepository) GetUserByUserNameAndPassWord(userName, password string) (UserAccount, error) {
	user := &UserAccount{}
	err := a.dbpg.Model(user).Where("user_name = ? AND user_password = ?", userName, password).Select()
	if err != nil {
		return UserAccount{}, err
	}
	return *user, nil
}

func (a *authRepository) GetUserByID(userId int) (UserAccount, error) {
	user := &UserAccount{ID: userId}
	err := a.dbpg.Model(user).Select()
	if err != nil {
		return UserAccount{}, err
	}
	return *user, nil
}

func (a *authRepository) InsertUser(dto UserAccount) (UserAccount, error) {
	_, err := a.dbpg.Model(&dto).Insert()
	if err != nil {
		return UserAccount{}, err
	}
	return dto, nil
}

func (a *authRepository) GetListUsers() ([]UserAccount, error) {
	var users []UserAccount
	err := a.dbpg.Model(users).Select()
	if err != nil {
		return []UserAccount{}, err
	}
	return users, nil
}

func (a *authRepository) GetUserInfo(userAccountID int) (UserInfo, error) {
	userInfo := &UserInfo{UserAccountId: userAccountID}
	err := a.dbpg.Model(userInfo).Select()
	if err != nil {
		return UserInfo{}, err
	}
	return *userInfo, nil
}

func (a *authRepository) GetListUserInfoByUserAccountIDs(userAccountIDs []int) ([]UserInfo, error) {
	userInfo := []UserInfo{}
	err := a.dbpg.Model(userInfo).Where("user_account_id in (?)", pg.In(userAccountIDs)).Select()
	if err != nil {
		return []UserInfo{}, err
	}
	return userInfo, nil
}
