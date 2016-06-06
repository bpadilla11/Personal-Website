/* Personal website handler in GoLang.
 * Author: Brian Padilla, bpadilla614@gmail.com
 * This work is licensed under a Creative Commons Attribution 4.0 International License.
 */

package personalWebsite

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

// handle the main page
func index(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.html")
	if err != nil { // if file does not exist, give user a error
		log.Fatalln(err) // stops program if file does not exist
	}
	tpl.Execute(res, nil) // execute the html file
}

// initalize the website
func init() {
	r := mux.NewRouter() // create new router

	// handle URLs
	r.HandleFunc("/", index)
	http.Handle("/", r)

	// handle elements
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("assets"))))
	http.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("images"))))
}
