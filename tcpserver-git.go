package main

import ("fmt"; "net"; "os";"encoding/json";"io/ioutil";"bufio")


// var for configuration
var (
	ConfigFile	= "config.json"
	CONN_HOST	= ""
	CONN_PORT	= ""
	CONN_TYPE	= ""
	MSG		= ""
	c 		Config
)


// struct for loading configuration information
type Config struct {

    CONN_HOST           string
    CONN_PORT           string
    CONN_TYPE		string
    MSG			string
}

func init() {
 // open our config and parse
   confBytes, e := ioutil.ReadFile(ConfigFile)
    if e != nil {
        fmt.Println(e.Error())
        os.Exit(1)
    }

    if e := json.Unmarshal(confBytes, &c); e != nil {
        fmt.Println(e.Error())
        os.Exit(2)
    }
fmt.Println(c)
}

func main() {
	// listen for incoming connections.
	l, err := net.Listen(c.CONN_TYPE, c.CONN_HOST + ":" + c.CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

// Close the listener when application closes.
defer l.Close()
fmt.Println("listening on ", c.CONN_PORT)
for {
	//listen for an incoming connection.
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting: ", err.Error())
		os.Exit(1)
	}
	// Handle connections in a new goroutine.
	go handleRequest(conn)
	}
}


//Handles incoming requests.
func handleRequest(conn net.Conn) {

	// display incoming message
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message from client: "+message)

	//Send a response back to person contacting us.
	conn.Write([]byte(c.MSG + "\n"))
	//Close the connection when done
	conn.Close()
}
