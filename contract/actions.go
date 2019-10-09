package contract

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/opentracing/opentracing-go"
	"github.com/videocoin/cloud-pkg/stream"
)

func (c *ContractClient) ValidateProof(ctx context.Context, streamContractAddress string, profileId, inputChunkId int64) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "ValidateProof")
	defer span.Finish()

	stream, err := stream.NewStream(common.HexToAddress(streamContractAddress), c.client)
	if err != nil {
		return fmt.Errorf("failed to create new stream: %s", err.Error())
	}

	tx, err := stream.ValidateProof(c.transactOpts, big.NewInt(profileId), big.NewInt(inputChunkId))
	if err != nil {
		return fmt.Errorf("failed to validate proof: %s", err.Error())
	}

	err = c.waitMinedAndCheck(tx)
	if err != nil {
		return fmt.Errorf("failed to wait mined: %s", err.Error())
	}

	return nil
}
