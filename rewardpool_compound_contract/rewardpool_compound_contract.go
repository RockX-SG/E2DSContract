// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rewardpool_compound_contract

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
)

// RewardpoolCompoundContractMetaData contains all meta data concerning the RewardpoolCompoundContract contract.
var RewardpoolCompoundContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"CONTROLLER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MANAGER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ORACLE_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimRewards\",\"inputs\":[{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimRewardsFor\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAccountedBalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingManagerRevenue\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingReward\",\"inputs\":[{\"name\":\"claimaddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalShare\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"managerFeeShare\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setManagerFeeShare\",\"inputs\":[{\"name\":\"milli\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updatePool\",\"inputs\":[{\"name\":\"claimaddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"deltas\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userInfo\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"accSharePoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"rewardBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawManagerRevenue\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Claimed\",\"inputs\":[{\"name\":\"beneficiary\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ManagerFeeSet\",\"inputs\":[{\"name\":\"milli\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ManagerFeeWithdrawed\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolUpdate\",\"inputs\":[{\"name\":\"claimaddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// RewardpoolCompoundContractABI is the input ABI used to generate the binding from.
// Deprecated: Use RewardpoolCompoundContractMetaData.ABI instead.
var RewardpoolCompoundContractABI = RewardpoolCompoundContractMetaData.ABI

// RewardpoolCompoundContract is an auto generated Go binding around an Ethereum contract.
type RewardpoolCompoundContract struct {
	RewardpoolCompoundContractCaller     // Read-only binding to the contract
	RewardpoolCompoundContractTransactor // Write-only binding to the contract
	RewardpoolCompoundContractFilterer   // Log filterer for contract events
}

// RewardpoolCompoundContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type RewardpoolCompoundContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardpoolCompoundContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RewardpoolCompoundContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardpoolCompoundContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewardpoolCompoundContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardpoolCompoundContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewardpoolCompoundContractSession struct {
	Contract     *RewardpoolCompoundContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// RewardpoolCompoundContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewardpoolCompoundContractCallerSession struct {
	Contract *RewardpoolCompoundContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// RewardpoolCompoundContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewardpoolCompoundContractTransactorSession struct {
	Contract     *RewardpoolCompoundContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// RewardpoolCompoundContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type RewardpoolCompoundContractRaw struct {
	Contract *RewardpoolCompoundContract // Generic contract binding to access the raw methods on
}

// RewardpoolCompoundContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewardpoolCompoundContractCallerRaw struct {
	Contract *RewardpoolCompoundContractCaller // Generic read-only contract binding to access the raw methods on
}

// RewardpoolCompoundContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewardpoolCompoundContractTransactorRaw struct {
	Contract *RewardpoolCompoundContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRewardpoolCompoundContract creates a new instance of RewardpoolCompoundContract, bound to a specific deployed contract.
func NewRewardpoolCompoundContract(address common.Address, backend bind.ContractBackend) (*RewardpoolCompoundContract, error) {
	contract, err := bindRewardpoolCompoundContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RewardpoolCompoundContract{RewardpoolCompoundContractCaller: RewardpoolCompoundContractCaller{contract: contract}, RewardpoolCompoundContractTransactor: RewardpoolCompoundContractTransactor{contract: contract}, RewardpoolCompoundContractFilterer: RewardpoolCompoundContractFilterer{contract: contract}}, nil
}

// NewRewardpoolCompoundContractCaller creates a new read-only instance of RewardpoolCompoundContract, bound to a specific deployed contract.
func NewRewardpoolCompoundContractCaller(address common.Address, caller bind.ContractCaller) (*RewardpoolCompoundContractCaller, error) {
	contract, err := bindRewardpoolCompoundContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewardpoolCompoundContractCaller{contract: contract}, nil
}

// NewRewardpoolCompoundContractTransactor creates a new write-only instance of RewardpoolCompoundContract, bound to a specific deployed contract.
func NewRewardpoolCompoundContractTransactor(address common.Address, transactor bind.ContractTransactor) (*RewardpoolCompoundContractTransactor, error) {
	contract, err := bindRewardpoolCompoundContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewardpoolCompoundContractTransactor{contract: contract}, nil
}

// NewRewardpoolCompoundContractFilterer creates a new log filterer instance of RewardpoolCompoundContract, bound to a specific deployed contract.
func NewRewardpoolCompoundContractFilterer(address common.Address, filterer bind.ContractFilterer) (*RewardpoolCompoundContractFilterer, error) {
	contract, err := bindRewardpoolCompoundContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewardpoolCompoundContractFilterer{contract: contract}, nil
}

// bindRewardpoolCompoundContract binds a generic wrapper to an already deployed contract.
func bindRewardpoolCompoundContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RewardpoolCompoundContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardpoolCompoundContract *RewardpoolCompoundContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardpoolCompoundContract.Contract.RewardpoolCompoundContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardpoolCompoundContract *RewardpoolCompoundContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardpoolCompoundContract.Contract.RewardpoolCompoundContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardpoolCompoundContract *RewardpoolCompoundContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardpoolCompoundContract.Contract.RewardpoolCompoundContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardpoolCompoundContract *RewardpoolCompoundContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardpoolCompoundContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardpoolCompoundContract *RewardpoolCompoundContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardpoolCompoundContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardpoolCompoundContract *RewardpoolCompoundContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardpoolCompoundContract.Contract.contract.Transact(opts, method, params...)
}

// CONTROLLERROLE is a free data retrieval call binding the contract method 0x092c5b3b.
//
// Solidity: function CONTROLLER_ROLE() view returns(bytes32)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractCaller) CONTROLLERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RewardpoolCompoundContract.contract.Call(opts, &out, "CONTROLLER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CONTROLLERROLE is a free data retrieval call binding the contract method 0x092c5b3b.
//
// Solidity: function CONTROLLER_ROLE() view returns(bytes32)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractSession) CONTROLLERROLE() ([32]byte, error) {
	return _RewardpoolCompoundContract.Contract.CONTROLLERROLE(&_RewardpoolCompoundContract.CallOpts)
}

// CONTROLLERROLE is a free data retrieval call binding the contract method 0x092c5b3b.
//
// Solidity: function CONTROLLER_ROLE() view returns(bytes32)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractCallerSession) CONTROLLERROLE() ([32]byte, error) {
	return _RewardpoolCompoundContract.Contract.CONTROLLERROLE(&_RewardpoolCompoundContract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RewardpoolCompoundContract.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _RewardpoolCompoundContract.Contract.DEFAULTADMINROLE(&_RewardpoolCompoundContract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _RewardpoolCompoundContract.Contract.DEFAULTADMINROLE(&_RewardpoolCompoundContract.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractCaller) MANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RewardpoolCompoundContract.contract.Call(opts, &out, "MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractSession) MANAGERROLE() ([32]byte, error) {
	return _RewardpoolCompoundContract.Contract.MANAGERROLE(&_RewardpoolCompoundContract.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractCallerSession) MANAGERROLE() ([32]byte, error) {
	return _RewardpoolCompoundContract.Contract.MANAGERROLE(&_RewardpoolCompoundContract.CallOpts)
}

// ORACLEROLE is a free data retrieval call binding the contract method 0x07e2cea5.
//
// Solidity: function ORACLE_ROLE() view returns(bytes32)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractCaller) ORACLEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RewardpoolCompoundContract.contract.Call(opts, &out, "ORACLE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ORACLEROLE is a free data retrieval call binding the contract method 0x07e2cea5.
//
// Solidity: function ORACLE_ROLE() view returns(bytes32)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractSession) ORACLEROLE() ([32]byte, error) {
	return _RewardpoolCompoundContract.Contract.ORACLEROLE(&_RewardpoolCompoundContract.CallOpts)
}

// ORACLEROLE is a free data retrieval call binding the contract method 0x07e2cea5.
//
// Solidity: function ORACLE_ROLE() view returns(bytes32)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractCallerSession) ORACLEROLE() ([32]byte, error) {
	return _RewardpoolCompoundContract.Contract.ORACLEROLE(&_RewardpoolCompoundContract.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RewardpoolCompoundContract.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractSession) PAUSERROLE() ([32]byte, error) {
	return _RewardpoolCompoundContract.Contract.PAUSERROLE(&_RewardpoolCompoundContract.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractCallerSession) PAUSERROLE() ([32]byte, error) {
	return _RewardpoolCompoundContract.Contract.PAUSERROLE(&_RewardpoolCompoundContract.CallOpts)
}

// GetAccountedBalance is a free data retrieval call binding the contract method 0x33e5761f.
//
// Solidity: function getAccountedBalance() view returns(uint256)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractCaller) GetAccountedBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardpoolCompoundContract.contract.Call(opts, &out, "getAccountedBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountedBalance is a free data retrieval call binding the contract method 0x33e5761f.
//
// Solidity: function getAccountedBalance() view returns(uint256)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractSession) GetAccountedBalance() (*big.Int, error) {
	return _RewardpoolCompoundContract.Contract.GetAccountedBalance(&_RewardpoolCompoundContract.CallOpts)
}

// GetAccountedBalance is a free data retrieval call binding the contract method 0x33e5761f.
//
// Solidity: function getAccountedBalance() view returns(uint256)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractCallerSession) GetAccountedBalance() (*big.Int, error) {
	return _RewardpoolCompoundContract.Contract.GetAccountedBalance(&_RewardpoolCompoundContract.CallOpts)
}

// GetPendingManagerRevenue is a free data retrieval call binding the contract method 0xb8f932f6.
//
// Solidity: function getPendingManagerRevenue() view returns(uint256)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractCaller) GetPendingManagerRevenue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardpoolCompoundContract.contract.Call(opts, &out, "getPendingManagerRevenue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingManagerRevenue is a free data retrieval call binding the contract method 0xb8f932f6.
//
// Solidity: function getPendingManagerRevenue() view returns(uint256)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractSession) GetPendingManagerRevenue() (*big.Int, error) {
	return _RewardpoolCompoundContract.Contract.GetPendingManagerRevenue(&_RewardpoolCompoundContract.CallOpts)
}

// GetPendingManagerRevenue is a free data retrieval call binding the contract method 0xb8f932f6.
//
// Solidity: function getPendingManagerRevenue() view returns(uint256)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractCallerSession) GetPendingManagerRevenue() (*big.Int, error) {
	return _RewardpoolCompoundContract.Contract.GetPendingManagerRevenue(&_RewardpoolCompoundContract.CallOpts)
}

// GetPendingReward is a free data retrieval call binding the contract method 0x4df9d6ba.
//
// Solidity: function getPendingReward(address claimaddr) view returns(uint256)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractCaller) GetPendingReward(opts *bind.CallOpts, claimaddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardpoolCompoundContract.contract.Call(opts, &out, "getPendingReward", claimaddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingReward is a free data retrieval call binding the contract method 0x4df9d6ba.
//
// Solidity: function getPendingReward(address claimaddr) view returns(uint256)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractSession) GetPendingReward(claimaddr common.Address) (*big.Int, error) {
	return _RewardpoolCompoundContract.Contract.GetPendingReward(&_RewardpoolCompoundContract.CallOpts, claimaddr)
}

// GetPendingReward is a free data retrieval call binding the contract method 0x4df9d6ba.
//
// Solidity: function getPendingReward(address claimaddr) view returns(uint256)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractCallerSession) GetPendingReward(claimaddr common.Address) (*big.Int, error) {
	return _RewardpoolCompoundContract.Contract.GetPendingReward(&_RewardpoolCompoundContract.CallOpts, claimaddr)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _RewardpoolCompoundContract.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RewardpoolCompoundContract.Contract.GetRoleAdmin(&_RewardpoolCompoundContract.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RewardpoolCompoundContract.Contract.GetRoleAdmin(&_RewardpoolCompoundContract.CallOpts, role)
}

// GetTotalShare is a free data retrieval call binding the contract method 0x3f06cd66.
//
// Solidity: function getTotalShare() view returns(uint256)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractCaller) GetTotalShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardpoolCompoundContract.contract.Call(opts, &out, "getTotalShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalShare is a free data retrieval call binding the contract method 0x3f06cd66.
//
// Solidity: function getTotalShare() view returns(uint256)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractSession) GetTotalShare() (*big.Int, error) {
	return _RewardpoolCompoundContract.Contract.GetTotalShare(&_RewardpoolCompoundContract.CallOpts)
}

// GetTotalShare is a free data retrieval call binding the contract method 0x3f06cd66.
//
// Solidity: function getTotalShare() view returns(uint256)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractCallerSession) GetTotalShare() (*big.Int, error) {
	return _RewardpoolCompoundContract.Contract.GetTotalShare(&_RewardpoolCompoundContract.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _RewardpoolCompoundContract.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RewardpoolCompoundContract.Contract.HasRole(&_RewardpoolCompoundContract.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RewardpoolCompoundContract.Contract.HasRole(&_RewardpoolCompoundContract.CallOpts, role, account)
}

// ManagerFeeShare is a free data retrieval call binding the contract method 0xe43a4954.
//
// Solidity: function managerFeeShare() view returns(uint256)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractCaller) ManagerFeeShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardpoolCompoundContract.contract.Call(opts, &out, "managerFeeShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ManagerFeeShare is a free data retrieval call binding the contract method 0xe43a4954.
//
// Solidity: function managerFeeShare() view returns(uint256)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractSession) ManagerFeeShare() (*big.Int, error) {
	return _RewardpoolCompoundContract.Contract.ManagerFeeShare(&_RewardpoolCompoundContract.CallOpts)
}

// ManagerFeeShare is a free data retrieval call binding the contract method 0xe43a4954.
//
// Solidity: function managerFeeShare() view returns(uint256)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractCallerSession) ManagerFeeShare() (*big.Int, error) {
	return _RewardpoolCompoundContract.Contract.ManagerFeeShare(&_RewardpoolCompoundContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RewardpoolCompoundContract.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractSession) Paused() (bool, error) {
	return _RewardpoolCompoundContract.Contract.Paused(&_RewardpoolCompoundContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractCallerSession) Paused() (bool, error) {
	return _RewardpoolCompoundContract.Contract.Paused(&_RewardpoolCompoundContract.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _RewardpoolCompoundContract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RewardpoolCompoundContract.Contract.SupportsInterface(&_RewardpoolCompoundContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RewardpoolCompoundContract.Contract.SupportsInterface(&_RewardpoolCompoundContract.CallOpts, interfaceId)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 accSharePoint, int256 amount, uint256 rewardBalance)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractCaller) UserInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	AccSharePoint *big.Int
	Amount        *big.Int
	RewardBalance *big.Int
}, error) {
	var out []interface{}
	err := _RewardpoolCompoundContract.contract.Call(opts, &out, "userInfo", arg0)

	outstruct := new(struct {
		AccSharePoint *big.Int
		Amount        *big.Int
		RewardBalance *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AccSharePoint = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RewardBalance = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 accSharePoint, int256 amount, uint256 rewardBalance)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractSession) UserInfo(arg0 common.Address) (struct {
	AccSharePoint *big.Int
	Amount        *big.Int
	RewardBalance *big.Int
}, error) {
	return _RewardpoolCompoundContract.Contract.UserInfo(&_RewardpoolCompoundContract.CallOpts, arg0)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 accSharePoint, int256 amount, uint256 rewardBalance)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractCallerSession) UserInfo(arg0 common.Address) (struct {
	AccSharePoint *big.Int
	Amount        *big.Int
	RewardBalance *big.Int
}, error) {
	return _RewardpoolCompoundContract.Contract.UserInfo(&_RewardpoolCompoundContract.CallOpts, arg0)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x9a99b4f0.
//
// Solidity: function claimRewards(address beneficiary, uint256 amount) returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractTransactor) ClaimRewards(opts *bind.TransactOpts, beneficiary common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RewardpoolCompoundContract.contract.Transact(opts, "claimRewards", beneficiary, amount)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x9a99b4f0.
//
// Solidity: function claimRewards(address beneficiary, uint256 amount) returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractSession) ClaimRewards(beneficiary common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RewardpoolCompoundContract.Contract.ClaimRewards(&_RewardpoolCompoundContract.TransactOpts, beneficiary, amount)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x9a99b4f0.
//
// Solidity: function claimRewards(address beneficiary, uint256 amount) returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractTransactorSession) ClaimRewards(beneficiary common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RewardpoolCompoundContract.Contract.ClaimRewards(&_RewardpoolCompoundContract.TransactOpts, beneficiary, amount)
}

// ClaimRewardsFor is a paid mutator transaction binding the contract method 0x1ac6d19d.
//
// Solidity: function claimRewardsFor(address account) returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractTransactor) ClaimRewardsFor(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _RewardpoolCompoundContract.contract.Transact(opts, "claimRewardsFor", account)
}

// ClaimRewardsFor is a paid mutator transaction binding the contract method 0x1ac6d19d.
//
// Solidity: function claimRewardsFor(address account) returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractSession) ClaimRewardsFor(account common.Address) (*types.Transaction, error) {
	return _RewardpoolCompoundContract.Contract.ClaimRewardsFor(&_RewardpoolCompoundContract.TransactOpts, account)
}

// ClaimRewardsFor is a paid mutator transaction binding the contract method 0x1ac6d19d.
//
// Solidity: function claimRewardsFor(address account) returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractTransactorSession) ClaimRewardsFor(account common.Address) (*types.Transaction, error) {
	return _RewardpoolCompoundContract.Contract.ClaimRewardsFor(&_RewardpoolCompoundContract.TransactOpts, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewardpoolCompoundContract.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewardpoolCompoundContract.Contract.GrantRole(&_RewardpoolCompoundContract.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewardpoolCompoundContract.Contract.GrantRole(&_RewardpoolCompoundContract.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardpoolCompoundContract.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractSession) Initialize() (*types.Transaction, error) {
	return _RewardpoolCompoundContract.Contract.Initialize(&_RewardpoolCompoundContract.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractTransactorSession) Initialize() (*types.Transaction, error) {
	return _RewardpoolCompoundContract.Contract.Initialize(&_RewardpoolCompoundContract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardpoolCompoundContract.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractSession) Pause() (*types.Transaction, error) {
	return _RewardpoolCompoundContract.Contract.Pause(&_RewardpoolCompoundContract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractTransactorSession) Pause() (*types.Transaction, error) {
	return _RewardpoolCompoundContract.Contract.Pause(&_RewardpoolCompoundContract.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewardpoolCompoundContract.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewardpoolCompoundContract.Contract.RenounceRole(&_RewardpoolCompoundContract.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewardpoolCompoundContract.Contract.RenounceRole(&_RewardpoolCompoundContract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewardpoolCompoundContract.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewardpoolCompoundContract.Contract.RevokeRole(&_RewardpoolCompoundContract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewardpoolCompoundContract.Contract.RevokeRole(&_RewardpoolCompoundContract.TransactOpts, role, account)
}

// SetManagerFeeShare is a paid mutator transaction binding the contract method 0x755d7dd3.
//
// Solidity: function setManagerFeeShare(uint256 milli) returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractTransactor) SetManagerFeeShare(opts *bind.TransactOpts, milli *big.Int) (*types.Transaction, error) {
	return _RewardpoolCompoundContract.contract.Transact(opts, "setManagerFeeShare", milli)
}

// SetManagerFeeShare is a paid mutator transaction binding the contract method 0x755d7dd3.
//
// Solidity: function setManagerFeeShare(uint256 milli) returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractSession) SetManagerFeeShare(milli *big.Int) (*types.Transaction, error) {
	return _RewardpoolCompoundContract.Contract.SetManagerFeeShare(&_RewardpoolCompoundContract.TransactOpts, milli)
}

// SetManagerFeeShare is a paid mutator transaction binding the contract method 0x755d7dd3.
//
// Solidity: function setManagerFeeShare(uint256 milli) returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractTransactorSession) SetManagerFeeShare(milli *big.Int) (*types.Transaction, error) {
	return _RewardpoolCompoundContract.Contract.SetManagerFeeShare(&_RewardpoolCompoundContract.TransactOpts, milli)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardpoolCompoundContract.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractSession) Unpause() (*types.Transaction, error) {
	return _RewardpoolCompoundContract.Contract.Unpause(&_RewardpoolCompoundContract.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractTransactorSession) Unpause() (*types.Transaction, error) {
	return _RewardpoolCompoundContract.Contract.Unpause(&_RewardpoolCompoundContract.TransactOpts)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x2a17c7e5.
//
// Solidity: function updatePool(address[] claimaddrs, int256[] deltas) returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractTransactor) UpdatePool(opts *bind.TransactOpts, claimaddrs []common.Address, deltas []*big.Int) (*types.Transaction, error) {
	return _RewardpoolCompoundContract.contract.Transact(opts, "updatePool", claimaddrs, deltas)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x2a17c7e5.
//
// Solidity: function updatePool(address[] claimaddrs, int256[] deltas) returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractSession) UpdatePool(claimaddrs []common.Address, deltas []*big.Int) (*types.Transaction, error) {
	return _RewardpoolCompoundContract.Contract.UpdatePool(&_RewardpoolCompoundContract.TransactOpts, claimaddrs, deltas)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x2a17c7e5.
//
// Solidity: function updatePool(address[] claimaddrs, int256[] deltas) returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractTransactorSession) UpdatePool(claimaddrs []common.Address, deltas []*big.Int) (*types.Transaction, error) {
	return _RewardpoolCompoundContract.Contract.UpdatePool(&_RewardpoolCompoundContract.TransactOpts, claimaddrs, deltas)
}

// WithdrawManagerRevenue is a paid mutator transaction binding the contract method 0xac9a5f56.
//
// Solidity: function withdrawManagerRevenue(uint256 amount, address to) returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractTransactor) WithdrawManagerRevenue(opts *bind.TransactOpts, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _RewardpoolCompoundContract.contract.Transact(opts, "withdrawManagerRevenue", amount, to)
}

// WithdrawManagerRevenue is a paid mutator transaction binding the contract method 0xac9a5f56.
//
// Solidity: function withdrawManagerRevenue(uint256 amount, address to) returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractSession) WithdrawManagerRevenue(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _RewardpoolCompoundContract.Contract.WithdrawManagerRevenue(&_RewardpoolCompoundContract.TransactOpts, amount, to)
}

// WithdrawManagerRevenue is a paid mutator transaction binding the contract method 0xac9a5f56.
//
// Solidity: function withdrawManagerRevenue(uint256 amount, address to) returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractTransactorSession) WithdrawManagerRevenue(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _RewardpoolCompoundContract.Contract.WithdrawManagerRevenue(&_RewardpoolCompoundContract.TransactOpts, amount, to)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardpoolCompoundContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractSession) Receive() (*types.Transaction, error) {
	return _RewardpoolCompoundContract.Contract.Receive(&_RewardpoolCompoundContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RewardpoolCompoundContract *RewardpoolCompoundContractTransactorSession) Receive() (*types.Transaction, error) {
	return _RewardpoolCompoundContract.Contract.Receive(&_RewardpoolCompoundContract.TransactOpts)
}

// RewardpoolCompoundContractClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the RewardpoolCompoundContract contract.
type RewardpoolCompoundContractClaimedIterator struct {
	Event *RewardpoolCompoundContractClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardpoolCompoundContractClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardpoolCompoundContractClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardpoolCompoundContractClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardpoolCompoundContractClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardpoolCompoundContractClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardpoolCompoundContractClaimed represents a Claimed event raised by the RewardpoolCompoundContract contract.
type RewardpoolCompoundContractClaimed struct {
	Beneficiary common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address beneficiary, uint256 amount)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractFilterer) FilterClaimed(opts *bind.FilterOpts) (*RewardpoolCompoundContractClaimedIterator, error) {

	logs, sub, err := _RewardpoolCompoundContract.contract.FilterLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return &RewardpoolCompoundContractClaimedIterator{contract: _RewardpoolCompoundContract.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address beneficiary, uint256 amount)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *RewardpoolCompoundContractClaimed) (event.Subscription, error) {

	logs, sub, err := _RewardpoolCompoundContract.contract.WatchLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardpoolCompoundContractClaimed)
				if err := _RewardpoolCompoundContract.contract.UnpackLog(event, "Claimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimed is a log parse operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address beneficiary, uint256 amount)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractFilterer) ParseClaimed(log types.Log) (*RewardpoolCompoundContractClaimed, error) {
	event := new(RewardpoolCompoundContractClaimed)
	if err := _RewardpoolCompoundContract.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardpoolCompoundContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the RewardpoolCompoundContract contract.
type RewardpoolCompoundContractInitializedIterator struct {
	Event *RewardpoolCompoundContractInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardpoolCompoundContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardpoolCompoundContractInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardpoolCompoundContractInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardpoolCompoundContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardpoolCompoundContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardpoolCompoundContractInitialized represents a Initialized event raised by the RewardpoolCompoundContract contract.
type RewardpoolCompoundContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*RewardpoolCompoundContractInitializedIterator, error) {

	logs, sub, err := _RewardpoolCompoundContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RewardpoolCompoundContractInitializedIterator{contract: _RewardpoolCompoundContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RewardpoolCompoundContractInitialized) (event.Subscription, error) {

	logs, sub, err := _RewardpoolCompoundContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardpoolCompoundContractInitialized)
				if err := _RewardpoolCompoundContract.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractFilterer) ParseInitialized(log types.Log) (*RewardpoolCompoundContractInitialized, error) {
	event := new(RewardpoolCompoundContractInitialized)
	if err := _RewardpoolCompoundContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardpoolCompoundContractManagerFeeSetIterator is returned from FilterManagerFeeSet and is used to iterate over the raw logs and unpacked data for ManagerFeeSet events raised by the RewardpoolCompoundContract contract.
type RewardpoolCompoundContractManagerFeeSetIterator struct {
	Event *RewardpoolCompoundContractManagerFeeSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardpoolCompoundContractManagerFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardpoolCompoundContractManagerFeeSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardpoolCompoundContractManagerFeeSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardpoolCompoundContractManagerFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardpoolCompoundContractManagerFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardpoolCompoundContractManagerFeeSet represents a ManagerFeeSet event raised by the RewardpoolCompoundContract contract.
type RewardpoolCompoundContractManagerFeeSet struct {
	Milli *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterManagerFeeSet is a free log retrieval operation binding the contract event 0x4de90ec86e1bc56c192e2399bacbd10bdaba720caca606354d66c5cb33d6802b.
//
// Solidity: event ManagerFeeSet(uint256 milli)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractFilterer) FilterManagerFeeSet(opts *bind.FilterOpts) (*RewardpoolCompoundContractManagerFeeSetIterator, error) {

	logs, sub, err := _RewardpoolCompoundContract.contract.FilterLogs(opts, "ManagerFeeSet")
	if err != nil {
		return nil, err
	}
	return &RewardpoolCompoundContractManagerFeeSetIterator{contract: _RewardpoolCompoundContract.contract, event: "ManagerFeeSet", logs: logs, sub: sub}, nil
}

// WatchManagerFeeSet is a free log subscription operation binding the contract event 0x4de90ec86e1bc56c192e2399bacbd10bdaba720caca606354d66c5cb33d6802b.
//
// Solidity: event ManagerFeeSet(uint256 milli)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractFilterer) WatchManagerFeeSet(opts *bind.WatchOpts, sink chan<- *RewardpoolCompoundContractManagerFeeSet) (event.Subscription, error) {

	logs, sub, err := _RewardpoolCompoundContract.contract.WatchLogs(opts, "ManagerFeeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardpoolCompoundContractManagerFeeSet)
				if err := _RewardpoolCompoundContract.contract.UnpackLog(event, "ManagerFeeSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseManagerFeeSet is a log parse operation binding the contract event 0x4de90ec86e1bc56c192e2399bacbd10bdaba720caca606354d66c5cb33d6802b.
//
// Solidity: event ManagerFeeSet(uint256 milli)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractFilterer) ParseManagerFeeSet(log types.Log) (*RewardpoolCompoundContractManagerFeeSet, error) {
	event := new(RewardpoolCompoundContractManagerFeeSet)
	if err := _RewardpoolCompoundContract.contract.UnpackLog(event, "ManagerFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardpoolCompoundContractManagerFeeWithdrawedIterator is returned from FilterManagerFeeWithdrawed and is used to iterate over the raw logs and unpacked data for ManagerFeeWithdrawed events raised by the RewardpoolCompoundContract contract.
type RewardpoolCompoundContractManagerFeeWithdrawedIterator struct {
	Event *RewardpoolCompoundContractManagerFeeWithdrawed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardpoolCompoundContractManagerFeeWithdrawedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardpoolCompoundContractManagerFeeWithdrawed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardpoolCompoundContractManagerFeeWithdrawed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardpoolCompoundContractManagerFeeWithdrawedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardpoolCompoundContractManagerFeeWithdrawedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardpoolCompoundContractManagerFeeWithdrawed represents a ManagerFeeWithdrawed event raised by the RewardpoolCompoundContract contract.
type RewardpoolCompoundContractManagerFeeWithdrawed struct {
	Amount *big.Int
	To     common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterManagerFeeWithdrawed is a free log retrieval operation binding the contract event 0x2425aa1fadefc5c570850aa9c9e3dfa4fc6b43ccd1c05b47db38dd6518a743b3.
//
// Solidity: event ManagerFeeWithdrawed(uint256 amount, address to)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractFilterer) FilterManagerFeeWithdrawed(opts *bind.FilterOpts) (*RewardpoolCompoundContractManagerFeeWithdrawedIterator, error) {

	logs, sub, err := _RewardpoolCompoundContract.contract.FilterLogs(opts, "ManagerFeeWithdrawed")
	if err != nil {
		return nil, err
	}
	return &RewardpoolCompoundContractManagerFeeWithdrawedIterator{contract: _RewardpoolCompoundContract.contract, event: "ManagerFeeWithdrawed", logs: logs, sub: sub}, nil
}

// WatchManagerFeeWithdrawed is a free log subscription operation binding the contract event 0x2425aa1fadefc5c570850aa9c9e3dfa4fc6b43ccd1c05b47db38dd6518a743b3.
//
// Solidity: event ManagerFeeWithdrawed(uint256 amount, address to)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractFilterer) WatchManagerFeeWithdrawed(opts *bind.WatchOpts, sink chan<- *RewardpoolCompoundContractManagerFeeWithdrawed) (event.Subscription, error) {

	logs, sub, err := _RewardpoolCompoundContract.contract.WatchLogs(opts, "ManagerFeeWithdrawed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardpoolCompoundContractManagerFeeWithdrawed)
				if err := _RewardpoolCompoundContract.contract.UnpackLog(event, "ManagerFeeWithdrawed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseManagerFeeWithdrawed is a log parse operation binding the contract event 0x2425aa1fadefc5c570850aa9c9e3dfa4fc6b43ccd1c05b47db38dd6518a743b3.
//
// Solidity: event ManagerFeeWithdrawed(uint256 amount, address to)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractFilterer) ParseManagerFeeWithdrawed(log types.Log) (*RewardpoolCompoundContractManagerFeeWithdrawed, error) {
	event := new(RewardpoolCompoundContractManagerFeeWithdrawed)
	if err := _RewardpoolCompoundContract.contract.UnpackLog(event, "ManagerFeeWithdrawed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardpoolCompoundContractPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the RewardpoolCompoundContract contract.
type RewardpoolCompoundContractPausedIterator struct {
	Event *RewardpoolCompoundContractPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardpoolCompoundContractPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardpoolCompoundContractPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardpoolCompoundContractPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardpoolCompoundContractPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardpoolCompoundContractPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardpoolCompoundContractPaused represents a Paused event raised by the RewardpoolCompoundContract contract.
type RewardpoolCompoundContractPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractFilterer) FilterPaused(opts *bind.FilterOpts) (*RewardpoolCompoundContractPausedIterator, error) {

	logs, sub, err := _RewardpoolCompoundContract.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &RewardpoolCompoundContractPausedIterator{contract: _RewardpoolCompoundContract.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *RewardpoolCompoundContractPaused) (event.Subscription, error) {

	logs, sub, err := _RewardpoolCompoundContract.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardpoolCompoundContractPaused)
				if err := _RewardpoolCompoundContract.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractFilterer) ParsePaused(log types.Log) (*RewardpoolCompoundContractPaused, error) {
	event := new(RewardpoolCompoundContractPaused)
	if err := _RewardpoolCompoundContract.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardpoolCompoundContractPoolUpdateIterator is returned from FilterPoolUpdate and is used to iterate over the raw logs and unpacked data for PoolUpdate events raised by the RewardpoolCompoundContract contract.
type RewardpoolCompoundContractPoolUpdateIterator struct {
	Event *RewardpoolCompoundContractPoolUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardpoolCompoundContractPoolUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardpoolCompoundContractPoolUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardpoolCompoundContractPoolUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardpoolCompoundContractPoolUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardpoolCompoundContractPoolUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardpoolCompoundContractPoolUpdate represents a PoolUpdate event raised by the RewardpoolCompoundContract contract.
type RewardpoolCompoundContractPoolUpdate struct {
	Claimaddr common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPoolUpdate is a free log retrieval operation binding the contract event 0x4bdf48368af4e58208498c0a0ed00f8fd5cb933a966be8e72600dc296742662b.
//
// Solidity: event PoolUpdate(address claimaddr, int256 amount)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractFilterer) FilterPoolUpdate(opts *bind.FilterOpts) (*RewardpoolCompoundContractPoolUpdateIterator, error) {

	logs, sub, err := _RewardpoolCompoundContract.contract.FilterLogs(opts, "PoolUpdate")
	if err != nil {
		return nil, err
	}
	return &RewardpoolCompoundContractPoolUpdateIterator{contract: _RewardpoolCompoundContract.contract, event: "PoolUpdate", logs: logs, sub: sub}, nil
}

// WatchPoolUpdate is a free log subscription operation binding the contract event 0x4bdf48368af4e58208498c0a0ed00f8fd5cb933a966be8e72600dc296742662b.
//
// Solidity: event PoolUpdate(address claimaddr, int256 amount)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractFilterer) WatchPoolUpdate(opts *bind.WatchOpts, sink chan<- *RewardpoolCompoundContractPoolUpdate) (event.Subscription, error) {

	logs, sub, err := _RewardpoolCompoundContract.contract.WatchLogs(opts, "PoolUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardpoolCompoundContractPoolUpdate)
				if err := _RewardpoolCompoundContract.contract.UnpackLog(event, "PoolUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePoolUpdate is a log parse operation binding the contract event 0x4bdf48368af4e58208498c0a0ed00f8fd5cb933a966be8e72600dc296742662b.
//
// Solidity: event PoolUpdate(address claimaddr, int256 amount)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractFilterer) ParsePoolUpdate(log types.Log) (*RewardpoolCompoundContractPoolUpdate, error) {
	event := new(RewardpoolCompoundContractPoolUpdate)
	if err := _RewardpoolCompoundContract.contract.UnpackLog(event, "PoolUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardpoolCompoundContractRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the RewardpoolCompoundContract contract.
type RewardpoolCompoundContractRoleAdminChangedIterator struct {
	Event *RewardpoolCompoundContractRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardpoolCompoundContractRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardpoolCompoundContractRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardpoolCompoundContractRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardpoolCompoundContractRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardpoolCompoundContractRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardpoolCompoundContractRoleAdminChanged represents a RoleAdminChanged event raised by the RewardpoolCompoundContract contract.
type RewardpoolCompoundContractRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*RewardpoolCompoundContractRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _RewardpoolCompoundContract.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &RewardpoolCompoundContractRoleAdminChangedIterator{contract: _RewardpoolCompoundContract.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *RewardpoolCompoundContractRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _RewardpoolCompoundContract.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardpoolCompoundContractRoleAdminChanged)
				if err := _RewardpoolCompoundContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractFilterer) ParseRoleAdminChanged(log types.Log) (*RewardpoolCompoundContractRoleAdminChanged, error) {
	event := new(RewardpoolCompoundContractRoleAdminChanged)
	if err := _RewardpoolCompoundContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardpoolCompoundContractRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the RewardpoolCompoundContract contract.
type RewardpoolCompoundContractRoleGrantedIterator struct {
	Event *RewardpoolCompoundContractRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardpoolCompoundContractRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardpoolCompoundContractRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardpoolCompoundContractRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardpoolCompoundContractRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardpoolCompoundContractRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardpoolCompoundContractRoleGranted represents a RoleGranted event raised by the RewardpoolCompoundContract contract.
type RewardpoolCompoundContractRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RewardpoolCompoundContractRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RewardpoolCompoundContract.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RewardpoolCompoundContractRoleGrantedIterator{contract: _RewardpoolCompoundContract.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RewardpoolCompoundContractRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RewardpoolCompoundContract.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardpoolCompoundContractRoleGranted)
				if err := _RewardpoolCompoundContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractFilterer) ParseRoleGranted(log types.Log) (*RewardpoolCompoundContractRoleGranted, error) {
	event := new(RewardpoolCompoundContractRoleGranted)
	if err := _RewardpoolCompoundContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardpoolCompoundContractRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the RewardpoolCompoundContract contract.
type RewardpoolCompoundContractRoleRevokedIterator struct {
	Event *RewardpoolCompoundContractRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardpoolCompoundContractRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardpoolCompoundContractRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardpoolCompoundContractRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardpoolCompoundContractRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardpoolCompoundContractRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardpoolCompoundContractRoleRevoked represents a RoleRevoked event raised by the RewardpoolCompoundContract contract.
type RewardpoolCompoundContractRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RewardpoolCompoundContractRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RewardpoolCompoundContract.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RewardpoolCompoundContractRoleRevokedIterator{contract: _RewardpoolCompoundContract.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RewardpoolCompoundContractRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RewardpoolCompoundContract.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardpoolCompoundContractRoleRevoked)
				if err := _RewardpoolCompoundContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractFilterer) ParseRoleRevoked(log types.Log) (*RewardpoolCompoundContractRoleRevoked, error) {
	event := new(RewardpoolCompoundContractRoleRevoked)
	if err := _RewardpoolCompoundContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardpoolCompoundContractUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the RewardpoolCompoundContract contract.
type RewardpoolCompoundContractUnpausedIterator struct {
	Event *RewardpoolCompoundContractUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardpoolCompoundContractUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardpoolCompoundContractUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardpoolCompoundContractUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardpoolCompoundContractUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardpoolCompoundContractUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardpoolCompoundContractUnpaused represents a Unpaused event raised by the RewardpoolCompoundContract contract.
type RewardpoolCompoundContractUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractFilterer) FilterUnpaused(opts *bind.FilterOpts) (*RewardpoolCompoundContractUnpausedIterator, error) {

	logs, sub, err := _RewardpoolCompoundContract.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &RewardpoolCompoundContractUnpausedIterator{contract: _RewardpoolCompoundContract.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *RewardpoolCompoundContractUnpaused) (event.Subscription, error) {

	logs, sub, err := _RewardpoolCompoundContract.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardpoolCompoundContractUnpaused)
				if err := _RewardpoolCompoundContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_RewardpoolCompoundContract *RewardpoolCompoundContractFilterer) ParseUnpaused(log types.Log) (*RewardpoolCompoundContractUnpaused, error) {
	event := new(RewardpoolCompoundContractUnpaused)
	if err := _RewardpoolCompoundContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
