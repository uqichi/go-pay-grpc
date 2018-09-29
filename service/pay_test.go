package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/uqichi/payjp/proto"
	"testing"
)

func TestPayManager_Charge(t *testing.T) {
	manager := NewPayManager()

	res, err:= manager.Charge(context.Background(), &proto.ChargeRequest{
		Id: 123456,
		Token: "tok_f98454bd3db9be5d41602e06b65a",
		Amount: 98000,
		Name: "Poppy",
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
