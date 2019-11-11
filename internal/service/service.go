package service

import (
	"context"
	"github.com/heyimblake/charbalancer-service/internal/balancer"
	"github.com/heyimblake/charbalancer-service/pkg/api/proto"
)

// Holds Possible Character Groups
var groups = make([]*balancer.CharacterGroup, 4)

func init() {
	groups = append(groups, balancer.AngleBracketsGroup, balancer.CurlyBracketsGroup, balancer.ParenthesisGroup, balancer.SquareBracketsGroup)
}

type BalancerService struct {
	// Empty
}

// New creates and returns a new BalancerService
func New() *BalancerService {
	return &BalancerService{}
}

// BalanceString balances a string from a BalanceStringRequest and returns the result in a BalanceStringResponse
func (service *BalancerService) BalanceString(ctx context.Context, request *proto.BalanceStringRequest) (*proto.BalanceStringResponse, error) {
	// Create Stack
	stack := balancer.NewCharacterGroupStack()

	// Get possible character entries
	entries := stringToCharacterEntries(request.GetData())

	// Push into stack
	for _, entry := range entries {
		if entry == nil {
			continue
		}

		stack.Push(entry)
	}

	// Return Response
	return &proto.BalanceStringResponse{Balanced:stack.IsBalanced()}, nil
}

// stringToCharacterEntries returns an ordered array of CharacterEntry corresponding to the provided string.
// If a character in the string does not have a corresponding CharacterGroup, then it is not included.
func stringToCharacterEntries(str string) []*balancer.CharacterEntry {
	entries := make([]*balancer.CharacterEntry, len(str))

	for _, char := range []rune(str) {
		group := getCharacterGroup(&char)

		if group != nil {
			entries = append(entries, balancer.NewCharacterEntry(char, group))
		}
	}

	return entries
}

// getCharacterGroup returns the corresponding CharacterGroup of the specified character.
func getCharacterGroup(char *rune) *balancer.CharacterGroup {
	for _, group := range groups {
		if group == nil {
			continue
		}

		if group.Opening == *char || group.Closing == *char {
			return group
		}
	}

	return nil
}