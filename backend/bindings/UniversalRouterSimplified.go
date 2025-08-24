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

// UniversalRouterSimplifiedMetaData contains all meta data concerning the UniversalRouterSimplified contract.
var UniversalRouterSimplifiedMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"commandIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"ExecutionFailed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"commands\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"inputs\",\"type\":\"bytes[]\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"commands\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"inputs\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506107a08061001c5f395ff3fe60806040526004361061002c575f3560e01c806324856bc3146100375780633593564c1461005357610033565b3661003357005b5f5ffd5b610051600480360381019061004c919061036b565b61006f565b005b61006d6004803603810190610068919061041c565b610255565b005b5f8484905090508083839050146100bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100b290610507565b60405180910390fd5b5f5f90505b8181101561015e575f8686838181106100dc576100db610525565b5b9050013560f81c60f81b9050365f8686858181106100fd576100fc610525565b5b905060200281019061010f919061055e565b915091507f2c4029e985980dfd49697129a5665bb97f9ecbe343afce80cc9d0fdac2e1629384838360405161014693929190610629565b60405180910390a183806001019450505050506100c0565b505f4790505f8111801561019e57503073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b1561024d575f3373ffffffffffffffffffffffffffffffffffffffff16826040516101c890610686565b5f6040518083038185875af1925050503d805f8114610202576040519150601f19603f3d011682016040523d82523d5f602084013e610207565b606091505b505090508061024b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610242906106e4565b60405180910390fd5b505b505050505050565b8080421115610299576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102909061074c565b60405180910390fd5b6102a58686868661006f565b505050505050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f8401126102d6576102d56102b5565b5b8235905067ffffffffffffffff8111156102f3576102f26102b9565b5b60208301915083600182028301111561030f5761030e6102bd565b5b9250929050565b5f5f83601f84011261032b5761032a6102b5565b5b8235905067ffffffffffffffff811115610348576103476102b9565b5b602083019150836020820283011115610364576103636102bd565b5b9250929050565b5f5f5f5f60408587031215610383576103826102ad565b5b5f85013567ffffffffffffffff8111156103a05761039f6102b1565b5b6103ac878288016102c1565b9450945050602085013567ffffffffffffffff8111156103cf576103ce6102b1565b5b6103db87828801610316565b925092505092959194509250565b5f819050919050565b6103fb816103e9565b8114610405575f5ffd5b50565b5f81359050610416816103f2565b92915050565b5f5f5f5f5f60608688031215610435576104346102ad565b5b5f86013567ffffffffffffffff811115610452576104516102b1565b5b61045e888289016102c1565b9550955050602086013567ffffffffffffffff811115610481576104806102b1565b5b61048d88828901610316565b935093505060406104a088828901610408565b9150509295509295909350565b5f82825260208201905092915050565b7f4c656e6774684d69736d617463680000000000000000000000000000000000005f82015250565b5f6104f1600e836104ad565b91506104fc826104bd565b602082019050919050565b5f6020820190508181035f83015261051e816104e5565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5f8335600160200384360303811261057a57610579610552565b5b80840192508235915067ffffffffffffffff82111561059c5761059b610556565b5b6020830192506001820236038313156105b8576105b761055a565b5b509250929050565b6105c9816103e9565b82525050565b5f82825260208201905092915050565b828183375f83830152505050565b5f601f19601f8301169050919050565b5f61060883856105cf565b93506106158385846105df565b61061e836105ed565b840190509392505050565b5f60408201905061063c5f8301866105c0565b818103602083015261064f8184866105fd565b9050949350505050565b5f81905092915050565b50565b5f6106715f83610659565b915061067c82610663565b5f82019050919050565b5f61069082610666565b9150819050919050565b7f455448207472616e73666572206661696c6564000000000000000000000000005f82015250565b5f6106ce6013836104ad565b91506106d98261069a565b602082019050919050565b5f6020820190508181035f8301526106fb816106c2565b9050919050565b7f5472616e73616374696f6e446561646c696e65506173736564000000000000005f82015250565b5f6107366019836104ad565b915061074182610702565b602082019050919050565b5f6020820190508181035f8301526107638161072a565b905091905056fea26469706673582212200df2d4f1cc14bf27dcce22fb3ca52328a6b42ced5ad86e7748a81d3929b5285864736f6c634300081e0033",
}

// UniversalRouterSimplifiedABI is the input ABI used to generate the binding from.
// Deprecated: Use UniversalRouterSimplifiedMetaData.ABI instead.
var UniversalRouterSimplifiedABI = UniversalRouterSimplifiedMetaData.ABI

// UniversalRouterSimplifiedBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UniversalRouterSimplifiedMetaData.Bin instead.
var UniversalRouterSimplifiedBin = UniversalRouterSimplifiedMetaData.Bin

// DeployUniversalRouterSimplified deploys a new Ethereum contract, binding an instance of UniversalRouterSimplified to it.
func DeployUniversalRouterSimplified(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UniversalRouterSimplified, error) {
	parsed, err := UniversalRouterSimplifiedMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UniversalRouterSimplifiedBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UniversalRouterSimplified{UniversalRouterSimplifiedCaller: UniversalRouterSimplifiedCaller{contract: contract}, UniversalRouterSimplifiedTransactor: UniversalRouterSimplifiedTransactor{contract: contract}, UniversalRouterSimplifiedFilterer: UniversalRouterSimplifiedFilterer{contract: contract}}, nil
}

// UniversalRouterSimplified is an auto generated Go binding around an Ethereum contract.
type UniversalRouterSimplified struct {
	UniversalRouterSimplifiedCaller     // Read-only binding to the contract
	UniversalRouterSimplifiedTransactor // Write-only binding to the contract
	UniversalRouterSimplifiedFilterer   // Log filterer for contract events
}

// UniversalRouterSimplifiedCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniversalRouterSimplifiedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalRouterSimplifiedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniversalRouterSimplifiedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalRouterSimplifiedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniversalRouterSimplifiedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalRouterSimplifiedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniversalRouterSimplifiedSession struct {
	Contract     *UniversalRouterSimplified // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// UniversalRouterSimplifiedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniversalRouterSimplifiedCallerSession struct {
	Contract *UniversalRouterSimplifiedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// UniversalRouterSimplifiedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniversalRouterSimplifiedTransactorSession struct {
	Contract     *UniversalRouterSimplifiedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// UniversalRouterSimplifiedRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniversalRouterSimplifiedRaw struct {
	Contract *UniversalRouterSimplified // Generic contract binding to access the raw methods on
}

// UniversalRouterSimplifiedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniversalRouterSimplifiedCallerRaw struct {
	Contract *UniversalRouterSimplifiedCaller // Generic read-only contract binding to access the raw methods on
}

// UniversalRouterSimplifiedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniversalRouterSimplifiedTransactorRaw struct {
	Contract *UniversalRouterSimplifiedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniversalRouterSimplified creates a new instance of UniversalRouterSimplified, bound to a specific deployed contract.
func NewUniversalRouterSimplified(address common.Address, backend bind.ContractBackend) (*UniversalRouterSimplified, error) {
	contract, err := bindUniversalRouterSimplified(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniversalRouterSimplified{UniversalRouterSimplifiedCaller: UniversalRouterSimplifiedCaller{contract: contract}, UniversalRouterSimplifiedTransactor: UniversalRouterSimplifiedTransactor{contract: contract}, UniversalRouterSimplifiedFilterer: UniversalRouterSimplifiedFilterer{contract: contract}}, nil
}

// NewUniversalRouterSimplifiedCaller creates a new read-only instance of UniversalRouterSimplified, bound to a specific deployed contract.
func NewUniversalRouterSimplifiedCaller(address common.Address, caller bind.ContractCaller) (*UniversalRouterSimplifiedCaller, error) {
	contract, err := bindUniversalRouterSimplified(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniversalRouterSimplifiedCaller{contract: contract}, nil
}

// NewUniversalRouterSimplifiedTransactor creates a new write-only instance of UniversalRouterSimplified, bound to a specific deployed contract.
func NewUniversalRouterSimplifiedTransactor(address common.Address, transactor bind.ContractTransactor) (*UniversalRouterSimplifiedTransactor, error) {
	contract, err := bindUniversalRouterSimplified(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniversalRouterSimplifiedTransactor{contract: contract}, nil
}

// NewUniversalRouterSimplifiedFilterer creates a new log filterer instance of UniversalRouterSimplified, bound to a specific deployed contract.
func NewUniversalRouterSimplifiedFilterer(address common.Address, filterer bind.ContractFilterer) (*UniversalRouterSimplifiedFilterer, error) {
	contract, err := bindUniversalRouterSimplified(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniversalRouterSimplifiedFilterer{contract: contract}, nil
}

// bindUniversalRouterSimplified binds a generic wrapper to an already deployed contract.
func bindUniversalRouterSimplified(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniversalRouterSimplifiedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniversalRouterSimplified *UniversalRouterSimplifiedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniversalRouterSimplified.Contract.UniversalRouterSimplifiedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniversalRouterSimplified *UniversalRouterSimplifiedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniversalRouterSimplified.Contract.UniversalRouterSimplifiedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniversalRouterSimplified *UniversalRouterSimplifiedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniversalRouterSimplified.Contract.UniversalRouterSimplifiedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniversalRouterSimplified *UniversalRouterSimplifiedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniversalRouterSimplified.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniversalRouterSimplified *UniversalRouterSimplifiedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniversalRouterSimplified.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniversalRouterSimplified *UniversalRouterSimplifiedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniversalRouterSimplified.Contract.contract.Transact(opts, method, params...)
}

// Execute is a paid mutator transaction binding the contract method 0x24856bc3.
//
// Solidity: function execute(bytes commands, bytes[] inputs) payable returns()
func (_UniversalRouterSimplified *UniversalRouterSimplifiedTransactor) Execute(opts *bind.TransactOpts, commands []byte, inputs [][]byte) (*types.Transaction, error) {
	return _UniversalRouterSimplified.contract.Transact(opts, "execute", commands, inputs)
}

// Execute is a paid mutator transaction binding the contract method 0x24856bc3.
//
// Solidity: function execute(bytes commands, bytes[] inputs) payable returns()
func (_UniversalRouterSimplified *UniversalRouterSimplifiedSession) Execute(commands []byte, inputs [][]byte) (*types.Transaction, error) {
	return _UniversalRouterSimplified.Contract.Execute(&_UniversalRouterSimplified.TransactOpts, commands, inputs)
}

// Execute is a paid mutator transaction binding the contract method 0x24856bc3.
//
// Solidity: function execute(bytes commands, bytes[] inputs) payable returns()
func (_UniversalRouterSimplified *UniversalRouterSimplifiedTransactorSession) Execute(commands []byte, inputs [][]byte) (*types.Transaction, error) {
	return _UniversalRouterSimplified.Contract.Execute(&_UniversalRouterSimplified.TransactOpts, commands, inputs)
}

// Execute0 is a paid mutator transaction binding the contract method 0x3593564c.
//
// Solidity: function execute(bytes commands, bytes[] inputs, uint256 deadline) payable returns()
func (_UniversalRouterSimplified *UniversalRouterSimplifiedTransactor) Execute0(opts *bind.TransactOpts, commands []byte, inputs [][]byte, deadline *big.Int) (*types.Transaction, error) {
	return _UniversalRouterSimplified.contract.Transact(opts, "execute0", commands, inputs, deadline)
}

// Execute0 is a paid mutator transaction binding the contract method 0x3593564c.
//
// Solidity: function execute(bytes commands, bytes[] inputs, uint256 deadline) payable returns()
func (_UniversalRouterSimplified *UniversalRouterSimplifiedSession) Execute0(commands []byte, inputs [][]byte, deadline *big.Int) (*types.Transaction, error) {
	return _UniversalRouterSimplified.Contract.Execute0(&_UniversalRouterSimplified.TransactOpts, commands, inputs, deadline)
}

// Execute0 is a paid mutator transaction binding the contract method 0x3593564c.
//
// Solidity: function execute(bytes commands, bytes[] inputs, uint256 deadline) payable returns()
func (_UniversalRouterSimplified *UniversalRouterSimplifiedTransactorSession) Execute0(commands []byte, inputs [][]byte, deadline *big.Int) (*types.Transaction, error) {
	return _UniversalRouterSimplified.Contract.Execute0(&_UniversalRouterSimplified.TransactOpts, commands, inputs, deadline)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniversalRouterSimplified *UniversalRouterSimplifiedTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniversalRouterSimplified.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniversalRouterSimplified *UniversalRouterSimplifiedSession) Receive() (*types.Transaction, error) {
	return _UniversalRouterSimplified.Contract.Receive(&_UniversalRouterSimplified.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniversalRouterSimplified *UniversalRouterSimplifiedTransactorSession) Receive() (*types.Transaction, error) {
	return _UniversalRouterSimplified.Contract.Receive(&_UniversalRouterSimplified.TransactOpts)
}

// UniversalRouterSimplifiedExecutionFailedIterator is returned from FilterExecutionFailed and is used to iterate over the raw logs and unpacked data for ExecutionFailed events raised by the UniversalRouterSimplified contract.
type UniversalRouterSimplifiedExecutionFailedIterator struct {
	Event *UniversalRouterSimplifiedExecutionFailed // Event containing the contract specifics and raw log

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
func (it *UniversalRouterSimplifiedExecutionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalRouterSimplifiedExecutionFailed)
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
		it.Event = new(UniversalRouterSimplifiedExecutionFailed)
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
func (it *UniversalRouterSimplifiedExecutionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalRouterSimplifiedExecutionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalRouterSimplifiedExecutionFailed represents a ExecutionFailed event raised by the UniversalRouterSimplified contract.
type UniversalRouterSimplifiedExecutionFailed struct {
	CommandIndex *big.Int
	Message      []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterExecutionFailed is a free log retrieval operation binding the contract event 0x2c4029e985980dfd49697129a5665bb97f9ecbe343afce80cc9d0fdac2e16293.
//
// Solidity: event ExecutionFailed(uint256 commandIndex, bytes message)
func (_UniversalRouterSimplified *UniversalRouterSimplifiedFilterer) FilterExecutionFailed(opts *bind.FilterOpts) (*UniversalRouterSimplifiedExecutionFailedIterator, error) {

	logs, sub, err := _UniversalRouterSimplified.contract.FilterLogs(opts, "ExecutionFailed")
	if err != nil {
		return nil, err
	}
	return &UniversalRouterSimplifiedExecutionFailedIterator{contract: _UniversalRouterSimplified.contract, event: "ExecutionFailed", logs: logs, sub: sub}, nil
}

// WatchExecutionFailed is a free log subscription operation binding the contract event 0x2c4029e985980dfd49697129a5665bb97f9ecbe343afce80cc9d0fdac2e16293.
//
// Solidity: event ExecutionFailed(uint256 commandIndex, bytes message)
func (_UniversalRouterSimplified *UniversalRouterSimplifiedFilterer) WatchExecutionFailed(opts *bind.WatchOpts, sink chan<- *UniversalRouterSimplifiedExecutionFailed) (event.Subscription, error) {

	logs, sub, err := _UniversalRouterSimplified.contract.WatchLogs(opts, "ExecutionFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalRouterSimplifiedExecutionFailed)
				if err := _UniversalRouterSimplified.contract.UnpackLog(event, "ExecutionFailed", log); err != nil {
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

// ParseExecutionFailed is a log parse operation binding the contract event 0x2c4029e985980dfd49697129a5665bb97f9ecbe343afce80cc9d0fdac2e16293.
//
// Solidity: event ExecutionFailed(uint256 commandIndex, bytes message)
func (_UniversalRouterSimplified *UniversalRouterSimplifiedFilterer) ParseExecutionFailed(log types.Log) (*UniversalRouterSimplifiedExecutionFailed, error) {
	event := new(UniversalRouterSimplifiedExecutionFailed)
	if err := _UniversalRouterSimplified.contract.UnpackLog(event, "ExecutionFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
