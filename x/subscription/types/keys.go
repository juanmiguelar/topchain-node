package types

const (
	// ModuleName defines the module name
	ModuleName = "subscription"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_subscription"

	SubscriptionKeyPrefix = "Subscription/value"
	
	// old prefix
	DealKeyPrefix         = "Deal/value"
	// DealKeyPrefix defines key prefixes for different deal states
    DealScheduledPrefix   = "Deal/scheduled"
    DealInitialisedPrefix = "Deal/initialised"
    DealActivePrefix      = "Deal/active"
    DealInactivePrefix    = "Deal/inactive"
    DealCancelledPrefix   = "Deal/cancelled"
    DealExpiredPrefix     = "Deal/expired"
)

var (
	ParamsKey = []byte("p_subscription")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
