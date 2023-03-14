package e2e

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/gogo/protobuf/proto"
	"github.com/ory/dockertest/v3/docker"

	oracletypes "github.com/umee-network/umee/v4/x/oracle/types"
	"github.com/umee-network/umee/v4/x/uibc"
)

func (s *IntegrationTestSuite) umeeREST() string {
	return fmt.Sprintf("http://%s", s.valResources[0].GetHostPort("1317/tcp"))
}

func (s *IntegrationTestSuite) gaiaREST() string {
	return fmt.Sprintf("http://%s", s.gaiaResource.GetHostPort("1317/tcp"))
}

func (s *IntegrationTestSuite) connectIBCChains() {
	s.T().Logf("connecting %s and %s chains via IBC", s.chain.id, gaiaChainID)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	exec, err := s.dkrPool.Client.CreateExec(docker.CreateExecOptions{
		Context:      ctx,
		AttachStdout: true,
		AttachStderr: true,
		Container:    s.hermesResource.Container.ID,
		User:         "root",
		Cmd: []string{
			"hermes",
			"create",
			"channel",
			s.chain.id,
			gaiaChainID,
			"--port-a=transfer",
			"--port-b=transfer",
		},
	})
	s.Require().NoError(err)

	var (
		outBuf bytes.Buffer
		errBuf bytes.Buffer
	)

	err = s.dkrPool.Client.StartExec(exec.ID, docker.StartExecOptions{
		Context:      ctx,
		Detach:       false,
		OutputStream: &outBuf,
		ErrorStream:  &errBuf,
	})
	s.Require().NoErrorf(
		err,
		"failed connect chains; stdout: %s, stderr: %s", outBuf.String(), errBuf.String(),
	)

	s.Require().Containsf(
		errBuf.String(),
		"successfully opened init channel",
		"failed to connect chains via IBC: %s", errBuf.String(),
	)

	s.T().Logf("connected %s and %s chains via IBC", s.chain.id, gaiaChainID)
}

func (s *IntegrationTestSuite) sendIBC(srcChainID, dstChainID, recipient string, token sdk.Coin) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	s.T().Logf("sending %s from %s to %s (%s)", token, srcChainID, dstChainID, recipient)
	cmd := []string{
		"hermes",
		"tx",
		"raw",
		"ft-transfer",
		dstChainID,
		srcChainID,
		"transfer",  // source chain port ID
		"channel-0", // since only one connection/channel exists, assume 0
		token.Amount.String(),
		fmt.Sprintf("--denom=%s", token.Denom),
		"--timeout-height-offset=1000",
	}

	if len(recipient) != 0 {
		cmd = append(cmd, fmt.Sprintf("--receiver=%s", recipient))
	}

	exec, err := s.dkrPool.Client.CreateExec(docker.CreateExecOptions{
		Context:      ctx,
		AttachStdout: true,
		AttachStderr: true,
		Container:    s.hermesResource.Container.ID,
		Cmd:          cmd,
	})
	s.Require().NoError(err)

	var (
		outBuf bytes.Buffer
		errBuf bytes.Buffer
	)

	err = s.dkrPool.Client.StartExec(exec.ID, docker.StartExecOptions{
		Context:      ctx,
		Detach:       false,
		OutputStream: &outBuf,
		ErrorStream:  &errBuf,
	})

	s.Require().NoErrorf(
		err,
		"failed to send IBC tokens; stdout: %s, stderr: %s", outBuf.String(), errBuf.String(),
	)
	s.T().Log("successfully sent IBC tokens")
	s.T().Log("Waiting for 12 seconds to make sure trasaction is processed or include in the block")
	time.Sleep(time.Second * 12)
}

// queryREST make http query to grpc-web endpoint and tries to decode valPtr using proto-JSON
// decoder if valPtr implements proto.Message. Otherwise standard JSON decoder is used.
// valPtr must be a pointer.
func queryREST(endpoint string, valPtr interface{}) error {
	resp, err := http.Get(endpoint)
	if err != nil {
		return fmt.Errorf("failed to execute HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("tx query returned non-200 status: %d (%s)", resp.StatusCode, endpoint)
	}

	if valProto, ok := valPtr.(proto.Message); ok {
		bz, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %w, endpoint: %s", err, endpoint)
		}
		if err = cdc.UnmarshalJSON(bz, valProto); err != nil {
			return fmt.Errorf("failed to protoJSON.decode response body: %w, endpoint: %s", err, endpoint)
		}
	} else {
		if err := json.NewDecoder(resp.Body).Decode(valPtr); err != nil {
			return fmt.Errorf("failed to json.decode response body: %w, endpoint: %s", err, endpoint)
		}
	}

	return nil
}

func queryUmeeTx(endpoint, txHash string) error {
	endpoint = fmt.Sprintf("%s/cosmos/tx/v1beta1/txs/%s", endpoint, txHash)
	var result map[string]interface{}
	if err := queryREST(endpoint, &result); err != nil {
		return err
	}

	txResp := result["tx_response"].(map[string]interface{})
	if v := txResp["code"]; v.(float64) != 0 {
		return fmt.Errorf("tx %s failed with status code %v", txHash, v)
	}
	return nil
}

func queryUmeeAllBalances(endpoint, addr string) (sdk.Coins, error) {
	endpoint = fmt.Sprintf("%s/cosmos/bank/v1beta1/balances/%s", endpoint, addr)
	var balancesResp banktypes.QueryAllBalancesResponse
	if err := queryREST(endpoint, &balancesResp); err != nil {
		return nil, err
	}

	return balancesResp.Balances, nil
}

func queryTotalSupply(endpoint string) (sdk.Coins, error) {
	endpoint = fmt.Sprintf("%s/cosmos/bank/v1beta1/supply", endpoint)
	var balancesResp banktypes.QueryTotalSupplyResponse
	if err := queryREST(endpoint, &balancesResp); err != nil {
		return nil, err
	}

	return balancesResp.Supply, nil
}

func queryExchangeRate(endpoint, denom string) (sdk.DecCoins, error) {
	endpoint = fmt.Sprintf("%s/umee/oracle/v1/denoms/exchange_rates/%s", endpoint, denom)
	var resp oracletypes.QueryExchangeRatesResponse
	if err := queryREST(endpoint, &resp); err != nil {
		return nil, err
	}

	return resp.ExchangeRates, nil
}

func queryHistAvgPrice(endpoint, denom string) (sdk.Dec, error) {
	endpoint = fmt.Sprintf("%s/umee/historacle/v1/avg_price/%s", endpoint, strings.ToUpper(denom))
	var resp oracletypes.QueryAvgPriceResponse
	if err := queryREST(endpoint, &resp); err != nil {
		return sdk.Dec{}, err
	}

	return resp.Price, nil
}

func queryOutflows(endpoint, denom string) (sdk.Dec, error) {
	endpoint = fmt.Sprintf("%s/umee/uibc/v1/outflows?denom=%s", endpoint, denom)
	var resp uibc.QueryOutflowsResponse
	if err := queryREST(endpoint, &resp); err != nil {
		return sdk.Dec{}, err
	}

	return resp.Amount, nil
}

func queryUmeeDenomBalance(endpoint, addr, denom string) (sdk.Coin, error) {
	endpoint = fmt.Sprintf("%s/cosmos/bank/v1beta1/balances/%s/by_denom?denom=%s", endpoint, addr, denom)
	var resp banktypes.QueryBalanceResponse
	if err := queryREST(endpoint, &resp); err != nil {
		return sdk.Coin{}, err
	}

	return *resp.Balance, nil
}

func (s *IntegrationTestSuite) queryUmeeBalance(
	umeeValIdx int,
	umeeTokenDenom string,
) (umeeBalance sdk.Coin, umeeAddr string) {
	umeeEndpoint := fmt.Sprintf("http://%s", s.valResources[umeeValIdx].GetHostPort("1317/tcp"))
	umeeAddress, err := s.chain.validators[umeeValIdx].keyInfo.GetAddress()
	s.Require().NoError(err)
	umeeAddr = umeeAddress.String()

	umeeBalance, err = queryUmeeDenomBalance(umeeEndpoint, umeeAddr, umeeTokenDenom)
	s.Require().NoError(err)
	s.T().Logf(
		"Umee Balance of tokens validator; index: %d, addr: %s, amount: %s, denom: %s",
		umeeValIdx, umeeAddr, umeeBalance.String(), umeeTokenDenom,
	)

	return umeeBalance, umeeAddr
}
