package core

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/vm"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts"
)

var (
	ethereumMainnetGenesisBalance = new(big.Int).Mul(new(big.Int).SetUint64(630720000), common.BigEther)
	ethereumTestnetGenesisBalance = new(big.Int).Mul(new(big.Int).SetUint64(100000000), common.BigEther)
)

var ethereumMainnetAllocData = map[common.Address]GenesisAccount{
	common.HexToAddress("0x45Ec182EDC6774c9A2926172F1Fd996e59b58CED"): {
		Balance: ethereumMainnetGenesisBalance,
	},

	contracts.CreatorAddress: {
		Balance: common.Big0,
		Code:    common.FromHex(contracts.CreatorBinRuntime),
		Storage: map[common.Hash]common.Hash{
			vm.CreatorSlotKey(common.HexToAddress("0x1000000000000000000000000000000000000001")): common.HexToHash("0x0000000000000000000000002000000000000000000000000000000000000001"),
			vm.CreatorSlotKey(common.HexToAddress("0x1000000000000000000000000000000000000002")): common.HexToHash("0x0000000000000000000000002000000000000000000000000000000000000001"),
			vm.CreatorSlotKey(common.HexToAddress("0x1000000000000000000000000000000000000003")): common.HexToHash("0x0000000000000000000000002000000000000000000000000000000000000001"),
			vm.CreatorSlotKey(common.HexToAddress("0x1000000000000000000000000000000000000004")): common.HexToHash("0x0000000000000000000000002000000000000000000000000000000000000001"),
			vm.CreatorSlotKey(common.HexToAddress("0x1000000000000000000000000000000000000005")): common.HexToHash("0x0000000000000000000000002000000000000000000000000000000000000001"),
			vm.CreatorSlotKey(common.HexToAddress("0x1000000000000000000000000000000000000006")): common.HexToHash("0x0000000000000000000000002000000000000000000000000000000000000001"),
			vm.CreatorSlotKey(common.HexToAddress("0x1000000000000000000000000000000000000007")): common.HexToHash("0x0000000000000000000000002000000000000000000000000000000000000001"),
			vm.CreatorSlotKey(common.HexToAddress("0x1000000000000000000000000000000000000008")): common.HexToHash("0x0000000000000000000000002000000000000000000000000000000000000001"),
			vm.CreatorSlotKey(common.HexToAddress("0x1000000000000000000000000000000000000009")): common.HexToHash("0x0000000000000000000000002000000000000000000000000000000000000001"),
		},
	},
	contracts.MinerAddress: {
		Balance: common.Big0,
		Code:    common.FromHex(contracts.MinerBinRuntime),
		Storage: map[common.Hash]common.Hash{
			common.HexToHash("0x06ff3c55f357d4545a14dcc167670bf1dcc8bb45dcd90fa4a085a02a39da3a8a"): common.HexToHash("0x45ec182edc6774c9a2926172f1fd996e59b58ced000000640000000000000001"),
			common.HexToHash("0x06ff3c55f357d4545a14dcc167670bf1dcc8bb45dcd90fa4a085a02a39da3a8b"): common.HexToHash("0xf88a8d844c217531a38d6019ea671652340fe0d899996250bccce13af99933de"),
			common.HexToHash("0x06ff3c55f357d4545a14dcc167670bf1dcc8bb45dcd90fa4a085a02a39da3a8c"): common.HexToHash("0x6e8f4a7c7651766722dd7fb9d7a97cd28678a1cefb12631580a7ffe90a910b8f"),
		},
	},
	contracts.AuthorityAddress: {
		Balance: common.Big0,
		Code:    common.FromHex(contracts.AuthorityBinRuntime),
	},
	contracts.DelegationAddress: {
		Balance: common.Big0,
		Code:    common.FromHex(contracts.DelegationBinRuntime),
	},
}

var ethereumTestnetAllocData = map[common.Address]GenesisAccount{
	common.HexToAddress("0x48F155527f25EB1d4cb2aa32b7e84692AA0025C0"): {
		Balance: ethereumTestnetGenesisBalance,
	},

	contracts.CreatorAddress: {
		Balance: common.Big0,
		Code:    common.FromHex(contracts.CreatorBinRuntime),
		Storage: map[common.Hash]common.Hash{
			vm.CreatorSlotKey(common.HexToAddress("0x1000000000000000000000000000000000000001")): common.HexToHash("0x0000000000000000000000002000000000000000000000000000000000000001"),
			vm.CreatorSlotKey(common.HexToAddress("0x1000000000000000000000000000000000000002")): common.HexToHash("0x0000000000000000000000002000000000000000000000000000000000000001"),
			vm.CreatorSlotKey(common.HexToAddress("0x1000000000000000000000000000000000000003")): common.HexToHash("0x0000000000000000000000002000000000000000000000000000000000000001"),
			vm.CreatorSlotKey(common.HexToAddress("0x1000000000000000000000000000000000000004")): common.HexToHash("0x0000000000000000000000002000000000000000000000000000000000000001"),
			vm.CreatorSlotKey(common.HexToAddress("0x1000000000000000000000000000000000000005")): common.HexToHash("0x0000000000000000000000002000000000000000000000000000000000000001"),
			vm.CreatorSlotKey(common.HexToAddress("0x1000000000000000000000000000000000000006")): common.HexToHash("0x0000000000000000000000002000000000000000000000000000000000000001"),
			vm.CreatorSlotKey(common.HexToAddress("0x1000000000000000000000000000000000000007")): common.HexToHash("0x0000000000000000000000002000000000000000000000000000000000000001"),
			vm.CreatorSlotKey(common.HexToAddress("0x1000000000000000000000000000000000000008")): common.HexToHash("0x0000000000000000000000002000000000000000000000000000000000000001"),
			vm.CreatorSlotKey(common.HexToAddress("0x1000000000000000000000000000000000000009")): common.HexToHash("0x0000000000000000000000002000000000000000000000000000000000000001"),
		},
	},
	contracts.MinerAddress: {
		Balance: common.Big0,
		Code:    common.FromHex(contracts.MinerBinRuntime),
		Storage: map[common.Hash]common.Hash{
			common.HexToHash("0x82f1831a162a7f3e29811d2195e3e69849199225cd98cdaac26cbc717f24fcaf"): common.HexToHash("0x48F155527f25EB1d4cb2aa32b7e84692AA0025C0000000640000000000000001"),
			common.HexToHash("0x82f1831a162a7f3e29811d2195e3e69849199225cd98cdaac26cbc717f24fcb0"): common.HexToHash("0x5acfc834316080bb6158cc3f2ba4abc2b5dfe14865dfaa979aa6d8937bbe21b7"),
			common.HexToHash("0x82f1831a162a7f3e29811d2195e3e69849199225cd98cdaac26cbc717f24fcb1"): common.HexToHash("0xcc126d7ac641ea94a756ad4ddd4bcf92b683521bc92f7a261c21e116e62a9083"),
		},
	},
	contracts.AuthorityAddress: {
		Balance: common.Big0,
		Code:    common.FromHex(contracts.AuthorityBinRuntime),
	},
	contracts.DelegationAddress: {
		Balance: common.Big0,
		Code:    common.FromHex(contracts.DelegationBinRuntime),
	},
}
