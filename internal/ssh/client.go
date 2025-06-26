package ssh

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// PrepareConnectionCommand creates an *exec.Cmd to run the ssh command for a given host.
// It sets up stdin, stdout, and stderr to take over the current terminal.
// By building the command and arguments separately, we prevent shell injection vulnerabilities.
func PrepareConnectionCommand(hostName string) *exec.Cmd {
	// #nosec G204
	// This is the core functionality of the application. We must execute the ssh
	// command. We are passing the hostname as a distinct argument to exec.Command,
	// which prevents command injection vulnerabilities. The hostName is sourced
	// from the user's own SSH config file, which is a trusted source in this context.
	cmd := exec.Command("ssh", hostName)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}

// ConnectToHost executes the SSH command for the given host and handles any errors
func ConnectToHost(hostName string) error {
	// Special handling for GitHub - just test the connection
	if strings.Contains(hostName, "github.com") {
		testCmd := exec.Command("ssh", "-T", hostName)
		testCmd.Stdout = os.Stdout
		testCmd.Stderr = os.Stderr
		err := testCmd.Run()
		// GitHub always returns exit status 1 even on successful auth
		// We check if the error is just the exit status 1
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 1 {
			return nil
		}
		return err
	}

	cmd := PrepareConnectionCommand(hostName)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to connect to %s: %w", hostName, err)
	}
	return nil
}
