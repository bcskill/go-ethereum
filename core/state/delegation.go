// Copyright (c) 2019 The ethereum Authors
// This file is part of ethereum
//
// ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with ethereum. If not, see <https://www.gnu.org/licenses/>.

package state

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts"
	"github.com/ethereum/go-ethereum/crypto"
)

func idSlotKey(id *common.Address) common.Hash {
	return common.BytesToHash(
		crypto.Keccak256(append(id.Hash().Bytes(), common.HexToHash("0x0").Bytes()...)),
	)
}

func addrSlotKey(addr *common.Address) common.Hash {
	return common.BytesToHash(
		crypto.Keccak256(append(addr.Hash().Bytes(), common.HexToHash("0x1").Bytes()...)),
	)
}

// GetTotalReceived 取抵押给id的总数量
func GetTotalReceived(statedb *StateDB, id common.Address) *big.Int {
	return statedb.GetState(contracts.DelegationAddress, idSlotKey(&id)).Big()
}

// GetTotalFund 取地址已抵押的总数量
func GetTotalFund(statedb *StateDB, addr common.Address) *big.Int {
	return statedb.GetState(contracts.DelegationAddress, addrSlotKey(&addr)).Big()
}

// GetBalanceWithFund returns a value summed by balance and fund
func GetBalanceWithFund(statedb *StateDB, addr common.Address) *big.Int {
	balance := new(big.Int).Set(statedb.GetBalance(addr))
	balance.Add(balance, GetTotalFund(statedb, addr))

	return balance
}
