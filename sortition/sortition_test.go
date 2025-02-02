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

package sortition

import (
	"crypto/rand"
	"testing"

	"github.com/ethereum/go-ethereum/crypto/ed25519"
)

func TestSortition(t *testing.T) {
	hitcount := uint64(0)
	const nTimes = 10000

	const expectedWeight = 20
	const ownWeight = 100
	const totalWeight = 200
	for i := 0; i < nTimes; i++ {
		var vrfOutput ed25519.VrfOutput256
		_, _ = rand.Read(vrfOutput[:])
		selected := Choose(vrfOutput, ownWeight, 2, expectedWeight, totalWeight)
		hitcount += selected
	}

	expected := uint64(nTimes * expectedWeight / 2)
	var diff uint64
	if expected > hitcount {
		diff = expected - hitcount
	} else {
		diff = hitcount - expected
	}

	// within 2% good enough
	maxDiff := expected / 50
	if diff > maxDiff {
		t.Errorf("wanted %d weight but got %d, diff=%d, maxDiff=%d", expected, hitcount, diff, maxDiff)
	}
}

func BenchmarkSortition(b *testing.B) {
	b.StopTimer()
	keys := make([]ed25519.VrfOutput256, b.N)
	for i := 0; i < b.N; i++ {
		_, _ = rand.Read(keys[i][:])
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Choose(keys[i], 1000000, 2, 2500, 1000000000000)
	}
}
