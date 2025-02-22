package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"goModule/snippetbox/pkg/models/mysql" // New import

	_ "github.com/go-sql-driver/mysql"
	"github.com/golangcollege/sessions" // New import
)

type contextKey string

var contextKeyUser = contextKey("user")



// To make the logger available to all the files we create a custom struct against which we will be calling handlers
type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	session       *sessions.Session
	users         *mysql.UserModel
	snippets      *mysql.SnippetModel
	templateCache map[string]*template.Template
}

// The openDB() function wraps sql.Open() and returns a sql.DB connection pool
// for a given DSN.
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func main() {

	err := godotenv.Load()

	 if err != nil {
        log.Println("No .env file found, using system environment variables.")
    }
	
	// For configurations
	addr := flag.String("addr", ":"+os.Getenv("PORT"), "HTTP Network Address")
	dsn := flag.String("dsn",os.Getenv("DB_DSN"), "MySQL database")
	secret := flag.String("secret", os.Getenv("SECRET_KEY"), "Secret")
	flag.Parse()

	//Now for improved leveled logging below are the loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	//We defer db.Close(), so that the connection pool is closed before the main() function exits.
	defer db.Close()

	// Initialize a new template cache...
	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour

	// And add it to the application dependencies.
	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		session:       session,
		users:         &mysql.UserModel{DB: db},
		snippets:      &mysql.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	//tls.Config struct to hold the non-default TLS settings we want the server to use.
	// tlsConfig := &tls.Config{
	// 	PreferServerCipherSuites: true,
	// 	CipherSuites: []uint16{
	// 		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
	// 		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
	// 		tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
	// 		tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
	// 		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
	// 		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
	// 	},
	// }

	//new http Server instead of using default one cause it use standard logger for http error so we creating
	// new server and use it so that we can user our custiom error logger to log error message
	srv := http.Server{
		Addr:         *addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		// TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	errorLog.Fatal(err)
}
