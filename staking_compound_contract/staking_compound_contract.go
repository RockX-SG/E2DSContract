// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staking_compound_contract

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

// DirectStakingStakeParams is an auto generated low-level Go binding around an user-defined struct.
type DirectStakingStakeParams struct {
	ExtraData    *big.Int
	Amount       *big.Int
	ClaimAddr    common.Address
	WithdrawAddr common.Address
	Pubkeys      [][]byte
	Signatures   [][]byte
	ParamsSig    []byte
}

// StakingCompoundContractMetaData contains all meta data concerning the StakingCompoundContract contract.
var StakingCompoundContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEPOSIT_SIZE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_DEPOSIT_SIZE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batchEmergencyExit\",\"inputs\":[{\"name\":\"validatorIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"exitToClaimAddress\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchExit\",\"inputs\":[{\"name\":\"validatorIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"emergencyExit\",\"inputs\":[{\"name\":\"validatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"exitToClaimAddress\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ethDepositContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"exit\",\"inputs\":[{\"name\":\"validatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getExitQueue\",\"inputs\":[{\"name\":\"from\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExitQueueLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorInfo\",\"inputs\":[{\"name\":\"idx\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"claimAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraData\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorInfos\",\"inputs\":[{\"name\":\"from\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"pubkeys\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"claimAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"extraDatas\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rewardPool\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setRewardPool\",\"inputs\":[{\"name\":\"_rewardPool\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSigner\",\"inputs\":[{\"name\":\"_signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTips\",\"inputs\":[{\"name\":\"_tips\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structDirectStaking.StakeParams\",\"components\":[{\"name\":\"extraData\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pubkeys\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"signatures\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"paramsSig\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sysSigner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tips\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifySigner\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structDirectStaking.StakeParams\",\"components\":[{\"name\":\"extraData\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pubkeys\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"signatures\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"paramsSig\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardPoolContractSet\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignerSet\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Staked\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TipsSet\",\"inputs\":[{\"name\":\"oldTips\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newTips\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// StakingCompoundContractABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingCompoundContractMetaData.ABI instead.
var StakingCompoundContractABI = StakingCompoundContractMetaData.ABI

// StakingCompoundContract is an auto generated Go binding around an Ethereum contract.
type StakingCompoundContract struct {
	StakingCompoundContractCaller     // Read-only binding to the contract
	StakingCompoundContractTransactor // Write-only binding to the contract
	StakingCompoundContractFilterer   // Log filterer for contract events
}

// StakingCompoundContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCompoundContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingCompoundContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingCompoundContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingCompoundContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingCompoundContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingCompoundContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingCompoundContractSession struct {
	Contract     *StakingCompoundContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StakingCompoundContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCompoundContractCallerSession struct {
	Contract *StakingCompoundContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// StakingCompoundContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingCompoundContractTransactorSession struct {
	Contract     *StakingCompoundContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// StakingCompoundContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingCompoundContractRaw struct {
	Contract *StakingCompoundContract // Generic contract binding to access the raw methods on
}

// StakingCompoundContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCompoundContractCallerRaw struct {
	Contract *StakingCompoundContractCaller // Generic read-only contract binding to access the raw methods on
}

// StakingCompoundContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingCompoundContractTransactorRaw struct {
	Contract *StakingCompoundContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingCompoundContract creates a new instance of StakingCompoundContract, bound to a specific deployed contract.
func NewStakingCompoundContract(address common.Address, backend bind.ContractBackend) (*StakingCompoundContract, error) {
	contract, err := bindStakingCompoundContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingCompoundContract{StakingCompoundContractCaller: StakingCompoundContractCaller{contract: contract}, StakingCompoundContractTransactor: StakingCompoundContractTransactor{contract: contract}, StakingCompoundContractFilterer: StakingCompoundContractFilterer{contract: contract}}, nil
}

// NewStakingCompoundContractCaller creates a new read-only instance of StakingCompoundContract, bound to a specific deployed contract.
func NewStakingCompoundContractCaller(address common.Address, caller bind.ContractCaller) (*StakingCompoundContractCaller, error) {
	contract, err := bindStakingCompoundContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCompoundContractCaller{contract: contract}, nil
}

// NewStakingCompoundContractTransactor creates a new write-only instance of StakingCompoundContract, bound to a specific deployed contract.
func NewStakingCompoundContractTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingCompoundContractTransactor, error) {
	contract, err := bindStakingCompoundContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCompoundContractTransactor{contract: contract}, nil
}

// NewStakingCompoundContractFilterer creates a new log filterer instance of StakingCompoundContract, bound to a specific deployed contract.
func NewStakingCompoundContractFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingCompoundContractFilterer, error) {
	contract, err := bindStakingCompoundContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingCompoundContractFilterer{contract: contract}, nil
}

// bindStakingCompoundContract binds a generic wrapper to an already deployed contract.
func bindStakingCompoundContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingCompoundContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingCompoundContract *StakingCompoundContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingCompoundContract.Contract.StakingCompoundContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingCompoundContract *StakingCompoundContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.StakingCompoundContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingCompoundContract *StakingCompoundContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.StakingCompoundContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingCompoundContract *StakingCompoundContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingCompoundContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingCompoundContract *StakingCompoundContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingCompoundContract *StakingCompoundContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakingCompoundContract *StakingCompoundContractCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingCompoundContract.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakingCompoundContract *StakingCompoundContractSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StakingCompoundContract.Contract.DEFAULTADMINROLE(&_StakingCompoundContract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakingCompoundContract *StakingCompoundContractCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StakingCompoundContract.Contract.DEFAULTADMINROLE(&_StakingCompoundContract.CallOpts)
}

// DEPOSITSIZE is a free data retrieval call binding the contract method 0x36bf3325.
//
// Solidity: function DEPOSIT_SIZE() view returns(uint256)
func (_StakingCompoundContract *StakingCompoundContractCaller) DEPOSITSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingCompoundContract.contract.Call(opts, &out, "DEPOSIT_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEPOSITSIZE is a free data retrieval call binding the contract method 0x36bf3325.
//
// Solidity: function DEPOSIT_SIZE() view returns(uint256)
func (_StakingCompoundContract *StakingCompoundContractSession) DEPOSITSIZE() (*big.Int, error) {
	return _StakingCompoundContract.Contract.DEPOSITSIZE(&_StakingCompoundContract.CallOpts)
}

// DEPOSITSIZE is a free data retrieval call binding the contract method 0x36bf3325.
//
// Solidity: function DEPOSIT_SIZE() view returns(uint256)
func (_StakingCompoundContract *StakingCompoundContractCallerSession) DEPOSITSIZE() (*big.Int, error) {
	return _StakingCompoundContract.Contract.DEPOSITSIZE(&_StakingCompoundContract.CallOpts)
}

// MAXDEPOSITSIZE is a free data retrieval call binding the contract method 0x48c0ee2d.
//
// Solidity: function MAX_DEPOSIT_SIZE() view returns(uint256)
func (_StakingCompoundContract *StakingCompoundContractCaller) MAXDEPOSITSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingCompoundContract.contract.Call(opts, &out, "MAX_DEPOSIT_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXDEPOSITSIZE is a free data retrieval call binding the contract method 0x48c0ee2d.
//
// Solidity: function MAX_DEPOSIT_SIZE() view returns(uint256)
func (_StakingCompoundContract *StakingCompoundContractSession) MAXDEPOSITSIZE() (*big.Int, error) {
	return _StakingCompoundContract.Contract.MAXDEPOSITSIZE(&_StakingCompoundContract.CallOpts)
}

// MAXDEPOSITSIZE is a free data retrieval call binding the contract method 0x48c0ee2d.
//
// Solidity: function MAX_DEPOSIT_SIZE() view returns(uint256)
func (_StakingCompoundContract *StakingCompoundContractCallerSession) MAXDEPOSITSIZE() (*big.Int, error) {
	return _StakingCompoundContract.Contract.MAXDEPOSITSIZE(&_StakingCompoundContract.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_StakingCompoundContract *StakingCompoundContractCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingCompoundContract.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_StakingCompoundContract *StakingCompoundContractSession) PAUSERROLE() ([32]byte, error) {
	return _StakingCompoundContract.Contract.PAUSERROLE(&_StakingCompoundContract.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_StakingCompoundContract *StakingCompoundContractCallerSession) PAUSERROLE() ([32]byte, error) {
	return _StakingCompoundContract.Contract.PAUSERROLE(&_StakingCompoundContract.CallOpts)
}

// EthDepositContract is a free data retrieval call binding the contract method 0x3884545d.
//
// Solidity: function ethDepositContract() view returns(address)
func (_StakingCompoundContract *StakingCompoundContractCaller) EthDepositContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingCompoundContract.contract.Call(opts, &out, "ethDepositContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthDepositContract is a free data retrieval call binding the contract method 0x3884545d.
//
// Solidity: function ethDepositContract() view returns(address)
func (_StakingCompoundContract *StakingCompoundContractSession) EthDepositContract() (common.Address, error) {
	return _StakingCompoundContract.Contract.EthDepositContract(&_StakingCompoundContract.CallOpts)
}

// EthDepositContract is a free data retrieval call binding the contract method 0x3884545d.
//
// Solidity: function ethDepositContract() view returns(address)
func (_StakingCompoundContract *StakingCompoundContractCallerSession) EthDepositContract() (common.Address, error) {
	return _StakingCompoundContract.Contract.EthDepositContract(&_StakingCompoundContract.CallOpts)
}

// GetExitQueue is a free data retrieval call binding the contract method 0x43931256.
//
// Solidity: function getExitQueue(uint256 from, uint256 to) view returns(uint256[])
func (_StakingCompoundContract *StakingCompoundContractCaller) GetExitQueue(opts *bind.CallOpts, from *big.Int, to *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _StakingCompoundContract.contract.Call(opts, &out, "getExitQueue", from, to)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetExitQueue is a free data retrieval call binding the contract method 0x43931256.
//
// Solidity: function getExitQueue(uint256 from, uint256 to) view returns(uint256[])
func (_StakingCompoundContract *StakingCompoundContractSession) GetExitQueue(from *big.Int, to *big.Int) ([]*big.Int, error) {
	return _StakingCompoundContract.Contract.GetExitQueue(&_StakingCompoundContract.CallOpts, from, to)
}

// GetExitQueue is a free data retrieval call binding the contract method 0x43931256.
//
// Solidity: function getExitQueue(uint256 from, uint256 to) view returns(uint256[])
func (_StakingCompoundContract *StakingCompoundContractCallerSession) GetExitQueue(from *big.Int, to *big.Int) ([]*big.Int, error) {
	return _StakingCompoundContract.Contract.GetExitQueue(&_StakingCompoundContract.CallOpts, from, to)
}

// GetExitQueueLength is a free data retrieval call binding the contract method 0x559230a0.
//
// Solidity: function getExitQueueLength() view returns(uint256)
func (_StakingCompoundContract *StakingCompoundContractCaller) GetExitQueueLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingCompoundContract.contract.Call(opts, &out, "getExitQueueLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExitQueueLength is a free data retrieval call binding the contract method 0x559230a0.
//
// Solidity: function getExitQueueLength() view returns(uint256)
func (_StakingCompoundContract *StakingCompoundContractSession) GetExitQueueLength() (*big.Int, error) {
	return _StakingCompoundContract.Contract.GetExitQueueLength(&_StakingCompoundContract.CallOpts)
}

// GetExitQueueLength is a free data retrieval call binding the contract method 0x559230a0.
//
// Solidity: function getExitQueueLength() view returns(uint256)
func (_StakingCompoundContract *StakingCompoundContractCallerSession) GetExitQueueLength() (*big.Int, error) {
	return _StakingCompoundContract.Contract.GetExitQueueLength(&_StakingCompoundContract.CallOpts)
}

// GetNextValidators is a free data retrieval call binding the contract method 0x40cddab3.
//
// Solidity: function getNextValidators() view returns(uint256)
func (_StakingCompoundContract *StakingCompoundContractCaller) GetNextValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingCompoundContract.contract.Call(opts, &out, "getNextValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextValidators is a free data retrieval call binding the contract method 0x40cddab3.
//
// Solidity: function getNextValidators() view returns(uint256)
func (_StakingCompoundContract *StakingCompoundContractSession) GetNextValidators() (*big.Int, error) {
	return _StakingCompoundContract.Contract.GetNextValidators(&_StakingCompoundContract.CallOpts)
}

// GetNextValidators is a free data retrieval call binding the contract method 0x40cddab3.
//
// Solidity: function getNextValidators() view returns(uint256)
func (_StakingCompoundContract *StakingCompoundContractCallerSession) GetNextValidators() (*big.Int, error) {
	return _StakingCompoundContract.Contract.GetNextValidators(&_StakingCompoundContract.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakingCompoundContract *StakingCompoundContractCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _StakingCompoundContract.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakingCompoundContract *StakingCompoundContractSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StakingCompoundContract.Contract.GetRoleAdmin(&_StakingCompoundContract.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakingCompoundContract *StakingCompoundContractCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StakingCompoundContract.Contract.GetRoleAdmin(&_StakingCompoundContract.CallOpts, role)
}

// GetValidatorInfo is a free data retrieval call binding the contract method 0xb7797537.
//
// Solidity: function getValidatorInfo(uint256 idx) view returns(bytes pubkey, address claimAddress, uint256 extraData)
func (_StakingCompoundContract *StakingCompoundContractCaller) GetValidatorInfo(opts *bind.CallOpts, idx *big.Int) (struct {
	Pubkey       []byte
	ClaimAddress common.Address
	ExtraData    *big.Int
}, error) {
	var out []interface{}
	err := _StakingCompoundContract.contract.Call(opts, &out, "getValidatorInfo", idx)

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
func (_StakingCompoundContract *StakingCompoundContractSession) GetValidatorInfo(idx *big.Int) (struct {
	Pubkey       []byte
	ClaimAddress common.Address
	ExtraData    *big.Int
}, error) {
	return _StakingCompoundContract.Contract.GetValidatorInfo(&_StakingCompoundContract.CallOpts, idx)
}

// GetValidatorInfo is a free data retrieval call binding the contract method 0xb7797537.
//
// Solidity: function getValidatorInfo(uint256 idx) view returns(bytes pubkey, address claimAddress, uint256 extraData)
func (_StakingCompoundContract *StakingCompoundContractCallerSession) GetValidatorInfo(idx *big.Int) (struct {
	Pubkey       []byte
	ClaimAddress common.Address
	ExtraData    *big.Int
}, error) {
	return _StakingCompoundContract.Contract.GetValidatorInfo(&_StakingCompoundContract.CallOpts, idx)
}

// GetValidatorInfos is a free data retrieval call binding the contract method 0x9e054533.
//
// Solidity: function getValidatorInfos(uint256 from, uint256 to) view returns(bytes[] pubkeys, address[] claimAddresses, uint256[] extraDatas)
func (_StakingCompoundContract *StakingCompoundContractCaller) GetValidatorInfos(opts *bind.CallOpts, from *big.Int, to *big.Int) (struct {
	Pubkeys        [][]byte
	ClaimAddresses []common.Address
	ExtraDatas     []*big.Int
}, error) {
	var out []interface{}
	err := _StakingCompoundContract.contract.Call(opts, &out, "getValidatorInfos", from, to)

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
func (_StakingCompoundContract *StakingCompoundContractSession) GetValidatorInfos(from *big.Int, to *big.Int) (struct {
	Pubkeys        [][]byte
	ClaimAddresses []common.Address
	ExtraDatas     []*big.Int
}, error) {
	return _StakingCompoundContract.Contract.GetValidatorInfos(&_StakingCompoundContract.CallOpts, from, to)
}

// GetValidatorInfos is a free data retrieval call binding the contract method 0x9e054533.
//
// Solidity: function getValidatorInfos(uint256 from, uint256 to) view returns(bytes[] pubkeys, address[] claimAddresses, uint256[] extraDatas)
func (_StakingCompoundContract *StakingCompoundContractCallerSession) GetValidatorInfos(from *big.Int, to *big.Int) (struct {
	Pubkeys        [][]byte
	ClaimAddresses []common.Address
	ExtraDatas     []*big.Int
}, error) {
	return _StakingCompoundContract.Contract.GetValidatorInfos(&_StakingCompoundContract.CallOpts, from, to)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakingCompoundContract *StakingCompoundContractCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _StakingCompoundContract.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakingCompoundContract *StakingCompoundContractSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StakingCompoundContract.Contract.HasRole(&_StakingCompoundContract.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakingCompoundContract *StakingCompoundContractCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StakingCompoundContract.Contract.HasRole(&_StakingCompoundContract.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakingCompoundContract *StakingCompoundContractCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingCompoundContract.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakingCompoundContract *StakingCompoundContractSession) Paused() (bool, error) {
	return _StakingCompoundContract.Contract.Paused(&_StakingCompoundContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakingCompoundContract *StakingCompoundContractCallerSession) Paused() (bool, error) {
	return _StakingCompoundContract.Contract.Paused(&_StakingCompoundContract.CallOpts)
}

// RewardPool is a free data retrieval call binding the contract method 0x66666aa9.
//
// Solidity: function rewardPool() view returns(address)
func (_StakingCompoundContract *StakingCompoundContractCaller) RewardPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingCompoundContract.contract.Call(opts, &out, "rewardPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardPool is a free data retrieval call binding the contract method 0x66666aa9.
//
// Solidity: function rewardPool() view returns(address)
func (_StakingCompoundContract *StakingCompoundContractSession) RewardPool() (common.Address, error) {
	return _StakingCompoundContract.Contract.RewardPool(&_StakingCompoundContract.CallOpts)
}

// RewardPool is a free data retrieval call binding the contract method 0x66666aa9.
//
// Solidity: function rewardPool() view returns(address)
func (_StakingCompoundContract *StakingCompoundContractCallerSession) RewardPool() (common.Address, error) {
	return _StakingCompoundContract.Contract.RewardPool(&_StakingCompoundContract.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakingCompoundContract *StakingCompoundContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _StakingCompoundContract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakingCompoundContract *StakingCompoundContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StakingCompoundContract.Contract.SupportsInterface(&_StakingCompoundContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakingCompoundContract *StakingCompoundContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StakingCompoundContract.Contract.SupportsInterface(&_StakingCompoundContract.CallOpts, interfaceId)
}

// SysSigner is a free data retrieval call binding the contract method 0xfa734b1c.
//
// Solidity: function sysSigner() view returns(address)
func (_StakingCompoundContract *StakingCompoundContractCaller) SysSigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingCompoundContract.contract.Call(opts, &out, "sysSigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SysSigner is a free data retrieval call binding the contract method 0xfa734b1c.
//
// Solidity: function sysSigner() view returns(address)
func (_StakingCompoundContract *StakingCompoundContractSession) SysSigner() (common.Address, error) {
	return _StakingCompoundContract.Contract.SysSigner(&_StakingCompoundContract.CallOpts)
}

// SysSigner is a free data retrieval call binding the contract method 0xfa734b1c.
//
// Solidity: function sysSigner() view returns(address)
func (_StakingCompoundContract *StakingCompoundContractCallerSession) SysSigner() (common.Address, error) {
	return _StakingCompoundContract.Contract.SysSigner(&_StakingCompoundContract.CallOpts)
}

// Tips is a free data retrieval call binding the contract method 0x1a4e1e78.
//
// Solidity: function tips() view returns(uint256)
func (_StakingCompoundContract *StakingCompoundContractCaller) Tips(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingCompoundContract.contract.Call(opts, &out, "tips")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tips is a free data retrieval call binding the contract method 0x1a4e1e78.
//
// Solidity: function tips() view returns(uint256)
func (_StakingCompoundContract *StakingCompoundContractSession) Tips() (*big.Int, error) {
	return _StakingCompoundContract.Contract.Tips(&_StakingCompoundContract.CallOpts)
}

// Tips is a free data retrieval call binding the contract method 0x1a4e1e78.
//
// Solidity: function tips() view returns(uint256)
func (_StakingCompoundContract *StakingCompoundContractCallerSession) Tips() (*big.Int, error) {
	return _StakingCompoundContract.Contract.Tips(&_StakingCompoundContract.CallOpts)
}

// VerifySigner is a free data retrieval call binding the contract method 0xb73bfdd2.
//
// Solidity: function verifySigner((uint256,uint256,address,address,bytes[],bytes[],bytes) params) view returns(bool)
func (_StakingCompoundContract *StakingCompoundContractCaller) VerifySigner(opts *bind.CallOpts, params DirectStakingStakeParams) (bool, error) {
	var out []interface{}
	err := _StakingCompoundContract.contract.Call(opts, &out, "verifySigner", params)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifySigner is a free data retrieval call binding the contract method 0xb73bfdd2.
//
// Solidity: function verifySigner((uint256,uint256,address,address,bytes[],bytes[],bytes) params) view returns(bool)
func (_StakingCompoundContract *StakingCompoundContractSession) VerifySigner(params DirectStakingStakeParams) (bool, error) {
	return _StakingCompoundContract.Contract.VerifySigner(&_StakingCompoundContract.CallOpts, params)
}

// VerifySigner is a free data retrieval call binding the contract method 0xb73bfdd2.
//
// Solidity: function verifySigner((uint256,uint256,address,address,bytes[],bytes[],bytes) params) view returns(bool)
func (_StakingCompoundContract *StakingCompoundContractCallerSession) VerifySigner(params DirectStakingStakeParams) (bool, error) {
	return _StakingCompoundContract.Contract.VerifySigner(&_StakingCompoundContract.CallOpts, params)
}

// BatchEmergencyExit is a paid mutator transaction binding the contract method 0x2393df84.
//
// Solidity: function batchEmergencyExit(uint256[] validatorIds, bool exitToClaimAddress) returns()
func (_StakingCompoundContract *StakingCompoundContractTransactor) BatchEmergencyExit(opts *bind.TransactOpts, validatorIds []*big.Int, exitToClaimAddress bool) (*types.Transaction, error) {
	return _StakingCompoundContract.contract.Transact(opts, "batchEmergencyExit", validatorIds, exitToClaimAddress)
}

// BatchEmergencyExit is a paid mutator transaction binding the contract method 0x2393df84.
//
// Solidity: function batchEmergencyExit(uint256[] validatorIds, bool exitToClaimAddress) returns()
func (_StakingCompoundContract *StakingCompoundContractSession) BatchEmergencyExit(validatorIds []*big.Int, exitToClaimAddress bool) (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.BatchEmergencyExit(&_StakingCompoundContract.TransactOpts, validatorIds, exitToClaimAddress)
}

// BatchEmergencyExit is a paid mutator transaction binding the contract method 0x2393df84.
//
// Solidity: function batchEmergencyExit(uint256[] validatorIds, bool exitToClaimAddress) returns()
func (_StakingCompoundContract *StakingCompoundContractTransactorSession) BatchEmergencyExit(validatorIds []*big.Int, exitToClaimAddress bool) (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.BatchEmergencyExit(&_StakingCompoundContract.TransactOpts, validatorIds, exitToClaimAddress)
}

// BatchExit is a paid mutator transaction binding the contract method 0x4dd1fb0b.
//
// Solidity: function batchExit(uint256[] validatorIds) returns()
func (_StakingCompoundContract *StakingCompoundContractTransactor) BatchExit(opts *bind.TransactOpts, validatorIds []*big.Int) (*types.Transaction, error) {
	return _StakingCompoundContract.contract.Transact(opts, "batchExit", validatorIds)
}

// BatchExit is a paid mutator transaction binding the contract method 0x4dd1fb0b.
//
// Solidity: function batchExit(uint256[] validatorIds) returns()
func (_StakingCompoundContract *StakingCompoundContractSession) BatchExit(validatorIds []*big.Int) (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.BatchExit(&_StakingCompoundContract.TransactOpts, validatorIds)
}

// BatchExit is a paid mutator transaction binding the contract method 0x4dd1fb0b.
//
// Solidity: function batchExit(uint256[] validatorIds) returns()
func (_StakingCompoundContract *StakingCompoundContractTransactorSession) BatchExit(validatorIds []*big.Int) (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.BatchExit(&_StakingCompoundContract.TransactOpts, validatorIds)
}

// EmergencyExit is a paid mutator transaction binding the contract method 0xf806e87a.
//
// Solidity: function emergencyExit(uint256 validatorId, bool exitToClaimAddress) returns()
func (_StakingCompoundContract *StakingCompoundContractTransactor) EmergencyExit(opts *bind.TransactOpts, validatorId *big.Int, exitToClaimAddress bool) (*types.Transaction, error) {
	return _StakingCompoundContract.contract.Transact(opts, "emergencyExit", validatorId, exitToClaimAddress)
}

// EmergencyExit is a paid mutator transaction binding the contract method 0xf806e87a.
//
// Solidity: function emergencyExit(uint256 validatorId, bool exitToClaimAddress) returns()
func (_StakingCompoundContract *StakingCompoundContractSession) EmergencyExit(validatorId *big.Int, exitToClaimAddress bool) (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.EmergencyExit(&_StakingCompoundContract.TransactOpts, validatorId, exitToClaimAddress)
}

// EmergencyExit is a paid mutator transaction binding the contract method 0xf806e87a.
//
// Solidity: function emergencyExit(uint256 validatorId, bool exitToClaimAddress) returns()
func (_StakingCompoundContract *StakingCompoundContractTransactorSession) EmergencyExit(validatorId *big.Int, exitToClaimAddress bool) (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.EmergencyExit(&_StakingCompoundContract.TransactOpts, validatorId, exitToClaimAddress)
}

// Exit is a paid mutator transaction binding the contract method 0x7f8661a1.
//
// Solidity: function exit(uint256 validatorId) returns()
func (_StakingCompoundContract *StakingCompoundContractTransactor) Exit(opts *bind.TransactOpts, validatorId *big.Int) (*types.Transaction, error) {
	return _StakingCompoundContract.contract.Transact(opts, "exit", validatorId)
}

// Exit is a paid mutator transaction binding the contract method 0x7f8661a1.
//
// Solidity: function exit(uint256 validatorId) returns()
func (_StakingCompoundContract *StakingCompoundContractSession) Exit(validatorId *big.Int) (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.Exit(&_StakingCompoundContract.TransactOpts, validatorId)
}

// Exit is a paid mutator transaction binding the contract method 0x7f8661a1.
//
// Solidity: function exit(uint256 validatorId) returns()
func (_StakingCompoundContract *StakingCompoundContractTransactorSession) Exit(validatorId *big.Int) (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.Exit(&_StakingCompoundContract.TransactOpts, validatorId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakingCompoundContract *StakingCompoundContractTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingCompoundContract.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakingCompoundContract *StakingCompoundContractSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.GrantRole(&_StakingCompoundContract.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakingCompoundContract *StakingCompoundContractTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.GrantRole(&_StakingCompoundContract.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_StakingCompoundContract *StakingCompoundContractTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingCompoundContract.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_StakingCompoundContract *StakingCompoundContractSession) Initialize() (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.Initialize(&_StakingCompoundContract.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_StakingCompoundContract *StakingCompoundContractTransactorSession) Initialize() (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.Initialize(&_StakingCompoundContract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StakingCompoundContract *StakingCompoundContractTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingCompoundContract.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StakingCompoundContract *StakingCompoundContractSession) Pause() (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.Pause(&_StakingCompoundContract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StakingCompoundContract *StakingCompoundContractTransactorSession) Pause() (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.Pause(&_StakingCompoundContract.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StakingCompoundContract *StakingCompoundContractTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingCompoundContract.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StakingCompoundContract *StakingCompoundContractSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.RenounceRole(&_StakingCompoundContract.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StakingCompoundContract *StakingCompoundContractTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.RenounceRole(&_StakingCompoundContract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakingCompoundContract *StakingCompoundContractTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingCompoundContract.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakingCompoundContract *StakingCompoundContractSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.RevokeRole(&_StakingCompoundContract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakingCompoundContract *StakingCompoundContractTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.RevokeRole(&_StakingCompoundContract.TransactOpts, role, account)
}

// SetRewardPool is a paid mutator transaction binding the contract method 0x78238c37.
//
// Solidity: function setRewardPool(address _rewardPool) returns()
func (_StakingCompoundContract *StakingCompoundContractTransactor) SetRewardPool(opts *bind.TransactOpts, _rewardPool common.Address) (*types.Transaction, error) {
	return _StakingCompoundContract.contract.Transact(opts, "setRewardPool", _rewardPool)
}

// SetRewardPool is a paid mutator transaction binding the contract method 0x78238c37.
//
// Solidity: function setRewardPool(address _rewardPool) returns()
func (_StakingCompoundContract *StakingCompoundContractSession) SetRewardPool(_rewardPool common.Address) (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.SetRewardPool(&_StakingCompoundContract.TransactOpts, _rewardPool)
}

// SetRewardPool is a paid mutator transaction binding the contract method 0x78238c37.
//
// Solidity: function setRewardPool(address _rewardPool) returns()
func (_StakingCompoundContract *StakingCompoundContractTransactorSession) SetRewardPool(_rewardPool common.Address) (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.SetRewardPool(&_StakingCompoundContract.TransactOpts, _rewardPool)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address _signer) returns()
func (_StakingCompoundContract *StakingCompoundContractTransactor) SetSigner(opts *bind.TransactOpts, _signer common.Address) (*types.Transaction, error) {
	return _StakingCompoundContract.contract.Transact(opts, "setSigner", _signer)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address _signer) returns()
func (_StakingCompoundContract *StakingCompoundContractSession) SetSigner(_signer common.Address) (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.SetSigner(&_StakingCompoundContract.TransactOpts, _signer)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address _signer) returns()
func (_StakingCompoundContract *StakingCompoundContractTransactorSession) SetSigner(_signer common.Address) (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.SetSigner(&_StakingCompoundContract.TransactOpts, _signer)
}

// SetTips is a paid mutator transaction binding the contract method 0xcc90a64f.
//
// Solidity: function setTips(uint256 _tips) returns()
func (_StakingCompoundContract *StakingCompoundContractTransactor) SetTips(opts *bind.TransactOpts, _tips *big.Int) (*types.Transaction, error) {
	return _StakingCompoundContract.contract.Transact(opts, "setTips", _tips)
}

// SetTips is a paid mutator transaction binding the contract method 0xcc90a64f.
//
// Solidity: function setTips(uint256 _tips) returns()
func (_StakingCompoundContract *StakingCompoundContractSession) SetTips(_tips *big.Int) (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.SetTips(&_StakingCompoundContract.TransactOpts, _tips)
}

// SetTips is a paid mutator transaction binding the contract method 0xcc90a64f.
//
// Solidity: function setTips(uint256 _tips) returns()
func (_StakingCompoundContract *StakingCompoundContractTransactorSession) SetTips(_tips *big.Int) (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.SetTips(&_StakingCompoundContract.TransactOpts, _tips)
}

// Stake is a paid mutator transaction binding the contract method 0xaa3a4e66.
//
// Solidity: function stake((uint256,uint256,address,address,bytes[],bytes[],bytes) params) payable returns()
func (_StakingCompoundContract *StakingCompoundContractTransactor) Stake(opts *bind.TransactOpts, params DirectStakingStakeParams) (*types.Transaction, error) {
	return _StakingCompoundContract.contract.Transact(opts, "stake", params)
}

// Stake is a paid mutator transaction binding the contract method 0xaa3a4e66.
//
// Solidity: function stake((uint256,uint256,address,address,bytes[],bytes[],bytes) params) payable returns()
func (_StakingCompoundContract *StakingCompoundContractSession) Stake(params DirectStakingStakeParams) (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.Stake(&_StakingCompoundContract.TransactOpts, params)
}

// Stake is a paid mutator transaction binding the contract method 0xaa3a4e66.
//
// Solidity: function stake((uint256,uint256,address,address,bytes[],bytes[],bytes) params) payable returns()
func (_StakingCompoundContract *StakingCompoundContractTransactorSession) Stake(params DirectStakingStakeParams) (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.Stake(&_StakingCompoundContract.TransactOpts, params)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StakingCompoundContract *StakingCompoundContractTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingCompoundContract.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StakingCompoundContract *StakingCompoundContractSession) Unpause() (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.Unpause(&_StakingCompoundContract.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StakingCompoundContract *StakingCompoundContractTransactorSession) Unpause() (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.Unpause(&_StakingCompoundContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingCompoundContract *StakingCompoundContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingCompoundContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingCompoundContract *StakingCompoundContractSession) Receive() (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.Receive(&_StakingCompoundContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingCompoundContract *StakingCompoundContractTransactorSession) Receive() (*types.Transaction, error) {
	return _StakingCompoundContract.Contract.Receive(&_StakingCompoundContract.TransactOpts)
}

// StakingCompoundContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StakingCompoundContract contract.
type StakingCompoundContractInitializedIterator struct {
	Event *StakingCompoundContractInitialized // Event containing the contract specifics and raw log

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
func (it *StakingCompoundContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingCompoundContractInitialized)
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
		it.Event = new(StakingCompoundContractInitialized)
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
func (it *StakingCompoundContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingCompoundContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingCompoundContractInitialized represents a Initialized event raised by the StakingCompoundContract contract.
type StakingCompoundContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StakingCompoundContract *StakingCompoundContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakingCompoundContractInitializedIterator, error) {

	logs, sub, err := _StakingCompoundContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakingCompoundContractInitializedIterator{contract: _StakingCompoundContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StakingCompoundContract *StakingCompoundContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakingCompoundContractInitialized) (event.Subscription, error) {

	logs, sub, err := _StakingCompoundContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingCompoundContractInitialized)
				if err := _StakingCompoundContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_StakingCompoundContract *StakingCompoundContractFilterer) ParseInitialized(log types.Log) (*StakingCompoundContractInitialized, error) {
	event := new(StakingCompoundContractInitialized)
	if err := _StakingCompoundContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingCompoundContractPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the StakingCompoundContract contract.
type StakingCompoundContractPausedIterator struct {
	Event *StakingCompoundContractPaused // Event containing the contract specifics and raw log

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
func (it *StakingCompoundContractPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingCompoundContractPaused)
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
		it.Event = new(StakingCompoundContractPaused)
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
func (it *StakingCompoundContractPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingCompoundContractPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingCompoundContractPaused represents a Paused event raised by the StakingCompoundContract contract.
type StakingCompoundContractPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StakingCompoundContract *StakingCompoundContractFilterer) FilterPaused(opts *bind.FilterOpts) (*StakingCompoundContractPausedIterator, error) {

	logs, sub, err := _StakingCompoundContract.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StakingCompoundContractPausedIterator{contract: _StakingCompoundContract.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StakingCompoundContract *StakingCompoundContractFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StakingCompoundContractPaused) (event.Subscription, error) {

	logs, sub, err := _StakingCompoundContract.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingCompoundContractPaused)
				if err := _StakingCompoundContract.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_StakingCompoundContract *StakingCompoundContractFilterer) ParsePaused(log types.Log) (*StakingCompoundContractPaused, error) {
	event := new(StakingCompoundContractPaused)
	if err := _StakingCompoundContract.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingCompoundContractRewardPoolContractSetIterator is returned from FilterRewardPoolContractSet and is used to iterate over the raw logs and unpacked data for RewardPoolContractSet events raised by the StakingCompoundContract contract.
type StakingCompoundContractRewardPoolContractSetIterator struct {
	Event *StakingCompoundContractRewardPoolContractSet // Event containing the contract specifics and raw log

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
func (it *StakingCompoundContractRewardPoolContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingCompoundContractRewardPoolContractSet)
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
		it.Event = new(StakingCompoundContractRewardPoolContractSet)
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
func (it *StakingCompoundContractRewardPoolContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingCompoundContractRewardPoolContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingCompoundContractRewardPoolContractSet represents a RewardPoolContractSet event raised by the StakingCompoundContract contract.
type StakingCompoundContractRewardPoolContractSet struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRewardPoolContractSet is a free log retrieval operation binding the contract event 0x7e2602838378170d01002df4f80426d046e484841f67339edd3213b1875b1a89.
//
// Solidity: event RewardPoolContractSet(address addr)
func (_StakingCompoundContract *StakingCompoundContractFilterer) FilterRewardPoolContractSet(opts *bind.FilterOpts) (*StakingCompoundContractRewardPoolContractSetIterator, error) {

	logs, sub, err := _StakingCompoundContract.contract.FilterLogs(opts, "RewardPoolContractSet")
	if err != nil {
		return nil, err
	}
	return &StakingCompoundContractRewardPoolContractSetIterator{contract: _StakingCompoundContract.contract, event: "RewardPoolContractSet", logs: logs, sub: sub}, nil
}

// WatchRewardPoolContractSet is a free log subscription operation binding the contract event 0x7e2602838378170d01002df4f80426d046e484841f67339edd3213b1875b1a89.
//
// Solidity: event RewardPoolContractSet(address addr)
func (_StakingCompoundContract *StakingCompoundContractFilterer) WatchRewardPoolContractSet(opts *bind.WatchOpts, sink chan<- *StakingCompoundContractRewardPoolContractSet) (event.Subscription, error) {

	logs, sub, err := _StakingCompoundContract.contract.WatchLogs(opts, "RewardPoolContractSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingCompoundContractRewardPoolContractSet)
				if err := _StakingCompoundContract.contract.UnpackLog(event, "RewardPoolContractSet", log); err != nil {
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
func (_StakingCompoundContract *StakingCompoundContractFilterer) ParseRewardPoolContractSet(log types.Log) (*StakingCompoundContractRewardPoolContractSet, error) {
	event := new(StakingCompoundContractRewardPoolContractSet)
	if err := _StakingCompoundContract.contract.UnpackLog(event, "RewardPoolContractSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingCompoundContractRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the StakingCompoundContract contract.
type StakingCompoundContractRoleAdminChangedIterator struct {
	Event *StakingCompoundContractRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StakingCompoundContractRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingCompoundContractRoleAdminChanged)
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
		it.Event = new(StakingCompoundContractRoleAdminChanged)
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
func (it *StakingCompoundContractRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingCompoundContractRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingCompoundContractRoleAdminChanged represents a RoleAdminChanged event raised by the StakingCompoundContract contract.
type StakingCompoundContractRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakingCompoundContract *StakingCompoundContractFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StakingCompoundContractRoleAdminChangedIterator, error) {

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

	logs, sub, err := _StakingCompoundContract.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StakingCompoundContractRoleAdminChangedIterator{contract: _StakingCompoundContract.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakingCompoundContract *StakingCompoundContractFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StakingCompoundContractRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _StakingCompoundContract.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingCompoundContractRoleAdminChanged)
				if err := _StakingCompoundContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_StakingCompoundContract *StakingCompoundContractFilterer) ParseRoleAdminChanged(log types.Log) (*StakingCompoundContractRoleAdminChanged, error) {
	event := new(StakingCompoundContractRoleAdminChanged)
	if err := _StakingCompoundContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingCompoundContractRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the StakingCompoundContract contract.
type StakingCompoundContractRoleGrantedIterator struct {
	Event *StakingCompoundContractRoleGranted // Event containing the contract specifics and raw log

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
func (it *StakingCompoundContractRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingCompoundContractRoleGranted)
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
		it.Event = new(StakingCompoundContractRoleGranted)
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
func (it *StakingCompoundContractRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingCompoundContractRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingCompoundContractRoleGranted represents a RoleGranted event raised by the StakingCompoundContract contract.
type StakingCompoundContractRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingCompoundContract *StakingCompoundContractFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakingCompoundContractRoleGrantedIterator, error) {

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

	logs, sub, err := _StakingCompoundContract.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakingCompoundContractRoleGrantedIterator{contract: _StakingCompoundContract.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingCompoundContract *StakingCompoundContractFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StakingCompoundContractRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StakingCompoundContract.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingCompoundContractRoleGranted)
				if err := _StakingCompoundContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_StakingCompoundContract *StakingCompoundContractFilterer) ParseRoleGranted(log types.Log) (*StakingCompoundContractRoleGranted, error) {
	event := new(StakingCompoundContractRoleGranted)
	if err := _StakingCompoundContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingCompoundContractRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the StakingCompoundContract contract.
type StakingCompoundContractRoleRevokedIterator struct {
	Event *StakingCompoundContractRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StakingCompoundContractRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingCompoundContractRoleRevoked)
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
		it.Event = new(StakingCompoundContractRoleRevoked)
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
func (it *StakingCompoundContractRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingCompoundContractRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingCompoundContractRoleRevoked represents a RoleRevoked event raised by the StakingCompoundContract contract.
type StakingCompoundContractRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingCompoundContract *StakingCompoundContractFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakingCompoundContractRoleRevokedIterator, error) {

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

	logs, sub, err := _StakingCompoundContract.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakingCompoundContractRoleRevokedIterator{contract: _StakingCompoundContract.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingCompoundContract *StakingCompoundContractFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StakingCompoundContractRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StakingCompoundContract.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingCompoundContractRoleRevoked)
				if err := _StakingCompoundContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_StakingCompoundContract *StakingCompoundContractFilterer) ParseRoleRevoked(log types.Log) (*StakingCompoundContractRoleRevoked, error) {
	event := new(StakingCompoundContractRoleRevoked)
	if err := _StakingCompoundContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingCompoundContractSignerSetIterator is returned from FilterSignerSet and is used to iterate over the raw logs and unpacked data for SignerSet events raised by the StakingCompoundContract contract.
type StakingCompoundContractSignerSetIterator struct {
	Event *StakingCompoundContractSignerSet // Event containing the contract specifics and raw log

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
func (it *StakingCompoundContractSignerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingCompoundContractSignerSet)
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
		it.Event = new(StakingCompoundContractSignerSet)
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
func (it *StakingCompoundContractSignerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingCompoundContractSignerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingCompoundContractSignerSet represents a SignerSet event raised by the StakingCompoundContract contract.
type StakingCompoundContractSignerSet struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSignerSet is a free log retrieval operation binding the contract event 0x9eaa897564d022fb8c5efaf0acdb5d9d27b440b2aad44400b6e1c702e65b9ed3.
//
// Solidity: event SignerSet(address addr)
func (_StakingCompoundContract *StakingCompoundContractFilterer) FilterSignerSet(opts *bind.FilterOpts) (*StakingCompoundContractSignerSetIterator, error) {

	logs, sub, err := _StakingCompoundContract.contract.FilterLogs(opts, "SignerSet")
	if err != nil {
		return nil, err
	}
	return &StakingCompoundContractSignerSetIterator{contract: _StakingCompoundContract.contract, event: "SignerSet", logs: logs, sub: sub}, nil
}

// WatchSignerSet is a free log subscription operation binding the contract event 0x9eaa897564d022fb8c5efaf0acdb5d9d27b440b2aad44400b6e1c702e65b9ed3.
//
// Solidity: event SignerSet(address addr)
func (_StakingCompoundContract *StakingCompoundContractFilterer) WatchSignerSet(opts *bind.WatchOpts, sink chan<- *StakingCompoundContractSignerSet) (event.Subscription, error) {

	logs, sub, err := _StakingCompoundContract.contract.WatchLogs(opts, "SignerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingCompoundContractSignerSet)
				if err := _StakingCompoundContract.contract.UnpackLog(event, "SignerSet", log); err != nil {
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
func (_StakingCompoundContract *StakingCompoundContractFilterer) ParseSignerSet(log types.Log) (*StakingCompoundContractSignerSet, error) {
	event := new(StakingCompoundContractSignerSet)
	if err := _StakingCompoundContract.contract.UnpackLog(event, "SignerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingCompoundContractStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the StakingCompoundContract contract.
type StakingCompoundContractStakedIterator struct {
	Event *StakingCompoundContractStaked // Event containing the contract specifics and raw log

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
func (it *StakingCompoundContractStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingCompoundContractStaked)
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
		it.Event = new(StakingCompoundContractStaked)
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
func (it *StakingCompoundContractStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingCompoundContractStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingCompoundContractStaked represents a Staked event raised by the StakingCompoundContract contract.
type StakingCompoundContractStaked struct {
	Addr   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address addr, uint256 amount)
func (_StakingCompoundContract *StakingCompoundContractFilterer) FilterStaked(opts *bind.FilterOpts) (*StakingCompoundContractStakedIterator, error) {

	logs, sub, err := _StakingCompoundContract.contract.FilterLogs(opts, "Staked")
	if err != nil {
		return nil, err
	}
	return &StakingCompoundContractStakedIterator{contract: _StakingCompoundContract.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address addr, uint256 amount)
func (_StakingCompoundContract *StakingCompoundContractFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *StakingCompoundContractStaked) (event.Subscription, error) {

	logs, sub, err := _StakingCompoundContract.contract.WatchLogs(opts, "Staked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingCompoundContractStaked)
				if err := _StakingCompoundContract.contract.UnpackLog(event, "Staked", log); err != nil {
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
func (_StakingCompoundContract *StakingCompoundContractFilterer) ParseStaked(log types.Log) (*StakingCompoundContractStaked, error) {
	event := new(StakingCompoundContractStaked)
	if err := _StakingCompoundContract.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingCompoundContractTipsSetIterator is returned from FilterTipsSet and is used to iterate over the raw logs and unpacked data for TipsSet events raised by the StakingCompoundContract contract.
type StakingCompoundContractTipsSetIterator struct {
	Event *StakingCompoundContractTipsSet // Event containing the contract specifics and raw log

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
func (it *StakingCompoundContractTipsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingCompoundContractTipsSet)
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
		it.Event = new(StakingCompoundContractTipsSet)
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
func (it *StakingCompoundContractTipsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingCompoundContractTipsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingCompoundContractTipsSet represents a TipsSet event raised by the StakingCompoundContract contract.
type StakingCompoundContractTipsSet struct {
	OldTips *big.Int
	NewTips *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTipsSet is a free log retrieval operation binding the contract event 0xce10e156265be187b2e1b5b553328a9a368e94c2a1825ea93046809494aad258.
//
// Solidity: event TipsSet(uint256 oldTips, uint256 newTips)
func (_StakingCompoundContract *StakingCompoundContractFilterer) FilterTipsSet(opts *bind.FilterOpts) (*StakingCompoundContractTipsSetIterator, error) {

	logs, sub, err := _StakingCompoundContract.contract.FilterLogs(opts, "TipsSet")
	if err != nil {
		return nil, err
	}
	return &StakingCompoundContractTipsSetIterator{contract: _StakingCompoundContract.contract, event: "TipsSet", logs: logs, sub: sub}, nil
}

// WatchTipsSet is a free log subscription operation binding the contract event 0xce10e156265be187b2e1b5b553328a9a368e94c2a1825ea93046809494aad258.
//
// Solidity: event TipsSet(uint256 oldTips, uint256 newTips)
func (_StakingCompoundContract *StakingCompoundContractFilterer) WatchTipsSet(opts *bind.WatchOpts, sink chan<- *StakingCompoundContractTipsSet) (event.Subscription, error) {

	logs, sub, err := _StakingCompoundContract.contract.WatchLogs(opts, "TipsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingCompoundContractTipsSet)
				if err := _StakingCompoundContract.contract.UnpackLog(event, "TipsSet", log); err != nil {
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

// ParseTipsSet is a log parse operation binding the contract event 0xce10e156265be187b2e1b5b553328a9a368e94c2a1825ea93046809494aad258.
//
// Solidity: event TipsSet(uint256 oldTips, uint256 newTips)
func (_StakingCompoundContract *StakingCompoundContractFilterer) ParseTipsSet(log types.Log) (*StakingCompoundContractTipsSet, error) {
	event := new(StakingCompoundContractTipsSet)
	if err := _StakingCompoundContract.contract.UnpackLog(event, "TipsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingCompoundContractUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the StakingCompoundContract contract.
type StakingCompoundContractUnpausedIterator struct {
	Event *StakingCompoundContractUnpaused // Event containing the contract specifics and raw log

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
func (it *StakingCompoundContractUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingCompoundContractUnpaused)
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
		it.Event = new(StakingCompoundContractUnpaused)
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
func (it *StakingCompoundContractUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingCompoundContractUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingCompoundContractUnpaused represents a Unpaused event raised by the StakingCompoundContract contract.
type StakingCompoundContractUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StakingCompoundContract *StakingCompoundContractFilterer) FilterUnpaused(opts *bind.FilterOpts) (*StakingCompoundContractUnpausedIterator, error) {

	logs, sub, err := _StakingCompoundContract.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &StakingCompoundContractUnpausedIterator{contract: _StakingCompoundContract.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StakingCompoundContract *StakingCompoundContractFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StakingCompoundContractUnpaused) (event.Subscription, error) {

	logs, sub, err := _StakingCompoundContract.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingCompoundContractUnpaused)
				if err := _StakingCompoundContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_StakingCompoundContract *StakingCompoundContractFilterer) ParseUnpaused(log types.Log) (*StakingCompoundContractUnpaused, error) {
	event := new(StakingCompoundContractUnpaused)
	if err := _StakingCompoundContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
