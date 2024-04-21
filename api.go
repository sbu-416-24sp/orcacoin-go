package main

import (
	"fmt"
	"os/exec"
	"sync"
	"time"
)

// Commands
const (
	getBalanceCommand       = "getbalance"
	sendToAddressCommand    = "sendtoaddress"
	generateCommand         = "generate"
	walletPassphraseCommand = "walletpassphrase"
)

func main() {
	// Create WaitGroup to wait for both commands to finish
	var wg sync.WaitGroup

	wg.Add(1)
	go runBtcd()

	wg.Add(1)
	go runBtcWallet()

	// calls the getbalance from btcctl after waiting for 10 seconds to ensure wallet connects
	time.Sleep(10 * time.Second)
	args := []string{getBalanceCommand}
	apiCall(args)

	// Wait for both commands to finish
	wg.Wait()

	fmt.Println("Both btcd and btcwallet have finished running.")
}

func apiCall(args []string) {
	argsLen := len(args)
	if argsLen == 0 {
		fmt.Println("No commands were provided")
		return
	}

	for argIdx := 0; argIdx < argsLen; argIdx++ {
		arg := args[argIdx]
		switch arg {
		case getBalanceCommand:
			err := executeCommand(getBalanceCommand)
			if err != nil {
				return
			}
		case sendToAddressCommand, generateCommand, walletPassphraseCommand:
			if len(args[argIdx:]) < 2 {
				fmt.Printf("Not enough arguments for the command '%s'\n", arg)
				return
			}

			err := executeCommand(arg, args[argIdx+1:]...)
			if err != nil {
				return
			}
			argIdx++
		}
	}
}

// executeCommand executes the given command with the provided arguments
func executeCommand(command string, args ...string) error {
	cmd := exec.Command("btcctl", command)
	cmd.Args = append(cmd.Args, args...)
	// cmd.Args = append(cmd.Args, "--notls")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("Error running '%s': %w", command, err)
	}
	fmt.Println(string(output))
	return nil
}

func runBtcd() error {
	cmd := exec.Command("btcd")

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("Error starting btcd: %v", err)
	}

	return nil
}

func runBtcWallet() error {
	cmd := exec.Command("btcwallet")

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("Error starting btcwallet: %v", err)
	}

	return nil
}
