




func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseFrorm()
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")

		hash, _ := pcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		_, err := db.Exec("INSER INTO users (username , email , password) VALUES (?, ?, ?)", username, email, hash)
		if err != nil {
			http.Error(w, "Username or email already exists", http.StatusBadRequest)
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	templ := template.Must(template.ParseFiles("templates/register.html"))
	tmpl.Execute(w, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		email := r.FormValue("email")
		password := r.FormValue("password")

		var hash string
		err := db.QueryRow("SELECT password_hash FROM users WHERE email = ?", email).Scan(&hash)
		if err != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
		if err != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/login.html"))
	tmpl.Execute(w, nil)
}