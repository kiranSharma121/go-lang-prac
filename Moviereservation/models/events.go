package models

import (
	"fmt"

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

func (s *Showtime) Save() error {
	query := `INSERT INTO shows (movieid, userid, startat, capacity)VALUES(?,?,?,?)`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(s.Movieid, s.Userid, s.StartAT, s.Capacity)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	s.Showtimeid = id
	return err
}
func (s Showtime) Updateshowtime() error {
	query := `UPDATE shows SET showtimeid=?,movieid=?,userid=?,startat=?,capacity=?`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(s.Showtimeid, s.Movieid, s.Userid, s.StartAT, s.Capacity)
	if err != nil {
		return err
	}
	return nil
}
func Getshowsbyid(showtimeid int) (*Showtime, error) {
	query := `SELECT * FROM shows WHERE showtimeid=?`
	row := database.DB.QueryRow(query, showtimeid)
	var showtime Showtime
	err := row.Scan(&showtime.Showtimeid, &showtime.Movieid, &showtime.Userid, &showtime.StartAT, &showtime.Capacity)
	if err != nil {
		return nil, err
	}
	return &showtime, err
}
func Getallshowtime() ([]Showtime, error) {
	query := `SELECT * FROM  shows`
	row, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	var shows []Showtime
	for row.Next() {
		var show Showtime
		err := row.Scan(&show.Showtimeid, &show.Movieid, &show.Userid, &show.StartAT, &show.Capacity)
		if err != nil {
			return nil, err
		}
		shows = append(shows, show)
	}
	return shows, err
}
func (s *Seat) Save() error {
	query := `INSERT INTO seats(seatnumber,isbooked)VALUES(?,?)`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(s.SeatNumber, s.IsBooked)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	s.Seatid = id
	return err
}

func Allseats() ([]Seat, error) {
	query := `SELECT * FROM seats`
	row, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var seats []Seat
	for row.Next() {
		var seat Seat
		err := row.Scan(&seat.Seatid, &seat.SeatNumber, &seat.IsBooked)
		if err != nil {
			return nil, err
		}
		seats = append(seats, seat)
	}
	return seats, nil
}
func (b *Booking) Save() error {
	checkQuery := `SELECT COUNT(*) FROM booking WHERE seatid = ? AND showtimeid = ?`
	var count int
	err := database.DB.QueryRow(checkQuery, b.SeatID, b.ShowtimeID).Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to check seat availability: %w", err)
	}

	if count > 0 {
		return fmt.Errorf("seat already booked for this showtime")
	}

	
	query := `INSERT INTO booking(userid, seatid, showtimeid, bookat) VALUES (?, ?, ?, ?)`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(b.UserID, b.SeatID, b.ShowtimeID, b.BookedAt)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	b.BookingID = id
	return nil
}
