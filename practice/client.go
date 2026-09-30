package main


import ("bufio"
	"fmt"
	"net"
	"flag"
)

func read(conn net.Conn) {
	reader:= bufio.NewReader(conn)
	msg,_ :=reader.ReadString('\n')
	fmt.Printf(msg)
}

func main() {
	msgP := flag.String("msg", "Default","Message to send")
	flag.Parse()
	conn, _ := net.Dial("tcp", "localhost:8030")
	fmt.Fprintln(conn,*msgP)
	
}
