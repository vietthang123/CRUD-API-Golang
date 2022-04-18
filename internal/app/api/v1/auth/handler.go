package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"developer-orientenergy-golang/internal/pkg/util"
)

type authController struct {
	authService IAuthService
	authToken   Authentication
}

func NewAuthController(authService IAuthService) *authController {
	authToken := NewAuthentication()
	return &authController{authService: authService, authToken: authToken}
}

func (c *authController) CheckLogin(w http.ResponseWriter, r *http.Request) {
	userName := r.Header.Get("username")
	passWord := r.Header.Get("password")
	account, err := c.authService.GetUserByUserNameAndPassWord(userName, passWord)
	if err != nil {
		util.RespondJSONError(w, http.StatusUnauthorized, err.Error())
		return
	}
	userInfo, err := c.authService.GetUserInfoByID(account.ID)
	tokenString, err := c.authToken.GenerateToken(userInfo)
	if err != nil {
		util.RespondJSONError(w, http.StatusUnauthorized, err.Error())
		return
	}
	var tokenResp = map[string]string{}
	tokenResp["token"] = tokenString
	util.RespondJSON(w, http.StatusOK, tokenResp)
	return
}

func (c *authController) InsertUser(w http.ResponseWriter, r *http.Request) {
	var formData UserAccount
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &formData)
	if err != nil {
		util.RespondJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	user, err := c.authService.InsertUser(formData)
	if err != nil {
		util.RespondJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	util.RespondJSONSuccess(w, fmt.Sprintf("Create User Success %d", user.ID))
	return
}
