package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	orderbookmodulekeeper "github.com/sge-network/sge/x/orderbook/keeper"
)

var _ orderbookmodulekeeper.Hook = Keeper{}

func (k Keeper) AfterBettorWin(ctx sdk.Context, bettor sdk.AccAddress, originalAmount sdk.Int, profit sdk.Int) {
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
	err = k.bankKeeper.SendCoins(ctx, owner, bettor, sdk.NewCoins(sdk.NewCoin(k.GetParams(ctx).LockedBalanceDenom, originalAmount)))
	if err != nil {
		panic(err)
	}
	k.SetBalance(ctx, bettor, balance)
}

func (k Keeper) AfterBettorLoss(ctx sdk.Context, bettor sdk.AccAddress, originalAmount sdk.Int) {
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

func (k Keeper) AfterBettorRefund(ctx sdk.Context, bettor sdk.AccAddress, originalAmount, fee sdk.Int) {
	balance, exists := k.GetBalance(ctx, bettor)
	if !exists {
		return
	}
	err := balance.Unspend(originalAmount.Add(fee))
	if err != nil {
		panic(err)
	}
	k.SetBalance(ctx, bettor, balance)
}
