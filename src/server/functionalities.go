package server

import (
	"database/sql"
	"forum/database"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var (
	templates *template.Template
)

func init() {
	templates = template.Must(template.ParseGlob(filepath.Join("templates", "*.html")))
}

type ErrorPageData struct {
	Code     string
	ErrorMsg string
}

type PageData struct {
	HasSession      bool
	UserID          int
	UserName        string
	Avatar          string
	RoleName        string
	TotalLikes      int
	TotalPosts      int
	Categories      []database.Category
	Users           []database.User
	Posts           []database.Post
	Reports         []database.Report
	TotalUsers      int
	TotalCategories int
	UserLogs        []database.UserLog
	UserSessions    []database.UserSession
	Notifications   []database.Notification
	RoleID          int
	Post            database.Post
	Comments        []database.Comment
	SelectedTab     string
	SelectedFilter  string
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func LikePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println("Method not allowed")
		err := ErrorPageData{Code: "405", ErrorMsg: "METHOD NOT ALLOWED"}
		errHandler(w, r, &err)
		return
	}

	postID, err := strconv.Atoi(r.FormValue("post_id"))
	if err != nil {
		log.Println("Invalid post ID")
		err := ErrorPageData{Code: "400", ErrorMsg: "BAD REQUEST"}
		errHandler(w, r, &err)
		return
	}

	userID, err := strconv.Atoi(r.FormValue("user"))
	if err != nil {
		log.Println("Error converting userID to int:", err)
		err := ErrorPageData{Code: "500", ErrorMsg: "INTERNAL SERVER ERROR"}
		errHandler(w, r, &err)
		return
	}

	db, err := sql.Open("sqlite3", "./database/main.db")
	if err != nil {
		log.Println("Error opening database:", err)
		err := ErrorPageData{Code: "500", ErrorMsg: "INTERNAL SERVER ERROR"}
		errHandler(w, r, &err)
		return
	}
	defer db.Close()

	err = database.ToggleLike(db, postID, userID)
	if err != nil {
		log.Println("Error toggling like:", err)
		err := ErrorPageData{Code: "500", ErrorMsg: "INTERNAL SERVER ERROR"}
		errHandler(w, r, &err)
		return
	}
	log.Println("Like toggled")
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
}

func DislikePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println("Method not allowed")
		err := ErrorPageData{Code: "405", ErrorMsg: "METHOD NOT ALLOWED"}
		errHandler(w, r, &err)
		return
	}

	postID, err := strconv.Atoi(r.FormValue("post_id"))
	if err != nil {
		log.Println("Invalid post ID")
		err := ErrorPageData{Code: "400", ErrorMsg: "BAD REQUEST"}
		errHandler(w, r, &err)
		return
	}

	userID, err := strconv.Atoi(r.FormValue("user"))
	if err != nil {
		log.Println("Error converting userID to int:", err)
		err := ErrorPageData{Code: "500", ErrorMsg: "INTERNAL SERVER ERROR"}
		errHandler(w, r, &err)
		return
	}

	db, err := sql.Open("sqlite3", "./database/main.db")
	if err != nil {
		log.Println("Error opening database:", err)
		err := ErrorPageData{Code: "500", ErrorMsg: "INTERNAL SERVER ERROR"}
		errHandler(w, r, &err)
		return
	}
	defer db.Close()

	err = database.ToggleDislike(db, postID, userID)
	if err != nil {
		log.Println("Error toggling dislike:", err)
		err := ErrorPageData{Code: "500", ErrorMsg: "INTERNAL SERVER ERROR"}
		errHandler(w, r, &err)
		return
	}
	log.Println("Dislike toggled")
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
}

func LikeComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println("Method not allowed")
		err := ErrorPageData{Code: "405", ErrorMsg: "METHOD NOT ALLOWED"}
		errHandler(w, r, &err)
		return
	}

	commentID, err := strconv.Atoi(r.FormValue("comment_id"))
	if err != nil {
		log.Println("Invalid comment ID")
		err := ErrorPageData{Code: "400", ErrorMsg: "BAD REQUEST"}
		errHandler(w, r, &err)
		return
	}

	userID, err := strconv.Atoi(r.FormValue("user"))
	if err != nil {
		log.Println("Error converting userID to int:", err)
		err := ErrorPageData{Code: "500", ErrorMsg: "INTERNAL SERVER ERROR"}
		errHandler(w, r, &err)
		return
	}

	db, err := sql.Open("sqlite3", "./database/main.db")
	if err != nil {
		log.Println("Error opening database:", err)
		err := ErrorPageData{Code: "500", ErrorMsg: "INTERNAL SERVER ERROR"}
		errHandler(w, r, &err)
		return
	}
	defer db.Close()

	err = database.ToggleCommentLike(db, commentID, userID)
	if err != nil {
		log.Println("Error toggling like:", err)
		err := ErrorPageData{Code: "500", ErrorMsg: "INTERNAL SERVER ERROR"}
		errHandler(w, r, &err)
		return
	}
	log.Println("Like toggled")
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
}

func DislikeComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println("Method not allowed")
		err := ErrorPageData{Code: "405", ErrorMsg: "METHOD NOT ALLOWED"}
		errHandler(w, r, &err)
		return
	}

	commentID, err := strconv.Atoi(r.FormValue("comment_id"))
	if err != nil {
		log.Println("Invalid comment ID")
		err := ErrorPageData{Code: "400", ErrorMsg: "BAD REQUEST"}
		errHandler(w, r, &err)
		return
	}

	userID, err := strconv.Atoi(r.FormValue("user"))
	if err != nil {
		log.Println("Error converting userID to int:", err)
		err := ErrorPageData{Code: "500", ErrorMsg: "INTERNAL SERVER ERROR"}
		errHandler(w, r, &err)
		return
	}

	db, err := sql.Open("sqlite3", "./database/main.db")
	if err != nil {
		log.Println("Error opening database:", err)
		err := ErrorPageData{Code: "500", ErrorMsg: "INTERNAL SERVER ERROR"}
		errHandler(w, r, &err)
		return
	}
	defer db.Close()

	err = database.ToggleCommentDislike(db, commentID, userID)
	if err != nil {
		log.Println("Error toggling dislike:", err)
		err := ErrorPageData{Code: "500", ErrorMsg: "INTERNAL SERVER ERROR"}
		errHandler(w, r, &err)
		return
	}
	log.Println("Dislike toggled")
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Println("Method not allowed")
		err := ErrorPageData{Code: "405", ErrorMsg: "METHOD NOT ALLOWED"}
		errHandler(w, r, &err)
		return
	}

	postID := r.URL.Query().Get("id")
	if postID == "" {
		log.Println("Post ID is missing")
		err := ErrorPageData{Code: "400", ErrorMsg: "BAD REQUEST"}
		errHandler(w, r, &err)
		return
	}

	db, err := sql.Open("sqlite3", "./database/main.db")
	if err != nil {
		log.Println("Error opening database:", err)
		err := ErrorPageData{Code: "500", ErrorMsg: "INTERNAL SERVER ERROR"}
		errHandler(w, r, &err)
		return
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM post WHERE postid = ?", postID)
	if err != nil {
		log.Println("Error deleting post:", err)
		err := ErrorPageData{Code: "500", ErrorMsg: "INTERNAL SERVER ERROR"}
		errHandler(w, r, &err)
		return
	}

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func ReportPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Println("Method not allowed")
		err := ErrorPageData{Code: "405", ErrorMsg: "METHOD NOT ALLOWED"}
		errHandler(w, r, &err)
		return
	}

	postID := r.URL.Query().Get("id")
	if postID == "" {
		log.Println("Post ID is missing")
		err := ErrorPageData{Code: "400", ErrorMsg: "BAD REQUEST"}
		errHandler(w, r, &err)
		return
	}

	db, err := sql.Open("sqlite3", "./database/main.db")
	if err != nil {
		log.Println("Error opening database:", err)
		err := ErrorPageData{Code: "500", ErrorMsg: "INTERNAL SERVER ERROR"}
		errHandler(w, r, &err)
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO reports (post_id, reported_by, report_reason) VALUES (?, ?, ?)", postID, r.URL.Query().Get("user"), "Reported by moderator")
	if err != nil {
		log.Println("Error reporting post:", err)
		err := ErrorPageData{Code: "500", ErrorMsg: "INTERNAL SERVER ERROR"}
		errHandler(w, r, &err)
		return
	}

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}


func CheckFilter(filter string,categoryNames []string) bool {
	for _, category := range categoryNames {
		if filter == category {
			return true

		}
	}
	return false
}