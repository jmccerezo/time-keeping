package views

import (
	"net/http"
	"text/template"

	"github.com/jmccerezo/time-keeping/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/dashboard.html"))
	data := map[string]interface{}{}

	//* Database connection
	database_connection := "root:Allen is Great 200%@tcp(localhost:3306)/time_keeping?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(database_connection), &gorm.Config{})
	if err != nil {
		panic("Failed to Connect to the Database")
	}

	users := []models.Users{}
	db.Find(&users)

	username := r.FormValue("username")
	if r.Method == "POST" {
		if username != "" {
			username = "%" + username + "%"
			db.Where("username LIKE ?", username).Find(&users)
		}
	}

	data["Title"] = "Dashboard | Time Keeping"
	data["Users"] = users
	tmpl.Execute(w, data)
}
