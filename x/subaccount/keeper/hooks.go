package keeper

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	orderbookmodulekeeper "github.com/sge-network/sge/x/orderbook/keeper"
)

var _ orderbookmodulekeeper.Hook = Keeper{}

func (k Keeper) AfterBettorWin(ctx sdk.Context, bettor sdk.AccAddress, originalAmount, profit math.Int) {
	balance, exists := k.GetBalance(ctx, bettor)
	if !exists {
		return
	}
	err := balance.Unspend(originalAmount)
	if err != nil {
		panic(err)
	}
	// send profits to subaccount owner
	owner, exists := k.GetSubAccountOwner(ctx, bettor)
	if !exists {
		panic("subaccount owner not found")
	}
	err = k.bankKeeper.SendCoins(ctx, bettor, owner, sdk.NewCoins(sdk.NewCoin(k.GetParams(ctx).LockedBalanceDenom, profit)))
	if err != nil {
		panic(err)
	}
	k.SetBalance(ctx, bettor, balance)
}

func (k Keeper) AfterBettorLoss(ctx sdk.Context, bettor sdk.AccAddress, originalAmount math.Int) {
	balance, exists := k.GetBalance(ctx, bettor)
	if !exists {
		return
	}
	err := balance.Unspend(originalAmount)
	if err != nil {
		panic(err)
	}
	err = balance.AddLoss(originalAmount)
	if err != nil {
		panic(err)
	}
	k.SetBalance(ctx, bettor, balance)
}

func (k Keeper) AfterBettorRefund(ctx sdk.Context, bettor sdk.AccAddress, originalAmount, fee math.Int) {
	balance, exists := k.GetBalance(ctx, bettor)
	if !exists {
		return
	}
	totalUnspent := originalAmount.Add(fee)
	err := balance.Unspend(totalUnspent)
	if err != nil {
		panic(err)
	}
	k.SetBalance(ctx, bettor, balance)
}

func (k Keeper) AfterHouseWin(ctx sdk.Context, house sdk.AccAddress, originalAmount, profit math.Int) {
	// update balance
	balance, exists := k.GetBalance(ctx, house)
	if !exists {
		return
	}

	err := balance.Unspend(originalAmount)
	if err != nil {
		panic(err)
	}
	k.SetBalance(ctx, house, balance)

	// send profits
	subAccountOwner, exists := k.GetSubAccountOwner(ctx, house)
	if !exists {
		panic("data corruption: subaccount owner not found")
	}
	err = k.bankKeeper.SendCoins(ctx, house, subAccountOwner, sdk.NewCoins(sdk.NewCoin(k.GetParams(ctx).LockedBalanceDenom, profit)))
	if err != nil {
		panic(err)
	}
}

func (k Keeper) AfterHouseLoss(ctx sdk.Context, house sdk.AccAddress, originalAmount, lostAmt math.Int) {
	balance, exists := k.GetBalance(ctx, house)
	if !exists {
		return
	}

	err := balance.Unspend(originalAmount)
	if err != nil {
		panic(err)
	}
	err = balance.AddLoss(lostAmt)
	if err != nil {
		panic(err)
	}

	k.SetBalance(ctx, house, balance)
}

func (k Keeper) AfterHouseRefund(ctx sdk.Context, house sdk.AccAddress, originalAmount math.Int) {
	balance, exists := k.GetBalance(ctx, house)
	if !exists {
		return
	}

	err := balance.Unspend(originalAmount)
	if err != nil {
		panic(err)
	}

	k.SetBalance(ctx, house, balance)
}

func (k Keeper) AfterHouseFeeRefund(ctx sdk.Context, house sdk.AccAddress, fee math.Int) {
	balance, exists := k.GetBalance(ctx, house)
	if !exists {
		return
	}

	err := balance.Unspend(fee)
	if err != nil {
		panic(err)
	}

	k.SetBalance(ctx, house, balance)
}