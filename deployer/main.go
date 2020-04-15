package main

import (
	"crypto/ecdsa"
	"deployer/config"
	"deployer/contracts"
	"deployer/helpers"
	"flag"
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"

	"github.com/ethereum/go-ethereum/common/hexutil"
	wavesplatform "github.com/wavesplatform/go-lib-crypto"
	"github.com/wavesplatform/gowaves/pkg/client"
	"github.com/wavesplatform/gowaves/pkg/crypto"
	"github.com/wavesplatform/gowaves/pkg/proto"
)

type State struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
	Type  string      `json:"type"`
}

const (
	defaultConfigFileName = "config.json"
	defaultDbPath         = "db"
	defaultHost           = "127.0.0.1:8080"
)

var pubKey crypto.PublicKey
var cfg config.Config
var secret crypto.SecretKey
var nodeClient *client.Client
var ethNode *ethclient.Client
var ethPrivateKey *ecdsa.PrivateKey

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	var confFileName string
	flag.StringVar(&confFileName, "config", defaultConfigFileName, "set config path")
	flag.Parse()

	var err error
	cfg, err = config.Load(confFileName)
	if err != nil {
		panic(err)
	}
	wCrypto := wavesplatform.NewWavesCrypto()
	distributorWavesSeed := wavesplatform.Seed(cfg.DistributorWavesSeed)
	distributorSecret, err := crypto.NewSecretKeyFromBase58(string(wCrypto.PrivateKey(distributorWavesSeed)))
	checkErr(err)

	seed, wavesContractAddress := createWavesContractAddress(distributorSecret, wCrypto)
	secret, err = crypto.NewSecretKeyFromBase58(string(wCrypto.PrivateKey(seed)))
	checkErr(err)

	pubKey = crypto.GeneratePublicKey(secret)
	wCrypto.RandomSeed()
	nodeClient, err = client.NewClient(client.Options{ApiKey: "", BaseUrl: cfg.WavesNodeUrl})
	checkErr(err)

	scriptBytes, err := ioutil.ReadFile("/tmp/dat")
	fmt.Print(string(scriptBytes))

	//-----------------------------------
	ethNode, err = ethclient.DialContext(nil, cfg.EthNodeUrl)
	checkErr(err)

	privBytes, err := hexutil.Decode(cfg.EthereumPrivKey)
	checkErr(err)

	ethPrivateKey = &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: secp256k1.S256(),
		},
		D: new(big.Int),
	}
	ethPrivateKey.D.SetBytes(privBytes)
	ethPrivateKey.PublicKey.X, ethPrivateKey.PublicKey.Y = ethPrivateKey.PublicKey.Curve.ScalarBaseMult(privBytes)

	//------------------------------------
	ethAssetId := newEthereumAssetInWaves()
	contractAddress, wavesERC20Address := deployEthereumContractAndWavesERC20(ethAssetId)
	deployWavesContract(scriptBytes)
	setDataWavesContract(ethAssetId, wavesERC20Address)

	println("Waves Contract address:" + wavesContractAddress)
	println("Seed Waves Contract address:" + seed)
	println("Eth Contract address:" + contractAddress)
	println("Eth asset id:" + ethAssetId)
	println("Waves erc20:" + wavesERC20Address)

}

func createWavesContractAddress(distributorSecret crypto.SecretKey, wCrypto wavesplatform.WavesCrypto) (wavesplatform.Seed, string) {
	contractSeed := wCrypto.RandomSeed()
	address := wCrypto.Address(wCrypto.PublicKey(contractSeed), wavesplatform.WavesChainID(cfg.ChainId[0]))
	ad, err := proto.NewAddressFromString(string(address))
	checkErr(err)

	tx := &proto.TransferV2{
		Type:    proto.SetScriptTransaction,
		Version: 1,
		Transfer: proto.Transfer{
			Amount:    1000000000,
			Recipient: proto.NewRecipientFromAddress(ad),
			SenderPK:  pubKey,
			Fee:       10000000,
			Timestamp: client.NewTimestampFromTime(time.Now()),
		},
	}
	err = tx.Sign(distributorSecret)
	checkErr(err)

	_, err = nodeClient.Transactions.Broadcast(nil, tx)
	checkErr(err)

	err = <-helpers.WaitTx(nodeClient, tx.ID)
	checkErr(err)

	return contractSeed, string(address)
}

func newEthereumAssetInWaves() string {
	tx := &proto.IssueV2{
		Type:    proto.SetScriptTransaction,
		Version: 1,
		Issue: proto.Issue{
			Decimals:    8,
			Quantity:    11053382600000000,
			Reissuable:  true,
			Description: "Susy ETH",
			Fee:         100000000,
			Name:        "sETH",
			SenderPK:    pubKey,
			Timestamp:   client.NewTimestampFromTime(time.Now()),
		},
		ChainID: cfg.ChainId[0],
	}
	err := tx.Sign(secret)
	checkErr(err)

	_, err = nodeClient.Transactions.Broadcast(nil, tx)
	checkErr(err)

	err = <-helpers.WaitTx(nodeClient, tx.ID)
	checkErr(err)

	return tx.ID.String()
}

func deployEthereumContractAndWavesERC20(ethAssetId string) (contractAddress string, wavesERC20 string) {
	transactOpt := bind.NewKeyedTransactor(ethPrivateKey)
	var admins [5]common.Address
	for i, admin := range cfg.Admins {
		admins[i] = common.HexToAddress(admin)
	}
	address, _, contract, err := contracts.DeploySupersymmetry(transactOpt, nil, admins, uint8(cfg.BftCoefficient), ethAssetId, uint8(cfg.RqTimeout))
	checkErr(err)

	_, err = contract.NewWavesToken(transactOpt)
	checkErr(err)

	filter, err := contract.FilterNewApprovedToken(nil)
	checkErr(err)
	filter.Next()
	return address.String(), filter.Event.TokenAddress.String()
}

func deployWavesContract(script []byte) {
	tx := &proto.SetScriptV1{
		Type:      proto.SetScriptTransaction,
		Version:   1,
		SenderPK:  pubKey,
		ChainID:   cfg.ChainId[0],
		Script:    script,
		Fee:       10000000,
		Timestamp: client.NewTimestampFromTime(time.Now()),
	}
	err := tx.Sign(secret)
	checkErr(err)

	_, err = nodeClient.Transactions.Broadcast(nil, tx)
	checkErr(err)

	err = <-helpers.WaitTx(nodeClient, tx.ID)
	checkErr(err)
}

func setDataWavesContract(ethAssetId string, wavesERC20Address string) {
	tx := &proto.DataV1{
		Type:     proto.SetScriptTransaction,
		Version:  1,
		SenderPK: pubKey,
		Entries: proto.DataEntries{
			&proto.StringDataEntry{
				Key:   "admins",
				Value: strings.Join(cfg.Admins, ","),
			},
			&proto.IntegerDataEntry{
				Key:   "bftCoefficient",
				Value: cfg.BftCoefficient,
			},
			&proto.IntegerDataEntry{
				Key:   "rq_timeout",
				Value: cfg.RqTimeout,
			},
			//---------------------------------------
			&proto.StringDataEntry{
				Key:   "erc20_address_WAVES",
				Value: wavesERC20Address,
			},
			&proto.IntegerDataEntry{
				Key:   "asset_type_WAVES",
				Value: 0,
			},
			&proto.IntegerDataEntry{
				Key:   "asset_status_WAVES",
				Value: 3,
			},
			&proto.IntegerDataEntry{
				Key:   "asset_decimals_WAVES",
				Value: 8,
			},
			//----------------------------------------
			&proto.StringDataEntry{
				Key:   "erc20_address_" + ethAssetId,
				Value: "0x0000000000000000000000000000000000000000",
			},
			&proto.IntegerDataEntry{
				Key:   "asset_type_" + ethAssetId,
				Value: 1,
			},
			&proto.IntegerDataEntry{
				Key:   "asset_status_" + ethAssetId,
				Value: 3,
			},
			&proto.IntegerDataEntry{
				Key:   "asset_decimals_" + ethAssetId,
				Value: 8,
			},
		},
		Fee:       10000000,
		Timestamp: client.NewTimestampFromTime(time.Now()),
	}
	err := tx.Sign(secret)
	checkErr(err)

	_, err = nodeClient.Transactions.Broadcast(nil, tx)
	checkErr(err)

	err = <-helpers.WaitTx(nodeClient, tx.ID)
	checkErr(err)
}
