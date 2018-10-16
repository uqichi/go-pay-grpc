package service

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/uqichi/go-pay-grpc/proto"
)

func TestPayService_Charge(t *testing.T) {
	service := NewPayService()

	res, err := service.Charge(context.Background(), &pb.ChargeRequest{
		Id:          "8eba8568-8898-4b6f-9b50-d9670375bc0e",
		Token:       "tok_f98454bd3db9be5d41602e06b65a",
		Amount:      98000,
		Name:        "Poppy",
		Description: "fancy delicious candy",
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v\n", func(v interface{}) string {
		b, _ := json.MarshalIndent(v, "-", "\t")
		return string(b)
	}(res))
}
