package main

import (
	
	"net/http"
	"log"
	"fmt"
	"os"
	"github.com/go-chi/chi/v5"
	"github.com/farisbrandone/chirpy_project/internal/database"
	"github.com/joho/godotenv"
	"github.com/golang-jwt/jwt/v5"
    
)

type MyCustomClaims struct {
	UserToken database.UserEmailPassword `json:user`
	jwt.RegisteredClaims
}


type apiConfig struct {
	fileserverHits int
	jwtSecret string
	jwtRefreshSecret string
	apiKey string
	
}
type Bol struct {
	Body string `json:"body"`
	Id int `json:"id"`
	AuthorId int `json:"author_id"`
}
type User struct {
	Email string `json:"email"`
	Id int `json:"id"`
	Password string `json:"password"`
}
 type ArrayBol []Bol
 type ArrayUser []User

 type dbConfig struct {
	config apiConfig
	db     *(database.DB)
 }

 




func main() {
	// by default, godotenv will look for a file named .env in the current directory
    err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
	jwtSecret := os.Getenv("JWT_SECRET")
    jwtRefreshSecret := os.Getenv("REFRESH_TOKEN")
	apiKey :=os.Getenv("API_KEY")

	myDb, err:=database.NewDB("file.json")
		if err!=nil {
			log.Println("Something went wrong when we create data")
				return
		 }
	apiCfg:=apiConfig{
		fileserverHits:0,
		jwtSecret : jwtSecret,
		jwtRefreshSecret: jwtRefreshSecret,
		apiKey: apiKey,
	}
	totalConfig:=dbConfig{
		config: apiCfg,
		db: myDb,
	}
	//data:=ArrayBol{}
	//user:=ArrayUser{}
	
    const filepathRoot = "."
	const port = "8080"

    r := chi.NewRouter()
    //mux := http.NewServeMux()  
	corsMux := middlewareCors(r)
	r.Get("/healthz",  handlerReadiness)
	//r.Post(handlerNo)
    r.HandleFunc("/healthz",  handlerReadiness) // juse une fonction passer en callback
	fsHandler :=totalConfig.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))
    r.Handle("/app/*" ,fsHandler )// r.Handle for middleware because all pass
    //r.Handle( "/app", fsHandler)
    r.Get("/metrics",  totalConfig.handlerHits)
	r.Get("/reset",  totalConfig.handlerReset)
	r.Mount("/api", totalConfig.adminRouter())
	r.Mount("/admin", totalConfig.admin())
	/* mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))))
	//ici une function qui retourne une handlefunc passer en callbbesck, la handlefunc retourner recupere la requete.
	mux.HandleFunc("/metrics",  apiCfg.handlerHits)
	mux.HandleFunc("/reset",  apiCfg.handlerReset) */

	/* apiRouter := chi.NewRouter()
	apiRouter.Get("/healthz", handlerReadiness)
	apiRouter.Get("/metrics", apiCfg.handlerMetrics)
	apiRouter.Get("/reset", apiCfg.handlerReset)
	router.Mount("/api", apiRouter) */
	
	
	srv := &http.Server{ 
		Addr:    ":" + port,
		Handler: corsMux,
	}
	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe()) 
	
}
 func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {//handler retourné, recois les requète les traites et pass a next paramètre d'entrée
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		if r.Method == "POST" && (r.URL.Path=="/healthz" || r.URL.Path=="/metrics") {
			log.Println("Serving on port: ", r.URL.Path)
	        w.WriteHeader(http.StatusMethodNotAllowed)
	        w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
			return
		}
		next.ServeHTTP(w, r)
	})
} 

func handlerReadiness(w http.ResponseWriter, r *http.Request) { // founction a executer pour une requette donnée
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}



func (cfg *dbConfig) handlerHits(w http.ResponseWriter, r *http.Request) { // founction a executer pour une requette donnée
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	a:=fmt.Sprintf("Hits: %d", cfg.config.fileserverHits)
	w.Write([]byte(a))
}
func (cfg *dbConfig) handlerReset(w http.ResponseWriter, r *http.Request) { // founction a executer pour une requette donnée
	cfg.config.fileserverHits=0
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	//w.Write([]byte("Hits: x"))
}
 
func (cfg *dbConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	const filepathRoot = "."
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {//recois les requete exécute et pass a next son parametre d'entrée
		cfg.config.fileserverHits ++
		w.Header().Add("Cache-Control", "no-cache")
		//w.WriteHeader(http.StatusOK)
		next.ServeHTTP(w, r)// ici, http.StripPrefix recupere la requete et la traite pour renvoyé les fichier statique
	})
} 
func (cfg *dbConfig) handlerHitsHtml(w http.ResponseWriter, r *http.Request) { // founction a executer pour une requette donnée
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	a:=fmt.Sprintf("<html><body><h1>Welcome, Chirpy Admin</h1><p>Chirpy has been visited %d times!</p></body></html>", cfg.config.fileserverHits)
	w.Write([]byte(a))
}

func (cfg *dbConfig) adminRouter(/* data ArrayBol, user ArrayUser */) chi.Router {
	r := chi.NewRouter()
	//r.Use(AdminOnly)
	r.Get("/healthz",  handlerReadiness)
	r.Get("/metrics",  cfg.handlerHits)
	r.Get("/reset",  cfg.handlerReset)
	r.Post("/chirps",  cfg.handlerPostVal)//data
	r.Get("/chirps",  cfg.handlerPostValGet)
	r.Get("/chirps/{chirpId}",  cfg.handlerPostValGetId)
	r.Post("/users",  cfg.handlerPostUser)
	r.Put("/users", cfg.handlerPutUser)
	r.Post("/login",  cfg.handlerLoginUser)
	r.Post("/refresh",  cfg.handlerPostRefresh)
	r.Post("/revoke",  cfg.handlerPostRevoque)
	r.Post("/polka/webhooks",  cfg.handlerPostWebhooks)
	r.Delete("/chirps/{chirpId}", cfg.handlerDelete)
	
	return r
}
func (cfg *dbConfig) admin() chi.Router {
	r := chi.NewRouter()
	r.Get("/metrics",  cfg.handlerHitsHtml)
	return r
}




	

	

		
		
	






