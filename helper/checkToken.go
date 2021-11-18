package helper

import (
	"fmt"
	// "reflect"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func CheckToken(tokenstr string) map[string]string {
	respMessage := make(map[string]string)
	if tokenstr == ""{
		respMessage["warning"] = "You are unauthorized, please login."
		return respMessage
	}

	tokenSlice := strings.Split(tokenstr, " ")

	if tokenSlice[0] != "Bearer"{
		respMessage["warning"] = "Token is not Berear token!"
		return respMessage
	}

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenSlice[1], claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("secret_key")), nil
	})

	secureIss := "go-auth-system"

	if _, ok := claims["iss"]; !ok {
		fmt.Println("ok: ", ok)
		respMessage["warning"] = "You have not permission, not secure token!"
		return respMessage
	}
	if claims["iss"] != secureIss {
		respMessage["warning"] = "You have not permission, not secure token!"
		return respMessage
	}

	if token.Valid {
		fmt.Println("Token is valid!")
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			respMessage["warning"] = "That's not even a token!"
			fmt.Println("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Token is expired or not active!")
			respMessage["warning"] = "Token is expired or not active!"
		} else {
			respMessage["warning"] = "Couldn't handle this token!"
			fmt.Println("Couldn't handle this token:", err)
		}
	} else {
		respMessage["warning"] = "Couldn't handle this token!"
		fmt.Println("Couldn't handle this token:", err)
	}

	// fmt.Println("only for print token: ", token)
	
	
	claimsMap := make(map[string]string)

	for key, val := range claims {
		
		switch v := val.(type) { 
		default:
			fmt.Printf("unexpected type %T\n", v)
		case float64:
			fmt.Println("")
		case string:
			strVal := val
			if key == "Username"{
				claimsMap["Username"] = strVal.(string)
			}
			if key == "Usertype"{
				claimsMap["Usertype"] = strVal.(string)
			}
			if key == "iss"{
				claimsMap["iss"] = strVal.(string)
			}
		}
	}

	if len(respMessage) != 0 {
		return respMessage
	}
	return claimsMap

}