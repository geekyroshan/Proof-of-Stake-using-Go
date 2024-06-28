# Proof of Stake Algorithm using Go

This repository contains a Proof of Stake (PoS) algorithm implemented in Go. PoS is a consensus algorithm used in blockchain networks to achieve distributed consensus.
## Files

- `generateNewBlock.go` - Handles the creation of new blocks.
- `go.mod` - Module file for managing dependencies.
- `hashing.go` - Implements hashing functions.
- `main.go` - Entry point for the application.
- `newNode.go` - Manages new node creation and registration.
- `selectwinner.go` - Selects the winner node for block creation.
- `structs.go` - Defines the data structures used in the application.
- `Validateblockcandidate.go` - Validates potential block candidates.
- `validateBlockchain.go` - Validates the integrity of the blockchain.

## Getting Started

### Prerequisites

- Go (version 1.16 or higher)

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/your-username/your-repo-name.git
    cd your-repo-name
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

### Usage

To run the application:
```sh
go run *.go
