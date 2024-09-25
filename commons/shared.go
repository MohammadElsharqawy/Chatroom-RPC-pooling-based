package commons

type Message struct {
	User    string
	Message string
}

func Get_server_address() string {
	return "0.0.0.0:9999"
}