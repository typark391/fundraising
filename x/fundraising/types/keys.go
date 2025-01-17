package types

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (
	// ModuleName defines the module name
	ModuleName = "fundraising"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for the fundraising module
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_fundraising"
)

var (
	AuctionIdKey = []byte{0x11} // key for the latest auction id
	SequenceKey  = []byte{0x12} // key for the latest sequence number

	AuctionKeyPrefix    = []byte{0x21} // key for auction id to find an auction
	AuctioneerKeyPrefix = []byte{0x22} // key for auctineer address to find an auction id

	SequenceKeyPrefix = []byte{0x31} // key for auction id with the sequence number to find the bid
	BidKeyPrefix      = []byte{0x32} // key for auction id with bidder address to find the sequence number
	BidderKeyPrefix   = []byte{0x33} // key for bidder address to find the sequence number
)

// GetAuctionKey returns key/value indexing key of the auction.
func GetAuctionKey(auctionID uint64) []byte {
	return append(AuctionKeyPrefix, sdk.Uint64ToBigEndian(auctionID)...)
}

// GetAuctioneerKey returns key/value indexing key of the auction.
func GetAuctioneerKey(auctioneerAcc sdk.AccAddress) []byte {
	return append(AuctioneerKeyPrefix, address.MustLengthPrefix(auctioneerAcc)...)
}

func GetSequenceKey(auctionID uint64, sequence uint64) []byte {
	return append(append(SequenceKeyPrefix, sdk.Uint64ToBigEndian(auctionID)...), sdk.Uint64ToBigEndian(sequence)...)
}

func GetBidKey(auctionID uint64, auctioneerAcc sdk.AccAddress) []byte {
	return append(append(BidKeyPrefix, sdk.Uint64ToBigEndian(auctionID)...), address.MustLengthPrefix(auctioneerAcc)...)
}

func GetBidderKey(auctioneerAcc sdk.AccAddress) []byte {
	return append(BidderKeyPrefix, address.MustLengthPrefix(auctioneerAcc)...)
}

func ParseSequenceKey(key []byte) (auctionID uint64, sequence uint64) {
	if !bytes.HasPrefix(key, SequenceKeyPrefix) {
		panic("key does not have proper prefix")
	}
	// auctionID = sdk.BigEndianToUint64(key[2:])
	// sequence = sdk.BigEndianToUint64(key[2:])
	return
}

func ParseBidKey(key []byte) (auctionID uint64, auctioneerAcc sdk.AccAddress) {
	if !bytes.HasPrefix(key, BidKeyPrefix) {
		panic("key does not have proper prefix")
	}
	// TODO: not implemented yet
	return
}
