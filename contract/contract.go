package contract

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/videocoin/cloud-pkg/bcops"
	sm "github.com/videocoin/cloud-pkg/streamManager"
)

type ContractClientOpts struct {
	RPCNodeHTTPAddr string
	ContractAddr    string

	Key    string
	Secret string

	Logger *logrus.Entry
}

type ContractClient struct {
	client        *ethclient.Client
	streamManager *sm.Manager
	transactOpts  *bind.TransactOpts

	logger *logrus.Entry
}

func NewContractClient(opts *ContractClientOpts) (*ContractClient, error) {
	client, err := ethclient.Dial(opts.RPCNodeHTTPAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to dial eth client: %s", err.Error())
	}

	address := common.HexToAddress(opts.ContractAddr)
	manager, err := sm.NewManager(address, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create smart contract stream manager: %s", err.Error())
	}

	transactOpts, err := getTransactOpts(context.Background(), client, opts.Key, opts.Secret)
	if err != nil {
		return nil, fmt.Errorf("failed to get transact opts: %s", err.Error())
	}

	return &ContractClient{
		client:        client,
		streamManager: manager,
		transactOpts:  transactOpts,
		logger:        opts.Logger,
	}, nil
}

func getTransactOpts(ctx context.Context, client *ethclient.Client, key, secret string) (*bind.TransactOpts, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "getTransactOpts")
	defer span.Finish()

	decrypted, err := keystore.DecryptKey([]byte(key), secret)
	if err != nil {
		return nil, err
	}

	transactOpts, err := bcops.GetBCAuth(client, decrypted)
	if err != nil {
		return nil, err
	}

	return transactOpts, nil
}

func (c *ContractClient) waitMinedAndCheck(tx *types.Transaction) error {
	cancelCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	receipt, err := bind.WaitMined(cancelCtx, c.client, tx)
	if err != nil {
		return err
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("transaction %s failed", tx.Hash().String())
	}

	return nil
}
