package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// Data Models
type Movie struct {
	ID     int
	Title  string
	Genre  string
	Duration int
}

type TheaterRoom struct {
	ID       int
	Name     string
	Capacity int
	Amenities string
}

type Showtime struct {
	ID       int
	MovieID  int
	TheaterRoomID int
	Date     time.Time
	Time     time.Time
}

type Ticket struct {
	ID       int
	ShowtimeID int
	CustomerID int
	SeatNumber int
}

// Database Connection
var db *sqlx.DB

func initDB() {
	db, err := sqlx.Open("sqlite3", "./movietheater.db")
	if err != nil {
		log.Fatal(err)
	}

	db.MustExec(`CREATE TABLE IF NOT EXISTS movies (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		genre TEXT,
		duration INTEGER
	)`)

	db.MustExec(`CREATE TABLE IF NOT EXISTS theater_rooms (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		capacity INTEGER,
		amenities TEXT
	)`)

	db.MustExec(`CREATE TABLE IF NOT EXISTS showtimes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		movie_id INTEGER,
		theater_room_id INTEGER,
		date TEXT,
		time TEXT
	)`)

	db.MustExec(`CREATE TABLE IF NOT EXISTS tickets (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		showtime_id INTEGER,
		customer_id INTEGER,
		seat_number INTEGER
	)`)
}

// Functions for Movie Management
func addMovie(movie *Movie) {
	_, err := db.Exec("INSERT INTO movies (title, genre, duration) VALUES (?, ?, ?)", movie.Title, movie.Genre, movie.Duration)
	if err != nil {
		log.Fatal(err)
	}
}

func getMovies() []Movie {
	var movies []Movie
	err := db.Select(&movies, "SELECT * FROM movies")
	if err != nil {
		log.Fatal(err)
	}
	return movies
}

// Function for Theater Room Management
func addTheaterRoom(theaterRoom *TheaterRoom) {
	_, err := db.Exec("INSERT INTO theater_rooms (name, capacity, amenities) VALUES (?, ?, ?)", theaterRoom.Name, theaterRoom.Capacity, theaterRoom.Amenities)
	if err != nil {
		log.Fatal(err)
	}
}

// Main Function
func main() {
	initDB()
	defer db.Close()

	movie := &Movie{Title: "Inception", Genre: "Sci-Fi", Duration: 148}
	addMovie(movie)

	theaterRoom := &TheaterRoom{Name: "Room 1", Capacity: 50, Amenities: "3D"}
	addTheaterRoom(theaterRoom)

	movies := getMovies()
	fmt.Println("Movies:")
	for _, m := range movies {
		fmt.Println(m.Title)
	}
}
