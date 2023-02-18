package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"

	"github.com/sge-network/sge/utils"
)

const (
	// ModuleName is the name of the house module
	ModuleName = "house"

	// StoreKey is the string store representation
	StoreKey = ModuleName

	// QuerierRoute is the querier route for the house module
	QuerierRoute = ModuleName

	// RouterKey is the msg router key for the house module
	RouterKey = ModuleName
)

// module accounts constants
const (
	// HouseParticipationFeeName defines the account name for house participation fee
	HouseParticipationFeeName = "house_participation_fee_pool"
)

var (
	DepositKeyPrefix    = []byte{0x00} // prefix for keys that store deposits
	WithdrawalKeyPrefix = []byte{0x01} // prefix for keys that store withdrawals
)

// GetDepositKey creates the key for deposit bond with sport event and participant
func GetDepositKey(depAddr sdk.AccAddress, sportEventUID string, participantID uint64) []byte {
	return append(GetDepositsKey(depAddr), append(utils.StrBytes(sportEventUID), utils.Uint64ToBytes(participantID)...)...)
}

// GetDepositsKey creates the key for deposit bond with sport event
func GetDepositsKey(depAddr sdk.AccAddress) []byte {
	return address.MustLengthPrefix(depAddr)
}

// GetWithdrawalKey creates the key for withdrawal bond with sport event and deposit
func GetWithdrawalKey(depAddr sdk.AccAddress, sportEventUID string, participantID uint64, withdrawalNumber uint64) []byte {
	return append(GetDepositKey(depAddr, sportEventUID, participantID), utils.Uint64ToBytes(withdrawalNumber)...)
}

// GetWithdrawalsKey creates the key for withdrawals bond with sport event
func GetWithdrawalsKey(depAddr sdk.AccAddress) []byte {
	return address.MustLengthPrefix(depAddr)
}