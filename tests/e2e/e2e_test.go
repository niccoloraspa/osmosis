package e2e

import (
	coretypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/osmosis-labs/osmosis/v7/tests/e2e/chain"
)

func (s *IntegrationTestSuite) TestIBCTokenTransfer() {
	chainA := s.networks[0].GetChain()
	chainB := s.networks[1].GetChain()

	// compare coins of reciever pre and post IBC send
	// diff should only be the amount sent
	s.sendIBC(chainA, chainB, chainB.Validators[0].PublicAddress, chain.OsmoToken)
}

func (s *IntegrationTestSuite) TestStateSync() {
	_, err := s.networks[0].RunValidator(3)
	s.Require().NoError(err)

	doneCondition := func(syncInfo coretypes.SyncInfo) bool {
		return !syncInfo.CatchingUp
	}

	err = s.networks[0].WaitUntil(3, doneCondition)
	s.Require().NoError(err)
}
