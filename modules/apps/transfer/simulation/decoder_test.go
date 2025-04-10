package simulation_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/cosmos/ibc-go/v10/modules/apps/transfer/simulation"
	"github.com/cosmos/ibc-go/v10/modules/apps/transfer/types"
)

func TestDecodeStore(t *testing.T) {
	dec := simulation.NewDecodeStore()
	denom := types.NewDenom("uatom", types.NewHop("transfer", "channelToA"))

	kvPairs := kv.Pairs{
		Pairs: []kv.Pair{
			{
				Key:   types.PortKey,
				Value: []byte(types.PortID),
			},
			{
				Key:   types.DenomKey,
				Value: types.ModuleCdc.MustMarshal(&denom),
			},
			{
				Key:   []byte{0x99},
				Value: []byte{0x99},
			},
		},
	}
	tests := []struct {
		name        string
		expectedLog string
	}{
		{"PortID", fmt.Sprintf("Port A: %s\nPort B: %s", types.PortID, types.PortID)},
		{"Denom", fmt.Sprintf("Denom A: %s\nDenom B: %s", denom.IBCDenom(), denom.IBCDenom())},
		{"other", ""},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if i == len(tests)-1 {
				require.Panics(t, func() { dec(kvPairs.Pairs[i], kvPairs.Pairs[i]) }, tt.name)
			} else {
				require.Equal(t, tt.expectedLog, dec(kvPairs.Pairs[i], kvPairs.Pairs[i]), tt.name)
			}
		})
	}
}
