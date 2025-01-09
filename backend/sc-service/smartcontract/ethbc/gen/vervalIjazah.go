// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verval_ijazah

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

// VervalIjazahMetaData contains all meta data concerning the VervalIjazah contract.
var VervalIjazahMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nisn\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_nama\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_tahunLulus\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"_nomor\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_nilaiTranskrip\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_cidURL\",\"type\":\"string\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nomor\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nomor\",\"type\":\"string\"}],\"name\":\"validasiIjazah\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610fc7806100616000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063693ec85e146100465780637d4392ee1461007b5780638f1c4cae14610097575b600080fd5b610060600480360381019061005b91906107da565b6100c7565b604051610072969594939291906108f5565b60405180910390f35b610095600480360381019061009091906109de565b610422565b005b6100b160048036038101906100ac91906107da565b610649565b6040516100be9190610adb565b60405180910390f35b60608060006060600080600080886040516100e29190610b32565b90815260200160405180910390206040518060e001604052908160008201805461010b90610b78565b80601f016020809104026020016040519081016040528092919081815260200182805461013790610b78565b80156101845780601f1061015957610100808354040283529160200191610184565b820191906000526020600020905b81548152906001019060200180831161016757829003601f168201915b5050505050815260200160018201805461019d90610b78565b80601f01602080910402602001604051908101604052809291908181526020018280546101c990610b78565b80156102165780601f106101eb57610100808354040283529160200191610216565b820191906000526020600020905b8154815290600101906020018083116101f957829003601f168201915b5050505050815260200160028201805461022f90610b78565b80601f016020809104026020016040519081016040528092919081815260200182805461025b90610b78565b80156102a85780601f1061027d576101008083540402835291602001916102a8565b820191906000526020600020905b81548152906001019060200180831161028b57829003601f168201915b505050505081526020016003820160009054906101000a900461ffff1661ffff1661ffff1681526020016003820160029054906101000a900460ff1660ff1660ff1681526020016004820180546102fe90610b78565b80601f016020809104026020016040519081016040528092919081815260200182805461032a90610b78565b80156103775780601f1061034c57610100808354040283529160200191610377565b820191906000526020600020905b81548152906001019060200180831161035a57829003601f168201915b505050505081526020016005820160009054906101000a900460ff16151515158152505090508060c001516103e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103d890610bf5565b60405180910390fd5b806000015181602001518260600151836040015184608001518560c001518361ffff1693508160ff1691509650965096509650965096505091939550919395565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104a990610c87565b60405180910390fd5b6000836040516104c29190610b32565b908152602001604051809103902060050160009054906101000a900460ff1615610521576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161051890610cf3565b60405180910390fd5b60006040518060e001604052808881526020018781526020018581526020018661ffff1681526020018460ff168152602001838152602001600115158152509050806000856040516105739190610b32565b908152602001604051809103902060008201518160000190816105969190610ebf565b5060208201518160010190816105ac9190610ebf565b5060408201518160020190816105c29190610ebf565b5060608201518160030160006101000a81548161ffff021916908361ffff16021790555060808201518160030160026101000a81548160ff021916908360ff16021790555060a082015181600401908161061c9190610ebf565b5060c08201518160050160006101000a81548160ff02191690831515021790555090505050505050505050565b6000808260405161065a9190610b32565b908152602001604051809103902060050160009054906101000a900460ff169050919050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6106e78261069e565b810181811067ffffffffffffffff82111715610706576107056106af565b5b80604052505050565b6000610719610680565b905061072582826106de565b919050565b600067ffffffffffffffff821115610745576107446106af565b5b61074e8261069e565b9050602081019050919050565b82818337600083830152505050565b600061077d6107788461072a565b61070f565b90508281526020810184848401111561079957610798610699565b5b6107a484828561075b565b509392505050565b600082601f8301126107c1576107c0610694565b5b81356107d184826020860161076a565b91505092915050565b6000602082840312156107f0576107ef61068a565b5b600082013567ffffffffffffffff81111561080e5761080d61068f565b5b61081a848285016107ac565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561085d578082015181840152602081019050610842565b60008484015250505050565b600061087482610823565b61087e818561082e565b935061088e81856020860161083f565b6108978161069e565b840191505092915050565b6000819050919050565b6108b5816108a2565b82525050565b600063ffffffff82169050919050565b6108d4816108bb565b82525050565b60008115159050919050565b6108ef816108da565b82525050565b600060c082019050818103600083015261090f8189610869565b905081810360208301526109238188610869565b905061093260408301876108ac565b81810360608301526109448186610869565b905061095360808301856108cb565b61096060a08301846108e6565b979650505050505050565b600061ffff82169050919050565b6109828161096b565b811461098d57600080fd5b50565b60008135905061099f81610979565b92915050565b600060ff82169050919050565b6109bb816109a5565b81146109c657600080fd5b50565b6000813590506109d8816109b2565b92915050565b60008060008060008060c087890312156109fb576109fa61068a565b5b600087013567ffffffffffffffff811115610a1957610a1861068f565b5b610a2589828a016107ac565b965050602087013567ffffffffffffffff811115610a4657610a4561068f565b5b610a5289828a016107ac565b9550506040610a6389828a01610990565b945050606087013567ffffffffffffffff811115610a8457610a8361068f565b5b610a9089828a016107ac565b9350506080610aa189828a016109c9565b92505060a087013567ffffffffffffffff811115610ac257610ac161068f565b5b610ace89828a016107ac565b9150509295509295509295565b6000602082019050610af060008301846108e6565b92915050565b600081905092915050565b6000610b0c82610823565b610b168185610af6565b9350610b2681856020860161083f565b80840191505092915050565b6000610b3e8284610b01565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610b9057607f821691505b602082108103610ba357610ba2610b49565b5b50919050565b7f496a617a616820746964616b2076616c69640000000000000000000000000000600082015250565b6000610bdf60128361082e565b9150610bea82610ba9565b602082019050919050565b60006020820190508181036000830152610c0e81610bd2565b9050919050565b7f48616e79612070656d696c696b206b6f6e7472616b2079616e6720626973612060008201527f6d656e616d6261686b616e000000000000000000000000000000000000000000602082015250565b6000610c71602b8361082e565b9150610c7c82610c15565b604082019050919050565b60006020820190508181036000830152610ca081610c64565b9050919050565b7f496a617a61682073756461682061646100000000000000000000000000000000600082015250565b6000610cdd60108361082e565b9150610ce882610ca7565b602082019050919050565b60006020820190508181036000830152610d0c81610cd0565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610d757fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610d38565b610d7f8683610d38565b95508019841693508086168417925050509392505050565b6000819050919050565b6000610dbc610db7610db2846108a2565b610d97565b6108a2565b9050919050565b6000819050919050565b610dd683610da1565b610dea610de282610dc3565b848454610d45565b825550505050565b600090565b610dff610df2565b610e0a818484610dcd565b505050565b5b81811015610e2e57610e23600082610df7565b600181019050610e10565b5050565b601f821115610e7357610e4481610d13565b610e4d84610d28565b81016020851015610e5c578190505b610e70610e6885610d28565b830182610e0f565b50505b505050565b600082821c905092915050565b6000610e9660001984600802610e78565b1980831691505092915050565b6000610eaf8383610e85565b9150826002028217905092915050565b610ec882610823565b67ffffffffffffffff811115610ee157610ee06106af565b5b610eeb8254610b78565b610ef6828285610e32565b600060209050601f831160018114610f295760008415610f17578287015190505b610f218582610ea3565b865550610f89565b601f198416610f3786610d13565b60005b82811015610f5f57848901518255600182019150602085019450602081019050610f3a565b86831015610f7c5784890151610f78601f891682610e85565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220282006335808f0b9f2752cad0b2562883150af30110c58afcb8a9f828e4fa68464736f6c63430008130033",
}

// VervalIjazahABI is the input ABI used to generate the binding from.
// Deprecated: Use VervalIjazahMetaData.ABI instead.
var VervalIjazahABI = VervalIjazahMetaData.ABI

// VervalIjazahBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VervalIjazahMetaData.Bin instead.
var VervalIjazahBin = VervalIjazahMetaData.Bin

// DeployVervalIjazah deploys a new Ethereum contract, binding an instance of VervalIjazah to it.
func DeployVervalIjazah(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VervalIjazah, error) {
	parsed, err := VervalIjazahMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VervalIjazahBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VervalIjazah{VervalIjazahCaller: VervalIjazahCaller{contract: contract}, VervalIjazahTransactor: VervalIjazahTransactor{contract: contract}, VervalIjazahFilterer: VervalIjazahFilterer{contract: contract}}, nil
}

// VervalIjazah is an auto generated Go binding around an Ethereum contract.
type VervalIjazah struct {
	VervalIjazahCaller     // Read-only binding to the contract
	VervalIjazahTransactor // Write-only binding to the contract
	VervalIjazahFilterer   // Log filterer for contract events
}

// VervalIjazahCaller is an auto generated read-only Go binding around an Ethereum contract.
type VervalIjazahCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VervalIjazahTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VervalIjazahTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VervalIjazahFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VervalIjazahFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VervalIjazahSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VervalIjazahSession struct {
	Contract     *VervalIjazah     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VervalIjazahCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VervalIjazahCallerSession struct {
	Contract *VervalIjazahCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// VervalIjazahTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VervalIjazahTransactorSession struct {
	Contract     *VervalIjazahTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// VervalIjazahRaw is an auto generated low-level Go binding around an Ethereum contract.
type VervalIjazahRaw struct {
	Contract *VervalIjazah // Generic contract binding to access the raw methods on
}

// VervalIjazahCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VervalIjazahCallerRaw struct {
	Contract *VervalIjazahCaller // Generic read-only contract binding to access the raw methods on
}

// VervalIjazahTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VervalIjazahTransactorRaw struct {
	Contract *VervalIjazahTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVervalIjazah creates a new instance of VervalIjazah, bound to a specific deployed contract.
func NewVervalIjazah(address common.Address, backend bind.ContractBackend) (*VervalIjazah, error) {
	contract, err := bindVervalIjazah(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VervalIjazah{VervalIjazahCaller: VervalIjazahCaller{contract: contract}, VervalIjazahTransactor: VervalIjazahTransactor{contract: contract}, VervalIjazahFilterer: VervalIjazahFilterer{contract: contract}}, nil
}

// NewVervalIjazahCaller creates a new read-only instance of VervalIjazah, bound to a specific deployed contract.
func NewVervalIjazahCaller(address common.Address, caller bind.ContractCaller) (*VervalIjazahCaller, error) {
	contract, err := bindVervalIjazah(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VervalIjazahCaller{contract: contract}, nil
}

// NewVervalIjazahTransactor creates a new write-only instance of VervalIjazah, bound to a specific deployed contract.
func NewVervalIjazahTransactor(address common.Address, transactor bind.ContractTransactor) (*VervalIjazahTransactor, error) {
	contract, err := bindVervalIjazah(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VervalIjazahTransactor{contract: contract}, nil
}

// NewVervalIjazahFilterer creates a new log filterer instance of VervalIjazah, bound to a specific deployed contract.
func NewVervalIjazahFilterer(address common.Address, filterer bind.ContractFilterer) (*VervalIjazahFilterer, error) {
	contract, err := bindVervalIjazah(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VervalIjazahFilterer{contract: contract}, nil
}

// bindVervalIjazah binds a generic wrapper to an already deployed contract.
func bindVervalIjazah(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VervalIjazahMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VervalIjazah *VervalIjazahRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VervalIjazah.Contract.VervalIjazahCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VervalIjazah *VervalIjazahRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VervalIjazah.Contract.VervalIjazahTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VervalIjazah *VervalIjazahRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VervalIjazah.Contract.VervalIjazahTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VervalIjazah *VervalIjazahCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VervalIjazah.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VervalIjazah *VervalIjazahTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VervalIjazah.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VervalIjazah *VervalIjazahTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VervalIjazah.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string _nomor) view returns(string, string, uint256, string, uint32, bool)
func (_VervalIjazah *VervalIjazahCaller) Get(opts *bind.CallOpts, _nomor string) (string, string, *big.Int, string, uint32, bool, error) {
	var out []interface{}
	err := _VervalIjazah.contract.Call(opts, &out, "get", _nomor)

	if err != nil {
		return *new(string), *new(string), *new(*big.Int), *new(string), *new(uint32), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(string)).(*string)
	out4 := *abi.ConvertType(out[4], new(uint32)).(*uint32)
	out5 := *abi.ConvertType(out[5], new(bool)).(*bool)

	return out0, out1, out2, out3, out4, out5, err

}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string _nomor) view returns(string, string, uint256, string, uint32, bool)
func (_VervalIjazah *VervalIjazahSession) Get(_nomor string) (string, string, *big.Int, string, uint32, bool, error) {
	return _VervalIjazah.Contract.Get(&_VervalIjazah.CallOpts, _nomor)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string _nomor) view returns(string, string, uint256, string, uint32, bool)
func (_VervalIjazah *VervalIjazahCallerSession) Get(_nomor string) (string, string, *big.Int, string, uint32, bool, error) {
	return _VervalIjazah.Contract.Get(&_VervalIjazah.CallOpts, _nomor)
}

// ValidasiIjazah is a free data retrieval call binding the contract method 0x8f1c4cae.
//
// Solidity: function validasiIjazah(string _nomor) view returns(bool)
func (_VervalIjazah *VervalIjazahCaller) ValidasiIjazah(opts *bind.CallOpts, _nomor string) (bool, error) {
	var out []interface{}
	err := _VervalIjazah.contract.Call(opts, &out, "validasiIjazah", _nomor)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidasiIjazah is a free data retrieval call binding the contract method 0x8f1c4cae.
//
// Solidity: function validasiIjazah(string _nomor) view returns(bool)
func (_VervalIjazah *VervalIjazahSession) ValidasiIjazah(_nomor string) (bool, error) {
	return _VervalIjazah.Contract.ValidasiIjazah(&_VervalIjazah.CallOpts, _nomor)
}

// ValidasiIjazah is a free data retrieval call binding the contract method 0x8f1c4cae.
//
// Solidity: function validasiIjazah(string _nomor) view returns(bool)
func (_VervalIjazah *VervalIjazahCallerSession) ValidasiIjazah(_nomor string) (bool, error) {
	return _VervalIjazah.Contract.ValidasiIjazah(&_VervalIjazah.CallOpts, _nomor)
}

// Add is a paid mutator transaction binding the contract method 0x7d4392ee.
//
// Solidity: function add(string _nisn, string _nama, uint16 _tahunLulus, string _nomor, uint8 _nilaiTranskrip, string _cidURL) returns()
func (_VervalIjazah *VervalIjazahTransactor) Add(opts *bind.TransactOpts, _nisn string, _nama string, _tahunLulus uint16, _nomor string, _nilaiTranskrip uint8, _cidURL string) (*types.Transaction, error) {
	return _VervalIjazah.contract.Transact(opts, "add", _nisn, _nama, _tahunLulus, _nomor, _nilaiTranskrip, _cidURL)
}

// Add is a paid mutator transaction binding the contract method 0x7d4392ee.
//
// Solidity: function add(string _nisn, string _nama, uint16 _tahunLulus, string _nomor, uint8 _nilaiTranskrip, string _cidURL) returns()
func (_VervalIjazah *VervalIjazahSession) Add(_nisn string, _nama string, _tahunLulus uint16, _nomor string, _nilaiTranskrip uint8, _cidURL string) (*types.Transaction, error) {
	return _VervalIjazah.Contract.Add(&_VervalIjazah.TransactOpts, _nisn, _nama, _tahunLulus, _nomor, _nilaiTranskrip, _cidURL)
}

// Add is a paid mutator transaction binding the contract method 0x7d4392ee.
//
// Solidity: function add(string _nisn, string _nama, uint16 _tahunLulus, string _nomor, uint8 _nilaiTranskrip, string _cidURL) returns()
func (_VervalIjazah *VervalIjazahTransactorSession) Add(_nisn string, _nama string, _tahunLulus uint16, _nomor string, _nilaiTranskrip uint8, _cidURL string) (*types.Transaction, error) {
	return _VervalIjazah.Contract.Add(&_VervalIjazah.TransactOpts, _nisn, _nama, _tahunLulus, _nomor, _nilaiTranskrip, _cidURL)
}
