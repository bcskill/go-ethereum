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

package algorand

import "fmt"

const ProtocolName = "algorand"
const Version = 0x1

// Supported versions of the eth protocol (first is primary).
var ProtocolVersions = []uint{Version}

// Number of implemented message corresponding to different protocol versions.
var ProtocolLengths = []uint64{16}

const ProtocolMaxMsgSize = 10 * 1024 * 1024 // Maximum cap on the size of a protocol message

type errCode int

const (
	ErrMsgTooLarge = iota
	ErrDecode
	ErrInvalidMsgCode
	ErrProtocolVersionMismatch
	ErrNoStatusMsg
	ErrExtraHandshakeMsg
	ErrSuspendedPeer
)

func (e errCode) String() string {
	return errorToString[int(e)]
}

// XXX change once legacy code is out
var errorToString = map[int]string{
	ErrMsgTooLarge:             "Message too long",
	ErrDecode:                  "Invalid message",
	ErrInvalidMsgCode:          "Invalid message code",
	ErrProtocolVersionMismatch: "Protocol version mismatch",
	ErrNoStatusMsg:             "No status message",
	ErrExtraHandshakeMsg:       "Extra handshake message",
	ErrSuspendedPeer:           "Suspended peer",
}

func errResp(code errCode, format string, v ...interface{}) error {
	return fmt.Errorf("%v - %v", code, fmt.Sprintf(format, v...))
}
