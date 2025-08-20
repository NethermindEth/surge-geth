package core

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	taikoGenesis "github.com/ethereum/go-ethereum/core/taiko_genesis"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

var (
	InternalDevnetOntakeBlock = new(big.Int).SetUint64(0)
	PreconfDevnetOntakeBlock  = common.Big0
	MasayaDevnetOntakeBlock   = common.Big0
	HeklaOntakeBlock          = new(big.Int).SetUint64(840_512)
	MainnetOntakeBlock        = new(big.Int).SetUint64(538_304)

	// Surge
	SurgeTestnetOntakeBlock = common.Big1
	SurgeDevnetOntakeBlock  = common.Big1
	SurgeMainnetOntakeBlock = common.Big1

	InternalDevnetPacayaBlock = new(big.Int).SetUint64(0)
	PreconfDevnetPacayaBlock  = common.Big0
	MasayaDevnetPacayaBlock   = common.Big0
	HeklaPacayaBlock          = new(big.Int).SetUint64(1_299_888)
	MainnetPacayaBlock        = new(big.Int).SetUint64(1_166_000)

	// Surge
	SurgeTestnetPacayaBlock = common.Big1
	SurgeDevnetPacayaBlock  = common.Big1
	SurgeMainnetPacayaBlock = common.Big1

	// ContractOwner is the address that will be embedded in the genesis extraData
	ContractOwner = common.HexToAddress("0xdf08f82de32b8d460adbe8d72043e3a7e25a3b39")
)

// generateCliqueExtraData creates the extraData structure
func generateCliqueExtraData(signer common.Address) []byte {
	const extraVanity = 32

	totalSize := extraVanity + common.AddressLength + crypto.SignatureLength
	extraData := make([]byte, totalSize)

	// Signer address starts after vanity
	copy(extraData[extraVanity:extraVanity+common.AddressLength], signer.Bytes())

	return extraData
}

// TaikoGenesisBlock returns the Taiko network genesis block configs.
func TaikoGenesisBlock(networkID uint64) *Genesis {
	var chainConfig *params.ChainConfig
	switch networkID {
	case params.SurgeMainnetNetworkID.Uint64():
		chainConfig = params.SurgeMainnetChainConfig
	case params.SurgeDevnetNetworkID.Uint64():
		chainConfig = params.SurgeDevnetChainConfig
	case params.SurgeTestnetNetworkID.Uint64():
		chainConfig = params.SurgeTestnetChainConfig
	default:
		chainConfig = params.TaikoChainConfig
	}

	var allocJSON []byte
	switch networkID {
	case params.TaikoMainnetNetworkID.Uint64():
		chainConfig.ChainID = params.TaikoMainnetNetworkID
		chainConfig.OntakeBlock = MainnetOntakeBlock
		chainConfig.PacayaBlock = MainnetPacayaBlock
		allocJSON = taikoGenesis.MainnetGenesisAllocJSON
	case params.TaikoInternalL2ANetworkID.Uint64():
		chainConfig.ChainID = params.TaikoInternalL2ANetworkID
		chainConfig.OntakeBlock = InternalDevnetOntakeBlock
		chainConfig.PacayaBlock = InternalDevnetPacayaBlock
		allocJSON = taikoGenesis.InternalL2AGenesisAllocJSON
	case params.TaikoInternalL2BNetworkID.Uint64():
		chainConfig.ChainID = params.TaikoInternalL2BNetworkID
		allocJSON = taikoGenesis.InternalL2BGenesisAllocJSON
	case params.SnaefellsjokullNetworkID.Uint64():
		chainConfig.ChainID = params.SnaefellsjokullNetworkID
		allocJSON = taikoGenesis.SnaefellsjokullGenesisAllocJSON
	case params.AskjaNetworkID.Uint64():
		chainConfig.ChainID = params.AskjaNetworkID
		allocJSON = taikoGenesis.AskjaGenesisAllocJSON
	case params.GrimsvotnNetworkID.Uint64():
		chainConfig.ChainID = params.GrimsvotnNetworkID
		allocJSON = taikoGenesis.GrimsvotnGenesisAllocJSON
	case params.EldfellNetworkID.Uint64():
		chainConfig.ChainID = params.EldfellNetworkID
		allocJSON = taikoGenesis.EldfellGenesisAllocJSON
	case params.JolnirNetworkID.Uint64():
		chainConfig.ChainID = params.JolnirNetworkID
		allocJSON = taikoGenesis.JolnirGenesisAllocJSON
	case params.KatlaNetworkID.Uint64():
		chainConfig.ChainID = params.KatlaNetworkID
		allocJSON = taikoGenesis.KatlaGenesisAllocJSON
	case params.HeklaNetworkID.Uint64():
		chainConfig.ChainID = params.HeklaNetworkID
		chainConfig.OntakeBlock = HeklaOntakeBlock
		chainConfig.PacayaBlock = HeklaPacayaBlock
		allocJSON = taikoGenesis.HeklaGenesisAllocJSON
	case params.PreconfDevnetNetworkID.Uint64():
		chainConfig.ChainID = params.PreconfDevnetNetworkID
		chainConfig.OntakeBlock = PreconfDevnetOntakeBlock
		chainConfig.PacayaBlock = PreconfDevnetPacayaBlock
		allocJSON = taikoGenesis.PreconfDevnetGenesisAllocJSON
	case params.MasayaDevnetNetworkID.Uint64():
		chainConfig.ChainID = params.MasayaDevnetNetworkID
		chainConfig.OntakeBlock = MasayaDevnetOntakeBlock
		chainConfig.PacayaBlock = MasayaDevnetPacayaBlock
		allocJSON = taikoGenesis.MasayaGenesisAllocJSON
	case params.SurgeMainnetNetworkID.Uint64():
		chainConfig.ChainID = params.SurgeMainnetNetworkID
		chainConfig.OntakeBlock = SurgeMainnetOntakeBlock
		chainConfig.PacayaBlock = SurgeMainnetPacayaBlock
		allocJSON = taikoGenesis.SurgeMainnetGenesisAllocJSON
	case params.SurgeTestnetNetworkID.Uint64():
		chainConfig.ChainID = params.SurgeTestnetNetworkID
		chainConfig.OntakeBlock = SurgeTestnetOntakeBlock
		chainConfig.PacayaBlock = SurgeTestnetPacayaBlock
		allocJSON = taikoGenesis.SurgeTestnetGenesisAllocJSON
	case params.SurgeDevnetNetworkID.Uint64():
		chainConfig.ChainID = params.SurgeDevnetNetworkID
		chainConfig.OntakeBlock = SurgeDevnetOntakeBlock
		chainConfig.PacayaBlock = SurgeDevnetPacayaBlock
		allocJSON = taikoGenesis.SurgeDevnetGenesisAllocJSON
	default:
		chainConfig.ChainID = params.TaikoInternalL2ANetworkID
		chainConfig.OntakeBlock = InternalDevnetOntakeBlock
		chainConfig.PacayaBlock = InternalDevnetPacayaBlock
		allocJSON = taikoGenesis.InternalL2AGenesisAllocJSON
	}

	var alloc GenesisAlloc
	if err := alloc.UnmarshalJSON(allocJSON); err != nil {
		log.Crit("unmarshal alloc json error", "error", err)
	}

	extraData := generateCliqueExtraData(ContractOwner)

	return &Genesis{
		Config:     chainConfig,
		ExtraData:  extraData,
		GasLimit:   uint64(30_000_000),
		Difficulty: common.Big0,
		Alloc:      alloc,
		GasUsed:    0,
	}
}
