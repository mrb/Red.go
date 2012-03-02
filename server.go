package main

import (
    "fmt" 
    "net"
    "os"
)

type Server struct {
    Db *Db
}

func NewServer(db *Db) *Server {
    return &Server{Db: db}
}

func (s *Server) Start() {
    listener, err := net.Listen("tcp", "0.0.0.0:6380")
    if err != nil {
        fmt.Printf("Can't start on 6380\n")
        os.Exit(1)
    }

    // Listen
    fmt.Printf("Listening on port 6380\n")
    // fmt.Printf("value or whatev: ", listener)
    for {
        // var read = true
        conn, err := listener.Accept()
        if err != nil {
            fmt.Printf("Can't accept connections\n")
            os.Exit(1)
        }

        fmt.Printf("New connection from: %s\n", conn.RemoteAddr())
        
        // Launch some go routine with connection
        fmt.Printf("NewConnFromStart: ", conn)
        fmt.Printf("NewConnFromStartWithPointer: ", &conn)
        go s.handleConn(conn)
    }
}

func (s *Server) handleConn(conn net.Conn) {
    fmt.Printf("NewConnFromCreateClient: ", conn)
    
    c := NewClient(s.Db, conn)
    fmt.Printf("NewClient: ", c)
    
    // c.InjectDb(db)

    c.ProcessRequest()
    fmt.Printf("Client DUMP: %#v", c)
}