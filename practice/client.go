package main


import ("bufio"
	"fmt"
	"net"
	"os"
)

func read(conn net.Conn) {
	reader:= bufio.NewReader(conn)
	msg,_ :=reader.ReadString('\n')
	fmt.Printf(msg)
}

func main() {
	// msgP := flag.String("msg", "CDS","Message to send")
	// flag.Parse()
	stdin := bufio.NewReader(os.Stdin)
	conn, _ := net.Dial("tcp", "localhost:8030")
	for {
		fmt.Println("Enter message to send:")
		msg,_ := stdin.ReadString('\n')
		fmt.Fprintln(conn,msg)
		read(conn)
	}
}
