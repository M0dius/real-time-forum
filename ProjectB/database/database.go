package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func DataBase() {
	db, err := sql.Open("sqlite3", "./database/main.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var tableName string
	err = db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='categories'").Scan(&tableName)
	if err == nil && tableName == "categories" {
		log.Println("Database already exists. Skipping table creation.")
		return
	}

	const CreateCategoriesTable = `
		CREATE TABLE IF NOT EXISTS categories (
			idcategories INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NULL
		);`

	const CreateCommentTable = `
		CREATE TABLE IF NOT EXISTS comment (
			commentid INTEGER PRIMARY KEY AUTOINCREMENT,
			content TEXT NULL,
			comment_at DATETIME NULL,
			post_postid INTEGER NOT NULL,
			user_userid INTEGER NOT NULL,
			FOREIGN KEY (post_postid) REFERENCES post(postid),
			FOREIGN KEY (user_userid) REFERENCES user(userid)
		);`

	const CreateDislikeTable = `
		CREATE TABLE IF NOT EXISTS dislikes (
			dislikeid INTEGER PRIMARY KEY AUTOINCREMENT,
			dislike_at DATETIME NULL,
			post_postid INTEGER NOT NULL,
			user_userid INTEGER NOT NULL,
			FOREIGN KEY (post_postid) REFERENCES post(postid),
			FOREIGN KEY (user_userid) REFERENCES user(userid)
		);`

	const CreateLikeTable = `
		CREATE TABLE IF NOT EXISTS likes (
			likeid INTEGER PRIMARY KEY AUTOINCREMENT,
			like_at DATETIME NULL,
			post_postid INTEGER NOT NULL,
			user_userid INTEGER NOT NULL,
			FOREIGN KEY (post_postid) REFERENCES post(postid),
			FOREIGN KEY (user_userid) REFERENCES user(userid)
		);`

	const CreateCommentDislikeTable = `
		CREATE TABLE IF NOT EXISTS comment_dislikes (
			dislikeid INTEGER PRIMARY KEY AUTOINCREMENT,
			dislike_at DATETIME NULL,
			commentid INTEGER NOT NULL,
			userid INTEGER NOT NULL,
			FOREIGN KEY (commentid) REFERENCES comment(commentid)
			FOREIGN KEY (userid) REFERENCES user(userid)
		);`

	const CreateCommentLikeTable = `
		CREATE TABLE IF NOT EXISTS comment_likes (
			likeid INTEGER PRIMARY KEY AUTOINCREMENT,
			like_at DATETIME NULL,
			commentid INTEGER NOT NULL,
			userid INTEGER NOT NULL,
			FOREIGN KEY (commentid) REFERENCES comment(commentid)
			FOREIGN KEY (userid) REFERENCES user(userid)
		);`

	const CreatePostTable = `
		CREATE TABLE IF NOT EXISTS post (
			postid INTEGER PRIMARY KEY AUTOINCREMENT,
			image TEXT NULL,
			content TEXT NULL,
			post_at DATETIME NOT NULL,
			user_userid INTEGER NOT NULL,
			FOREIGN KEY (user_userid) REFERENCES user(userid)
		);`

	const CreatePostHasCategoriesTable = `
		CREATE TABLE IF NOT EXISTS post_has_categories (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			post_postid INTEGER NOT NULL,
			categories_idcategories INTEGER NOT NULL,
			FOREIGN KEY (post_postid) REFERENCES post(postid),
			FOREIGN KEY (categories_idcategories) REFERENCES categories(idcategories)
		);`

	const CreateSessionsTable = `
		CREATE TABLE IF NOT EXISTS session (
			sessionid TEXT PRIMARY KEY,
			userid INTEGER NOT NULL UNIQUE,
			endtime DATETIME NOT NULL,
			FOREIGN KEY (userid) REFERENCES user(userid)
		);
	`

	const CreateUserTable = `
		CREATE TABLE IF NOT EXISTS user (
			userid INTEGER PRIMARY KEY AUTOINCREMENT,
			F_name TEXT NOT NULL,
			L_name TEXT NOT NULL,
			Username TEXT NOT NULL,
			Email TEXT NOT NULL,
			password TEXT NOT NULL,
			current_session TEXT,
			role_id INTEGER NOT NULL,
			Avatar TEXT,
			FOREIGN KEY (current_session) REFERENCES session(sessionid),
			FOREIGN KEY (role_id) REFERENCES user_roles(roleid)
		);
		`

	const CreateUserRolesTable = `
		CREATE TABLE IF NOT EXISTS user_roles (
			roleid INTEGER PRIMARY KEY AUTOINCREMENT,
			role_name TEXT NOT NULL
		);
		

		`

	const CreateFriendsTable = `
		CREATE TABLE IF NOT EXISTS friends (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_userid INTEGER NOT NULL,
			friend_userid INTEGER NOT NULL,
			FOREIGN KEY (user_userid) REFERENCES user(userid),
			FOREIGN KEY (friend_userid) REFERENCES user(userid)
		);`

	const CreateFollowersTable = `
		CREATE TABLE IF NOT EXISTS followers (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_userid INTEGER NOT NULL,
			follower_userid INTEGER NOT NULL,
			FOREIGN KEY (user_userid) REFERENCES user(userid),
			FOREIGN KEY (follower_userid) REFERENCES user(userid)
		);`

	const CreateNotificationsTable = `
		CREATE TABLE IF NOT EXISTS notifications (
			notificationid INTEGER PRIMARY KEY AUTOINCREMENT,
			user_userid INTEGER NOT NULL,
			post_id INTEGER NOT NULL,
			message TEXT NOT NULL,
			created_at DATETIME default CURRENT_TIMESTAMP,
			FOREIGN KEY (user_userid) REFERENCES user(userid),
			FOREIGN KEY (post_id) REFERENCES post(postid)
		);`

	const CreateFollowingTable = `
		CREATE TABLE IF NOT EXISTS following (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_userid INTEGER NOT NULL,
			following_userid INTEGER NOT NULL,
			FOREIGN KEY (user_userid) REFERENCES user(userid),
			FOREIGN KEY (following_userid) REFERENCES user(userid)
		);`

	const CreateReportsTable = `
		CREATE TABLE IF NOT EXISTS reports (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			post_id INTEGER NOT NULL,
			reported_by INTEGER NOT NULL,
			report_reason TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (post_id) REFERENCES post(postid),
			FOREIGN KEY (reported_by) REFERENCES user(userid)
		);`

	createTableStatements := []string{
		CreateCategoriesTable,
		CreateCommentTable,
		CreateDislikeTable,
		CreateLikeTable,
		CreateCommentDislikeTable,
		CreateCommentLikeTable,
		CreatePostTable,
		CreatePostHasCategoriesTable,
		CreateSessionsTable,
		CreateUserTable,
		CreateUserRolesTable,
		CreateFriendsTable,
		CreateFollowersTable,
		CreateNotificationsTable,
		CreateFollowingTable,
		CreateReportsTable,
	}

	for _, stmt := range createTableStatements {
		_, err = db.Exec(stmt)
		if err != nil {
			log.Fatal(err)
		}
	}

	insertCategories := []string{
		`INSERT INTO categories (name, description) VALUES ('AI & ML', 'All about Artificial Intelligence and Machine Learning');`,
		`INSERT INTO categories (name, description) VALUES ('Cloud & DevOps', 'Cloud infrastructure and DevOps best practices');`,
		`INSERT INTO categories (name, description) VALUES ('Cybersecurity', 'Guides and insights on staying secure online');`,
		`INSERT INTO categories (name, description) VALUES ('Blockchain & Web3', 'Decentralized networks and blockchain technologies');`,
		`INSERT INTO categories (name, description) VALUES ('AR/VR & Gaming', 'Immersive technologies and game development');`,
		`INSERT INTO categories (name, description) VALUES ('UI/UX Design', 'Improving digital experiences through design');`,
		`INSERT INTO categories (name, description) VALUES ('IoT & Edge Computing', 'Internet of Things and edge computing trends');`,
		`INSERT INTO categories (name, description) VALUES ('Data Analytics', 'Extracting insights from big data');`,
		`INSERT INTO categories (name, description) VALUES ('Quantum Computing', 'Next-gen computing paradigms and qubits');`,
		`INSERT INTO categories (name, description) VALUES ('SRE & Observability', 'Site Reliability Engineering and system observability');`,
	}

	insertUserRoles := []string{
		`INSERT INTO user_roles (role_name) VALUES ('Admin');`,
		`INSERT INTO user_roles (role_name) VALUES ('Moderator');`,
		`INSERT INTO user_roles (role_name) VALUES ('User');`,
		`INSERT INTO user_roles (role_name) VALUES ('Guest');`,
	}

	insertUser := `INSERT INTO user (F_name, L_name, Username, Email, password, current_session, role_id, Avatar) VALUES ('Alicia', 'Nguyen', 'aliceN', 'aliceN@example.com', 'alicePass', 1, 1, 'https://randomuser.me/api/portraits/women/1.jpg');`

	db.Exec(insertUser)

	allInserts := [][]string{
		insertCategories,
		insertUserRoles,
	}

	fmt.Print(allInserts)

	for _, group := range allInserts {
		for _, stmt := range group {
			_, err := db.Exec(stmt)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
