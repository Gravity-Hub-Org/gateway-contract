package helpers

import (
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
    "strings"
	"math/big"
	"log"
	"crypto/ecdsa"
	"context"
)

type NebulaCaller struct {
	contract *bind.BoundContract
	Backend bind.ContractBackend
	OraclePK []string
}

func NewNebula(address string, abiString string, backend bind.ContractBackend, OraclePK []string) (NebulaCaller) {
	connectAddress := common.HexToAddress(address)
	parsedAbi, _ := abi.JSON(strings.NewReader(abiString))
	return NebulaCaller{
		contract: bind.NewBoundContract(connectAddress, parsedAbi, backend, backend, backend),
		Backend: backend,
		OraclePK: OraclePK,
	}
}

func (caller *NebulaCaller) GetPulseHash(pulseNum *big.Int) ([32]byte) {
    var out [32]byte
    caller.contract.Call(nil, &out, "pulses", pulseNum)
    return out
}

func (caller *NebulaCaller) SendData(value uint64, blockNumber big.Int, subscriptionId [32]byte) (bool) {
    privateKey, err := crypto.HexToECDSA(caller.OraclePK[0])
    if err != nil {
        log.Fatal(err)
    }
    auth := bind.NewKeyedTransactor(privateKey)
    _, err = caller.contract.Transact(auth, "sendData", value, blockNumber, subscriptionId)
    if err != nil {
        return false
    }
    return true
}

func (caller *NebulaCaller) SignData(dataHash [32]byte, validSignsCount int) bool {
    var r [5][32]byte
    var s [5][32]byte
    var v [5]uint8

    privateKey, err := crypto.HexToECDSA(caller.OraclePK[0])
    if err != nil {
        log.Fatal(err)
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("error casting public key to ECDSA")
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := caller.Backend.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }

    gasPrice, err := caller.Backend.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    auth := bind.NewKeyedTransactor(privateKey)
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)     // in wei
    auth.GasLimit = uint64(300000) // in units
    auth.GasPrice = gasPrice

    for position, validatorKey := range caller.OraclePK {
        validatorEthKey, _ := crypto.HexToECDSA(validatorKey)
        sign, _ := crypto.Sign(dataHash[:], validatorEthKey)

        copy(r[position][:], sign[:32])
        copy(s[position][:], sign[32:64])

        if (position < validSignsCount) {
            v[position] = sign[64] + 27
        } else {
            v[position] = 0 // generate invalid signature
        }
    }

    _, err = caller.contract.Transact(auth, "confirmData", dataHash, v[:], r[:], s[:])
    if err != nil {
        return false
	}
	
    return true
}
