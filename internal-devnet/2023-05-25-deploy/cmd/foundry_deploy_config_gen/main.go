package main

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"io/ioutil"
)

type DeployConfigParams struct {
	BatchSenderAddress                common.Address `json:"batchSenderAddress"`
	Controller                        common.Address `json:"controller"`
	FinalSystemOwner                  common.Address `json:"finalSystemOwner"`
	FinalizationPeriodSeconds         uint64         `json:"finalizationPeriodSeconds"`
	GasPriceOracleOverhead            uint64         `json:"gasPriceOracleOverhead"`
	GasPriceOracleScalar              uint64         `json:"gasPriceOracleScalar"`
	L2BlockTime                       uint64         `json:"l2BlockTime"`
	L2GenesisBlockGasLimit            hexutil.Uint64 `json:"l2GenesisBlockGasLimit"`
	L2OutputOracleChallenger          common.Address `json:"l2OutputOracleChallenger"`
	L2OutputOracleProposer            common.Address `json:"l2OutputOracleProposer"`
	L2OutputOracleStartingBlockNumber uint64         `json:"l2OutputOracleStartingBlockNumber"`
	L2OutputOracleSubmissionInterval  uint64         `json:"l2OutputOracleSubmissionInterval"`
	L2OutputOracleStartingTimestamp   uint64         `json:"l2OutputOracleStartingTimestamp"`
	P2PSequencerAddress               common.Address `json:"p2pSequencerAddress"`
	ProxyAdminOwner                   common.Address `json:"proxyAdminOwner"`
}

type MiscConfigParams struct {
	DeployerAddress common.Address `json:"deployerAddress"`
}

type FoundryConfigParams struct {
	BatchSenderAddress                common.Address `json:"batchSenderAddress"`
	Controller                        common.Address `json:"controller"`
	DeployerAddress                   common.Address `json:"deployerAddress"`
	FinalSystemOwner                  common.Address `json:"finalSystemOwner"`
	FinalizationPeriodSeconds         uint64         `json:"finalizationPeriodSeconds"`
	GasPriceOracleOverhead            uint64         `json:"gasPriceOracleOverhead"`
	GasPriceOracleScalar              uint64         `json:"gasPriceOracleScalar"`
	L2BlockTime                       uint64         `json:"l2BlockTime"`
	L2GenesisBlockGasLimit            uint64         `json:"l2GenesisBlockGasLimit"`
	L2OutputOracleChallenger          common.Address `json:"l2OutputOracleChallenger"`
	L2OutputOracleProposer            common.Address `json:"l2OutputOracleProposer"`
	L2OutputOracleStartingBlockNumber uint64         `json:"l2OutputOracleStartingBlockNumber"`
	L2OutputOracleStartingTimestamp   uint64         `json:"l2OutputOracleStartingTimestamp"`
	L2OutputOracleSubmissionInterval  uint64         `json:"l2OutputOracleSubmissionInterval"`
	P2PSequencerAddress               common.Address `json:"p2pSequencerAddress"`
	ProxyAdminOwnerL2                 common.Address `json:"proxyAdminOwner"`
}

type FoundryConfig struct {
	DeployConfig FoundryConfigParams `json:"deployConfig"`
}

func main() {
	deployConfigJsonBytes, err := ioutil.ReadFile("inputs/deploy-config.json")
	if err != nil {
		panic(err)
	}

	var deployConfigParams DeployConfigParams
	err = json.Unmarshal(deployConfigJsonBytes, &deployConfigParams)
	if err != nil {
		panic(err)
	}

	miscConfigJsonByes, err := ioutil.ReadFile("inputs/misc-config.json")
	if err != nil {
		panic(err)
	}
	var miscConfigParams MiscConfigParams
	err = json.Unmarshal(miscConfigJsonByes, &miscConfigParams)
	if err != nil {
		panic(err)
	}

	var foundryConfigParams FoundryConfigParams

	// Params copied from deploy-config.json
	foundryConfigParams.BatchSenderAddress = deployConfigParams.BatchSenderAddress
	foundryConfigParams.Controller = deployConfigParams.Controller
	foundryConfigParams.FinalSystemOwner = deployConfigParams.FinalSystemOwner
	foundryConfigParams.FinalizationPeriodSeconds = deployConfigParams.FinalizationPeriodSeconds
	foundryConfigParams.GasPriceOracleOverhead = deployConfigParams.GasPriceOracleOverhead
	foundryConfigParams.GasPriceOracleScalar = deployConfigParams.GasPriceOracleScalar
	foundryConfigParams.L2BlockTime = deployConfigParams.L2BlockTime
	foundryConfigParams.L2GenesisBlockGasLimit = uint64(deployConfigParams.L2GenesisBlockGasLimit)
	foundryConfigParams.L2OutputOracleChallenger = deployConfigParams.L2OutputOracleChallenger
	foundryConfigParams.L2OutputOracleProposer = deployConfigParams.L2OutputOracleProposer
	foundryConfigParams.L2OutputOracleStartingBlockNumber = deployConfigParams.L2OutputOracleStartingBlockNumber
	foundryConfigParams.L2OutputOracleStartingTimestamp = deployConfigParams.L2OutputOracleStartingTimestamp
	foundryConfigParams.L2OutputOracleSubmissionInterval = deployConfigParams.L2OutputOracleSubmissionInterval
	foundryConfigParams.P2PSequencerAddress = deployConfigParams.P2PSequencerAddress
	foundryConfigParams.ProxyAdminOwnerL2 = deployConfigParams.ProxyAdminOwner

	// Params copied from misc-config.json
	foundryConfigParams.DeployerAddress = miscConfigParams.DeployerAddress

	config := &FoundryConfig{
		DeployConfig: foundryConfigParams,
	}

	file, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		panic(err)
	}

	outputFileName := "inputs/foundry-config.json"
	err = ioutil.WriteFile(outputFileName, file, 0600)
	if err != nil {
		panic(err)
	}
}