package main

import(
	"database/sql"
	"log"
	"net/http"
	"os"
	"text/template"
	_"github.com/lib/pq"
	
)

type Mahasiswa struct {
	Nim			int 	
	Nama 		string 	
	Jk			int 	
	Alamat 		string 	
	TempatLahir	string 	
	TglLahir 	string 	
	Agama 		int 	
}

type Jurusan struct {
	IdJurusan 	int 	
	NamaJurusan	string	
}

func dbConn() (db *sql.DB)  {
	if len(os.Getenv("ROACH_HOST")) ==0 {
		os.Setenv("ROACH_HOST", "localhost")
	}

	dbUser := "root"
	dbName := "kampus"
	db, err := sql.Open("postgres", "postgresql://"+dbUser+"@"+os.Getenv("ROACH_HOST")+":26257/"+dbName+"?sslmode=enable")
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl, err = template.ParseGlob("view/*")

func index(w http.ResponseWriter, r *http.Request) {
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title": "golang CRUD",
	}

	err = tmpl.ExecuteTemplate(w, "index", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("bulma/css/"))))
}

func dataJurusan(w http.ResponseWriter, r *http.Request) {
	var data = map[string]interface{}{
		"title": "golang CRUD",
	}

	err = tmpl.ExecuteTemplate(w, "dataJurusan", data)
}

//membuat main
func main() {
	log.Println("Server started on: http://localhost:8010")
	http.HandleFunc("/", index)
	http.HandleFunc("/dataJurusan", dataJurusan)

	http.ListenAndServe(":8010", nil)
}