tcpsock_posix.go
dialTCP(ctx context.Context, laddr, raddr *TCPAddr) (*TCPConn, error)
listenTCP(ctx context.Context, laddr *TCPAddr) (*TCPListener, error)
dialUDP(ctx context.Context, laddr, raddr *UDPAddr) (*UDPConn, error)
listenUDP(ctx context.Context, laddr *UDPAddr) (*UDPConn, error)

sql.go
QueryContext(ctx context.Context, query string, args ...interface{}) 
PrepareContext(ctx context.Context, query string)
