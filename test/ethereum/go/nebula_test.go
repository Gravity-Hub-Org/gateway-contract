package nebula_test

import (
	"context"
	"math/rand"
	"nebula-test/helpers"
	"crypto/ecdsa"
	"os"
	"testing"
	"math/big"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ctx context.Context
var ethConnection *ethclient.Client
var nebulaContract helpers.NebulaCaller
var ibportContract helpers.IBPortCaller
var config helpers.Config
var nebulaData helpers.NebulaData

func TestMain(m *testing.M) {
	var err error
	config, nebulaData, err = helpers.LoadConfiguration()
	if err != nil {
		panic(err)
	}

	ctx = context.Background()
	ethConnection, err = ethclient.DialContext(ctx, config.Endpoint)
	if err != nil {
		panic(err)
	}

	nebulaContract = helpers.NewNebula(nebulaData.Address, nebulaData.Abi, ethConnection, config.OraclePK[:])
	ibportContract = helpers.NewIBPort(nebulaData.IbportAddress, nebulaData.IbportAbi, ethConnection, config.OraclePK[:])
	os.Exit(m.Run())
}

func Random32Byte() [32]byte {
	var out [32]byte
	rand.Read(out[:])
	return out
}

func filluint(data []byte, pos uint, val uint64) {
	var i uint
	for i = 0; i < 32; i++ {
		data[i + pos] = byte(val % 256)
		val = val / 256
	}
}

func filladdress(data []byte, pos uint, addressStr string) {
	address := common.HexToAddress(addressStr)
	copy(data[pos:], address[:])
}

func bytes32fromhex(s string) ([32]byte) {
	var ret [32]byte
	decoded, err := hex.DecodeString(s[2:])
	if err != nil {
		return ret
	}
	copy(ret[:], decoded[:])
	return ret
}

func TestPulseSaved(t *testing.T) {
	d := Random32Byte()
	height := nebulaContract.SignData(d, 5)
	if height == nil {
		t.Error("can't send signed data")
	} else {
		pulseData := nebulaContract.GetPulseHash(height)

		if d != pulseData {
			t.Error("data mismatch")
		}
	}
}

func TestPulseCorrect3(t *testing.T) {
	d := Random32Byte()
	if nebulaContract.SignData(d, 3) == nil {
		t.Error("can't send signed data")
	}
}

func TestPulseInCorrect2(t *testing.T) {
	d := Random32Byte()
	if nebulaContract.SignData(d, 2) != nil {
		t.Error("transaction 2/3 valid sigs should be rejected")
	}
}

func TestInvalidHash(t *testing.T) {
	var attachedData [2]byte
	attachedData[0] = 1
	attachedData[1] = 2

	// generate invalid proof
	proof := crypto.Keccak256Hash(attachedData[:])

	height := nebulaContract.SignData(proof, 5)
	if height == nil {
		t.Error("can't submit proof")
	} else if nebulaContract.SendData(attachedData[:], height, bytes32fromhex(nebulaData.SubscriptionId)) {
		t.Error("this tx should fail because of invalid hash")
	}
}

func TestMint(t *testing.T) {
	var attachedData [1+32+32+20]byte

	attachedData[0] = 'm' // mint
	filluint(attachedData[:], 1, 123456789) // req id
	filluint(attachedData[:], 1+32, 10) // amount = 10
	filladdress(attachedData[:], 1+32+32, "9561C133DD8580860B6b7E504bC5Aa500f0f0103") // address

	proof := crypto.Keccak256Hash(attachedData[:])
	height := nebulaContract.SignData(proof, 5)
	if height == nil {
		t.Error("can't submit proof")
	} else if !nebulaContract.SendData(attachedData[:], height, bytes32fromhex(nebulaData.SubscriptionId)) {
		t.Error("can't submit data")
	}
}

func TestChangeStatusFail(t *testing.T) {
	var attachedData [1+32+1]byte

	attachedData[0] = 'c' // change status
	filluint(attachedData[:], 1, 111111111) // req id = 111111111
	attachedData[1+32] = 1

	proof := crypto.Keccak256Hash(attachedData[:])
	height := nebulaContract.SignData(proof, 5)
	if height == nil {
		t.Error("can't submit proof")
	}

	if nebulaContract.SendData(attachedData[:], height, bytes32fromhex(nebulaData.SubscriptionId)) {
		t.Error("request should fail")
	}
}

func TestChangeStatusOk(t *testing.T) {
	var attachedData [1+32+1]byte
	var address [32]byte

	privateKey, err := crypto.HexToECDSA(config.OraclePK[0])
	if err != nil {
		t.Error(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		t.Error("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	copy(address[:], fromAddress[:])

	reqId := ibportContract.CreateTransferUnwrapRequest(big.NewInt(1), address)

	if reqId == nil {
		t.Error("request failed")
	} else {
		attachedData[0] = 'c' // change status
		filluint(attachedData[:], 1, 0)
		attachedData[32] = reqId.Bytes()[0]

		attachedData[1+32] = 2 // next status

		proof := crypto.Keccak256Hash(attachedData[:])
		height := nebulaContract.SignData(proof, 5)
		if height == nil {
			t.Error("can't submit proof")
		}

		if !nebulaContract.SendData(attachedData[:], height, bytes32fromhex(nebulaData.SubscriptionId)) {
			t.Error("request failed")
		}
	}
}
