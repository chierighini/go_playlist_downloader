package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	print("Hello World")
	cmd := &cli.Command{
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fmt.Printf("Hello %q", cmd.Args().Get(0))
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

	requestURL := fmt.Sprintf("a")
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Print(res)
}
