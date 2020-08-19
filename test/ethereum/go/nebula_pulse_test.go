package nebula_test

import (
	"context"
	"math/rand"
	"nebula-test/helpers"
	"os"
	"testing"

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

func TestPulseSaved(t *testing.T) {
	d := Random32Byte()
	if !nebulaContract.SignData(d, 5) {
		t.Error("can't send signed data")
	}
	height, err := ethConnection.BlockByNumber(ctx, nil)
	if err != nil {
		t.Error(err)
	}
	pulseData := nebulaContract.GetPulseHash(height.Number())

	if d != pulseData {
		t.Error("data mismatch")
	}
}

func TestPulseCorrect3(t *testing.T) {
	d := Random32Byte()
	if !nebulaContract.SignData(d, 3) {
		t.Error("can't send signed data")
	}
}

func TestPulseInCorrect2(t *testing.T) {
	d := Random32Byte()
	if nebulaContract.SignData(d, 2) {
		t.Error("transaction 2/3 valid sigs should be rejected")
	}
}
