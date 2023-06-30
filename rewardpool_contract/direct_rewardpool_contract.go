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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"DepositContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"RewardPoolContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"ShangHaiStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"SignerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEPOSIT_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REGISTRY_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"validatorIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"exitToClaimAddress\",\"type\":\"bool\"}],\"name\":\"batchEmergencyExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"validatorIds\",\"type\":\"uint256[]\"}],\"name\":\"batchExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exitToClaimAddress\",\"type\":\"bool\"}],\"name\":\"emergencyExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethDepositContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"}],\"name\":\"getExitQueue\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExitQueueLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getValidatorInfo\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"claimAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"extraData\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"}],\"name\":\"getValidatorInfos\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"claimAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"extraDatas\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethDepositContract\",\"type\":\"address\"}],\"name\":\"setETHDepositContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardPool\",\"type\":\"address\"}],\"name\":\"setRewardPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"setSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"claimaddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"withdrawaddr\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"paramsSig\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"extradata\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tips\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sysSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleShangHai\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"extraData\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"claimaddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"withdrawaddr\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"paramsSig\",\"type\":\"bytes\"}],\"name\":\"verifySigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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

// DEPOSITSIZE is a free data retrieval call binding the contract method 0x36bf3325.
//
// Solidity: function DEPOSIT_SIZE() view returns(uint256)
func (_RewardpoolContract *RewardpoolContractCaller) DEPOSITSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardpoolContract.contract.Call(opts, &out, "DEPOSIT_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEPOSITSIZE is a free data retrieval call binding the contract method 0x36bf3325.
//
// Solidity: function DEPOSIT_SIZE() view returns(uint256)
func (_RewardpoolContract *RewardpoolContractSession) DEPOSITSIZE() (*big.Int, error) {
	return _RewardpoolContract.Contract.DEPOSITSIZE(&_RewardpoolContract.CallOpts)
}

// DEPOSITSIZE is a free data retrieval call binding the contract method 0x36bf3325.
//
// Solidity: function DEPOSIT_SIZE() view returns(uint256)
func (_RewardpoolContract *RewardpoolContractCallerSession) DEPOSITSIZE() (*big.Int, error) {
	return _RewardpoolContract.Contract.DEPOSITSIZE(&_RewardpoolContract.CallOpts)
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

// REGISTRYROLE is a free data retrieval call binding the contract method 0x42f1e879.
//
// Solidity: function REGISTRY_ROLE() view returns(bytes32)
func (_RewardpoolContract *RewardpoolContractCaller) REGISTRYROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RewardpoolContract.contract.Call(opts, &out, "REGISTRY_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REGISTRYROLE is a free data retrieval call binding the contract method 0x42f1e879.
//
// Solidity: function REGISTRY_ROLE() view returns(bytes32)
func (_RewardpoolContract *RewardpoolContractSession) REGISTRYROLE() ([32]byte, error) {
	return _RewardpoolContract.Contract.REGISTRYROLE(&_RewardpoolContract.CallOpts)
}

// REGISTRYROLE is a free data retrieval call binding the contract method 0x42f1e879.
//
// Solidity: function REGISTRY_ROLE() view returns(bytes32)
func (_RewardpoolContract *RewardpoolContractCallerSession) REGISTRYROLE() ([32]byte, error) {
	return _RewardpoolContract.Contract.REGISTRYROLE(&_RewardpoolContract.CallOpts)
}

// EthDepositContract is a free data retrieval call binding the contract method 0x3884545d.
//
// Solidity: function ethDepositContract() view returns(address)
func (_RewardpoolContract *RewardpoolContractCaller) EthDepositContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardpoolContract.contract.Call(opts, &out, "ethDepositContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthDepositContract is a free data retrieval call binding the contract method 0x3884545d.
//
// Solidity: function ethDepositContract() view returns(address)
func (_RewardpoolContract *RewardpoolContractSession) EthDepositContract() (common.Address, error) {
	return _RewardpoolContract.Contract.EthDepositContract(&_RewardpoolContract.CallOpts)
}

// EthDepositContract is a free data retrieval call binding the contract method 0x3884545d.
//
// Solidity: function ethDepositContract() view returns(address)
func (_RewardpoolContract *RewardpoolContractCallerSession) EthDepositContract() (common.Address, error) {
	return _RewardpoolContract.Contract.EthDepositContract(&_RewardpoolContract.CallOpts)
}

// GetExitQueue is a free data retrieval call binding the contract method 0x43931256.
//
// Solidity: function getExitQueue(uint256 from, uint256 to) view returns(uint256[])
func (_RewardpoolContract *RewardpoolContractCaller) GetExitQueue(opts *bind.CallOpts, from *big.Int, to *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _RewardpoolContract.contract.Call(opts, &out, "getExitQueue", from, to)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetExitQueue is a free data retrieval call binding the contract method 0x43931256.
//
// Solidity: function getExitQueue(uint256 from, uint256 to) view returns(uint256[])
func (_RewardpoolContract *RewardpoolContractSession) GetExitQueue(from *big.Int, to *big.Int) ([]*big.Int, error) {
	return _RewardpoolContract.Contract.GetExitQueue(&_RewardpoolContract.CallOpts, from, to)
}

// GetExitQueue is a free data retrieval call binding the contract method 0x43931256.
//
// Solidity: function getExitQueue(uint256 from, uint256 to) view returns(uint256[])
func (_RewardpoolContract *RewardpoolContractCallerSession) GetExitQueue(from *big.Int, to *big.Int) ([]*big.Int, error) {
	return _RewardpoolContract.Contract.GetExitQueue(&_RewardpoolContract.CallOpts, from, to)
}

// GetExitQueueLength is a free data retrieval call binding the contract method 0x559230a0.
//
// Solidity: function getExitQueueLength() view returns(uint256)
func (_RewardpoolContract *RewardpoolContractCaller) GetExitQueueLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardpoolContract.contract.Call(opts, &out, "getExitQueueLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExitQueueLength is a free data retrieval call binding the contract method 0x559230a0.
//
// Solidity: function getExitQueueLength() view returns(uint256)
func (_RewardpoolContract *RewardpoolContractSession) GetExitQueueLength() (*big.Int, error) {
	return _RewardpoolContract.Contract.GetExitQueueLength(&_RewardpoolContract.CallOpts)
}

// GetExitQueueLength is a free data retrieval call binding the contract method 0x559230a0.
//
// Solidity: function getExitQueueLength() view returns(uint256)
func (_RewardpoolContract *RewardpoolContractCallerSession) GetExitQueueLength() (*big.Int, error) {
	return _RewardpoolContract.Contract.GetExitQueueLength(&_RewardpoolContract.CallOpts)
}

// GetNextValidators is a free data retrieval call binding the contract method 0x40cddab3.
//
// Solidity: function getNextValidators() view returns(uint256)
func (_RewardpoolContract *RewardpoolContractCaller) GetNextValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardpoolContract.contract.Call(opts, &out, "getNextValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextValidators is a free data retrieval call binding the contract method 0x40cddab3.
//
// Solidity: function getNextValidators() view returns(uint256)
func (_RewardpoolContract *RewardpoolContractSession) GetNextValidators() (*big.Int, error) {
	return _RewardpoolContract.Contract.GetNextValidators(&_RewardpoolContract.CallOpts)
}

// GetNextValidators is a free data retrieval call binding the contract method 0x40cddab3.
//
// Solidity: function getNextValidators() view returns(uint256)
func (_RewardpoolContract *RewardpoolContractCallerSession) GetNextValidators() (*big.Int, error) {
	return _RewardpoolContract.Contract.GetNextValidators(&_RewardpoolContract.CallOpts)
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

// GetValidatorInfo is a free data retrieval call binding the contract method 0xb7797537.
//
// Solidity: function getValidatorInfo(uint256 idx) view returns(bytes pubkey, address claimAddress, uint256 extraData)
func (_RewardpoolContract *RewardpoolContractCaller) GetValidatorInfo(opts *bind.CallOpts, idx *big.Int) (struct {
	Pubkey       []byte
	ClaimAddress common.Address
	ExtraData    *big.Int
}, error) {
	var out []interface{}
	err := _RewardpoolContract.contract.Call(opts, &out, "getValidatorInfo", idx)

	outstruct := new(struct {
		Pubkey       []byte
		ClaimAddress common.Address
		ExtraData    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pubkey = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.ClaimAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ExtraData = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetValidatorInfo is a free data retrieval call binding the contract method 0xb7797537.
//
// Solidity: function getValidatorInfo(uint256 idx) view returns(bytes pubkey, address claimAddress, uint256 extraData)
func (_RewardpoolContract *RewardpoolContractSession) GetValidatorInfo(idx *big.Int) (struct {
	Pubkey       []byte
	ClaimAddress common.Address
	ExtraData    *big.Int
}, error) {
	return _RewardpoolContract.Contract.GetValidatorInfo(&_RewardpoolContract.CallOpts, idx)
}

// GetValidatorInfo is a free data retrieval call binding the contract method 0xb7797537.
//
// Solidity: function getValidatorInfo(uint256 idx) view returns(bytes pubkey, address claimAddress, uint256 extraData)
func (_RewardpoolContract *RewardpoolContractCallerSession) GetValidatorInfo(idx *big.Int) (struct {
	Pubkey       []byte
	ClaimAddress common.Address
	ExtraData    *big.Int
}, error) {
	return _RewardpoolContract.Contract.GetValidatorInfo(&_RewardpoolContract.CallOpts, idx)
}

// GetValidatorInfos is a free data retrieval call binding the contract method 0x9e054533.
//
// Solidity: function getValidatorInfos(uint256 from, uint256 to) view returns(bytes[] pubkeys, address[] claimAddresses, uint256[] extraDatas)
func (_RewardpoolContract *RewardpoolContractCaller) GetValidatorInfos(opts *bind.CallOpts, from *big.Int, to *big.Int) (struct {
	Pubkeys        [][]byte
	ClaimAddresses []common.Address
	ExtraDatas     []*big.Int
}, error) {
	var out []interface{}
	err := _RewardpoolContract.contract.Call(opts, &out, "getValidatorInfos", from, to)

	outstruct := new(struct {
		Pubkeys        [][]byte
		ClaimAddresses []common.Address
		ExtraDatas     []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pubkeys = *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)
	outstruct.ClaimAddresses = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.ExtraDatas = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetValidatorInfos is a free data retrieval call binding the contract method 0x9e054533.
//
// Solidity: function getValidatorInfos(uint256 from, uint256 to) view returns(bytes[] pubkeys, address[] claimAddresses, uint256[] extraDatas)
func (_RewardpoolContract *RewardpoolContractSession) GetValidatorInfos(from *big.Int, to *big.Int) (struct {
	Pubkeys        [][]byte
	ClaimAddresses []common.Address
	ExtraDatas     []*big.Int
}, error) {
	return _RewardpoolContract.Contract.GetValidatorInfos(&_RewardpoolContract.CallOpts, from, to)
}

// GetValidatorInfos is a free data retrieval call binding the contract method 0x9e054533.
//
// Solidity: function getValidatorInfos(uint256 from, uint256 to) view returns(bytes[] pubkeys, address[] claimAddresses, uint256[] extraDatas)
func (_RewardpoolContract *RewardpoolContractCallerSession) GetValidatorInfos(from *big.Int, to *big.Int) (struct {
	Pubkeys        [][]byte
	ClaimAddresses []common.Address
	ExtraDatas     []*big.Int
}, error) {
	return _RewardpoolContract.Contract.GetValidatorInfos(&_RewardpoolContract.CallOpts, from, to)
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

// RewardPool is a free data retrieval call binding the contract method 0x66666aa9.
//
// Solidity: function rewardPool() view returns(address)
func (_RewardpoolContract *RewardpoolContractCaller) RewardPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardpoolContract.contract.Call(opts, &out, "rewardPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardPool is a free data retrieval call binding the contract method 0x66666aa9.
//
// Solidity: function rewardPool() view returns(address)
func (_RewardpoolContract *RewardpoolContractSession) RewardPool() (common.Address, error) {
	return _RewardpoolContract.Contract.RewardPool(&_RewardpoolContract.CallOpts)
}

// RewardPool is a free data retrieval call binding the contract method 0x66666aa9.
//
// Solidity: function rewardPool() view returns(address)
func (_RewardpoolContract *RewardpoolContractCallerSession) RewardPool() (common.Address, error) {
	return _RewardpoolContract.Contract.RewardPool(&_RewardpoolContract.CallOpts)
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

// SysSigner is a free data retrieval call binding the contract method 0xfa734b1c.
//
// Solidity: function sysSigner() view returns(address)
func (_RewardpoolContract *RewardpoolContractCaller) SysSigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardpoolContract.contract.Call(opts, &out, "sysSigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SysSigner is a free data retrieval call binding the contract method 0xfa734b1c.
//
// Solidity: function sysSigner() view returns(address)
func (_RewardpoolContract *RewardpoolContractSession) SysSigner() (common.Address, error) {
	return _RewardpoolContract.Contract.SysSigner(&_RewardpoolContract.CallOpts)
}

// SysSigner is a free data retrieval call binding the contract method 0xfa734b1c.
//
// Solidity: function sysSigner() view returns(address)
func (_RewardpoolContract *RewardpoolContractCallerSession) SysSigner() (common.Address, error) {
	return _RewardpoolContract.Contract.SysSigner(&_RewardpoolContract.CallOpts)
}

// VerifySigner is a free data retrieval call binding the contract method 0x24361e1b.
//
// Solidity: function verifySigner(uint256 extraData, address claimaddr, address withdrawaddr, bytes[] pubkeys, bytes[] signatures, bytes paramsSig) view returns(bool)
func (_RewardpoolContract *RewardpoolContractCaller) VerifySigner(opts *bind.CallOpts, extraData *big.Int, claimaddr common.Address, withdrawaddr common.Address, pubkeys [][]byte, signatures [][]byte, paramsSig []byte) (bool, error) {
	var out []interface{}
	err := _RewardpoolContract.contract.Call(opts, &out, "verifySigner", extraData, claimaddr, withdrawaddr, pubkeys, signatures, paramsSig)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifySigner is a free data retrieval call binding the contract method 0x24361e1b.
//
// Solidity: function verifySigner(uint256 extraData, address claimaddr, address withdrawaddr, bytes[] pubkeys, bytes[] signatures, bytes paramsSig) view returns(bool)
func (_RewardpoolContract *RewardpoolContractSession) VerifySigner(extraData *big.Int, claimaddr common.Address, withdrawaddr common.Address, pubkeys [][]byte, signatures [][]byte, paramsSig []byte) (bool, error) {
	return _RewardpoolContract.Contract.VerifySigner(&_RewardpoolContract.CallOpts, extraData, claimaddr, withdrawaddr, pubkeys, signatures, paramsSig)
}

// VerifySigner is a free data retrieval call binding the contract method 0x24361e1b.
//
// Solidity: function verifySigner(uint256 extraData, address claimaddr, address withdrawaddr, bytes[] pubkeys, bytes[] signatures, bytes paramsSig) view returns(bool)
func (_RewardpoolContract *RewardpoolContractCallerSession) VerifySigner(extraData *big.Int, claimaddr common.Address, withdrawaddr common.Address, pubkeys [][]byte, signatures [][]byte, paramsSig []byte) (bool, error) {
	return _RewardpoolContract.Contract.VerifySigner(&_RewardpoolContract.CallOpts, extraData, claimaddr, withdrawaddr, pubkeys, signatures, paramsSig)
}

// BatchEmergencyExit is a paid mutator transaction binding the contract method 0x2393df84.
//
// Solidity: function batchEmergencyExit(uint256[] validatorIds, bool exitToClaimAddress) returns()
func (_RewardpoolContract *RewardpoolContractTransactor) BatchEmergencyExit(opts *bind.TransactOpts, validatorIds []*big.Int, exitToClaimAddress bool) (*types.Transaction, error) {
	return _RewardpoolContract.contract.Transact(opts, "batchEmergencyExit", validatorIds, exitToClaimAddress)
}

// BatchEmergencyExit is a paid mutator transaction binding the contract method 0x2393df84.
//
// Solidity: function batchEmergencyExit(uint256[] validatorIds, bool exitToClaimAddress) returns()
func (_RewardpoolContract *RewardpoolContractSession) BatchEmergencyExit(validatorIds []*big.Int, exitToClaimAddress bool) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.BatchEmergencyExit(&_RewardpoolContract.TransactOpts, validatorIds, exitToClaimAddress)
}

// BatchEmergencyExit is a paid mutator transaction binding the contract method 0x2393df84.
//
// Solidity: function batchEmergencyExit(uint256[] validatorIds, bool exitToClaimAddress) returns()
func (_RewardpoolContract *RewardpoolContractTransactorSession) BatchEmergencyExit(validatorIds []*big.Int, exitToClaimAddress bool) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.BatchEmergencyExit(&_RewardpoolContract.TransactOpts, validatorIds, exitToClaimAddress)
}

// BatchExit is a paid mutator transaction binding the contract method 0x4dd1fb0b.
//
// Solidity: function batchExit(uint256[] validatorIds) returns()
func (_RewardpoolContract *RewardpoolContractTransactor) BatchExit(opts *bind.TransactOpts, validatorIds []*big.Int) (*types.Transaction, error) {
	return _RewardpoolContract.contract.Transact(opts, "batchExit", validatorIds)
}

// BatchExit is a paid mutator transaction binding the contract method 0x4dd1fb0b.
//
// Solidity: function batchExit(uint256[] validatorIds) returns()
func (_RewardpoolContract *RewardpoolContractSession) BatchExit(validatorIds []*big.Int) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.BatchExit(&_RewardpoolContract.TransactOpts, validatorIds)
}

// BatchExit is a paid mutator transaction binding the contract method 0x4dd1fb0b.
//
// Solidity: function batchExit(uint256[] validatorIds) returns()
func (_RewardpoolContract *RewardpoolContractTransactorSession) BatchExit(validatorIds []*big.Int) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.BatchExit(&_RewardpoolContract.TransactOpts, validatorIds)
}

// EmergencyExit is a paid mutator transaction binding the contract method 0xf806e87a.
//
// Solidity: function emergencyExit(uint256 validatorId, bool exitToClaimAddress) returns()
func (_RewardpoolContract *RewardpoolContractTransactor) EmergencyExit(opts *bind.TransactOpts, validatorId *big.Int, exitToClaimAddress bool) (*types.Transaction, error) {
	return _RewardpoolContract.contract.Transact(opts, "emergencyExit", validatorId, exitToClaimAddress)
}

// EmergencyExit is a paid mutator transaction binding the contract method 0xf806e87a.
//
// Solidity: function emergencyExit(uint256 validatorId, bool exitToClaimAddress) returns()
func (_RewardpoolContract *RewardpoolContractSession) EmergencyExit(validatorId *big.Int, exitToClaimAddress bool) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.EmergencyExit(&_RewardpoolContract.TransactOpts, validatorId, exitToClaimAddress)
}

// EmergencyExit is a paid mutator transaction binding the contract method 0xf806e87a.
//
// Solidity: function emergencyExit(uint256 validatorId, bool exitToClaimAddress) returns()
func (_RewardpoolContract *RewardpoolContractTransactorSession) EmergencyExit(validatorId *big.Int, exitToClaimAddress bool) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.EmergencyExit(&_RewardpoolContract.TransactOpts, validatorId, exitToClaimAddress)
}

// Exit is a paid mutator transaction binding the contract method 0x7f8661a1.
//
// Solidity: function exit(uint256 validatorId) returns()
func (_RewardpoolContract *RewardpoolContractTransactor) Exit(opts *bind.TransactOpts, validatorId *big.Int) (*types.Transaction, error) {
	return _RewardpoolContract.contract.Transact(opts, "exit", validatorId)
}

// Exit is a paid mutator transaction binding the contract method 0x7f8661a1.
//
// Solidity: function exit(uint256 validatorId) returns()
func (_RewardpoolContract *RewardpoolContractSession) Exit(validatorId *big.Int) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.Exit(&_RewardpoolContract.TransactOpts, validatorId)
}

// Exit is a paid mutator transaction binding the contract method 0x7f8661a1.
//
// Solidity: function exit(uint256 validatorId) returns()
func (_RewardpoolContract *RewardpoolContractTransactorSession) Exit(validatorId *big.Int) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.Exit(&_RewardpoolContract.TransactOpts, validatorId)
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

// SetETHDepositContract is a paid mutator transaction binding the contract method 0x91b66caa.
//
// Solidity: function setETHDepositContract(address _ethDepositContract) returns()
func (_RewardpoolContract *RewardpoolContractTransactor) SetETHDepositContract(opts *bind.TransactOpts, _ethDepositContract common.Address) (*types.Transaction, error) {
	return _RewardpoolContract.contract.Transact(opts, "setETHDepositContract", _ethDepositContract)
}

// SetETHDepositContract is a paid mutator transaction binding the contract method 0x91b66caa.
//
// Solidity: function setETHDepositContract(address _ethDepositContract) returns()
func (_RewardpoolContract *RewardpoolContractSession) SetETHDepositContract(_ethDepositContract common.Address) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.SetETHDepositContract(&_RewardpoolContract.TransactOpts, _ethDepositContract)
}

// SetETHDepositContract is a paid mutator transaction binding the contract method 0x91b66caa.
//
// Solidity: function setETHDepositContract(address _ethDepositContract) returns()
func (_RewardpoolContract *RewardpoolContractTransactorSession) SetETHDepositContract(_ethDepositContract common.Address) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.SetETHDepositContract(&_RewardpoolContract.TransactOpts, _ethDepositContract)
}

// SetRewardPool is a paid mutator transaction binding the contract method 0x78238c37.
//
// Solidity: function setRewardPool(address _rewardPool) returns()
func (_RewardpoolContract *RewardpoolContractTransactor) SetRewardPool(opts *bind.TransactOpts, _rewardPool common.Address) (*types.Transaction, error) {
	return _RewardpoolContract.contract.Transact(opts, "setRewardPool", _rewardPool)
}

// SetRewardPool is a paid mutator transaction binding the contract method 0x78238c37.
//
// Solidity: function setRewardPool(address _rewardPool) returns()
func (_RewardpoolContract *RewardpoolContractSession) SetRewardPool(_rewardPool common.Address) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.SetRewardPool(&_RewardpoolContract.TransactOpts, _rewardPool)
}

// SetRewardPool is a paid mutator transaction binding the contract method 0x78238c37.
//
// Solidity: function setRewardPool(address _rewardPool) returns()
func (_RewardpoolContract *RewardpoolContractTransactorSession) SetRewardPool(_rewardPool common.Address) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.SetRewardPool(&_RewardpoolContract.TransactOpts, _rewardPool)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address _signer) returns()
func (_RewardpoolContract *RewardpoolContractTransactor) SetSigner(opts *bind.TransactOpts, _signer common.Address) (*types.Transaction, error) {
	return _RewardpoolContract.contract.Transact(opts, "setSigner", _signer)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address _signer) returns()
func (_RewardpoolContract *RewardpoolContractSession) SetSigner(_signer common.Address) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.SetSigner(&_RewardpoolContract.TransactOpts, _signer)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address _signer) returns()
func (_RewardpoolContract *RewardpoolContractTransactorSession) SetSigner(_signer common.Address) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.SetSigner(&_RewardpoolContract.TransactOpts, _signer)
}

// Stake is a paid mutator transaction binding the contract method 0x238d11bf.
//
// Solidity: function stake(address claimaddr, address withdrawaddr, bytes[] pubkeys, bytes[] signatures, bytes paramsSig, uint256 extradata, uint256 tips) payable returns()
func (_RewardpoolContract *RewardpoolContractTransactor) Stake(opts *bind.TransactOpts, claimaddr common.Address, withdrawaddr common.Address, pubkeys [][]byte, signatures [][]byte, paramsSig []byte, extradata *big.Int, tips *big.Int) (*types.Transaction, error) {
	return _RewardpoolContract.contract.Transact(opts, "stake", claimaddr, withdrawaddr, pubkeys, signatures, paramsSig, extradata, tips)
}

// Stake is a paid mutator transaction binding the contract method 0x238d11bf.
//
// Solidity: function stake(address claimaddr, address withdrawaddr, bytes[] pubkeys, bytes[] signatures, bytes paramsSig, uint256 extradata, uint256 tips) payable returns()
func (_RewardpoolContract *RewardpoolContractSession) Stake(claimaddr common.Address, withdrawaddr common.Address, pubkeys [][]byte, signatures [][]byte, paramsSig []byte, extradata *big.Int, tips *big.Int) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.Stake(&_RewardpoolContract.TransactOpts, claimaddr, withdrawaddr, pubkeys, signatures, paramsSig, extradata, tips)
}

// Stake is a paid mutator transaction binding the contract method 0x238d11bf.
//
// Solidity: function stake(address claimaddr, address withdrawaddr, bytes[] pubkeys, bytes[] signatures, bytes paramsSig, uint256 extradata, uint256 tips) payable returns()
func (_RewardpoolContract *RewardpoolContractTransactorSession) Stake(claimaddr common.Address, withdrawaddr common.Address, pubkeys [][]byte, signatures [][]byte, paramsSig []byte, extradata *big.Int, tips *big.Int) (*types.Transaction, error) {
	return _RewardpoolContract.Contract.Stake(&_RewardpoolContract.TransactOpts, claimaddr, withdrawaddr, pubkeys, signatures, paramsSig, extradata, tips)
}

// ToggleShangHai is a paid mutator transaction binding the contract method 0x15f6476f.
//
// Solidity: function toggleShangHai() returns()
func (_RewardpoolContract *RewardpoolContractTransactor) ToggleShangHai(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardpoolContract.contract.Transact(opts, "toggleShangHai")
}

// ToggleShangHai is a paid mutator transaction binding the contract method 0x15f6476f.
//
// Solidity: function toggleShangHai() returns()
func (_RewardpoolContract *RewardpoolContractSession) ToggleShangHai() (*types.Transaction, error) {
	return _RewardpoolContract.Contract.ToggleShangHai(&_RewardpoolContract.TransactOpts)
}

// ToggleShangHai is a paid mutator transaction binding the contract method 0x15f6476f.
//
// Solidity: function toggleShangHai() returns()
func (_RewardpoolContract *RewardpoolContractTransactorSession) ToggleShangHai() (*types.Transaction, error) {
	return _RewardpoolContract.Contract.ToggleShangHai(&_RewardpoolContract.TransactOpts)
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

// RewardpoolContractDepositContractSetIterator is returned from FilterDepositContractSet and is used to iterate over the raw logs and unpacked data for DepositContractSet events raised by the RewardpoolContract contract.
type RewardpoolContractDepositContractSetIterator struct {
	Event *RewardpoolContractDepositContractSet // Event containing the contract specifics and raw log

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
func (it *RewardpoolContractDepositContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardpoolContractDepositContractSet)
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
		it.Event = new(RewardpoolContractDepositContractSet)
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
func (it *RewardpoolContractDepositContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardpoolContractDepositContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardpoolContractDepositContractSet represents a DepositContractSet event raised by the RewardpoolContract contract.
type RewardpoolContractDepositContractSet struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDepositContractSet is a free log retrieval operation binding the contract event 0x1781ac9526b978975dba0fd26a33e044a55a7ace054a3ee7efa5f8459513bead.
//
// Solidity: event DepositContractSet(address addr)
func (_RewardpoolContract *RewardpoolContractFilterer) FilterDepositContractSet(opts *bind.FilterOpts) (*RewardpoolContractDepositContractSetIterator, error) {

	logs, sub, err := _RewardpoolContract.contract.FilterLogs(opts, "DepositContractSet")
	if err != nil {
		return nil, err
	}
	return &RewardpoolContractDepositContractSetIterator{contract: _RewardpoolContract.contract, event: "DepositContractSet", logs: logs, sub: sub}, nil
}

// WatchDepositContractSet is a free log subscription operation binding the contract event 0x1781ac9526b978975dba0fd26a33e044a55a7ace054a3ee7efa5f8459513bead.
//
// Solidity: event DepositContractSet(address addr)
func (_RewardpoolContract *RewardpoolContractFilterer) WatchDepositContractSet(opts *bind.WatchOpts, sink chan<- *RewardpoolContractDepositContractSet) (event.Subscription, error) {

	logs, sub, err := _RewardpoolContract.contract.WatchLogs(opts, "DepositContractSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardpoolContractDepositContractSet)
				if err := _RewardpoolContract.contract.UnpackLog(event, "DepositContractSet", log); err != nil {
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
func (_RewardpoolContract *RewardpoolContractFilterer) ParseDepositContractSet(log types.Log) (*RewardpoolContractDepositContractSet, error) {
	event := new(RewardpoolContractDepositContractSet)
	if err := _RewardpoolContract.contract.UnpackLog(event, "DepositContractSet", log); err != nil {
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

// RewardpoolContractRewardPoolContractSetIterator is returned from FilterRewardPoolContractSet and is used to iterate over the raw logs and unpacked data for RewardPoolContractSet events raised by the RewardpoolContract contract.
type RewardpoolContractRewardPoolContractSetIterator struct {
	Event *RewardpoolContractRewardPoolContractSet // Event containing the contract specifics and raw log

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
func (it *RewardpoolContractRewardPoolContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardpoolContractRewardPoolContractSet)
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
		it.Event = new(RewardpoolContractRewardPoolContractSet)
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
func (it *RewardpoolContractRewardPoolContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardpoolContractRewardPoolContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardpoolContractRewardPoolContractSet represents a RewardPoolContractSet event raised by the RewardpoolContract contract.
type RewardpoolContractRewardPoolContractSet struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRewardPoolContractSet is a free log retrieval operation binding the contract event 0x7e2602838378170d01002df4f80426d046e484841f67339edd3213b1875b1a89.
//
// Solidity: event RewardPoolContractSet(address addr)
func (_RewardpoolContract *RewardpoolContractFilterer) FilterRewardPoolContractSet(opts *bind.FilterOpts) (*RewardpoolContractRewardPoolContractSetIterator, error) {

	logs, sub, err := _RewardpoolContract.contract.FilterLogs(opts, "RewardPoolContractSet")
	if err != nil {
		return nil, err
	}
	return &RewardpoolContractRewardPoolContractSetIterator{contract: _RewardpoolContract.contract, event: "RewardPoolContractSet", logs: logs, sub: sub}, nil
}

// WatchRewardPoolContractSet is a free log subscription operation binding the contract event 0x7e2602838378170d01002df4f80426d046e484841f67339edd3213b1875b1a89.
//
// Solidity: event RewardPoolContractSet(address addr)
func (_RewardpoolContract *RewardpoolContractFilterer) WatchRewardPoolContractSet(opts *bind.WatchOpts, sink chan<- *RewardpoolContractRewardPoolContractSet) (event.Subscription, error) {

	logs, sub, err := _RewardpoolContract.contract.WatchLogs(opts, "RewardPoolContractSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardpoolContractRewardPoolContractSet)
				if err := _RewardpoolContract.contract.UnpackLog(event, "RewardPoolContractSet", log); err != nil {
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
func (_RewardpoolContract *RewardpoolContractFilterer) ParseRewardPoolContractSet(log types.Log) (*RewardpoolContractRewardPoolContractSet, error) {
	event := new(RewardpoolContractRewardPoolContractSet)
	if err := _RewardpoolContract.contract.UnpackLog(event, "RewardPoolContractSet", log); err != nil {
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

// RewardpoolContractShangHaiStatusIterator is returned from FilterShangHaiStatus and is used to iterate over the raw logs and unpacked data for ShangHaiStatus events raised by the RewardpoolContract contract.
type RewardpoolContractShangHaiStatusIterator struct {
	Event *RewardpoolContractShangHaiStatus // Event containing the contract specifics and raw log

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
func (it *RewardpoolContractShangHaiStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardpoolContractShangHaiStatus)
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
		it.Event = new(RewardpoolContractShangHaiStatus)
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
func (it *RewardpoolContractShangHaiStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardpoolContractShangHaiStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardpoolContractShangHaiStatus represents a ShangHaiStatus event raised by the RewardpoolContract contract.
type RewardpoolContractShangHaiStatus struct {
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterShangHaiStatus is a free log retrieval operation binding the contract event 0x7741ff6843f2be5a10355f0eaa2f114a7532b9748536e5fc2d1e73e836db3028.
//
// Solidity: event ShangHaiStatus(bool status)
func (_RewardpoolContract *RewardpoolContractFilterer) FilterShangHaiStatus(opts *bind.FilterOpts) (*RewardpoolContractShangHaiStatusIterator, error) {

	logs, sub, err := _RewardpoolContract.contract.FilterLogs(opts, "ShangHaiStatus")
	if err != nil {
		return nil, err
	}
	return &RewardpoolContractShangHaiStatusIterator{contract: _RewardpoolContract.contract, event: "ShangHaiStatus", logs: logs, sub: sub}, nil
}

// WatchShangHaiStatus is a free log subscription operation binding the contract event 0x7741ff6843f2be5a10355f0eaa2f114a7532b9748536e5fc2d1e73e836db3028.
//
// Solidity: event ShangHaiStatus(bool status)
func (_RewardpoolContract *RewardpoolContractFilterer) WatchShangHaiStatus(opts *bind.WatchOpts, sink chan<- *RewardpoolContractShangHaiStatus) (event.Subscription, error) {

	logs, sub, err := _RewardpoolContract.contract.WatchLogs(opts, "ShangHaiStatus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardpoolContractShangHaiStatus)
				if err := _RewardpoolContract.contract.UnpackLog(event, "ShangHaiStatus", log); err != nil {
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

// ParseShangHaiStatus is a log parse operation binding the contract event 0x7741ff6843f2be5a10355f0eaa2f114a7532b9748536e5fc2d1e73e836db3028.
//
// Solidity: event ShangHaiStatus(bool status)
func (_RewardpoolContract *RewardpoolContractFilterer) ParseShangHaiStatus(log types.Log) (*RewardpoolContractShangHaiStatus, error) {
	event := new(RewardpoolContractShangHaiStatus)
	if err := _RewardpoolContract.contract.UnpackLog(event, "ShangHaiStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardpoolContractSignerSetIterator is returned from FilterSignerSet and is used to iterate over the raw logs and unpacked data for SignerSet events raised by the RewardpoolContract contract.
type RewardpoolContractSignerSetIterator struct {
	Event *RewardpoolContractSignerSet // Event containing the contract specifics and raw log

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
func (it *RewardpoolContractSignerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardpoolContractSignerSet)
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
		it.Event = new(RewardpoolContractSignerSet)
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
func (it *RewardpoolContractSignerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardpoolContractSignerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardpoolContractSignerSet represents a SignerSet event raised by the RewardpoolContract contract.
type RewardpoolContractSignerSet struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSignerSet is a free log retrieval operation binding the contract event 0x9eaa897564d022fb8c5efaf0acdb5d9d27b440b2aad44400b6e1c702e65b9ed3.
//
// Solidity: event SignerSet(address addr)
func (_RewardpoolContract *RewardpoolContractFilterer) FilterSignerSet(opts *bind.FilterOpts) (*RewardpoolContractSignerSetIterator, error) {

	logs, sub, err := _RewardpoolContract.contract.FilterLogs(opts, "SignerSet")
	if err != nil {
		return nil, err
	}
	return &RewardpoolContractSignerSetIterator{contract: _RewardpoolContract.contract, event: "SignerSet", logs: logs, sub: sub}, nil
}

// WatchSignerSet is a free log subscription operation binding the contract event 0x9eaa897564d022fb8c5efaf0acdb5d9d27b440b2aad44400b6e1c702e65b9ed3.
//
// Solidity: event SignerSet(address addr)
func (_RewardpoolContract *RewardpoolContractFilterer) WatchSignerSet(opts *bind.WatchOpts, sink chan<- *RewardpoolContractSignerSet) (event.Subscription, error) {

	logs, sub, err := _RewardpoolContract.contract.WatchLogs(opts, "SignerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardpoolContractSignerSet)
				if err := _RewardpoolContract.contract.UnpackLog(event, "SignerSet", log); err != nil {
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

// ParseSignerSet is a log parse operation binding the contract event 0x9eaa897564d022fb8c5efaf0acdb5d9d27b440b2aad44400b6e1c702e65b9ed3.
//
// Solidity: event SignerSet(address addr)
func (_RewardpoolContract *RewardpoolContractFilterer) ParseSignerSet(log types.Log) (*RewardpoolContractSignerSet, error) {
	event := new(RewardpoolContractSignerSet)
	if err := _RewardpoolContract.contract.UnpackLog(event, "SignerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardpoolContractStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the RewardpoolContract contract.
type RewardpoolContractStakedIterator struct {
	Event *RewardpoolContractStaked // Event containing the contract specifics and raw log

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
func (it *RewardpoolContractStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardpoolContractStaked)
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
		it.Event = new(RewardpoolContractStaked)
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
func (it *RewardpoolContractStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardpoolContractStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardpoolContractStaked represents a Staked event raised by the RewardpoolContract contract.
type RewardpoolContractStaked struct {
	Addr   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address addr, uint256 amount)
func (_RewardpoolContract *RewardpoolContractFilterer) FilterStaked(opts *bind.FilterOpts) (*RewardpoolContractStakedIterator, error) {

	logs, sub, err := _RewardpoolContract.contract.FilterLogs(opts, "Staked")
	if err != nil {
		return nil, err
	}
	return &RewardpoolContractStakedIterator{contract: _RewardpoolContract.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address addr, uint256 amount)
func (_RewardpoolContract *RewardpoolContractFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *RewardpoolContractStaked) (event.Subscription, error) {

	logs, sub, err := _RewardpoolContract.contract.WatchLogs(opts, "Staked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardpoolContractStaked)
				if err := _RewardpoolContract.contract.UnpackLog(event, "Staked", log); err != nil {
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
func (_RewardpoolContract *RewardpoolContractFilterer) ParseStaked(log types.Log) (*RewardpoolContractStaked, error) {
	event := new(RewardpoolContractStaked)
	if err := _RewardpoolContract.contract.UnpackLog(event, "Staked", log); err != nil {
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
