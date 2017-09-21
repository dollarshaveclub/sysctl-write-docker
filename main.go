package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"strconv"
	"time"
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
		stderrBuf := &bytes.Buffer{}
		command.Stderr = stderrBuf

		if err := command.Run(); err != nil {
			if _, ok := err.(*exec.ExitError); ok {
				log.Fatalf("error setting kernel parameter %s=%s: %s", key, value, stderrBuf.String())
			}
			log.Fatalf("error setting kernel parameter %s=%s: %s", key, value, err)
		}
	}

	// Optionally block forever. This usefule when this tool is
	// ran as a DaemonSet in Kubernetes.
	if ok, _ := strconv.ParseBool(os.Getenv("SYSCTL_BLOCK")); ok {
		<-time.After(time.Duration(math.MaxInt64))
	}
}
