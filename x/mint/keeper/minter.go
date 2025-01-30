package keeper

import (
	"context"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sge-network/sge/x/mint/types"
)

// StakingTokenSupply implements an alias call to the underlying staking keeper's
// to be used in BeginBlocker.
func (k Keeper) StakingTokenSupply(ctx context.Context) (sdkmath.Int, error) {
	return k.stakingKeeper.StakingTokenSupply(ctx)
}

// TokenSupply implements an alias call to the underlying bank keeper's
// to be used in BeginBlocker.
func (k Keeper) TokenSupply(ctx context.Context, denom string) sdkmath.Int {
	return k.bankKeeper.GetSupply(ctx, denom).Amount
}

// BondedRatio implements an alias call to the underlying staking keeper's
// to be used in BeginBlocker.
func (k Keeper) BondedRatio(ctx context.Context) (sdkmath.LegacyDec, error) {
	return k.stakingKeeper.BondedRatio(ctx)
}

// MintCoins implements an alias call to the underlying supply keeper's
// to be used in BeginBlocker.
func (k Keeper) MintCoins(ctx context.Context, newCoins sdk.Coins) error {
	if newCoins.Empty() {
		// skip as no coins need to be minted
		return nil
	}

	return k.bankKeeper.MintCoins(ctx, types.ModuleName, newCoins)
}

// AddCollectedFees implements an alias call to the underlying supply keeper's
// to be used in BeginBlocker.
func (k Keeper) AddCollectedFees(ctx context.Context, fees sdk.Coins) error {
	return k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, k.feeCollectorName, fees)
}
