// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://7bed18c87054f807bc9096501bc78f737363f357af831791bab07c4fa6c5a1a67cdcf0a097dc2cc918262ef04fb1c05c26026df5c11a6a56666f9b1fb4072210@18.178.30.66:37888",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
var TestnetBootnodes = []string{
	"enode://d4d3b5d143f3e75cc0c68fd9bb40013deddbc72f1977637a71958784c83c94780f0c8c3217b6afb74af181388c500b3fbc4270036351bf323de5bbbc9274315d@47.241.95.187:37888",
}

var V5Bootnodes = []string{}

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
