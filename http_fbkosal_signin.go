package main

import "net/http"

// handleFbKosalSignin 는 kosalaam 용 Firebase 로그인 페이지를 응답해주는 handler다.
func handleFbKosalSignin(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	err := TEMPLATES.ExecuteTemplate(w, "signin.html", "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
