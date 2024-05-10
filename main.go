package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	contract "github.com/gpnull/be-earning-interview/contracts"
)

type EthClientInterface interface {
	bind.ContractBackend
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
}

// Giả sử đây là khóa riêng, chỉ sử dụng cho mục đích phát triển
var privateKeyHex = "43f1917060aaa466eec53ba542193cf9c006b719c7d1331985462ccd7230d1d9"

// Giả sử ChainID cho mạng Ethereum Ropsten
var chainID = int64(3) // Ropsten Testnet ID

var privateKey *ecdsa.PrivateKey

func init() {
	// Chuyển đổi khóa riêng từ hex string sang *ecdsa.PrivateKey
	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		log.Fatalf("Invalid private key hex: %v", err)
	}

	privateKey, err = crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		log.Fatalf("Failed to create private key: %v", err)
	}
}

func main() {
	router := gin.Default()

	// Khởi tạo kết nối với Ethereum node
	client, err := ethclient.Dial("http://localhost:9545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	defer client.Close()

	// Đăng ký các route
	router.POST("/api/attendance", func(c *gin.Context) {
		recordAttendance(c, client)
	})
	router.GET("/api/attendance/:employeeId", func(c *gin.Context) {
		getAttendanceByEmployee(c, client)
	})
	router.PUT("/api/attendance/:employeeId", func(c *gin.Context) {
		updateAttendance(c, client)
	})

	router.Run(":8080")
}

func recordAttendance(c *gin.Context, client EthClientInterface) {
	// Giả sử địa chỉ của smart contract và ABI đã biết
	contractAddress := common.HexToAddress("0x4D665FE987608077b4273C62D80002e4EEFC2130")
	contract, err := contract.NewContract(contractAddress, client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Contract instance error"})
		return
	}

	// Giả sử chúng ta có các tham số cần thiết từ body request
	var json struct {
		EmployeeID  uint
		CheckInTime uint64
		Details     string
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Gọi hàm recordAttendance trong smart contract
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	if err != nil {
		// Xử lý lỗi ở đây, ví dụ:
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	// Gọi hàm RecordAttendance với auth đã được tạo
	tx, err := contract.RecordAttendance(
		auth,
		big.NewInt(int64(json.EmployeeID)),
		big.NewInt(int64(json.CheckInTime)),
		json.Details,
	)
	if err != nil {
		// Xử lý lỗi ở đây, ví dụ:
		log.Fatalf("Failed to record attendance: %v", err)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record attendance"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Attendance recorded", "txHash": tx.Hash().Hex()})
}

func getAttendanceByEmployee(c *gin.Context, client EthClientInterface) {
	employeeId := c.Param("employeeId")
	employeeIdInt, err := strconv.ParseInt(employeeId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	// Giả sử địa chỉ của smart contract và ABI đã biết
	contractAddress := common.HexToAddress("0x4D665FE987608077b4273C62D80002e4EEFC2130")
	contractInstance, err := contract.NewContract(contractAddress, client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load contract"})
		return
	}

	// Gọi hàm GetAttendanceByEmployee của smart contract
	attendanceRecord, err := contractInstance.GetAttendanceByEmployee(nil, big.NewInt(employeeIdInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve attendance record"})
		return
	}

	// Chuyển đổi kết quả trả về thành một định dạng phù hợp để gửi lại cho client
	// Giả sử attendanceRecord là một struct chứa thông tin điểm danh
	recordMap := map[string]interface{}{
		"date":    attendanceRecord.CheckInTime,
		"details": attendanceRecord.Details,
	}

	c.JSON(http.StatusOK, gin.H{"employeeId": employeeId, "record": recordMap})
}

func updateAttendance(c *gin.Context, client EthClientInterface) {
	employeeId := c.Param("employeeId")
	employeeIdInt, err := strconv.ParseInt(employeeId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	// Giả sử địa chỉ của smart contract và ABI đã biết
	contractAddress := common.HexToAddress("0x4D665FE987608077b4273C62D80002e4EEFC2130")
	contractInstance, err := contract.NewContract(contractAddress, client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load contract"})
		return
	}

	// Giả sử chúng ta có các tham số cần thiết từ body request
	var json struct {
		CheckInTime uint64
		Details     string
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Khởi tạo auth cho giao dịch
	privateKey, err := crypto.HexToECDSA("0xd9274f0e0965a76380a8d5267fe2407bd720f771b334085280819acdc3192bd6")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse private key"})
		return
	}
	chainID := big.NewInt(3) // Ví dụ: Ropsten Testnet ID
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create authorized transactor"})
		return
	}

	// Gọi hàm UpdateAttendance của smart contract
	tx, err := contractInstance.UpdateAttendance(
		auth,
		big.NewInt(employeeIdInt),
		big.NewInt(int64(json.CheckInTime)),
		json.Details,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update attendance"})
		return
	}

	// Trả về transaction hash để theo dõi
	c.JSON(http.StatusOK, gin.H{"employeeId": employeeId, "message": "Attendance updated", "txHash": tx.Hash().Hex()})
}

func fakePrivateKeyHex() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := fmt.Sprintf("%x", privateKeyBytes)
	fmt.Println("Private Key:", privateKeyHex)
}
