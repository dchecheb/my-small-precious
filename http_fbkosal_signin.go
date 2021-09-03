package main

import (
	"net/http"
	"os"
)

// handleFbKosalSignin 는 kosalaam 용 Firebase 로그인 페이지를 응답해주는 handler다.
func handleFbKosalSignin(w http.ResponseWriter, r *http.Request) {

	// Firebase SDK 인증키
	type FirebaseKey struct {
		ApiKey            string
		AuthDomain        string
		ProjectId         string
		StorageBucket     string
		MessagingSenderId string
		AppId             string
		MeasurementId     string
	}

	f := FirebaseKey{}
	f.ApiKey = os.Getenv("FB_API_KEY")
	f.AuthDomain = os.Getenv("FB_AUTH_DOMAIN")
	f.ProjectId = os.Getenv("FB_PROJECT_ID")
	f.StorageBucket = os.Getenv("FB_STORAGE_BUCKET")
	f.MessagingSenderId = os.Getenv("FB_MESSAGING_SENDER_ID")
	f.AppId = os.Getenv("FB_APP_ID")
	f.MessagingSenderId = os.Getenv("FB_MEASUREMENT_ID")

	w.Header().Set("Content-Type", "text/html")
	err := TEMPLATES.ExecuteTemplate(w, "sign-in", f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
