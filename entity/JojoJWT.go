package entity

import "github.com/dgrijalva/jwt-go"

type JojoJWT struct {
	jwt.StandardClaims
	SessID int `json:"sess_id" mapstructure:"sess_id"`
	User   struct {
		ID            int    `json:"id" bson:"id"`
		Email         string `json:"email" bson:"email"`
		CompanyID     int    `json:"company_id" bson:"company_id" mapstructure:"company_id"`
		UserCompanyID int    `json:"user_company_id" bson:"user_company_id" mapstructure:"user_company_id"`
		UserRole      int    `json:"user_role" bson:"user_role" mapstructure:"user_role"`
		UserRoleName  string `json:"user_role_name" bson:"user_role_name" mapstructure:"user_role_name"`
	} `json:"user"`
	Lang           string `json:"lang"`
	SessionSetting int    `json:"session_setting" bson:"session_setting" mapstructure:"session_setting"`
	Token          string `json:"token" bson:"token"`
}

type UserInfo struct {
	UserID        int    `json:"user_id" mapstructure:"user_id"`
	Email         string `json:"email" mapstructure:"email"`
	CompanyID     int    `json:"company_id" mapstructure:"company_id"`
	UserCompanyID int    `json:"user_company_id" mapstructure:"user_company_id"`
	Token         string `json:"token" mapstructure:"token"`
}
