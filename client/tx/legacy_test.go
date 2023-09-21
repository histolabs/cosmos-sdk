package tx_test

import (
	"github.com/cosmos/cosmos-sdk/v2/testutil/testdata"
	"github.com/cosmos/cosmos-sdk/v2/types"
	banktypes "github.com/cosmos/cosmos-sdk/v2/x/bank/types"
)

const (
	memo          = "waboom"
	timeoutHeight = uint64(5)
)

var (
	_, pub1, addr1 = testdata.KeyTestPubAddr()
	_, _, addr2    = testdata.KeyTestPubAddr()
	rawSig         = []byte("dummy")
	msg1           = banktypes.NewMsgSend(addr1, addr2, types.NewCoins(types.NewInt64Coin("wack", 2)))

	chainID = "test-chain"
)
