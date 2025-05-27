// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rewardpool_contract

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

// RewardpoolContractMetaData contains all meta data concerning the RewardpoolContract contract.
var RewardpoolContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"CONTROLLER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MANAGER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ORACLE_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimRewards\",\"inputs\":[{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimRewardsFor\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAccountedBalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingManagerRevenue\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingReward\",\"inputs\":[{\"name\":\"claimaddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalShare\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"managerFeeShare\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setManagerFeeShare\",\"inputs\":[{\"name\":\"milli\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updatePool\",\"inputs\":[{\"name\":\"claimaddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"deltas\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userInfo\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"accSharePoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rewardBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawManagerRevenue\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Claimed\",\"inputs\":[{\"name\":\"beneficiary\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ManagerFeeSet\",\"inputs\":[{\"name\":\"milli\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ManagerFeeWithdrawed\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolUpdate\",\"inputs\":[{\"name\":\"claimaddr\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"amount\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// RewardpoolContractABI is the input ABI used to generate the binding from.
// Deprecated: Use RewardpoolContractMetaData.ABI instead.
var RewardpoolContractABI = RewardpoolContractMetaData.ABI

// RewardpoolContract is an auto generated Go binding around an Ethereum contract.
type RewardpoolContract struct {
	RewardpoolContractCaller     // Read-only binding to the contract
	RewardpoolContractTransactor // Write-only binding to the contract
	RewardpoolContractFilterer   // Log filterer for contract events
}

// RewardpoolContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type RewardpoolContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardpoolContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RewardpoolContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardpoolContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewardpoolContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardpoolContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewardpoolContractSession struct {
	Contract     *RewardpoolContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RewardpoolContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewardpoolContractCallerSession struct {
	Contract *RewardpoolContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// RewardpoolContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewardpoolContractTransactorSession struct {
	Contract     *RewardpoolContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// RewardpoolContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type RewardpoolContractRaw struct {
	Contract *RewardpoolContract // Generic contract binding to access the raw methods on
}

// RewardpoolContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewardpoolContractCallerRaw struct {
	Contract *RewardpoolContractCaller // Generic read-only contract binding to access the raw methods on
}

// RewardpoolContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewardpoolContractTransactorRaw struct {
	Contract *RewardpoolContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRewardpoolContract creates a new instance of RewardpoolContract, bound to a specific deployed contract.
func NewRewardpoolContract(address common.Address, backend bind.ContractBackend) (*RewardpoolContract, error) {
	contract, err := bindRewardpoolContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RewardpoolContract{RewardpoolContractCaller: RewardpoolContractCaller{contract: contract}, RewardpoolContractTransactor: RewardpoolContractTransactor{contract: contract}, RewardpoolContractFilterer: RewardpoolContractFilterer{contract: contract}}, nil
}

// NewRewardpoolContractCaller creates a new read-only instance of RewardpoolContract, bound to a specific deployed contract.
func NewRewardpoolContractCaller(address common.Address, caller bind.ContractCaller) (*RewardpoolContractCaller, error) {
	contract, err := bindRewardpoolContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewardpoolContractCaller{contract: contract}, nil
}

// NewRewardpoolContractTransactor creates a new write-only instance of RewardpoolContract, bound to a specific deployed contract.
func NewRewardpoolContractTransactor(address common.Address, transactor bind.ContractTransactor) (*RewardpoolContractTransactor, error) {
	contract, err := bindRewardpoolContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewardpoolContractTransactor{contract: contract}, nil
}

// NewRewardpoolContractFilterer creates a new log filterer instance of RewardpoolContract, bound to a specific deployed contract.
func NewRewardpoolContractFilterer(address common.Address, filterer bind.ContractFilterer) (*RewardpoolContractFilterer, error) {
	contract, err := bindRewardpoolContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewardpoolContractFilterer{contract: contract}, nil
}

// bindRewardpoolContract binds a generic wrapper to an already deployed contract.
func bindRewardpoolContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RewardpoolContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardpoolContract *RewardpoolContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardpoolContract.Contract.RewardpoolContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardpoolContract *RewardpoolContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.RewardpoolContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardpoolContract *RewardpoolContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.RewardpoolContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardpoolContract *RewardpoolContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardpoolContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardpoolContract *RewardpoolContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardpoolContract *RewardpoolContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.contract.Transact(opts, method, params...)
}

// CONTROLLERROLE is a free data retrieval call binding the contract method 0x092c5b3b.
//
// Solidity: function CONTROLLER_ROLE() view returns(bytes32)
func (_RewardpoolContract *RewardpoolContractCaller) CONTROLLERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RewardpoolContract.contract.Call(opts, &out, "CONTROLLER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CONTROLLERROLE is a free data retrieval call binding the contract method 0x092c5b3b.
//
// Solidity: function CONTROLLER_ROLE() view returns(bytes32)
func (_RewardpoolContract *RewardpoolContractSession) CONTROLLERROLE() ([32]byte, error) {
	return _RewardpoolContract.Contract.CONTROLLERROLE(&_RewardpoolContract.CallOpts)
}

// CONTROLLERROLE is a free data retrieval call binding the contract method 0x092c5b3b.
//
// Solidity: function CONTROLLER_ROLE() view returns(bytes32)
func (_RewardpoolContract *RewardpoolContractCallerSession) CONTROLLERROLE() ([32]byte, error) {
	return _RewardpoolContract.Contract.CONTROLLERROLE(&_RewardpoolContract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RewardpoolContract *RewardpoolContractCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RewardpoolContract.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RewardpoolContract *RewardpoolContractSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _RewardpoolContract.Contract.DEFAULTADMINROLE(&_RewardpoolContract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RewardpoolContract *RewardpoolContractCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _RewardpoolContract.Contract.DEFAULTADMINROLE(&_RewardpoolContract.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_RewardpoolContract *RewardpoolContractCaller) MANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RewardpoolContract.contract.Call(opts, &out, "MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_RewardpoolContract *RewardpoolContractSession) MANAGERROLE() ([32]byte, error) {
	return _RewardpoolContract.Contract.MANAGERROLE(&_RewardpoolContract.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_RewardpoolContract *RewardpoolContractCallerSession) MANAGERROLE() ([32]byte, error) {
	return _RewardpoolContract.Contract.MANAGERROLE(&_RewardpoolContract.CallOpts)
}

// ORACLEROLE is a free data retrieval call binding the contract method 0x07e2cea5.
//
// Solidity: function ORACLE_ROLE() view returns(bytes32)
func (_RewardpoolContract *RewardpoolContractCaller) ORACLEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RewardpoolContract.contract.Call(opts, &out, "ORACLE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ORACLEROLE is a free data retrieval call binding the contract method 0x07e2cea5.
//
// Solidity: function ORACLE_ROLE() view returns(bytes32)
func (_RewardpoolContract *RewardpoolContractSession) ORACLEROLE() ([32]byte, error) {
	return _RewardpoolContract.Contract.ORACLEROLE(&_RewardpoolContract.CallOpts)
}

// ORACLEROLE is a free data retrieval call binding the contract method 0x07e2cea5.
//
// Solidity: function ORACLE_ROLE() view returns(bytes32)
func (_RewardpoolContract *RewardpoolContractCallerSession) ORACLEROLE() ([32]byte, error) {
	return _RewardpoolContract.Contract.ORACLEROLE(&_RewardpoolContract.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_RewardpoolContract *RewardpoolContractCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RewardpoolContract.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_RewardpoolContract *RewardpoolContractSession) PAUSERROLE() ([32]byte, error) {
	return _RewardpoolContract.Contract.PAUSERROLE(&_RewardpoolContract.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_RewardpoolContract *RewardpoolContractCallerSession) PAUSERROLE() ([32]byte, error) {
	return _RewardpoolContract.Contract.PAUSERROLE(&_RewardpoolContract.CallOpts)
}

// GetAccountedBalance is a free data retrieval call binding the contract method 0x33e5761f.
//
// Solidity: function getAccountedBalance() view returns(uint256)
func (_RewardpoolContract *RewardpoolContractCaller) GetAccountedBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardpoolContract.contract.Call(opts, &out, "getAccountedBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountedBalance is a free data retrieval call binding the contract method 0x33e5761f.
//
// Solidity: function getAccountedBalance() view returns(uint256)
func (_RewardpoolContract *RewardpoolContractSession) GetAccountedBalance() (*big.Int, error) {
	return _RewardpoolContract.Contract.GetAccountedBalance(&_RewardpoolContract.CallOpts)
}

// GetAccountedBalance is a free data retrieval call binding the contract method 0x33e5761f.
//
// Solidity: function getAccountedBalance() view returns(uint256)
func (_RewardpoolContract *RewardpoolContractCallerSession) GetAccountedBalance() (*big.Int, error) {
	return _RewardpoolContract.Contract.GetAccountedBalance(&_RewardpoolContract.CallOpts)
}

// GetPendingManagerRevenue is a free data retrieval call binding the contract method 0xb8f932f6.
//
// Solidity: function getPendingManagerRevenue() view returns(uint256)
func (_RewardpoolContract *RewardpoolContractCaller) GetPendingManagerRevenue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardpoolContract.contract.Call(opts, &out, "getPendingManagerRevenue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingManagerRevenue is a free data retrieval call binding the contract method 0xb8f932f6.
//
// Solidity: function getPendingManagerRevenue() view returns(uint256)
func (_RewardpoolContract *RewardpoolContractSession) GetPendingManagerRevenue() (*big.Int, error) {
	return _RewardpoolContract.Contract.GetPendingManagerRevenue(&_RewardpoolContract.CallOpts)
}

// GetPendingManagerRevenue is a free data retrieval call binding the contract method 0xb8f932f6.
//
// Solidity: function getPendingManagerRevenue() view returns(uint256)
func (_RewardpoolContract *RewardpoolContractCallerSession) GetPendingManagerRevenue() (*big.Int, error) {
	return _RewardpoolContract.Contract.GetPendingManagerRevenue(&_RewardpoolContract.CallOpts)
}

// GetPendingReward is a free data retrieval call binding the contract method 0x4df9d6ba.
//
// Solidity: function getPendingReward(address claimaddr) view returns(uint256)
func (_RewardpoolContract *RewardpoolContractCaller) GetPendingReward(opts *bind.CallOpts, claimaddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardpoolContract.contract.Call(opts, &out, "getPendingReward", claimaddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingReward is a free data retrieval call binding the contract method 0x4df9d6ba.
//
// Solidity: function getPendingReward(address claimaddr) view returns(uint256)
func (_RewardpoolContract *RewardpoolContractSession) GetPendingReward(claimaddr common.Address) (*big.Int, error) {
	return _RewardpoolContract.Contract.GetPendingReward(&_RewardpoolContract.CallOpts, claimaddr)
}

// GetPendingReward is a free data retrieval call binding the contract method 0x4df9d6ba.
//
// Solidity: function getPendingReward(address claimaddr) view returns(uint256)
func (_RewardpoolContract *RewardpoolContractCallerSession) GetPendingReward(claimaddr common.Address) (*big.Int, error) {
	return _RewardpoolContract.Contract.GetPendingReward(&_RewardpoolContract.CallOpts, claimaddr)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RewardpoolContract *RewardpoolContractCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _RewardpoolContract.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RewardpoolContract *RewardpoolContractSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RewardpoolContract.Contract.GetRoleAdmin(&_RewardpoolContract.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RewardpoolContract *RewardpoolContractCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RewardpoolContract.Contract.GetRoleAdmin(&_RewardpoolContract.CallOpts, role)
}

// GetTotalShare is a free data retrieval call binding the contract method 0x3f06cd66.
//
// Solidity: function getTotalShare() view returns(uint256)
func (_RewardpoolContract *RewardpoolContractCaller) GetTotalShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardpoolContract.contract.Call(opts, &out, "getTotalShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalShare is a free data retrieval call binding the contract method 0x3f06cd66.
//
// Solidity: function getTotalShare() view returns(uint256)
func (_RewardpoolContract *RewardpoolContractSession) GetTotalShare() (*big.Int, error) {
	return _RewardpoolContract.Contract.GetTotalShare(&_RewardpoolContract.CallOpts)
}

// GetTotalShare is a free data retrieval call binding the contract method 0x3f06cd66.
//
// Solidity: function getTotalShare() view returns(uint256)
func (_RewardpoolContract *RewardpoolContractCallerSession) GetTotalShare() (*big.Int, error) {
	return _RewardpoolContract.Contract.GetTotalShare(&_RewardpoolContract.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RewardpoolContract *RewardpoolContractCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _RewardpoolContract.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RewardpoolContract *RewardpoolContractSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RewardpoolContract.Contract.HasRole(&_RewardpoolContract.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RewardpoolContract *RewardpoolContractCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RewardpoolContract.Contract.HasRole(&_RewardpoolContract.CallOpts, role, account)
}

// ManagerFeeShare is a free data retrieval call binding the contract method 0xe43a4954.
//
// Solidity: function managerFeeShare() view returns(uint256)
func (_RewardpoolContract *RewardpoolContractCaller) ManagerFeeShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardpoolContract.contract.Call(opts, &out, "managerFeeShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ManagerFeeShare is a free data retrieval call binding the contract method 0xe43a4954.
//
// Solidity: function managerFeeShare() view returns(uint256)
func (_RewardpoolContract *RewardpoolContractSession) ManagerFeeShare() (*big.Int, error) {
	return _RewardpoolContract.Contract.ManagerFeeShare(&_RewardpoolContract.CallOpts)
}

// ManagerFeeShare is a free data retrieval call binding the contract method 0xe43a4954.
//
// Solidity: function managerFeeShare() view returns(uint256)
func (_RewardpoolContract *RewardpoolContractCallerSession) ManagerFeeShare() (*big.Int, error) {
	return _RewardpoolContract.Contract.ManagerFeeShare(&_RewardpoolContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RewardpoolContract *RewardpoolContractCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RewardpoolContract.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RewardpoolContract *RewardpoolContractSession) Paused() (bool, error) {
	return _RewardpoolContract.Contract.Paused(&_RewardpoolContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RewardpoolContract *RewardpoolContractCallerSession) Paused() (bool, error) {
	return _RewardpoolContract.Contract.Paused(&_RewardpoolContract.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RewardpoolContract *RewardpoolContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _RewardpoolContract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RewardpoolContract *RewardpoolContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RewardpoolContract.Contract.SupportsInterface(&_RewardpoolContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RewardpoolContract *RewardpoolContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RewardpoolContract.Contract.SupportsInterface(&_RewardpoolContract.CallOpts, interfaceId)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 accSharePoint, uint256 amount, uint256 rewardBalance)
func (_RewardpoolContract *RewardpoolContractCaller) UserInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	AccSharePoint *big.Int
	Amount        *big.Int
	RewardBalance *big.Int
}, error) {
	var out []interface{}
	err := _RewardpoolContract.contract.Call(opts, &out, "userInfo", arg0)

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
// Solidity: function userInfo(address ) view returns(uint256 accSharePoint, uint256 amount, uint256 rewardBalance)
func (_RewardpoolContract *RewardpoolContractSession) UserInfo(arg0 common.Address) (struct {
	AccSharePoint *big.Int
	Amount        *big.Int
	RewardBalance *big.Int
}, error) {
	return _RewardpoolContract.Contract.UserInfo(&_RewardpoolContract.CallOpts, arg0)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 accSharePoint, uint256 amount, uint256 rewardBalance)
func (_RewardpoolContract *RewardpoolContractCallerSession) UserInfo(arg0 common.Address) (struct {
	AccSharePoint *big.Int
	Amount        *big.Int
	RewardBalance *big.Int
}, error) {
	return _RewardpoolContract.Contract.UserInfo(&_RewardpoolContract.CallOpts, arg0)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x9a99b4f0.
//
// Solidity: function claimRewards(address beneficiary, uint256 amount) returns()
func (_RewardpoolContract *RewardpoolContractTransactor) ClaimRewards(opts *bind.TransactOpts, beneficiary common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RewardpoolContract.contract.Transact(opts, "claimRewards", beneficiary, amount)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x9a99b4f0.
//
// Solidity: function claimRewards(address beneficiary, uint256 amount) returns()
func (_RewardpoolContract *RewardpoolContractSession) ClaimRewards(beneficiary common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.ClaimRewards(&_RewardpoolContract.TransactOpts, beneficiary, amount)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x9a99b4f0.
//
// Solidity: function claimRewards(address beneficiary, uint256 amount) returns()
func (_RewardpoolContract *RewardpoolContractTransactorSession) ClaimRewards(beneficiary common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.ClaimRewards(&_RewardpoolContract.TransactOpts, beneficiary, amount)
}

// ClaimRewardsFor is a paid mutator transaction binding the contract method 0x1ac6d19d.
//
// Solidity: function claimRewardsFor(address account) returns()
func (_RewardpoolContract *RewardpoolContractTransactor) ClaimRewardsFor(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _RewardpoolContract.contract.Transact(opts, "claimRewardsFor", account)
}

// ClaimRewardsFor is a paid mutator transaction binding the contract method 0x1ac6d19d.
//
// Solidity: function claimRewardsFor(address account) returns()
func (_RewardpoolContract *RewardpoolContractSession) ClaimRewardsFor(account common.Address) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.ClaimRewardsFor(&_RewardpoolContract.TransactOpts, account)
}

// ClaimRewardsFor is a paid mutator transaction binding the contract method 0x1ac6d19d.
//
// Solidity: function claimRewardsFor(address account) returns()
func (_RewardpoolContract *RewardpoolContractTransactorSession) ClaimRewardsFor(account common.Address) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.ClaimRewardsFor(&_RewardpoolContract.TransactOpts, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RewardpoolContract *RewardpoolContractTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewardpoolContract.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RewardpoolContract *RewardpoolContractSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.GrantRole(&_RewardpoolContract.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RewardpoolContract *RewardpoolContractTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.GrantRole(&_RewardpoolContract.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_RewardpoolContract *RewardpoolContractTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardpoolContract.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_RewardpoolContract *RewardpoolContractSession) Initialize() (*types.Transaction, error) {
	return _RewardpoolContract.Contract.Initialize(&_RewardpoolContract.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_RewardpoolContract *RewardpoolContractTransactorSession) Initialize() (*types.Transaction, error) {
	return _RewardpoolContract.Contract.Initialize(&_RewardpoolContract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RewardpoolContract *RewardpoolContractTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardpoolContract.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RewardpoolContract *RewardpoolContractSession) Pause() (*types.Transaction, error) {
	return _RewardpoolContract.Contract.Pause(&_RewardpoolContract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RewardpoolContract *RewardpoolContractTransactorSession) Pause() (*types.Transaction, error) {
	return _RewardpoolContract.Contract.Pause(&_RewardpoolContract.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RewardpoolContract *RewardpoolContractTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewardpoolContract.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RewardpoolContract *RewardpoolContractSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.RenounceRole(&_RewardpoolContract.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RewardpoolContract *RewardpoolContractTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.RenounceRole(&_RewardpoolContract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RewardpoolContract *RewardpoolContractTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewardpoolContract.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RewardpoolContract *RewardpoolContractSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.RevokeRole(&_RewardpoolContract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RewardpoolContract *RewardpoolContractTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.RevokeRole(&_RewardpoolContract.TransactOpts, role, account)
}

// SetManagerFeeShare is a paid mutator transaction binding the contract method 0x755d7dd3.
//
// Solidity: function setManagerFeeShare(uint256 milli) returns()
func (_RewardpoolContract *RewardpoolContractTransactor) SetManagerFeeShare(opts *bind.TransactOpts, milli *big.Int) (*types.Transaction, error) {
	return _RewardpoolContract.contract.Transact(opts, "setManagerFeeShare", milli)
}

// SetManagerFeeShare is a paid mutator transaction binding the contract method 0x755d7dd3.
//
// Solidity: function setManagerFeeShare(uint256 milli) returns()
func (_RewardpoolContract *RewardpoolContractSession) SetManagerFeeShare(milli *big.Int) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.SetManagerFeeShare(&_RewardpoolContract.TransactOpts, milli)
}

// SetManagerFeeShare is a paid mutator transaction binding the contract method 0x755d7dd3.
//
// Solidity: function setManagerFeeShare(uint256 milli) returns()
func (_RewardpoolContract *RewardpoolContractTransactorSession) SetManagerFeeShare(milli *big.Int) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.SetManagerFeeShare(&_RewardpoolContract.TransactOpts, milli)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RewardpoolContract *RewardpoolContractTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardpoolContract.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RewardpoolContract *RewardpoolContractSession) Unpause() (*types.Transaction, error) {
	return _RewardpoolContract.Contract.Unpause(&_RewardpoolContract.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RewardpoolContract *RewardpoolContractTransactorSession) Unpause() (*types.Transaction, error) {
	return _RewardpoolContract.Contract.Unpause(&_RewardpoolContract.TransactOpts)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x2a17c7e5.
//
// Solidity: function updatePool(address[] claimaddrs, int256[] deltas) returns()
func (_RewardpoolContract *RewardpoolContractTransactor) UpdatePool(opts *bind.TransactOpts, claimaddrs []common.Address, deltas []*big.Int) (*types.Transaction, error) {
	return _RewardpoolContract.contract.Transact(opts, "updatePool", claimaddrs, deltas)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x2a17c7e5.
//
// Solidity: function updatePool(address[] claimaddrs, int256[] deltas) returns()
func (_RewardpoolContract *RewardpoolContractSession) UpdatePool(claimaddrs []common.Address, deltas []*big.Int) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.UpdatePool(&_RewardpoolContract.TransactOpts, claimaddrs, deltas)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x2a17c7e5.
//
// Solidity: function updatePool(address[] claimaddrs, int256[] deltas) returns()
func (_RewardpoolContract *RewardpoolContractTransactorSession) UpdatePool(claimaddrs []common.Address, deltas []*big.Int) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.UpdatePool(&_RewardpoolContract.TransactOpts, claimaddrs, deltas)
}

// WithdrawManagerRevenue is a paid mutator transaction binding the contract method 0xac9a5f56.
//
// Solidity: function withdrawManagerRevenue(uint256 amount, address to) returns()
func (_RewardpoolContract *RewardpoolContractTransactor) WithdrawManagerRevenue(opts *bind.TransactOpts, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _RewardpoolContract.contract.Transact(opts, "withdrawManagerRevenue", amount, to)
}

// WithdrawManagerRevenue is a paid mutator transaction binding the contract method 0xac9a5f56.
//
// Solidity: function withdrawManagerRevenue(uint256 amount, address to) returns()
func (_RewardpoolContract *RewardpoolContractSession) WithdrawManagerRevenue(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.WithdrawManagerRevenue(&_RewardpoolContract.TransactOpts, amount, to)
}

// WithdrawManagerRevenue is a paid mutator transaction binding the contract method 0xac9a5f56.
//
// Solidity: function withdrawManagerRevenue(uint256 amount, address to) returns()
func (_RewardpoolContract *RewardpoolContractTransactorSession) WithdrawManagerRevenue(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.WithdrawManagerRevenue(&_RewardpoolContract.TransactOpts, amount, to)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RewardpoolContract *RewardpoolContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardpoolContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RewardpoolContract *RewardpoolContractSession) Receive() (*types.Transaction, error) {
	return _RewardpoolContract.Contract.Receive(&_RewardpoolContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RewardpoolContract *RewardpoolContractTransactorSession) Receive() (*types.Transaction, error) {
	return _RewardpoolContract.Contract.Receive(&_RewardpoolContract.TransactOpts)
}

// RewardpoolContractClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the RewardpoolContract contract.
type RewardpoolContractClaimedIterator struct {
	Event *RewardpoolContractClaimed // Event containing the contract specifics and raw log

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
func (it *RewardpoolContractClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardpoolContractClaimed)
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
		it.Event = new(RewardpoolContractClaimed)
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
func (it *RewardpoolContractClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardpoolContractClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardpoolContractClaimed represents a Claimed event raised by the RewardpoolContract contract.
type RewardpoolContractClaimed struct {
	Beneficiary common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address beneficiary, uint256 amount)
func (_RewardpoolContract *RewardpoolContractFilterer) FilterClaimed(opts *bind.FilterOpts) (*RewardpoolContractClaimedIterator, error) {

	logs, sub, err := _RewardpoolContract.contract.FilterLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return &RewardpoolContractClaimedIterator{contract: _RewardpoolContract.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address beneficiary, uint256 amount)
func (_RewardpoolContract *RewardpoolContractFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *RewardpoolContractClaimed) (event.Subscription, error) {

	logs, sub, err := _RewardpoolContract.contract.WatchLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardpoolContractClaimed)
				if err := _RewardpoolContract.contract.UnpackLog(event, "Claimed", log); err != nil {
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
func (_RewardpoolContract *RewardpoolContractFilterer) ParseClaimed(log types.Log) (*RewardpoolContractClaimed, error) {
	event := new(RewardpoolContractClaimed)
	if err := _RewardpoolContract.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardpoolContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the RewardpoolContract contract.
type RewardpoolContractInitializedIterator struct {
	Event *RewardpoolContractInitialized // Event containing the contract specifics and raw log

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
func (it *RewardpoolContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardpoolContractInitialized)
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
		it.Event = new(RewardpoolContractInitialized)
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
func (it *RewardpoolContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardpoolContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardpoolContractInitialized represents a Initialized event raised by the RewardpoolContract contract.
type RewardpoolContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RewardpoolContract *RewardpoolContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*RewardpoolContractInitializedIterator, error) {

	logs, sub, err := _RewardpoolContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RewardpoolContractInitializedIterator{contract: _RewardpoolContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RewardpoolContract *RewardpoolContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RewardpoolContractInitialized) (event.Subscription, error) {

	logs, sub, err := _RewardpoolContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardpoolContractInitialized)
				if err := _RewardpoolContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_RewardpoolContract *RewardpoolContractFilterer) ParseInitialized(log types.Log) (*RewardpoolContractInitialized, error) {
	event := new(RewardpoolContractInitialized)
	if err := _RewardpoolContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardpoolContractManagerFeeSetIterator is returned from FilterManagerFeeSet and is used to iterate over the raw logs and unpacked data for ManagerFeeSet events raised by the RewardpoolContract contract.
type RewardpoolContractManagerFeeSetIterator struct {
	Event *RewardpoolContractManagerFeeSet // Event containing the contract specifics and raw log

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
func (it *RewardpoolContractManagerFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardpoolContractManagerFeeSet)
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
		it.Event = new(RewardpoolContractManagerFeeSet)
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
func (it *RewardpoolContractManagerFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardpoolContractManagerFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardpoolContractManagerFeeSet represents a ManagerFeeSet event raised by the RewardpoolContract contract.
type RewardpoolContractManagerFeeSet struct {
	Milli *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterManagerFeeSet is a free log retrieval operation binding the contract event 0x4de90ec86e1bc56c192e2399bacbd10bdaba720caca606354d66c5cb33d6802b.
//
// Solidity: event ManagerFeeSet(uint256 milli)
func (_RewardpoolContract *RewardpoolContractFilterer) FilterManagerFeeSet(opts *bind.FilterOpts) (*RewardpoolContractManagerFeeSetIterator, error) {

	logs, sub, err := _RewardpoolContract.contract.FilterLogs(opts, "ManagerFeeSet")
	if err != nil {
		return nil, err
	}
	return &RewardpoolContractManagerFeeSetIterator{contract: _RewardpoolContract.contract, event: "ManagerFeeSet", logs: logs, sub: sub}, nil
}

// WatchManagerFeeSet is a free log subscription operation binding the contract event 0x4de90ec86e1bc56c192e2399bacbd10bdaba720caca606354d66c5cb33d6802b.
//
// Solidity: event ManagerFeeSet(uint256 milli)
func (_RewardpoolContract *RewardpoolContractFilterer) WatchManagerFeeSet(opts *bind.WatchOpts, sink chan<- *RewardpoolContractManagerFeeSet) (event.Subscription, error) {

	logs, sub, err := _RewardpoolContract.contract.WatchLogs(opts, "ManagerFeeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardpoolContractManagerFeeSet)
				if err := _RewardpoolContract.contract.UnpackLog(event, "ManagerFeeSet", log); err != nil {
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
func (_RewardpoolContract *RewardpoolContractFilterer) ParseManagerFeeSet(log types.Log) (*RewardpoolContractManagerFeeSet, error) {
	event := new(RewardpoolContractManagerFeeSet)
	if err := _RewardpoolContract.contract.UnpackLog(event, "ManagerFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardpoolContractManagerFeeWithdrawedIterator is returned from FilterManagerFeeWithdrawed and is used to iterate over the raw logs and unpacked data for ManagerFeeWithdrawed events raised by the RewardpoolContract contract.
type RewardpoolContractManagerFeeWithdrawedIterator struct {
	Event *RewardpoolContractManagerFeeWithdrawed // Event containing the contract specifics and raw log

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
func (it *RewardpoolContractManagerFeeWithdrawedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardpoolContractManagerFeeWithdrawed)
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
		it.Event = new(RewardpoolContractManagerFeeWithdrawed)
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
func (it *RewardpoolContractManagerFeeWithdrawedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardpoolContractManagerFeeWithdrawedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardpoolContractManagerFeeWithdrawed represents a ManagerFeeWithdrawed event raised by the RewardpoolContract contract.
type RewardpoolContractManagerFeeWithdrawed struct {
	Amount *big.Int
	To     common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterManagerFeeWithdrawed is a free log retrieval operation binding the contract event 0x2425aa1fadefc5c570850aa9c9e3dfa4fc6b43ccd1c05b47db38dd6518a743b3.
//
// Solidity: event ManagerFeeWithdrawed(uint256 amount, address to)
func (_RewardpoolContract *RewardpoolContractFilterer) FilterManagerFeeWithdrawed(opts *bind.FilterOpts) (*RewardpoolContractManagerFeeWithdrawedIterator, error) {

	logs, sub, err := _RewardpoolContract.contract.FilterLogs(opts, "ManagerFeeWithdrawed")
	if err != nil {
		return nil, err
	}
	return &RewardpoolContractManagerFeeWithdrawedIterator{contract: _RewardpoolContract.contract, event: "ManagerFeeWithdrawed", logs: logs, sub: sub}, nil
}

// WatchManagerFeeWithdrawed is a free log subscription operation binding the contract event 0x2425aa1fadefc5c570850aa9c9e3dfa4fc6b43ccd1c05b47db38dd6518a743b3.
//
// Solidity: event ManagerFeeWithdrawed(uint256 amount, address to)
func (_RewardpoolContract *RewardpoolContractFilterer) WatchManagerFeeWithdrawed(opts *bind.WatchOpts, sink chan<- *RewardpoolContractManagerFeeWithdrawed) (event.Subscription, error) {

	logs, sub, err := _RewardpoolContract.contract.WatchLogs(opts, "ManagerFeeWithdrawed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardpoolContractManagerFeeWithdrawed)
				if err := _RewardpoolContract.contract.UnpackLog(event, "ManagerFeeWithdrawed", log); err != nil {
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
func (_RewardpoolContract *RewardpoolContractFilterer) ParseManagerFeeWithdrawed(log types.Log) (*RewardpoolContractManagerFeeWithdrawed, error) {
	event := new(RewardpoolContractManagerFeeWithdrawed)
	if err := _RewardpoolContract.contract.UnpackLog(event, "ManagerFeeWithdrawed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardpoolContractPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the RewardpoolContract contract.
type RewardpoolContractPausedIterator struct {
	Event *RewardpoolContractPaused // Event containing the contract specifics and raw log

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
func (it *RewardpoolContractPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardpoolContractPaused)
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
		it.Event = new(RewardpoolContractPaused)
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
func (it *RewardpoolContractPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardpoolContractPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardpoolContractPaused represents a Paused event raised by the RewardpoolContract contract.
type RewardpoolContractPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_RewardpoolContract *RewardpoolContractFilterer) FilterPaused(opts *bind.FilterOpts) (*RewardpoolContractPausedIterator, error) {

	logs, sub, err := _RewardpoolContract.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &RewardpoolContractPausedIterator{contract: _RewardpoolContract.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_RewardpoolContract *RewardpoolContractFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *RewardpoolContractPaused) (event.Subscription, error) {

	logs, sub, err := _RewardpoolContract.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardpoolContractPaused)
				if err := _RewardpoolContract.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_RewardpoolContract *RewardpoolContractFilterer) ParsePaused(log types.Log) (*RewardpoolContractPaused, error) {
	event := new(RewardpoolContractPaused)
	if err := _RewardpoolContract.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardpoolContractPoolUpdateIterator is returned from FilterPoolUpdate and is used to iterate over the raw logs and unpacked data for PoolUpdate events raised by the RewardpoolContract contract.
type RewardpoolContractPoolUpdateIterator struct {
	Event *RewardpoolContractPoolUpdate // Event containing the contract specifics and raw log

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
func (it *RewardpoolContractPoolUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardpoolContractPoolUpdate)
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
		it.Event = new(RewardpoolContractPoolUpdate)
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
func (it *RewardpoolContractPoolUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardpoolContractPoolUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardpoolContractPoolUpdate represents a PoolUpdate event raised by the RewardpoolContract contract.
type RewardpoolContractPoolUpdate struct {
	Claimaddr []common.Address
	Amount    []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPoolUpdate is a free log retrieval operation binding the contract event 0x2d12baed952e32deda21166272a112c0347541bb1ba77c3d58389be6ec3bfcff.
//
// Solidity: event PoolUpdate(address[] claimaddr, int256[] amount)
func (_RewardpoolContract *RewardpoolContractFilterer) FilterPoolUpdate(opts *bind.FilterOpts) (*RewardpoolContractPoolUpdateIterator, error) {

	logs, sub, err := _RewardpoolContract.contract.FilterLogs(opts, "PoolUpdate")
	if err != nil {
		return nil, err
	}
	return &RewardpoolContractPoolUpdateIterator{contract: _RewardpoolContract.contract, event: "PoolUpdate", logs: logs, sub: sub}, nil
}

// WatchPoolUpdate is a free log subscription operation binding the contract event 0x2d12baed952e32deda21166272a112c0347541bb1ba77c3d58389be6ec3bfcff.
//
// Solidity: event PoolUpdate(address[] claimaddr, int256[] amount)
func (_RewardpoolContract *RewardpoolContractFilterer) WatchPoolUpdate(opts *bind.WatchOpts, sink chan<- *RewardpoolContractPoolUpdate) (event.Subscription, error) {

	logs, sub, err := _RewardpoolContract.contract.WatchLogs(opts, "PoolUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardpoolContractPoolUpdate)
				if err := _RewardpoolContract.contract.UnpackLog(event, "PoolUpdate", log); err != nil {
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

// ParsePoolUpdate is a log parse operation binding the contract event 0x2d12baed952e32deda21166272a112c0347541bb1ba77c3d58389be6ec3bfcff.
//
// Solidity: event PoolUpdate(address[] claimaddr, int256[] amount)
func (_RewardpoolContract *RewardpoolContractFilterer) ParsePoolUpdate(log types.Log) (*RewardpoolContractPoolUpdate, error) {
	event := new(RewardpoolContractPoolUpdate)
	if err := _RewardpoolContract.contract.UnpackLog(event, "PoolUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardpoolContractRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the RewardpoolContract contract.
type RewardpoolContractRoleAdminChangedIterator struct {
	Event *RewardpoolContractRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *RewardpoolContractRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardpoolContractRoleAdminChanged)
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
		it.Event = new(RewardpoolContractRoleAdminChanged)
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
func (it *RewardpoolContractRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardpoolContractRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardpoolContractRoleAdminChanged represents a RoleAdminChanged event raised by the RewardpoolContract contract.
type RewardpoolContractRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RewardpoolContract *RewardpoolContractFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*RewardpoolContractRoleAdminChangedIterator, error) {

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

	logs, sub, err := _RewardpoolContract.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &RewardpoolContractRoleAdminChangedIterator{contract: _RewardpoolContract.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RewardpoolContract *RewardpoolContractFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *RewardpoolContractRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _RewardpoolContract.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardpoolContractRoleAdminChanged)
				if err := _RewardpoolContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_RewardpoolContract *RewardpoolContractFilterer) ParseRoleAdminChanged(log types.Log) (*RewardpoolContractRoleAdminChanged, error) {
	event := new(RewardpoolContractRoleAdminChanged)
	if err := _RewardpoolContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardpoolContractRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the RewardpoolContract contract.
type RewardpoolContractRoleGrantedIterator struct {
	Event *RewardpoolContractRoleGranted // Event containing the contract specifics and raw log

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
func (it *RewardpoolContractRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardpoolContractRoleGranted)
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
		it.Event = new(RewardpoolContractRoleGranted)
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
func (it *RewardpoolContractRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardpoolContractRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardpoolContractRoleGranted represents a RoleGranted event raised by the RewardpoolContract contract.
type RewardpoolContractRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RewardpoolContract *RewardpoolContractFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RewardpoolContractRoleGrantedIterator, error) {

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

	logs, sub, err := _RewardpoolContract.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RewardpoolContractRoleGrantedIterator{contract: _RewardpoolContract.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RewardpoolContract *RewardpoolContractFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RewardpoolContractRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _RewardpoolContract.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardpoolContractRoleGranted)
				if err := _RewardpoolContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_RewardpoolContract *RewardpoolContractFilterer) ParseRoleGranted(log types.Log) (*RewardpoolContractRoleGranted, error) {
	event := new(RewardpoolContractRoleGranted)
	if err := _RewardpoolContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardpoolContractRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the RewardpoolContract contract.
type RewardpoolContractRoleRevokedIterator struct {
	Event *RewardpoolContractRoleRevoked // Event containing the contract specifics and raw log

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
func (it *RewardpoolContractRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardpoolContractRoleRevoked)
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
		it.Event = new(RewardpoolContractRoleRevoked)
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
func (it *RewardpoolContractRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardpoolContractRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardpoolContractRoleRevoked represents a RoleRevoked event raised by the RewardpoolContract contract.
type RewardpoolContractRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RewardpoolContract *RewardpoolContractFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RewardpoolContractRoleRevokedIterator, error) {

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

	logs, sub, err := _RewardpoolContract.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RewardpoolContractRoleRevokedIterator{contract: _RewardpoolContract.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RewardpoolContract *RewardpoolContractFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RewardpoolContractRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _RewardpoolContract.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardpoolContractRoleRevoked)
				if err := _RewardpoolContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_RewardpoolContract *RewardpoolContractFilterer) ParseRoleRevoked(log types.Log) (*RewardpoolContractRoleRevoked, error) {
	event := new(RewardpoolContractRoleRevoked)
	if err := _RewardpoolContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardpoolContractUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the RewardpoolContract contract.
type RewardpoolContractUnpausedIterator struct {
	Event *RewardpoolContractUnpaused // Event containing the contract specifics and raw log

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
func (it *RewardpoolContractUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardpoolContractUnpaused)
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
		it.Event = new(RewardpoolContractUnpaused)
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
func (it *RewardpoolContractUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardpoolContractUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardpoolContractUnpaused represents a Unpaused event raised by the RewardpoolContract contract.
type RewardpoolContractUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_RewardpoolContract *RewardpoolContractFilterer) FilterUnpaused(opts *bind.FilterOpts) (*RewardpoolContractUnpausedIterator, error) {

	logs, sub, err := _RewardpoolContract.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &RewardpoolContractUnpausedIterator{contract: _RewardpoolContract.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_RewardpoolContract *RewardpoolContractFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *RewardpoolContractUnpaused) (event.Subscription, error) {

	logs, sub, err := _RewardpoolContract.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardpoolContractUnpaused)
				if err := _RewardpoolContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_RewardpoolContract *RewardpoolContractFilterer) ParseUnpaused(log types.Log) (*RewardpoolContractUnpaused, error) {
	event := new(RewardpoolContractUnpaused)
	if err := _RewardpoolContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
