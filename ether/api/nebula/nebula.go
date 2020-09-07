// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nebula

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// GravityABI is the input ABI used to generate the binding from.
const GravityABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newConsuls\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"newBftValue\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"bftValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"consuls\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getConsuls\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newConsuls\",\"type\":\"address[]\"}],\"name\":\"hashNewConsuls\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rounds\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newConsuls\",\"type\":\"address[]\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"newLastRound\",\"type\":\"uint256\"}],\"name\":\"updateConsuls\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GravityFuncSigs maps the 4-byte function signature to its string representation.
var GravityFuncSigs = map[string]string{
	"3cec1bdd": "bftValue()",
	"a2c2c617": "consuls(uint256)",
	"ad595b1a": "getConsuls()",
	"4dea5eba": "hashNewConsuls(address[])",
	"8c65c81f": "rounds(uint256)",
	"92c388ab": "updateConsuls(address[],uint8[],bytes32[],bytes32[],uint256)",
}

// GravityBin is the compiled bytecode used for deploying new contracts.
var GravityBin = "0x608060405234801561001057600080fd5b5060405161094c38038061094c8339818101604052604081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b90830190602082018581111561006857600080fd5b825186602082028301116401000000008211171561008557600080fd5b82525081516020918201928201910280838360005b838110156100b257818101518382015260200161009a565b50505050919091016040525060209081015184519093506100d992506001918501906100e3565b506002555061016f565b828054828255906000526020600020908101928215610138579160200282015b8281111561013857825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190610103565b50610144929150610148565b5090565b61016c91905b808211156101445780546001600160a01b031916815560010161014e565b90565b6107ce8061017e6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80633cec1bdd146100675780634dea5eba146100815780638c65c81f1461012257806392c388ab14610153578063a2c2c6171461037e578063ad595b1a146103b7575b600080fd5b61006f61040f565b60408051918252519081900360200190f35b61006f6004803603602081101561009757600080fd5b810190602081018135600160201b8111156100b157600080fd5b8201836020820111156100c357600080fd5b803590602001918460208302840111600160201b831117156100e457600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610415945050505050565b61013f6004803603602081101561013857600080fd5b50356104d9565b604080519115158252519081900360200190f35b61037c600480360360a081101561016957600080fd5b810190602081018135600160201b81111561018357600080fd5b82018360208201111561019557600080fd5b803590602001918460208302840111600160201b831117156101b657600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561020557600080fd5b82018360208201111561021757600080fd5b803590602001918460208302840111600160201b8311171561023857600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561028757600080fd5b82018360208201111561029957600080fd5b803590602001918460208302840111600160201b831117156102ba57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561030957600080fd5b82018360208201111561031b57600080fd5b803590602001918460208302840111600160201b8311171561033c57600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955050913592506104ee915050565b005b61039b6004803603602081101561039457600080fd5b5035610686565b604080516001600160a01b039092168252519081900360200190f35b6103bf6106ad565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156103fb5781810151838201526020016103e3565b505050509050019250505060405180910390f35b60025481565b60006060815b83518110156104ca578184828151811061043157fe5b60200260200101516040516020018083805190602001908083835b6020831061046b5780518252601f19909201916020918201910161044c565b6001836020036101000a038019825116818451168082178552505050505050905001826001600160a01b03166001600160a01b031660601b8152601401925050506040516020818303038152906040529150808060010191505061041b565b50805160209091012092915050565b60006020819052908152604090205460ff1681565b6000806104fa87610415565b905060005b600154811015610605576001818154811061051657fe5b9060005260206000200160009054906101000a90046001600160a01b03166001600160a01b031660018389848151811061054c57fe5b602002602001015189858151811061056057fe5b602002602001015189868151811061057457fe5b602002602001015160405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa1580156105d3573d6000803e3d6000fd5b505050602060405103516001600160a01b0316146105f25760006105f5565b60015b60ff1692909201916001016104ff565b50600254821015610651576040805162461bcd60e51b81526020600482015260116024820152701a5b9d985b1a590818999d0818dbdd5b9d607a1b604482015290519081900360640190fd5b86516106649060019060208a0190610710565b5050506000908152602081905260409020805460ff1916600117905550505050565b6001818154811061069357fe5b6000918252602090912001546001600160a01b0316905081565b6060600180548060200260200160405190810160405280929190818152602001828054801561070557602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116106e7575b505050505090505b90565b828054828255906000526020600020908101928215610765579160200282015b8281111561076557825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190610730565b50610771929150610775565b5090565b61070d91905b808211156107715780546001600160a01b031916815560010161077b56fea265627a7a72315820041ea00d3726c24fac67367605659993a0b9866676fb24f55410e779180a621564736f6c63430005100032"

// DeployGravity deploys a new Ethereum contract, binding an instance of Gravity to it.
func DeployGravity(auth *bind.TransactOpts, backend bind.ContractBackend, newConsuls []common.Address, newBftValue *big.Int) (common.Address, *types.Transaction, *Gravity, error) {
	parsed, err := abi.JSON(strings.NewReader(GravityABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GravityBin), backend, newConsuls, newBftValue)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Gravity{GravityCaller: GravityCaller{contract: contract}, GravityTransactor: GravityTransactor{contract: contract}, GravityFilterer: GravityFilterer{contract: contract}}, nil
}

// Gravity is an auto generated Go binding around an Ethereum contract.
type Gravity struct {
	GravityCaller     // Read-only binding to the contract
	GravityTransactor // Write-only binding to the contract
	GravityFilterer   // Log filterer for contract events
}

// GravityCaller is an auto generated read-only Go binding around an Ethereum contract.
type GravityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GravityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GravityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GravityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GravityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GravitySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GravitySession struct {
	Contract     *Gravity          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GravityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GravityCallerSession struct {
	Contract *GravityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// GravityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GravityTransactorSession struct {
	Contract     *GravityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GravityRaw is an auto generated low-level Go binding around an Ethereum contract.
type GravityRaw struct {
	Contract *Gravity // Generic contract binding to access the raw methods on
}

// GravityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GravityCallerRaw struct {
	Contract *GravityCaller // Generic read-only contract binding to access the raw methods on
}

// GravityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GravityTransactorRaw struct {
	Contract *GravityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGravity creates a new instance of Gravity, bound to a specific deployed contract.
func NewGravity(address common.Address, backend bind.ContractBackend) (*Gravity, error) {
	contract, err := bindGravity(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gravity{GravityCaller: GravityCaller{contract: contract}, GravityTransactor: GravityTransactor{contract: contract}, GravityFilterer: GravityFilterer{contract: contract}}, nil
}

// NewGravityCaller creates a new read-only instance of Gravity, bound to a specific deployed contract.
func NewGravityCaller(address common.Address, caller bind.ContractCaller) (*GravityCaller, error) {
	contract, err := bindGravity(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GravityCaller{contract: contract}, nil
}

// NewGravityTransactor creates a new write-only instance of Gravity, bound to a specific deployed contract.
func NewGravityTransactor(address common.Address, transactor bind.ContractTransactor) (*GravityTransactor, error) {
	contract, err := bindGravity(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GravityTransactor{contract: contract}, nil
}

// NewGravityFilterer creates a new log filterer instance of Gravity, bound to a specific deployed contract.
func NewGravityFilterer(address common.Address, filterer bind.ContractFilterer) (*GravityFilterer, error) {
	contract, err := bindGravity(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GravityFilterer{contract: contract}, nil
}

// bindGravity binds a generic wrapper to an already deployed contract.
func bindGravity(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GravityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gravity *GravityRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Gravity.Contract.GravityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gravity *GravityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gravity.Contract.GravityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gravity *GravityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gravity.Contract.GravityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gravity *GravityCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Gravity.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gravity *GravityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gravity.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gravity *GravityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gravity.Contract.contract.Transact(opts, method, params...)
}

// BftValue is a free data retrieval call binding the contract method 0x3cec1bdd.
//
// Solidity: function bftValue() view returns(uint256)
func (_Gravity *GravityCaller) BftValue(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Gravity.contract.Call(opts, out, "bftValue")
	return *ret0, err
}

// BftValue is a free data retrieval call binding the contract method 0x3cec1bdd.
//
// Solidity: function bftValue() view returns(uint256)
func (_Gravity *GravitySession) BftValue() (*big.Int, error) {
	return _Gravity.Contract.BftValue(&_Gravity.CallOpts)
}

// BftValue is a free data retrieval call binding the contract method 0x3cec1bdd.
//
// Solidity: function bftValue() view returns(uint256)
func (_Gravity *GravityCallerSession) BftValue() (*big.Int, error) {
	return _Gravity.Contract.BftValue(&_Gravity.CallOpts)
}

// Consuls is a free data retrieval call binding the contract method 0xa2c2c617.
//
// Solidity: function consuls(uint256 ) view returns(address)
func (_Gravity *GravityCaller) Consuls(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Gravity.contract.Call(opts, out, "consuls", arg0)
	return *ret0, err
}

// Consuls is a free data retrieval call binding the contract method 0xa2c2c617.
//
// Solidity: function consuls(uint256 ) view returns(address)
func (_Gravity *GravitySession) Consuls(arg0 *big.Int) (common.Address, error) {
	return _Gravity.Contract.Consuls(&_Gravity.CallOpts, arg0)
}

// Consuls is a free data retrieval call binding the contract method 0xa2c2c617.
//
// Solidity: function consuls(uint256 ) view returns(address)
func (_Gravity *GravityCallerSession) Consuls(arg0 *big.Int) (common.Address, error) {
	return _Gravity.Contract.Consuls(&_Gravity.CallOpts, arg0)
}

// GetConsuls is a free data retrieval call binding the contract method 0xad595b1a.
//
// Solidity: function getConsuls() view returns(address[])
func (_Gravity *GravityCaller) GetConsuls(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Gravity.contract.Call(opts, out, "getConsuls")
	return *ret0, err
}

// GetConsuls is a free data retrieval call binding the contract method 0xad595b1a.
//
// Solidity: function getConsuls() view returns(address[])
func (_Gravity *GravitySession) GetConsuls() ([]common.Address, error) {
	return _Gravity.Contract.GetConsuls(&_Gravity.CallOpts)
}

// GetConsuls is a free data retrieval call binding the contract method 0xad595b1a.
//
// Solidity: function getConsuls() view returns(address[])
func (_Gravity *GravityCallerSession) GetConsuls() ([]common.Address, error) {
	return _Gravity.Contract.GetConsuls(&_Gravity.CallOpts)
}

// HashNewConsuls is a free data retrieval call binding the contract method 0x4dea5eba.
//
// Solidity: function hashNewConsuls(address[] newConsuls) pure returns(bytes32)
func (_Gravity *GravityCaller) HashNewConsuls(opts *bind.CallOpts, newConsuls []common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Gravity.contract.Call(opts, out, "hashNewConsuls", newConsuls)
	return *ret0, err
}

// HashNewConsuls is a free data retrieval call binding the contract method 0x4dea5eba.
//
// Solidity: function hashNewConsuls(address[] newConsuls) pure returns(bytes32)
func (_Gravity *GravitySession) HashNewConsuls(newConsuls []common.Address) ([32]byte, error) {
	return _Gravity.Contract.HashNewConsuls(&_Gravity.CallOpts, newConsuls)
}

// HashNewConsuls is a free data retrieval call binding the contract method 0x4dea5eba.
//
// Solidity: function hashNewConsuls(address[] newConsuls) pure returns(bytes32)
func (_Gravity *GravityCallerSession) HashNewConsuls(newConsuls []common.Address) ([32]byte, error) {
	return _Gravity.Contract.HashNewConsuls(&_Gravity.CallOpts, newConsuls)
}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(bool)
func (_Gravity *GravityCaller) Rounds(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Gravity.contract.Call(opts, out, "rounds", arg0)
	return *ret0, err
}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(bool)
func (_Gravity *GravitySession) Rounds(arg0 *big.Int) (bool, error) {
	return _Gravity.Contract.Rounds(&_Gravity.CallOpts, arg0)
}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(bool)
func (_Gravity *GravityCallerSession) Rounds(arg0 *big.Int) (bool, error) {
	return _Gravity.Contract.Rounds(&_Gravity.CallOpts, arg0)
}

// UpdateConsuls is a paid mutator transaction binding the contract method 0x92c388ab.
//
// Solidity: function updateConsuls(address[] newConsuls, uint8[] v, bytes32[] r, bytes32[] s, uint256 newLastRound) returns()
func (_Gravity *GravityTransactor) UpdateConsuls(opts *bind.TransactOpts, newConsuls []common.Address, v []uint8, r [][32]byte, s [][32]byte, newLastRound *big.Int) (*types.Transaction, error) {
	return _Gravity.contract.Transact(opts, "updateConsuls", newConsuls, v, r, s, newLastRound)
}

// UpdateConsuls is a paid mutator transaction binding the contract method 0x92c388ab.
//
// Solidity: function updateConsuls(address[] newConsuls, uint8[] v, bytes32[] r, bytes32[] s, uint256 newLastRound) returns()
func (_Gravity *GravitySession) UpdateConsuls(newConsuls []common.Address, v []uint8, r [][32]byte, s [][32]byte, newLastRound *big.Int) (*types.Transaction, error) {
	return _Gravity.Contract.UpdateConsuls(&_Gravity.TransactOpts, newConsuls, v, r, s, newLastRound)
}

// UpdateConsuls is a paid mutator transaction binding the contract method 0x92c388ab.
//
// Solidity: function updateConsuls(address[] newConsuls, uint8[] v, bytes32[] r, bytes32[] s, uint256 newLastRound) returns()
func (_Gravity *GravityTransactorSession) UpdateConsuls(newConsuls []common.Address, v []uint8, r [][32]byte, s [][32]byte, newLastRound *big.Int) (*types.Transaction, error) {
	return _Gravity.Contract.UpdateConsuls(&_Gravity.TransactOpts, newConsuls, v, r, s, newLastRound)
}

// ISubscriptionABI is the input ABI used to generate the binding from.
const ISubscriptionABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"attachData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ISubscriptionFuncSigs maps the 4-byte function signature to its string representation.
var ISubscriptionFuncSigs = map[string]string{
	"37ef679a": "attachData(bytes)",
}

// ISubscription is an auto generated Go binding around an Ethereum contract.
type ISubscription struct {
	ISubscriptionCaller     // Read-only binding to the contract
	ISubscriptionTransactor // Write-only binding to the contract
	ISubscriptionFilterer   // Log filterer for contract events
}

// ISubscriptionCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISubscriptionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISubscriptionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISubscriptionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISubscriptionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISubscriptionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISubscriptionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISubscriptionSession struct {
	Contract     *ISubscription    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISubscriptionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISubscriptionCallerSession struct {
	Contract *ISubscriptionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ISubscriptionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISubscriptionTransactorSession struct {
	Contract     *ISubscriptionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ISubscriptionRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISubscriptionRaw struct {
	Contract *ISubscription // Generic contract binding to access the raw methods on
}

// ISubscriptionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISubscriptionCallerRaw struct {
	Contract *ISubscriptionCaller // Generic read-only contract binding to access the raw methods on
}

// ISubscriptionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISubscriptionTransactorRaw struct {
	Contract *ISubscriptionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISubscription creates a new instance of ISubscription, bound to a specific deployed contract.
func NewISubscription(address common.Address, backend bind.ContractBackend) (*ISubscription, error) {
	contract, err := bindISubscription(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISubscription{ISubscriptionCaller: ISubscriptionCaller{contract: contract}, ISubscriptionTransactor: ISubscriptionTransactor{contract: contract}, ISubscriptionFilterer: ISubscriptionFilterer{contract: contract}}, nil
}

// NewISubscriptionCaller creates a new read-only instance of ISubscription, bound to a specific deployed contract.
func NewISubscriptionCaller(address common.Address, caller bind.ContractCaller) (*ISubscriptionCaller, error) {
	contract, err := bindISubscription(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISubscriptionCaller{contract: contract}, nil
}

// NewISubscriptionTransactor creates a new write-only instance of ISubscription, bound to a specific deployed contract.
func NewISubscriptionTransactor(address common.Address, transactor bind.ContractTransactor) (*ISubscriptionTransactor, error) {
	contract, err := bindISubscription(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISubscriptionTransactor{contract: contract}, nil
}

// NewISubscriptionFilterer creates a new log filterer instance of ISubscription, bound to a specific deployed contract.
func NewISubscriptionFilterer(address common.Address, filterer bind.ContractFilterer) (*ISubscriptionFilterer, error) {
	contract, err := bindISubscription(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISubscriptionFilterer{contract: contract}, nil
}

// bindISubscription binds a generic wrapper to an already deployed contract.
func bindISubscription(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISubscriptionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISubscription *ISubscriptionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ISubscription.Contract.ISubscriptionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISubscription *ISubscriptionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISubscription.Contract.ISubscriptionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISubscription *ISubscriptionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISubscription.Contract.ISubscriptionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISubscription *ISubscriptionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ISubscription.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISubscription *ISubscriptionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISubscription.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISubscription *ISubscriptionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISubscription.Contract.contract.Transact(opts, method, params...)
}

// AttachData is a paid mutator transaction binding the contract method 0x37ef679a.
//
// Solidity: function attachData(bytes data) returns()
func (_ISubscription *ISubscriptionTransactor) AttachData(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _ISubscription.contract.Transact(opts, "attachData", data)
}

// AttachData is a paid mutator transaction binding the contract method 0x37ef679a.
//
// Solidity: function attachData(bytes data) returns()
func (_ISubscription *ISubscriptionSession) AttachData(data []byte) (*types.Transaction, error) {
	return _ISubscription.Contract.AttachData(&_ISubscription.TransactOpts, data)
}

// AttachData is a paid mutator transaction binding the contract method 0x37ef679a.
//
// Solidity: function attachData(bytes data) returns()
func (_ISubscription *ISubscriptionTransactorSession) AttachData(data []byte) (*types.Transaction, error) {
	return _ISubscription.Contract.AttachData(&_ISubscription.TransactOpts, data)
}

// ModelsABI is the input ABI used to generate the binding from.
const ModelsABI = "[]"

// ModelsBin is the compiled bytecode used for deploying new contracts.
var ModelsBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158200135446abaa192896351489925ceb3bf0590b2ec496a4c9ab94d53d9aa8bb10764736f6c63430005100032"

// DeployModels deploys a new Ethereum contract, binding an instance of Models to it.
func DeployModels(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Models, error) {
	parsed, err := abi.JSON(strings.NewReader(ModelsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ModelsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Models{ModelsCaller: ModelsCaller{contract: contract}, ModelsTransactor: ModelsTransactor{contract: contract}, ModelsFilterer: ModelsFilterer{contract: contract}}, nil
}

// Models is an auto generated Go binding around an Ethereum contract.
type Models struct {
	ModelsCaller     // Read-only binding to the contract
	ModelsTransactor // Write-only binding to the contract
	ModelsFilterer   // Log filterer for contract events
}

// ModelsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ModelsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ModelsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ModelsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ModelsSession struct {
	Contract     *Models           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModelsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ModelsCallerSession struct {
	Contract *ModelsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ModelsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ModelsTransactorSession struct {
	Contract     *ModelsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModelsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ModelsRaw struct {
	Contract *Models // Generic contract binding to access the raw methods on
}

// ModelsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ModelsCallerRaw struct {
	Contract *ModelsCaller // Generic read-only contract binding to access the raw methods on
}

// ModelsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ModelsTransactorRaw struct {
	Contract *ModelsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewModels creates a new instance of Models, bound to a specific deployed contract.
func NewModels(address common.Address, backend bind.ContractBackend) (*Models, error) {
	contract, err := bindModels(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Models{ModelsCaller: ModelsCaller{contract: contract}, ModelsTransactor: ModelsTransactor{contract: contract}, ModelsFilterer: ModelsFilterer{contract: contract}}, nil
}

// NewModelsCaller creates a new read-only instance of Models, bound to a specific deployed contract.
func NewModelsCaller(address common.Address, caller bind.ContractCaller) (*ModelsCaller, error) {
	contract, err := bindModels(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ModelsCaller{contract: contract}, nil
}

// NewModelsTransactor creates a new write-only instance of Models, bound to a specific deployed contract.
func NewModelsTransactor(address common.Address, transactor bind.ContractTransactor) (*ModelsTransactor, error) {
	contract, err := bindModels(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ModelsTransactor{contract: contract}, nil
}

// NewModelsFilterer creates a new log filterer instance of Models, bound to a specific deployed contract.
func NewModelsFilterer(address common.Address, filterer bind.ContractFilterer) (*ModelsFilterer, error) {
	contract, err := bindModels(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ModelsFilterer{contract: contract}, nil
}

// bindModels binds a generic wrapper to an already deployed contract.
func bindModels(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ModelsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Models *ModelsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Models.Contract.ModelsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Models *ModelsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Models.Contract.ModelsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Models *ModelsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Models.Contract.ModelsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Models *ModelsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Models.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Models *ModelsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Models.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Models *ModelsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Models.Contract.contract.Transact(opts, method, params...)
}

// NModelsABI is the input ABI used to generate the binding from.
const NModelsABI = "[]"

// NModelsBin is the compiled bytecode used for deploying new contracts.
var NModelsBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582003ae8e9bedeff21cbcc2594a833309f9b105329852faa2af7650a4db11cb659c64736f6c63430005100032"

// DeployNModels deploys a new Ethereum contract, binding an instance of NModels to it.
func DeployNModels(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NModels, error) {
	parsed, err := abi.JSON(strings.NewReader(NModelsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NModelsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NModels{NModelsCaller: NModelsCaller{contract: contract}, NModelsTransactor: NModelsTransactor{contract: contract}, NModelsFilterer: NModelsFilterer{contract: contract}}, nil
}

// NModels is an auto generated Go binding around an Ethereum contract.
type NModels struct {
	NModelsCaller     // Read-only binding to the contract
	NModelsTransactor // Write-only binding to the contract
	NModelsFilterer   // Log filterer for contract events
}

// NModelsCaller is an auto generated read-only Go binding around an Ethereum contract.
type NModelsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NModelsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NModelsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NModelsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NModelsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NModelsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NModelsSession struct {
	Contract     *NModels          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NModelsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NModelsCallerSession struct {
	Contract *NModelsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// NModelsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NModelsTransactorSession struct {
	Contract     *NModelsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// NModelsRaw is an auto generated low-level Go binding around an Ethereum contract.
type NModelsRaw struct {
	Contract *NModels // Generic contract binding to access the raw methods on
}

// NModelsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NModelsCallerRaw struct {
	Contract *NModelsCaller // Generic read-only contract binding to access the raw methods on
}

// NModelsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NModelsTransactorRaw struct {
	Contract *NModelsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNModels creates a new instance of NModels, bound to a specific deployed contract.
func NewNModels(address common.Address, backend bind.ContractBackend) (*NModels, error) {
	contract, err := bindNModels(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NModels{NModelsCaller: NModelsCaller{contract: contract}, NModelsTransactor: NModelsTransactor{contract: contract}, NModelsFilterer: NModelsFilterer{contract: contract}}, nil
}

// NewNModelsCaller creates a new read-only instance of NModels, bound to a specific deployed contract.
func NewNModelsCaller(address common.Address, caller bind.ContractCaller) (*NModelsCaller, error) {
	contract, err := bindNModels(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NModelsCaller{contract: contract}, nil
}

// NewNModelsTransactor creates a new write-only instance of NModels, bound to a specific deployed contract.
func NewNModelsTransactor(address common.Address, transactor bind.ContractTransactor) (*NModelsTransactor, error) {
	contract, err := bindNModels(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NModelsTransactor{contract: contract}, nil
}

// NewNModelsFilterer creates a new log filterer instance of NModels, bound to a specific deployed contract.
func NewNModelsFilterer(address common.Address, filterer bind.ContractFilterer) (*NModelsFilterer, error) {
	contract, err := bindNModels(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NModelsFilterer{contract: contract}, nil
}

// bindNModels binds a generic wrapper to an already deployed contract.
func bindNModels(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NModelsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NModels *NModelsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NModels.Contract.NModelsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NModels *NModelsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NModels.Contract.NModelsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NModels *NModelsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NModels.Contract.NModelsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NModels *NModelsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NModels.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NModels *NModelsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NModels.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NModels *NModelsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NModels.Contract.contract.Transact(opts, method, params...)
}

// NebulaABI is the input ABI used to generate the binding from.
const NebulaABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newOracle\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"newGravityContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newBftValue\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"NewPulse\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"NewSubscriber\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"bftValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"confirmData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOracles\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSubscriptionIds\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newOracles\",\"type\":\"address[]\"}],\"name\":\"hashNewOracles\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isPublseSubSent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracleQueue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"first\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"last\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"oracles\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pulseQueue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"first\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"last\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pulses\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rounds\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"subscriptionId\",\"type\":\"bytes32\"}],\"name\":\"sendData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"minConfirmations\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"subscribe\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"subscriptionIds\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"subscriptions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"minConfirmations\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"subscriptionsQueue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"first\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"last\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newOracles\",\"type\":\"address[]\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"newRound\",\"type\":\"uint256\"}],\"name\":\"updateOracles\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// NebulaFuncSigs maps the 4-byte function signature to its string representation.
var NebulaFuncSigs = map[string]string{
	"3cec1bdd": "bftValue()",
	"60be26d9": "confirmData(bytes32,uint8[],bytes32[],bytes32[])",
	"40884c52": "getOracles()",
	"35091ab1": "getSubscriptionIds()",
	"8bec345f": "hashNewOracles(address[])",
	"6148d3f3": "isPublseSubSent(uint256,bytes32)",
	"69a4246d": "oracleQueue()",
	"5b69a7d8": "oracles(uint256)",
	"1d11f944": "pulseQueue()",
	"0694fbb3": "pulses(uint256)",
	"8c65c81f": "rounds(uint256)",
	"1b81a640": "sendData(bytes,uint256,bytes32)",
	"3527715d": "subscribe(address,uint8,uint256)",
	"8cafc358": "subscriptionIds(uint256)",
	"94259c6c": "subscriptions(bytes32)",
	"b48a9c9b": "subscriptionsQueue()",
	"febae9ea": "updateOracles(address[],uint8[],bytes32[],bytes32[],uint256)",
}

// NebulaBin is the compiled bytecode used for deploying new contracts.
var NebulaBin = "0x60806040523480156200001157600080fd5b50604051620016cb380380620016cb833981810160405260608110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b9083019060208201858111156200006e57600080fd5b82518660208202830111640100000000821117156200008c57600080fd5b82525081516020918201928201910280838360005b83811015620000bb578181015183820152602001620000a1565b50505050919091016040908152602083810151939091015186519395509350620000ed92600d92509086019062000118565b50600e55600f80546001600160a01b0319166001600160a01b039290921691909117905550620001ac565b82805482825590600052602060002090810192821562000170579160200282015b828111156200017057825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000139565b506200017e92915062000182565b5090565b620001a991905b808211156200017e5780546001600160a01b031916815560010162000189565b90565b61150f80620001bc6000396000f3fe6080604052600436106100fe5760003560e01c806360be26d9116100955780638c65c81f116100645780638c65c81f146105c25780638cafc358146105ec57806394259c6c14610616578063b48a9c9b14610674578063febae9ea14610689576100fe565b806360be26d9146103025780636148d3f3146104bb57806369a4246d146104ff5780638bec345f14610514576100fe565b80633527715d116100d15780633527715d146102505780633cec1bdd1461029257806340884c52146102a75780635b69a7d8146102bc576100fe565b80630694fbb3146101005780631b81a6401461013c5780631d11f944146101bd57806335091ab1146101eb575b005b34801561010c57600080fd5b5061012a6004803603602081101561012357600080fd5b50356108bf565b60408051918252519081900360200190f35b34801561014857600080fd5b506100fe6004803603606081101561015f57600080fd5b810190602081018135600160201b81111561017957600080fd5b82018360208201111561018b57600080fd5b803590602001918460018302840111600160201b831117156101ac57600080fd5b9193509150803590602001356108d1565b3480156101c957600080fd5b506101d2610b20565b6040805192835260208301919091528051918290030190f35b3480156101f757600080fd5b50610200610b29565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561023c578181015183820152602001610224565b505050509050019250505060405180910390f35b34801561025c57600080fd5b506100fe6004803603606081101561027357600080fd5b506001600160a01b038135169060ff6020820135169060400135610b82565b34801561029e57600080fd5b5061012a610e10565b3480156102b357600080fd5b50610200610e16565b3480156102c857600080fd5b506102e6600480360360208110156102df57600080fd5b5035610e77565b604080516001600160a01b039092168252519081900360200190f35b34801561030e57600080fd5b506100fe6004803603608081101561032557600080fd5b81359190810190604081016020820135600160201b81111561034657600080fd5b82018360208201111561035857600080fd5b803590602001918460208302840111600160201b8311171561037957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156103c857600080fd5b8201836020820111156103da57600080fd5b803590602001918460208302840111600160201b831117156103fb57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561044a57600080fd5b82018360208201111561045c57600080fd5b803590602001918460208302840111600160201b8311171561047d57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610e9e945050505050565b3480156104c757600080fd5b506104eb600480360360408110156104de57600080fd5b5080359060200135611052565b604080519115158252519081900360200190f35b34801561050b57600080fd5b506101d2611072565b34801561052057600080fd5b5061012a6004803603602081101561053757600080fd5b810190602081018135600160201b81111561055157600080fd5b82018360208201111561056357600080fd5b803590602001918460208302840111600160201b8311171561058457600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061107b945050505050565b3480156105ce57600080fd5b506104eb600480360360208110156105e557600080fd5b503561113f565b3480156105f857600080fd5b5061012a6004803603602081101561060f57600080fd5b5035611154565b34801561062257600080fd5b506106406004803603602081101561063957600080fd5b5035611172565b604080516001600160a01b03958616815293909416602084015260ff90911682840152606082015290519081900360800190f35b34801561068057600080fd5b506101d26111ab565b34801561069557600080fd5b506100fe600480360360a08110156106ac57600080fd5b810190602081018135600160201b8111156106c657600080fd5b8201836020820111156106d857600080fd5b803590602001918460208302840111600160201b831117156106f957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561074857600080fd5b82018360208201111561075a57600080fd5b803590602001918460208302840111600160201b8311171561077b57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156107ca57600080fd5b8201836020820111156107dc57600080fd5b803590602001918460208302840111600160201b831117156107fd57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561084c57600080fd5b82018360208201111561085e57600080fd5b803590602001918460208302840111600160201b8311171561087f57600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955050913592506111b4915050565b60126020526000908152604090205481565b43600101821115610920576040805162461bcd60e51b815260206004820152601460248201527334b73b30b634b210313637b1b590373ab6b132b960611b604482015290519081900360640190fd5b600082815260126020908152604091829020549151869186910180838380828437808301925050509250505060405160208183030381529060405280519060200120146109a4576040805162461bcd60e51b815260206004820152600d60248201526c0d0c2e6d040dad2e6dac2e8c6d609b1b604482015290519081900360640190fd5b600082815260136020908152604080832084845290915290205460ff16156109fe576040805162461bcd60e51b81526020600482015260086024820152671cdd58881cd95b9d60c21b604482015290519081900360640190fd5b60008281526013602090815260408083208484528252808320805460ff191660019081179091556011835292819020909201549151631bf7b3cd60e11b8152600481019182526024810186905247926001600160a01b0316916337ef679a91889188918190604401848480828437600081840152601f19601f8201169050808301925050509350505050600060405180830381600087803b158015610aa257600080fd5b505af1158015610ab6573d6000803e3d6000fd5b50505060008381526011602052604090206002015447915082820390811015610b17576040805162461bcd60e51b815260206004820152600e60248201526d1a5b9d985b1a59081c995dd85c9960921b604482015290519081900360640190fd5b50505050505050565b600954600a5482565b60606010805480602002602001604051908101604052809291908181526020018280548015610b7757602002820191906000526020600020905b815481526020019060010190808311610b63575b505050505090505b90565b60408051600080356001600160e01b03191660208084019190915233606090811b602485015287901b6bffffffffffffffffffffffff1916603884015260f886901b6001600160f81b031916604c8401528351808403602d018152604d84019094528351919392606d019182918401908083835b60208310610c155780518252601f199092019160209182019101610bf6565b51815160209384036101000a60001901801990921691161790526040805192909401828103601f1901835284528151918101919091206000818152601190925292902054919450506001600160a01b0316159150610caa9050576040805162461bcd60e51b815260206004820152600b60248201526a1c9c481a5cc8195e1a5cdd60aa1b604482015290519081900360640190fd5b604080516080810182523381526001600160a01b03868116602080840191825260ff8881168587019081526060860189815260008981526011909452878420965187546001600160a01b0319908116918816919091178855945160018801805493519390961696169590951760ff60a01b1916600160a01b91909216021790915590516002909201919091558151632941b65560e21b81526005600482015260248101849052915173__$9e8e06cf00219ff9bf6fefc812eef484ba$__9263a506d954926044808301939192829003018186803b158015610d8a57600080fd5b505af4158015610d9e573d6000803e3d6000fd5b5050601080546001810182556000919091527f1b6847dc741a1b0cd08d278845f9d819d87b734759afb55fe2de5cb82a9ae6720183905550506040805182815290517fce45703d589912fa7e0caf69c208323a21733441d675118cd08e8b4479c61daf9181900360200190a150505050565b600e5481565b6060600d805480602002602001604051908101604052809291908181526020018280548015610b7757602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610e50575050505050905090565b600d8181548110610e8457fe5b6000918252602090912001546001600160a01b0316905081565b6000805b600d54811015610fa857600d8181548110610eb957fe5b9060005260206000200160009054906101000a90046001600160a01b03166001600160a01b0316600187878481518110610eef57fe5b6020026020010151878581518110610f0357fe5b6020026020010151878681518110610f1757fe5b602002602001015160405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015610f76573d6000803e3d6000fd5b505050602060405103516001600160a01b031614610f95576000610f98565b60015b60ff169190910190600101610ea2565b50600e54811015610ff4576040805162461bcd60e51b81526020600482015260116024820152701a5b9d985b1a590818999d0818dbdd5b9d607a1b604482015290519081900360640190fd5b60408051602080820183528782524360008181526012835284902092519092558251918252810187905281517fd616b4aa280263f7f493a9f9952600e59057001ce5ecaa1428e86f1b3e276d51929181900390910190a15050505050565b601360209081526000928352604080842090915290825290205460ff1681565b60015460025482565b60006060815b8351811015611130578184828151811061109757fe5b60200260200101516040516020018083805190602001908083835b602083106110d15780518252601f1990920191602091820191016110b2565b6001836020036101000a038019825116818451168082178552505050505050905001826001600160a01b03166001600160a01b031660601b81526014019250505060405160208183030381529060405291508080600101915050611081565b50805160209091012092915050565b60006020819052908152604090205460ff1681565b6010818154811061116157fe5b600091825260209091200154905081565b6011602052600090815260409020805460018201546002909201546001600160a01b039182169291821691600160a01b900460ff169084565b60055460065482565b6000806111c08761107b565b90506060600f60009054906101000a90046001600160a01b03166001600160a01b031663ad595b1a6040518163ffffffff1660e01b815260040160006040518083038186803b15801561121257600080fd5b505afa158015611226573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561124f57600080fd5b8101908080516040519392919084600160201b82111561126e57600080fd5b90830190602082018581111561128357600080fd5b82518660208202830111600160201b8211171561129f57600080fd5b82525081516020918201928201910280838360005b838110156112cc5781810151838201526020016112b4565b50505050905001604052505050905060008090505b81518110156113cf578181815181106112f657fe5b60200260200101516001600160a01b03166001848a848151811061131657fe5b60200260200101518a858151811061132a57fe5b60200260200101518a868151811061133e57fe5b602002602001015160405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa15801561139d573d6000803e3d6000fd5b505050602060405103516001600160a01b0316146113bc5760006113bf565b60015b60ff1693909301926001016112e1565b50600e5483101561141b576040805162461bcd60e51b81526020600482015260116024820152701a5b9d985b1a590818999d0818dbdd5b9d607a1b604482015290519081900360640190fd5b875161142e90600d9060208b0190611451565b505050600091825250602081905260409020805460ff1916600117905550505050565b8280548282559060005260206000209081019282156114a6579160200282015b828111156114a657825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190611471565b506114b29291506114b6565b5090565b610b7f91905b808211156114b25780546001600160a01b03191681556001016114bc56fea265627a7a72315820d9e44476477f661161d540a7ff8eff54666840ef1011061cd96a878be67e0e6764736f6c63430005100032"

// DeployNebula deploys a new Ethereum contract, binding an instance of Nebula to it.
func DeployNebula(auth *bind.TransactOpts, backend bind.ContractBackend, newOracle []common.Address, newGravityContract common.Address, newBftValue *big.Int) (common.Address, *types.Transaction, *Nebula, error) {
	parsed, err := abi.JSON(strings.NewReader(NebulaABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	queueLibAddr, _, _, _ := DeployQueueLib(auth, backend)
	NebulaBin = strings.Replace(NebulaBin, "__$9e8e06cf00219ff9bf6fefc812eef484ba$__", queueLibAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NebulaBin), backend, newOracle, newGravityContract, newBftValue)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Nebula{NebulaCaller: NebulaCaller{contract: contract}, NebulaTransactor: NebulaTransactor{contract: contract}, NebulaFilterer: NebulaFilterer{contract: contract}}, nil
}

// Nebula is an auto generated Go binding around an Ethereum contract.
type Nebula struct {
	NebulaCaller     // Read-only binding to the contract
	NebulaTransactor // Write-only binding to the contract
	NebulaFilterer   // Log filterer for contract events
}

// NebulaCaller is an auto generated read-only Go binding around an Ethereum contract.
type NebulaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NebulaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NebulaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NebulaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NebulaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NebulaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NebulaSession struct {
	Contract     *Nebula           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NebulaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NebulaCallerSession struct {
	Contract *NebulaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NebulaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NebulaTransactorSession struct {
	Contract     *NebulaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NebulaRaw is an auto generated low-level Go binding around an Ethereum contract.
type NebulaRaw struct {
	Contract *Nebula // Generic contract binding to access the raw methods on
}

// NebulaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NebulaCallerRaw struct {
	Contract *NebulaCaller // Generic read-only contract binding to access the raw methods on
}

// NebulaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NebulaTransactorRaw struct {
	Contract *NebulaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNebula creates a new instance of Nebula, bound to a specific deployed contract.
func NewNebula(address common.Address, backend bind.ContractBackend) (*Nebula, error) {
	contract, err := bindNebula(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nebula{NebulaCaller: NebulaCaller{contract: contract}, NebulaTransactor: NebulaTransactor{contract: contract}, NebulaFilterer: NebulaFilterer{contract: contract}}, nil
}

// NewNebulaCaller creates a new read-only instance of Nebula, bound to a specific deployed contract.
func NewNebulaCaller(address common.Address, caller bind.ContractCaller) (*NebulaCaller, error) {
	contract, err := bindNebula(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NebulaCaller{contract: contract}, nil
}

// NewNebulaTransactor creates a new write-only instance of Nebula, bound to a specific deployed contract.
func NewNebulaTransactor(address common.Address, transactor bind.ContractTransactor) (*NebulaTransactor, error) {
	contract, err := bindNebula(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NebulaTransactor{contract: contract}, nil
}

// NewNebulaFilterer creates a new log filterer instance of Nebula, bound to a specific deployed contract.
func NewNebulaFilterer(address common.Address, filterer bind.ContractFilterer) (*NebulaFilterer, error) {
	contract, err := bindNebula(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NebulaFilterer{contract: contract}, nil
}

// bindNebula binds a generic wrapper to an already deployed contract.
func bindNebula(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NebulaABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nebula *NebulaRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Nebula.Contract.NebulaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nebula *NebulaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nebula.Contract.NebulaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nebula *NebulaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nebula.Contract.NebulaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nebula *NebulaCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Nebula.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nebula *NebulaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nebula.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nebula *NebulaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nebula.Contract.contract.Transact(opts, method, params...)
}

// BftValue is a free data retrieval call binding the contract method 0x3cec1bdd.
//
// Solidity: function bftValue() view returns(uint256)
func (_Nebula *NebulaCaller) BftValue(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Nebula.contract.Call(opts, out, "bftValue")
	return *ret0, err
}

// BftValue is a free data retrieval call binding the contract method 0x3cec1bdd.
//
// Solidity: function bftValue() view returns(uint256)
func (_Nebula *NebulaSession) BftValue() (*big.Int, error) {
	return _Nebula.Contract.BftValue(&_Nebula.CallOpts)
}

// BftValue is a free data retrieval call binding the contract method 0x3cec1bdd.
//
// Solidity: function bftValue() view returns(uint256)
func (_Nebula *NebulaCallerSession) BftValue() (*big.Int, error) {
	return _Nebula.Contract.BftValue(&_Nebula.CallOpts)
}

// GetOracles is a free data retrieval call binding the contract method 0x40884c52.
//
// Solidity: function getOracles() view returns(address[])
func (_Nebula *NebulaCaller) GetOracles(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Nebula.contract.Call(opts, out, "getOracles")
	return *ret0, err
}

// GetOracles is a free data retrieval call binding the contract method 0x40884c52.
//
// Solidity: function getOracles() view returns(address[])
func (_Nebula *NebulaSession) GetOracles() ([]common.Address, error) {
	return _Nebula.Contract.GetOracles(&_Nebula.CallOpts)
}

// GetOracles is a free data retrieval call binding the contract method 0x40884c52.
//
// Solidity: function getOracles() view returns(address[])
func (_Nebula *NebulaCallerSession) GetOracles() ([]common.Address, error) {
	return _Nebula.Contract.GetOracles(&_Nebula.CallOpts)
}

// GetSubscriptionIds is a free data retrieval call binding the contract method 0x35091ab1.
//
// Solidity: function getSubscriptionIds() view returns(bytes32[])
func (_Nebula *NebulaCaller) GetSubscriptionIds(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _Nebula.contract.Call(opts, out, "getSubscriptionIds")
	return *ret0, err
}

// GetSubscriptionIds is a free data retrieval call binding the contract method 0x35091ab1.
//
// Solidity: function getSubscriptionIds() view returns(bytes32[])
func (_Nebula *NebulaSession) GetSubscriptionIds() ([][32]byte, error) {
	return _Nebula.Contract.GetSubscriptionIds(&_Nebula.CallOpts)
}

// GetSubscriptionIds is a free data retrieval call binding the contract method 0x35091ab1.
//
// Solidity: function getSubscriptionIds() view returns(bytes32[])
func (_Nebula *NebulaCallerSession) GetSubscriptionIds() ([][32]byte, error) {
	return _Nebula.Contract.GetSubscriptionIds(&_Nebula.CallOpts)
}

// HashNewOracles is a free data retrieval call binding the contract method 0x8bec345f.
//
// Solidity: function hashNewOracles(address[] newOracles) pure returns(bytes32)
func (_Nebula *NebulaCaller) HashNewOracles(opts *bind.CallOpts, newOracles []common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Nebula.contract.Call(opts, out, "hashNewOracles", newOracles)
	return *ret0, err
}

// HashNewOracles is a free data retrieval call binding the contract method 0x8bec345f.
//
// Solidity: function hashNewOracles(address[] newOracles) pure returns(bytes32)
func (_Nebula *NebulaSession) HashNewOracles(newOracles []common.Address) ([32]byte, error) {
	return _Nebula.Contract.HashNewOracles(&_Nebula.CallOpts, newOracles)
}

// HashNewOracles is a free data retrieval call binding the contract method 0x8bec345f.
//
// Solidity: function hashNewOracles(address[] newOracles) pure returns(bytes32)
func (_Nebula *NebulaCallerSession) HashNewOracles(newOracles []common.Address) ([32]byte, error) {
	return _Nebula.Contract.HashNewOracles(&_Nebula.CallOpts, newOracles)
}

// IsPublseSubSent is a free data retrieval call binding the contract method 0x6148d3f3.
//
// Solidity: function isPublseSubSent(uint256 , bytes32 ) view returns(bool)
func (_Nebula *NebulaCaller) IsPublseSubSent(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Nebula.contract.Call(opts, out, "isPublseSubSent", arg0, arg1)
	return *ret0, err
}

// IsPublseSubSent is a free data retrieval call binding the contract method 0x6148d3f3.
//
// Solidity: function isPublseSubSent(uint256 , bytes32 ) view returns(bool)
func (_Nebula *NebulaSession) IsPublseSubSent(arg0 *big.Int, arg1 [32]byte) (bool, error) {
	return _Nebula.Contract.IsPublseSubSent(&_Nebula.CallOpts, arg0, arg1)
}

// IsPublseSubSent is a free data retrieval call binding the contract method 0x6148d3f3.
//
// Solidity: function isPublseSubSent(uint256 , bytes32 ) view returns(bool)
func (_Nebula *NebulaCallerSession) IsPublseSubSent(arg0 *big.Int, arg1 [32]byte) (bool, error) {
	return _Nebula.Contract.IsPublseSubSent(&_Nebula.CallOpts, arg0, arg1)
}

// OracleQueue is a free data retrieval call binding the contract method 0x69a4246d.
//
// Solidity: function oracleQueue() view returns(bytes32 first, bytes32 last)
func (_Nebula *NebulaCaller) OracleQueue(opts *bind.CallOpts) (struct {
	First [32]byte
	Last  [32]byte
}, error) {
	ret := new(struct {
		First [32]byte
		Last  [32]byte
	})
	out := ret
	err := _Nebula.contract.Call(opts, out, "oracleQueue")
	return *ret, err
}

// OracleQueue is a free data retrieval call binding the contract method 0x69a4246d.
//
// Solidity: function oracleQueue() view returns(bytes32 first, bytes32 last)
func (_Nebula *NebulaSession) OracleQueue() (struct {
	First [32]byte
	Last  [32]byte
}, error) {
	return _Nebula.Contract.OracleQueue(&_Nebula.CallOpts)
}

// OracleQueue is a free data retrieval call binding the contract method 0x69a4246d.
//
// Solidity: function oracleQueue() view returns(bytes32 first, bytes32 last)
func (_Nebula *NebulaCallerSession) OracleQueue() (struct {
	First [32]byte
	Last  [32]byte
}, error) {
	return _Nebula.Contract.OracleQueue(&_Nebula.CallOpts)
}

// Oracles is a free data retrieval call binding the contract method 0x5b69a7d8.
//
// Solidity: function oracles(uint256 ) view returns(address)
func (_Nebula *NebulaCaller) Oracles(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Nebula.contract.Call(opts, out, "oracles", arg0)
	return *ret0, err
}

// Oracles is a free data retrieval call binding the contract method 0x5b69a7d8.
//
// Solidity: function oracles(uint256 ) view returns(address)
func (_Nebula *NebulaSession) Oracles(arg0 *big.Int) (common.Address, error) {
	return _Nebula.Contract.Oracles(&_Nebula.CallOpts, arg0)
}

// Oracles is a free data retrieval call binding the contract method 0x5b69a7d8.
//
// Solidity: function oracles(uint256 ) view returns(address)
func (_Nebula *NebulaCallerSession) Oracles(arg0 *big.Int) (common.Address, error) {
	return _Nebula.Contract.Oracles(&_Nebula.CallOpts, arg0)
}

// PulseQueue is a free data retrieval call binding the contract method 0x1d11f944.
//
// Solidity: function pulseQueue() view returns(bytes32 first, bytes32 last)
func (_Nebula *NebulaCaller) PulseQueue(opts *bind.CallOpts) (struct {
	First [32]byte
	Last  [32]byte
}, error) {
	ret := new(struct {
		First [32]byte
		Last  [32]byte
	})
	out := ret
	err := _Nebula.contract.Call(opts, out, "pulseQueue")
	return *ret, err
}

// PulseQueue is a free data retrieval call binding the contract method 0x1d11f944.
//
// Solidity: function pulseQueue() view returns(bytes32 first, bytes32 last)
func (_Nebula *NebulaSession) PulseQueue() (struct {
	First [32]byte
	Last  [32]byte
}, error) {
	return _Nebula.Contract.PulseQueue(&_Nebula.CallOpts)
}

// PulseQueue is a free data retrieval call binding the contract method 0x1d11f944.
//
// Solidity: function pulseQueue() view returns(bytes32 first, bytes32 last)
func (_Nebula *NebulaCallerSession) PulseQueue() (struct {
	First [32]byte
	Last  [32]byte
}, error) {
	return _Nebula.Contract.PulseQueue(&_Nebula.CallOpts)
}

// Pulses is a free data retrieval call binding the contract method 0x0694fbb3.
//
// Solidity: function pulses(uint256 ) view returns(bytes32 dataHash)
func (_Nebula *NebulaCaller) Pulses(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Nebula.contract.Call(opts, out, "pulses", arg0)
	return *ret0, err
}

// Pulses is a free data retrieval call binding the contract method 0x0694fbb3.
//
// Solidity: function pulses(uint256 ) view returns(bytes32 dataHash)
func (_Nebula *NebulaSession) Pulses(arg0 *big.Int) ([32]byte, error) {
	return _Nebula.Contract.Pulses(&_Nebula.CallOpts, arg0)
}

// Pulses is a free data retrieval call binding the contract method 0x0694fbb3.
//
// Solidity: function pulses(uint256 ) view returns(bytes32 dataHash)
func (_Nebula *NebulaCallerSession) Pulses(arg0 *big.Int) ([32]byte, error) {
	return _Nebula.Contract.Pulses(&_Nebula.CallOpts, arg0)
}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(bool)
func (_Nebula *NebulaCaller) Rounds(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Nebula.contract.Call(opts, out, "rounds", arg0)
	return *ret0, err
}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(bool)
func (_Nebula *NebulaSession) Rounds(arg0 *big.Int) (bool, error) {
	return _Nebula.Contract.Rounds(&_Nebula.CallOpts, arg0)
}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(bool)
func (_Nebula *NebulaCallerSession) Rounds(arg0 *big.Int) (bool, error) {
	return _Nebula.Contract.Rounds(&_Nebula.CallOpts, arg0)
}

// SubscriptionIds is a free data retrieval call binding the contract method 0x8cafc358.
//
// Solidity: function subscriptionIds(uint256 ) view returns(bytes32)
func (_Nebula *NebulaCaller) SubscriptionIds(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Nebula.contract.Call(opts, out, "subscriptionIds", arg0)
	return *ret0, err
}

// SubscriptionIds is a free data retrieval call binding the contract method 0x8cafc358.
//
// Solidity: function subscriptionIds(uint256 ) view returns(bytes32)
func (_Nebula *NebulaSession) SubscriptionIds(arg0 *big.Int) ([32]byte, error) {
	return _Nebula.Contract.SubscriptionIds(&_Nebula.CallOpts, arg0)
}

// SubscriptionIds is a free data retrieval call binding the contract method 0x8cafc358.
//
// Solidity: function subscriptionIds(uint256 ) view returns(bytes32)
func (_Nebula *NebulaCallerSession) SubscriptionIds(arg0 *big.Int) ([32]byte, error) {
	return _Nebula.Contract.SubscriptionIds(&_Nebula.CallOpts, arg0)
}

// Subscriptions is a free data retrieval call binding the contract method 0x94259c6c.
//
// Solidity: function subscriptions(bytes32 ) view returns(address owner, address contractAddress, uint8 minConfirmations, uint256 reward)
func (_Nebula *NebulaCaller) Subscriptions(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Owner            common.Address
	ContractAddress  common.Address
	MinConfirmations uint8
	Reward           *big.Int
}, error) {
	ret := new(struct {
		Owner            common.Address
		ContractAddress  common.Address
		MinConfirmations uint8
		Reward           *big.Int
	})
	out := ret
	err := _Nebula.contract.Call(opts, out, "subscriptions", arg0)
	return *ret, err
}

// Subscriptions is a free data retrieval call binding the contract method 0x94259c6c.
//
// Solidity: function subscriptions(bytes32 ) view returns(address owner, address contractAddress, uint8 minConfirmations, uint256 reward)
func (_Nebula *NebulaSession) Subscriptions(arg0 [32]byte) (struct {
	Owner            common.Address
	ContractAddress  common.Address
	MinConfirmations uint8
	Reward           *big.Int
}, error) {
	return _Nebula.Contract.Subscriptions(&_Nebula.CallOpts, arg0)
}

// Subscriptions is a free data retrieval call binding the contract method 0x94259c6c.
//
// Solidity: function subscriptions(bytes32 ) view returns(address owner, address contractAddress, uint8 minConfirmations, uint256 reward)
func (_Nebula *NebulaCallerSession) Subscriptions(arg0 [32]byte) (struct {
	Owner            common.Address
	ContractAddress  common.Address
	MinConfirmations uint8
	Reward           *big.Int
}, error) {
	return _Nebula.Contract.Subscriptions(&_Nebula.CallOpts, arg0)
}

// SubscriptionsQueue is a free data retrieval call binding the contract method 0xb48a9c9b.
//
// Solidity: function subscriptionsQueue() view returns(bytes32 first, bytes32 last)
func (_Nebula *NebulaCaller) SubscriptionsQueue(opts *bind.CallOpts) (struct {
	First [32]byte
	Last  [32]byte
}, error) {
	ret := new(struct {
		First [32]byte
		Last  [32]byte
	})
	out := ret
	err := _Nebula.contract.Call(opts, out, "subscriptionsQueue")
	return *ret, err
}

// SubscriptionsQueue is a free data retrieval call binding the contract method 0xb48a9c9b.
//
// Solidity: function subscriptionsQueue() view returns(bytes32 first, bytes32 last)
func (_Nebula *NebulaSession) SubscriptionsQueue() (struct {
	First [32]byte
	Last  [32]byte
}, error) {
	return _Nebula.Contract.SubscriptionsQueue(&_Nebula.CallOpts)
}

// SubscriptionsQueue is a free data retrieval call binding the contract method 0xb48a9c9b.
//
// Solidity: function subscriptionsQueue() view returns(bytes32 first, bytes32 last)
func (_Nebula *NebulaCallerSession) SubscriptionsQueue() (struct {
	First [32]byte
	Last  [32]byte
}, error) {
	return _Nebula.Contract.SubscriptionsQueue(&_Nebula.CallOpts)
}

// ConfirmData is a paid mutator transaction binding the contract method 0x60be26d9.
//
// Solidity: function confirmData(bytes32 dataHash, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_Nebula *NebulaTransactor) ConfirmData(opts *bind.TransactOpts, dataHash [32]byte, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Nebula.contract.Transact(opts, "confirmData", dataHash, v, r, s)
}

// ConfirmData is a paid mutator transaction binding the contract method 0x60be26d9.
//
// Solidity: function confirmData(bytes32 dataHash, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_Nebula *NebulaSession) ConfirmData(dataHash [32]byte, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Nebula.Contract.ConfirmData(&_Nebula.TransactOpts, dataHash, v, r, s)
}

// ConfirmData is a paid mutator transaction binding the contract method 0x60be26d9.
//
// Solidity: function confirmData(bytes32 dataHash, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_Nebula *NebulaTransactorSession) ConfirmData(dataHash [32]byte, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Nebula.Contract.ConfirmData(&_Nebula.TransactOpts, dataHash, v, r, s)
}

// SendData is a paid mutator transaction binding the contract method 0x1b81a640.
//
// Solidity: function sendData(bytes value, uint256 blockNumber, bytes32 subscriptionId) returns()
func (_Nebula *NebulaTransactor) SendData(opts *bind.TransactOpts, value []byte, blockNumber *big.Int, subscriptionId [32]byte) (*types.Transaction, error) {
	return _Nebula.contract.Transact(opts, "sendData", value, blockNumber, subscriptionId)
}

// SendData is a paid mutator transaction binding the contract method 0x1b81a640.
//
// Solidity: function sendData(bytes value, uint256 blockNumber, bytes32 subscriptionId) returns()
func (_Nebula *NebulaSession) SendData(value []byte, blockNumber *big.Int, subscriptionId [32]byte) (*types.Transaction, error) {
	return _Nebula.Contract.SendData(&_Nebula.TransactOpts, value, blockNumber, subscriptionId)
}

// SendData is a paid mutator transaction binding the contract method 0x1b81a640.
//
// Solidity: function sendData(bytes value, uint256 blockNumber, bytes32 subscriptionId) returns()
func (_Nebula *NebulaTransactorSession) SendData(value []byte, blockNumber *big.Int, subscriptionId [32]byte) (*types.Transaction, error) {
	return _Nebula.Contract.SendData(&_Nebula.TransactOpts, value, blockNumber, subscriptionId)
}

// Subscribe is a paid mutator transaction binding the contract method 0x3527715d.
//
// Solidity: function subscribe(address contractAddress, uint8 minConfirmations, uint256 reward) returns()
func (_Nebula *NebulaTransactor) Subscribe(opts *bind.TransactOpts, contractAddress common.Address, minConfirmations uint8, reward *big.Int) (*types.Transaction, error) {
	return _Nebula.contract.Transact(opts, "subscribe", contractAddress, minConfirmations, reward)
}

// Subscribe is a paid mutator transaction binding the contract method 0x3527715d.
//
// Solidity: function subscribe(address contractAddress, uint8 minConfirmations, uint256 reward) returns()
func (_Nebula *NebulaSession) Subscribe(contractAddress common.Address, minConfirmations uint8, reward *big.Int) (*types.Transaction, error) {
	return _Nebula.Contract.Subscribe(&_Nebula.TransactOpts, contractAddress, minConfirmations, reward)
}

// Subscribe is a paid mutator transaction binding the contract method 0x3527715d.
//
// Solidity: function subscribe(address contractAddress, uint8 minConfirmations, uint256 reward) returns()
func (_Nebula *NebulaTransactorSession) Subscribe(contractAddress common.Address, minConfirmations uint8, reward *big.Int) (*types.Transaction, error) {
	return _Nebula.Contract.Subscribe(&_Nebula.TransactOpts, contractAddress, minConfirmations, reward)
}

// UpdateOracles is a paid mutator transaction binding the contract method 0xfebae9ea.
//
// Solidity: function updateOracles(address[] newOracles, uint8[] v, bytes32[] r, bytes32[] s, uint256 newRound) returns()
func (_Nebula *NebulaTransactor) UpdateOracles(opts *bind.TransactOpts, newOracles []common.Address, v []uint8, r [][32]byte, s [][32]byte, newRound *big.Int) (*types.Transaction, error) {
	return _Nebula.contract.Transact(opts, "updateOracles", newOracles, v, r, s, newRound)
}

// UpdateOracles is a paid mutator transaction binding the contract method 0xfebae9ea.
//
// Solidity: function updateOracles(address[] newOracles, uint8[] v, bytes32[] r, bytes32[] s, uint256 newRound) returns()
func (_Nebula *NebulaSession) UpdateOracles(newOracles []common.Address, v []uint8, r [][32]byte, s [][32]byte, newRound *big.Int) (*types.Transaction, error) {
	return _Nebula.Contract.UpdateOracles(&_Nebula.TransactOpts, newOracles, v, r, s, newRound)
}

// UpdateOracles is a paid mutator transaction binding the contract method 0xfebae9ea.
//
// Solidity: function updateOracles(address[] newOracles, uint8[] v, bytes32[] r, bytes32[] s, uint256 newRound) returns()
func (_Nebula *NebulaTransactorSession) UpdateOracles(newOracles []common.Address, v []uint8, r [][32]byte, s [][32]byte, newRound *big.Int) (*types.Transaction, error) {
	return _Nebula.Contract.UpdateOracles(&_Nebula.TransactOpts, newOracles, v, r, s, newRound)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Nebula *NebulaTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Nebula.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Nebula *NebulaSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Nebula.Contract.Fallback(&_Nebula.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Nebula *NebulaTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Nebula.Contract.Fallback(&_Nebula.TransactOpts, calldata)
}

// NebulaNewPulseIterator is returned from FilterNewPulse and is used to iterate over the raw logs and unpacked data for NewPulse events raised by the Nebula contract.
type NebulaNewPulseIterator struct {
	Event *NebulaNewPulse // Event containing the contract specifics and raw log

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
func (it *NebulaNewPulseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NebulaNewPulse)
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
		it.Event = new(NebulaNewPulse)
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
func (it *NebulaNewPulseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NebulaNewPulseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NebulaNewPulse represents a NewPulse event raised by the Nebula contract.
type NebulaNewPulse struct {
	Height   *big.Int
	DataHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewPulse is a free log retrieval operation binding the contract event 0xd616b4aa280263f7f493a9f9952600e59057001ce5ecaa1428e86f1b3e276d51.
//
// Solidity: event NewPulse(uint256 height, bytes32 dataHash)
func (_Nebula *NebulaFilterer) FilterNewPulse(opts *bind.FilterOpts) (*NebulaNewPulseIterator, error) {

	logs, sub, err := _Nebula.contract.FilterLogs(opts, "NewPulse")
	if err != nil {
		return nil, err
	}
	return &NebulaNewPulseIterator{contract: _Nebula.contract, event: "NewPulse", logs: logs, sub: sub}, nil
}

// WatchNewPulse is a free log subscription operation binding the contract event 0xd616b4aa280263f7f493a9f9952600e59057001ce5ecaa1428e86f1b3e276d51.
//
// Solidity: event NewPulse(uint256 height, bytes32 dataHash)
func (_Nebula *NebulaFilterer) WatchNewPulse(opts *bind.WatchOpts, sink chan<- *NebulaNewPulse) (event.Subscription, error) {

	logs, sub, err := _Nebula.contract.WatchLogs(opts, "NewPulse")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NebulaNewPulse)
				if err := _Nebula.contract.UnpackLog(event, "NewPulse", log); err != nil {
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

// ParseNewPulse is a log parse operation binding the contract event 0xd616b4aa280263f7f493a9f9952600e59057001ce5ecaa1428e86f1b3e276d51.
//
// Solidity: event NewPulse(uint256 height, bytes32 dataHash)
func (_Nebula *NebulaFilterer) ParseNewPulse(log types.Log) (*NebulaNewPulse, error) {
	event := new(NebulaNewPulse)
	if err := _Nebula.contract.UnpackLog(event, "NewPulse", log); err != nil {
		return nil, err
	}
	return event, nil
}

// NebulaNewSubscriberIterator is returned from FilterNewSubscriber and is used to iterate over the raw logs and unpacked data for NewSubscriber events raised by the Nebula contract.
type NebulaNewSubscriberIterator struct {
	Event *NebulaNewSubscriber // Event containing the contract specifics and raw log

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
func (it *NebulaNewSubscriberIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NebulaNewSubscriber)
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
		it.Event = new(NebulaNewSubscriber)
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
func (it *NebulaNewSubscriberIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NebulaNewSubscriberIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NebulaNewSubscriber represents a NewSubscriber event raised by the Nebula contract.
type NebulaNewSubscriber struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNewSubscriber is a free log retrieval operation binding the contract event 0xce45703d589912fa7e0caf69c208323a21733441d675118cd08e8b4479c61daf.
//
// Solidity: event NewSubscriber(bytes32 id)
func (_Nebula *NebulaFilterer) FilterNewSubscriber(opts *bind.FilterOpts) (*NebulaNewSubscriberIterator, error) {

	logs, sub, err := _Nebula.contract.FilterLogs(opts, "NewSubscriber")
	if err != nil {
		return nil, err
	}
	return &NebulaNewSubscriberIterator{contract: _Nebula.contract, event: "NewSubscriber", logs: logs, sub: sub}, nil
}

// WatchNewSubscriber is a free log subscription operation binding the contract event 0xce45703d589912fa7e0caf69c208323a21733441d675118cd08e8b4479c61daf.
//
// Solidity: event NewSubscriber(bytes32 id)
func (_Nebula *NebulaFilterer) WatchNewSubscriber(opts *bind.WatchOpts, sink chan<- *NebulaNewSubscriber) (event.Subscription, error) {

	logs, sub, err := _Nebula.contract.WatchLogs(opts, "NewSubscriber")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NebulaNewSubscriber)
				if err := _Nebula.contract.UnpackLog(event, "NewSubscriber", log); err != nil {
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

// ParseNewSubscriber is a log parse operation binding the contract event 0xce45703d589912fa7e0caf69c208323a21733441d675118cd08e8b4479c61daf.
//
// Solidity: event NewSubscriber(bytes32 id)
func (_Nebula *NebulaFilterer) ParseNewSubscriber(log types.Log) (*NebulaNewSubscriber, error) {
	event := new(NebulaNewSubscriber)
	if err := _Nebula.contract.UnpackLog(event, "NewSubscriber", log); err != nil {
		return nil, err
	}
	return event, nil
}

// QueueLibABI is the input ABI used to generate the binding from.
const QueueLibABI = "[]"

// QueueLibFuncSigs maps the 4-byte function signature to its string representation.
var QueueLibFuncSigs = map[string]string{
	"9d6ad84b": "drop(QueueLib.Queue storage,bytes32)",
	"870a61ff": "next(QueueLib.Queue storage,bytes32)",
	"a506d954": "push(QueueLib.Queue storage,bytes32)",
}

// QueueLibBin is the compiled bytecode used for deploying new contracts.
var QueueLibBin = "0x6101fb610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004b5760003560e01c8063870a61ff146100505780639d6ad84b14610085578063a506d954146100b7575b600080fd5b6100736004803603604081101561006657600080fd5b50803590602001356100e7565b60408051918252519081900360200190f35b81801561009157600080fd5b506100b5600480360360408110156100a857600080fd5b508035906020013561010f565b005b8180156100c357600080fd5b506100b5600480360360408110156100da57600080fd5b508035906020013561017e565b6000816100f657508154610109565b5060008181526002830160205260409020545b92915050565b8154811480156101225750808260010154145b15610136576000808355600183015561017a565b8154811415610157576000818152600283016020526040902054825561017a565b808260010154141561017a57600081815260038301602052604090205460018301555b5050565b8154610193578082556001820181905561017a565b6001820180546000908152600284016020908152604080832085905583548584526003870190925290912055819055505056fea265627a7a72315820e8968104c29099920ef41af6e996a34a9fbc159920bb00ff717e28d3042d8ee264736f6c63430005100032"

// DeployQueueLib deploys a new Ethereum contract, binding an instance of QueueLib to it.
func DeployQueueLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *QueueLib, error) {
	parsed, err := abi.JSON(strings.NewReader(QueueLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(QueueLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &QueueLib{QueueLibCaller: QueueLibCaller{contract: contract}, QueueLibTransactor: QueueLibTransactor{contract: contract}, QueueLibFilterer: QueueLibFilterer{contract: contract}}, nil
}

// QueueLib is an auto generated Go binding around an Ethereum contract.
type QueueLib struct {
	QueueLibCaller     // Read-only binding to the contract
	QueueLibTransactor // Write-only binding to the contract
	QueueLibFilterer   // Log filterer for contract events
}

// QueueLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type QueueLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QueueLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QueueLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QueueLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QueueLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QueueLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QueueLibSession struct {
	Contract     *QueueLib         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QueueLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QueueLibCallerSession struct {
	Contract *QueueLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// QueueLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QueueLibTransactorSession struct {
	Contract     *QueueLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// QueueLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type QueueLibRaw struct {
	Contract *QueueLib // Generic contract binding to access the raw methods on
}

// QueueLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QueueLibCallerRaw struct {
	Contract *QueueLibCaller // Generic read-only contract binding to access the raw methods on
}

// QueueLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QueueLibTransactorRaw struct {
	Contract *QueueLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQueueLib creates a new instance of QueueLib, bound to a specific deployed contract.
func NewQueueLib(address common.Address, backend bind.ContractBackend) (*QueueLib, error) {
	contract, err := bindQueueLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QueueLib{QueueLibCaller: QueueLibCaller{contract: contract}, QueueLibTransactor: QueueLibTransactor{contract: contract}, QueueLibFilterer: QueueLibFilterer{contract: contract}}, nil
}

// NewQueueLibCaller creates a new read-only instance of QueueLib, bound to a specific deployed contract.
func NewQueueLibCaller(address common.Address, caller bind.ContractCaller) (*QueueLibCaller, error) {
	contract, err := bindQueueLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QueueLibCaller{contract: contract}, nil
}

// NewQueueLibTransactor creates a new write-only instance of QueueLib, bound to a specific deployed contract.
func NewQueueLibTransactor(address common.Address, transactor bind.ContractTransactor) (*QueueLibTransactor, error) {
	contract, err := bindQueueLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QueueLibTransactor{contract: contract}, nil
}

// NewQueueLibFilterer creates a new log filterer instance of QueueLib, bound to a specific deployed contract.
func NewQueueLibFilterer(address common.Address, filterer bind.ContractFilterer) (*QueueLibFilterer, error) {
	contract, err := bindQueueLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QueueLibFilterer{contract: contract}, nil
}

// bindQueueLib binds a generic wrapper to an already deployed contract.
func bindQueueLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(QueueLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QueueLib *QueueLibRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _QueueLib.Contract.QueueLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QueueLib *QueueLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QueueLib.Contract.QueueLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QueueLib *QueueLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QueueLib.Contract.QueueLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QueueLib *QueueLibCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _QueueLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QueueLib *QueueLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QueueLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QueueLib *QueueLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QueueLib.Contract.contract.Transact(opts, method, params...)
}

