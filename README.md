Orcacoin
====


Orcacoin is an alternative full node bitcoin implementation written in Go (golang). 

This project is currently under active development.

It is extremely stable.

It properly downloads, validates, and serves the block chain using the exact
rules (including consensus bugs) for block acceptance as Bitcoin Core.  

It also properly relays newly mined blocks, maintains a transaction pool, and
relays individual transactions that have not yet made it into a block.  

Orcacoin itself does **NOT** include a wallet implementation and we have a separate repository where we have done work on a [OrcaWallet](https://github.com/sbu-416-24sp/orcanet-go/tree/main/wallet)

## Requirements

[Go](http://golang.org) 1.17 or newer.

## Installation

#### Linux/BSD/MacOSX/POSIX - Build from Source

- Install Go according to the installation instructions here:
  http://golang.org/doc/install

- Ensure Go was installed properly and is a supported version:

```bash
$ go version
$ go env GOROOT GOPATH
```

NOTE: The `GOROOT` and `GOPATH` above must not be the same path.  It is
recommended that `GOPATH` is set to a directory in your home directory such as
`~/goprojects` to avoid write permission issues.  It is also recommended to add
`$GOPATH/bin` to your `PATH` at this point.

Since orcacoin was implemented on top of btcd, our dependencies will be the same as btcd.
- Run the following commands to obtain btcd, all dependencies, and install it:

```bash
$ cd $GOPATH/src/github.com/btcsuite/btcd
$ go install -v . ./cmd/...
```

- btcd (and utilities) will now be installed in ```$GOPATH/bin```.  If you did
  not already add the bin directory to your system path during Go installation,
  we recommend you do so now.

## Updating

#### Linux/BSD/MacOSX/POSIX - Build from Source

- Run the following commands to update btcd, all dependencies, and install it:

```bash
$ cd $GOPATH/src/github.com/btcsuite/btcd
$ git pull
$ go install -v . ./cmd/...
```

## Setting Up User Info

### sample-btcd.conf
Include this information in the file:
```bash
rpcuser=set-your-username
rpcpass=set-your-password

miningaddr=set-your-wallet-addres-1
miningaddr=set-your-wallet-address-2
```
(You can have multiple wallet addresses associated with a single wallet)

### cmd/btcctl/sample-btcctl.conf
Include this information in the file:
```bash
rpcuser=same-username-as-above
rpcpass=same-password-as-above
```

## Running The Node

To run an OrcaCoin node, there are several configurations. The two most important ones are the default config (you have to manually mine blocks), and the config that lets your node mine automatically.

***Running the default config***
```bash
$ go install
$ go build
$ ./btcd
```

***Running the mining config***
```bash
$ go install
$ go build
$ ./btcd --generate
```

## api.go Commands

***generateNewWalletCommand***: Generates a wallet address for the specified wallet.

- *Argument* (String) → Name of the wallet account, usually “default”
- Generates wallet address.
- **Returns** wallet address.
- NOTE: Running this command will generate a new wallet address for the same account (it should not be used for retrieving wallet address)

***getWalletAddressCommand***: Retrieves the wallet address of the specified wallet.

- *Argument* (String)→ Name of the wallet account, usually “default”
- **Returns** the wallet address

***getAllAccountsCommand***: Lists all the wallet accounts' names.

- No arguments needed
- **Returns** list of all the accounts which can then be used for retrieving wallet address

***getBalanceCommand***: Retrieves the balance of the current running wallet.

- No arguments needed
- **Returns** the balance 

***walletPassphraseCommand***: Unlocks the wallet for the specified amount of time.

- *Argument 1* (String)→ Wallet passphrase which was set when wallet was created 
- *Argument 2* (String)→ Amount of time to keep the wallet unlocked in millisecond 
- Unlocks the wallet for specified amount of time
- NOTE : Current wallet needs to be unlocked before proceeding with money transfer

***sendToAddressCommand***: Transfer Orca Coin from current wallet to the specified wallet address.

- Reminder to unlock the wallet first! (**run the walletPassphraseCommand**)
- *Argument 1* (String)→ receiving wallet address
- *Argument 2* (String) → Amount of money
- **Returns** the transaction hash


***generateCommand***: Mines the specified number of blocks and sends rewards to the running wallet.

- *Argument* (String)→ Number of blocks to generate
- **Returns** the array of block hashes

***getAllTransactionsCommand***: **Returns** the list of transactions.

- No argument
- **Returns** a list of all the transactions (including confirmed and unconfirmed transactions)
- NOTE: It can take few hours for the transaction to be confirmed and appear on blockchain