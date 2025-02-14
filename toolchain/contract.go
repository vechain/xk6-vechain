// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package toolchain

import (
	"context"
	"errors"
	"math/big"
	"strings"

	"github.com/darrenvechain/thorgo/accounts"
	"github.com/darrenvechain/thorgo/blocks"
	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/darrenvechain/thorgo/transactions"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = bind.Bind
	_ = common.Big1
	_ = abi.ConvertType
	_ = hexutil.MustDecode
	_ = context.Background
	_ = tx.NewClause
	_ = blocks.New
)

// ToolchainTodo is an auto generated low-level Go binding around a user-defined struct.
type ToolchainTodo struct {
	Text string

	Completed bool
}

// ToolchainMetaData contains all meta data concerning the Toolchain contract.
var ToolchainMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"b\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"c\",\"type\":\"bytes32\"}],\"name\":\"ToolchainEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTodo\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"text\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"completed\",\"type\":\"bool\"}],\"internalType\":\"structToolchain.Todo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payMe\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomFunc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"a\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"b\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"c\",\"type\":\"bytes32\"}],\"name\":\"setBytes32\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610563806100206000396000f3fe60806040526004361061004a5760003560e01c8063722713f71461004f5780637fa1680a1461007a578063a1b9243d146100a5578063aecb29bf146100d1578063d997ccb3146100fa575b600080fd5b34801561005b57600080fd5b50610064610104565b6040516100719190610230565b60405180910390f35b34801561008657600080fd5b5061008f61010c565b60405161009c9190610333565b60405180910390f35b3480156100b157600080fd5b506100ba610166565b6040516100c8929190610396565b60405180910390f35b3480156100dd57600080fd5b506100f860048036038101906100f39190610433565b610175565b005b6101026101b6565b005b600047905090565b6101146101fb565b60405180604001604052806040518060400160405280600d81526020017f48656c6c6f2c20576f726c642100000000000000000000000000000000000000815250815260200160001515815250905090565b60008060016000915091509091565b818360ff167f80f229eeb0808ec0efca1b655fc050dbf966be96a3f44ff4bf2df6b948e61c93836040516101a99190610495565b60405180910390a3505050565b600034116101f9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101f09061050d565b60405180910390fd5b565b6040518060400160405280606081526020016000151581525090565b6000819050919050565b61022a81610217565b82525050565b60006020820190506102456000830184610221565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561028557808201518184015260208101905061026a565b60008484015250505050565b6000601f19601f8301169050919050565b60006102ad8261024b565b6102b78185610256565b93506102c7818560208601610267565b6102d081610291565b840191505092915050565b60008115159050919050565b6102f0816102db565b82525050565b6000604083016000830151848203600086015261031382826102a2565b915050602083015161032860208601826102e7565b508091505092915050565b6000602082019050818103600083015261034d81846102f6565b905092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061038082610355565b9050919050565b61039081610375565b82525050565b60006040820190506103ab6000830185610221565b6103b86020830184610387565b9392505050565b600080fd5b600060ff82169050919050565b6103da816103c4565b81146103e557600080fd5b50565b6000813590506103f7816103d1565b92915050565b6000819050919050565b610410816103fd565b811461041b57600080fd5b50565b60008135905061042d81610407565b92915050565b60008060006060848603121561044c5761044b6103bf565b5b600061045a868287016103e8565b935050602061046b8682870161041e565b925050604061047c8682870161041e565b9150509250925092565b61048f816103fd565b82525050565b60006020820190506104aa6000830184610486565b92915050565b600082825260208201905092915050565b7f4e6f206d6f6e65792073656e7400000000000000000000000000000000000000600082015250565b60006104f7600d836104b0565b9150610502826104c1565b602082019050919050565b60006020820190508181036000830152610526816104ea565b905091905056fea264697066735822122024fbbd2e99212972ddf1764a4dbd0ed735cdfb0ac7cfda89a66c5bae5cd596ec64736f6c63430008130033",
}

// DeployToolchain deploys a new Ethereum contract, binding an instance of Toolchain to it.
func DeployToolchain(ctx context.Context, thor *thorest.Client, sender accounts.TxManager, opts *transactions.Options) (common.Hash, *ToolchainTransactor, error) {
	parsed, err := ToolchainMetaData.GetAbi()
	if err != nil {
		return common.Hash{}, nil, err
	}
	if parsed == nil {
		return common.Hash{}, nil, errors.New("GetABI returned nil")
	}

	bytes, err := hexutil.Decode(ToolchainMetaData.Bin)
	if err != nil {
		return common.Hash{}, nil, err
	}
	contract, txID, err := accounts.NewDeployer(thor, bytes, parsed).Deploy(ctx, sender, opts)
	if err != nil {
		return common.Hash{}, nil, err
	}
	return txID, &ToolchainTransactor{&Toolchain{thor: thor, contract: contract}, contract.Transactor(sender), sender}, nil
}

// Toolchain is an auto generated Go binding around an Ethereum contract, allowing you to query and create clauses.
type Toolchain struct {
	thor     *thorest.Client    // Thor client connection to use
	contract *accounts.Contract // Generic contract wrapper for the low level calls
}

// ToolchainTransactor is an auto generated Go binding around an Ethereum, allowing you to transact with the contract.
type ToolchainTransactor struct {
	*Toolchain
	contract *accounts.ContractTransactor // Generic contract wrapper for the low level calls
	manager  accounts.TxManager           // TxManager to use
}

// NewToolchain creates a new instance of Toolchain, bound to a specific deployed contract.
func NewToolchain(address common.Address, thor *thorest.Client) (*Toolchain, error) {
	parsed, err := ToolchainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	contract := accounts.New(thor, address).Contract(parsed)
	return &Toolchain{thor: thor, contract: contract}, nil
}

// NewToolchainTransactor creates a new instance of ToolchainTransactor, bound to a specific deployed contract.
func NewToolchainTransactor(address common.Address, thor *thorest.Client, manager accounts.TxManager) (*ToolchainTransactor, error) {
	base, err := NewToolchain(address, thor)
	if err != nil {
		return nil, err
	}
	return &ToolchainTransactor{Toolchain: base, contract: base.contract.Transactor(manager), manager: manager}, nil
}

// Address returns the address of the contract.
func (_Toolchain *Toolchain) Address() common.Address {
	return _Toolchain.contract.Address
}

// Transactor constructs a new transactor for the contract, which allows to send transactions.
func (_Toolchain *Toolchain) Transactor(manager accounts.TxManager) *ToolchainTransactor {
	return &ToolchainTransactor{Toolchain: _Toolchain, contract: _Toolchain.contract.Transactor(manager), manager: manager}
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Toolchain *Toolchain) Call(revision thorest.Revision, result *[]interface{}, method string, params ...interface{}) error {
	return _Toolchain.contract.CallAt(revision, method, result, params...)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ToolchainTransactor *ToolchainTransactor) Transact(opts *transactions.Options, method string, params ...interface{}) (*transactions.Visitor, error) {
	return _ToolchainTransactor.contract.Send(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x722713f7.
//
// Solidity: function balanceOf() view returns(uint256)
func (_Toolchain *Toolchain) BalanceOf(revision ...thorest.Revision) (*big.Int, error) {
	var rev thorest.Revision
	if len(revision) > 0 {
		rev = revision[0]
	} else {
		rev = thorest.RevisionBest()
	}

	var out []interface{}
	err := _Toolchain.Call(rev, &out, "balanceOf")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// GetTodo is a free data retrieval call binding the contract method 0x7fa1680a.
//
// Solidity: function getTodo() pure returns((string,bool))
func (_Toolchain *Toolchain) GetTodo(revision ...thorest.Revision) (ToolchainTodo, error) {
	var rev thorest.Revision
	if len(revision) > 0 {
		rev = revision[0]
	} else {
		rev = thorest.RevisionBest()
	}

	var out []interface{}
	err := _Toolchain.Call(rev, &out, "getTodo")

	if err != nil {
		return *new(ToolchainTodo), err
	}

	out0 := *abi.ConvertType(out[0], new(ToolchainTodo)).(*ToolchainTodo)

	return out0, err
}

// RandomFunc is a free data retrieval call binding the contract method 0xa1b9243d.
//
// Solidity: function randomFunc() pure returns(uint256, address)
func (_Toolchain *Toolchain) RandomFunc(revision ...thorest.Revision) (*big.Int, common.Address, error) {
	var rev thorest.Revision
	if len(revision) > 0 {
		rev = revision[0]
	} else {
		rev = thorest.RevisionBest()
	}

	var out []interface{}
	err := _Toolchain.Call(rev, &out, "randomFunc")

	if err != nil {
		return *new(*big.Int), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err
}

// PayMe is a paid mutator transaction binding the contract method 0xd997ccb3.
//
// Solidity: function payMe() payable returns()
func (_ToolchainTransactor *ToolchainTransactor) PayMe(opts *transactions.Options) (*transactions.Visitor, error) {
	return _ToolchainTransactor.Transact(opts, "payMe")
}

// PayMeAsClause is a transaction clause generator 0xd997ccb3.
//
// Solidity: function payMe() payable returns()
func (_Toolchain *Toolchain) PayMeAsClause(vetValue ...*big.Int) (*tx.Clause, error) {
	var val *big.Int
	if len(vetValue) > 0 {
		val = vetValue[0]
	} else {
		val = big.NewInt(0)
	}
	return _Toolchain.contract.AsClauseWithVET(val, "payMe")
}

// SetBytes32 is a paid mutator transaction binding the contract method 0xaecb29bf.
//
// Solidity: function setBytes32(uint8 a, bytes32 b, bytes32 c) returns()
func (_ToolchainTransactor *ToolchainTransactor) SetBytes32(a uint8, b [32]byte, c [32]byte, opts *transactions.Options) (*transactions.Visitor, error) {
	return _ToolchainTransactor.Transact(opts, "setBytes32", a, b, c)
}

// SetBytes32AsClause is a transaction clause generator 0xaecb29bf.
//
// Solidity: function setBytes32(uint8 a, bytes32 b, bytes32 c) returns()
func (_Toolchain *Toolchain) SetBytes32AsClause(a uint8, b [32]byte, c [32]byte, vetValue ...*big.Int) (*tx.Clause, error) {
	var val *big.Int
	if len(vetValue) > 0 {
		val = vetValue[0]
	} else {
		val = big.NewInt(0)
	}
	return _Toolchain.contract.AsClauseWithVET(val, "setBytes32", a, b, c)
}

// ToolchainToolchainEvent represents a ToolchainEvent event raised by the Toolchain contract.
type ToolchainToolchainEvent struct {
	A   *big.Int
	B   [32]byte
	C   [32]byte
	Log thorest.EventLog
}

type ToolchainToolchainEventCriteria struct {
	A *big.Int  `abi:"a"`
	B *[32]byte `abi:"b"`
}

// FilterToolchainEvent is a free log retrieval operation binding the contract event 0x80f229eeb0808ec0efca1b655fc050dbf966be96a3f44ff4bf2df6b948e61c93.
//
// Solidity: event ToolchainEvent(uint256 indexed a, bytes32 indexed b, bytes32 c)
func (_Toolchain *Toolchain) FilterToolchainEvent(criteria []ToolchainToolchainEventCriteria, filters *thorest.LogFilters) ([]ToolchainToolchainEvent, error) {
	topicHash := _Toolchain.contract.ABI.Events["ToolchainEvent"].ID

	criteriaSet := make([]thorest.EventCriteria, len(criteria))
	for i, c := range criteria {
		crteria := thorest.EventCriteria{
			Address: &_Toolchain.contract.Address,
			Topic0:  &topicHash,
		}
		if c.A != nil {
			matcher := c.A
			topics, err := abi.MakeTopics([]interface{}{matcher})
			if err != nil {
				return nil, err
			}
			crteria.Topic1 = &topics[0][0]
		}
		if c.B != nil {
			matcher := *c.B
			topics, err := abi.MakeTopics([]interface{}{matcher})
			if err != nil {
				return nil, err
			}
			crteria.Topic2 = &topics[0][0]
		}

		criteriaSet[i] = crteria
	}

	if len(criteriaSet) == 0 {
		criteriaSet = append(criteriaSet, thorest.EventCriteria{
			Address: &_Toolchain.contract.Address,
			Topic0:  &topicHash,
		})
	}

	logs, err := _Toolchain.thor.FilterEvents(criteriaSet, filters)
	if err != nil {
		return nil, err
	}

	inputs := _Toolchain.contract.ABI.Events["ToolchainEvent"].Inputs
	var indexed abi.Arguments
	for _, arg := range inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}

	events := make([]ToolchainToolchainEvent, len(logs))
	for i, log := range logs {
		event := new(ToolchainToolchainEvent)
		if err := _Toolchain.contract.UnpackLog(event, "ToolchainEvent", log); err != nil {
			return nil, err
		}
		event.Log = log
		events[i] = *event
	}

	return events, nil
}

// WatchToolchainEvent listens for on chain events binding the contract event 0x80f229eeb0808ec0efca1b655fc050dbf966be96a3f44ff4bf2df6b948e61c93.
//
// Solidity: event ToolchainEvent(uint256 indexed a, bytes32 indexed b, bytes32 c)
func (_Toolchain *Toolchain) WatchToolchainEvent(criteria []ToolchainToolchainEventCriteria, ctx context.Context, bufferSize int64) (chan *ToolchainToolchainEvent, error) {
	topicHash := _Toolchain.contract.ABI.Events["ToolchainEvent"].ID

	criteriaSet := make([]thorest.EventCriteria, len(criteria))

	for i, c := range criteria {
		crteria := thorest.EventCriteria{
			Address: &_Toolchain.contract.Address,
			Topic0:  &topicHash,
		}
		if c.A != nil {
			matcher := c.A
			topics, err := abi.MakeTopics([]interface{}{matcher})
			if err != nil {
				return nil, err
			}
			crteria.Topic1 = &topics[0][0]
		}
		if c.B != nil {
			matcher := *c.B
			topics, err := abi.MakeTopics([]interface{}{matcher})
			if err != nil {
				return nil, err
			}
			crteria.Topic2 = &topics[0][0]
		}

		criteriaSet[i] = crteria
	}

	eventChan := make(chan *ToolchainToolchainEvent, bufferSize)
	blocks := blocks.New(ctx, _Toolchain.thor)
	ticker := blocks.Ticker()

	go func() {
		defer close(eventChan)

		for {
			select {
			case <-ticker.C():
				block, err := blocks.Best()
				if err != nil {
					continue
				}
				for _, tx := range block.Transactions {
					for index, outputs := range tx.Outputs {
						for _, event := range outputs.Events {
							for _, c := range criteriaSet {
								if !c.Matches(event) {
									continue
								}
							}

							log := thorest.EventLog{
								Address: &_Toolchain.contract.Address,
								Topics:  event.Topics,
								Data:    event.Data,
								Meta: thorest.LogMeta{
									BlockID:     block.ID,
									BlockNumber: block.Number,
									BlockTime:   block.Timestamp,
									TxID:        tx.ID,
									TxOrigin:    tx.Origin,
									ClauseIndex: int64(index),
								},
							}

							ev := new(ToolchainToolchainEvent)
							if err := _Toolchain.contract.UnpackLog(ev, "ToolchainEvent", log); err != nil {
								continue
							}
							ev.Log = log
							eventChan <- ev
						}
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return eventChan, nil
}
