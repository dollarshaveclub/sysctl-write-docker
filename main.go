package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var (
	sysctlCmd = "/sbin/sysctl"
)

type sysctlMapping map[string]string

func main() {
	if cmd := os.Getenv("SYSCTL_CMD"); cmd != "" {
		sysctlCmd = cmd
	}

	unparsedMapping := os.Getenv("SYSCTL")
	if unparsedMapping == "" {
		log.Fatalf("`SYSCTL` variable not set.")
	}

	var mapping sysctlMapping
	if err := json.Unmarshal([]byte(unparsedMapping), &mapping); err != nil {
		log.Fatalf("error parsing `SYSCTL` variable as JSON: %s", err)
	}

	for key, value := range mapping {
		log.Printf("setting kernel parameter with %s: %s=%s", sysctlCmd, key, value)
		command := exec.Command(
			sysctlCmd,
			"-w",
			fmt.Sprintf("%s=%s", key, value),
		)
		if err := command.Run(); err != nil {
			log.Fatalf("error setting kernel parameter %s=%s: %s", key, value, err)
		}
	}
}
