package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
)

var (
	TEMPLATES     = template.New("")
	flagHTTPPort  = flag.String("http", "", "Web Service Port Number") // 웹서비스 포트
	flagCookieAge = flag.Int("cookieage", 4, "Cookie age (hour)")      // MPAA 기준 4시간
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("my-small-precious: ")
	flag.Parse()

	if *flagHTTPPort != "" {

		// 웹서버 실행
		ip, err := serviceIP()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Service start: http://%s\n", ip)
		webserver()

	} else {
		flag.PrintDefaults()
		os.Exit(1)
	}
}
