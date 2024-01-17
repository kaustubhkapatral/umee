package ante_test

import (
	"fmt"
	"testing"

	tmrand "github.com/cometbft/cometbft/libs/rand"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/stretchr/testify/suite"

	umeeapp "github.com/umee-network/umee/v6/app"
)

type IntegrationTestSuite struct {
	suite.Suite
	app       *umeeapp.UmeeApp
	ctx       sdk.Context
	clientCtx client.Context
	txBuilder client.TxBuilder
}

func (s *IntegrationTestSuite) SetupTest() {
	app := umeeapp.Setup(s.T())
	ctx := app.BaseApp.NewContextLegacy(false, tmproto.Header{
		ChainID: fmt.Sprintf("test-chain-%s", tmrand.Str(4)),
		Height:  1,
	})

	s.app = app
	s.ctx = ctx

	encodingConfig := testutil.MakeTestEncodingConfig()
	encodingConfig.Amino.RegisterConcrete(&testdata.TestMsg{}, "testdata.TestMsg", nil)
	testdata.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	s.clientCtx = client.Context{}.
		WithTxConfig(encodingConfig.TxConfig)
}

// CreateTestTx is a helper function to create a tx given multiple inputs.
func (suite *IntegrationTestSuite) CreateTestTx(privs []cryptotypes.PrivKey, accNums, accSeqs []uint64, chainID string) (xauthsigning.Tx, error) {
	defaultSignMode, err := authsigning.APISignModeToInternal(suite.clientCtx.TxConfig.SignModeHandler().DefaultMode())
	if err != nil {
		return nil, err
	}

	var sigsV2 []signing.SignatureV2
	for i, priv := range privs {
		sigV2 := signing.SignatureV2{
			PubKey: priv.PubKey(),
			Data: &signing.SingleSignatureData{
				SignMode:  defaultSignMode,
				Signature: nil,
			},
			Sequence: accSeqs[i],
		}

		sigsV2 = append(sigsV2, sigV2)
	}
	err = suite.txBuilder.SetSignatures(sigsV2...)
	if err != nil {
		return nil, err
	}

	sigsV2 = []signing.SignatureV2{}
	for i, priv := range privs {
		signerData := xauthsigning.SignerData{
			ChainID:       chainID,
			AccountNumber: accNums[i],
			Sequence:      accSeqs[i],
		}
		sigV2, err := tx.SignWithPrivKey(suite.ctx,
			defaultSignMode,
			signerData,
			suite.txBuilder, priv, suite.clientCtx.TxConfig, accSeqs[i])
		if err != nil {
			return nil, err
		}

		sigsV2 = append(sigsV2, sigV2)
	}
	err = suite.txBuilder.SetSignatures(sigsV2...)
	if err != nil {
		return nil, err
	}

	return suite.txBuilder.GetTx(), nil
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
