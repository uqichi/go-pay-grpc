package service

import (
	"context"
	"fmt"
	"github.com/uqichi/payjp/proto"
	"os"

	"github.com/payjp/payjp-go/v1"
)

type PayManager struct {
	apiKey string
}

func NewPayManager() *PayManager {
	return &PayManager {
		apiKey: os.Getenv("PAYJP_API_KEY"),
	}
}

func (m *PayManager) Charge(ctx context.Context, req *proto.ChargeRequest) ( *proto.ChargeResponse, error) {
	pay := payjp.New(m.apiKey, nil)

	charge, err := pay.Charge.Create(int(req.Amount), payjp.Charge{
		Currency: "jpy",
		CardToken: req.Token,
		Capture: true,
		Description: fmt.Sprintf("%s:%s", req.Name, req.Description),
		Metadata: nil,
	})
	if err != nil {
		return nil, err
	}

	res := &proto.ChargeResponse{
		Paid: charge.Paid,
		Amount: int32(charge.Amount),
	}

	return res, nil
}
