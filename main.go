package main

import (
	"rh_tests/contracts"
	"rh_tests/helpers"
	"crypto/ecdsa"
	"context"
	"log"
	"math/big"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func pubFromPK(pk string) (common.Address) {
	privateKey, err := crypto.HexToECDSA(pk)
    if err != nil {
        log.Fatal(err)
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("error casting public key to ECDSA")
    }

    return crypto.PubkeyToAddress(*publicKeyECDSA)
}

func main() {
	var addresses helpers.DeployedAddresses

	config, err := helpers.LoadConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("using endpoint", config.Endpoint)

	ctx := context.Background()
	ethConnection, err := ethclient.DialContext(ctx, config.Endpoint)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(config.OraclePK[0])
	if err != nil {
		log.Fatal(err)
	}

	transactor := bind.NewKeyedTransactor(privateKey)
	erc20addr, tx, token, err := contracts.DeployToken(transactor, ethConnection, "TST", "TST", 8)

	if err != nil {
		log.Fatal(err)
	}
	bind.WaitMined(context.Background(), ethConnection, tx)

	addresses.ERC20 = common.Bytes2Hex(erc20addr.Bytes())

	var oracles [5]common.Address
	for i := 0; i < 5; i++ {
		oracles[i] = pubFromPK(config.OraclePK[i])
	}
	nebulaAddr, tx, nebula, err := contracts.DeployNebula(transactor, ethConnection, oracles[:], oracles[0], big.NewInt(3))
	if err != nil {
		log.Fatal(err)
	}
	bind.WaitMined(context.Background(), ethConnection, tx)
	addresses.Nebula = common.Bytes2Hex(nebulaAddr.Bytes())

	ibportAddress, tx, _, err := contracts.DeployIBPort(transactor, ethConnection, nebulaAddr, erc20addr)
	if err != nil {
		log.Fatal(err)
	}
	bind.WaitMined(context.Background(), ethConnection, tx)
	addresses.IBPort = common.Bytes2Hex(ibportAddress.Bytes())

	token.AddMinter(transactor, ibportAddress)
	// token mint
	// token approve

	tx, err = nebula.Subscribe(transactor, ibportAddress, 1, big.NewInt(0))
	receipt, err := bind.WaitMined(context.Background(), ethConnection, tx)
	if err != nil {
		log.Fatal(err)
	}
	
	newSubEvent, err := nebula.NebulaFilterer.ParseNewSubscriber(*receipt.Logs[0])
	if err != nil {
		log.Fatal(err)
	}

	addresses.SubscriptionId = common.Bytes2Hex(newSubEvent.Id[:])

	log.Println(helpers.SaveAddresses(addresses))
}