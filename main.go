package main

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

type Client struct {
	ID      int    `json:"id"`
	Fio     string `json:"fio"`
	City    string `json:"city"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	IP      string `json:"ip"`
	Login   string `json:"login"`
}

var db *sql.DB

// --- СЕССИИ ---
var store = sessions.NewCookieStore([]byte("super-secret-key"))

func init() {
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400, // сутки
		HttpOnly: true,
	}
}

// --- ТЕСТОВЫЕ ПОЛЬЗОВАТЕЛИ ---
var users = map[string]string{
	"admin": "1234",
	"artem": "qwerty",
}

func main() {
	godotenv.Load()

	var err error
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") +
		"@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" +
		os.Getenv("DB_NAME")

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("БД подключена 🔥")

	// --- AUTH ---
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	// --- API (ЗАЩИЩЕНО) ---
	http.HandleFunc("/api/clients", authMiddleware(clientsHandler))
	http.HandleFunc("/api/clients/create", authMiddleware(createClient))
	http.HandleFunc("/api/clients/update", authMiddleware(updateClient))
	http.HandleFunc("/api/clients/delete", authMiddleware(deleteClient))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	
	// --- ГЛАВНАЯ ---
	http.HandleFunc("/", authMiddleware(dashboard))

	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// --- MIDDLEWARE ---
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")
		_, ok := session.Values["user"].(string)

		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next(w, r)
	}
}

// --- LOGIN ---
func login(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/login.html"))

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		log.Println("LOGIN TRY:", username, password)

		if pass, ok := users[username]; ok && pass == password {
			session, _ := store.Get(r, "session")
			session.Values["user"] = username
			session.Save(r, w)
			log.Println("LOGIN SUCCESS:", username)

			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	tmpl.Execute(w, nil)
}

// --- LOGOUT ---
func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	delete(session.Values, "user")
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// --- DASHBOARD ---
func dashboard(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	user, _ := session.Values["user"].(string)

	tmpl := template.Must(template.ParseFiles("templates/dashboard.html"))
	tmpl.Execute(w, user)
}

// --- API ---

func clientsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, fio, city, address, phone, ip, login FROM clients")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	clients := []Client{}
	for rows.Next() {
		var c Client
		if err := rows.Scan(&c.ID, &c.Fio, &c.City, &c.Address, &c.Phone, &c.IP, &c.Login); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		clients = append(clients, c)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": clients,
	})
}

func createClient(w http.ResponseWriter, r *http.Request) {
	var c Client
	json.NewDecoder(r.Body).Decode(&c)

	_, err := db.Exec(`
		INSERT INTO clients (fio, city, address, phone, ip, login)
		VALUES (?, ?, ?, ?, ?, ?)`,
		c.Fio, c.City, c.Address, c.Phone, c.IP, c.Login)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte(`{"status":"ok"}`))
}

func updateClient(w http.ResponseWriter, r *http.Request) {
	var c Client
	json.NewDecoder(r.Body).Decode(&c)

	_, err := db.Exec(`
		UPDATE clients SET fio=?, city=?, address=?, phone=?, ip=?, login=?
		WHERE id=?`,
		c.Fio, c.City, c.Address, c.Phone, c.IP, c.Login, c.ID)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte(`{"status":"ok"}`))
}

func deleteClient(w http.ResponseWriter, r *http.Request) {
	var data struct {
		ID int `json:"id"`
	}

	json.NewDecoder(r.Body).Decode(&data)

	_, err := db.Exec("DELETE FROM clients WHERE id=?", data.ID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte(`{"status":"ok"}`))
}