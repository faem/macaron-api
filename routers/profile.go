package routers

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/macaron.v1"
	"log"
	"macaron-api/models"
	"net/http"
	"strings"
	"time"
)

func ReadProfile(c *macaron.Context) {
	profileByte, err := json.MarshalIndent(models.ReadProfile(), "", "\t")
	if err!= nil{
		c.Status(http.StatusInternalServerError)
		c.Write([]byte(err.Error()))
	}
	c.Write(profileByte)
	return
}

func CreateProfile(c *macaron.Context) {
	var p models.Profile
	json.NewDecoder(c.Req.Request.Body).Decode(&p)
	log.Println(p.UserName, p.Name, p.Company, p.Position, p.Password)
	err := models.CreateProfile(p.UserName, p.Name, p.Company, p.Position, p.Password)
	if err != nil {
		c.Write([]byte(err.Error()))
		return
	}
	return
}

func Login(c *macaron.Context) {
	username, password, ok := c.Req.BasicAuth()
	log.Println(username,password,ok)
	if !ok {
		c.Status(http.StatusUnauthorized)
		c.Write([]byte("Unauthorized!"))
		return
	}

	ok, err := models.MatchUsernamePass(username, password)
	if ok {
		mySigningKey := []byte("secretkey")
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)

		claims["admin"] = true
		claims["user"] = username
		claims["exp"] = time.Now().Add(time.Minute * time.Duration(60)).Unix()
		token.Claims = claims

		tokenString, _ := token.SignedString(mySigningKey)
		c.Write([]byte(tokenString+"\n"))
	} else {
		c.Status(http.StatusUnauthorized)
		c.Write([]byte("Username & Password doesn't match!"))
	}

	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.Write([]byte(err.Error()))
	}
}

func CheckAuth(c *macaron.Context)  {
	authHeader := c.Req.Header.Get("Authorization")
	if authHeader == "" {
		c.WriteHeader(http.StatusUnauthorized)
		c.Write([]byte("Need Bearer authorization! Generate token using your username and password here: http://0.0.0.0:<port>/login\n"))
		return
	}
	token, err:= jwt.Parse(strings.Split(authHeader, " ")[1], func(token *jwt.Token) (interface{}, error) {
		return []byte("secretkey"), nil
	})

	if token.Valid {
		fmt.Println("Valid Token")
		//next.ServeHTTP(w, r)
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			c.Write([]byte(fmt.Sprintf("That's not even a token\n")))
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			c.Write([]byte(fmt.Sprintf("The token is Expired! Please issue a new token!")))
		} else {
			c.Write([]byte(fmt.Sprintf("Couldn't handle this token: ", err)))
		}
	} else {
		c.Write([]byte(fmt.Sprintf("Couldn't handle this token: ", err)))
	}
}