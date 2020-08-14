package nebula_test

import (
	"nebula-test/helpers"
	"bytes"
	"os"
	"log"
	"testing"
	"math/rand"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ethConnection *ethclient.Client
var nebulaContract helpers.NebulaCaller
var config helpers.Config
var nebulaData helpers.NebulaData

func TestMain(m *testing.M) {
    config, nebulaData = helpers.LoadConfiguration()
    ethConnection = helpers.ConnectClient(config.Endpoint)
    nebulaContract = helpers.NewNebula(nebulaData.Address, nebulaData.Abi, ethConnection, config.OraclePK[:])
	os.Exit(m.Run())
}

func Random32Byte() ([32]byte) {
    var out [32]byte
    rand.Read(out[:])
    return out
}

func TestPulseSaved(t *testing.T) {
    d := Random32Byte()
    if !nebulaContract.SignData(d, 5) {
		log.Fatal("can't send signed data")
        t.Fail()
    }
    height := helpers.GetCurrentHeight(ethConnection)
    pulseData := nebulaContract.GetPulseHash(height)

    if bytes.Compare(d[:], pulseData[:]) != 0 {
		log.Fatal("data mismatch")
        t.Fail()
    }
}

func TestPulseCorect3(t *testing.T) {
    d := Random32Byte()
    if !nebulaContract.SignData(d, 3) {
		log.Fatal("can't send signed data")
        t.Fail()
    }
}

func TestPulseInCorect2(t *testing.T) {
    d := Random32Byte()
    if nebulaContract.SignData(d, 2) {
		log.Fatal("transaction 2/3 valid sigs should be rejected")
        t.Fail()
    }
}
