package service

import (
	"context"
	"fmt"
	"os"

	"github.com/uqichi/go-pay-grpc/proto"

	"github.com/payjp/payjp-go/v1"
)

type PayService struct {
	apiKey string
}

func NewPayService() *PayService {
	return &PayService{
		apiKey: os.Getenv("PAYJP_API_KEY"),
	}
}

func (s *PayService) Charge(ctx context.Context, req *pb.ChargeRequest) (*pb.ChargeResponse, error) {
	pay := payjp.New(s.apiKey, nil)

	charge, err := pay.Charge.Create(int(req.Amount), payjp.Charge{
		Currency:    "jpy",
		CardToken:   req.Token,
		Capture:     true,
		Description: fmt.Sprintf("%s:%s", req.Name, req.Description),
		Metadata:    nil,
	})
	if err != nil {
		return nil, err
	}

	res := &pb.ChargeResponse{
		Paid:   charge.Paid,
		Amount: int32(charge.Amount),
	}

	return res, nil
}
