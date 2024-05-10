// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// AttendanceTrackerAttendanceRecord is an auto generated low-level Go binding around an user-defined struct.
type AttendanceTrackerAttendanceRecord struct {
	EmployeeId  *big.Int
	CheckInTime *big.Int
	Details     string
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"attendanceRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"employeeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"checkInTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"employeeAttendanceLog\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"employeeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"checkInTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"name\":\"recordAttendance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"employeeId\",\"type\":\"uint256\"}],\"name\":\"getAttendanceByEmployee\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"employeeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"checkInTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"internalType\":\"structAttendanceTracker.AttendanceRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"employeeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"checkInTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"name\":\"updateAttendance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"employeeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"getAttendanceByDateRange\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"employeeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"checkInTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"internalType\":\"structAttendanceTracker.AttendanceRecord[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b503360025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506112b08061005c5f395ff3fe608060405234801561000f575f80fd5b506004361061007b575f3560e01c80635763c184116100595780635763c184146100fb5780638da5cb5b1461012d5780639f43f46f1461014b578063d72e70e91461017b5761007b565b806308c220c21461007f5780632330f9ac1461009b5780632e74c140146100cb575b5f80fd5b61009960048036038101906100949190610a13565b610197565b005b6100b560048036038101906100b09190610a7f565b6102bb565b6040516100c29190610acc565b60405180910390f35b6100e560048036038101906100e09190610ae5565b6102e6565b6040516100f29190610bcc565b60405180910390f35b61011560048036038101906101109190610ae5565b6103b3565b60405161012493929190610c34565b60405180910390f35b61013561045e565b6040516101429190610caf565b60405180910390f35b61016560048036038101906101609190610cc8565b610483565b6040516101729190610e20565b60405180910390f35b61019560048036038101906101909190610a13565b610735565b005b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610226576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161021d90610e8a565b60405180910390fd5b5f6040518060600160405280858152602001848152602001838152509050805f808681526020019081526020015f205f820151815f015560208201518160010155604082015181600201908161027c91906110a2565b5090505060015f8581526020019081526020015f2083908060018154018082558091505060019003905f5260205f20015f909190919091505550505050565b6001602052815f5260405f2081815481106102d4575f80fd5b905f5260205f20015f91509150505481565b6102ee610874565b5f808381526020019081526020015f206040518060600160405290815f82015481526020016001820154815260200160028201805461032c90610ed5565b80601f016020809104026020016040519081016040528092919081815260200182805461035890610ed5565b80156103a35780601f1061037a576101008083540402835291602001916103a3565b820191905f5260205f20905b81548152906001019060200180831161038657829003601f168201915b5050505050815250509050919050565b5f602052805f5260405f205f91509050805f0154908060010154908060020180546103dd90610ed5565b80601f016020809104026020016040519081016040528092919081815260200182805461040990610ed5565b80156104545780601f1061042b57610100808354040283529160200191610454565b820191905f5260205f20905b81548152906001019060200180831161043757829003601f168201915b5050505050905083565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60605f60015f8681526020019081526020015f208054806020026020016040519081016040528092919081815260200182805480156104df57602002820191905f5260205f20905b8154815260200190600101908083116104cb575b505050505090505f815167ffffffffffffffff811115610502576105016108ef565b5b60405190808252806020026020018201604052801561053b57816020015b610528610874565b8152602001906001900390816105205790505b5090505f805b835181101561067e575f805f8684815181106105605761055f611171565b5b602002602001015181526020019081526020015f206040518060600160405290815f8201548152602001600182015481526020016002820180546105a390610ed5565b80601f01602080910402602001604051908101604052809291908181526020018280546105cf90610ed5565b801561061a5780601f106105f15761010080835404028352916020019161061a565b820191905f5260205f20905b8154815290600101906020018083116105fd57829003601f168201915b50505050508152505090508781602001511015801561063d575086816020015111155b15610670578084848151811061065657610655611171565b5b6020026020010181905250828061066c906111cb565b9350505b508080600101915050610541565b505f8167ffffffffffffffff81111561069a576106996108ef565b5b6040519080825280602002602001820160405280156106d357816020015b6106c0610874565b8152602001906001900390816106b85790505b5090505f5b82811015610726578381815181106106f3576106f2611171565b5b602002602001015182828151811061070e5761070d611171565b5b602002602001018190525080806001019150506106d8565b50809450505050509392505050565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107bb90610e8a565b60405180910390fd5b5f805f8581526020019081526020015f206001015403610819576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108109061125c565b60405180910390fd5b6040518060600160405280848152602001838152602001828152505f808581526020019081526020015f205f820151815f015560208201518160010155604082015181600201908161086b91906110a2565b50905050505050565b60405180606001604052805f81526020015f8152602001606081525090565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b6108b6816108a4565b81146108c0575f80fd5b50565b5f813590506108d1816108ad565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610925826108df565b810181811067ffffffffffffffff82111715610944576109436108ef565b5b80604052505050565b5f610956610893565b9050610962828261091c565b919050565b5f67ffffffffffffffff821115610981576109806108ef565b5b61098a826108df565b9050602081019050919050565b828183375f83830152505050565b5f6109b76109b284610967565b61094d565b9050828152602081018484840111156109d3576109d26108db565b5b6109de848285610997565b509392505050565b5f82601f8301126109fa576109f96108d7565b5b8135610a0a8482602086016109a5565b91505092915050565b5f805f60608486031215610a2a57610a2961089c565b5b5f610a37868287016108c3565b9350506020610a48868287016108c3565b925050604084013567ffffffffffffffff811115610a6957610a686108a0565b5b610a75868287016109e6565b9150509250925092565b5f8060408385031215610a9557610a9461089c565b5b5f610aa2858286016108c3565b9250506020610ab3858286016108c3565b9150509250929050565b610ac6816108a4565b82525050565b5f602082019050610adf5f830184610abd565b92915050565b5f60208284031215610afa57610af961089c565b5b5f610b07848285016108c3565b91505092915050565b610b19816108a4565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f610b5182610b1f565b610b5b8185610b29565b9350610b6b818560208601610b39565b610b74816108df565b840191505092915050565b5f606083015f830151610b945f860182610b10565b506020830151610ba76020860182610b10565b5060408301518482036040860152610bbf8282610b47565b9150508091505092915050565b5f6020820190508181035f830152610be48184610b7f565b905092915050565b5f82825260208201905092915050565b5f610c0682610b1f565b610c108185610bec565b9350610c20818560208601610b39565b610c29816108df565b840191505092915050565b5f606082019050610c475f830186610abd565b610c546020830185610abd565b8181036040830152610c668184610bfc565b9050949350505050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610c9982610c70565b9050919050565b610ca981610c8f565b82525050565b5f602082019050610cc25f830184610ca0565b92915050565b5f805f60608486031215610cdf57610cde61089c565b5b5f610cec868287016108c3565b9350506020610cfd868287016108c3565b9250506040610d0e868287016108c3565b9150509250925092565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f606083015f830151610d565f860182610b10565b506020830151610d696020860182610b10565b5060408301518482036040860152610d818282610b47565b9150508091505092915050565b5f610d998383610d41565b905092915050565b5f602082019050919050565b5f610db782610d18565b610dc18185610d22565b935083602082028501610dd385610d32565b805f5b85811015610e0e5784840389528151610def8582610d8e565b9450610dfa83610da1565b925060208a01995050600181019050610dd6565b50829750879550505050505092915050565b5f6020820190508181035f830152610e388184610dad565b905092915050565b7f556e617574686f72697a656400000000000000000000000000000000000000005f82015250565b5f610e74600c83610bec565b9150610e7f82610e40565b602082019050919050565b5f6020820190508181035f830152610ea181610e68565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610eec57607f821691505b602082108103610eff57610efe610ea8565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610f617fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610f26565b610f6b8683610f26565b95508019841693508086168417925050509392505050565b5f819050919050565b5f610fa6610fa1610f9c846108a4565b610f83565b6108a4565b9050919050565b5f819050919050565b610fbf83610f8c565b610fd3610fcb82610fad565b848454610f32565b825550505050565b5f90565b610fe7610fdb565b610ff2818484610fb6565b505050565b5b818110156110155761100a5f82610fdf565b600181019050610ff8565b5050565b601f82111561105a5761102b81610f05565b61103484610f17565b81016020851015611043578190505b61105761104f85610f17565b830182610ff7565b50505b505050565b5f82821c905092915050565b5f61107a5f198460080261105f565b1980831691505092915050565b5f611092838361106b565b9150826002028217905092915050565b6110ab82610b1f565b67ffffffffffffffff8111156110c4576110c36108ef565b5b6110ce8254610ed5565b6110d9828285611019565b5f60209050601f83116001811461110a575f84156110f8578287015190505b6111028582611087565b865550611169565b601f19841661111886610f05565b5f5b8281101561113f5784890151825560018201915060208501945060208101905061111a565b8683101561115c5784890151611158601f89168261106b565b8355505b6001600288020188555050505b505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6111d5826108a4565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036112075761120661119e565b5b600182019050919050565b7f5265636f726420646f6573206e6f742065786973742e000000000000000000005f82015250565b5f611246601683610bec565b915061125182611212565b602082019050919050565b5f6020820190508181035f8301526112738161123a565b905091905056fea264697066735822122000259dcf51b0124d3c4d1d75aebc85a5201f0de2dfa8e507d0338dfd9cde48f664736f6c63430008190033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// AttendanceRecords is a free data retrieval call binding the contract method 0x5763c184.
//
// Solidity: function attendanceRecords(uint256 ) view returns(uint256 employeeId, uint256 checkInTime, string details)
func (_Contract *ContractCaller) AttendanceRecords(opts *bind.CallOpts, arg0 *big.Int) (struct {
	EmployeeId  *big.Int
	CheckInTime *big.Int
	Details     string
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "attendanceRecords", arg0)

	outstruct := new(struct {
		EmployeeId  *big.Int
		CheckInTime *big.Int
		Details     string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EmployeeId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CheckInTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Details = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// AttendanceRecords is a free data retrieval call binding the contract method 0x5763c184.
//
// Solidity: function attendanceRecords(uint256 ) view returns(uint256 employeeId, uint256 checkInTime, string details)
func (_Contract *ContractSession) AttendanceRecords(arg0 *big.Int) (struct {
	EmployeeId  *big.Int
	CheckInTime *big.Int
	Details     string
}, error) {
	return _Contract.Contract.AttendanceRecords(&_Contract.CallOpts, arg0)
}

// AttendanceRecords is a free data retrieval call binding the contract method 0x5763c184.
//
// Solidity: function attendanceRecords(uint256 ) view returns(uint256 employeeId, uint256 checkInTime, string details)
func (_Contract *ContractCallerSession) AttendanceRecords(arg0 *big.Int) (struct {
	EmployeeId  *big.Int
	CheckInTime *big.Int
	Details     string
}, error) {
	return _Contract.Contract.AttendanceRecords(&_Contract.CallOpts, arg0)
}

// EmployeeAttendanceLog is a free data retrieval call binding the contract method 0x2330f9ac.
//
// Solidity: function employeeAttendanceLog(uint256 , uint256 ) view returns(uint256)
func (_Contract *ContractCaller) EmployeeAttendanceLog(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "employeeAttendanceLog", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EmployeeAttendanceLog is a free data retrieval call binding the contract method 0x2330f9ac.
//
// Solidity: function employeeAttendanceLog(uint256 , uint256 ) view returns(uint256)
func (_Contract *ContractSession) EmployeeAttendanceLog(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Contract.Contract.EmployeeAttendanceLog(&_Contract.CallOpts, arg0, arg1)
}

// EmployeeAttendanceLog is a free data retrieval call binding the contract method 0x2330f9ac.
//
// Solidity: function employeeAttendanceLog(uint256 , uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) EmployeeAttendanceLog(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Contract.Contract.EmployeeAttendanceLog(&_Contract.CallOpts, arg0, arg1)
}

// GetAttendanceByDateRange is a free data retrieval call binding the contract method 0x9f43f46f.
//
// Solidity: function getAttendanceByDateRange(uint256 employeeId, uint256 startTime, uint256 endTime) view returns((uint256,uint256,string)[])
func (_Contract *ContractCaller) GetAttendanceByDateRange(opts *bind.CallOpts, employeeId *big.Int, startTime *big.Int, endTime *big.Int) ([]AttendanceTrackerAttendanceRecord, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getAttendanceByDateRange", employeeId, startTime, endTime)

	if err != nil {
		return *new([]AttendanceTrackerAttendanceRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]AttendanceTrackerAttendanceRecord)).(*[]AttendanceTrackerAttendanceRecord)

	return out0, err

}

// GetAttendanceByDateRange is a free data retrieval call binding the contract method 0x9f43f46f.
//
// Solidity: function getAttendanceByDateRange(uint256 employeeId, uint256 startTime, uint256 endTime) view returns((uint256,uint256,string)[])
func (_Contract *ContractSession) GetAttendanceByDateRange(employeeId *big.Int, startTime *big.Int, endTime *big.Int) ([]AttendanceTrackerAttendanceRecord, error) {
	return _Contract.Contract.GetAttendanceByDateRange(&_Contract.CallOpts, employeeId, startTime, endTime)
}

// GetAttendanceByDateRange is a free data retrieval call binding the contract method 0x9f43f46f.
//
// Solidity: function getAttendanceByDateRange(uint256 employeeId, uint256 startTime, uint256 endTime) view returns((uint256,uint256,string)[])
func (_Contract *ContractCallerSession) GetAttendanceByDateRange(employeeId *big.Int, startTime *big.Int, endTime *big.Int) ([]AttendanceTrackerAttendanceRecord, error) {
	return _Contract.Contract.GetAttendanceByDateRange(&_Contract.CallOpts, employeeId, startTime, endTime)
}

// GetAttendanceByEmployee is a free data retrieval call binding the contract method 0x2e74c140.
//
// Solidity: function getAttendanceByEmployee(uint256 employeeId) view returns((uint256,uint256,string))
func (_Contract *ContractCaller) GetAttendanceByEmployee(opts *bind.CallOpts, employeeId *big.Int) (AttendanceTrackerAttendanceRecord, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getAttendanceByEmployee", employeeId)

	if err != nil {
		return *new(AttendanceTrackerAttendanceRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(AttendanceTrackerAttendanceRecord)).(*AttendanceTrackerAttendanceRecord)

	return out0, err

}

// GetAttendanceByEmployee is a free data retrieval call binding the contract method 0x2e74c140.
//
// Solidity: function getAttendanceByEmployee(uint256 employeeId) view returns((uint256,uint256,string))
func (_Contract *ContractSession) GetAttendanceByEmployee(employeeId *big.Int) (AttendanceTrackerAttendanceRecord, error) {
	return _Contract.Contract.GetAttendanceByEmployee(&_Contract.CallOpts, employeeId)
}

// GetAttendanceByEmployee is a free data retrieval call binding the contract method 0x2e74c140.
//
// Solidity: function getAttendanceByEmployee(uint256 employeeId) view returns((uint256,uint256,string))
func (_Contract *ContractCallerSession) GetAttendanceByEmployee(employeeId *big.Int) (AttendanceTrackerAttendanceRecord, error) {
	return _Contract.Contract.GetAttendanceByEmployee(&_Contract.CallOpts, employeeId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// RecordAttendance is a paid mutator transaction binding the contract method 0x08c220c2.
//
// Solidity: function recordAttendance(uint256 employeeId, uint256 checkInTime, string details) returns()
func (_Contract *ContractTransactor) RecordAttendance(opts *bind.TransactOpts, employeeId *big.Int, checkInTime *big.Int, details string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "recordAttendance", employeeId, checkInTime, details)
}

// RecordAttendance is a paid mutator transaction binding the contract method 0x08c220c2.
//
// Solidity: function recordAttendance(uint256 employeeId, uint256 checkInTime, string details) returns()
func (_Contract *ContractSession) RecordAttendance(employeeId *big.Int, checkInTime *big.Int, details string) (*types.Transaction, error) {
	return _Contract.Contract.RecordAttendance(&_Contract.TransactOpts, employeeId, checkInTime, details)
}

// RecordAttendance is a paid mutator transaction binding the contract method 0x08c220c2.
//
// Solidity: function recordAttendance(uint256 employeeId, uint256 checkInTime, string details) returns()
func (_Contract *ContractTransactorSession) RecordAttendance(employeeId *big.Int, checkInTime *big.Int, details string) (*types.Transaction, error) {
	return _Contract.Contract.RecordAttendance(&_Contract.TransactOpts, employeeId, checkInTime, details)
}

// UpdateAttendance is a paid mutator transaction binding the contract method 0xd72e70e9.
//
// Solidity: function updateAttendance(uint256 employeeId, uint256 checkInTime, string details) returns()
func (_Contract *ContractTransactor) UpdateAttendance(opts *bind.TransactOpts, employeeId *big.Int, checkInTime *big.Int, details string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateAttendance", employeeId, checkInTime, details)
}

// UpdateAttendance is a paid mutator transaction binding the contract method 0xd72e70e9.
//
// Solidity: function updateAttendance(uint256 employeeId, uint256 checkInTime, string details) returns()
func (_Contract *ContractSession) UpdateAttendance(employeeId *big.Int, checkInTime *big.Int, details string) (*types.Transaction, error) {
	return _Contract.Contract.UpdateAttendance(&_Contract.TransactOpts, employeeId, checkInTime, details)
}

// UpdateAttendance is a paid mutator transaction binding the contract method 0xd72e70e9.
//
// Solidity: function updateAttendance(uint256 employeeId, uint256 checkInTime, string details) returns()
func (_Contract *ContractTransactorSession) UpdateAttendance(employeeId *big.Int, checkInTime *big.Int, details string) (*types.Transaction, error) {
	return _Contract.Contract.UpdateAttendance(&_Contract.TransactOpts, employeeId, checkInTime, details)
}
