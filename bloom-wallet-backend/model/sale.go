// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package model

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

// C2NSaleMetaData contains all meta data concerning the C2NSale contract.
var C2NSaleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_allocationStaking\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxParticipation\",\"type\":\"uint256\"}],\"name\":\"MaxParticipationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"registrationTimeStarts\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"registrationTimeEnds\",\"type\":\"uint256\"}],\"name\":\"RegistrationTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"saleOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenPriceInETH\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOfTokensToSell\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"saleEnd\",\"type\":\"uint256\"}],\"name\":\"SaleCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"name\":\"StartTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"TokenPriceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensSold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"UserRegistered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"contractIAdmin\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allocationStakingContract\",\"outputs\":[{\"internalType\":\"contractIAllocationStaking\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"checkParticipationSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"checkRegistrationSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeToAdd\",\"type\":\"uint256\"}],\"name\":\"extendRegistrationPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"contractISalesFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumberOfRegisteredUsers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getParticipation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getParticipationSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVestingInfo\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isParticipated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxVestingTimeShift\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numberOfParticipants\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"participate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"portionVestingPrecision\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeToShift\",\"type\":\"uint256\"}],\"name\":\"postponeSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"}],\"name\":\"registerForSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"registrationTimeStarts\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"registrationTimeEnds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numberOfRegistrants\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sale\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isCreated\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"earningsWithdrawn\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"leftoverWithdrawn\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"tokensDeposited\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"saleOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenPriceInETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOfTokensToSell\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalTokensSold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalETHRaised\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"saleStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"saleEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensUnlockTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxParticipation\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"setCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_registrationTimeStarts\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_registrationTimeEnds\",\"type\":\"uint256\"}],\"name\":\"setRegistrationTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_saleOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenPriceInETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountOfTokensToSell\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_saleEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokensUnlockTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_portionVestingPrecision\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxParticipation\",\"type\":\"uint256\"}],\"name\":\"setSaleParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starTime\",\"type\":\"uint256\"}],\"name\":\"setSaleStart\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"saleToken\",\"type\":\"address\"}],\"name\":\"setSaleToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_unlockingTimes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_percents\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_maxVestingTimeShift\",\"type\":\"uint256\"}],\"name\":\"setVestingParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeToShift\",\"type\":\"uint256\"}],\"name\":\"shiftVestingUnlockingTimes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"updateTokenPriceInETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userToParticipation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountBought\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeParticipated\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vestingPercentPerPortion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vestingPortionsUnlockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawEarnings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawEarningsAndLeftover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawLeftover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"portionIds\",\"type\":\"uint256[]\"}],\"name\":\"withdrawMultiplePortions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"portionId\",\"type\":\"uint256\"}],\"name\":\"withdrawTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// C2NSaleABI is the input ABI used to generate the binding from.
// Deprecated: Use C2NSaleMetaData.ABI instead.
var C2NSaleABI = C2NSaleMetaData.ABI

// C2NSale is an auto generated Go binding around an Ethereum contract.
type C2NSale struct {
	C2NSaleCaller     // Read-only binding to the contract
	C2NSaleTransactor // Write-only binding to the contract
	C2NSaleFilterer   // Log filterer for contract events
}

// C2NSaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type C2NSaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// C2NSaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type C2NSaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// C2NSaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type C2NSaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// C2NSaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type C2NSaleSession struct {
	Contract     *C2NSale          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// C2NSaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type C2NSaleCallerSession struct {
	Contract *C2NSaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// C2NSaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type C2NSaleTransactorSession struct {
	Contract     *C2NSaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// C2NSaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type C2NSaleRaw struct {
	Contract *C2NSale // Generic contract binding to access the raw methods on
}

// C2NSaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type C2NSaleCallerRaw struct {
	Contract *C2NSaleCaller // Generic read-only contract binding to access the raw methods on
}

// C2NSaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type C2NSaleTransactorRaw struct {
	Contract *C2NSaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewC2NSale creates a new instance of C2NSale, bound to a specific deployed contract.
func NewC2NSale(address common.Address, backend bind.ContractBackend) (*C2NSale, error) {
	contract, err := bindC2NSale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &C2NSale{C2NSaleCaller: C2NSaleCaller{contract: contract}, C2NSaleTransactor: C2NSaleTransactor{contract: contract}, C2NSaleFilterer: C2NSaleFilterer{contract: contract}}, nil
}

// NewC2NSaleCaller creates a new read-only instance of C2NSale, bound to a specific deployed contract.
func NewC2NSaleCaller(address common.Address, caller bind.ContractCaller) (*C2NSaleCaller, error) {
	contract, err := bindC2NSale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &C2NSaleCaller{contract: contract}, nil
}

// NewC2NSaleTransactor creates a new write-only instance of C2NSale, bound to a specific deployed contract.
func NewC2NSaleTransactor(address common.Address, transactor bind.ContractTransactor) (*C2NSaleTransactor, error) {
	contract, err := bindC2NSale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &C2NSaleTransactor{contract: contract}, nil
}

// NewC2NSaleFilterer creates a new log filterer instance of C2NSale, bound to a specific deployed contract.
func NewC2NSaleFilterer(address common.Address, filterer bind.ContractFilterer) (*C2NSaleFilterer, error) {
	contract, err := bindC2NSale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &C2NSaleFilterer{contract: contract}, nil
}

// bindC2NSale binds a generic wrapper to an already deployed contract.
func bindC2NSale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := C2NSaleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_C2NSale *C2NSaleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _C2NSale.Contract.C2NSaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_C2NSale *C2NSaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _C2NSale.Contract.C2NSaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_C2NSale *C2NSaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _C2NSale.Contract.C2NSaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_C2NSale *C2NSaleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _C2NSale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_C2NSale *C2NSaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _C2NSale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_C2NSale *C2NSaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _C2NSale.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_C2NSale *C2NSaleCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _C2NSale.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_C2NSale *C2NSaleSession) Admin() (common.Address, error) {
	return _C2NSale.Contract.Admin(&_C2NSale.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_C2NSale *C2NSaleCallerSession) Admin() (common.Address, error) {
	return _C2NSale.Contract.Admin(&_C2NSale.CallOpts)
}

// AllocationStakingContract is a free data retrieval call binding the contract method 0x059ea172.
//
// Solidity: function allocationStakingContract() view returns(address)
func (_C2NSale *C2NSaleCaller) AllocationStakingContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _C2NSale.contract.Call(opts, &out, "allocationStakingContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationStakingContract is a free data retrieval call binding the contract method 0x059ea172.
//
// Solidity: function allocationStakingContract() view returns(address)
func (_C2NSale *C2NSaleSession) AllocationStakingContract() (common.Address, error) {
	return _C2NSale.Contract.AllocationStakingContract(&_C2NSale.CallOpts)
}

// AllocationStakingContract is a free data retrieval call binding the contract method 0x059ea172.
//
// Solidity: function allocationStakingContract() view returns(address)
func (_C2NSale *C2NSaleCallerSession) AllocationStakingContract() (common.Address, error) {
	return _C2NSale.Contract.AllocationStakingContract(&_C2NSale.CallOpts)
}

// CheckParticipationSignature is a free data retrieval call binding the contract method 0x482a5dc7.
//
// Solidity: function checkParticipationSignature(bytes signature, address user, uint256 amount) view returns(bool)
func (_C2NSale *C2NSaleCaller) CheckParticipationSignature(opts *bind.CallOpts, signature []byte, user common.Address, amount *big.Int) (bool, error) {
	var out []interface{}
	err := _C2NSale.contract.Call(opts, &out, "checkParticipationSignature", signature, user, amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckParticipationSignature is a free data retrieval call binding the contract method 0x482a5dc7.
//
// Solidity: function checkParticipationSignature(bytes signature, address user, uint256 amount) view returns(bool)
func (_C2NSale *C2NSaleSession) CheckParticipationSignature(signature []byte, user common.Address, amount *big.Int) (bool, error) {
	return _C2NSale.Contract.CheckParticipationSignature(&_C2NSale.CallOpts, signature, user, amount)
}

// CheckParticipationSignature is a free data retrieval call binding the contract method 0x482a5dc7.
//
// Solidity: function checkParticipationSignature(bytes signature, address user, uint256 amount) view returns(bool)
func (_C2NSale *C2NSaleCallerSession) CheckParticipationSignature(signature []byte, user common.Address, amount *big.Int) (bool, error) {
	return _C2NSale.Contract.CheckParticipationSignature(&_C2NSale.CallOpts, signature, user, amount)
}

// CheckRegistrationSignature is a free data retrieval call binding the contract method 0x735deffa.
//
// Solidity: function checkRegistrationSignature(bytes signature, address user) view returns(bool)
func (_C2NSale *C2NSaleCaller) CheckRegistrationSignature(opts *bind.CallOpts, signature []byte, user common.Address) (bool, error) {
	var out []interface{}
	err := _C2NSale.contract.Call(opts, &out, "checkRegistrationSignature", signature, user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckRegistrationSignature is a free data retrieval call binding the contract method 0x735deffa.
//
// Solidity: function checkRegistrationSignature(bytes signature, address user) view returns(bool)
func (_C2NSale *C2NSaleSession) CheckRegistrationSignature(signature []byte, user common.Address) (bool, error) {
	return _C2NSale.Contract.CheckRegistrationSignature(&_C2NSale.CallOpts, signature, user)
}

// CheckRegistrationSignature is a free data retrieval call binding the contract method 0x735deffa.
//
// Solidity: function checkRegistrationSignature(bytes signature, address user) view returns(bool)
func (_C2NSale *C2NSaleCallerSession) CheckRegistrationSignature(signature []byte, user common.Address) (bool, error) {
	return _C2NSale.Contract.CheckRegistrationSignature(&_C2NSale.CallOpts, signature, user)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_C2NSale *C2NSaleCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _C2NSale.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_C2NSale *C2NSaleSession) Factory() (common.Address, error) {
	return _C2NSale.Contract.Factory(&_C2NSale.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_C2NSale *C2NSaleCallerSession) Factory() (common.Address, error) {
	return _C2NSale.Contract.Factory(&_C2NSale.CallOpts)
}

// GetNumberOfRegisteredUsers is a free data retrieval call binding the contract method 0xab7589b5.
//
// Solidity: function getNumberOfRegisteredUsers() view returns(uint256)
func (_C2NSale *C2NSaleCaller) GetNumberOfRegisteredUsers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _C2NSale.contract.Call(opts, &out, "getNumberOfRegisteredUsers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberOfRegisteredUsers is a free data retrieval call binding the contract method 0xab7589b5.
//
// Solidity: function getNumberOfRegisteredUsers() view returns(uint256)
func (_C2NSale *C2NSaleSession) GetNumberOfRegisteredUsers() (*big.Int, error) {
	return _C2NSale.Contract.GetNumberOfRegisteredUsers(&_C2NSale.CallOpts)
}

// GetNumberOfRegisteredUsers is a free data retrieval call binding the contract method 0xab7589b5.
//
// Solidity: function getNumberOfRegisteredUsers() view returns(uint256)
func (_C2NSale *C2NSaleCallerSession) GetNumberOfRegisteredUsers() (*big.Int, error) {
	return _C2NSale.Contract.GetNumberOfRegisteredUsers(&_C2NSale.CallOpts)
}

// GetParticipation is a free data retrieval call binding the contract method 0xcad925ef.
//
// Solidity: function getParticipation(address _user) view returns(uint256, uint256, uint256, bool[])
func (_C2NSale *C2NSaleCaller) GetParticipation(opts *bind.CallOpts, _user common.Address) (*big.Int, *big.Int, *big.Int, []bool, error) {
	var out []interface{}
	err := _C2NSale.contract.Call(opts, &out, "getParticipation", _user)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new([]bool)).(*[]bool)

	return out0, out1, out2, out3, err

}

// GetParticipation is a free data retrieval call binding the contract method 0xcad925ef.
//
// Solidity: function getParticipation(address _user) view returns(uint256, uint256, uint256, bool[])
func (_C2NSale *C2NSaleSession) GetParticipation(_user common.Address) (*big.Int, *big.Int, *big.Int, []bool, error) {
	return _C2NSale.Contract.GetParticipation(&_C2NSale.CallOpts, _user)
}

// GetParticipation is a free data retrieval call binding the contract method 0xcad925ef.
//
// Solidity: function getParticipation(address _user) view returns(uint256, uint256, uint256, bool[])
func (_C2NSale *C2NSaleCallerSession) GetParticipation(_user common.Address) (*big.Int, *big.Int, *big.Int, []bool, error) {
	return _C2NSale.Contract.GetParticipation(&_C2NSale.CallOpts, _user)
}

// GetParticipationSigner is a free data retrieval call binding the contract method 0x4897f4c3.
//
// Solidity: function getParticipationSigner(bytes signature, address user, uint256 amount) view returns(address)
func (_C2NSale *C2NSaleCaller) GetParticipationSigner(opts *bind.CallOpts, signature []byte, user common.Address, amount *big.Int) (common.Address, error) {
	var out []interface{}
	err := _C2NSale.contract.Call(opts, &out, "getParticipationSigner", signature, user, amount)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetParticipationSigner is a free data retrieval call binding the contract method 0x4897f4c3.
//
// Solidity: function getParticipationSigner(bytes signature, address user, uint256 amount) view returns(address)
func (_C2NSale *C2NSaleSession) GetParticipationSigner(signature []byte, user common.Address, amount *big.Int) (common.Address, error) {
	return _C2NSale.Contract.GetParticipationSigner(&_C2NSale.CallOpts, signature, user, amount)
}

// GetParticipationSigner is a free data retrieval call binding the contract method 0x4897f4c3.
//
// Solidity: function getParticipationSigner(bytes signature, address user, uint256 amount) view returns(address)
func (_C2NSale *C2NSaleCallerSession) GetParticipationSigner(signature []byte, user common.Address, amount *big.Int) (common.Address, error) {
	return _C2NSale.Contract.GetParticipationSigner(&_C2NSale.CallOpts, signature, user, amount)
}

// GetVestingInfo is a free data retrieval call binding the contract method 0xdc25a300.
//
// Solidity: function getVestingInfo() view returns(uint256[], uint256[])
func (_C2NSale *C2NSaleCaller) GetVestingInfo(opts *bind.CallOpts) ([]*big.Int, []*big.Int, error) {
	var out []interface{}
	err := _C2NSale.contract.Call(opts, &out, "getVestingInfo")

	if err != nil {
		return *new([]*big.Int), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetVestingInfo is a free data retrieval call binding the contract method 0xdc25a300.
//
// Solidity: function getVestingInfo() view returns(uint256[], uint256[])
func (_C2NSale *C2NSaleSession) GetVestingInfo() ([]*big.Int, []*big.Int, error) {
	return _C2NSale.Contract.GetVestingInfo(&_C2NSale.CallOpts)
}

// GetVestingInfo is a free data retrieval call binding the contract method 0xdc25a300.
//
// Solidity: function getVestingInfo() view returns(uint256[], uint256[])
func (_C2NSale *C2NSaleCallerSession) GetVestingInfo() ([]*big.Int, []*big.Int, error) {
	return _C2NSale.Contract.GetVestingInfo(&_C2NSale.CallOpts)
}

// IsParticipated is a free data retrieval call binding the contract method 0xcf5b8d4b.
//
// Solidity: function isParticipated(address ) view returns(bool)
func (_C2NSale *C2NSaleCaller) IsParticipated(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _C2NSale.contract.Call(opts, &out, "isParticipated", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsParticipated is a free data retrieval call binding the contract method 0xcf5b8d4b.
//
// Solidity: function isParticipated(address ) view returns(bool)
func (_C2NSale *C2NSaleSession) IsParticipated(arg0 common.Address) (bool, error) {
	return _C2NSale.Contract.IsParticipated(&_C2NSale.CallOpts, arg0)
}

// IsParticipated is a free data retrieval call binding the contract method 0xcf5b8d4b.
//
// Solidity: function isParticipated(address ) view returns(bool)
func (_C2NSale *C2NSaleCallerSession) IsParticipated(arg0 common.Address) (bool, error) {
	return _C2NSale.Contract.IsParticipated(&_C2NSale.CallOpts, arg0)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address ) view returns(bool)
func (_C2NSale *C2NSaleCaller) IsRegistered(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _C2NSale.contract.Call(opts, &out, "isRegistered", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address ) view returns(bool)
func (_C2NSale *C2NSaleSession) IsRegistered(arg0 common.Address) (bool, error) {
	return _C2NSale.Contract.IsRegistered(&_C2NSale.CallOpts, arg0)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address ) view returns(bool)
func (_C2NSale *C2NSaleCallerSession) IsRegistered(arg0 common.Address) (bool, error) {
	return _C2NSale.Contract.IsRegistered(&_C2NSale.CallOpts, arg0)
}

// MaxVestingTimeShift is a free data retrieval call binding the contract method 0xccc171f5.
//
// Solidity: function maxVestingTimeShift() view returns(uint256)
func (_C2NSale *C2NSaleCaller) MaxVestingTimeShift(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _C2NSale.contract.Call(opts, &out, "maxVestingTimeShift")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxVestingTimeShift is a free data retrieval call binding the contract method 0xccc171f5.
//
// Solidity: function maxVestingTimeShift() view returns(uint256)
func (_C2NSale *C2NSaleSession) MaxVestingTimeShift() (*big.Int, error) {
	return _C2NSale.Contract.MaxVestingTimeShift(&_C2NSale.CallOpts)
}

// MaxVestingTimeShift is a free data retrieval call binding the contract method 0xccc171f5.
//
// Solidity: function maxVestingTimeShift() view returns(uint256)
func (_C2NSale *C2NSaleCallerSession) MaxVestingTimeShift() (*big.Int, error) {
	return _C2NSale.Contract.MaxVestingTimeShift(&_C2NSale.CallOpts)
}

// NumberOfParticipants is a free data retrieval call binding the contract method 0x7417040e.
//
// Solidity: function numberOfParticipants() view returns(uint256)
func (_C2NSale *C2NSaleCaller) NumberOfParticipants(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _C2NSale.contract.Call(opts, &out, "numberOfParticipants")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberOfParticipants is a free data retrieval call binding the contract method 0x7417040e.
//
// Solidity: function numberOfParticipants() view returns(uint256)
func (_C2NSale *C2NSaleSession) NumberOfParticipants() (*big.Int, error) {
	return _C2NSale.Contract.NumberOfParticipants(&_C2NSale.CallOpts)
}

// NumberOfParticipants is a free data retrieval call binding the contract method 0x7417040e.
//
// Solidity: function numberOfParticipants() view returns(uint256)
func (_C2NSale *C2NSaleCallerSession) NumberOfParticipants() (*big.Int, error) {
	return _C2NSale.Contract.NumberOfParticipants(&_C2NSale.CallOpts)
}

// PortionVestingPrecision is a free data retrieval call binding the contract method 0x2a7c35de.
//
// Solidity: function portionVestingPrecision() view returns(uint256)
func (_C2NSale *C2NSaleCaller) PortionVestingPrecision(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _C2NSale.contract.Call(opts, &out, "portionVestingPrecision")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PortionVestingPrecision is a free data retrieval call binding the contract method 0x2a7c35de.
//
// Solidity: function portionVestingPrecision() view returns(uint256)
func (_C2NSale *C2NSaleSession) PortionVestingPrecision() (*big.Int, error) {
	return _C2NSale.Contract.PortionVestingPrecision(&_C2NSale.CallOpts)
}

// PortionVestingPrecision is a free data retrieval call binding the contract method 0x2a7c35de.
//
// Solidity: function portionVestingPrecision() view returns(uint256)
func (_C2NSale *C2NSaleCallerSession) PortionVestingPrecision() (*big.Int, error) {
	return _C2NSale.Contract.PortionVestingPrecision(&_C2NSale.CallOpts)
}

// Registration is a free data retrieval call binding the contract method 0x443bd1d0.
//
// Solidity: function registration() view returns(uint256 registrationTimeStarts, uint256 registrationTimeEnds, uint256 numberOfRegistrants)
func (_C2NSale *C2NSaleCaller) Registration(opts *bind.CallOpts) (struct {
	RegistrationTimeStarts *big.Int
	RegistrationTimeEnds   *big.Int
	NumberOfRegistrants    *big.Int
}, error) {
	var out []interface{}
	err := _C2NSale.contract.Call(opts, &out, "registration")

	outstruct := new(struct {
		RegistrationTimeStarts *big.Int
		RegistrationTimeEnds   *big.Int
		NumberOfRegistrants    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RegistrationTimeStarts = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RegistrationTimeEnds = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NumberOfRegistrants = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Registration is a free data retrieval call binding the contract method 0x443bd1d0.
//
// Solidity: function registration() view returns(uint256 registrationTimeStarts, uint256 registrationTimeEnds, uint256 numberOfRegistrants)
func (_C2NSale *C2NSaleSession) Registration() (struct {
	RegistrationTimeStarts *big.Int
	RegistrationTimeEnds   *big.Int
	NumberOfRegistrants    *big.Int
}, error) {
	return _C2NSale.Contract.Registration(&_C2NSale.CallOpts)
}

// Registration is a free data retrieval call binding the contract method 0x443bd1d0.
//
// Solidity: function registration() view returns(uint256 registrationTimeStarts, uint256 registrationTimeEnds, uint256 numberOfRegistrants)
func (_C2NSale *C2NSaleCallerSession) Registration() (struct {
	RegistrationTimeStarts *big.Int
	RegistrationTimeEnds   *big.Int
	NumberOfRegistrants    *big.Int
}, error) {
	return _C2NSale.Contract.Registration(&_C2NSale.CallOpts)
}

// Sale is a free data retrieval call binding the contract method 0x6ad1fe02.
//
// Solidity: function sale() view returns(address token, bool isCreated, bool earningsWithdrawn, bool leftoverWithdrawn, bool tokensDeposited, address saleOwner, uint256 tokenPriceInETH, uint256 amountOfTokensToSell, uint256 totalTokensSold, uint256 totalETHRaised, uint256 saleStart, uint256 saleEnd, uint256 tokensUnlockTime, uint256 maxParticipation)
func (_C2NSale *C2NSaleCaller) Sale(opts *bind.CallOpts) (struct {
	Token                common.Address
	IsCreated            bool
	EarningsWithdrawn    bool
	LeftoverWithdrawn    bool
	TokensDeposited      bool
	SaleOwner            common.Address
	TokenPriceInETH      *big.Int
	AmountOfTokensToSell *big.Int
	TotalTokensSold      *big.Int
	TotalETHRaised       *big.Int
	SaleStart            *big.Int
	SaleEnd              *big.Int
	TokensUnlockTime     *big.Int
	MaxParticipation     *big.Int
}, error) {
	var out []interface{}
	err := _C2NSale.contract.Call(opts, &out, "sale")

	outstruct := new(struct {
		Token                common.Address
		IsCreated            bool
		EarningsWithdrawn    bool
		LeftoverWithdrawn    bool
		TokensDeposited      bool
		SaleOwner            common.Address
		TokenPriceInETH      *big.Int
		AmountOfTokensToSell *big.Int
		TotalTokensSold      *big.Int
		TotalETHRaised       *big.Int
		SaleStart            *big.Int
		SaleEnd              *big.Int
		TokensUnlockTime     *big.Int
		MaxParticipation     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.IsCreated = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.EarningsWithdrawn = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.LeftoverWithdrawn = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.TokensDeposited = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.SaleOwner = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.TokenPriceInETH = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.AmountOfTokensToSell = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.TotalTokensSold = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.TotalETHRaised = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.SaleStart = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.SaleEnd = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)
	outstruct.TokensUnlockTime = *abi.ConvertType(out[12], new(*big.Int)).(**big.Int)
	outstruct.MaxParticipation = *abi.ConvertType(out[13], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Sale is a free data retrieval call binding the contract method 0x6ad1fe02.
//
// Solidity: function sale() view returns(address token, bool isCreated, bool earningsWithdrawn, bool leftoverWithdrawn, bool tokensDeposited, address saleOwner, uint256 tokenPriceInETH, uint256 amountOfTokensToSell, uint256 totalTokensSold, uint256 totalETHRaised, uint256 saleStart, uint256 saleEnd, uint256 tokensUnlockTime, uint256 maxParticipation)
func (_C2NSale *C2NSaleSession) Sale() (struct {
	Token                common.Address
	IsCreated            bool
	EarningsWithdrawn    bool
	LeftoverWithdrawn    bool
	TokensDeposited      bool
	SaleOwner            common.Address
	TokenPriceInETH      *big.Int
	AmountOfTokensToSell *big.Int
	TotalTokensSold      *big.Int
	TotalETHRaised       *big.Int
	SaleStart            *big.Int
	SaleEnd              *big.Int
	TokensUnlockTime     *big.Int
	MaxParticipation     *big.Int
}, error) {
	return _C2NSale.Contract.Sale(&_C2NSale.CallOpts)
}

// Sale is a free data retrieval call binding the contract method 0x6ad1fe02.
//
// Solidity: function sale() view returns(address token, bool isCreated, bool earningsWithdrawn, bool leftoverWithdrawn, bool tokensDeposited, address saleOwner, uint256 tokenPriceInETH, uint256 amountOfTokensToSell, uint256 totalTokensSold, uint256 totalETHRaised, uint256 saleStart, uint256 saleEnd, uint256 tokensUnlockTime, uint256 maxParticipation)
func (_C2NSale *C2NSaleCallerSession) Sale() (struct {
	Token                common.Address
	IsCreated            bool
	EarningsWithdrawn    bool
	LeftoverWithdrawn    bool
	TokensDeposited      bool
	SaleOwner            common.Address
	TokenPriceInETH      *big.Int
	AmountOfTokensToSell *big.Int
	TotalTokensSold      *big.Int
	TotalETHRaised       *big.Int
	SaleStart            *big.Int
	SaleEnd              *big.Int
	TokensUnlockTime     *big.Int
	MaxParticipation     *big.Int
}, error) {
	return _C2NSale.Contract.Sale(&_C2NSale.CallOpts)
}

// UserToParticipation is a free data retrieval call binding the contract method 0x5e7464f6.
//
// Solidity: function userToParticipation(address ) view returns(uint256 amountBought, uint256 amountETHPaid, uint256 timeParticipated)
func (_C2NSale *C2NSaleCaller) UserToParticipation(opts *bind.CallOpts, arg0 common.Address) (struct {
	AmountBought     *big.Int
	AmountETHPaid    *big.Int
	TimeParticipated *big.Int
}, error) {
	var out []interface{}
	err := _C2NSale.contract.Call(opts, &out, "userToParticipation", arg0)

	outstruct := new(struct {
		AmountBought     *big.Int
		AmountETHPaid    *big.Int
		TimeParticipated *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountBought = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountETHPaid = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TimeParticipated = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserToParticipation is a free data retrieval call binding the contract method 0x5e7464f6.
//
// Solidity: function userToParticipation(address ) view returns(uint256 amountBought, uint256 amountETHPaid, uint256 timeParticipated)
func (_C2NSale *C2NSaleSession) UserToParticipation(arg0 common.Address) (struct {
	AmountBought     *big.Int
	AmountETHPaid    *big.Int
	TimeParticipated *big.Int
}, error) {
	return _C2NSale.Contract.UserToParticipation(&_C2NSale.CallOpts, arg0)
}

// UserToParticipation is a free data retrieval call binding the contract method 0x5e7464f6.
//
// Solidity: function userToParticipation(address ) view returns(uint256 amountBought, uint256 amountETHPaid, uint256 timeParticipated)
func (_C2NSale *C2NSaleCallerSession) UserToParticipation(arg0 common.Address) (struct {
	AmountBought     *big.Int
	AmountETHPaid    *big.Int
	TimeParticipated *big.Int
}, error) {
	return _C2NSale.Contract.UserToParticipation(&_C2NSale.CallOpts, arg0)
}

// VestingPercentPerPortion is a free data retrieval call binding the contract method 0x927f6aee.
//
// Solidity: function vestingPercentPerPortion(uint256 ) view returns(uint256)
func (_C2NSale *C2NSaleCaller) VestingPercentPerPortion(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _C2NSale.contract.Call(opts, &out, "vestingPercentPerPortion", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VestingPercentPerPortion is a free data retrieval call binding the contract method 0x927f6aee.
//
// Solidity: function vestingPercentPerPortion(uint256 ) view returns(uint256)
func (_C2NSale *C2NSaleSession) VestingPercentPerPortion(arg0 *big.Int) (*big.Int, error) {
	return _C2NSale.Contract.VestingPercentPerPortion(&_C2NSale.CallOpts, arg0)
}

// VestingPercentPerPortion is a free data retrieval call binding the contract method 0x927f6aee.
//
// Solidity: function vestingPercentPerPortion(uint256 ) view returns(uint256)
func (_C2NSale *C2NSaleCallerSession) VestingPercentPerPortion(arg0 *big.Int) (*big.Int, error) {
	return _C2NSale.Contract.VestingPercentPerPortion(&_C2NSale.CallOpts, arg0)
}

// VestingPortionsUnlockTime is a free data retrieval call binding the contract method 0xf1ef7ff2.
//
// Solidity: function vestingPortionsUnlockTime(uint256 ) view returns(uint256)
func (_C2NSale *C2NSaleCaller) VestingPortionsUnlockTime(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _C2NSale.contract.Call(opts, &out, "vestingPortionsUnlockTime", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VestingPortionsUnlockTime is a free data retrieval call binding the contract method 0xf1ef7ff2.
//
// Solidity: function vestingPortionsUnlockTime(uint256 ) view returns(uint256)
func (_C2NSale *C2NSaleSession) VestingPortionsUnlockTime(arg0 *big.Int) (*big.Int, error) {
	return _C2NSale.Contract.VestingPortionsUnlockTime(&_C2NSale.CallOpts, arg0)
}

// VestingPortionsUnlockTime is a free data retrieval call binding the contract method 0xf1ef7ff2.
//
// Solidity: function vestingPortionsUnlockTime(uint256 ) view returns(uint256)
func (_C2NSale *C2NSaleCallerSession) VestingPortionsUnlockTime(arg0 *big.Int) (*big.Int, error) {
	return _C2NSale.Contract.VestingPortionsUnlockTime(&_C2NSale.CallOpts, arg0)
}

// DepositTokens is a paid mutator transaction binding the contract method 0x7c4b414d.
//
// Solidity: function depositTokens() returns()
func (_C2NSale *C2NSaleTransactor) DepositTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _C2NSale.contract.Transact(opts, "depositTokens")
}

// DepositTokens is a paid mutator transaction binding the contract method 0x7c4b414d.
//
// Solidity: function depositTokens() returns()
func (_C2NSale *C2NSaleSession) DepositTokens() (*types.Transaction, error) {
	return _C2NSale.Contract.DepositTokens(&_C2NSale.TransactOpts)
}

// DepositTokens is a paid mutator transaction binding the contract method 0x7c4b414d.
//
// Solidity: function depositTokens() returns()
func (_C2NSale *C2NSaleTransactorSession) DepositTokens() (*types.Transaction, error) {
	return _C2NSale.Contract.DepositTokens(&_C2NSale.TransactOpts)
}

// ExtendRegistrationPeriod is a paid mutator transaction binding the contract method 0x5164d871.
//
// Solidity: function extendRegistrationPeriod(uint256 timeToAdd) returns()
func (_C2NSale *C2NSaleTransactor) ExtendRegistrationPeriod(opts *bind.TransactOpts, timeToAdd *big.Int) (*types.Transaction, error) {
	return _C2NSale.contract.Transact(opts, "extendRegistrationPeriod", timeToAdd)
}

// ExtendRegistrationPeriod is a paid mutator transaction binding the contract method 0x5164d871.
//
// Solidity: function extendRegistrationPeriod(uint256 timeToAdd) returns()
func (_C2NSale *C2NSaleSession) ExtendRegistrationPeriod(timeToAdd *big.Int) (*types.Transaction, error) {
	return _C2NSale.Contract.ExtendRegistrationPeriod(&_C2NSale.TransactOpts, timeToAdd)
}

// ExtendRegistrationPeriod is a paid mutator transaction binding the contract method 0x5164d871.
//
// Solidity: function extendRegistrationPeriod(uint256 timeToAdd) returns()
func (_C2NSale *C2NSaleTransactorSession) ExtendRegistrationPeriod(timeToAdd *big.Int) (*types.Transaction, error) {
	return _C2NSale.Contract.ExtendRegistrationPeriod(&_C2NSale.TransactOpts, timeToAdd)
}

// Participate is a paid mutator transaction binding the contract method 0x931d81c9.
//
// Solidity: function participate(bytes signature, uint256 amount) payable returns()
func (_C2NSale *C2NSaleTransactor) Participate(opts *bind.TransactOpts, signature []byte, amount *big.Int) (*types.Transaction, error) {
	return _C2NSale.contract.Transact(opts, "participate", signature, amount)
}

// Participate is a paid mutator transaction binding the contract method 0x931d81c9.
//
// Solidity: function participate(bytes signature, uint256 amount) payable returns()
func (_C2NSale *C2NSaleSession) Participate(signature []byte, amount *big.Int) (*types.Transaction, error) {
	return _C2NSale.Contract.Participate(&_C2NSale.TransactOpts, signature, amount)
}

// Participate is a paid mutator transaction binding the contract method 0x931d81c9.
//
// Solidity: function participate(bytes signature, uint256 amount) payable returns()
func (_C2NSale *C2NSaleTransactorSession) Participate(signature []byte, amount *big.Int) (*types.Transaction, error) {
	return _C2NSale.Contract.Participate(&_C2NSale.TransactOpts, signature, amount)
}

// PostponeSale is a paid mutator transaction binding the contract method 0x1f11cb1e.
//
// Solidity: function postponeSale(uint256 timeToShift) returns()
func (_C2NSale *C2NSaleTransactor) PostponeSale(opts *bind.TransactOpts, timeToShift *big.Int) (*types.Transaction, error) {
	return _C2NSale.contract.Transact(opts, "postponeSale", timeToShift)
}

// PostponeSale is a paid mutator transaction binding the contract method 0x1f11cb1e.
//
// Solidity: function postponeSale(uint256 timeToShift) returns()
func (_C2NSale *C2NSaleSession) PostponeSale(timeToShift *big.Int) (*types.Transaction, error) {
	return _C2NSale.Contract.PostponeSale(&_C2NSale.TransactOpts, timeToShift)
}

// PostponeSale is a paid mutator transaction binding the contract method 0x1f11cb1e.
//
// Solidity: function postponeSale(uint256 timeToShift) returns()
func (_C2NSale *C2NSaleTransactorSession) PostponeSale(timeToShift *big.Int) (*types.Transaction, error) {
	return _C2NSale.Contract.PostponeSale(&_C2NSale.TransactOpts, timeToShift)
}

// RegisterForSale is a paid mutator transaction binding the contract method 0xe9d8479e.
//
// Solidity: function registerForSale(bytes signature, uint256 pid) returns()
func (_C2NSale *C2NSaleTransactor) RegisterForSale(opts *bind.TransactOpts, signature []byte, pid *big.Int) (*types.Transaction, error) {
	return _C2NSale.contract.Transact(opts, "registerForSale", signature, pid)
}

// RegisterForSale is a paid mutator transaction binding the contract method 0xe9d8479e.
//
// Solidity: function registerForSale(bytes signature, uint256 pid) returns()
func (_C2NSale *C2NSaleSession) RegisterForSale(signature []byte, pid *big.Int) (*types.Transaction, error) {
	return _C2NSale.Contract.RegisterForSale(&_C2NSale.TransactOpts, signature, pid)
}

// RegisterForSale is a paid mutator transaction binding the contract method 0xe9d8479e.
//
// Solidity: function registerForSale(bytes signature, uint256 pid) returns()
func (_C2NSale *C2NSaleTransactorSession) RegisterForSale(signature []byte, pid *big.Int) (*types.Transaction, error) {
	return _C2NSale.Contract.RegisterForSale(&_C2NSale.TransactOpts, signature, pid)
}

// SetCap is a paid mutator transaction binding the contract method 0x47786d37.
//
// Solidity: function setCap(uint256 cap) returns()
func (_C2NSale *C2NSaleTransactor) SetCap(opts *bind.TransactOpts, cap *big.Int) (*types.Transaction, error) {
	return _C2NSale.contract.Transact(opts, "setCap", cap)
}

// SetCap is a paid mutator transaction binding the contract method 0x47786d37.
//
// Solidity: function setCap(uint256 cap) returns()
func (_C2NSale *C2NSaleSession) SetCap(cap *big.Int) (*types.Transaction, error) {
	return _C2NSale.Contract.SetCap(&_C2NSale.TransactOpts, cap)
}

// SetCap is a paid mutator transaction binding the contract method 0x47786d37.
//
// Solidity: function setCap(uint256 cap) returns()
func (_C2NSale *C2NSaleTransactorSession) SetCap(cap *big.Int) (*types.Transaction, error) {
	return _C2NSale.Contract.SetCap(&_C2NSale.TransactOpts, cap)
}

// SetRegistrationTime is a paid mutator transaction binding the contract method 0xe099cf64.
//
// Solidity: function setRegistrationTime(uint256 _registrationTimeStarts, uint256 _registrationTimeEnds) returns()
func (_C2NSale *C2NSaleTransactor) SetRegistrationTime(opts *bind.TransactOpts, _registrationTimeStarts *big.Int, _registrationTimeEnds *big.Int) (*types.Transaction, error) {
	return _C2NSale.contract.Transact(opts, "setRegistrationTime", _registrationTimeStarts, _registrationTimeEnds)
}

// SetRegistrationTime is a paid mutator transaction binding the contract method 0xe099cf64.
//
// Solidity: function setRegistrationTime(uint256 _registrationTimeStarts, uint256 _registrationTimeEnds) returns()
func (_C2NSale *C2NSaleSession) SetRegistrationTime(_registrationTimeStarts *big.Int, _registrationTimeEnds *big.Int) (*types.Transaction, error) {
	return _C2NSale.Contract.SetRegistrationTime(&_C2NSale.TransactOpts, _registrationTimeStarts, _registrationTimeEnds)
}

// SetRegistrationTime is a paid mutator transaction binding the contract method 0xe099cf64.
//
// Solidity: function setRegistrationTime(uint256 _registrationTimeStarts, uint256 _registrationTimeEnds) returns()
func (_C2NSale *C2NSaleTransactorSession) SetRegistrationTime(_registrationTimeStarts *big.Int, _registrationTimeEnds *big.Int) (*types.Transaction, error) {
	return _C2NSale.Contract.SetRegistrationTime(&_C2NSale.TransactOpts, _registrationTimeStarts, _registrationTimeEnds)
}

// SetSaleParams is a paid mutator transaction binding the contract method 0xc4fbe091.
//
// Solidity: function setSaleParams(address _token, address _saleOwner, uint256 _tokenPriceInETH, uint256 _amountOfTokensToSell, uint256 _saleEnd, uint256 _tokensUnlockTime, uint256 _portionVestingPrecision, uint256 _maxParticipation) returns()
func (_C2NSale *C2NSaleTransactor) SetSaleParams(opts *bind.TransactOpts, _token common.Address, _saleOwner common.Address, _tokenPriceInETH *big.Int, _amountOfTokensToSell *big.Int, _saleEnd *big.Int, _tokensUnlockTime *big.Int, _portionVestingPrecision *big.Int, _maxParticipation *big.Int) (*types.Transaction, error) {
	return _C2NSale.contract.Transact(opts, "setSaleParams", _token, _saleOwner, _tokenPriceInETH, _amountOfTokensToSell, _saleEnd, _tokensUnlockTime, _portionVestingPrecision, _maxParticipation)
}

// SetSaleParams is a paid mutator transaction binding the contract method 0xc4fbe091.
//
// Solidity: function setSaleParams(address _token, address _saleOwner, uint256 _tokenPriceInETH, uint256 _amountOfTokensToSell, uint256 _saleEnd, uint256 _tokensUnlockTime, uint256 _portionVestingPrecision, uint256 _maxParticipation) returns()
func (_C2NSale *C2NSaleSession) SetSaleParams(_token common.Address, _saleOwner common.Address, _tokenPriceInETH *big.Int, _amountOfTokensToSell *big.Int, _saleEnd *big.Int, _tokensUnlockTime *big.Int, _portionVestingPrecision *big.Int, _maxParticipation *big.Int) (*types.Transaction, error) {
	return _C2NSale.Contract.SetSaleParams(&_C2NSale.TransactOpts, _token, _saleOwner, _tokenPriceInETH, _amountOfTokensToSell, _saleEnd, _tokensUnlockTime, _portionVestingPrecision, _maxParticipation)
}

// SetSaleParams is a paid mutator transaction binding the contract method 0xc4fbe091.
//
// Solidity: function setSaleParams(address _token, address _saleOwner, uint256 _tokenPriceInETH, uint256 _amountOfTokensToSell, uint256 _saleEnd, uint256 _tokensUnlockTime, uint256 _portionVestingPrecision, uint256 _maxParticipation) returns()
func (_C2NSale *C2NSaleTransactorSession) SetSaleParams(_token common.Address, _saleOwner common.Address, _tokenPriceInETH *big.Int, _amountOfTokensToSell *big.Int, _saleEnd *big.Int, _tokensUnlockTime *big.Int, _portionVestingPrecision *big.Int, _maxParticipation *big.Int) (*types.Transaction, error) {
	return _C2NSale.Contract.SetSaleParams(&_C2NSale.TransactOpts, _token, _saleOwner, _tokenPriceInETH, _amountOfTokensToSell, _saleEnd, _tokensUnlockTime, _portionVestingPrecision, _maxParticipation)
}

// SetSaleStart is a paid mutator transaction binding the contract method 0x2f181f54.
//
// Solidity: function setSaleStart(uint256 starTime) returns()
func (_C2NSale *C2NSaleTransactor) SetSaleStart(opts *bind.TransactOpts, starTime *big.Int) (*types.Transaction, error) {
	return _C2NSale.contract.Transact(opts, "setSaleStart", starTime)
}

// SetSaleStart is a paid mutator transaction binding the contract method 0x2f181f54.
//
// Solidity: function setSaleStart(uint256 starTime) returns()
func (_C2NSale *C2NSaleSession) SetSaleStart(starTime *big.Int) (*types.Transaction, error) {
	return _C2NSale.Contract.SetSaleStart(&_C2NSale.TransactOpts, starTime)
}

// SetSaleStart is a paid mutator transaction binding the contract method 0x2f181f54.
//
// Solidity: function setSaleStart(uint256 starTime) returns()
func (_C2NSale *C2NSaleTransactorSession) SetSaleStart(starTime *big.Int) (*types.Transaction, error) {
	return _C2NSale.Contract.SetSaleStart(&_C2NSale.TransactOpts, starTime)
}

// SetSaleToken is a paid mutator transaction binding the contract method 0xa29f481c.
//
// Solidity: function setSaleToken(address saleToken) returns()
func (_C2NSale *C2NSaleTransactor) SetSaleToken(opts *bind.TransactOpts, saleToken common.Address) (*types.Transaction, error) {
	return _C2NSale.contract.Transact(opts, "setSaleToken", saleToken)
}

// SetSaleToken is a paid mutator transaction binding the contract method 0xa29f481c.
//
// Solidity: function setSaleToken(address saleToken) returns()
func (_C2NSale *C2NSaleSession) SetSaleToken(saleToken common.Address) (*types.Transaction, error) {
	return _C2NSale.Contract.SetSaleToken(&_C2NSale.TransactOpts, saleToken)
}

// SetSaleToken is a paid mutator transaction binding the contract method 0xa29f481c.
//
// Solidity: function setSaleToken(address saleToken) returns()
func (_C2NSale *C2NSaleTransactorSession) SetSaleToken(saleToken common.Address) (*types.Transaction, error) {
	return _C2NSale.Contract.SetSaleToken(&_C2NSale.TransactOpts, saleToken)
}

// SetVestingParams is a paid mutator transaction binding the contract method 0xd937d456.
//
// Solidity: function setVestingParams(uint256[] _unlockingTimes, uint256[] _percents, uint256 _maxVestingTimeShift) returns()
func (_C2NSale *C2NSaleTransactor) SetVestingParams(opts *bind.TransactOpts, _unlockingTimes []*big.Int, _percents []*big.Int, _maxVestingTimeShift *big.Int) (*types.Transaction, error) {
	return _C2NSale.contract.Transact(opts, "setVestingParams", _unlockingTimes, _percents, _maxVestingTimeShift)
}

// SetVestingParams is a paid mutator transaction binding the contract method 0xd937d456.
//
// Solidity: function setVestingParams(uint256[] _unlockingTimes, uint256[] _percents, uint256 _maxVestingTimeShift) returns()
func (_C2NSale *C2NSaleSession) SetVestingParams(_unlockingTimes []*big.Int, _percents []*big.Int, _maxVestingTimeShift *big.Int) (*types.Transaction, error) {
	return _C2NSale.Contract.SetVestingParams(&_C2NSale.TransactOpts, _unlockingTimes, _percents, _maxVestingTimeShift)
}

// SetVestingParams is a paid mutator transaction binding the contract method 0xd937d456.
//
// Solidity: function setVestingParams(uint256[] _unlockingTimes, uint256[] _percents, uint256 _maxVestingTimeShift) returns()
func (_C2NSale *C2NSaleTransactorSession) SetVestingParams(_unlockingTimes []*big.Int, _percents []*big.Int, _maxVestingTimeShift *big.Int) (*types.Transaction, error) {
	return _C2NSale.Contract.SetVestingParams(&_C2NSale.TransactOpts, _unlockingTimes, _percents, _maxVestingTimeShift)
}

// ShiftVestingUnlockingTimes is a paid mutator transaction binding the contract method 0xabee6927.
//
// Solidity: function shiftVestingUnlockingTimes(uint256 timeToShift) returns()
func (_C2NSale *C2NSaleTransactor) ShiftVestingUnlockingTimes(opts *bind.TransactOpts, timeToShift *big.Int) (*types.Transaction, error) {
	return _C2NSale.contract.Transact(opts, "shiftVestingUnlockingTimes", timeToShift)
}

// ShiftVestingUnlockingTimes is a paid mutator transaction binding the contract method 0xabee6927.
//
// Solidity: function shiftVestingUnlockingTimes(uint256 timeToShift) returns()
func (_C2NSale *C2NSaleSession) ShiftVestingUnlockingTimes(timeToShift *big.Int) (*types.Transaction, error) {
	return _C2NSale.Contract.ShiftVestingUnlockingTimes(&_C2NSale.TransactOpts, timeToShift)
}

// ShiftVestingUnlockingTimes is a paid mutator transaction binding the contract method 0xabee6927.
//
// Solidity: function shiftVestingUnlockingTimes(uint256 timeToShift) returns()
func (_C2NSale *C2NSaleTransactorSession) ShiftVestingUnlockingTimes(timeToShift *big.Int) (*types.Transaction, error) {
	return _C2NSale.Contract.ShiftVestingUnlockingTimes(&_C2NSale.TransactOpts, timeToShift)
}

// UpdateTokenPriceInETH is a paid mutator transaction binding the contract method 0x7843990a.
//
// Solidity: function updateTokenPriceInETH(uint256 price) returns()
func (_C2NSale *C2NSaleTransactor) UpdateTokenPriceInETH(opts *bind.TransactOpts, price *big.Int) (*types.Transaction, error) {
	return _C2NSale.contract.Transact(opts, "updateTokenPriceInETH", price)
}

// UpdateTokenPriceInETH is a paid mutator transaction binding the contract method 0x7843990a.
//
// Solidity: function updateTokenPriceInETH(uint256 price) returns()
func (_C2NSale *C2NSaleSession) UpdateTokenPriceInETH(price *big.Int) (*types.Transaction, error) {
	return _C2NSale.Contract.UpdateTokenPriceInETH(&_C2NSale.TransactOpts, price)
}

// UpdateTokenPriceInETH is a paid mutator transaction binding the contract method 0x7843990a.
//
// Solidity: function updateTokenPriceInETH(uint256 price) returns()
func (_C2NSale *C2NSaleTransactorSession) UpdateTokenPriceInETH(price *big.Int) (*types.Transaction, error) {
	return _C2NSale.Contract.UpdateTokenPriceInETH(&_C2NSale.TransactOpts, price)
}

// WithdrawEarnings is a paid mutator transaction binding the contract method 0xb73c6ce9.
//
// Solidity: function withdrawEarnings() returns()
func (_C2NSale *C2NSaleTransactor) WithdrawEarnings(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _C2NSale.contract.Transact(opts, "withdrawEarnings")
}

// WithdrawEarnings is a paid mutator transaction binding the contract method 0xb73c6ce9.
//
// Solidity: function withdrawEarnings() returns()
func (_C2NSale *C2NSaleSession) WithdrawEarnings() (*types.Transaction, error) {
	return _C2NSale.Contract.WithdrawEarnings(&_C2NSale.TransactOpts)
}

// WithdrawEarnings is a paid mutator transaction binding the contract method 0xb73c6ce9.
//
// Solidity: function withdrawEarnings() returns()
func (_C2NSale *C2NSaleTransactorSession) WithdrawEarnings() (*types.Transaction, error) {
	return _C2NSale.Contract.WithdrawEarnings(&_C2NSale.TransactOpts)
}

// WithdrawEarningsAndLeftover is a paid mutator transaction binding the contract method 0xda4d4fbf.
//
// Solidity: function withdrawEarningsAndLeftover() returns()
func (_C2NSale *C2NSaleTransactor) WithdrawEarningsAndLeftover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _C2NSale.contract.Transact(opts, "withdrawEarningsAndLeftover")
}

// WithdrawEarningsAndLeftover is a paid mutator transaction binding the contract method 0xda4d4fbf.
//
// Solidity: function withdrawEarningsAndLeftover() returns()
func (_C2NSale *C2NSaleSession) WithdrawEarningsAndLeftover() (*types.Transaction, error) {
	return _C2NSale.Contract.WithdrawEarningsAndLeftover(&_C2NSale.TransactOpts)
}

// WithdrawEarningsAndLeftover is a paid mutator transaction binding the contract method 0xda4d4fbf.
//
// Solidity: function withdrawEarningsAndLeftover() returns()
func (_C2NSale *C2NSaleTransactorSession) WithdrawEarningsAndLeftover() (*types.Transaction, error) {
	return _C2NSale.Contract.WithdrawEarningsAndLeftover(&_C2NSale.TransactOpts)
}

// WithdrawLeftover is a paid mutator transaction binding the contract method 0xa525d237.
//
// Solidity: function withdrawLeftover() returns()
func (_C2NSale *C2NSaleTransactor) WithdrawLeftover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _C2NSale.contract.Transact(opts, "withdrawLeftover")
}

// WithdrawLeftover is a paid mutator transaction binding the contract method 0xa525d237.
//
// Solidity: function withdrawLeftover() returns()
func (_C2NSale *C2NSaleSession) WithdrawLeftover() (*types.Transaction, error) {
	return _C2NSale.Contract.WithdrawLeftover(&_C2NSale.TransactOpts)
}

// WithdrawLeftover is a paid mutator transaction binding the contract method 0xa525d237.
//
// Solidity: function withdrawLeftover() returns()
func (_C2NSale *C2NSaleTransactorSession) WithdrawLeftover() (*types.Transaction, error) {
	return _C2NSale.Contract.WithdrawLeftover(&_C2NSale.TransactOpts)
}

// WithdrawMultiplePortions is a paid mutator transaction binding the contract method 0x718af7e6.
//
// Solidity: function withdrawMultiplePortions(uint256[] portionIds) returns()
func (_C2NSale *C2NSaleTransactor) WithdrawMultiplePortions(opts *bind.TransactOpts, portionIds []*big.Int) (*types.Transaction, error) {
	return _C2NSale.contract.Transact(opts, "withdrawMultiplePortions", portionIds)
}

// WithdrawMultiplePortions is a paid mutator transaction binding the contract method 0x718af7e6.
//
// Solidity: function withdrawMultiplePortions(uint256[] portionIds) returns()
func (_C2NSale *C2NSaleSession) WithdrawMultiplePortions(portionIds []*big.Int) (*types.Transaction, error) {
	return _C2NSale.Contract.WithdrawMultiplePortions(&_C2NSale.TransactOpts, portionIds)
}

// WithdrawMultiplePortions is a paid mutator transaction binding the contract method 0x718af7e6.
//
// Solidity: function withdrawMultiplePortions(uint256[] portionIds) returns()
func (_C2NSale *C2NSaleTransactorSession) WithdrawMultiplePortions(portionIds []*big.Int) (*types.Transaction, error) {
	return _C2NSale.Contract.WithdrawMultiplePortions(&_C2NSale.TransactOpts, portionIds)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x315a095d.
//
// Solidity: function withdrawTokens(uint256 portionId) returns()
func (_C2NSale *C2NSaleTransactor) WithdrawTokens(opts *bind.TransactOpts, portionId *big.Int) (*types.Transaction, error) {
	return _C2NSale.contract.Transact(opts, "withdrawTokens", portionId)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x315a095d.
//
// Solidity: function withdrawTokens(uint256 portionId) returns()
func (_C2NSale *C2NSaleSession) WithdrawTokens(portionId *big.Int) (*types.Transaction, error) {
	return _C2NSale.Contract.WithdrawTokens(&_C2NSale.TransactOpts, portionId)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x315a095d.
//
// Solidity: function withdrawTokens(uint256 portionId) returns()
func (_C2NSale *C2NSaleTransactorSession) WithdrawTokens(portionId *big.Int) (*types.Transaction, error) {
	return _C2NSale.Contract.WithdrawTokens(&_C2NSale.TransactOpts, portionId)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_C2NSale *C2NSaleTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _C2NSale.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_C2NSale *C2NSaleSession) Receive() (*types.Transaction, error) {
	return _C2NSale.Contract.Receive(&_C2NSale.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_C2NSale *C2NSaleTransactorSession) Receive() (*types.Transaction, error) {
	return _C2NSale.Contract.Receive(&_C2NSale.TransactOpts)
}

// C2NSaleMaxParticipationSetIterator is returned from FilterMaxParticipationSet and is used to iterate over the raw logs and unpacked data for MaxParticipationSet events raised by the C2NSale contract.
type C2NSaleMaxParticipationSetIterator struct {
	Event *C2NSaleMaxParticipationSet // Event containing the contract specifics and raw log

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
func (it *C2NSaleMaxParticipationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(C2NSaleMaxParticipationSet)
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
		it.Event = new(C2NSaleMaxParticipationSet)
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
func (it *C2NSaleMaxParticipationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *C2NSaleMaxParticipationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// C2NSaleMaxParticipationSet represents a MaxParticipationSet event raised by the C2NSale contract.
type C2NSaleMaxParticipationSet struct {
	MaxParticipation *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMaxParticipationSet is a free log retrieval operation binding the contract event 0x37cfb0744bdb923d0300a0c37fac13cfec5fc2ee313cb9b217d284122980bada.
//
// Solidity: event MaxParticipationSet(uint256 maxParticipation)
func (_C2NSale *C2NSaleFilterer) FilterMaxParticipationSet(opts *bind.FilterOpts) (*C2NSaleMaxParticipationSetIterator, error) {

	logs, sub, err := _C2NSale.contract.FilterLogs(opts, "MaxParticipationSet")
	if err != nil {
		return nil, err
	}
	return &C2NSaleMaxParticipationSetIterator{contract: _C2NSale.contract, event: "MaxParticipationSet", logs: logs, sub: sub}, nil
}

// WatchMaxParticipationSet is a free log subscription operation binding the contract event 0x37cfb0744bdb923d0300a0c37fac13cfec5fc2ee313cb9b217d284122980bada.
//
// Solidity: event MaxParticipationSet(uint256 maxParticipation)
func (_C2NSale *C2NSaleFilterer) WatchMaxParticipationSet(opts *bind.WatchOpts, sink chan<- *C2NSaleMaxParticipationSet) (event.Subscription, error) {

	logs, sub, err := _C2NSale.contract.WatchLogs(opts, "MaxParticipationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(C2NSaleMaxParticipationSet)
				if err := _C2NSale.contract.UnpackLog(event, "MaxParticipationSet", log); err != nil {
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

// ParseMaxParticipationSet is a log parse operation binding the contract event 0x37cfb0744bdb923d0300a0c37fac13cfec5fc2ee313cb9b217d284122980bada.
//
// Solidity: event MaxParticipationSet(uint256 maxParticipation)
func (_C2NSale *C2NSaleFilterer) ParseMaxParticipationSet(log types.Log) (*C2NSaleMaxParticipationSet, error) {
	event := new(C2NSaleMaxParticipationSet)
	if err := _C2NSale.contract.UnpackLog(event, "MaxParticipationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// C2NSaleRegistrationTimeSetIterator is returned from FilterRegistrationTimeSet and is used to iterate over the raw logs and unpacked data for RegistrationTimeSet events raised by the C2NSale contract.
type C2NSaleRegistrationTimeSetIterator struct {
	Event *C2NSaleRegistrationTimeSet // Event containing the contract specifics and raw log

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
func (it *C2NSaleRegistrationTimeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(C2NSaleRegistrationTimeSet)
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
		it.Event = new(C2NSaleRegistrationTimeSet)
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
func (it *C2NSaleRegistrationTimeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *C2NSaleRegistrationTimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// C2NSaleRegistrationTimeSet represents a RegistrationTimeSet event raised by the C2NSale contract.
type C2NSaleRegistrationTimeSet struct {
	RegistrationTimeStarts *big.Int
	RegistrationTimeEnds   *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterRegistrationTimeSet is a free log retrieval operation binding the contract event 0xaf6e91c17885891414abbe7fa95074976b3e429980b0d8790034468e69650dd9.
//
// Solidity: event RegistrationTimeSet(uint256 registrationTimeStarts, uint256 registrationTimeEnds)
func (_C2NSale *C2NSaleFilterer) FilterRegistrationTimeSet(opts *bind.FilterOpts) (*C2NSaleRegistrationTimeSetIterator, error) {

	logs, sub, err := _C2NSale.contract.FilterLogs(opts, "RegistrationTimeSet")
	if err != nil {
		return nil, err
	}
	return &C2NSaleRegistrationTimeSetIterator{contract: _C2NSale.contract, event: "RegistrationTimeSet", logs: logs, sub: sub}, nil
}

// WatchRegistrationTimeSet is a free log subscription operation binding the contract event 0xaf6e91c17885891414abbe7fa95074976b3e429980b0d8790034468e69650dd9.
//
// Solidity: event RegistrationTimeSet(uint256 registrationTimeStarts, uint256 registrationTimeEnds)
func (_C2NSale *C2NSaleFilterer) WatchRegistrationTimeSet(opts *bind.WatchOpts, sink chan<- *C2NSaleRegistrationTimeSet) (event.Subscription, error) {

	logs, sub, err := _C2NSale.contract.WatchLogs(opts, "RegistrationTimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(C2NSaleRegistrationTimeSet)
				if err := _C2NSale.contract.UnpackLog(event, "RegistrationTimeSet", log); err != nil {
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

// ParseRegistrationTimeSet is a log parse operation binding the contract event 0xaf6e91c17885891414abbe7fa95074976b3e429980b0d8790034468e69650dd9.
//
// Solidity: event RegistrationTimeSet(uint256 registrationTimeStarts, uint256 registrationTimeEnds)
func (_C2NSale *C2NSaleFilterer) ParseRegistrationTimeSet(log types.Log) (*C2NSaleRegistrationTimeSet, error) {
	event := new(C2NSaleRegistrationTimeSet)
	if err := _C2NSale.contract.UnpackLog(event, "RegistrationTimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// C2NSaleSaleCreatedIterator is returned from FilterSaleCreated and is used to iterate over the raw logs and unpacked data for SaleCreated events raised by the C2NSale contract.
type C2NSaleSaleCreatedIterator struct {
	Event *C2NSaleSaleCreated // Event containing the contract specifics and raw log

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
func (it *C2NSaleSaleCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(C2NSaleSaleCreated)
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
		it.Event = new(C2NSaleSaleCreated)
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
func (it *C2NSaleSaleCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *C2NSaleSaleCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// C2NSaleSaleCreated represents a SaleCreated event raised by the C2NSale contract.
type C2NSaleSaleCreated struct {
	SaleOwner            common.Address
	TokenPriceInETH      *big.Int
	AmountOfTokensToSell *big.Int
	SaleEnd              *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSaleCreated is a free log retrieval operation binding the contract event 0x2cd2dfcdeb2b58c4b80527e9df5e12da537fa4f6c958a4fb623a83ab74eeab63.
//
// Solidity: event SaleCreated(address saleOwner, uint256 tokenPriceInETH, uint256 amountOfTokensToSell, uint256 saleEnd)
func (_C2NSale *C2NSaleFilterer) FilterSaleCreated(opts *bind.FilterOpts) (*C2NSaleSaleCreatedIterator, error) {

	logs, sub, err := _C2NSale.contract.FilterLogs(opts, "SaleCreated")
	if err != nil {
		return nil, err
	}
	return &C2NSaleSaleCreatedIterator{contract: _C2NSale.contract, event: "SaleCreated", logs: logs, sub: sub}, nil
}

// WatchSaleCreated is a free log subscription operation binding the contract event 0x2cd2dfcdeb2b58c4b80527e9df5e12da537fa4f6c958a4fb623a83ab74eeab63.
//
// Solidity: event SaleCreated(address saleOwner, uint256 tokenPriceInETH, uint256 amountOfTokensToSell, uint256 saleEnd)
func (_C2NSale *C2NSaleFilterer) WatchSaleCreated(opts *bind.WatchOpts, sink chan<- *C2NSaleSaleCreated) (event.Subscription, error) {

	logs, sub, err := _C2NSale.contract.WatchLogs(opts, "SaleCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(C2NSaleSaleCreated)
				if err := _C2NSale.contract.UnpackLog(event, "SaleCreated", log); err != nil {
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

// ParseSaleCreated is a log parse operation binding the contract event 0x2cd2dfcdeb2b58c4b80527e9df5e12da537fa4f6c958a4fb623a83ab74eeab63.
//
// Solidity: event SaleCreated(address saleOwner, uint256 tokenPriceInETH, uint256 amountOfTokensToSell, uint256 saleEnd)
func (_C2NSale *C2NSaleFilterer) ParseSaleCreated(log types.Log) (*C2NSaleSaleCreated, error) {
	event := new(C2NSaleSaleCreated)
	if err := _C2NSale.contract.UnpackLog(event, "SaleCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// C2NSaleStartTimeSetIterator is returned from FilterStartTimeSet and is used to iterate over the raw logs and unpacked data for StartTimeSet events raised by the C2NSale contract.
type C2NSaleStartTimeSetIterator struct {
	Event *C2NSaleStartTimeSet // Event containing the contract specifics and raw log

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
func (it *C2NSaleStartTimeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(C2NSaleStartTimeSet)
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
		it.Event = new(C2NSaleStartTimeSet)
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
func (it *C2NSaleStartTimeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *C2NSaleStartTimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// C2NSaleStartTimeSet represents a StartTimeSet event raised by the C2NSale contract.
type C2NSaleStartTimeSet struct {
	StartTime *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStartTimeSet is a free log retrieval operation binding the contract event 0xaad53c4362ef2fe5a5390cc046e71fd8423a0a8dceebc156ac9bbcd15997eec2.
//
// Solidity: event StartTimeSet(uint256 startTime)
func (_C2NSale *C2NSaleFilterer) FilterStartTimeSet(opts *bind.FilterOpts) (*C2NSaleStartTimeSetIterator, error) {

	logs, sub, err := _C2NSale.contract.FilterLogs(opts, "StartTimeSet")
	if err != nil {
		return nil, err
	}
	return &C2NSaleStartTimeSetIterator{contract: _C2NSale.contract, event: "StartTimeSet", logs: logs, sub: sub}, nil
}

// WatchStartTimeSet is a free log subscription operation binding the contract event 0xaad53c4362ef2fe5a5390cc046e71fd8423a0a8dceebc156ac9bbcd15997eec2.
//
// Solidity: event StartTimeSet(uint256 startTime)
func (_C2NSale *C2NSaleFilterer) WatchStartTimeSet(opts *bind.WatchOpts, sink chan<- *C2NSaleStartTimeSet) (event.Subscription, error) {

	logs, sub, err := _C2NSale.contract.WatchLogs(opts, "StartTimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(C2NSaleStartTimeSet)
				if err := _C2NSale.contract.UnpackLog(event, "StartTimeSet", log); err != nil {
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

// ParseStartTimeSet is a log parse operation binding the contract event 0xaad53c4362ef2fe5a5390cc046e71fd8423a0a8dceebc156ac9bbcd15997eec2.
//
// Solidity: event StartTimeSet(uint256 startTime)
func (_C2NSale *C2NSaleFilterer) ParseStartTimeSet(log types.Log) (*C2NSaleStartTimeSet, error) {
	event := new(C2NSaleStartTimeSet)
	if err := _C2NSale.contract.UnpackLog(event, "StartTimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// C2NSaleTokenPriceSetIterator is returned from FilterTokenPriceSet and is used to iterate over the raw logs and unpacked data for TokenPriceSet events raised by the C2NSale contract.
type C2NSaleTokenPriceSetIterator struct {
	Event *C2NSaleTokenPriceSet // Event containing the contract specifics and raw log

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
func (it *C2NSaleTokenPriceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(C2NSaleTokenPriceSet)
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
		it.Event = new(C2NSaleTokenPriceSet)
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
func (it *C2NSaleTokenPriceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *C2NSaleTokenPriceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// C2NSaleTokenPriceSet represents a TokenPriceSet event raised by the C2NSale contract.
type C2NSaleTokenPriceSet struct {
	NewPrice *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTokenPriceSet is a free log retrieval operation binding the contract event 0x4b59d61d9ffdc3db926d0ce7e06ebabb6bd1bf9dcdae262667e48be368227216.
//
// Solidity: event TokenPriceSet(uint256 newPrice)
func (_C2NSale *C2NSaleFilterer) FilterTokenPriceSet(opts *bind.FilterOpts) (*C2NSaleTokenPriceSetIterator, error) {

	logs, sub, err := _C2NSale.contract.FilterLogs(opts, "TokenPriceSet")
	if err != nil {
		return nil, err
	}
	return &C2NSaleTokenPriceSetIterator{contract: _C2NSale.contract, event: "TokenPriceSet", logs: logs, sub: sub}, nil
}

// WatchTokenPriceSet is a free log subscription operation binding the contract event 0x4b59d61d9ffdc3db926d0ce7e06ebabb6bd1bf9dcdae262667e48be368227216.
//
// Solidity: event TokenPriceSet(uint256 newPrice)
func (_C2NSale *C2NSaleFilterer) WatchTokenPriceSet(opts *bind.WatchOpts, sink chan<- *C2NSaleTokenPriceSet) (event.Subscription, error) {

	logs, sub, err := _C2NSale.contract.WatchLogs(opts, "TokenPriceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(C2NSaleTokenPriceSet)
				if err := _C2NSale.contract.UnpackLog(event, "TokenPriceSet", log); err != nil {
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

// ParseTokenPriceSet is a log parse operation binding the contract event 0x4b59d61d9ffdc3db926d0ce7e06ebabb6bd1bf9dcdae262667e48be368227216.
//
// Solidity: event TokenPriceSet(uint256 newPrice)
func (_C2NSale *C2NSaleFilterer) ParseTokenPriceSet(log types.Log) (*C2NSaleTokenPriceSet, error) {
	event := new(C2NSaleTokenPriceSet)
	if err := _C2NSale.contract.UnpackLog(event, "TokenPriceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// C2NSaleTokensSoldIterator is returned from FilterTokensSold and is used to iterate over the raw logs and unpacked data for TokensSold events raised by the C2NSale contract.
type C2NSaleTokensSoldIterator struct {
	Event *C2NSaleTokensSold // Event containing the contract specifics and raw log

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
func (it *C2NSaleTokensSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(C2NSaleTokensSold)
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
		it.Event = new(C2NSaleTokensSold)
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
func (it *C2NSaleTokensSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *C2NSaleTokensSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// C2NSaleTokensSold represents a TokensSold event raised by the C2NSale contract.
type C2NSaleTokensSold struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensSold is a free log retrieval operation binding the contract event 0x57d61f3ccd4ccd25ec5d234d6049553a586fac134c85c98d0b0d9d5724f4e43e.
//
// Solidity: event TokensSold(address user, uint256 amount)
func (_C2NSale *C2NSaleFilterer) FilterTokensSold(opts *bind.FilterOpts) (*C2NSaleTokensSoldIterator, error) {

	logs, sub, err := _C2NSale.contract.FilterLogs(opts, "TokensSold")
	if err != nil {
		return nil, err
	}
	return &C2NSaleTokensSoldIterator{contract: _C2NSale.contract, event: "TokensSold", logs: logs, sub: sub}, nil
}

// WatchTokensSold is a free log subscription operation binding the contract event 0x57d61f3ccd4ccd25ec5d234d6049553a586fac134c85c98d0b0d9d5724f4e43e.
//
// Solidity: event TokensSold(address user, uint256 amount)
func (_C2NSale *C2NSaleFilterer) WatchTokensSold(opts *bind.WatchOpts, sink chan<- *C2NSaleTokensSold) (event.Subscription, error) {

	logs, sub, err := _C2NSale.contract.WatchLogs(opts, "TokensSold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(C2NSaleTokensSold)
				if err := _C2NSale.contract.UnpackLog(event, "TokensSold", log); err != nil {
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

// ParseTokensSold is a log parse operation binding the contract event 0x57d61f3ccd4ccd25ec5d234d6049553a586fac134c85c98d0b0d9d5724f4e43e.
//
// Solidity: event TokensSold(address user, uint256 amount)
func (_C2NSale *C2NSaleFilterer) ParseTokensSold(log types.Log) (*C2NSaleTokensSold, error) {
	event := new(C2NSaleTokensSold)
	if err := _C2NSale.contract.UnpackLog(event, "TokensSold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// C2NSaleTokensWithdrawnIterator is returned from FilterTokensWithdrawn and is used to iterate over the raw logs and unpacked data for TokensWithdrawn events raised by the C2NSale contract.
type C2NSaleTokensWithdrawnIterator struct {
	Event *C2NSaleTokensWithdrawn // Event containing the contract specifics and raw log

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
func (it *C2NSaleTokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(C2NSaleTokensWithdrawn)
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
		it.Event = new(C2NSaleTokensWithdrawn)
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
func (it *C2NSaleTokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *C2NSaleTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// C2NSaleTokensWithdrawn represents a TokensWithdrawn event raised by the C2NSale contract.
type C2NSaleTokensWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensWithdrawn is a free log retrieval operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address user, uint256 amount)
func (_C2NSale *C2NSaleFilterer) FilterTokensWithdrawn(opts *bind.FilterOpts) (*C2NSaleTokensWithdrawnIterator, error) {

	logs, sub, err := _C2NSale.contract.FilterLogs(opts, "TokensWithdrawn")
	if err != nil {
		return nil, err
	}
	return &C2NSaleTokensWithdrawnIterator{contract: _C2NSale.contract, event: "TokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokensWithdrawn is a free log subscription operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address user, uint256 amount)
func (_C2NSale *C2NSaleFilterer) WatchTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *C2NSaleTokensWithdrawn) (event.Subscription, error) {

	logs, sub, err := _C2NSale.contract.WatchLogs(opts, "TokensWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(C2NSaleTokensWithdrawn)
				if err := _C2NSale.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
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

// ParseTokensWithdrawn is a log parse operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address user, uint256 amount)
func (_C2NSale *C2NSaleFilterer) ParseTokensWithdrawn(log types.Log) (*C2NSaleTokensWithdrawn, error) {
	event := new(C2NSaleTokensWithdrawn)
	if err := _C2NSale.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// C2NSaleUserRegisteredIterator is returned from FilterUserRegistered and is used to iterate over the raw logs and unpacked data for UserRegistered events raised by the C2NSale contract.
type C2NSaleUserRegisteredIterator struct {
	Event *C2NSaleUserRegistered // Event containing the contract specifics and raw log

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
func (it *C2NSaleUserRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(C2NSaleUserRegistered)
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
		it.Event = new(C2NSaleUserRegistered)
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
func (it *C2NSaleUserRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *C2NSaleUserRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// C2NSaleUserRegistered represents a UserRegistered event raised by the C2NSale contract.
type C2NSaleUserRegistered struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUserRegistered is a free log retrieval operation binding the contract event 0x54db7a5cb4735e1aac1f53db512d3390390bb6637bd30ad4bf9fc98667d9b9b9.
//
// Solidity: event UserRegistered(address user)
func (_C2NSale *C2NSaleFilterer) FilterUserRegistered(opts *bind.FilterOpts) (*C2NSaleUserRegisteredIterator, error) {

	logs, sub, err := _C2NSale.contract.FilterLogs(opts, "UserRegistered")
	if err != nil {
		return nil, err
	}
	return &C2NSaleUserRegisteredIterator{contract: _C2NSale.contract, event: "UserRegistered", logs: logs, sub: sub}, nil
}

// WatchUserRegistered is a free log subscription operation binding the contract event 0x54db7a5cb4735e1aac1f53db512d3390390bb6637bd30ad4bf9fc98667d9b9b9.
//
// Solidity: event UserRegistered(address user)
func (_C2NSale *C2NSaleFilterer) WatchUserRegistered(opts *bind.WatchOpts, sink chan<- *C2NSaleUserRegistered) (event.Subscription, error) {

	logs, sub, err := _C2NSale.contract.WatchLogs(opts, "UserRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(C2NSaleUserRegistered)
				if err := _C2NSale.contract.UnpackLog(event, "UserRegistered", log); err != nil {
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

// ParseUserRegistered is a log parse operation binding the contract event 0x54db7a5cb4735e1aac1f53db512d3390390bb6637bd30ad4bf9fc98667d9b9b9.
//
// Solidity: event UserRegistered(address user)
func (_C2NSale *C2NSaleFilterer) ParseUserRegistered(log types.Log) (*C2NSaleUserRegistered, error) {
	event := new(C2NSaleUserRegistered)
	if err := _C2NSale.contract.UnpackLog(event, "UserRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
