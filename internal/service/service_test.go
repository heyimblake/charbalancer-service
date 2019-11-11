package service

import (
	"github.com/heyimblake/charbalancer-service/pkg/api/proto"
	"testing"
)

func compareBoolean(t **testing.T, actual, expected bool) {
	if actual != expected {
		(*t).Fatalf("Expected %t but got %t.\n", expected, actual)
	}
}

func TestBalancerService_BalanceString(t *testing.T) {
	service := New()
	request := &proto.BalanceStringRequest{Data: "{({})}"}
	actual, _ := service.BalanceString(nil, request)
	compareBoolean(&t, actual.GetBalanced(), true)

	request = &proto.BalanceStringRequest{Data: "(a[bc)"}
	actual, _ = service.BalanceString(nil, request)
	compareBoolean(&t, actual.GetBalanced(), false)
}
