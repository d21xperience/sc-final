// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verval

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

// VervalMetaData contains all meta data concerning the Verval contract.
var VervalMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"degreeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sekolah\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"issueDate\",\"type\":\"uint256\"}],\"name\":\"DegreeIssued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"degreeHash\",\"type\":\"bytes32\"}],\"name\":\"TranscriptAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_degreeHash\",\"type\":\"bytes32\"},{\"internalType\":\"string[]\",\"name\":\"_mataPelajaran\",\"type\":\"string[]\"},{\"internalType\":\"uint8[]\",\"name\":\"_nilai\",\"type\":\"uint8[]\"}],\"name\":\"addTranscript\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"degrees\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"degreeHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"sekolah\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"issueDate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_degreeHash\",\"type\":\"bytes32\"}],\"name\":\"getTranscript\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint8[]\",\"name\":\"\",\"type\":\"uint8[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_degreeHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_sekolah\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_issueDate\",\"type\":\"uint256\"}],\"name\":\"issueDegree\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_degreeHash\",\"type\":\"bytes32\"}],\"name\":\"verifyDegree\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061155b806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80631a7169461461005c578063568ef99114610078578063a2e5a9e2146100a9578063bc50b3cd146100c5578063eba8f9df146100f7575b600080fd5b61007660048036038101906100719190610a5c565b610128565b005b610092600480360381019061008d9190610acb565b61021b565b6040516100a0929190610b86565b60405180910390f35b6100c360048036038101906100be9190610d98565b610342565b005b6100df60048036038101906100da9190610acb565b610477565b6040516100ee93929190610e32565b60405180910390f35b610111600480360381019061010c9190610acb565b610529565b60405161011f92919061103a565b60405180910390f35b6000801b6000808581526020019081526020016000206000015414610182576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610179906110bd565b60405180910390fd5b6040518060600160405280848152602001838152602001828152506000808581526020019081526020016000206000820151816000015560208201518160010190816101ce91906112e9565b5060408201518160020155905050827f8cf374d846fd98968a82fec7083832ce7335a9d89f143c431161be4816cea942838360405161020e929190610b86565b60405180910390a2505050565b6060600080600080858152602001908152602001600020604051806060016040529081600082015481526020016001820180546102579061110c565b80601f01602080910402602001604051908101604052809291908181526020018280546102839061110c565b80156102d05780601f106102a5576101008083540402835291602001916102d0565b820191906000526020600020905b8154815290600101906020018083116102b357829003601f168201915b5050505050815260200160028201548152505090506000801b81600001510361032e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161032590611407565b60405180910390fd5b806020015181604001519250925050915091565b6000801b600080858152602001908152602001600020600001540361039c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161039390611407565b60405180910390fd5b80518251146103e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103d790611499565b60405180910390fd5b604051806040016040528083815260200182815250600160008581526020019081526020016000206000820151816000019080519060200190610424929190610715565b50602082015181600101908051906020019061044192919061076e565b50905050827f8cf28ea2c10881cba28163c6c2fe0009757176b7ff05f1bcaa370c926526172160405160405180910390a2505050565b60006020528060005260406000206000915090508060000154908060010180546104a09061110c565b80601f01602080910402602001604051908101604052809291908181526020018280546104cc9061110c565b80156105195780601f106104ee57610100808354040283529160200191610519565b820191906000526020600020905b8154815290600101906020018083116104fc57829003601f168201915b5050505050908060020154905083565b6060806000600160008581526020019081526020016000206000018054905011610588576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161057f90611505565b60405180910390fd5b60006001600085815260200190815260200160002060405180604001604052908160008201805480602002602001604051908101604052809291908181526020016000905b828210156106795783829060005260206000200180546105ec9061110c565b80601f01602080910402602001604051908101604052809291908181526020018280546106189061110c565b80156106655780601f1061063a57610100808354040283529160200191610665565b820191906000526020600020905b81548152906001019060200180831161064857829003601f168201915b5050505050815260200190600101906105cd565b505050508152602001600182018054806020026020016040519081016040528092919081815260200182805480156106f657602002820191906000526020600020906000905b82829054906101000a900460ff1660ff16815260200190600101906020826000010492830192600103820291508084116106bf5790505b5050505050815250509050806000015181602001519250925050915091565b82805482825590600052602060002090810192821561075d579160200282015b8281111561075c57825182908161074c91906112e9565b5091602001919060010190610735565b5b50905061076a9190610815565b5090565b82805482825590600052602060002090601f016020900481019282156108045791602002820160005b838211156107d557835183826101000a81548160ff021916908360ff1602179055509260200192600101602081600001049283019260010302610797565b80156108025782816101000a81549060ff02191690556001016020816000010492830192600103026107d5565b505b5090506108119190610839565b5090565b5b80821115610835576000818161082c9190610856565b50600101610816565b5090565b5b8082111561085257600081600090555060010161083a565b5090565b5080546108629061110c565b6000825580601f106108745750610893565b601f0160209004906000526020600020908101906108929190610839565b5b50565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b6108bd816108aa565b81146108c857600080fd5b50565b6000813590506108da816108b4565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610933826108ea565b810181811067ffffffffffffffff82111715610952576109516108fb565b5b80604052505050565b6000610965610896565b9050610971828261092a565b919050565b600067ffffffffffffffff821115610991576109906108fb565b5b61099a826108ea565b9050602081019050919050565b82818337600083830152505050565b60006109c96109c484610976565b61095b565b9050828152602081018484840111156109e5576109e46108e5565b5b6109f08482856109a7565b509392505050565b600082601f830112610a0d57610a0c6108e0565b5b8135610a1d8482602086016109b6565b91505092915050565b6000819050919050565b610a3981610a26565b8114610a4457600080fd5b50565b600081359050610a5681610a30565b92915050565b600080600060608486031215610a7557610a746108a0565b5b6000610a83868287016108cb565b935050602084013567ffffffffffffffff811115610aa457610aa36108a5565b5b610ab0868287016109f8565b9250506040610ac186828701610a47565b9150509250925092565b600060208284031215610ae157610ae06108a0565b5b6000610aef848285016108cb565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610b32578082015181840152602081019050610b17565b60008484015250505050565b6000610b4982610af8565b610b538185610b03565b9350610b63818560208601610b14565b610b6c816108ea565b840191505092915050565b610b8081610a26565b82525050565b60006040820190508181036000830152610ba08185610b3e565b9050610baf6020830184610b77565b9392505050565b600067ffffffffffffffff821115610bd157610bd06108fb565b5b602082029050602081019050919050565b600080fd5b6000610bfa610bf584610bb6565b61095b565b90508083825260208201905060208402830185811115610c1d57610c1c610be2565b5b835b81811015610c6457803567ffffffffffffffff811115610c4257610c416108e0565b5b808601610c4f89826109f8565b85526020850194505050602081019050610c1f565b5050509392505050565b600082601f830112610c8357610c826108e0565b5b8135610c93848260208601610be7565b91505092915050565b600067ffffffffffffffff821115610cb757610cb66108fb565b5b602082029050602081019050919050565b600060ff82169050919050565b610cde81610cc8565b8114610ce957600080fd5b50565b600081359050610cfb81610cd5565b92915050565b6000610d14610d0f84610c9c565b61095b565b90508083825260208201905060208402830185811115610d3757610d36610be2565b5b835b81811015610d605780610d4c8882610cec565b845260208401935050602081019050610d39565b5050509392505050565b600082601f830112610d7f57610d7e6108e0565b5b8135610d8f848260208601610d01565b91505092915050565b600080600060608486031215610db157610db06108a0565b5b6000610dbf868287016108cb565b935050602084013567ffffffffffffffff811115610de057610ddf6108a5565b5b610dec86828701610c6e565b925050604084013567ffffffffffffffff811115610e0d57610e0c6108a5565b5b610e1986828701610d6a565b9150509250925092565b610e2c816108aa565b82525050565b6000606082019050610e476000830186610e23565b8181036020830152610e598185610b3e565b9050610e686040830184610b77565b949350505050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600082825260208201905092915050565b6000610eb882610af8565b610ec28185610e9c565b9350610ed2818560208601610b14565b610edb816108ea565b840191505092915050565b6000610ef28383610ead565b905092915050565b6000602082019050919050565b6000610f1282610e70565b610f1c8185610e7b565b935083602082028501610f2e85610e8c565b8060005b85811015610f6a5784840389528151610f4b8582610ee6565b9450610f5683610efa565b925060208a01995050600181019050610f32565b50829750879550505050505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610fb181610cc8565b82525050565b6000610fc38383610fa8565b60208301905092915050565b6000602082019050919050565b6000610fe782610f7c565b610ff18185610f87565b9350610ffc83610f98565b8060005b8381101561102d5781516110148882610fb7565b975061101f83610fcf565b925050600181019050611000565b5085935050505092915050565b600060408201905081810360008301526110548185610f07565b905081810360208301526110688184610fdc565b90509392505050565b7f496a617a6168207375646168207465726461667461722e000000000000000000600082015250565b60006110a7601783610b03565b91506110b282611071565b602082019050919050565b600060208201905081810360008301526110d68161109a565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061112457607f821691505b602082108103611137576111366110dd565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b60006008830261119f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611162565b6111a98683611162565b95508019841693508086168417925050509392505050565b6000819050919050565b60006111e66111e16111dc84610a26565b6111c1565b610a26565b9050919050565b6000819050919050565b611200836111cb565b61121461120c826111ed565b84845461116f565b825550505050565b600090565b61122961121c565b6112348184846111f7565b505050565b5b818110156112585761124d600082611221565b60018101905061123a565b5050565b601f82111561129d5761126e8161113d565b61127784611152565b81016020851015611286578190505b61129a61129285611152565b830182611239565b50505b505050565b600082821c905092915050565b60006112c0600019846008026112a2565b1980831691505092915050565b60006112d983836112af565b9150826002028217905092915050565b6112f282610af8565b67ffffffffffffffff81111561130b5761130a6108fb565b5b611315825461110c565b61132082828561125c565b600060209050601f8311600181146113535760008415611341578287015190505b61134b85826112cd565b8655506113b3565b601f1984166113618661113d565b60005b8281101561138957848901518255600182019150602085019450602081019050611364565b868310156113a657848901516113a2601f8916826112af565b8355505b6001600288020188555050505b505050505050565b7f496a617a616820746964616b20646974656d756b616e2e000000000000000000600082015250565b60006113f1601783610b03565b91506113fc826113bb565b602082019050919050565b60006020820190508181036000830152611420816113e4565b9050919050565b7f4a756d6c6168206d6174612070656c616a6172616e2064616e206e696c61692060008201527f68617275732073616d612e000000000000000000000000000000000000000000602082015250565b6000611483602b83610b03565b915061148e82611427565b604082019050919050565b600060208201905081810360008301526114b281611476565b9050919050565b7f5472616e736b72697020746964616b20646974656d756b616e2e000000000000600082015250565b60006114ef601a83610b03565b91506114fa826114b9565b602082019050919050565b6000602082019050818103600083015261151e816114e2565b905091905056fea2646970667358221220644f6053085e77e931288042f7d87f0039cf3ff69ea48d226272557440a23a9d64736f6c63430008130033",
}

// VervalABI is the input ABI used to generate the binding from.
// Deprecated: Use VervalMetaData.ABI instead.
var VervalABI = VervalMetaData.ABI

// VervalBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VervalMetaData.Bin instead.
var VervalBin = VervalMetaData.Bin

// DeployVerval deploys a new Ethereum contract, binding an instance of Verval to it.
func DeployVerval(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Verval, error) {
	parsed, err := VervalMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VervalBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Verval{VervalCaller: VervalCaller{contract: contract}, VervalTransactor: VervalTransactor{contract: contract}, VervalFilterer: VervalFilterer{contract: contract}}, nil
}

// Verval is an auto generated Go binding around an Ethereum contract.
type Verval struct {
	VervalCaller     // Read-only binding to the contract
	VervalTransactor // Write-only binding to the contract
	VervalFilterer   // Log filterer for contract events
}

// VervalCaller is an auto generated read-only Go binding around an Ethereum contract.
type VervalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VervalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VervalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VervalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VervalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VervalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VervalSession struct {
	Contract     *Verval           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VervalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VervalCallerSession struct {
	Contract *VervalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VervalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VervalTransactorSession struct {
	Contract     *VervalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VervalRaw is an auto generated low-level Go binding around an Ethereum contract.
type VervalRaw struct {
	Contract *Verval // Generic contract binding to access the raw methods on
}

// VervalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VervalCallerRaw struct {
	Contract *VervalCaller // Generic read-only contract binding to access the raw methods on
}

// VervalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VervalTransactorRaw struct {
	Contract *VervalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerval creates a new instance of Verval, bound to a specific deployed contract.
func NewVerval(address common.Address, backend bind.ContractBackend) (*Verval, error) {
	contract, err := bindVerval(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Verval{VervalCaller: VervalCaller{contract: contract}, VervalTransactor: VervalTransactor{contract: contract}, VervalFilterer: VervalFilterer{contract: contract}}, nil
}

// NewVervalCaller creates a new read-only instance of Verval, bound to a specific deployed contract.
func NewVervalCaller(address common.Address, caller bind.ContractCaller) (*VervalCaller, error) {
	contract, err := bindVerval(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VervalCaller{contract: contract}, nil
}

// NewVervalTransactor creates a new write-only instance of Verval, bound to a specific deployed contract.
func NewVervalTransactor(address common.Address, transactor bind.ContractTransactor) (*VervalTransactor, error) {
	contract, err := bindVerval(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VervalTransactor{contract: contract}, nil
}

// NewVervalFilterer creates a new log filterer instance of Verval, bound to a specific deployed contract.
func NewVervalFilterer(address common.Address, filterer bind.ContractFilterer) (*VervalFilterer, error) {
	contract, err := bindVerval(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VervalFilterer{contract: contract}, nil
}

// bindVerval binds a generic wrapper to an already deployed contract.
func bindVerval(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VervalMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verval *VervalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verval.Contract.VervalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verval *VervalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verval.Contract.VervalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verval *VervalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verval.Contract.VervalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verval *VervalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verval.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verval *VervalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verval.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verval *VervalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verval.Contract.contract.Transact(opts, method, params...)
}

// Degrees is a free data retrieval call binding the contract method 0xbc50b3cd.
//
// Solidity: function degrees(bytes32 ) view returns(bytes32 degreeHash, string sekolah, uint256 issueDate)
func (_Verval *VervalCaller) Degrees(opts *bind.CallOpts, arg0 [32]byte) (struct {
	DegreeHash [32]byte
	Sekolah    string
	IssueDate  *big.Int
}, error) {
	var out []interface{}
	err := _Verval.contract.Call(opts, &out, "degrees", arg0)

	outstruct := new(struct {
		DegreeHash [32]byte
		Sekolah    string
		IssueDate  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DegreeHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Sekolah = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.IssueDate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Degrees is a free data retrieval call binding the contract method 0xbc50b3cd.
//
// Solidity: function degrees(bytes32 ) view returns(bytes32 degreeHash, string sekolah, uint256 issueDate)
func (_Verval *VervalSession) Degrees(arg0 [32]byte) (struct {
	DegreeHash [32]byte
	Sekolah    string
	IssueDate  *big.Int
}, error) {
	return _Verval.Contract.Degrees(&_Verval.CallOpts, arg0)
}

// Degrees is a free data retrieval call binding the contract method 0xbc50b3cd.
//
// Solidity: function degrees(bytes32 ) view returns(bytes32 degreeHash, string sekolah, uint256 issueDate)
func (_Verval *VervalCallerSession) Degrees(arg0 [32]byte) (struct {
	DegreeHash [32]byte
	Sekolah    string
	IssueDate  *big.Int
}, error) {
	return _Verval.Contract.Degrees(&_Verval.CallOpts, arg0)
}

// GetTranscript is a free data retrieval call binding the contract method 0xeba8f9df.
//
// Solidity: function getTranscript(bytes32 _degreeHash) view returns(string[], uint8[])
func (_Verval *VervalCaller) GetTranscript(opts *bind.CallOpts, _degreeHash [32]byte) ([]string, []uint8, error) {
	var out []interface{}
	err := _Verval.contract.Call(opts, &out, "getTranscript", _degreeHash)

	if err != nil {
		return *new([]string), *new([]uint8), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)
	out1 := *abi.ConvertType(out[1], new([]uint8)).(*[]uint8)

	return out0, out1, err

}

// GetTranscript is a free data retrieval call binding the contract method 0xeba8f9df.
//
// Solidity: function getTranscript(bytes32 _degreeHash) view returns(string[], uint8[])
func (_Verval *VervalSession) GetTranscript(_degreeHash [32]byte) ([]string, []uint8, error) {
	return _Verval.Contract.GetTranscript(&_Verval.CallOpts, _degreeHash)
}

// GetTranscript is a free data retrieval call binding the contract method 0xeba8f9df.
//
// Solidity: function getTranscript(bytes32 _degreeHash) view returns(string[], uint8[])
func (_Verval *VervalCallerSession) GetTranscript(_degreeHash [32]byte) ([]string, []uint8, error) {
	return _Verval.Contract.GetTranscript(&_Verval.CallOpts, _degreeHash)
}

// VerifyDegree is a free data retrieval call binding the contract method 0x568ef991.
//
// Solidity: function verifyDegree(bytes32 _degreeHash) view returns(string, uint256)
func (_Verval *VervalCaller) VerifyDegree(opts *bind.CallOpts, _degreeHash [32]byte) (string, *big.Int, error) {
	var out []interface{}
	err := _Verval.contract.Call(opts, &out, "verifyDegree", _degreeHash)

	if err != nil {
		return *new(string), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// VerifyDegree is a free data retrieval call binding the contract method 0x568ef991.
//
// Solidity: function verifyDegree(bytes32 _degreeHash) view returns(string, uint256)
func (_Verval *VervalSession) VerifyDegree(_degreeHash [32]byte) (string, *big.Int, error) {
	return _Verval.Contract.VerifyDegree(&_Verval.CallOpts, _degreeHash)
}

// VerifyDegree is a free data retrieval call binding the contract method 0x568ef991.
//
// Solidity: function verifyDegree(bytes32 _degreeHash) view returns(string, uint256)
func (_Verval *VervalCallerSession) VerifyDegree(_degreeHash [32]byte) (string, *big.Int, error) {
	return _Verval.Contract.VerifyDegree(&_Verval.CallOpts, _degreeHash)
}

// AddTranscript is a paid mutator transaction binding the contract method 0xa2e5a9e2.
//
// Solidity: function addTranscript(bytes32 _degreeHash, string[] _mataPelajaran, uint8[] _nilai) returns()
func (_Verval *VervalTransactor) AddTranscript(opts *bind.TransactOpts, _degreeHash [32]byte, _mataPelajaran []string, _nilai []uint8) (*types.Transaction, error) {
	return _Verval.contract.Transact(opts, "addTranscript", _degreeHash, _mataPelajaran, _nilai)
}

// AddTranscript is a paid mutator transaction binding the contract method 0xa2e5a9e2.
//
// Solidity: function addTranscript(bytes32 _degreeHash, string[] _mataPelajaran, uint8[] _nilai) returns()
func (_Verval *VervalSession) AddTranscript(_degreeHash [32]byte, _mataPelajaran []string, _nilai []uint8) (*types.Transaction, error) {
	return _Verval.Contract.AddTranscript(&_Verval.TransactOpts, _degreeHash, _mataPelajaran, _nilai)
}

// AddTranscript is a paid mutator transaction binding the contract method 0xa2e5a9e2.
//
// Solidity: function addTranscript(bytes32 _degreeHash, string[] _mataPelajaran, uint8[] _nilai) returns()
func (_Verval *VervalTransactorSession) AddTranscript(_degreeHash [32]byte, _mataPelajaran []string, _nilai []uint8) (*types.Transaction, error) {
	return _Verval.Contract.AddTranscript(&_Verval.TransactOpts, _degreeHash, _mataPelajaran, _nilai)
}

// IssueDegree is a paid mutator transaction binding the contract method 0x1a716946.
//
// Solidity: function issueDegree(bytes32 _degreeHash, string _sekolah, uint256 _issueDate) returns()
func (_Verval *VervalTransactor) IssueDegree(opts *bind.TransactOpts, _degreeHash [32]byte, _sekolah string, _issueDate *big.Int) (*types.Transaction, error) {
	return _Verval.contract.Transact(opts, "issueDegree", _degreeHash, _sekolah, _issueDate)
}

// IssueDegree is a paid mutator transaction binding the contract method 0x1a716946.
//
// Solidity: function issueDegree(bytes32 _degreeHash, string _sekolah, uint256 _issueDate) returns()
func (_Verval *VervalSession) IssueDegree(_degreeHash [32]byte, _sekolah string, _issueDate *big.Int) (*types.Transaction, error) {
	return _Verval.Contract.IssueDegree(&_Verval.TransactOpts, _degreeHash, _sekolah, _issueDate)
}

// IssueDegree is a paid mutator transaction binding the contract method 0x1a716946.
//
// Solidity: function issueDegree(bytes32 _degreeHash, string _sekolah, uint256 _issueDate) returns()
func (_Verval *VervalTransactorSession) IssueDegree(_degreeHash [32]byte, _sekolah string, _issueDate *big.Int) (*types.Transaction, error) {
	return _Verval.Contract.IssueDegree(&_Verval.TransactOpts, _degreeHash, _sekolah, _issueDate)
}

// VervalDegreeIssuedIterator is returned from FilterDegreeIssued and is used to iterate over the raw logs and unpacked data for DegreeIssued events raised by the Verval contract.
type VervalDegreeIssuedIterator struct {
	Event *VervalDegreeIssued // Event containing the contract specifics and raw log

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
func (it *VervalDegreeIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VervalDegreeIssued)
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
		it.Event = new(VervalDegreeIssued)
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
func (it *VervalDegreeIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VervalDegreeIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VervalDegreeIssued represents a DegreeIssued event raised by the Verval contract.
type VervalDegreeIssued struct {
	DegreeHash [32]byte
	Sekolah    string
	IssueDate  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDegreeIssued is a free log retrieval operation binding the contract event 0x8cf374d846fd98968a82fec7083832ce7335a9d89f143c431161be4816cea942.
//
// Solidity: event DegreeIssued(bytes32 indexed degreeHash, string sekolah, uint256 issueDate)
func (_Verval *VervalFilterer) FilterDegreeIssued(opts *bind.FilterOpts, degreeHash [][32]byte) (*VervalDegreeIssuedIterator, error) {

	var degreeHashRule []interface{}
	for _, degreeHashItem := range degreeHash {
		degreeHashRule = append(degreeHashRule, degreeHashItem)
	}

	logs, sub, err := _Verval.contract.FilterLogs(opts, "DegreeIssued", degreeHashRule)
	if err != nil {
		return nil, err
	}
	return &VervalDegreeIssuedIterator{contract: _Verval.contract, event: "DegreeIssued", logs: logs, sub: sub}, nil
}

// WatchDegreeIssued is a free log subscription operation binding the contract event 0x8cf374d846fd98968a82fec7083832ce7335a9d89f143c431161be4816cea942.
//
// Solidity: event DegreeIssued(bytes32 indexed degreeHash, string sekolah, uint256 issueDate)
func (_Verval *VervalFilterer) WatchDegreeIssued(opts *bind.WatchOpts, sink chan<- *VervalDegreeIssued, degreeHash [][32]byte) (event.Subscription, error) {

	var degreeHashRule []interface{}
	for _, degreeHashItem := range degreeHash {
		degreeHashRule = append(degreeHashRule, degreeHashItem)
	}

	logs, sub, err := _Verval.contract.WatchLogs(opts, "DegreeIssued", degreeHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VervalDegreeIssued)
				if err := _Verval.contract.UnpackLog(event, "DegreeIssued", log); err != nil {
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

// ParseDegreeIssued is a log parse operation binding the contract event 0x8cf374d846fd98968a82fec7083832ce7335a9d89f143c431161be4816cea942.
//
// Solidity: event DegreeIssued(bytes32 indexed degreeHash, string sekolah, uint256 issueDate)
func (_Verval *VervalFilterer) ParseDegreeIssued(log types.Log) (*VervalDegreeIssued, error) {
	event := new(VervalDegreeIssued)
	if err := _Verval.contract.UnpackLog(event, "DegreeIssued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VervalTranscriptAddedIterator is returned from FilterTranscriptAdded and is used to iterate over the raw logs and unpacked data for TranscriptAdded events raised by the Verval contract.
type VervalTranscriptAddedIterator struct {
	Event *VervalTranscriptAdded // Event containing the contract specifics and raw log

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
func (it *VervalTranscriptAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VervalTranscriptAdded)
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
		it.Event = new(VervalTranscriptAdded)
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
func (it *VervalTranscriptAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VervalTranscriptAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VervalTranscriptAdded represents a TranscriptAdded event raised by the Verval contract.
type VervalTranscriptAdded struct {
	DegreeHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTranscriptAdded is a free log retrieval operation binding the contract event 0x8cf28ea2c10881cba28163c6c2fe0009757176b7ff05f1bcaa370c9265261721.
//
// Solidity: event TranscriptAdded(bytes32 indexed degreeHash)
func (_Verval *VervalFilterer) FilterTranscriptAdded(opts *bind.FilterOpts, degreeHash [][32]byte) (*VervalTranscriptAddedIterator, error) {

	var degreeHashRule []interface{}
	for _, degreeHashItem := range degreeHash {
		degreeHashRule = append(degreeHashRule, degreeHashItem)
	}

	logs, sub, err := _Verval.contract.FilterLogs(opts, "TranscriptAdded", degreeHashRule)
	if err != nil {
		return nil, err
	}
	return &VervalTranscriptAddedIterator{contract: _Verval.contract, event: "TranscriptAdded", logs: logs, sub: sub}, nil
}

// WatchTranscriptAdded is a free log subscription operation binding the contract event 0x8cf28ea2c10881cba28163c6c2fe0009757176b7ff05f1bcaa370c9265261721.
//
// Solidity: event TranscriptAdded(bytes32 indexed degreeHash)
func (_Verval *VervalFilterer) WatchTranscriptAdded(opts *bind.WatchOpts, sink chan<- *VervalTranscriptAdded, degreeHash [][32]byte) (event.Subscription, error) {

	var degreeHashRule []interface{}
	for _, degreeHashItem := range degreeHash {
		degreeHashRule = append(degreeHashRule, degreeHashItem)
	}

	logs, sub, err := _Verval.contract.WatchLogs(opts, "TranscriptAdded", degreeHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VervalTranscriptAdded)
				if err := _Verval.contract.UnpackLog(event, "TranscriptAdded", log); err != nil {
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

// ParseTranscriptAdded is a log parse operation binding the contract event 0x8cf28ea2c10881cba28163c6c2fe0009757176b7ff05f1bcaa370c9265261721.
//
// Solidity: event TranscriptAdded(bytes32 indexed degreeHash)
func (_Verval *VervalFilterer) ParseTranscriptAdded(log types.Log) (*VervalTranscriptAdded, error) {
	event := new(VervalTranscriptAdded)
	if err := _Verval.contract.UnpackLog(event, "TranscriptAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
