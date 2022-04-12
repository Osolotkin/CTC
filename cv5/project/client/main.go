package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	gogo "project/project/gogo"
)

const CMD_PREFIX = "> ";

var client gogo.GoGoServiceClient;

func main() {

	var conn *grpc.ClientConn;
	conn, err := grpc.Dial(":8080", grpc.WithInsecure());
	if err != nil {
		log.Fatalf("Could not connect: %s", err);
		return;
	}
	defer conn.Close();

	client = gogo.NewGoGoServiceClient(conn);

	fmt.Printf("\nClient started\n==============\n");

	// loop input
	reader := bufio.NewReader(os.Stdin);

	for {
		
		fmt.Print(CMD_PREFIX);
		input, _ := reader.ReadString('\n');
		input = strings.TrimRight(input, "\n\r ");

		args := strings.Fields(input);
		cmd := args[0];

		fmt.Print(CMD_PREFIX);
		if strcmp("HELP", cmd) {

			printHelp();
		
		} else if strcmp("PING", cmd) {

			ping();
	
		} else if strcmp("GET", cmd) {

			get(args[1]);

		} else if strcmp("POST", cmd) {

			post(args[1], args[2]);
		
		} else if strcmp("LIST", cmd) {

			list(args[1]);

		} else if strcmp("DELETE", cmd) {

			delete(args[1]);
	
		} else {

			fmt.Print("No such command! Try HELP");

		}

		fmt.Println();
	
	}	

}

func strcmp(strA string, strB string) bool {

	return strings.EqualFold(strA, strB);

}

func printHelp() {

	fmt.Printf("Commands\n  --------\n\tPING\n\tGET\t <key>\n\tPOST\t <key> <value>\n\tLIST\t <key>\n\tDELETE\t <key>\n");

}

func ping() {

	rsp, err := client.Ping(context.Background(), &gogo.Message{Body: "Ping"});
	if err != nil {
		log.Fatalf("Error when calling Ping: %s", err);
		return;
	}
	fmt.Printf(rsp.Body);

}

func get(key string) {

	fmt.Printf("GET ans:\n");

	rsp, err := client.Get(context.Background(), &gogo.Message{Body: key});
	if err != nil {
		log.Fatalf("Get failed: %s", err);
		return;
	}
	fmt.Printf(rsp.Body);

}

func post(key string, val string) {

	fmt.Printf("POST ans:\n");

	rsp, err := client.Post(context.Background(), &gogo.KeyValuePair{Key: key, Val: val});
	if err != nil {
		log.Fatalf("Post failed: %s", err);
		return;
	}
	fmt.Printf(rsp.Body);

}

func list(key string) {

	fmt.Printf("LIST ans:\n");

	rsp, err := client.List(context.Background(), &gogo.Message{Body: key});
	if err != nil {
		log.Fatalf("List failed: %s", err);
		return;
	}
	fmt.Printf(rsp.Body);

}

func delete(key string) {

	fmt.Printf("DELETE ans:\n");

	rsp, err := client.Delete(context.Background(), &gogo.Message{Body: key});
	if err != nil {
		log.Fatalf("Delete failed: %s", err);
		return;
	}
	fmt.Printf(rsp.Body);

}
