package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// Data Models
type Employee struct {
	ID        int
	Name      string
	Position  string
	StartDate time.Time
}

type WorkSchedule struct {
	ID        int
	EmployeeID int
	Date      time.Time
	Shift     string
}

// Database Connection
var db *sqlx.DB

func initDB() {
	db, err := sqlx.Open("sqlite3", "./employeemanagement.db")
	if err != nil {
		log.Fatal(err)
	}

	db.MustExec(`CREATE TABLE IF NOT EXISTS employees (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		position TEXT,
		start_date TEXT
	)`)

	db.MustExec(`CREATE TABLE IF NOT EXISTS work_schedules (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		employee_id INTEGER,
		date TEXT,
		shift TEXT
	)`)
}

// Functions for Employee Management
func addEmployee(employee *Employee) {
	_, err := db.Exec("INSERT INTO employees (name, position, start_date) VALUES (?, ?, ?)", employee.Name, employee.Position, employee.StartDate)
	if err != nil {
		log.Fatal(err)
	}
}

func getEmployees() []Employee {
	var employees []Employee
	err := db.Select(&employees, "SELECT * FROM employees")
	if err != nil {
		log.Fatal(err)
	}
	return employees
}

// Functions for Work Schedule Management
func addWorkSchedule(workSchedule *WorkSchedule) {
	_, err := db.Exec("INSERT INTO work_schedules (employee_id, date, shift) VALUES (?, ?, ?)", workSchedule.EmployeeID, workSchedule.Date, workSchedule.Shift)
	if err != nil {
		log.Fatal(err)
	}
}

// Main Function
func main() {
	initDB()
	defer db.Close()

	employee := &Employee{Name: "Alice", Position: "Manager", StartDate: time.Now()}
	addEmployee(employee)

	employees := getEmployees()
	fmt.Println("Employees:")
	for _, emp := range employees {
		fmt.Println(emp.Name)
	}
}
