package helpers

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"context"
    "log"
)

func ConnectClient(endpoint string) (*ethclient.Client) {
    conn, err := ethclient.Dial(endpoint)
    if err != nil {
        log.Fatal(err)
		return nil
	}
	return conn
}

func GetCurrentHeight(ethConnection *ethclient.Client) (*big.Int) {
    header, err := ethConnection.HeaderByNumber(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
        return big.NewInt(0)
    }
    return header.Number
}
