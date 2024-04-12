package main

import (
	"fmt"
	"os/exec"
)

func btcctlCommand(args []string) {
	argsLen := len(args)
	if argsLen == 0 {
		fmt.Println("No commands were provided")
		return
	}

	btcctlConfPath := "/Users/draku/go/src/orcacointest/cmd/btcctl/sample-btcctl.conf"
	argIdx := 0
	for argIdx < argsLen {
		arg := args[argIdx]
		switch arg {
		case "getbalance":
			cmd := exec.Command("btcctl", "--configfile="+btcctlConfPath, "getbalance", "--notls")
			output, err := cmd.Output()
			if err != nil {
				fmt.Println("Error when running 'getbalance':", err)
				return
			}

			fmt.Println(string(output))
		case "sendtoaddress":
			if len(args[argIdx:]) < 2 {
				fmt.Print("Not enough arguments for the command 'sendtoaddress'")
				return
			}

			cmd := exec.Command("btcctl", "--configfile="+btcctlConfPath, "sendtoaddress",
				args[argIdx+1], args[argIdx+2], "--notls")
			output, err := cmd.Output()
			if err != nil {
				fmt.Println("Error running 'sendtoaddress'", err)
				return
			}

			argIdx += 2
			fmt.Println(string(output))
		case "generate":
			if len(args[argIdx:]) < 1 {
				fmt.Print("Not enough arguments for the command 'generate'")
				return
			}

			cmd := exec.Command("btcctl", "--configfile="+btcctlConfPath, "generate", args[argIdx+1], "--notls")
			output, err := cmd.Output()
			if err != nil {
				fmt.Println("Error running 'generate'", err)
				return
			}

			argIdx += 1
			fmt.Println(string(output))
		case "walletpassphrase":
			if len(args[argIdx:]) < 2 {
				fmt.Print("Not enough arguments for the command 'walletpassphrase'")
				return
			}

			cmd := exec.Command("btcctl", "--configfile="+btcctlConfPath, "walletpassphrase",
				args[argIdx+1], args[argIdx+2], "--notls")
			output, err := cmd.Output()
			if err != nil {
				fmt.Println("Error running 'walletpassphrase'", err)
				return
			}

			argIdx += 2
			fmt.Println(string(output))
		}

		argIdx += 1
	}
}
