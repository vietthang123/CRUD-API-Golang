package auth

import (
	"encoding/json"

	"developer-orientenergy-golang/internal/app/models"
)

type UserRole struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Permission  string `json:"permission"`
	ClassList   string `json:"class"`
	models.BaseModel
}

type UserRoleInfo struct {
	UserAccountId int             `json:"userAccountID"`
	RoleId        string          `json:"roleId"`
	RoleTitle     string          `json:"roleTitle"`
	Department    string          `json:"department"`
	School        string          `json:"school"`
	Permissions   json.RawMessage `json:"permission"`
	IsActive      bool            `json:"isActive"`
	ClassIds      []string        `json:"classIds"`
	models.BaseModel
}

type UserInfo struct {
	tableName     struct{}        `pg:"user_info"`
	UserAccountId int             `json:"userAccountID"`
	RoleId        int             `json:"roleID"`
	DepartmentId  int             `json:"departmentID"`
	SchoolId      int             `json:"schoolID"`
	IsActive      bool            `json:"isActive"`
	ClassId       []int           `json:"classIds" pg:"-"`
	Status        int             `json:"status"`
	Permissions   json.RawMessage `json:"permission" pg:"-"`
}

type UserAccount struct {
	tableName    struct{} `pg:"user_account"`
	ID           int      `json:"id"`
	UserName     string   `json:"userName"`
	UserEmail    string   `json:"userEmail"`
	UserPassword string   `json:"userPassword"`
	IsActive     bool     `json:"isActive"`
	PhoneNumber  string   `json:"phoneNumber"`
	LoginCount   int      `json:"loginCount"`
	IsFirstLogin bool     `json:"isFirstLogin"`
	Surname      string   `json:"surname"`
	FirstName    string   `json:"firstName"`
	Avatar       string   `json:"avatar"`
	models.BaseModel
}
