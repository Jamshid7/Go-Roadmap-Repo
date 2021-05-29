package main
import (
	"log"
	
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"example.com/chat-server/chat"	
)

func main() {
	var conn *grpc.ClientConn
	conn,err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldn't connect: %s", err)
	}
	defer conn.Close()
	
	c := chat.NewChatServiceClient(conn)
	
	message := chat.Message {
		Body: "Hello from ContactList client",
	}
	
	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	
	log.Printf("Response from Server: %s", response.Body)
}
