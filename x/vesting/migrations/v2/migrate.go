// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)
package v2

import (
	"bytes"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/evmos/v13/x/vesting/types"
)

// MigrateStore migrates the x/vesting module state from the consensus version 1 to
// version 2. Specifically, it adds a new store key to track team accounts subject to
// clawback from governance.
func MigrateStore(
	ctx sdk.Context,
	storeKey storetypes.StoreKey,
) error {
	store := ctx.KVStore(storeKey)
	accAddr := sdk.MustAccAddressFromBech32("evmos19mqtl7pyvtazl85jlre9jltpuff9enjdn9m7hz")
	store.Set(bytes.Join([][]byte{types.KeyPrefixClawbackKey, accAddr.Bytes()}, []byte{}), []byte{0x01})

	return nil
}
