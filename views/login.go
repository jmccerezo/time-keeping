package views

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/jmccerezo/time_keeping/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func LogInHander(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/login.html"))
	data := map[string]interface{}{}

	//* Database connection
	dsn := "root:Allen is Great 200%@tcp(127.0.0.1:3306)/time_keeping?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Faied to Connect to the Database ", err)
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	users := []models.Users{}

	if r.Method == "POST" {
		rows := db.Where("username = ? and password = ?", username, password).Find(&users)
		if rows.RowsAffected > 0 {
			_ = db.Exec("USE time_keeping;")
			_ = db.Exec("UPDATE users SET time_in = ? WHERE username = ?", time.Now(), username)

			expires := time.Now().AddDate(1, 0, 0)
			ck := http.Cookie{
				Name:    "username",
				Path:    "/",
				Expires: expires,
			}
			ck.Value = username
			http.SetCookie(w, &ck)
			http.Redirect(w, r, "/dashboard/", http.StatusSeeOther)
		}
	}

	data["Title"] = "Log In | Time Keeping"
	tmpl.Execute(w, data)
}
