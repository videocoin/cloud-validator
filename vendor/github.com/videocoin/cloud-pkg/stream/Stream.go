// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stream

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StreamABI is the input ABI used to generate the binding from.
const StreamABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getprofiles\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"client\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ended\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isChunk\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"profile\",\"type\":\"uint256\"},{\"name\":\"chunkId\",\"type\":\"uint256\"},{\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getProof\",\"outputs\":[{\"name\":\"miner\",\"type\":\"address\"},{\"name\":\"outputChunkId\",\"type\":\"uint256\"},{\"name\":\"proof\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"profile\",\"type\":\"uint256\"},{\"name\":\"chunkId\",\"type\":\"uint256\"}],\"name\":\"getCandidateProof\",\"outputs\":[{\"name\":\"miner\",\"type\":\"address\"},{\"name\":\"outputChunkId\",\"type\":\"uint256\"},{\"name\":\"proof\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"chunkId\",\"type\":\"uint256\"},{\"name\":\"wattage\",\"type\":\"uint256[]\"}],\"name\":\"addInputChunkId\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getProfileCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"refundAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"profile\",\"type\":\"uint256\"},{\"name\":\"chunkId\",\"type\":\"uint256\"}],\"name\":\"hasValidProof\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"refund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"profile\",\"type\":\"uint256\"}],\"name\":\"getOutChunks\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInChunkCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"profile\",\"type\":\"uint256\"},{\"name\":\"chunkId\",\"type\":\"uint256\"},{\"name\":\"proof\",\"type\":\"uint256\"},{\"name\":\"outChunkId\",\"type\":\"uint256\"}],\"name\":\"submitProof\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"profile\",\"type\":\"uint256\"},{\"name\":\"chunkId\",\"type\":\"uint256\"}],\"name\":\"getProofCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"profile\",\"type\":\"uint256\"}],\"name\":\"isBitrateTranscoded\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"id\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"endStream\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"profile\",\"type\":\"uint256\"},{\"name\":\"chunkId\",\"type\":\"uint256\"}],\"name\":\"scrapProof\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"profile\",\"type\":\"uint256\"},{\"name\":\"chunkId\",\"type\":\"uint256\"}],\"name\":\"getValidProof\",\"outputs\":[{\"name\":\"miner\",\"type\":\"address\"},{\"name\":\"validator\",\"type\":\"address\"},{\"name\":\"outputChunkId\",\"type\":\"uint256\"},{\"name\":\"proof\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"outStreams\",\"outputs\":[{\"name\":\"required\",\"type\":\"bool\"},{\"name\":\"index\",\"type\":\"uint256\"},{\"name\":\"validatedChunks\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInChunks\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"profile\",\"type\":\"uint256\"},{\"name\":\"chunkId\",\"type\":\"uint256\"}],\"name\":\"validateProof\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"wattages\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isTranscodingDone\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"client\",\"type\":\"address\"},{\"name\":\"profiles\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"chunkId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"profile\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"ChunkProofSubmited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"profile\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"chunkId\",\"type\":\"uint256\"}],\"name\":\"ChunkProofValidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"profile\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"chunkId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"ChunkProofScrapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"weiAmount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"weiAmount\",\"type\":\"uint256\"}],\"name\":\"Refunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"weiAmount\",\"type\":\"uint256\"}],\"name\":\"AccountFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"OutOfFunds\",\"type\":\"event\"}]"

// Stream is an auto generated Go binding around an Ethereum contract.
type Stream struct {
	StreamCaller     // Read-only binding to the contract
	StreamTransactor // Write-only binding to the contract
	StreamFilterer   // Log filterer for contract events
}

// StreamCaller is an auto generated read-only Go binding around an Ethereum contract.
type StreamCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StreamTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StreamTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StreamFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StreamFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StreamSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StreamSession struct {
	Contract     *Stream           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StreamCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StreamCallerSession struct {
	Contract *StreamCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StreamTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StreamTransactorSession struct {
	Contract     *StreamTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StreamRaw is an auto generated low-level Go binding around an Ethereum contract.
type StreamRaw struct {
	Contract *Stream // Generic contract binding to access the raw methods on
}

// StreamCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StreamCallerRaw struct {
	Contract *StreamCaller // Generic read-only contract binding to access the raw methods on
}

// StreamTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StreamTransactorRaw struct {
	Contract *StreamTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStream creates a new instance of Stream, bound to a specific deployed contract.
func NewStream(address common.Address, backend bind.ContractBackend) (*Stream, error) {
	contract, err := bindStream(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stream{StreamCaller: StreamCaller{contract: contract}, StreamTransactor: StreamTransactor{contract: contract}, StreamFilterer: StreamFilterer{contract: contract}}, nil
}

// NewStreamCaller creates a new read-only instance of Stream, bound to a specific deployed contract.
func NewStreamCaller(address common.Address, caller bind.ContractCaller) (*StreamCaller, error) {
	contract, err := bindStream(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StreamCaller{contract: contract}, nil
}

// NewStreamTransactor creates a new write-only instance of Stream, bound to a specific deployed contract.
func NewStreamTransactor(address common.Address, transactor bind.ContractTransactor) (*StreamTransactor, error) {
	contract, err := bindStream(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StreamTransactor{contract: contract}, nil
}

// NewStreamFilterer creates a new log filterer instance of Stream, bound to a specific deployed contract.
func NewStreamFilterer(address common.Address, filterer bind.ContractFilterer) (*StreamFilterer, error) {
	contract, err := bindStream(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StreamFilterer{contract: contract}, nil
}

// bindStream binds a generic wrapper to an already deployed contract.
func bindStream(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StreamABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stream *StreamRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Stream.Contract.StreamCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stream *StreamRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stream.Contract.StreamTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stream *StreamRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stream.Contract.StreamTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stream *StreamCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Stream.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stream *StreamTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stream.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stream *StreamTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stream.Contract.contract.Transact(opts, method, params...)
}

// Client is a free data retrieval call binding the contract method 0x109e94cf.
//
// Solidity: function client() constant returns(address)
func (_Stream *StreamCaller) Client(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "client")
	return *ret0, err
}

// Client is a free data retrieval call binding the contract method 0x109e94cf.
//
// Solidity: function client() constant returns(address)
func (_Stream *StreamSession) Client() (common.Address, error) {
	return _Stream.Contract.Client(&_Stream.CallOpts)
}

// Client is a free data retrieval call binding the contract method 0x109e94cf.
//
// Solidity: function client() constant returns(address)
func (_Stream *StreamCallerSession) Client() (common.Address, error) {
	return _Stream.Contract.Client(&_Stream.CallOpts)
}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() constant returns(bool)
func (_Stream *StreamCaller) Ended(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "ended")
	return *ret0, err
}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() constant returns(bool)
func (_Stream *StreamSession) Ended() (bool, error) {
	return _Stream.Contract.Ended(&_Stream.CallOpts)
}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() constant returns(bool)
func (_Stream *StreamCallerSession) Ended() (bool, error) {
	return _Stream.Contract.Ended(&_Stream.CallOpts)
}

// GetCandidateProof is a free data retrieval call binding the contract method 0x2f750f20.
//
// Solidity: function getCandidateProof(uint256 profile, uint256 chunkId) constant returns(address miner, uint256 outputChunkId, uint256 proof)
func (_Stream *StreamCaller) GetCandidateProof(opts *bind.CallOpts, profile *big.Int, chunkId *big.Int) (struct {
	Miner         common.Address
	OutputChunkId *big.Int
	Proof         *big.Int
}, error) {
	ret := new(struct {
		Miner         common.Address
		OutputChunkId *big.Int
		Proof         *big.Int
	})
	out := ret
	err := _Stream.contract.Call(opts, out, "getCandidateProof", profile, chunkId)
	return *ret, err
}

// GetCandidateProof is a free data retrieval call binding the contract method 0x2f750f20.
//
// Solidity: function getCandidateProof(uint256 profile, uint256 chunkId) constant returns(address miner, uint256 outputChunkId, uint256 proof)
func (_Stream *StreamSession) GetCandidateProof(profile *big.Int, chunkId *big.Int) (struct {
	Miner         common.Address
	OutputChunkId *big.Int
	Proof         *big.Int
}, error) {
	return _Stream.Contract.GetCandidateProof(&_Stream.CallOpts, profile, chunkId)
}

// GetCandidateProof is a free data retrieval call binding the contract method 0x2f750f20.
//
// Solidity: function getCandidateProof(uint256 profile, uint256 chunkId) constant returns(address miner, uint256 outputChunkId, uint256 proof)
func (_Stream *StreamCallerSession) GetCandidateProof(profile *big.Int, chunkId *big.Int) (struct {
	Miner         common.Address
	OutputChunkId *big.Int
	Proof         *big.Int
}, error) {
	return _Stream.Contract.GetCandidateProof(&_Stream.CallOpts, profile, chunkId)
}

// GetInChunkCount is a free data retrieval call binding the contract method 0x73f93b2a.
//
// Solidity: function getInChunkCount() constant returns(uint256)
func (_Stream *StreamCaller) GetInChunkCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "getInChunkCount")
	return *ret0, err
}

// GetInChunkCount is a free data retrieval call binding the contract method 0x73f93b2a.
//
// Solidity: function getInChunkCount() constant returns(uint256)
func (_Stream *StreamSession) GetInChunkCount() (*big.Int, error) {
	return _Stream.Contract.GetInChunkCount(&_Stream.CallOpts)
}

// GetInChunkCount is a free data retrieval call binding the contract method 0x73f93b2a.
//
// Solidity: function getInChunkCount() constant returns(uint256)
func (_Stream *StreamCallerSession) GetInChunkCount() (*big.Int, error) {
	return _Stream.Contract.GetInChunkCount(&_Stream.CallOpts)
}

// GetInChunks is a free data retrieval call binding the contract method 0xc6175193.
//
// Solidity: function getInChunks() constant returns(uint256[])
func (_Stream *StreamCaller) GetInChunks(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "getInChunks")
	return *ret0, err
}

// GetInChunks is a free data retrieval call binding the contract method 0xc6175193.
//
// Solidity: function getInChunks() constant returns(uint256[])
func (_Stream *StreamSession) GetInChunks() ([]*big.Int, error) {
	return _Stream.Contract.GetInChunks(&_Stream.CallOpts)
}

// GetInChunks is a free data retrieval call binding the contract method 0xc6175193.
//
// Solidity: function getInChunks() constant returns(uint256[])
func (_Stream *StreamCallerSession) GetInChunks() ([]*big.Int, error) {
	return _Stream.Contract.GetInChunks(&_Stream.CallOpts)
}

// GetOutChunks is a free data retrieval call binding the contract method 0x62372298.
//
// Solidity: function getOutChunks(uint256 profile) constant returns(uint256[])
func (_Stream *StreamCaller) GetOutChunks(opts *bind.CallOpts, profile *big.Int) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "getOutChunks", profile)
	return *ret0, err
}

// GetOutChunks is a free data retrieval call binding the contract method 0x62372298.
//
// Solidity: function getOutChunks(uint256 profile) constant returns(uint256[])
func (_Stream *StreamSession) GetOutChunks(profile *big.Int) ([]*big.Int, error) {
	return _Stream.Contract.GetOutChunks(&_Stream.CallOpts, profile)
}

// GetOutChunks is a free data retrieval call binding the contract method 0x62372298.
//
// Solidity: function getOutChunks(uint256 profile) constant returns(uint256[])
func (_Stream *StreamCallerSession) GetOutChunks(profile *big.Int) ([]*big.Int, error) {
	return _Stream.Contract.GetOutChunks(&_Stream.CallOpts, profile)
}

// GetProfileCount is a free data retrieval call binding the contract method 0x3697611a.
//
// Solidity: function getProfileCount() constant returns(uint256)
func (_Stream *StreamCaller) GetProfileCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "getProfileCount")
	return *ret0, err
}

// GetProfileCount is a free data retrieval call binding the contract method 0x3697611a.
//
// Solidity: function getProfileCount() constant returns(uint256)
func (_Stream *StreamSession) GetProfileCount() (*big.Int, error) {
	return _Stream.Contract.GetProfileCount(&_Stream.CallOpts)
}

// GetProfileCount is a free data retrieval call binding the contract method 0x3697611a.
//
// Solidity: function getProfileCount() constant returns(uint256)
func (_Stream *StreamCallerSession) GetProfileCount() (*big.Int, error) {
	return _Stream.Contract.GetProfileCount(&_Stream.CallOpts)
}

// GetProof is a free data retrieval call binding the contract method 0x28cc413a.
//
// Solidity: function getProof(uint256 profile, uint256 chunkId, uint256 idx) constant returns(address miner, uint256 outputChunkId, uint256 proof)
func (_Stream *StreamCaller) GetProof(opts *bind.CallOpts, profile *big.Int, chunkId *big.Int, idx *big.Int) (struct {
	Miner         common.Address
	OutputChunkId *big.Int
	Proof         *big.Int
}, error) {
	ret := new(struct {
		Miner         common.Address
		OutputChunkId *big.Int
		Proof         *big.Int
	})
	out := ret
	err := _Stream.contract.Call(opts, out, "getProof", profile, chunkId, idx)
	return *ret, err
}

// GetProof is a free data retrieval call binding the contract method 0x28cc413a.
//
// Solidity: function getProof(uint256 profile, uint256 chunkId, uint256 idx) constant returns(address miner, uint256 outputChunkId, uint256 proof)
func (_Stream *StreamSession) GetProof(profile *big.Int, chunkId *big.Int, idx *big.Int) (struct {
	Miner         common.Address
	OutputChunkId *big.Int
	Proof         *big.Int
}, error) {
	return _Stream.Contract.GetProof(&_Stream.CallOpts, profile, chunkId, idx)
}

// GetProof is a free data retrieval call binding the contract method 0x28cc413a.
//
// Solidity: function getProof(uint256 profile, uint256 chunkId, uint256 idx) constant returns(address miner, uint256 outputChunkId, uint256 proof)
func (_Stream *StreamCallerSession) GetProof(profile *big.Int, chunkId *big.Int, idx *big.Int) (struct {
	Miner         common.Address
	OutputChunkId *big.Int
	Proof         *big.Int
}, error) {
	return _Stream.Contract.GetProof(&_Stream.CallOpts, profile, chunkId, idx)
}

// GetProofCount is a free data retrieval call binding the contract method 0x7b40855d.
//
// Solidity: function getProofCount(uint256 profile, uint256 chunkId) constant returns(uint256)
func (_Stream *StreamCaller) GetProofCount(opts *bind.CallOpts, profile *big.Int, chunkId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "getProofCount", profile, chunkId)
	return *ret0, err
}

// GetProofCount is a free data retrieval call binding the contract method 0x7b40855d.
//
// Solidity: function getProofCount(uint256 profile, uint256 chunkId) constant returns(uint256)
func (_Stream *StreamSession) GetProofCount(profile *big.Int, chunkId *big.Int) (*big.Int, error) {
	return _Stream.Contract.GetProofCount(&_Stream.CallOpts, profile, chunkId)
}

// GetProofCount is a free data retrieval call binding the contract method 0x7b40855d.
//
// Solidity: function getProofCount(uint256 profile, uint256 chunkId) constant returns(uint256)
func (_Stream *StreamCallerSession) GetProofCount(profile *big.Int, chunkId *big.Int) (*big.Int, error) {
	return _Stream.Contract.GetProofCount(&_Stream.CallOpts, profile, chunkId)
}

// GetValidProof is a free data retrieval call binding the contract method 0xbf032f53.
//
// Solidity: function getValidProof(uint256 profile, uint256 chunkId) constant returns(address miner, address validator, uint256 outputChunkId, uint256 proof)
func (_Stream *StreamCaller) GetValidProof(opts *bind.CallOpts, profile *big.Int, chunkId *big.Int) (struct {
	Miner         common.Address
	Validator     common.Address
	OutputChunkId *big.Int
	Proof         *big.Int
}, error) {
	ret := new(struct {
		Miner         common.Address
		Validator     common.Address
		OutputChunkId *big.Int
		Proof         *big.Int
	})
	out := ret
	err := _Stream.contract.Call(opts, out, "getValidProof", profile, chunkId)
	return *ret, err
}

// GetValidProof is a free data retrieval call binding the contract method 0xbf032f53.
//
// Solidity: function getValidProof(uint256 profile, uint256 chunkId) constant returns(address miner, address validator, uint256 outputChunkId, uint256 proof)
func (_Stream *StreamSession) GetValidProof(profile *big.Int, chunkId *big.Int) (struct {
	Miner         common.Address
	Validator     common.Address
	OutputChunkId *big.Int
	Proof         *big.Int
}, error) {
	return _Stream.Contract.GetValidProof(&_Stream.CallOpts, profile, chunkId)
}

// GetValidProof is a free data retrieval call binding the contract method 0xbf032f53.
//
// Solidity: function getValidProof(uint256 profile, uint256 chunkId) constant returns(address miner, address validator, uint256 outputChunkId, uint256 proof)
func (_Stream *StreamCallerSession) GetValidProof(profile *big.Int, chunkId *big.Int) (struct {
	Miner         common.Address
	Validator     common.Address
	OutputChunkId *big.Int
	Proof         *big.Int
}, error) {
	return _Stream.Contract.GetValidProof(&_Stream.CallOpts, profile, chunkId)
}

// Getprofiles is a free data retrieval call binding the contract method 0x00ca5d92.
//
// Solidity: function getprofiles() constant returns(uint256[])
func (_Stream *StreamCaller) Getprofiles(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "getprofiles")
	return *ret0, err
}

// Getprofiles is a free data retrieval call binding the contract method 0x00ca5d92.
//
// Solidity: function getprofiles() constant returns(uint256[])
func (_Stream *StreamSession) Getprofiles() ([]*big.Int, error) {
	return _Stream.Contract.Getprofiles(&_Stream.CallOpts)
}

// Getprofiles is a free data retrieval call binding the contract method 0x00ca5d92.
//
// Solidity: function getprofiles() constant returns(uint256[])
func (_Stream *StreamCallerSession) Getprofiles() ([]*big.Int, error) {
	return _Stream.Contract.Getprofiles(&_Stream.CallOpts)
}

// HasValidProof is a free data retrieval call binding the contract method 0x4c0b715c.
//
// Solidity: function hasValidProof(uint256 profile, uint256 chunkId) constant returns(bool)
func (_Stream *StreamCaller) HasValidProof(opts *bind.CallOpts, profile *big.Int, chunkId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "hasValidProof", profile, chunkId)
	return *ret0, err
}

// HasValidProof is a free data retrieval call binding the contract method 0x4c0b715c.
//
// Solidity: function hasValidProof(uint256 profile, uint256 chunkId) constant returns(bool)
func (_Stream *StreamSession) HasValidProof(profile *big.Int, chunkId *big.Int) (bool, error) {
	return _Stream.Contract.HasValidProof(&_Stream.CallOpts, profile, chunkId)
}

// HasValidProof is a free data retrieval call binding the contract method 0x4c0b715c.
//
// Solidity: function hasValidProof(uint256 profile, uint256 chunkId) constant returns(bool)
func (_Stream *StreamCallerSession) HasValidProof(profile *big.Int, chunkId *big.Int) (bool, error) {
	return _Stream.Contract.HasValidProof(&_Stream.CallOpts, profile, chunkId)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() constant returns(uint256)
func (_Stream *StreamCaller) Id(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "id")
	return *ret0, err
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() constant returns(uint256)
func (_Stream *StreamSession) Id() (*big.Int, error) {
	return _Stream.Contract.Id(&_Stream.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() constant returns(uint256)
func (_Stream *StreamCallerSession) Id() (*big.Int, error) {
	return _Stream.Contract.Id(&_Stream.CallOpts)
}

// IsBitrateTranscoded is a free data retrieval call binding the contract method 0x963dce43.
//
// Solidity: function isBitrateTranscoded(uint256 profile) constant returns(bool)
func (_Stream *StreamCaller) IsBitrateTranscoded(opts *bind.CallOpts, profile *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "isBitrateTranscoded", profile)
	return *ret0, err
}

// IsBitrateTranscoded is a free data retrieval call binding the contract method 0x963dce43.
//
// Solidity: function isBitrateTranscoded(uint256 profile) constant returns(bool)
func (_Stream *StreamSession) IsBitrateTranscoded(profile *big.Int) (bool, error) {
	return _Stream.Contract.IsBitrateTranscoded(&_Stream.CallOpts, profile)
}

// IsBitrateTranscoded is a free data retrieval call binding the contract method 0x963dce43.
//
// Solidity: function isBitrateTranscoded(uint256 profile) constant returns(bool)
func (_Stream *StreamCallerSession) IsBitrateTranscoded(profile *big.Int) (bool, error) {
	return _Stream.Contract.IsBitrateTranscoded(&_Stream.CallOpts, profile)
}

// IsChunk is a free data retrieval call binding the contract method 0x1bb62fc4.
//
// Solidity: function isChunk(uint256 ) constant returns(bool)
func (_Stream *StreamCaller) IsChunk(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "isChunk", arg0)
	return *ret0, err
}

// IsChunk is a free data retrieval call binding the contract method 0x1bb62fc4.
//
// Solidity: function isChunk(uint256 ) constant returns(bool)
func (_Stream *StreamSession) IsChunk(arg0 *big.Int) (bool, error) {
	return _Stream.Contract.IsChunk(&_Stream.CallOpts, arg0)
}

// IsChunk is a free data retrieval call binding the contract method 0x1bb62fc4.
//
// Solidity: function isChunk(uint256 ) constant returns(bool)
func (_Stream *StreamCallerSession) IsChunk(arg0 *big.Int) (bool, error) {
	return _Stream.Contract.IsChunk(&_Stream.CallOpts, arg0)
}

// IsTranscodingDone is a free data retrieval call binding the contract method 0xfc1028bc.
//
// Solidity: function isTranscodingDone() constant returns(bool)
func (_Stream *StreamCaller) IsTranscodingDone(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "isTranscodingDone")
	return *ret0, err
}

// IsTranscodingDone is a free data retrieval call binding the contract method 0xfc1028bc.
//
// Solidity: function isTranscodingDone() constant returns(bool)
func (_Stream *StreamSession) IsTranscodingDone() (bool, error) {
	return _Stream.Contract.IsTranscodingDone(&_Stream.CallOpts)
}

// IsTranscodingDone is a free data retrieval call binding the contract method 0xfc1028bc.
//
// Solidity: function isTranscodingDone() constant returns(bool)
func (_Stream *StreamCallerSession) IsTranscodingDone() (bool, error) {
	return _Stream.Contract.IsTranscodingDone(&_Stream.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_Stream *StreamCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "manager")
	return *ret0, err
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_Stream *StreamSession) Manager() (common.Address, error) {
	return _Stream.Contract.Manager(&_Stream.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_Stream *StreamCallerSession) Manager() (common.Address, error) {
	return _Stream.Contract.Manager(&_Stream.CallOpts)
}

// OutStreams is a free data retrieval call binding the contract method 0xc5d0b14c.
//
// Solidity: function outStreams(uint256 ) constant returns(bool required, uint256 index, uint256 validatedChunks)
func (_Stream *StreamCaller) OutStreams(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Required        bool
	Index           *big.Int
	ValidatedChunks *big.Int
}, error) {
	ret := new(struct {
		Required        bool
		Index           *big.Int
		ValidatedChunks *big.Int
	})
	out := ret
	err := _Stream.contract.Call(opts, out, "outStreams", arg0)
	return *ret, err
}

// OutStreams is a free data retrieval call binding the contract method 0xc5d0b14c.
//
// Solidity: function outStreams(uint256 ) constant returns(bool required, uint256 index, uint256 validatedChunks)
func (_Stream *StreamSession) OutStreams(arg0 *big.Int) (struct {
	Required        bool
	Index           *big.Int
	ValidatedChunks *big.Int
}, error) {
	return _Stream.Contract.OutStreams(&_Stream.CallOpts, arg0)
}

// OutStreams is a free data retrieval call binding the contract method 0xc5d0b14c.
//
// Solidity: function outStreams(uint256 ) constant returns(bool required, uint256 index, uint256 validatedChunks)
func (_Stream *StreamCallerSession) OutStreams(arg0 *big.Int) (struct {
	Required        bool
	Index           *big.Int
	ValidatedChunks *big.Int
}, error) {
	return _Stream.Contract.OutStreams(&_Stream.CallOpts, arg0)
}

// RefundAllowed is a free data retrieval call binding the contract method 0x3fa911ae.
//
// Solidity: function refundAllowed() constant returns(bool)
func (_Stream *StreamCaller) RefundAllowed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "refundAllowed")
	return *ret0, err
}

// RefundAllowed is a free data retrieval call binding the contract method 0x3fa911ae.
//
// Solidity: function refundAllowed() constant returns(bool)
func (_Stream *StreamSession) RefundAllowed() (bool, error) {
	return _Stream.Contract.RefundAllowed(&_Stream.CallOpts)
}

// RefundAllowed is a free data retrieval call binding the contract method 0x3fa911ae.
//
// Solidity: function refundAllowed() constant returns(bool)
func (_Stream *StreamCallerSession) RefundAllowed() (bool, error) {
	return _Stream.Contract.RefundAllowed(&_Stream.CallOpts)
}

// Wattages is a free data retrieval call binding the contract method 0xeda0ce17.
//
// Solidity: function wattages(uint256 , uint256 ) constant returns(uint256)
func (_Stream *StreamCaller) Wattages(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "wattages", arg0, arg1)
	return *ret0, err
}

// Wattages is a free data retrieval call binding the contract method 0xeda0ce17.
//
// Solidity: function wattages(uint256 , uint256 ) constant returns(uint256)
func (_Stream *StreamSession) Wattages(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Stream.Contract.Wattages(&_Stream.CallOpts, arg0, arg1)
}

// Wattages is a free data retrieval call binding the contract method 0xeda0ce17.
//
// Solidity: function wattages(uint256 , uint256 ) constant returns(uint256)
func (_Stream *StreamCallerSession) Wattages(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Stream.Contract.Wattages(&_Stream.CallOpts, arg0, arg1)
}

// AddInputChunkId is a paid mutator transaction binding the contract method 0x356939ab.
//
// Solidity: function addInputChunkId(uint256 chunkId, uint256[] wattage) returns()
func (_Stream *StreamTransactor) AddInputChunkId(opts *bind.TransactOpts, chunkId *big.Int, wattage []*big.Int) (*types.Transaction, error) {
	return _Stream.contract.Transact(opts, "addInputChunkId", chunkId, wattage)
}

// AddInputChunkId is a paid mutator transaction binding the contract method 0x356939ab.
//
// Solidity: function addInputChunkId(uint256 chunkId, uint256[] wattage) returns()
func (_Stream *StreamSession) AddInputChunkId(chunkId *big.Int, wattage []*big.Int) (*types.Transaction, error) {
	return _Stream.Contract.AddInputChunkId(&_Stream.TransactOpts, chunkId, wattage)
}

// AddInputChunkId is a paid mutator transaction binding the contract method 0x356939ab.
//
// Solidity: function addInputChunkId(uint256 chunkId, uint256[] wattage) returns()
func (_Stream *StreamTransactorSession) AddInputChunkId(chunkId *big.Int, wattage []*big.Int) (*types.Transaction, error) {
	return _Stream.Contract.AddInputChunkId(&_Stream.TransactOpts, chunkId, wattage)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Stream *StreamTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stream.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Stream *StreamSession) Deposit() (*types.Transaction, error) {
	return _Stream.Contract.Deposit(&_Stream.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Stream *StreamTransactorSession) Deposit() (*types.Transaction, error) {
	return _Stream.Contract.Deposit(&_Stream.TransactOpts)
}

// EndStream is a paid mutator transaction binding the contract method 0xbb57a5ed.
//
// Solidity: function endStream() returns()
func (_Stream *StreamTransactor) EndStream(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stream.contract.Transact(opts, "endStream")
}

// EndStream is a paid mutator transaction binding the contract method 0xbb57a5ed.
//
// Solidity: function endStream() returns()
func (_Stream *StreamSession) EndStream() (*types.Transaction, error) {
	return _Stream.Contract.EndStream(&_Stream.TransactOpts)
}

// EndStream is a paid mutator transaction binding the contract method 0xbb57a5ed.
//
// Solidity: function endStream() returns()
func (_Stream *StreamTransactorSession) EndStream() (*types.Transaction, error) {
	return _Stream.Contract.EndStream(&_Stream.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_Stream *StreamTransactor) Refund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stream.contract.Transact(opts, "refund")
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_Stream *StreamSession) Refund() (*types.Transaction, error) {
	return _Stream.Contract.Refund(&_Stream.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_Stream *StreamTransactorSession) Refund() (*types.Transaction, error) {
	return _Stream.Contract.Refund(&_Stream.TransactOpts)
}

// ScrapProof is a paid mutator transaction binding the contract method 0xbbe58b0c.
//
// Solidity: function scrapProof(uint256 profile, uint256 chunkId) returns()
func (_Stream *StreamTransactor) ScrapProof(opts *bind.TransactOpts, profile *big.Int, chunkId *big.Int) (*types.Transaction, error) {
	return _Stream.contract.Transact(opts, "scrapProof", profile, chunkId)
}

// ScrapProof is a paid mutator transaction binding the contract method 0xbbe58b0c.
//
// Solidity: function scrapProof(uint256 profile, uint256 chunkId) returns()
func (_Stream *StreamSession) ScrapProof(profile *big.Int, chunkId *big.Int) (*types.Transaction, error) {
	return _Stream.Contract.ScrapProof(&_Stream.TransactOpts, profile, chunkId)
}

// ScrapProof is a paid mutator transaction binding the contract method 0xbbe58b0c.
//
// Solidity: function scrapProof(uint256 profile, uint256 chunkId) returns()
func (_Stream *StreamTransactorSession) ScrapProof(profile *big.Int, chunkId *big.Int) (*types.Transaction, error) {
	return _Stream.Contract.ScrapProof(&_Stream.TransactOpts, profile, chunkId)
}

// SubmitProof is a paid mutator transaction binding the contract method 0x747f7589.
//
// Solidity: function submitProof(uint256 profile, uint256 chunkId, uint256 proof, uint256 outChunkId) returns()
func (_Stream *StreamTransactor) SubmitProof(opts *bind.TransactOpts, profile *big.Int, chunkId *big.Int, proof *big.Int, outChunkId *big.Int) (*types.Transaction, error) {
	return _Stream.contract.Transact(opts, "submitProof", profile, chunkId, proof, outChunkId)
}

// SubmitProof is a paid mutator transaction binding the contract method 0x747f7589.
//
// Solidity: function submitProof(uint256 profile, uint256 chunkId, uint256 proof, uint256 outChunkId) returns()
func (_Stream *StreamSession) SubmitProof(profile *big.Int, chunkId *big.Int, proof *big.Int, outChunkId *big.Int) (*types.Transaction, error) {
	return _Stream.Contract.SubmitProof(&_Stream.TransactOpts, profile, chunkId, proof, outChunkId)
}

// SubmitProof is a paid mutator transaction binding the contract method 0x747f7589.
//
// Solidity: function submitProof(uint256 profile, uint256 chunkId, uint256 proof, uint256 outChunkId) returns()
func (_Stream *StreamTransactorSession) SubmitProof(profile *big.Int, chunkId *big.Int, proof *big.Int, outChunkId *big.Int) (*types.Transaction, error) {
	return _Stream.Contract.SubmitProof(&_Stream.TransactOpts, profile, chunkId, proof, outChunkId)
}

// ValidateProof is a paid mutator transaction binding the contract method 0xd78e647f.
//
// Solidity: function validateProof(uint256 profile, uint256 chunkId) returns()
func (_Stream *StreamTransactor) ValidateProof(opts *bind.TransactOpts, profile *big.Int, chunkId *big.Int) (*types.Transaction, error) {
	return _Stream.contract.Transact(opts, "validateProof", profile, chunkId)
}

// ValidateProof is a paid mutator transaction binding the contract method 0xd78e647f.
//
// Solidity: function validateProof(uint256 profile, uint256 chunkId) returns()
func (_Stream *StreamSession) ValidateProof(profile *big.Int, chunkId *big.Int) (*types.Transaction, error) {
	return _Stream.Contract.ValidateProof(&_Stream.TransactOpts, profile, chunkId)
}

// ValidateProof is a paid mutator transaction binding the contract method 0xd78e647f.
//
// Solidity: function validateProof(uint256 profile, uint256 chunkId) returns()
func (_Stream *StreamTransactorSession) ValidateProof(profile *big.Int, chunkId *big.Int) (*types.Transaction, error) {
	return _Stream.Contract.ValidateProof(&_Stream.TransactOpts, profile, chunkId)
}

// StreamAccountFundedIterator is returned from FilterAccountFunded and is used to iterate over the raw logs and unpacked data for AccountFunded events raised by the Stream contract.
type StreamAccountFundedIterator struct {
	Event *StreamAccountFunded // Event containing the contract specifics and raw log

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
func (it *StreamAccountFundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StreamAccountFunded)
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
		it.Event = new(StreamAccountFunded)
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
func (it *StreamAccountFundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StreamAccountFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StreamAccountFunded represents a AccountFunded event raised by the Stream contract.
type StreamAccountFunded struct {
	Account   common.Address
	WeiAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAccountFunded is a free log retrieval operation binding the contract event 0xbccbe05a3719eacef984f404dd2adff555adcc05fb72fb8b309bafbd462cd6f7.
//
// Solidity: event AccountFunded(address indexed account, uint256 weiAmount)
func (_Stream *StreamFilterer) FilterAccountFunded(opts *bind.FilterOpts, account []common.Address) (*StreamAccountFundedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Stream.contract.FilterLogs(opts, "AccountFunded", accountRule)
	if err != nil {
		return nil, err
	}
	return &StreamAccountFundedIterator{contract: _Stream.contract, event: "AccountFunded", logs: logs, sub: sub}, nil
}

// WatchAccountFunded is a free log subscription operation binding the contract event 0xbccbe05a3719eacef984f404dd2adff555adcc05fb72fb8b309bafbd462cd6f7.
//
// Solidity: event AccountFunded(address indexed account, uint256 weiAmount)
func (_Stream *StreamFilterer) WatchAccountFunded(opts *bind.WatchOpts, sink chan<- *StreamAccountFunded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Stream.contract.WatchLogs(opts, "AccountFunded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StreamAccountFunded)
				if err := _Stream.contract.UnpackLog(event, "AccountFunded", log); err != nil {
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

// StreamChunkProofScrappedIterator is returned from FilterChunkProofScrapped and is used to iterate over the raw logs and unpacked data for ChunkProofScrapped events raised by the Stream contract.
type StreamChunkProofScrappedIterator struct {
	Event *StreamChunkProofScrapped // Event containing the contract specifics and raw log

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
func (it *StreamChunkProofScrappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StreamChunkProofScrapped)
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
		it.Event = new(StreamChunkProofScrapped)
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
func (it *StreamChunkProofScrappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StreamChunkProofScrappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StreamChunkProofScrapped represents a ChunkProofScrapped event raised by the Stream contract.
type StreamChunkProofScrapped struct {
	Profile *big.Int
	ChunkId *big.Int
	Idx     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChunkProofScrapped is a free log retrieval operation binding the contract event 0xb25faf3bb93b5105c8bc184c6af20f251543b1fdd93cd5c6f9f31b13c138b91a.
//
// Solidity: event ChunkProofScrapped(uint256 indexed profile, uint256 indexed chunkId, uint256 indexed idx)
func (_Stream *StreamFilterer) FilterChunkProofScrapped(opts *bind.FilterOpts, profile []*big.Int, chunkId []*big.Int, idx []*big.Int) (*StreamChunkProofScrappedIterator, error) {

	var profileRule []interface{}
	for _, profileItem := range profile {
		profileRule = append(profileRule, profileItem)
	}
	var chunkIdRule []interface{}
	for _, chunkIdItem := range chunkId {
		chunkIdRule = append(chunkIdRule, chunkIdItem)
	}
	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _Stream.contract.FilterLogs(opts, "ChunkProofScrapped", profileRule, chunkIdRule, idxRule)
	if err != nil {
		return nil, err
	}
	return &StreamChunkProofScrappedIterator{contract: _Stream.contract, event: "ChunkProofScrapped", logs: logs, sub: sub}, nil
}

// WatchChunkProofScrapped is a free log subscription operation binding the contract event 0xb25faf3bb93b5105c8bc184c6af20f251543b1fdd93cd5c6f9f31b13c138b91a.
//
// Solidity: event ChunkProofScrapped(uint256 indexed profile, uint256 indexed chunkId, uint256 indexed idx)
func (_Stream *StreamFilterer) WatchChunkProofScrapped(opts *bind.WatchOpts, sink chan<- *StreamChunkProofScrapped, profile []*big.Int, chunkId []*big.Int, idx []*big.Int) (event.Subscription, error) {

	var profileRule []interface{}
	for _, profileItem := range profile {
		profileRule = append(profileRule, profileItem)
	}
	var chunkIdRule []interface{}
	for _, chunkIdItem := range chunkId {
		chunkIdRule = append(chunkIdRule, chunkIdItem)
	}
	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _Stream.contract.WatchLogs(opts, "ChunkProofScrapped", profileRule, chunkIdRule, idxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StreamChunkProofScrapped)
				if err := _Stream.contract.UnpackLog(event, "ChunkProofScrapped", log); err != nil {
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

// StreamChunkProofSubmitedIterator is returned from FilterChunkProofSubmited and is used to iterate over the raw logs and unpacked data for ChunkProofSubmited events raised by the Stream contract.
type StreamChunkProofSubmitedIterator struct {
	Event *StreamChunkProofSubmited // Event containing the contract specifics and raw log

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
func (it *StreamChunkProofSubmitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StreamChunkProofSubmited)
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
		it.Event = new(StreamChunkProofSubmited)
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
func (it *StreamChunkProofSubmitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StreamChunkProofSubmitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StreamChunkProofSubmited represents a ChunkProofSubmited event raised by the Stream contract.
type StreamChunkProofSubmited struct {
	ChunkId *big.Int
	Profile *big.Int
	Idx     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChunkProofSubmited is a free log retrieval operation binding the contract event 0xe176ecde2cccccb4849f78f5d9f8378a179f72033c2fc649f90820a324cfaa71.
//
// Solidity: event ChunkProofSubmited(uint256 indexed chunkId, uint256 indexed profile, uint256 indexed idx)
func (_Stream *StreamFilterer) FilterChunkProofSubmited(opts *bind.FilterOpts, chunkId []*big.Int, profile []*big.Int, idx []*big.Int) (*StreamChunkProofSubmitedIterator, error) {

	var chunkIdRule []interface{}
	for _, chunkIdItem := range chunkId {
		chunkIdRule = append(chunkIdRule, chunkIdItem)
	}
	var profileRule []interface{}
	for _, profileItem := range profile {
		profileRule = append(profileRule, profileItem)
	}
	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _Stream.contract.FilterLogs(opts, "ChunkProofSubmited", chunkIdRule, profileRule, idxRule)
	if err != nil {
		return nil, err
	}
	return &StreamChunkProofSubmitedIterator{contract: _Stream.contract, event: "ChunkProofSubmited", logs: logs, sub: sub}, nil
}

// WatchChunkProofSubmited is a free log subscription operation binding the contract event 0xe176ecde2cccccb4849f78f5d9f8378a179f72033c2fc649f90820a324cfaa71.
//
// Solidity: event ChunkProofSubmited(uint256 indexed chunkId, uint256 indexed profile, uint256 indexed idx)
func (_Stream *StreamFilterer) WatchChunkProofSubmited(opts *bind.WatchOpts, sink chan<- *StreamChunkProofSubmited, chunkId []*big.Int, profile []*big.Int, idx []*big.Int) (event.Subscription, error) {

	var chunkIdRule []interface{}
	for _, chunkIdItem := range chunkId {
		chunkIdRule = append(chunkIdRule, chunkIdItem)
	}
	var profileRule []interface{}
	for _, profileItem := range profile {
		profileRule = append(profileRule, profileItem)
	}
	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _Stream.contract.WatchLogs(opts, "ChunkProofSubmited", chunkIdRule, profileRule, idxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StreamChunkProofSubmited)
				if err := _Stream.contract.UnpackLog(event, "ChunkProofSubmited", log); err != nil {
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

// StreamChunkProofValidatedIterator is returned from FilterChunkProofValidated and is used to iterate over the raw logs and unpacked data for ChunkProofValidated events raised by the Stream contract.
type StreamChunkProofValidatedIterator struct {
	Event *StreamChunkProofValidated // Event containing the contract specifics and raw log

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
func (it *StreamChunkProofValidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StreamChunkProofValidated)
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
		it.Event = new(StreamChunkProofValidated)
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
func (it *StreamChunkProofValidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StreamChunkProofValidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StreamChunkProofValidated represents a ChunkProofValidated event raised by the Stream contract.
type StreamChunkProofValidated struct {
	Profile *big.Int
	ChunkId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChunkProofValidated is a free log retrieval operation binding the contract event 0x6e0dd2be4d72dbf12213042a47491e605ec38b66d076e6f74e9f458373080e37.
//
// Solidity: event ChunkProofValidated(uint256 indexed profile, uint256 indexed chunkId)
func (_Stream *StreamFilterer) FilterChunkProofValidated(opts *bind.FilterOpts, profile []*big.Int, chunkId []*big.Int) (*StreamChunkProofValidatedIterator, error) {

	var profileRule []interface{}
	for _, profileItem := range profile {
		profileRule = append(profileRule, profileItem)
	}
	var chunkIdRule []interface{}
	for _, chunkIdItem := range chunkId {
		chunkIdRule = append(chunkIdRule, chunkIdItem)
	}

	logs, sub, err := _Stream.contract.FilterLogs(opts, "ChunkProofValidated", profileRule, chunkIdRule)
	if err != nil {
		return nil, err
	}
	return &StreamChunkProofValidatedIterator{contract: _Stream.contract, event: "ChunkProofValidated", logs: logs, sub: sub}, nil
}

// WatchChunkProofValidated is a free log subscription operation binding the contract event 0x6e0dd2be4d72dbf12213042a47491e605ec38b66d076e6f74e9f458373080e37.
//
// Solidity: event ChunkProofValidated(uint256 indexed profile, uint256 indexed chunkId)
func (_Stream *StreamFilterer) WatchChunkProofValidated(opts *bind.WatchOpts, sink chan<- *StreamChunkProofValidated, profile []*big.Int, chunkId []*big.Int) (event.Subscription, error) {

	var profileRule []interface{}
	for _, profileItem := range profile {
		profileRule = append(profileRule, profileItem)
	}
	var chunkIdRule []interface{}
	for _, chunkIdItem := range chunkId {
		chunkIdRule = append(chunkIdRule, chunkIdItem)
	}

	logs, sub, err := _Stream.contract.WatchLogs(opts, "ChunkProofValidated", profileRule, chunkIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StreamChunkProofValidated)
				if err := _Stream.contract.UnpackLog(event, "ChunkProofValidated", log); err != nil {
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

// StreamDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Stream contract.
type StreamDepositedIterator struct {
	Event *StreamDeposited // Event containing the contract specifics and raw log

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
func (it *StreamDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StreamDeposited)
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
		it.Event = new(StreamDeposited)
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
func (it *StreamDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StreamDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StreamDeposited represents a Deposited event raised by the Stream contract.
type StreamDeposited struct {
	WeiAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2a89b2e3d580398d6dc2db5e0f336b52602bbaa51afa9bb5cdf59239cf0d2bea.
//
// Solidity: event Deposited(uint256 indexed weiAmount)
func (_Stream *StreamFilterer) FilterDeposited(opts *bind.FilterOpts, weiAmount []*big.Int) (*StreamDepositedIterator, error) {

	var weiAmountRule []interface{}
	for _, weiAmountItem := range weiAmount {
		weiAmountRule = append(weiAmountRule, weiAmountItem)
	}

	logs, sub, err := _Stream.contract.FilterLogs(opts, "Deposited", weiAmountRule)
	if err != nil {
		return nil, err
	}
	return &StreamDepositedIterator{contract: _Stream.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2a89b2e3d580398d6dc2db5e0f336b52602bbaa51afa9bb5cdf59239cf0d2bea.
//
// Solidity: event Deposited(uint256 indexed weiAmount)
func (_Stream *StreamFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *StreamDeposited, weiAmount []*big.Int) (event.Subscription, error) {

	var weiAmountRule []interface{}
	for _, weiAmountItem := range weiAmount {
		weiAmountRule = append(weiAmountRule, weiAmountItem)
	}

	logs, sub, err := _Stream.contract.WatchLogs(opts, "Deposited", weiAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StreamDeposited)
				if err := _Stream.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// StreamOutOfFundsIterator is returned from FilterOutOfFunds and is used to iterate over the raw logs and unpacked data for OutOfFunds events raised by the Stream contract.
type StreamOutOfFundsIterator struct {
	Event *StreamOutOfFunds // Event containing the contract specifics and raw log

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
func (it *StreamOutOfFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StreamOutOfFunds)
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
		it.Event = new(StreamOutOfFunds)
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
func (it *StreamOutOfFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StreamOutOfFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StreamOutOfFunds represents a OutOfFunds event raised by the Stream contract.
type StreamOutOfFunds struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOutOfFunds is a free log retrieval operation binding the contract event 0x5cc4d4e2397f909efdc05489f71c523c59913a6ffde292799faaca096635e16d.
//
// Solidity: event OutOfFunds()
func (_Stream *StreamFilterer) FilterOutOfFunds(opts *bind.FilterOpts) (*StreamOutOfFundsIterator, error) {

	logs, sub, err := _Stream.contract.FilterLogs(opts, "OutOfFunds")
	if err != nil {
		return nil, err
	}
	return &StreamOutOfFundsIterator{contract: _Stream.contract, event: "OutOfFunds", logs: logs, sub: sub}, nil
}

// WatchOutOfFunds is a free log subscription operation binding the contract event 0x5cc4d4e2397f909efdc05489f71c523c59913a6ffde292799faaca096635e16d.
//
// Solidity: event OutOfFunds()
func (_Stream *StreamFilterer) WatchOutOfFunds(opts *bind.WatchOpts, sink chan<- *StreamOutOfFunds) (event.Subscription, error) {

	logs, sub, err := _Stream.contract.WatchLogs(opts, "OutOfFunds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StreamOutOfFunds)
				if err := _Stream.contract.UnpackLog(event, "OutOfFunds", log); err != nil {
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

// StreamRefundedIterator is returned from FilterRefunded and is used to iterate over the raw logs and unpacked data for Refunded events raised by the Stream contract.
type StreamRefundedIterator struct {
	Event *StreamRefunded // Event containing the contract specifics and raw log

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
func (it *StreamRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StreamRefunded)
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
		it.Event = new(StreamRefunded)
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
func (it *StreamRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StreamRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StreamRefunded represents a Refunded event raised by the Stream contract.
type StreamRefunded struct {
	WeiAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRefunded is a free log retrieval operation binding the contract event 0x3d2a04f53164bedf9a8a46353305d6b2d2261410406df3b41f99ce6489dc003c.
//
// Solidity: event Refunded(uint256 weiAmount)
func (_Stream *StreamFilterer) FilterRefunded(opts *bind.FilterOpts) (*StreamRefundedIterator, error) {

	logs, sub, err := _Stream.contract.FilterLogs(opts, "Refunded")
	if err != nil {
		return nil, err
	}
	return &StreamRefundedIterator{contract: _Stream.contract, event: "Refunded", logs: logs, sub: sub}, nil
}

// WatchRefunded is a free log subscription operation binding the contract event 0x3d2a04f53164bedf9a8a46353305d6b2d2261410406df3b41f99ce6489dc003c.
//
// Solidity: event Refunded(uint256 weiAmount)
func (_Stream *StreamFilterer) WatchRefunded(opts *bind.WatchOpts, sink chan<- *StreamRefunded) (event.Subscription, error) {

	logs, sub, err := _Stream.contract.WatchLogs(opts, "Refunded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StreamRefunded)
				if err := _Stream.contract.UnpackLog(event, "Refunded", log); err != nil {
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
