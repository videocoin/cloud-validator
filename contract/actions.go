package contract

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/opentracing/opentracing-go"
	"github.com/videocoin/cloud-pkg/stream"
)

func (c *ContractClient) ValidateProof(
	ctx context.Context,
	streamContractAddress string,
	profileID,
	chunkID *big.Int,
) (*ethtypes.Transaction, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "ValidateProof")
	defer span.Finish()

	stream, err := stream.NewStream(common.HexToAddress(streamContractAddress), c.client)
	if err != nil {
		return nil, fmt.Errorf("failed to create new stream: %s", err.Error())
	}

	tx, err := stream.ValidateProof(c.transactOpts, profileID, chunkID)
	if err != nil {
		return nil, fmt.Errorf("failed to validate proof: %s", err.Error())
	}

	err = c.waitMinedAndCheck(tx)
	if err != nil {
		return nil, fmt.Errorf("failed to wait mined: %s", err.Error())
	}

	return tx, nil
}

func (c *ContractClient) ScrapProof(
	ctx context.Context,
	streamContractAddress string,
	profileID,
	chunkID *big.Int,
) (*ethtypes.Transaction, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "ScrapProof")
	defer span.Finish()

	stream, err := stream.NewStream(common.HexToAddress(streamContractAddress), c.client)
	if err != nil {
		return nil, fmt.Errorf("failed to create new stream: %s", err.Error())
	}

	tx, err := stream.ScrapProof(c.transactOpts, profileID, chunkID)
	if err != nil {
		return nil, fmt.Errorf("failed to scrap proof: %s", err.Error())
	}

	err = c.waitMinedAndCheck(tx)
	if err != nil {
		return nil, fmt.Errorf("failed to wait mined: %s", err.Error())
	}

	return tx, nil
}
