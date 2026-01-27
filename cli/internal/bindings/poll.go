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

// IPollVoteCommitment is an auto generated low-level Go binding around an user-defined struct.
type IPollVoteCommitment struct {
	Commitment [32]byte
	Timestamp  *big.Int
	Revealed   bool
}

// PollMetaData contains all meta data concerning the Poll contract.
var PollMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_question\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"optionsArray\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"_duration\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_voterMerkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_zkVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_oracle\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"_options\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"closePoll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"commitVote\",\"inputs\":[{\"name\":\"commitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"zkProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"merklePath\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"commitments\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"commitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revealed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createdAt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"creator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"duration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"endTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCommitment\",\"inputs\":[{\"name\":\"voter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPoll.VoteCommitment\",\"components\":[{\"name\":\"commitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revealed\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getResults\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"options\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"oracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOracle\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"question\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"revealVote\",\"inputs\":[{\"name\":\"choice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"state\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIPoll.PollState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tally\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalCommitted\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalRevealed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"voterMerkleRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"zkVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIZKVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"PollClosed\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ResultsTallied\",\"inputs\":[{\"name\":\"results\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VoteCommitted\",\"inputs\":[{\"name\":\"voter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"commitment\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VoteRevealed\",\"inputs\":[{\"name\":\"voter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"choice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyRevealed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyVoted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidChoice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMerkleProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidReveal\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoCommitment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PollAlreadyClosed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PollAlreadyTallied\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PollNotActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PollNotClosed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PollNotTallied\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedOracle\",\"inputs\":[]}]",
}

// PollABI is the input ABI used to generate the binding from.
// Deprecated: Use PollMetaData.ABI instead.
var PollABI = PollMetaData.ABI

// Poll is an auto generated Go binding around an Ethereum contract.
type Poll struct {
	PollCaller     // Read-only binding to the contract
	PollTransactor // Write-only binding to the contract
	PollFilterer   // Log filterer for contract events
}

// PollCaller is an auto generated read-only Go binding around an Ethereum contract.
type PollCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PollTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PollTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PollFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PollFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PollSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PollSession struct {
	Contract     *Poll             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PollCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PollCallerSession struct {
	Contract *PollCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PollTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PollTransactorSession struct {
	Contract     *PollTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PollRaw is an auto generated low-level Go binding around an Ethereum contract.
type PollRaw struct {
	Contract *Poll // Generic contract binding to access the raw methods on
}

// PollCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PollCallerRaw struct {
	Contract *PollCaller // Generic read-only contract binding to access the raw methods on
}

// PollTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PollTransactorRaw struct {
	Contract *PollTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoll creates a new instance of Poll, bound to a specific deployed contract.
func NewPoll(address common.Address, backend bind.ContractBackend) (*Poll, error) {
	contract, err := bindPoll(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Poll{PollCaller: PollCaller{contract: contract}, PollTransactor: PollTransactor{contract: contract}, PollFilterer: PollFilterer{contract: contract}}, nil
}

// NewPollCaller creates a new read-only instance of Poll, bound to a specific deployed contract.
func NewPollCaller(address common.Address, caller bind.ContractCaller) (*PollCaller, error) {
	contract, err := bindPoll(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PollCaller{contract: contract}, nil
}

// NewPollTransactor creates a new write-only instance of Poll, bound to a specific deployed contract.
func NewPollTransactor(address common.Address, transactor bind.ContractTransactor) (*PollTransactor, error) {
	contract, err := bindPoll(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PollTransactor{contract: contract}, nil
}

// NewPollFilterer creates a new log filterer instance of Poll, bound to a specific deployed contract.
func NewPollFilterer(address common.Address, filterer bind.ContractFilterer) (*PollFilterer, error) {
	contract, err := bindPoll(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PollFilterer{contract: contract}, nil
}

// bindPoll binds a generic wrapper to an already deployed contract.
func bindPoll(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PollMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Poll *PollRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Poll.Contract.PollCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Poll *PollRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poll.Contract.PollTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Poll *PollRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Poll.Contract.PollTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Poll *PollCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Poll.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Poll *PollTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poll.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Poll *PollTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Poll.Contract.contract.Transact(opts, method, params...)
}

// Options is a free data retrieval call binding the contract method 0xa2545761.
//
// Solidity: function _options(uint256 ) view returns(string)
func (_Poll *PollCaller) Options(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Poll.contract.Call(opts, &out, "_options", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Options is a free data retrieval call binding the contract method 0xa2545761.
//
// Solidity: function _options(uint256 ) view returns(string)
func (_Poll *PollSession) Options(arg0 *big.Int) (string, error) {
	return _Poll.Contract.Options(&_Poll.CallOpts, arg0)
}

// Options is a free data retrieval call binding the contract method 0xa2545761.
//
// Solidity: function _options(uint256 ) view returns(string)
func (_Poll *PollCallerSession) Options(arg0 *big.Int) (string, error) {
	return _Poll.Contract.Options(&_Poll.CallOpts, arg0)
}

// Commitments is a free data retrieval call binding the contract method 0xe8fcf723.
//
// Solidity: function commitments(address ) view returns(bytes32 commitment, uint256 timestamp, bool revealed)
func (_Poll *PollCaller) Commitments(opts *bind.CallOpts, arg0 common.Address) (struct {
	Commitment [32]byte
	Timestamp  *big.Int
	Revealed   bool
}, error) {
	var out []interface{}
	err := _Poll.contract.Call(opts, &out, "commitments", arg0)

	outstruct := new(struct {
		Commitment [32]byte
		Timestamp  *big.Int
		Revealed   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Commitment = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Revealed = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Commitments is a free data retrieval call binding the contract method 0xe8fcf723.
//
// Solidity: function commitments(address ) view returns(bytes32 commitment, uint256 timestamp, bool revealed)
func (_Poll *PollSession) Commitments(arg0 common.Address) (struct {
	Commitment [32]byte
	Timestamp  *big.Int
	Revealed   bool
}, error) {
	return _Poll.Contract.Commitments(&_Poll.CallOpts, arg0)
}

// Commitments is a free data retrieval call binding the contract method 0xe8fcf723.
//
// Solidity: function commitments(address ) view returns(bytes32 commitment, uint256 timestamp, bool revealed)
func (_Poll *PollCallerSession) Commitments(arg0 common.Address) (struct {
	Commitment [32]byte
	Timestamp  *big.Int
	Revealed   bool
}, error) {
	return _Poll.Contract.Commitments(&_Poll.CallOpts, arg0)
}

// CreatedAt is a free data retrieval call binding the contract method 0xcf09e0d0.
//
// Solidity: function createdAt() view returns(uint256)
func (_Poll *PollCaller) CreatedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poll.contract.Call(opts, &out, "createdAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreatedAt is a free data retrieval call binding the contract method 0xcf09e0d0.
//
// Solidity: function createdAt() view returns(uint256)
func (_Poll *PollSession) CreatedAt() (*big.Int, error) {
	return _Poll.Contract.CreatedAt(&_Poll.CallOpts)
}

// CreatedAt is a free data retrieval call binding the contract method 0xcf09e0d0.
//
// Solidity: function createdAt() view returns(uint256)
func (_Poll *PollCallerSession) CreatedAt() (*big.Int, error) {
	return _Poll.Contract.CreatedAt(&_Poll.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_Poll *PollCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Poll.contract.Call(opts, &out, "creator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_Poll *PollSession) Creator() (common.Address, error) {
	return _Poll.Contract.Creator(&_Poll.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_Poll *PollCallerSession) Creator() (common.Address, error) {
	return _Poll.Contract.Creator(&_Poll.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_Poll *PollCaller) Duration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poll.contract.Call(opts, &out, "duration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_Poll *PollSession) Duration() (*big.Int, error) {
	return _Poll.Contract.Duration(&_Poll.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_Poll *PollCallerSession) Duration() (*big.Int, error) {
	return _Poll.Contract.Duration(&_Poll.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Poll *PollCaller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poll.contract.Call(opts, &out, "endTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Poll *PollSession) EndTime() (*big.Int, error) {
	return _Poll.Contract.EndTime(&_Poll.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Poll *PollCallerSession) EndTime() (*big.Int, error) {
	return _Poll.Contract.EndTime(&_Poll.CallOpts)
}

// GetCommitment is a free data retrieval call binding the contract method 0xfa1026dd.
//
// Solidity: function getCommitment(address voter) view returns((bytes32,uint256,bool))
func (_Poll *PollCaller) GetCommitment(opts *bind.CallOpts, voter common.Address) (IPollVoteCommitment, error) {
	var out []interface{}
	err := _Poll.contract.Call(opts, &out, "getCommitment", voter)

	if err != nil {
		return *new(IPollVoteCommitment), err
	}

	out0 := *abi.ConvertType(out[0], new(IPollVoteCommitment)).(*IPollVoteCommitment)

	return out0, err

}

// GetCommitment is a free data retrieval call binding the contract method 0xfa1026dd.
//
// Solidity: function getCommitment(address voter) view returns((bytes32,uint256,bool))
func (_Poll *PollSession) GetCommitment(voter common.Address) (IPollVoteCommitment, error) {
	return _Poll.Contract.GetCommitment(&_Poll.CallOpts, voter)
}

// GetCommitment is a free data retrieval call binding the contract method 0xfa1026dd.
//
// Solidity: function getCommitment(address voter) view returns((bytes32,uint256,bool))
func (_Poll *PollCallerSession) GetCommitment(voter common.Address) (IPollVoteCommitment, error) {
	return _Poll.Contract.GetCommitment(&_Poll.CallOpts, voter)
}

// GetResults is a free data retrieval call binding the contract method 0x4717f97c.
//
// Solidity: function getResults() view returns(uint256[])
func (_Poll *PollCaller) GetResults(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Poll.contract.Call(opts, &out, "getResults")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetResults is a free data retrieval call binding the contract method 0x4717f97c.
//
// Solidity: function getResults() view returns(uint256[])
func (_Poll *PollSession) GetResults() ([]*big.Int, error) {
	return _Poll.Contract.GetResults(&_Poll.CallOpts)
}

// GetResults is a free data retrieval call binding the contract method 0x4717f97c.
//
// Solidity: function getResults() view returns(uint256[])
func (_Poll *PollCallerSession) GetResults() ([]*big.Int, error) {
	return _Poll.Contract.GetResults(&_Poll.CallOpts)
}

// PollOptions is a free data retrieval call binding the contract method 0x1069143a.
//
// Solidity: function options() view returns(string[])
func (_Poll *PollCaller) PollOptions(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Poll.contract.Call(opts, &out, "options")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// PollOptions is a free data retrieval call binding the contract method 0x1069143a.
//
// Solidity: function options() view returns(string[])
func (_Poll *PollSession) PollOptions() ([]string, error) {
	return _Poll.Contract.PollOptions(&_Poll.CallOpts)
}

// PollOptions is a free data retrieval call binding the contract method 0x1069143a.
//
// Solidity: function options() view returns(string[])
func (_Poll *PollCallerSession) PollOptions() ([]string, error) {
	return _Poll.Contract.PollOptions(&_Poll.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Poll *PollCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Poll.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Poll *PollSession) Oracle() (common.Address, error) {
	return _Poll.Contract.Oracle(&_Poll.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Poll *PollCallerSession) Oracle() (common.Address, error) {
	return _Poll.Contract.Oracle(&_Poll.CallOpts)
}

// Question is a free data retrieval call binding the contract method 0x3fad9ae0.
//
// Solidity: function question() view returns(string)
func (_Poll *PollCaller) Question(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Poll.contract.Call(opts, &out, "question")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Question is a free data retrieval call binding the contract method 0x3fad9ae0.
//
// Solidity: function question() view returns(string)
func (_Poll *PollSession) Question() (string, error) {
	return _Poll.Contract.Question(&_Poll.CallOpts)
}

// Question is a free data retrieval call binding the contract method 0x3fad9ae0.
//
// Solidity: function question() view returns(string)
func (_Poll *PollCallerSession) Question() (string, error) {
	return _Poll.Contract.Question(&_Poll.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Poll *PollCaller) State(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Poll.contract.Call(opts, &out, "state")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Poll *PollSession) State() (uint8, error) {
	return _Poll.Contract.State(&_Poll.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Poll *PollCallerSession) State() (uint8, error) {
	return _Poll.Contract.State(&_Poll.CallOpts)
}

// TotalCommitted is a free data retrieval call binding the contract method 0x1d3231d4.
//
// Solidity: function totalCommitted() view returns(uint256)
func (_Poll *PollCaller) TotalCommitted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poll.contract.Call(opts, &out, "totalCommitted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCommitted is a free data retrieval call binding the contract method 0x1d3231d4.
//
// Solidity: function totalCommitted() view returns(uint256)
func (_Poll *PollSession) TotalCommitted() (*big.Int, error) {
	return _Poll.Contract.TotalCommitted(&_Poll.CallOpts)
}

// TotalCommitted is a free data retrieval call binding the contract method 0x1d3231d4.
//
// Solidity: function totalCommitted() view returns(uint256)
func (_Poll *PollCallerSession) TotalCommitted() (*big.Int, error) {
	return _Poll.Contract.TotalCommitted(&_Poll.CallOpts)
}

// TotalRevealed is a free data retrieval call binding the contract method 0xf3738569.
//
// Solidity: function totalRevealed() view returns(uint256)
func (_Poll *PollCaller) TotalRevealed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poll.contract.Call(opts, &out, "totalRevealed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRevealed is a free data retrieval call binding the contract method 0xf3738569.
//
// Solidity: function totalRevealed() view returns(uint256)
func (_Poll *PollSession) TotalRevealed() (*big.Int, error) {
	return _Poll.Contract.TotalRevealed(&_Poll.CallOpts)
}

// TotalRevealed is a free data retrieval call binding the contract method 0xf3738569.
//
// Solidity: function totalRevealed() view returns(uint256)
func (_Poll *PollCallerSession) TotalRevealed() (*big.Int, error) {
	return _Poll.Contract.TotalRevealed(&_Poll.CallOpts)
}

// VoterMerkleRoot is a free data retrieval call binding the contract method 0x2a69fb46.
//
// Solidity: function voterMerkleRoot() view returns(bytes32)
func (_Poll *PollCaller) VoterMerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Poll.contract.Call(opts, &out, "voterMerkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VoterMerkleRoot is a free data retrieval call binding the contract method 0x2a69fb46.
//
// Solidity: function voterMerkleRoot() view returns(bytes32)
func (_Poll *PollSession) VoterMerkleRoot() ([32]byte, error) {
	return _Poll.Contract.VoterMerkleRoot(&_Poll.CallOpts)
}

// VoterMerkleRoot is a free data retrieval call binding the contract method 0x2a69fb46.
//
// Solidity: function voterMerkleRoot() view returns(bytes32)
func (_Poll *PollCallerSession) VoterMerkleRoot() ([32]byte, error) {
	return _Poll.Contract.VoterMerkleRoot(&_Poll.CallOpts)
}

// ZkVerifier is a free data retrieval call binding the contract method 0xd6df096d.
//
// Solidity: function zkVerifier() view returns(address)
func (_Poll *PollCaller) ZkVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Poll.contract.Call(opts, &out, "zkVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZkVerifier is a free data retrieval call binding the contract method 0xd6df096d.
//
// Solidity: function zkVerifier() view returns(address)
func (_Poll *PollSession) ZkVerifier() (common.Address, error) {
	return _Poll.Contract.ZkVerifier(&_Poll.CallOpts)
}

// ZkVerifier is a free data retrieval call binding the contract method 0xd6df096d.
//
// Solidity: function zkVerifier() view returns(address)
func (_Poll *PollCallerSession) ZkVerifier() (common.Address, error) {
	return _Poll.Contract.ZkVerifier(&_Poll.CallOpts)
}

// ClosePoll is a paid mutator transaction binding the contract method 0xed8c2aed.
//
// Solidity: function closePoll() returns()
func (_Poll *PollTransactor) ClosePoll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poll.contract.Transact(opts, "closePoll")
}

// ClosePoll is a paid mutator transaction binding the contract method 0xed8c2aed.
//
// Solidity: function closePoll() returns()
func (_Poll *PollSession) ClosePoll() (*types.Transaction, error) {
	return _Poll.Contract.ClosePoll(&_Poll.TransactOpts)
}

// ClosePoll is a paid mutator transaction binding the contract method 0xed8c2aed.
//
// Solidity: function closePoll() returns()
func (_Poll *PollTransactorSession) ClosePoll() (*types.Transaction, error) {
	return _Poll.Contract.ClosePoll(&_Poll.TransactOpts)
}

// CommitVote is a paid mutator transaction binding the contract method 0x713916bd.
//
// Solidity: function commitVote(bytes32 commitment, bytes zkProof, bytes32[] merklePath) returns()
func (_Poll *PollTransactor) CommitVote(opts *bind.TransactOpts, commitment [32]byte, zkProof []byte, merklePath [][32]byte) (*types.Transaction, error) {
	return _Poll.contract.Transact(opts, "commitVote", commitment, zkProof, merklePath)
}

// CommitVote is a paid mutator transaction binding the contract method 0x713916bd.
//
// Solidity: function commitVote(bytes32 commitment, bytes zkProof, bytes32[] merklePath) returns()
func (_Poll *PollSession) CommitVote(commitment [32]byte, zkProof []byte, merklePath [][32]byte) (*types.Transaction, error) {
	return _Poll.Contract.CommitVote(&_Poll.TransactOpts, commitment, zkProof, merklePath)
}

// CommitVote is a paid mutator transaction binding the contract method 0x713916bd.
//
// Solidity: function commitVote(bytes32 commitment, bytes zkProof, bytes32[] merklePath) returns()
func (_Poll *PollTransactorSession) CommitVote(commitment [32]byte, zkProof []byte, merklePath [][32]byte) (*types.Transaction, error) {
	return _Poll.Contract.CommitVote(&_Poll.TransactOpts, commitment, zkProof, merklePath)
}

// RevealVote is a paid mutator transaction binding the contract method 0x69edbcda.
//
// Solidity: function revealVote(uint256 choice, bytes32 salt) returns()
func (_Poll *PollTransactor) RevealVote(opts *bind.TransactOpts, choice *big.Int, salt [32]byte) (*types.Transaction, error) {
	return _Poll.contract.Transact(opts, "revealVote", choice, salt)
}

// RevealVote is a paid mutator transaction binding the contract method 0x69edbcda.
//
// Solidity: function revealVote(uint256 choice, bytes32 salt) returns()
func (_Poll *PollSession) RevealVote(choice *big.Int, salt [32]byte) (*types.Transaction, error) {
	return _Poll.Contract.RevealVote(&_Poll.TransactOpts, choice, salt)
}

// RevealVote is a paid mutator transaction binding the contract method 0x69edbcda.
//
// Solidity: function revealVote(uint256 choice, bytes32 salt) returns()
func (_Poll *PollTransactorSession) RevealVote(choice *big.Int, salt [32]byte) (*types.Transaction, error) {
	return _Poll.Contract.RevealVote(&_Poll.TransactOpts, choice, salt)
}

// Tally is a paid mutator transaction binding the contract method 0x410673e5.
//
// Solidity: function tally() returns(uint256[])
func (_Poll *PollTransactor) Tally(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poll.contract.Transact(opts, "tally")
}

// Tally is a paid mutator transaction binding the contract method 0x410673e5.
//
// Solidity: function tally() returns(uint256[])
func (_Poll *PollSession) Tally() (*types.Transaction, error) {
	return _Poll.Contract.Tally(&_Poll.TransactOpts)
}

// Tally is a paid mutator transaction binding the contract method 0x410673e5.
//
// Solidity: function tally() returns(uint256[])
func (_Poll *PollTransactorSession) Tally() (*types.Transaction, error) {
	return _Poll.Contract.Tally(&_Poll.TransactOpts)
}

// PollPollClosedIterator is returned from FilterPollClosed and is used to iterate over the raw logs and unpacked data for PollClosed events raised by the Poll contract.
type PollPollClosedIterator struct {
	Event *PollPollClosed // Event containing the contract specifics and raw log

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
func (it *PollPollClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PollPollClosed)
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
		it.Event = new(PollPollClosed)
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
func (it *PollPollClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PollPollClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PollPollClosed represents a PollClosed event raised by the Poll contract.
type PollPollClosed struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPollClosed is a free log retrieval operation binding the contract event 0x703c67c157a1afb7d386ce5212ee88bf0a690bca998060e9595ebdfd747100ce.
//
// Solidity: event PollClosed(uint256 timestamp)
func (_Poll *PollFilterer) FilterPollClosed(opts *bind.FilterOpts) (*PollPollClosedIterator, error) {

	logs, sub, err := _Poll.contract.FilterLogs(opts, "PollClosed")
	if err != nil {
		return nil, err
	}
	return &PollPollClosedIterator{contract: _Poll.contract, event: "PollClosed", logs: logs, sub: sub}, nil
}

// WatchPollClosed is a free log subscription operation binding the contract event 0x703c67c157a1afb7d386ce5212ee88bf0a690bca998060e9595ebdfd747100ce.
//
// Solidity: event PollClosed(uint256 timestamp)
func (_Poll *PollFilterer) WatchPollClosed(opts *bind.WatchOpts, sink chan<- *PollPollClosed) (event.Subscription, error) {

	logs, sub, err := _Poll.contract.WatchLogs(opts, "PollClosed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PollPollClosed)
				if err := _Poll.contract.UnpackLog(event, "PollClosed", log); err != nil {
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

// ParsePollClosed is a log parse operation binding the contract event 0x703c67c157a1afb7d386ce5212ee88bf0a690bca998060e9595ebdfd747100ce.
//
// Solidity: event PollClosed(uint256 timestamp)
func (_Poll *PollFilterer) ParsePollClosed(log types.Log) (*PollPollClosed, error) {
	event := new(PollPollClosed)
	if err := _Poll.contract.UnpackLog(event, "PollClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PollResultsTalliedIterator is returned from FilterResultsTallied and is used to iterate over the raw logs and unpacked data for ResultsTallied events raised by the Poll contract.
type PollResultsTalliedIterator struct {
	Event *PollResultsTallied // Event containing the contract specifics and raw log

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
func (it *PollResultsTalliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PollResultsTallied)
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
		it.Event = new(PollResultsTallied)
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
func (it *PollResultsTalliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PollResultsTalliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PollResultsTallied represents a ResultsTallied event raised by the Poll contract.
type PollResultsTallied struct {
	Results   []*big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterResultsTallied is a free log retrieval operation binding the contract event 0x5408f1d0b104f82c19a2174a2056a36345c26e44e0e7df271f5aff18e9a6ce44.
//
// Solidity: event ResultsTallied(uint256[] results, uint256 timestamp)
func (_Poll *PollFilterer) FilterResultsTallied(opts *bind.FilterOpts) (*PollResultsTalliedIterator, error) {

	logs, sub, err := _Poll.contract.FilterLogs(opts, "ResultsTallied")
	if err != nil {
		return nil, err
	}
	return &PollResultsTalliedIterator{contract: _Poll.contract, event: "ResultsTallied", logs: logs, sub: sub}, nil
}

// WatchResultsTallied is a free log subscription operation binding the contract event 0x5408f1d0b104f82c19a2174a2056a36345c26e44e0e7df271f5aff18e9a6ce44.
//
// Solidity: event ResultsTallied(uint256[] results, uint256 timestamp)
func (_Poll *PollFilterer) WatchResultsTallied(opts *bind.WatchOpts, sink chan<- *PollResultsTallied) (event.Subscription, error) {

	logs, sub, err := _Poll.contract.WatchLogs(opts, "ResultsTallied")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PollResultsTallied)
				if err := _Poll.contract.UnpackLog(event, "ResultsTallied", log); err != nil {
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

// ParseResultsTallied is a log parse operation binding the contract event 0x5408f1d0b104f82c19a2174a2056a36345c26e44e0e7df271f5aff18e9a6ce44.
//
// Solidity: event ResultsTallied(uint256[] results, uint256 timestamp)
func (_Poll *PollFilterer) ParseResultsTallied(log types.Log) (*PollResultsTallied, error) {
	event := new(PollResultsTallied)
	if err := _Poll.contract.UnpackLog(event, "ResultsTallied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PollVoteCommittedIterator is returned from FilterVoteCommitted and is used to iterate over the raw logs and unpacked data for VoteCommitted events raised by the Poll contract.
type PollVoteCommittedIterator struct {
	Event *PollVoteCommitted // Event containing the contract specifics and raw log

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
func (it *PollVoteCommittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PollVoteCommitted)
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
		it.Event = new(PollVoteCommitted)
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
func (it *PollVoteCommittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PollVoteCommittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PollVoteCommitted represents a VoteCommitted event raised by the Poll contract.
type PollVoteCommitted struct {
	Voter      common.Address
	Commitment [32]byte
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteCommitted is a free log retrieval operation binding the contract event 0x06fdcf30f0bb2c4ceab314db052ad51b198e3ff93ce4121183516fa1ae84fbc8.
//
// Solidity: event VoteCommitted(address indexed voter, bytes32 commitment, uint256 timestamp)
func (_Poll *PollFilterer) FilterVoteCommitted(opts *bind.FilterOpts, voter []common.Address) (*PollVoteCommittedIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Poll.contract.FilterLogs(opts, "VoteCommitted", voterRule)
	if err != nil {
		return nil, err
	}
	return &PollVoteCommittedIterator{contract: _Poll.contract, event: "VoteCommitted", logs: logs, sub: sub}, nil
}

// WatchVoteCommitted is a free log subscription operation binding the contract event 0x06fdcf30f0bb2c4ceab314db052ad51b198e3ff93ce4121183516fa1ae84fbc8.
//
// Solidity: event VoteCommitted(address indexed voter, bytes32 commitment, uint256 timestamp)
func (_Poll *PollFilterer) WatchVoteCommitted(opts *bind.WatchOpts, sink chan<- *PollVoteCommitted, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Poll.contract.WatchLogs(opts, "VoteCommitted", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PollVoteCommitted)
				if err := _Poll.contract.UnpackLog(event, "VoteCommitted", log); err != nil {
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

// ParseVoteCommitted is a log parse operation binding the contract event 0x06fdcf30f0bb2c4ceab314db052ad51b198e3ff93ce4121183516fa1ae84fbc8.
//
// Solidity: event VoteCommitted(address indexed voter, bytes32 commitment, uint256 timestamp)
func (_Poll *PollFilterer) ParseVoteCommitted(log types.Log) (*PollVoteCommitted, error) {
	event := new(PollVoteCommitted)
	if err := _Poll.contract.UnpackLog(event, "VoteCommitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PollVoteRevealedIterator is returned from FilterVoteRevealed and is used to iterate over the raw logs and unpacked data for VoteRevealed events raised by the Poll contract.
type PollVoteRevealedIterator struct {
	Event *PollVoteRevealed // Event containing the contract specifics and raw log

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
func (it *PollVoteRevealedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PollVoteRevealed)
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
		it.Event = new(PollVoteRevealed)
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
func (it *PollVoteRevealedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PollVoteRevealedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PollVoteRevealed represents a VoteRevealed event raised by the Poll contract.
type PollVoteRevealed struct {
	Voter     common.Address
	Choice    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVoteRevealed is a free log retrieval operation binding the contract event 0xf65a04be847d83385f9a7abcf32cfb35055023a4affcc7eaa319160746f44528.
//
// Solidity: event VoteRevealed(address indexed voter, uint256 choice, uint256 timestamp)
func (_Poll *PollFilterer) FilterVoteRevealed(opts *bind.FilterOpts, voter []common.Address) (*PollVoteRevealedIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Poll.contract.FilterLogs(opts, "VoteRevealed", voterRule)
	if err != nil {
		return nil, err
	}
	return &PollVoteRevealedIterator{contract: _Poll.contract, event: "VoteRevealed", logs: logs, sub: sub}, nil
}

// WatchVoteRevealed is a free log subscription operation binding the contract event 0xf65a04be847d83385f9a7abcf32cfb35055023a4affcc7eaa319160746f44528.
//
// Solidity: event VoteRevealed(address indexed voter, uint256 choice, uint256 timestamp)
func (_Poll *PollFilterer) WatchVoteRevealed(opts *bind.WatchOpts, sink chan<- *PollVoteRevealed, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Poll.contract.WatchLogs(opts, "VoteRevealed", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PollVoteRevealed)
				if err := _Poll.contract.UnpackLog(event, "VoteRevealed", log); err != nil {
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

// ParseVoteRevealed is a log parse operation binding the contract event 0xf65a04be847d83385f9a7abcf32cfb35055023a4affcc7eaa319160746f44528.
//
// Solidity: event VoteRevealed(address indexed voter, uint256 choice, uint256 timestamp)
func (_Poll *PollFilterer) ParseVoteRevealed(log types.Log) (*PollVoteRevealed, error) {
	event := new(PollVoteRevealed)
	if err := _Poll.contract.UnpackLog(event, "VoteRevealed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
