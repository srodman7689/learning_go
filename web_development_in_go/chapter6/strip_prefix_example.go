package main

import (
	"net/http"
)

func main() {
	// To serve a directory on disk (/public) under an alternative URL
	// path (/public/), use StripPrefix to modify the request
	// URL's path before the FileServer sees it:
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))
}
