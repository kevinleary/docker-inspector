package main

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func main() {
	// Create a new Docker client.
	// The client will automatically connect to the Docker daemon
	// based on standard environment variables.
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create docker client: %s\n", err)
		os.Exit(1)
	}
	defer cli.Close()

	// Create a context for our API calls.
	ctx := context.Background()

	// List all containers. The ListOptions can be used to filter the results.
	// We are asking for all containers, not just the running ones.
	containers, err := cli.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to list containers: %s\n", err)
		os.Exit(1)
	}

	// If no containers are found, print a message and exit.
	if len(containers) == 0 {
		fmt.Println("No containers found.")
		return
	}

	fmt.Println("--- Docker Container Inspector ---")

	// Iterate over the slice of containers.
	for _, cont := range containers {
		fmt.Println("----------------------------------------")
		// Use ContainerInspect to get detailed information about the container.
		json, err := cli.ContainerInspect(ctx, cont.ID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to inspect container %s: %s\n", cont.ID, err)
			continue
		}

		// Print out some basic information.
		fmt.Printf("ID: %s\n", cont.ID[:12])
		fmt.Printf("Name: %s\n", json.Name)
		fmt.Printf("Image: %s\n", cont.Image)
		fmt.Printf("State: %s\n", cont.State)
		fmt.Printf("Status: %s\n", cont.Status)

		// Print port mappings.
		if len(cont.Ports) > 0 {
			fmt.Println("Ports:")
			for _, port := range cont.Ports {
				// The port mapping can be complex, so we print it in a readable format.
				if port.PublicPort > 0 {
					fmt.Printf("  - %d:%d/%s -> %s:%d\n", port.PublicPort, port.PrivatePort, port.Type, port.IP, port.PrivatePort)
				} else {
					fmt.Printf("  - %d/%s\n", port.PrivatePort, port.Type)
				}
			}
		} else {
			fmt.Println("Ports: No ports exposed.")
		}
	}
	fmt.Println("----------------------------------------")
}
