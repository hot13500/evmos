package osmosis_test

import (
	"testing"
	"fmt"

	osmosisoutpost "github.com/evmos/evmos/v14/precompiles/outposts/osmosis"
	"github.com/stretchr/testify/require"
)

func TestCreateMemo(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name        string
		outputDenom string
		receiver    string
		contract	string
		slippage_percentage string
		window_seconds uint64
		expPass     bool
		errContains string
	}{
		{
			name:     "fail - slippage percentage cannot be negative",
			outputDenom: "uosmo",
			receiver: "receiveraddress",
			contract: "xcscontract",
			slippage_percentage: "-5",
			window_seconds: 10,
			expPass:  false,
		},
		{
			name:     "fail - slippage percentage cannot be higher than max slipppage",
			outputDenom: "uosmo",
			receiver: "receiveraddress",
			contract: "xcscontract",
			slippage_percentage: fmt.Sprint(osmosisoutpost.MaxSlippage + 1),
			window_seconds: 10,
			expPass:  false,
		},
		{
			name:     "fail - window seconds cannot be higher than max window seconds",
			outputDenom: "uosmo",
			receiver: "receiveraddress",
			contract: "xcscontract",
			slippage_percentage: "10",
			window_seconds: 61,
			expPass:  false,
		},
		{
			name:     "success - create memo",
			outputDenom: "uosmo",
			receiver: "receiveraddress",
			contract: "xcscontract",
			slippage_percentage: "5",
			window_seconds: 10,
			expPass:  true,
		},
	}

	for _, tc := range testcases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			memo, err := osmosisoutpost.CreateMemo(tc.outputDenom, tc.receiver, tc.contract, tc.slippage_percentage, tc.window_seconds)
			if tc.expPass {
				require.NoError(t, err, "expected no error while creating memo")
				require.NotEmpty(t, memo, "expected memo not to be empty")
			} else {
				require.Error(t, err, "expected error while creating memo")
				require.Contains(t, err.Error(), tc.errContains, "expected different error")
			}
		})
	}
}