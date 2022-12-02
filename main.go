package main

func init() {

}

func main() {
	StartServer()
}

type TwilioIncomingMessageBody struct {
	From string
	Body string
}

type Response struct {
	Message struct {
		Body string
	}
}
