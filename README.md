# Attendance Tracker Project Documentation

## Overview of Solution Architecture

### System Architecture

The Attendance Tracker project consists of two main components: a smart contract deployed on a blockchain network (e.g., Ethereum) and a backend service written in Go, using the `ethclient` library to interact with the blockchain.

#### Interaction Between Smart Contract and Backend

- **Smart Contract**: Responsible for storing and managing attendance data.
- **Backend Service**: Communicates with the smart contract through API calls to record and retrieve attendance data.

### Architecture Diagram

```plaintext
[Client] --HTTP--> [Backend Service] --Blockchain--> [Smart Contract]
```

## Deploying the Smart Contract

### Requirements

- Node.js and npm
- Truffle Suite
- Ganache (for development environment)

### Deployment Steps

1. **Install Truffle**:

```bash
npm install -g truffle
```

2. **Initialize and Configure Project**:

   - Initialize a new project with Truffle:

   ```bash
   truffle init
   ```

   - Configure network connections in `truffle-config.js`.

3. **Deploy Contract**:
   - Run the following command to deploy the contract to the selected network:
   ```bash
   truffle migrate --network development
   ```

## API Documentation

### Endpoint `/api/attendance`

- **POST**: Record attendance

  - **Request**:

    ```json
    {
      "employeeId": 1,
      "checkInTime": 1622554800,
      "details": "Checked in at main gate"
    }
    ```

  - **Response**:

    ```json
    {
      "message": "Attendance recorded",
      "txHash": "0x..."
    }
    ```

### Endpoint `/api/attendance/:employeeId`

- **GET**: Query attendance by employee ID
  - **Response**:
    ```json
    {
      "employeeId": 1,
      "records": [
        {
          "date": "2021-06-01",
          "details": "Checked in at main gate"
        }
      ]
    }
    ```

## Data Schema

### Attendance Schema

- **EmployeeID**: uint
- **CheckInTime**: uint (timestamp)
- **Details**: string

## Developer Guidelines

### Setting Up Development Environment

- Install dependencies:

```bash
npm install
go get -u github.com/ethereum/go-ethereum
```

### Contributing

- Fork and clone the repo.
- Create a new branch for your feature or fix.
- After completion, submit a Pull Request to the `main` branch.

## Hosting the Documentation

This documentation is hosted on GitHub Pages, accessible via [this link](https://yourgithubusername.github.io/attendance-tracker/).
