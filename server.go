package main

import (
	"controller"
	"db"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

//Model
type spaHandler struct {
	staticPath string
	indexPath  string
}

//Controller
func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func main() {
	// Init Router
	r := mux.NewRouter()
	// Set port from env file
	port := os.Getenv("PORT")

	originsOk := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"Accept", "application/json", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "X-CSRF-Token", "X-Requested-With"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// Create connection string and connect to db
	// db.ConnectionString()
	db.Lite()
	// Route Handlers / Endpoints
	r.HandleFunc("/api", apiHandler).Methods("GET")

	r.HandleFunc("/api/test/{path}/{id}", controller.Update2).Methods("GET")
	r.HandleFunc("/api/login/{username}/{password}", controller.LoginHandler).Methods("GET")

	r.HandleFunc("/api/funds", controller.FundsHandler).Methods("GET")
	r.HandleFunc("/api/fundstotals", controller.FundsTotalsHandler).Methods("GET")

	r.HandleFunc("/api/investments", controller.InvestmentsHandler).Methods("GET")
	r.HandleFunc("/api/investments/invid/{id}", controller.InvestmentHandler).Methods("GET")
	r.HandleFunc("/api/investmentsdetail", db.GetInvestmentsTbl).Methods("GET")

	r.HandleFunc("/api/cf/totals", controller.CFTotalsHandler).Methods("GET")
	r.HandleFunc("/api/cf/totals/monthly", controller.CFTotalsMonthHandler).Methods("GET")
	r.HandleFunc("/api/cf/totals/funds", controller.CFTotalsFundsHandler).Methods("GET")

	r.HandleFunc("/api/cf/invid/{id}", controller.CashflowInvidHandler).Methods("GET")
	r.HandleFunc("/api/cf/distro/invid", controller.CashflowInvidDistrosHandler).Methods("GET")

	r.HandleFunc("/api/investments/invid/{id}/update", controller.CashflowInvidUpdater).Methods("PUT")
	r.HandleFunc("/api/investments/invid/{id}/updatedetail", controller.DetailInvidUpdater).Methods("PUT")

	spa := spaHandler{staticPath: "client/dist", indexPath: "index.html"}
	r.PathPrefix("/").Handler(spa)

	// Start Server and Listen
	fmt.Println("Server Running On Port:", port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}
