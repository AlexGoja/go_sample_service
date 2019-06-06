package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"test/protos"
	"test/services"
	"test/storage/postgres"
)

var (
	grpcPort     = flag.Int("grpc_port", 1234, "GRPC port to bind to")
	httpPort     = flag.Int("http_port", 8080, "HTTP port to bind to")
	pgConnStr    = flag.String("pg_conn_str", "postgres://alexandrugoja:password@localhost/users", "posgres connection string")
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "users"
)

func init() {
	flag.Parse()
}

func startGRPCGateway(grpcServerAddress string, personService *services.PersonService) {

	log.Println("Starting HTTP1.1/GRPC gateway")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcMux := runtime.NewServeMux()

	// add services to gateway here
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := protos.RegisterPersonServiceHandlerFromEndpoint(ctx, grpcMux, grpcServerAddress, opts); err != nil {
		log.Fatal("Cannot start gateway: ", err)
	}

	router := mux.NewRouter()
	router.PathPrefix("/v1").Handler(grpcMux)
	router.PathPrefix("/").Handler(personService)

	http.ListenAndServe(fmt.Sprintf(":%d", *httpPort), router)
}

func main() {

	postgresStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	fmt.Println("Successfully connected!")

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcPort))
	if err != nil {
		log.Fatal("Count not start listener: ", err)
	}

	storage := postgres.NewPersonStorage(postgresStr)

	fmt.Println(storage)

	// Start RPC server with all services included
	rpcServer := grpc.NewServer()

	// create person service
	personService := services.NewPersonService(storage)

	// Add all services to this server
	protos.RegisterPersonServiceServer(rpcServer, personService)

	go startGRPCGateway(listener.Addr().String(), personService)

	//Start serving
	log.Println("Starting GRPC server")
	rpcServer.Serve(listener)
}
