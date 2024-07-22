package main
import (
	"log"
	"os"
	_"github.com/lib/pq"
	"Backend/project/Handler"
	"Backend/project/DB_collection"
	"Backend/project/DB"
	"Backend/project/Router"
	
	
)
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r:=router.New()

	// Set up PostgreSQL connection
	postgresDB, err := db.GetPostgresDB()
	if err != nil {
		log.Fatal(err)
	}
	defer postgresDB.Close()

	userstore := dbcollection.NewUserStore(postgresDB)

    h := handler.NewHandler(userstore)
	g :=r.Group("/api")
	h.RegisterRoutes(g)

	// Start server
	r.Logger.Fatal(r.Start("0.0.0.0:" + port))

}