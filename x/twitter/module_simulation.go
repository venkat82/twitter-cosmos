package twitter

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github/venkat/twitter/testutil/sample"
	twittersimulation "github/venkat/twitter/x/twitter/simulation"
	"github/venkat/twitter/x/twitter/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = twittersimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgAddTweet = "op_weight_msg_add_tweet"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddTweet int = 100

	opWeightMsgAddFollower = "op_weight_msg_add_follower"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddFollower int = 100

	opWeightMsgFetchFeed = "op_weight_msg_fetch_feed"
	// TODO: Determine the simulation weight value
	defaultWeightMsgFetchFeed int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	twitterGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&twitterGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgAddTweet int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddTweet, &weightMsgAddTweet, nil,
		func(_ *rand.Rand) {
			weightMsgAddTweet = defaultWeightMsgAddTweet
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddTweet,
		twittersimulation.SimulateMsgAddTweet(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddFollower int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddFollower, &weightMsgAddFollower, nil,
		func(_ *rand.Rand) {
			weightMsgAddFollower = defaultWeightMsgAddFollower
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddFollower,
		twittersimulation.SimulateMsgAddFollower(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgFetchFeed int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgFetchFeed, &weightMsgFetchFeed, nil,
		func(_ *rand.Rand) {
			weightMsgFetchFeed = defaultWeightMsgFetchFeed
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgFetchFeed,
		twittersimulation.SimulateMsgFetchFeed(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
