package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/Gamilkarr/store/internal/handlers"
	"github.com/Gamilkarr/store/internal/repository"
)

func main() {
	ctx := context.Background()
	DBurl := "postgres://postgres:postgres@localhost:5432/storage"
	conn, err := pgxpool.New(ctx, DBurl)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	server := rpc.NewServer()
	if err := server.Register(&handlers.Store{
		Repository: &repository.Repository{
			Conn: conn}}); err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	go server.Accept(lis)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serverCodec := jsonrpc.NewServerCodec(&httpConn{
			in:  r.Body,
			out: w,
		})
		w.Header().Set("Content-type", "application/json")
		if err := server.ServeRequest(serverCodec); err != nil {
			log.Printf("Error while serving JSON request: %v", err)
			http.Error(w, "can not serve request", http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	})
	log.Fatal(http.ListenAndServe("127.0.0.1:8081", nil))

}

type httpConn struct {
	in  io.Reader
	out io.Writer
}

func (c *httpConn) Read(p []byte) (n int, err error)  { return c.in.Read(p) }
func (c *httpConn) Write(d []byte) (n int, err error) { return c.out.Write(d) }
func (c *httpConn) Close() error                      { return nil }
