package nebula_test

import (
	"context"
	"math/rand"
	"nebula-test/helpers"
	"os"
	"testing"
	"encoding/hex"

	"golang.org/x/crypto/sha3"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ctx context.Context
var ethConnection *ethclient.Client
var nebulaContract helpers.NebulaCaller
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
	os.Exit(m.Run())
}

func Random32Byte() [32]byte {
	var out [32]byte
	rand.Read(out[:])
	return out
}

// func TestPulseSaved(t *testing.T) {
// 	d := Random32Byte()
// 	if !nebulaContract.SignData(d, 5) {
// 		t.Error("can't send signed data")
// 	}
// 	height, err := ethConnection.BlockByNumber(ctx, nil)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	pulseData := nebulaContract.GetPulseHash(height.Number())

// 	if d != pulseData {
// 		t.Error("data mismatch")
// 	}
// }

// func TestPulseCorrect3(t *testing.T) {
// 	d := Random32Byte()
// 	if !nebulaContract.SignData(d, 3) {
// 		t.Error("can't send signed data")
// 	}
// }

// func TestPulseInCorrect2(t *testing.T) {
// 	d := Random32Byte()
// 	if nebulaContract.SignData(d, 2) {
// 		t.Error("transaction 2/3 valid sigs should be rejected")
// 	}
// }

func filluint(data []byte, pos uint, val uint) {
	var i uint
	for i = 0; i < 32; i++ {
		data[i + pos] = byte(val % 256)
		val = val / 256
	}
}

func filladdress(data []byte, pos uint) {
	var i uint
	for i = 0; i < 20; i++ {
		data[i + pos] = 0
	}
}

func bytes32fromhex(s string) ([32]byte) {
	var ret [32]byte
	decoded, err := hex.DecodeString(s)
	if err != nil {
		return ret
	}
	copy(decoded, ret[:])
	return ret
}

// func TestMint(t *testing.T) {
// 	var attachedData [1+32+32+20]byte

// 	attachedData[0] = 'm' // mint
// 	filluint(attachedData[:], 1, 10) // req id = 10
// 	filluint(attachedData[:], 1+32, 10) // amount = 10
// 	filladdress(attachedData[:], 1+32+32) // address

// 	proof := sha3.Sum256(attachedData[:])
// 	if !nebulaContract.SignData(proof, 5) {
// 		t.Error("can't submit proof")
// 	}

// 	height, err := ethConnection.BlockByNumber(ctx, nil)
// 	if err != nil {
// 		t.Error(err)
// 	}
	
// 	if !nebulaContract.SendData(attachedData[:], height.Number(), bytes32fromhex(nebulaData.SubscriptionId)) {
// 		t.Error("can't submit data")
// 	}

// }

func TestChangeStatus(t *testing.T) {
	var attachedData [1+32+1]byte

	attachedData[0] = 'c' // mint
	filluint(attachedData[:], 1, 11) // req id = 10
	attachedData[1+32] = 1

	proof := sha3.Sum256(attachedData[:])
	if !nebulaContract.SignData(proof, 5) {
		t.Error("can't submit proof")
	}

	height, err := ethConnection.BlockByNumber(ctx, nil)
	if err != nil {
		t.Error(err)
	}
	
	if !nebulaContract.SendData(attachedData[:], height.Number(), bytes32fromhex(nebulaData.SubscriptionId)) {
		t.Error("can't submit data")
	}

}