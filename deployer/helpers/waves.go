package helpers

import (
	"bytes"
	"context"
	"deployer/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/wavesplatform/gowaves/pkg/crypto"

	"github.com/wavesplatform/gowaves/pkg/client"
)

const (
	WaitCount = 10
)

func IsUnconfirmedTx(nodeUrl string, id string) (bool, error) {
	_, code, err := sendRequest("GET", nodeUrl+"/transactions/unconfirmed/info/"+id, nil, "")
	if err != nil && code != 404 {
		return true, err
	}
	return code == 200, nil
}

func Broadcast(nodeUrl string, body []byte) (bool, error) {
	_, code, err := sendRequest("POST", nodeUrl+"/transactions/broadcast", body, "")
	if err != nil && code != 404 {
		return true, err
	}
	return code == 200, nil
}

func GetStateByAddress(nodeUrl string, address string) (map[string]models.State, error) {
	rsBody, _, err := sendRequest("GET", nodeUrl+"/addresses/data/"+address, nil, "")
	if err != nil {
		return nil, err
	}
	states := models.States{}
	if err := json.Unmarshal(rsBody, &states); err != nil {
		return nil, err
	}
	return states.Map(), nil
}

func WaitTx(client *client.Client, id *crypto.Digest, ctx context.Context) <-chan error {
	out := make(chan error)

	go func() {
		defer close(out)
		for i := 0; i < WaitCount; i++ {
			un, err := IsUnconfirmedTx(client.GetOptions().BaseUrl, id.String())
			if err != nil {
				out <- err
				break
			}

			if un == false {
				tx, _, err := client.Transactions.Info(ctx, *id)
				if err != nil {
					out <- err
				}
				if id, _ := tx.GetID(); len(id) == 0 {
					out <- errors.New("transaction not found")
				} else {
					out <- nil
				}
				break
			}

			if i == (WaitCount - 1) {
				out <- errors.New("transaction not found")
				break
			}

			time.Sleep(time.Second)
		}
	}()
	return out
}

func sendRequest(method string, url string, rqBody []byte, apiKey string) ([]byte, int, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(rqBody))
	req.Header.Add("content-type", "application/json")
	if apiKey != "" {
		req.Header.Add("X-API-Key", apiKey)
	}

	if err != nil {
		return nil, 0, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		if resp != nil {
			return nil, resp.StatusCode, err
		} else {
			return nil, 520, err
		}
	}

	defer resp.Body.Close()
	rsBody, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return rsBody, resp.StatusCode, errors.New(string(rsBody))
	}
	return rsBody, resp.StatusCode, nil
}
