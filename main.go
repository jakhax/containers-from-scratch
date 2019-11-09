package main;

import (
	"fmt"
	"os"
	"log"
	"os/exec"
);

// go run main.go run command args...
func main()  {
	if(len(os.Args) < 3){
		usage := "Usage: go run main.go cmd args...";
		log.Fatal(usage);
	}
	switch(os.Args[1]){
	case "run":
		run();
		break;
	default:
		log.Fatalf("Unknown command %s \n",os.Args[1])
	}
}


func run()  {
	fmt.Printf(" Running command: %s %v\n",os.Args[2],os.Args[3:]);
	cmd := 	exec.Command(os.Args[2],os.Args[3:]...);
	cmd.Stdin = os.Stdin;
	cmd.Stdout = os.Stdout;
	cmd.Stderr  = os.Stderr;
	err := cmd.Run();
	if(err != nil){
		log.Fatal(err);
	}
}