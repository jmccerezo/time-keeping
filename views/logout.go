package views

import (
	"net/http"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	//* Database connection
	database_connection := "root:Allen is Great 200%@tcp(localhost:3306)/time_keeping?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(database_connection), &gorm.Config{})
	if err != nil {
		panic("Failed to Connect to the Database")
	}

	cookie, _ := r.Cookie("username")
	_ = db.Exec("USE time_keeping;")
	_ = db.Debug().Exec("UPDATE users SET time_out = ? WHERE username = ?", time.Now(), cookie.Value)

	ck := http.Cookie{
		Name:    "username",
		Path:    "/",
		Expires: time.Now(),
	}

	http.SetCookie(w, &ck)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
