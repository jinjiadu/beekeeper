package api

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// PSSService represents Bee's PSS service
type PSSService service

// Sends a PSS message to a recipienct with a specific topic
func (p *PSSService) SendMessage(ctx context.Context, nodeAddress swarm.Address, nodePublicKey string, topic string, prefix int, data io.Reader, batchID string) error {
	h := http.Header{}
	h.Add(postageStampBatchHeader, batchID)

	url := fmt.Sprintf("/%s/pss/send/%s/%s?recipient=%s", apiVersion, topic, nodeAddress.String()[:prefix], nodePublicKey)

	return p.client.requestWithHeader(ctx, http.MethodPost, url, h, data, nil)
}
