// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staking_contract

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

// StakingContractMetaData contains all meta data concerning the StakingContract contract.
var StakingContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"DepositContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"RewardPoolContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEPOSIT_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REGISTRY_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"batchDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethDepositContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"}],\"name\":\"getExitQueue\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExitQueueLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextValidatorToBind\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextValidatorToDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextValidatorToRegister\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegisteredValidatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getValidatorInfo\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"withdrawalAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"claimAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"userid\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"registerValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"}],\"name\":\"registerValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"replaceValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethDepositContract\",\"type\":\"address\"}],\"name\":\"setETHDepositContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardPool\",\"type\":\"address\"}],\"name\":\"setRewardPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"withdrawaddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"claimaddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"extradata\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// StakingContractABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingContractMetaData.ABI instead.
var StakingContractABI = StakingContractMetaData.ABI

// StakingContract is an auto generated Go binding around an Ethereum contract.
type StakingContract struct {
	StakingContractCaller     // Read-only binding to the contract
	StakingContractTransactor // Write-only binding to the contract
	StakingContractFilterer   // Log filterer for contract events
}

// StakingContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingContractSession struct {
	Contract     *StakingContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingContractCallerSession struct {
	Contract *StakingContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// StakingContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingContractTransactorSession struct {
	Contract     *StakingContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// StakingContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingContractRaw struct {
	Contract *StakingContract // Generic contract binding to access the raw methods on
}

// StakingContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingContractCallerRaw struct {
	Contract *StakingContractCaller // Generic read-only contract binding to access the raw methods on
}

// StakingContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingContractTransactorRaw struct {
	Contract *StakingContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingContract creates a new instance of StakingContract, bound to a specific deployed contract.
func NewStakingContract(address common.Address, backend bind.ContractBackend) (*StakingContract, error) {
	contract, err := bindStakingContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingContract{StakingContractCaller: StakingContractCaller{contract: contract}, StakingContractTransactor: StakingContractTransactor{contract: contract}, StakingContractFilterer: StakingContractFilterer{contract: contract}}, nil
}

// NewStakingContractCaller creates a new read-only instance of StakingContract, bound to a specific deployed contract.
func NewStakingContractCaller(address common.Address, caller bind.ContractCaller) (*StakingContractCaller, error) {
	contract, err := bindStakingContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingContractCaller{contract: contract}, nil
}

// NewStakingContractTransactor creates a new write-only instance of StakingContract, bound to a specific deployed contract.
func NewStakingContractTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingContractTransactor, error) {
	contract, err := bindStakingContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingContractTransactor{contract: contract}, nil
}

// NewStakingContractFilterer creates a new log filterer instance of StakingContract, bound to a specific deployed contract.
func NewStakingContractFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingContractFilterer, error) {
	contract, err := bindStakingContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingContractFilterer{contract: contract}, nil
}

// bindStakingContract binds a generic wrapper to an already deployed contract.
func bindStakingContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingContract *StakingContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingContract.Contract.StakingContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingContract *StakingContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingContract.Contract.StakingContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingContract *StakingContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingContract.Contract.StakingContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingContract *StakingContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingContract *StakingContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingContract *StakingContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingContract.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakingContract *StakingContractCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakingContract *StakingContractSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StakingContract.Contract.DEFAULTADMINROLE(&_StakingContract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakingContract *StakingContractCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StakingContract.Contract.DEFAULTADMINROLE(&_StakingContract.CallOpts)
}

// DEPOSITSIZE is a free data retrieval call binding the contract method 0x36bf3325.
//
// Solidity: function DEPOSIT_SIZE() view returns(uint256)
func (_StakingContract *StakingContractCaller) DEPOSITSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "DEPOSIT_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEPOSITSIZE is a free data retrieval call binding the contract method 0x36bf3325.
//
// Solidity: function DEPOSIT_SIZE() view returns(uint256)
func (_StakingContract *StakingContractSession) DEPOSITSIZE() (*big.Int, error) {
	return _StakingContract.Contract.DEPOSITSIZE(&_StakingContract.CallOpts)
}

// DEPOSITSIZE is a free data retrieval call binding the contract method 0x36bf3325.
//
// Solidity: function DEPOSIT_SIZE() view returns(uint256)
func (_StakingContract *StakingContractCallerSession) DEPOSITSIZE() (*big.Int, error) {
	return _StakingContract.Contract.DEPOSITSIZE(&_StakingContract.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_StakingContract *StakingContractCaller) MANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_StakingContract *StakingContractSession) MANAGERROLE() ([32]byte, error) {
	return _StakingContract.Contract.MANAGERROLE(&_StakingContract.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_StakingContract *StakingContractCallerSession) MANAGERROLE() ([32]byte, error) {
	return _StakingContract.Contract.MANAGERROLE(&_StakingContract.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_StakingContract *StakingContractCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_StakingContract *StakingContractSession) PAUSERROLE() ([32]byte, error) {
	return _StakingContract.Contract.PAUSERROLE(&_StakingContract.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_StakingContract *StakingContractCallerSession) PAUSERROLE() ([32]byte, error) {
	return _StakingContract.Contract.PAUSERROLE(&_StakingContract.CallOpts)
}

// REGISTRYROLE is a free data retrieval call binding the contract method 0x42f1e879.
//
// Solidity: function REGISTRY_ROLE() view returns(bytes32)
func (_StakingContract *StakingContractCaller) REGISTRYROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "REGISTRY_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REGISTRYROLE is a free data retrieval call binding the contract method 0x42f1e879.
//
// Solidity: function REGISTRY_ROLE() view returns(bytes32)
func (_StakingContract *StakingContractSession) REGISTRYROLE() ([32]byte, error) {
	return _StakingContract.Contract.REGISTRYROLE(&_StakingContract.CallOpts)
}

// REGISTRYROLE is a free data retrieval call binding the contract method 0x42f1e879.
//
// Solidity: function REGISTRY_ROLE() view returns(bytes32)
func (_StakingContract *StakingContractCallerSession) REGISTRYROLE() ([32]byte, error) {
	return _StakingContract.Contract.REGISTRYROLE(&_StakingContract.CallOpts)
}

// EthDepositContract is a free data retrieval call binding the contract method 0x3884545d.
//
// Solidity: function ethDepositContract() view returns(address)
func (_StakingContract *StakingContractCaller) EthDepositContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "ethDepositContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthDepositContract is a free data retrieval call binding the contract method 0x3884545d.
//
// Solidity: function ethDepositContract() view returns(address)
func (_StakingContract *StakingContractSession) EthDepositContract() (common.Address, error) {
	return _StakingContract.Contract.EthDepositContract(&_StakingContract.CallOpts)
}

// EthDepositContract is a free data retrieval call binding the contract method 0x3884545d.
//
// Solidity: function ethDepositContract() view returns(address)
func (_StakingContract *StakingContractCallerSession) EthDepositContract() (common.Address, error) {
	return _StakingContract.Contract.EthDepositContract(&_StakingContract.CallOpts)
}

// GetExitQueue is a free data retrieval call binding the contract method 0x43931256.
//
// Solidity: function getExitQueue(uint256 from, uint256 to) view returns(uint256[])
func (_StakingContract *StakingContractCaller) GetExitQueue(opts *bind.CallOpts, from *big.Int, to *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "getExitQueue", from, to)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetExitQueue is a free data retrieval call binding the contract method 0x43931256.
//
// Solidity: function getExitQueue(uint256 from, uint256 to) view returns(uint256[])
func (_StakingContract *StakingContractSession) GetExitQueue(from *big.Int, to *big.Int) ([]*big.Int, error) {
	return _StakingContract.Contract.GetExitQueue(&_StakingContract.CallOpts, from, to)
}

// GetExitQueue is a free data retrieval call binding the contract method 0x43931256.
//
// Solidity: function getExitQueue(uint256 from, uint256 to) view returns(uint256[])
func (_StakingContract *StakingContractCallerSession) GetExitQueue(from *big.Int, to *big.Int) ([]*big.Int, error) {
	return _StakingContract.Contract.GetExitQueue(&_StakingContract.CallOpts, from, to)
}

// GetExitQueueLength is a free data retrieval call binding the contract method 0x559230a0.
//
// Solidity: function getExitQueueLength() view returns(uint256)
func (_StakingContract *StakingContractCaller) GetExitQueueLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "getExitQueueLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExitQueueLength is a free data retrieval call binding the contract method 0x559230a0.
//
// Solidity: function getExitQueueLength() view returns(uint256)
func (_StakingContract *StakingContractSession) GetExitQueueLength() (*big.Int, error) {
	return _StakingContract.Contract.GetExitQueueLength(&_StakingContract.CallOpts)
}

// GetExitQueueLength is a free data retrieval call binding the contract method 0x559230a0.
//
// Solidity: function getExitQueueLength() view returns(uint256)
func (_StakingContract *StakingContractCallerSession) GetExitQueueLength() (*big.Int, error) {
	return _StakingContract.Contract.GetExitQueueLength(&_StakingContract.CallOpts)
}

// GetNextValidatorToBind is a free data retrieval call binding the contract method 0x75fe2283.
//
// Solidity: function getNextValidatorToBind() view returns(uint256)
func (_StakingContract *StakingContractCaller) GetNextValidatorToBind(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "getNextValidatorToBind")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextValidatorToBind is a free data retrieval call binding the contract method 0x75fe2283.
//
// Solidity: function getNextValidatorToBind() view returns(uint256)
func (_StakingContract *StakingContractSession) GetNextValidatorToBind() (*big.Int, error) {
	return _StakingContract.Contract.GetNextValidatorToBind(&_StakingContract.CallOpts)
}

// GetNextValidatorToBind is a free data retrieval call binding the contract method 0x75fe2283.
//
// Solidity: function getNextValidatorToBind() view returns(uint256)
func (_StakingContract *StakingContractCallerSession) GetNextValidatorToBind() (*big.Int, error) {
	return _StakingContract.Contract.GetNextValidatorToBind(&_StakingContract.CallOpts)
}

// GetNextValidatorToDeposit is a free data retrieval call binding the contract method 0xa33eabfa.
//
// Solidity: function getNextValidatorToDeposit() view returns(uint256)
func (_StakingContract *StakingContractCaller) GetNextValidatorToDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "getNextValidatorToDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextValidatorToDeposit is a free data retrieval call binding the contract method 0xa33eabfa.
//
// Solidity: function getNextValidatorToDeposit() view returns(uint256)
func (_StakingContract *StakingContractSession) GetNextValidatorToDeposit() (*big.Int, error) {
	return _StakingContract.Contract.GetNextValidatorToDeposit(&_StakingContract.CallOpts)
}

// GetNextValidatorToDeposit is a free data retrieval call binding the contract method 0xa33eabfa.
//
// Solidity: function getNextValidatorToDeposit() view returns(uint256)
func (_StakingContract *StakingContractCallerSession) GetNextValidatorToDeposit() (*big.Int, error) {
	return _StakingContract.Contract.GetNextValidatorToDeposit(&_StakingContract.CallOpts)
}

// GetNextValidatorToRegister is a free data retrieval call binding the contract method 0x9a24db24.
//
// Solidity: function getNextValidatorToRegister() view returns(uint256)
func (_StakingContract *StakingContractCaller) GetNextValidatorToRegister(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "getNextValidatorToRegister")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextValidatorToRegister is a free data retrieval call binding the contract method 0x9a24db24.
//
// Solidity: function getNextValidatorToRegister() view returns(uint256)
func (_StakingContract *StakingContractSession) GetNextValidatorToRegister() (*big.Int, error) {
	return _StakingContract.Contract.GetNextValidatorToRegister(&_StakingContract.CallOpts)
}

// GetNextValidatorToRegister is a free data retrieval call binding the contract method 0x9a24db24.
//
// Solidity: function getNextValidatorToRegister() view returns(uint256)
func (_StakingContract *StakingContractCallerSession) GetNextValidatorToRegister() (*big.Int, error) {
	return _StakingContract.Contract.GetNextValidatorToRegister(&_StakingContract.CallOpts)
}

// GetRegisteredValidatorsCount is a free data retrieval call binding the contract method 0x30b12c8d.
//
// Solidity: function getRegisteredValidatorsCount() view returns(uint256)
func (_StakingContract *StakingContractCaller) GetRegisteredValidatorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "getRegisteredValidatorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRegisteredValidatorsCount is a free data retrieval call binding the contract method 0x30b12c8d.
//
// Solidity: function getRegisteredValidatorsCount() view returns(uint256)
func (_StakingContract *StakingContractSession) GetRegisteredValidatorsCount() (*big.Int, error) {
	return _StakingContract.Contract.GetRegisteredValidatorsCount(&_StakingContract.CallOpts)
}

// GetRegisteredValidatorsCount is a free data retrieval call binding the contract method 0x30b12c8d.
//
// Solidity: function getRegisteredValidatorsCount() view returns(uint256)
func (_StakingContract *StakingContractCallerSession) GetRegisteredValidatorsCount() (*big.Int, error) {
	return _StakingContract.Contract.GetRegisteredValidatorsCount(&_StakingContract.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakingContract *StakingContractCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakingContract *StakingContractSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StakingContract.Contract.GetRoleAdmin(&_StakingContract.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakingContract *StakingContractCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StakingContract.Contract.GetRoleAdmin(&_StakingContract.CallOpts, role)
}

// GetValidatorInfo is a free data retrieval call binding the contract method 0xb7797537.
//
// Solidity: function getValidatorInfo(uint256 idx) view returns(bytes pubkey, address withdrawalAddress, address claimAddress, uint256 userid)
func (_StakingContract *StakingContractCaller) GetValidatorInfo(opts *bind.CallOpts, idx *big.Int) (struct {
	Pubkey            []byte
	WithdrawalAddress common.Address
	ClaimAddress      common.Address
	Userid            *big.Int
}, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "getValidatorInfo", idx)

	outstruct := new(struct {
		Pubkey            []byte
		WithdrawalAddress common.Address
		ClaimAddress      common.Address
		Userid            *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pubkey = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.WithdrawalAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ClaimAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Userid = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetValidatorInfo is a free data retrieval call binding the contract method 0xb7797537.
//
// Solidity: function getValidatorInfo(uint256 idx) view returns(bytes pubkey, address withdrawalAddress, address claimAddress, uint256 userid)
func (_StakingContract *StakingContractSession) GetValidatorInfo(idx *big.Int) (struct {
	Pubkey            []byte
	WithdrawalAddress common.Address
	ClaimAddress      common.Address
	Userid            *big.Int
}, error) {
	return _StakingContract.Contract.GetValidatorInfo(&_StakingContract.CallOpts, idx)
}

// GetValidatorInfo is a free data retrieval call binding the contract method 0xb7797537.
//
// Solidity: function getValidatorInfo(uint256 idx) view returns(bytes pubkey, address withdrawalAddress, address claimAddress, uint256 userid)
func (_StakingContract *StakingContractCallerSession) GetValidatorInfo(idx *big.Int) (struct {
	Pubkey            []byte
	WithdrawalAddress common.Address
	ClaimAddress      common.Address
	Userid            *big.Int
}, error) {
	return _StakingContract.Contract.GetValidatorInfo(&_StakingContract.CallOpts, idx)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakingContract *StakingContractCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakingContract *StakingContractSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StakingContract.Contract.HasRole(&_StakingContract.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakingContract *StakingContractCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StakingContract.Contract.HasRole(&_StakingContract.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakingContract *StakingContractCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakingContract *StakingContractSession) Paused() (bool, error) {
	return _StakingContract.Contract.Paused(&_StakingContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakingContract *StakingContractCallerSession) Paused() (bool, error) {
	return _StakingContract.Contract.Paused(&_StakingContract.CallOpts)
}

// RewardPool is a free data retrieval call binding the contract method 0x66666aa9.
//
// Solidity: function rewardPool() view returns(address)
func (_StakingContract *StakingContractCaller) RewardPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "rewardPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardPool is a free data retrieval call binding the contract method 0x66666aa9.
//
// Solidity: function rewardPool() view returns(address)
func (_StakingContract *StakingContractSession) RewardPool() (common.Address, error) {
	return _StakingContract.Contract.RewardPool(&_StakingContract.CallOpts)
}

// RewardPool is a free data retrieval call binding the contract method 0x66666aa9.
//
// Solidity: function rewardPool() view returns(address)
func (_StakingContract *StakingContractCallerSession) RewardPool() (common.Address, error) {
	return _StakingContract.Contract.RewardPool(&_StakingContract.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakingContract *StakingContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakingContract *StakingContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StakingContract.Contract.SupportsInterface(&_StakingContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakingContract *StakingContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StakingContract.Contract.SupportsInterface(&_StakingContract.CallOpts, interfaceId)
}

// BatchDeposit is a paid mutator transaction binding the contract method 0x7e591740.
//
// Solidity: function batchDeposit(uint256 fromId, bytes[] signatures) returns()
func (_StakingContract *StakingContractTransactor) BatchDeposit(opts *bind.TransactOpts, fromId *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "batchDeposit", fromId, signatures)
}

// BatchDeposit is a paid mutator transaction binding the contract method 0x7e591740.
//
// Solidity: function batchDeposit(uint256 fromId, bytes[] signatures) returns()
func (_StakingContract *StakingContractSession) BatchDeposit(fromId *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _StakingContract.Contract.BatchDeposit(&_StakingContract.TransactOpts, fromId, signatures)
}

// BatchDeposit is a paid mutator transaction binding the contract method 0x7e591740.
//
// Solidity: function batchDeposit(uint256 fromId, bytes[] signatures) returns()
func (_StakingContract *StakingContractTransactorSession) BatchDeposit(fromId *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _StakingContract.Contract.BatchDeposit(&_StakingContract.TransactOpts, fromId, signatures)
}

// Exit is a paid mutator transaction binding the contract method 0x7f8661a1.
//
// Solidity: function exit(uint256 validatorId) returns()
func (_StakingContract *StakingContractTransactor) Exit(opts *bind.TransactOpts, validatorId *big.Int) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "exit", validatorId)
}

// Exit is a paid mutator transaction binding the contract method 0x7f8661a1.
//
// Solidity: function exit(uint256 validatorId) returns()
func (_StakingContract *StakingContractSession) Exit(validatorId *big.Int) (*types.Transaction, error) {
	return _StakingContract.Contract.Exit(&_StakingContract.TransactOpts, validatorId)
}

// Exit is a paid mutator transaction binding the contract method 0x7f8661a1.
//
// Solidity: function exit(uint256 validatorId) returns()
func (_StakingContract *StakingContractTransactorSession) Exit(validatorId *big.Int) (*types.Transaction, error) {
	return _StakingContract.Contract.Exit(&_StakingContract.TransactOpts, validatorId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakingContract *StakingContractTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakingContract *StakingContractSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingContract.Contract.GrantRole(&_StakingContract.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakingContract *StakingContractTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingContract.Contract.GrantRole(&_StakingContract.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_StakingContract *StakingContractTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_StakingContract *StakingContractSession) Initialize() (*types.Transaction, error) {
	return _StakingContract.Contract.Initialize(&_StakingContract.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_StakingContract *StakingContractTransactorSession) Initialize() (*types.Transaction, error) {
	return _StakingContract.Contract.Initialize(&_StakingContract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StakingContract *StakingContractTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StakingContract *StakingContractSession) Pause() (*types.Transaction, error) {
	return _StakingContract.Contract.Pause(&_StakingContract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StakingContract *StakingContractTransactorSession) Pause() (*types.Transaction, error) {
	return _StakingContract.Contract.Pause(&_StakingContract.TransactOpts)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x602a9eee.
//
// Solidity: function registerValidator(bytes pubkey) returns()
func (_StakingContract *StakingContractTransactor) RegisterValidator(opts *bind.TransactOpts, pubkey []byte) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "registerValidator", pubkey)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x602a9eee.
//
// Solidity: function registerValidator(bytes pubkey) returns()
func (_StakingContract *StakingContractSession) RegisterValidator(pubkey []byte) (*types.Transaction, error) {
	return _StakingContract.Contract.RegisterValidator(&_StakingContract.TransactOpts, pubkey)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x602a9eee.
//
// Solidity: function registerValidator(bytes pubkey) returns()
func (_StakingContract *StakingContractTransactorSession) RegisterValidator(pubkey []byte) (*types.Transaction, error) {
	return _StakingContract.Contract.RegisterValidator(&_StakingContract.TransactOpts, pubkey)
}

// RegisterValidators is a paid mutator transaction binding the contract method 0xdbd739ad.
//
// Solidity: function registerValidators(bytes[] pubkeys) returns()
func (_StakingContract *StakingContractTransactor) RegisterValidators(opts *bind.TransactOpts, pubkeys [][]byte) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "registerValidators", pubkeys)
}

// RegisterValidators is a paid mutator transaction binding the contract method 0xdbd739ad.
//
// Solidity: function registerValidators(bytes[] pubkeys) returns()
func (_StakingContract *StakingContractSession) RegisterValidators(pubkeys [][]byte) (*types.Transaction, error) {
	return _StakingContract.Contract.RegisterValidators(&_StakingContract.TransactOpts, pubkeys)
}

// RegisterValidators is a paid mutator transaction binding the contract method 0xdbd739ad.
//
// Solidity: function registerValidators(bytes[] pubkeys) returns()
func (_StakingContract *StakingContractTransactorSession) RegisterValidators(pubkeys [][]byte) (*types.Transaction, error) {
	return _StakingContract.Contract.RegisterValidators(&_StakingContract.TransactOpts, pubkeys)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StakingContract *StakingContractTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StakingContract *StakingContractSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingContract.Contract.RenounceRole(&_StakingContract.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StakingContract *StakingContractTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingContract.Contract.RenounceRole(&_StakingContract.TransactOpts, role, account)
}

// ReplaceValidator is a paid mutator transaction binding the contract method 0x970bf3a9.
//
// Solidity: function replaceValidator(uint256 idx, bytes pubkey) returns()
func (_StakingContract *StakingContractTransactor) ReplaceValidator(opts *bind.TransactOpts, idx *big.Int, pubkey []byte) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "replaceValidator", idx, pubkey)
}

// ReplaceValidator is a paid mutator transaction binding the contract method 0x970bf3a9.
//
// Solidity: function replaceValidator(uint256 idx, bytes pubkey) returns()
func (_StakingContract *StakingContractSession) ReplaceValidator(idx *big.Int, pubkey []byte) (*types.Transaction, error) {
	return _StakingContract.Contract.ReplaceValidator(&_StakingContract.TransactOpts, idx, pubkey)
}

// ReplaceValidator is a paid mutator transaction binding the contract method 0x970bf3a9.
//
// Solidity: function replaceValidator(uint256 idx, bytes pubkey) returns()
func (_StakingContract *StakingContractTransactorSession) ReplaceValidator(idx *big.Int, pubkey []byte) (*types.Transaction, error) {
	return _StakingContract.Contract.ReplaceValidator(&_StakingContract.TransactOpts, idx, pubkey)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakingContract *StakingContractTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakingContract *StakingContractSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingContract.Contract.RevokeRole(&_StakingContract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakingContract *StakingContractTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingContract.Contract.RevokeRole(&_StakingContract.TransactOpts, role, account)
}

// SetETHDepositContract is a paid mutator transaction binding the contract method 0x91b66caa.
//
// Solidity: function setETHDepositContract(address _ethDepositContract) returns()
func (_StakingContract *StakingContractTransactor) SetETHDepositContract(opts *bind.TransactOpts, _ethDepositContract common.Address) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "setETHDepositContract", _ethDepositContract)
}

// SetETHDepositContract is a paid mutator transaction binding the contract method 0x91b66caa.
//
// Solidity: function setETHDepositContract(address _ethDepositContract) returns()
func (_StakingContract *StakingContractSession) SetETHDepositContract(_ethDepositContract common.Address) (*types.Transaction, error) {
	return _StakingContract.Contract.SetETHDepositContract(&_StakingContract.TransactOpts, _ethDepositContract)
}

// SetETHDepositContract is a paid mutator transaction binding the contract method 0x91b66caa.
//
// Solidity: function setETHDepositContract(address _ethDepositContract) returns()
func (_StakingContract *StakingContractTransactorSession) SetETHDepositContract(_ethDepositContract common.Address) (*types.Transaction, error) {
	return _StakingContract.Contract.SetETHDepositContract(&_StakingContract.TransactOpts, _ethDepositContract)
}

// SetRewardPool is a paid mutator transaction binding the contract method 0x78238c37.
//
// Solidity: function setRewardPool(address _rewardPool) returns()
func (_StakingContract *StakingContractTransactor) SetRewardPool(opts *bind.TransactOpts, _rewardPool common.Address) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "setRewardPool", _rewardPool)
}

// SetRewardPool is a paid mutator transaction binding the contract method 0x78238c37.
//
// Solidity: function setRewardPool(address _rewardPool) returns()
func (_StakingContract *StakingContractSession) SetRewardPool(_rewardPool common.Address) (*types.Transaction, error) {
	return _StakingContract.Contract.SetRewardPool(&_StakingContract.TransactOpts, _rewardPool)
}

// SetRewardPool is a paid mutator transaction binding the contract method 0x78238c37.
//
// Solidity: function setRewardPool(address _rewardPool) returns()
func (_StakingContract *StakingContractTransactorSession) SetRewardPool(_rewardPool common.Address) (*types.Transaction, error) {
	return _StakingContract.Contract.SetRewardPool(&_StakingContract.TransactOpts, _rewardPool)
}

// Stake is a paid mutator transaction binding the contract method 0xf108fa93.
//
// Solidity: function stake(address withdrawaddr, address claimaddr, uint256 extradata, uint256 fee, uint256 deadline) payable returns()
func (_StakingContract *StakingContractTransactor) Stake(opts *bind.TransactOpts, withdrawaddr common.Address, claimaddr common.Address, extradata *big.Int, fee *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "stake", withdrawaddr, claimaddr, extradata, fee, deadline)
}

// Stake is a paid mutator transaction binding the contract method 0xf108fa93.
//
// Solidity: function stake(address withdrawaddr, address claimaddr, uint256 extradata, uint256 fee, uint256 deadline) payable returns()
func (_StakingContract *StakingContractSession) Stake(withdrawaddr common.Address, claimaddr common.Address, extradata *big.Int, fee *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _StakingContract.Contract.Stake(&_StakingContract.TransactOpts, withdrawaddr, claimaddr, extradata, fee, deadline)
}

// Stake is a paid mutator transaction binding the contract method 0xf108fa93.
//
// Solidity: function stake(address withdrawaddr, address claimaddr, uint256 extradata, uint256 fee, uint256 deadline) payable returns()
func (_StakingContract *StakingContractTransactorSession) Stake(withdrawaddr common.Address, claimaddr common.Address, extradata *big.Int, fee *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _StakingContract.Contract.Stake(&_StakingContract.TransactOpts, withdrawaddr, claimaddr, extradata, fee, deadline)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StakingContract *StakingContractTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StakingContract *StakingContractSession) Unpause() (*types.Transaction, error) {
	return _StakingContract.Contract.Unpause(&_StakingContract.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StakingContract *StakingContractTransactorSession) Unpause() (*types.Transaction, error) {
	return _StakingContract.Contract.Unpause(&_StakingContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingContract *StakingContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingContract *StakingContractSession) Receive() (*types.Transaction, error) {
	return _StakingContract.Contract.Receive(&_StakingContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingContract *StakingContractTransactorSession) Receive() (*types.Transaction, error) {
	return _StakingContract.Contract.Receive(&_StakingContract.TransactOpts)
}

// StakingContractDepositContractSetIterator is returned from FilterDepositContractSet and is used to iterate over the raw logs and unpacked data for DepositContractSet events raised by the StakingContract contract.
type StakingContractDepositContractSetIterator struct {
	Event *StakingContractDepositContractSet // Event containing the contract specifics and raw log

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
func (it *StakingContractDepositContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractDepositContractSet)
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
		it.Event = new(StakingContractDepositContractSet)
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
func (it *StakingContractDepositContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractDepositContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractDepositContractSet represents a DepositContractSet event raised by the StakingContract contract.
type StakingContractDepositContractSet struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDepositContractSet is a free log retrieval operation binding the contract event 0x1781ac9526b978975dba0fd26a33e044a55a7ace054a3ee7efa5f8459513bead.
//
// Solidity: event DepositContractSet(address addr)
func (_StakingContract *StakingContractFilterer) FilterDepositContractSet(opts *bind.FilterOpts) (*StakingContractDepositContractSetIterator, error) {

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "DepositContractSet")
	if err != nil {
		return nil, err
	}
	return &StakingContractDepositContractSetIterator{contract: _StakingContract.contract, event: "DepositContractSet", logs: logs, sub: sub}, nil
}

// WatchDepositContractSet is a free log subscription operation binding the contract event 0x1781ac9526b978975dba0fd26a33e044a55a7ace054a3ee7efa5f8459513bead.
//
// Solidity: event DepositContractSet(address addr)
func (_StakingContract *StakingContractFilterer) WatchDepositContractSet(opts *bind.WatchOpts, sink chan<- *StakingContractDepositContractSet) (event.Subscription, error) {

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "DepositContractSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractDepositContractSet)
				if err := _StakingContract.contract.UnpackLog(event, "DepositContractSet", log); err != nil {
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

// ParseDepositContractSet is a log parse operation binding the contract event 0x1781ac9526b978975dba0fd26a33e044a55a7ace054a3ee7efa5f8459513bead.
//
// Solidity: event DepositContractSet(address addr)
func (_StakingContract *StakingContractFilterer) ParseDepositContractSet(log types.Log) (*StakingContractDepositContractSet, error) {
	event := new(StakingContractDepositContractSet)
	if err := _StakingContract.contract.UnpackLog(event, "DepositContractSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StakingContract contract.
type StakingContractInitializedIterator struct {
	Event *StakingContractInitialized // Event containing the contract specifics and raw log

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
func (it *StakingContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractInitialized)
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
		it.Event = new(StakingContractInitialized)
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
func (it *StakingContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractInitialized represents a Initialized event raised by the StakingContract contract.
type StakingContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StakingContract *StakingContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakingContractInitializedIterator, error) {

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakingContractInitializedIterator{contract: _StakingContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StakingContract *StakingContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakingContractInitialized) (event.Subscription, error) {

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractInitialized)
				if err := _StakingContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_StakingContract *StakingContractFilterer) ParseInitialized(log types.Log) (*StakingContractInitialized, error) {
	event := new(StakingContractInitialized)
	if err := _StakingContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the StakingContract contract.
type StakingContractPausedIterator struct {
	Event *StakingContractPaused // Event containing the contract specifics and raw log

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
func (it *StakingContractPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractPaused)
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
		it.Event = new(StakingContractPaused)
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
func (it *StakingContractPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractPaused represents a Paused event raised by the StakingContract contract.
type StakingContractPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StakingContract *StakingContractFilterer) FilterPaused(opts *bind.FilterOpts) (*StakingContractPausedIterator, error) {

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StakingContractPausedIterator{contract: _StakingContract.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StakingContract *StakingContractFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StakingContractPaused) (event.Subscription, error) {

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractPaused)
				if err := _StakingContract.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_StakingContract *StakingContractFilterer) ParsePaused(log types.Log) (*StakingContractPaused, error) {
	event := new(StakingContractPaused)
	if err := _StakingContract.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractRewardPoolContractSetIterator is returned from FilterRewardPoolContractSet and is used to iterate over the raw logs and unpacked data for RewardPoolContractSet events raised by the StakingContract contract.
type StakingContractRewardPoolContractSetIterator struct {
	Event *StakingContractRewardPoolContractSet // Event containing the contract specifics and raw log

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
func (it *StakingContractRewardPoolContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractRewardPoolContractSet)
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
		it.Event = new(StakingContractRewardPoolContractSet)
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
func (it *StakingContractRewardPoolContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractRewardPoolContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractRewardPoolContractSet represents a RewardPoolContractSet event raised by the StakingContract contract.
type StakingContractRewardPoolContractSet struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRewardPoolContractSet is a free log retrieval operation binding the contract event 0x7e2602838378170d01002df4f80426d046e484841f67339edd3213b1875b1a89.
//
// Solidity: event RewardPoolContractSet(address addr)
func (_StakingContract *StakingContractFilterer) FilterRewardPoolContractSet(opts *bind.FilterOpts) (*StakingContractRewardPoolContractSetIterator, error) {

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "RewardPoolContractSet")
	if err != nil {
		return nil, err
	}
	return &StakingContractRewardPoolContractSetIterator{contract: _StakingContract.contract, event: "RewardPoolContractSet", logs: logs, sub: sub}, nil
}

// WatchRewardPoolContractSet is a free log subscription operation binding the contract event 0x7e2602838378170d01002df4f80426d046e484841f67339edd3213b1875b1a89.
//
// Solidity: event RewardPoolContractSet(address addr)
func (_StakingContract *StakingContractFilterer) WatchRewardPoolContractSet(opts *bind.WatchOpts, sink chan<- *StakingContractRewardPoolContractSet) (event.Subscription, error) {

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "RewardPoolContractSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractRewardPoolContractSet)
				if err := _StakingContract.contract.UnpackLog(event, "RewardPoolContractSet", log); err != nil {
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

// ParseRewardPoolContractSet is a log parse operation binding the contract event 0x7e2602838378170d01002df4f80426d046e484841f67339edd3213b1875b1a89.
//
// Solidity: event RewardPoolContractSet(address addr)
func (_StakingContract *StakingContractFilterer) ParseRewardPoolContractSet(log types.Log) (*StakingContractRewardPoolContractSet, error) {
	event := new(StakingContractRewardPoolContractSet)
	if err := _StakingContract.contract.UnpackLog(event, "RewardPoolContractSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the StakingContract contract.
type StakingContractRoleAdminChangedIterator struct {
	Event *StakingContractRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StakingContractRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractRoleAdminChanged)
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
		it.Event = new(StakingContractRoleAdminChanged)
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
func (it *StakingContractRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractRoleAdminChanged represents a RoleAdminChanged event raised by the StakingContract contract.
type StakingContractRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakingContract *StakingContractFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StakingContractRoleAdminChangedIterator, error) {

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

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StakingContractRoleAdminChangedIterator{contract: _StakingContract.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakingContract *StakingContractFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StakingContractRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractRoleAdminChanged)
				if err := _StakingContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_StakingContract *StakingContractFilterer) ParseRoleAdminChanged(log types.Log) (*StakingContractRoleAdminChanged, error) {
	event := new(StakingContractRoleAdminChanged)
	if err := _StakingContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the StakingContract contract.
type StakingContractRoleGrantedIterator struct {
	Event *StakingContractRoleGranted // Event containing the contract specifics and raw log

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
func (it *StakingContractRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractRoleGranted)
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
		it.Event = new(StakingContractRoleGranted)
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
func (it *StakingContractRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractRoleGranted represents a RoleGranted event raised by the StakingContract contract.
type StakingContractRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingContract *StakingContractFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakingContractRoleGrantedIterator, error) {

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

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakingContractRoleGrantedIterator{contract: _StakingContract.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingContract *StakingContractFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StakingContractRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractRoleGranted)
				if err := _StakingContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_StakingContract *StakingContractFilterer) ParseRoleGranted(log types.Log) (*StakingContractRoleGranted, error) {
	event := new(StakingContractRoleGranted)
	if err := _StakingContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the StakingContract contract.
type StakingContractRoleRevokedIterator struct {
	Event *StakingContractRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StakingContractRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractRoleRevoked)
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
		it.Event = new(StakingContractRoleRevoked)
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
func (it *StakingContractRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractRoleRevoked represents a RoleRevoked event raised by the StakingContract contract.
type StakingContractRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingContract *StakingContractFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakingContractRoleRevokedIterator, error) {

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

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakingContractRoleRevokedIterator{contract: _StakingContract.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingContract *StakingContractFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StakingContractRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractRoleRevoked)
				if err := _StakingContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_StakingContract *StakingContractFilterer) ParseRoleRevoked(log types.Log) (*StakingContractRoleRevoked, error) {
	event := new(StakingContractRoleRevoked)
	if err := _StakingContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the StakingContract contract.
type StakingContractStakedIterator struct {
	Event *StakingContractStaked // Event containing the contract specifics and raw log

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
func (it *StakingContractStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractStaked)
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
		it.Event = new(StakingContractStaked)
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
func (it *StakingContractStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractStaked represents a Staked event raised by the StakingContract contract.
type StakingContractStaked struct {
	Addr   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address addr, uint256 amount)
func (_StakingContract *StakingContractFilterer) FilterStaked(opts *bind.FilterOpts) (*StakingContractStakedIterator, error) {

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "Staked")
	if err != nil {
		return nil, err
	}
	return &StakingContractStakedIterator{contract: _StakingContract.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address addr, uint256 amount)
func (_StakingContract *StakingContractFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *StakingContractStaked) (event.Subscription, error) {

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "Staked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractStaked)
				if err := _StakingContract.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address addr, uint256 amount)
func (_StakingContract *StakingContractFilterer) ParseStaked(log types.Log) (*StakingContractStaked, error) {
	event := new(StakingContractStaked)
	if err := _StakingContract.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the StakingContract contract.
type StakingContractUnpausedIterator struct {
	Event *StakingContractUnpaused // Event containing the contract specifics and raw log

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
func (it *StakingContractUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractUnpaused)
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
		it.Event = new(StakingContractUnpaused)
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
func (it *StakingContractUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractUnpaused represents a Unpaused event raised by the StakingContract contract.
type StakingContractUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StakingContract *StakingContractFilterer) FilterUnpaused(opts *bind.FilterOpts) (*StakingContractUnpausedIterator, error) {

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &StakingContractUnpausedIterator{contract: _StakingContract.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StakingContract *StakingContractFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StakingContractUnpaused) (event.Subscription, error) {

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractUnpaused)
				if err := _StakingContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_StakingContract *StakingContractFilterer) ParseUnpaused(log types.Log) (*StakingContractUnpaused, error) {
	event := new(StakingContractUnpaused)
	if err := _StakingContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
