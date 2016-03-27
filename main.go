// docker-container-ip prints out the network IP Address of a container
package main

import (
	"fmt"
	"github.com/docker/engine-api/client"
	"golang.org/x/net/context"
	"os"
)

func main() {

	if len(os.Args[1:]) == 0 {
		fmt.Fprintln(os.Stderr, "You must provide a container id/name")
		os.Exit(1)
	}

	defaultHeaders := map[string]string{"User-Agent": "engine-api-cli-1.0"}
	cli, err := client.NewClient("unix:///var/run/docker.sock", "", nil, defaultHeaders)
	checkErr(err)

	container, err := cli.ContainerInspect(context.Background(), os.Args[1])
	checkErr(err)

	fmt.Printf("%s", container.NetworkSettings.DefaultNetworkSettings.IPAddress)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
