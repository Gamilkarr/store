package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/Gamilkarr/store/internal/handlers"
)

func main() {
	server := rpc.NewServer()
	if err := server.Register(new(handlers.Store)); err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatalf("net.Listen tcp :0: %v", err)
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
			http.Error(w, `{"error":"cant serve request"}`, 500)
		} else {
			w.WriteHeader(200)
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