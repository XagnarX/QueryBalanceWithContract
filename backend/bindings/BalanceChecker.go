// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// BalanceCheckerMetaData contains all meta data concerning the BalanceChecker contract.
var BalanceCheckerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getAddressBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ethBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenBalances\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"getERC20Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"}],\"name\":\"getERC20Balances\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"getETHBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"}],\"name\":\"getETHBalances\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getMultipleAddressBalances\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ethBalances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"tokenBalances\",\"type\":\"uint256[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"}],\"name\":\"getMultipleERC20Balances\",\"outputs\":[{\"internalType\":\"uint256[][]\",\"name\":\"\",\"type\":\"uint256[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BalanceCheckerABI is the input ABI used to generate the binding from.
// Deprecated: Use BalanceCheckerMetaData.ABI instead.
var BalanceCheckerABI = BalanceCheckerMetaData.ABI

// BalanceChecker is an auto generated Go binding around an Ethereum contract.
type BalanceChecker struct {
	BalanceCheckerCaller     // Read-only binding to the contract
	BalanceCheckerTransactor // Write-only binding to the contract
	BalanceCheckerFilterer   // Log filterer for contract events
}

// BalanceCheckerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalanceCheckerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceCheckerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalanceCheckerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceCheckerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalanceCheckerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceCheckerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalanceCheckerSession struct {
	Contract     *BalanceChecker   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth to use throughout this session
}

// BalanceCheckerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalanceCheckerCallerSession struct {
	Contract *BalanceCheckerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BalanceCheckerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalanceCheckerTransactorSession struct {
	Contract     *BalanceCheckerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth to use throughout this session
}

// BalanceCheckerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalanceCheckerRaw struct {
	Contract *BalanceChecker // Generic contract binding to access the raw methods on
}

// BalanceCheckerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalanceCheckerCallerRaw struct {
	Contract *BalanceCheckerCaller // Generic read-only contract binding to access the raw methods on
}

// BalanceCheckerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalanceCheckerTransactorRaw struct {
	Contract *BalanceCheckerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalanceChecker creates a new instance of BalanceChecker, bound to a specific deployed contract.
func NewBalanceChecker(address common.Address, backend bind.ContractBackend) (*BalanceChecker, error) {
	contract, err := bindBalanceChecker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BalanceChecker{BalanceCheckerCaller: BalanceCheckerCaller{contract: contract}, BalanceCheckerTransactor: BalanceCheckerTransactor{contract: contract}, BalanceCheckerFilterer: BalanceCheckerFilterer{contract: contract}}, nil
}

// NewBalanceCheckerCaller creates a new read-only instance of BalanceChecker, bound to a specific deployed contract.
func NewBalanceCheckerCaller(address common.Address, caller bind.ContractCaller) (*BalanceCheckerCaller, error) {
	contract, err := bindBalanceChecker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalanceCheckerCaller{contract: contract}, nil
}

// NewBalanceCheckerTransactor creates a new write-only instance of BalanceChecker, bound to a specific deployed contract.
func NewBalanceCheckerTransactor(address common.Address, transactor bind.ContractTransactor) (*BalanceCheckerTransactor, error) {
	contract, err := bindBalanceChecker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalanceCheckerTransactor{contract: contract}, nil
}

// NewBalanceCheckerFilterer creates a new log filterer instance of BalanceChecker, bound to a specific deployed contract.
func NewBalanceCheckerFilterer(address common.Address, filterer bind.ContractFilterer) (*BalanceCheckerFilterer, error) {
	contract, err := bindBalanceChecker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalanceCheckerFilterer{contract: contract}, nil
}

// bindBalanceChecker binds a generic wrapper to an already deployed contract.
func bindBalanceChecker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalanceCheckerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalanceChecker *BalanceCheckerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalanceChecker.Contract.BalanceCheckerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// the fallback function.
func (_BalanceChecker *BalanceCheckerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalanceChecker.Contract.BalanceCheckerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalanceChecker *BalanceCheckerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalanceChecker.Contract.BalanceCheckerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalanceChecker *BalanceCheckerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalanceChecker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// the fallback function.
func (_BalanceChecker *BalanceCheckerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalanceChecker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalanceChecker *BalanceCheckerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalanceChecker.Contract.contract.Transact(opts, method, params...)
}

// GetAddressBalances is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getAddressBalances(address target, address[] tokens) view returns(uint256 ethBalance, uint256[] tokenBalances)
func (_BalanceChecker *BalanceCheckerCaller) GetAddressBalances(opts *bind.CallOpts, target common.Address, tokens []common.Address) (struct {
	EthBalance    *big.Int
	TokenBalances []*big.Int
}, error) {
	var out []interface{}
	err := _BalanceChecker.contract.Call(opts, &out, "getAddressBalances", target, tokens)

	outstruct := new(struct {
		EthBalance    *big.Int
		TokenBalances []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EthBalance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TokenBalances = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err
}

// GetAddressBalances is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getAddressBalances(address target, address[] tokens) view returns(uint256 ethBalance, uint256[] tokenBalances)
func (_BalanceChecker *BalanceCheckerSession) GetAddressBalances(target common.Address, tokens []common.Address) (struct {
	EthBalance    *big.Int
	TokenBalances []*big.Int
}, error) {
	return _BalanceChecker.Contract.GetAddressBalances(&_BalanceChecker.CallOpts, target, tokens)
}

// GetAddressBalances is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getAddressBalances(address target, address[] tokens) view returns(uint256 ethBalance, uint256[] tokenBalances)
func (_BalanceChecker *BalanceCheckerCallerSession) GetAddressBalances(target common.Address, tokens []common.Address) (struct {
	EthBalance    *big.Int
	TokenBalances []*big.Int
}, error) {
	return _BalanceChecker.Contract.GetAddressBalances(&_BalanceChecker.CallOpts, target, tokens)
}

// GetERC20Balance is a free data retrieval call binding the contract method 0x59bf5bab.
//
// Solidity: function getERC20Balance(address token, address target) view returns(uint256)
func (_BalanceChecker *BalanceCheckerCaller) GetERC20Balance(opts *bind.CallOpts, token common.Address, target common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BalanceChecker.contract.Call(opts, &out, "getERC20Balance", token, target)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// GetERC20Balance is a free data retrieval call binding the contract method 0x59bf5bab.
//
// Solidity: function getERC20Balance(address token, address target) view returns(uint256)
func (_BalanceChecker *BalanceCheckerSession) GetERC20Balance(token common.Address, target common.Address) (*big.Int, error) {
	return _BalanceChecker.Contract.GetERC20Balance(&_BalanceChecker.CallOpts, token, target)
}

// GetERC20Balance is a free data retrieval call binding the contract method 0x59bf5bab.
//
// Solidity: function getERC20Balance(address token, address target) view returns(uint256)
func (_BalanceChecker *BalanceCheckerCallerSession) GetERC20Balance(token common.Address, target common.Address) (*big.Int, error) {
	return _BalanceChecker.Contract.GetERC20Balance(&_BalanceChecker.CallOpts, token, target)
}

// GetERC20Balances is a free data retrieval call binding the contract method 0xa0901e51.
//
// Solidity: function getERC20Balances(address token, address[] targets) view returns(uint256[])
func (_BalanceChecker *BalanceCheckerCaller) GetERC20Balances(opts *bind.CallOpts, token common.Address, targets []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _BalanceChecker.contract.Call(opts, &out, "getERC20Balances", token, targets)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err
}

// GetERC20Balances is a free data retrieval call binding the contract method 0xa0901e51.
//
// Solidity: function getERC20Balances(address token, address[] targets) view returns(uint256[])
func (_BalanceChecker *BalanceCheckerSession) GetERC20Balances(token common.Address, targets []common.Address) ([]*big.Int, error) {
	return _BalanceChecker.Contract.GetERC20Balances(&_BalanceChecker.CallOpts, token, targets)
}

// GetERC20Balances is a free data retrieval call binding the contract method 0xa0901e51.
//
// Solidity: function getERC20Balances(address token, address[] targets) view returns(uint256[])
func (_BalanceChecker *BalanceCheckerCallerSession) GetERC20Balances(token common.Address, targets []common.Address) ([]*big.Int, error) {
	return _BalanceChecker.Contract.GetERC20Balances(&_BalanceChecker.CallOpts, token, targets)
}

// GetETHBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getETHBalance(address target) view returns(uint256)
func (_BalanceChecker *BalanceCheckerCaller) GetETHBalance(opts *bind.CallOpts, target common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BalanceChecker.contract.Call(opts, &out, "getETHBalance", target)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// GetETHBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getETHBalance(address target) view returns(uint256)
func (_BalanceChecker *BalanceCheckerSession) GetETHBalance(target common.Address) (*big.Int, error) {
	return _BalanceChecker.Contract.GetETHBalance(&_BalanceChecker.CallOpts, target)
}

// GetETHBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getETHBalance(address target) view returns(uint256)
func (_BalanceChecker *BalanceCheckerCallerSession) GetETHBalance(target common.Address) (*big.Int, error) {
	return _BalanceChecker.Contract.GetETHBalance(&_BalanceChecker.CallOpts, target)
}

// GetETHBalances is a free data retrieval call binding the contract method 0x11b50c5d.
//
// Solidity: function getETHBalances(address[] targets) view returns(uint256[])
func (_BalanceChecker *BalanceCheckerCaller) GetETHBalances(opts *bind.CallOpts, targets []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _BalanceChecker.contract.Call(opts, &out, "getETHBalances", targets)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err
}

// GetETHBalances is a free data retrieval call binding the contract method 0x11b50c5d.
//
// Solidity: function getETHBalances(address[] targets) view returns(uint256[])
func (_BalanceChecker *BalanceCheckerSession) GetETHBalances(targets []common.Address) ([]*big.Int, error) {
	return _BalanceChecker.Contract.GetETHBalances(&_BalanceChecker.CallOpts, targets)
}

// GetETHBalances is a free data retrieval call binding the contract method 0x11b50c5d.
//
// Solidity: function getETHBalances(address[] targets) view returns(uint256[])
func (_BalanceChecker *BalanceCheckerCallerSession) GetETHBalances(targets []common.Address) ([]*big.Int, error) {
	return _BalanceChecker.Contract.GetETHBalances(&_BalanceChecker.CallOpts, targets)
}

// GetMultipleAddressBalances is a free data retrieval call binding the contract method 0x7cdcd02d.
//
// Solidity: function getMultipleAddressBalances(address[] targets, address[] tokens) view returns(uint256[] ethBalances, uint256[][] tokenBalances)
func (_BalanceChecker *BalanceCheckerCaller) GetMultipleAddressBalances(opts *bind.CallOpts, targets []common.Address, tokens []common.Address) (struct {
	EthBalances   []*big.Int
	TokenBalances [][]*big.Int
}, error) {
	var out []interface{}
	err := _BalanceChecker.contract.Call(opts, &out, "getMultipleAddressBalances", targets, tokens)

	outstruct := new(struct {
		EthBalances   []*big.Int
		TokenBalances [][]*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EthBalances = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.TokenBalances = *abi.ConvertType(out[1], new([][]*big.Int)).(*[][]*big.Int)

	return *outstruct, err
}

// GetMultipleAddressBalances is a free data retrieval call binding the contract method 0x7cdcd02d.
//
// Solidity: function getMultipleAddressBalances(address[] targets, address[] tokens) view returns(uint256[] ethBalances, uint256[][] tokenBalances)
func (_BalanceChecker *BalanceCheckerSession) GetMultipleAddressBalances(targets []common.Address, tokens []common.Address) (struct {
	EthBalances   []*big.Int
	TokenBalances [][]*big.Int
}, error) {
	return _BalanceChecker.Contract.GetMultipleAddressBalances(&_BalanceChecker.CallOpts, targets, tokens)
}

// GetMultipleAddressBalances is a free data retrieval call binding the contract method 0x7cdcd02d.
//
// Solidity: function getMultipleAddressBalances(address[] targets, address[] tokens) view returns(uint256[] ethBalances, uint256[][] tokenBalances)
func (_BalanceChecker *BalanceCheckerCallerSession) GetMultipleAddressBalances(targets []common.Address, tokens []common.Address) (struct {
	EthBalances   []*big.Int
	TokenBalances [][]*big.Int
}, error) {
	return _BalanceChecker.Contract.GetMultipleAddressBalances(&_BalanceChecker.CallOpts, targets, tokens)
}

// GetMultipleERC20Balances is a free data retrieval call binding the contract method 0xac92528d.
//
// Solidity: function getMultipleERC20Balances(address[] tokens, address[] targets) view returns(uint256[][])
func (_BalanceChecker *BalanceCheckerCaller) GetMultipleERC20Balances(opts *bind.CallOpts, tokens []common.Address, targets []common.Address) ([][]*big.Int, error) {
	var out []interface{}
	err := _BalanceChecker.contract.Call(opts, &out, "getMultipleERC20Balances", tokens, targets)

	if err != nil {
		return *new([][]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][]*big.Int)).(*[][]*big.Int)

	return out0, err
}

// GetMultipleERC20Balances is a free data retrieval call binding the contract method 0xac92528d.
//
// Solidity: function getMultipleERC20Balances(address[] tokens, address[] targets) view returns(uint256[][])
func (_BalanceChecker *BalanceCheckerSession) GetMultipleERC20Balances(tokens []common.Address, targets []common.Address) ([][]*big.Int, error) {
	return _BalanceChecker.Contract.GetMultipleERC20Balances(&_BalanceChecker.CallOpts, tokens, targets)
}

// GetMultipleERC20Balances is a free data retrieval call binding the contract method 0xac92528d.
//
// Solidity: function getMultipleERC20Balances(address[] tokens, address[] targets) view returns(uint256[][])
func (_BalanceChecker *BalanceCheckerCallerSession) GetMultipleERC20Balances(tokens []common.Address, targets []common.Address) ([][]*big.Int, error) {
	return _BalanceChecker.Contract.GetMultipleERC20Balances(&_BalanceChecker.CallOpts, tokens, targets)
}
