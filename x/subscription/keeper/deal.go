package keeper

import (
	"topchain/x/subscription/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"golang.org/x/sync/errgroup"
)

func (k Keeper) SetDeal(ctx sdk.Context, deal types.Deal) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.DealKeyPrefix))

	appendedValue := k.cdc.MustMarshal(&deal)
	store.Set([]byte(deal.Id), appendedValue)
}

func (k Keeper) GetDeal(ctx sdk.Context, dealId string) (deal types.Deal, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.DealKeyPrefix))

	dealBytes := store.Get([]byte(dealId))
	if dealBytes == nil {
		return deal, false
	}

	k.cdc.MustUnmarshal(dealBytes, &deal)
	return deal, true
}

// check if at least one subscription is active
func (k Keeper) IsDealActive(ctx sdk.Context, deal types.Deal) bool {
	for _, subscriptionId := range deal.SubscriptionIds {
		subscription, found := k.GetSubscription(ctx, subscriptionId)
		if !found {
			continue
		}
		if subscription.StartBlock <= uint64(ctx.BlockHeight()) && subscription.EndBlock >= uint64(ctx.BlockHeight()) {
			return true
		}
	}
	return false
}

func (k Keeper) GetAllActiveProviders(ctx sdk.Context, deal types.Deal) []string {
	providers := []string{}
	for _, subscriptionId := range deal.SubscriptionIds {
		subscription, found := k.GetSubscription(ctx, subscriptionId)
		if !found {
			continue
		}
		if subscription.StartBlock <= uint64(ctx.BlockHeight()) && subscription.EndBlock >= uint64(ctx.BlockHeight()) {
			providers = append(providers, subscription.Provider)
		}
	}
	return providers
}

// Need a formula
func (k Keeper) CalculateMinimumStake(ctx sdk.Context, deal types.Deal) int64 {
	return 0
}

func (k Keeper) CalculateBlockReward(ctx sdk.Context, deal types.Deal) int64 {
	remainingBlocks := deal.EndBlock - uint64(ctx.BlockHeight())
	return int64(deal.AvailableAmount) / int64(remainingBlocks)
}

// Iterate over all deals and apply the given callback function
func (k Keeper) IterateDeals(ctx sdk.Context, shouldBreak func(deal types.Deal) bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	iterator := prefix.NewStore(storeAdapter, types.KeyPrefix(types.DealKeyPrefix)).Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var deal types.Deal
		k.cdc.MustUnmarshal(iterator.Value(), &deal)
		if shouldBreak(deal) {
			break
		}
	}
}

// Iterate over all deals for a specific state and apply the given callback function
func (k Keeper) IterateDealsByState(ctx sdk.Context, statePrefix string, shouldBreak func(deal types.Deal) bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	// Create an iterator for the given state prefix
	iterator := prefix.NewStore(storeAdapter, types.KeyPrefix(statePrefix)).Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var deal types.Deal
		k.cdc.MustUnmarshal(iterator.Value(), &deal)
		if shouldBreak(deal) {
			break
		}
	}
}

// IterateDealsByPrefixesAsync concurrently iterates over deals with the provided state prefixes.
// It spawns a goroutine for each prefix and applies the given callback function to the deals.
func (k Keeper) IterateDealsByPrefixesAsync(ctx sdk.Context, statePrefixes []string, shouldBreak func(deal types.Deal) bool) error {
	var g errgroup.Group

	// Iterate over each prefix and spawn a goroutine to process the deals for that prefix
	for _, prefix := range statePrefixes {
		statePrefix := prefix // Capture loop variable

		// Add goroutine to the errgroup
		g.Go(func() error {
			// Call the original IterateDealsByState function for the given prefix
			k.IterateDealsByState(ctx, statePrefix, shouldBreak)
			return nil // Return nil as error if everything goes well
		})
	}

	// Wait for all goroutines to complete and handle any errors
	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

// IterateDealsByPrefixes iterates over deals with the provided state prefixes.
func (k Keeper) IterateDealsByPrefixes(ctx sdk.Context, statePrefixes []string, shouldBreak func(deal types.Deal) bool) error {
	// Iterate over each prefix and spawn a goroutine to process the deals for that prefix
	for _, prefix := range statePrefixes {
		statePrefix := prefix // Capture loop variable
		// Call the original IterateDealsByState function for the given prefix
		k.IterateDealsByState(ctx, statePrefix, shouldBreak)
	}
	return nil
}
