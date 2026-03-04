package main
import "html/template"

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// Abrir base de datos (usa la ruta completa para evitar problemas)
	db, err := sql.Open("sqlite3", "/mnt/c/Users/gongo/intro-go/series.db")
	if err != nil {
		log.Fatal("Error abriendo la base:", err)
	}
	defer db.Close()

	// Verificar conexión
	err = db.Ping()
	if err != nil {
		log.Fatal("No se pudo conectar a la base:", err)
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

    search := r.URL.Query().Get("search")

    var rows *sql.Rows
    var err error

    if search != "" {
        rows, err = db.Query(
            "SELECT name, current_episode, total_episodes FROM series WHERE name LIKE ?",
            "%"+search+"%",
        )
    } else {
        rows, err = db.Query(
            "SELECT name, current_episode, total_episodes FROM series",
        )
    }

    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    defer rows.Close()

    type Series struct {
        Name    string
        Current int
        Total   int
    }

    var seriesList []Series

    for rows.Next() {
        var s Series
        rows.Scan(&s.Name, &s.Current, &s.Total)
        seriesList = append(seriesList, s)
    }

    tmpl.Execute(w, seriesList)
})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
