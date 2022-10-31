package views

import (
	"net/http"
	"text/template"

	"github.com/jmccerezo/time-keeping/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/delete.html"))
	data := map[string]interface{}{}

	//* Database connection
	database_connection := "root:Allen is Great 200%@tcp(localhost:3306)/time_keeping?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(database_connection), &gorm.Config{})
	if err != nil {
		panic("Failed to Connect to the Database")
	}

	username := r.FormValue("username")
	users := []models.Users{}
	if r.Method == "POST" && username != "" {
		db.Where("username = ?", username).Delete(&users)
		data["Users"] = users

		http.Redirect(w, r, "/dashboard/", http.StatusSeeOther)
	}

	data["Title"] = "Delete | Time Keeping"
	tmpl.Execute(w, data)
}
