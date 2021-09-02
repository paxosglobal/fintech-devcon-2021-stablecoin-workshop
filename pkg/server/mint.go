package server

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	contracts "github.com/paxosglobal/fintech-devcon-2021-stablecoin-workshop/build"
	"github.com/shopspring/decimal"
	"log"
	"math/big"
	"strings"
)

const (
	// The address used to deploy the contract becomes the owner
	// We use the first address in the HD wallet that ganache creates with the seed in the docker-compose
	OwnerPrivateKey = "0xca3f21d12dbd7a6a301a70414c40c555761dfce4b3f84cb8ce5f9760a783e3f0"
	ContractAddr    = "0xc4680463046E64b10Da390d9049D24b8EC43AaAB"
)

var (
	OwnerTransactor *bind.TransactOpts
	DecimalToInt    = decimal.New(1, 18) // can also call decimals() on the contract to get the exponent (18)
	UsdkABI         abi.ABI
)

type ChainId *big.Int

var (
	ChainIDMainnet ChainId = big.NewInt(1)
	ChainIDRopsten ChainId = big.NewInt(3)
	ChainIDRinkeby ChainId = big.NewInt(4)
	ChainIDLocal   ChainId = big.NewInt(1337)
)

func init() {
	key, err := crypto.HexToECDSA(OwnerPrivateKey[2:])
	if err != nil {
		panic(err)
	}
	OwnerTransactor, err = bind.NewKeyedTransactorWithChainID(key, ChainIDLocal)
	if err != nil {
		panic(err)
	}
	log.Print("owner address: ", OwnerTransactor.From.String())

	UsdkABI, err = abi.JSON(strings.NewReader(contracts.USDKABI))
	if err != nil {
		panic(err)
	}
}

func (s *Server) mint(destination string, amount decimal.Decimal) error {
	return s.mintWithBindings(destination, amount)
}

func (s *Server) mintWithBindings(destination string, amount decimal.Decimal) error {
	usdkBindings, err := contracts.NewUSDK(common.HexToAddress(ContractAddr), s.ethClient)
	if err != nil {
		return err
	}
	_, err = usdkBindings.Mint(OwnerTransactor, common.HexToAddress(destination), amount.Mul(DecimalToInt).BigInt())
	return err
}

func (s *Server) mintWithExplicitSigning(destination string, amount decimal.Decimal) error {
	ctx := context.Background()
	x, err := s.createMintTransaction(ctx, destination, amount)
	if err != nil {
		return err
	}
	signedTx, err := signTransaction(OwnerTransactor, x)
	if err != nil {
		return err
	}
	if err := s.Broadcast(ctx, signedTx); err != nil {
		panic(err)
	}
	return nil
}

func (s *Server) createMintTransaction(ctx context.Context, destination string, amount decimal.Decimal) (*types.Transaction, error) {
	data, err := createMintTransactionData(destination, amount)
	if err != nil {
		return nil, err
	}
	contractAddress := common.HexToAddress(ContractAddr)
	latestBlockFilter := (*big.Int)(nil)
	// rpc call to read the blockchain node to get the next nonce to use
	currentNonce, err := s.ethClient.NonceAt(ctx, OwnerTransactor.From, latestBlockFilter)
	if err != nil {
		return nil, err
	}
	// rpc call to read the blockchain node to get the current base fee
	header, err := s.ethClient.HeaderByNumber(ctx, latestBlockFilter)
	if err != nil {
		return nil, err
	}
	// rpc call to read the blockchain node to get a suggested tip cap based on recent tip caps on previous blocks
	suggestedTipCap, err := s.ethClient.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, err
	}
	return types.NewTx(&types.DynamicFeeTx{
		ChainID:   ChainIDLocal,
		Nonce:     currentNonce,
		GasFeeCap: (&big.Int{}).Add(header.BaseFee, suggestedTipCap),
		GasTipCap: suggestedTipCap,          // 1 gwei isn't very much in real life, you'll likely pay more
		Gas:       OwnerTransactor.GasLimit, // gas limit - an upper bound on how much compute your transaction will use
		To:        &contractAddress,
		Value:     big.NewInt(0), // amount of Eth being transferred
		Data:      data,          // the smart contract call
	}), nil
}

func createMintTransactionData(destination string, amount decimal.Decimal) ([]byte, error) {
	addr := common.HexToAddress(destination)
	weiInt := amount.Mul(DecimalToInt).BigInt()
	method := "mint"
	return UsdkABI.Pack(method, addr, weiInt)
}

func signTransaction(keyedTransactor *bind.TransactOpts, x *types.Transaction) (*types.Transaction, error) {
	return keyedTransactor.Signer(keyedTransactor.From, x)
}

func (s *Server) Broadcast(ctx context.Context, x *types.Transaction) error {
	return s.ethClient.SendTransaction(ctx, x)
}
