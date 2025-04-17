package models

import (
	"errors"
	"fmt"

	"github.com/goVendor/database"
	"github.com/golang-jwt/jwt/v5"

	"golang.org/x/crypto/bcrypt"
)

func (u *User) Save() error {
	query := `INSERT INTO users(username,email,password)VALUES(?,?,?)`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hasedPassword, err := HasedPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.UserName, u.Email, hasedPassword) //insert data in the database exec
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	u.Id = id
	return err
}
func (u *User) ValidateCredentials() error {
	query := `SELECT id, email, password FROM users WHERE email=?`
	row := database.DB.QueryRow(query, u.Email)
	var id int64
	var email, retrivePassword string
	err := row.Scan(&id, &email, &retrivePassword)
	if err != nil {
		return err
	}
	isPasswordValid := ComparePassword(u.Password, retrivePassword)
	if !isPasswordValid {
		return fmt.Errorf("invalid password")
	}
	u.Id = id
	return nil
}
func (p *Post) Save() error {
	query := `INSERT INTO posts(authorid,author,title,content)VALUES(?,?,?,?)`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(p.Authorid, p.Author, p.Title, p.Content)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	p.Postid = id
	return err
}
func GetAllPost() ([]Post, error) {
	query := `SELECT * FROM posts`
	row, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var posts []Post
	for row.Next() {
		var post Post
		err := row.Scan(&post.Postid, &post.Authorid, &post.Author, &post.Title, &post.Content)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil

}
func GetPostById(postid int) (*Post, error) {
	query := "SELECT * FROM posts WHERE postid=?"
	row := database.DB.QueryRow(query, postid)
	var post Post
	err := row.Scan(&post.Postid, &post.Authorid, &post.Author, &post.Title, &post.Content)
	if err != nil {
		return nil, err
	}
	return &post, err
}
func (p Post) UpDatePost() error {
	query := `UPDATE posts SET authorid=?,author=?,title=?,content=? WHERE postid=? `
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(p.Authorid, p.Author, p.Title, p.Content, p.Postid)
	return err

}
func HasedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func ComparePassword(password, hasedpassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hasedpassword), []byte(password))
	return err == nil

}

var secretKey = []byte("kiran is my name")

func GenerateJwtToken(id int64, username, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"username": username,
		"email":    email,
	})
	return token.SignedString(secretKey)
}
func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unauthorized signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, err
	}
	if parsedToken == nil || !parsedToken.Valid {
		return 0, err
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}
	userIdFloat, ok := claims["id"].(float64)
	if !ok {
		return 0, errors.New("userId not found in token")
	}
	return int64(userIdFloat), nil
}
