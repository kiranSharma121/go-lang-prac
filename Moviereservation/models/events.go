package models

import (
	"github.com/movie/database"
	"github.com/movie/utils"
)

func (u *User) Save() error {
	query := `INSERT INTO users(username,email,password,role)VALUES(?,?,?,?)`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	hasedpassword, err := utils.HasedPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.UserName, u.Email, hasedpassword, u.Role)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	u.Userid = id
	return err
}
func (u *User) Validatecredentials() error {
	query := `SELECT userid,username,email,password,role FROM users WHERE email=?`
	row := database.DB.QueryRow(query, u.Email)
	var retrivepassword string
	err := row.Scan(&u.Userid, &u.UserName, &u.Email, &retrivepassword, &u.Role)
	if err != nil {
		return err
	}
	isPasswordValid := utils.CompareHasedPassword(u.Password, retrivepassword)
	if !isPasswordValid {
		return err
	}
	return nil
}
func (m *Movie) Save() error {
	query := `INSERT INTO movies(title,description,genre,userid) VALUES(?,?,?,?)`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(m.Title, m.Description, m.Genre, m.Userid)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	m.Movieid = id
	return err
}
func Getallmovies() ([]Movie, error) {
	query := `SELECT * FROM movies`
	row, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var movies []Movie
	for row.Next() {
		var movie Movie
		err := row.Scan(&movie.Movieid, &movie.Title, &movie.Description, &movie.Genre, &movie.Userid)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}
func GetMoviesById(movieid int) (*Movie, error) {
	query := `SELECT * FROM movies WHERE movieid=?`
	row := database.DB.QueryRow(query, movieid)
	var movie Movie
	err := row.Scan(&movie.Movieid, &movie.Title, &movie.Description, &movie.Genre, &movie.Userid)
	if err != nil {
		return nil, err
	}
	return &movie, nil

}
func (m Movie) Updatemovie() error {
	query := `UPDATE movies SET movieid=?,title=?,description=?,genre=?,userid=?`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(m.Movieid, m.Title, m.Description, m.Genre, m.Userid)
	return err
}
func (m Movie) Deletemovies() error {
	query := `DELETE FROM movies WHERE movieid=? `
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(m.Movieid)
	return err
}
