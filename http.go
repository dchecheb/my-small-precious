package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/shurcooL/httpfs/html/vfstemplate"
)

// LoadTemplates 는 asset 폴더 안의 모든 html 템플릿을 파싱하여 리턴하는 함수다.
func LoadTemplates() (*template.Template, error) {
	t := template.New("")
	t, err := vfstemplate.ParseGlob(assets, t, "/template/*.html")
	return t, err
}

// webserver 는 worktime의 웹서버를 실행하는 함수다. url에 따라 handler를 연결해주는 라우팅 기능도 수행한다.
func webserver() {

	// 템플릿 로딩을 위해 vfs(가상파일시스템) 로딩
	vfsTemplate, err := LoadTemplates()
	if err != nil {
		log.Fatal(err)
	}
	TEMPLATES = vfsTemplate

	// 리소스 로딩
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(assets)))

	// 웹주소 설정
	// http.HandleFunc("/", handleIndex)
	http.HandleFunc("/fb-kosal-signin", handleFbKosalSignin)
	// http.HandleFunc("/favicon.ico", handleIconFunc)

	// 웹서버 실행
	err = http.ListenAndServe(*flagHTTPPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}
