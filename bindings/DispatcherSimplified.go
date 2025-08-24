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
	_ = abi.ConvertType
)

// DispatcherSimplifiedMetaData contains all meta data concerning the DispatcherSimplified contract.
var DispatcherSimplifiedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"commandType\",\"type\":\"uint256\"}],\"name\":\"InvalidCommandType\",\"type\":\"error\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50603e80601a5f395ff3fe60806040525f5ffdfea2646970667358221220dfffe29f99a4d198371a6949e945497ae1467df885e118243cccf875190503b364736f6c634300081e0033",
}

// DispatcherSimplifiedABI is the input ABI used to generate the binding from.
// Deprecated: Use DispatcherSimplifiedMetaData.ABI instead.
var DispatcherSimplifiedABI = DispatcherSimplifiedMetaData.ABI

// DispatcherSimplifiedBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DispatcherSimplifiedMetaData.Bin instead.
var DispatcherSimplifiedBin = DispatcherSimplifiedMetaData.Bin

// DeployDispatcherSimplified deploys a new Ethereum contract, binding an instance of DispatcherSimplified to it.
func DeployDispatcherSimplified(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DispatcherSimplified, error) {
	parsed, err := DispatcherSimplifiedMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DispatcherSimplifiedBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DispatcherSimplified{DispatcherSimplifiedCaller: DispatcherSimplifiedCaller{contract: contract}, DispatcherSimplifiedTransactor: DispatcherSimplifiedTransactor{contract: contract}, DispatcherSimplifiedFilterer: DispatcherSimplifiedFilterer{contract: contract}}, nil
}

// DispatcherSimplified is an auto generated Go binding around an Ethereum contract.
type DispatcherSimplified struct {
	DispatcherSimplifiedCaller     // Read-only binding to the contract
	DispatcherSimplifiedTransactor // Write-only binding to the contract
	DispatcherSimplifiedFilterer   // Log filterer for contract events
}

// DispatcherSimplifiedCaller is an auto generated read-only Go binding around an Ethereum contract.
type DispatcherSimplifiedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DispatcherSimplifiedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DispatcherSimplifiedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DispatcherSimplifiedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DispatcherSimplifiedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DispatcherSimplifiedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DispatcherSimplifiedSession struct {
	Contract     *DispatcherSimplified // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DispatcherSimplifiedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DispatcherSimplifiedCallerSession struct {
	Contract *DispatcherSimplifiedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// DispatcherSimplifiedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DispatcherSimplifiedTransactorSession struct {
	Contract     *DispatcherSimplifiedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// DispatcherSimplifiedRaw is an auto generated low-level Go binding around an Ethereum contract.
type DispatcherSimplifiedRaw struct {
	Contract *DispatcherSimplified // Generic contract binding to access the raw methods on
}

// DispatcherSimplifiedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DispatcherSimplifiedCallerRaw struct {
	Contract *DispatcherSimplifiedCaller // Generic read-only contract binding to access the raw methods on
}

// DispatcherSimplifiedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DispatcherSimplifiedTransactorRaw struct {
	Contract *DispatcherSimplifiedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDispatcherSimplified creates a new instance of DispatcherSimplified, bound to a specific deployed contract.
func NewDispatcherSimplified(address common.Address, backend bind.ContractBackend) (*DispatcherSimplified, error) {
	contract, err := bindDispatcherSimplified(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DispatcherSimplified{DispatcherSimplifiedCaller: DispatcherSimplifiedCaller{contract: contract}, DispatcherSimplifiedTransactor: DispatcherSimplifiedTransactor{contract: contract}, DispatcherSimplifiedFilterer: DispatcherSimplifiedFilterer{contract: contract}}, nil
}

// NewDispatcherSimplifiedCaller creates a new read-only instance of DispatcherSimplified, bound to a specific deployed contract.
func NewDispatcherSimplifiedCaller(address common.Address, caller bind.ContractCaller) (*DispatcherSimplifiedCaller, error) {
	contract, err := bindDispatcherSimplified(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DispatcherSimplifiedCaller{contract: contract}, nil
}

// NewDispatcherSimplifiedTransactor creates a new write-only instance of DispatcherSimplified, bound to a specific deployed contract.
func NewDispatcherSimplifiedTransactor(address common.Address, transactor bind.ContractTransactor) (*DispatcherSimplifiedTransactor, error) {
	contract, err := bindDispatcherSimplified(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DispatcherSimplifiedTransactor{contract: contract}, nil
}

// NewDispatcherSimplifiedFilterer creates a new log filterer instance of DispatcherSimplified, bound to a specific deployed contract.
func NewDispatcherSimplifiedFilterer(address common.Address, filterer bind.ContractFilterer) (*DispatcherSimplifiedFilterer, error) {
	contract, err := bindDispatcherSimplified(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DispatcherSimplifiedFilterer{contract: contract}, nil
}

// bindDispatcherSimplified binds a generic wrapper to an already deployed contract.
func bindDispatcherSimplified(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DispatcherSimplifiedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DispatcherSimplified *DispatcherSimplifiedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DispatcherSimplified.Contract.DispatcherSimplifiedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DispatcherSimplified *DispatcherSimplifiedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DispatcherSimplified.Contract.DispatcherSimplifiedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DispatcherSimplified *DispatcherSimplifiedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DispatcherSimplified.Contract.DispatcherSimplifiedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DispatcherSimplified *DispatcherSimplifiedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DispatcherSimplified.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DispatcherSimplified *DispatcherSimplifiedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DispatcherSimplified.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DispatcherSimplified *DispatcherSimplifiedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DispatcherSimplified.Contract.contract.Transact(opts, method, params...)
}
