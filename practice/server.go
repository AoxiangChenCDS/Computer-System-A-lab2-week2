package main


import ("bufio"
	"fmt"
	"net"
	
)


func handleConnection(conn net.Conn) {
	for {
		reader:= bufio.NewReader(conn)
		msg,_ :=reader.ReadString('\n')
		fmt.Println(msg)
		fmt.Fprintln(conn,"Message received")
	}
}

func main() {
	ln,_ := net.Listen("tcp", ":8030")
	for{
		conn,_ :=ln.Accept()
		go handleConnection(conn)
	}

}
