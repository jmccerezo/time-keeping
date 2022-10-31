package views

import (
	"net/http"
	"text/template"

	"github.com/jmccerezo/time-keeping/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/update.html"))
	data := map[string]interface{}{}

	//* Database connection
	database_connection := "root:Allen is Great 200%@tcp(localhost:3306)/time_keeping?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(database_connection), &gorm.Config{})
	if err != nil {
		panic("Failed to Connect to the Database")
	}

	username := r.FormValue("username")
	firstname := r.FormValue("first_name")
	lastname := r.FormValue("last_name")

	users := []models.Users{}
	if r.Method == "POST" && username != "" {
		db.Model(&users).Where("username = ?", username).Updates(map[string]interface{}{"first_name": firstname, "last_name": lastname})
		data["Users"] = users
		http.Redirect(w, r, "/dashboard/", http.StatusSeeOther)
	}

	data["Title"] = "Update | Time Keeping"
	tmpl.Execute(w, data)
}
