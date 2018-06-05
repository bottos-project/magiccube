// Copyright 2017~2022 The Bottos Authors
// This file is part of the Bottos Chain library.
// Created by Rocket Core Team of Bottos.

//This program is free software: you can distribute it and/or modify
//it under the terms of the GNU General Public License as published by
//the Free Software Foundation, either version 3 of the License, or
//(at your option) any later version.

//This program is distributed in the hope that it will be useful,
//but WITHOUT ANY WARRANTY; without even the implied warranty of
//MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU General Public License for more details.

//You should have received a copy of the GNU General Public License
// along with bottos.  If not, see <http://www.gnu.org/licenses/>.

/*
 * file description:  general Hash type
 * @Author: Gong Zibin
 * @Date:   2017-12-12
 * @Last Modified by:
 * @Last Modified time:
 */

package common

import (
	"fmt"
	"testing"
)

func TestMerkleRootHash_Odd(t *testing.T) {
	var hs []Hash
	hs = append(hs, Sha256([]byte("1")))
	hs = append(hs, Sha256([]byte("2")))
	hs = append(hs, Sha256([]byte("3")))
	hs = append(hs, Sha256([]byte("4")))
	hs = append(hs, Sha256([]byte("5")))

	root := ComputeMerkleRootHash(hs)
	fmt.Printf("root hash: %x\n", root)
}

func TestMerkleRootHash_Even(t *testing.T) {
	var hs []Hash
	hs = append(hs, Sha256([]byte("1")))
	hs = append(hs, Sha256([]byte("2")))
	hs = append(hs, Sha256([]byte("3")))
	hs = append(hs, Sha256([]byte("4")))

	root := ComputeMerkleRootHash(hs)
	fmt.Printf("root hash: %x\n", root)
}

func TestMerkleRootHash_NULL(t *testing.T) {
	var hs []Hash

	root := ComputeMerkleRootHash(hs)
	fmt.Printf("empty root hash: %x\n", root)
}

// https://btc.com/0000000000000000003c4601e87ab5389c15d9c48c37076472aa7d45cdd4fa30
// bitcoin hash is of little-endian notation, Reverse before and after computing
func TestMerkleRootHash_BTC_0000000000000000003c4601e87ab5389c15d9c48c37076472aa7d45cdd4fa30(t *testing.T) {
	var hs []Hash
	hs = append(hs, bitcoinHashConvert("2fc7d439472c12c27b8eff3a758c01772d4ae0ed2d6139fc4bbddac9b8943c3c"))

	root := ComputeMerkleRootHash(hs)

	rootStr := reverseHash(root).ToHexString()
	expectedRootStr := "2fc7d439472c12c27b8eff3a758c01772d4ae0ed2d6139fc4bbddac9b8943c3c"
	if rootStr != expectedRootStr {
		fmt.Printf("fail, root: %s, expected: %s\n", rootStr, expectedRootStr)
	} else {
		fmt.Printf("suc, root: %s, expected: %s\n", rootStr, expectedRootStr)
	}
}

// https://btc.com/00000000000010e45e3d943559acb2be2323fceb24182e811eb4dffcf4b1f6c8
func TestMerkleRootHash_BTC_00000000000010e45e3d943559acb2be2323fceb24182e811eb4dffcf4b1f6c8(t *testing.T) {
	h1 := bitcoinHashConvert("a6cc01018576e05c253a47492574133e0572a7743b6d80ab16e21a71080ba15c")
	h2 := bitcoinHashConvert("0a3428333988f10759e017a64724501898d4bc37ed7887f09d719ba7a15c9d36")

	var hs []Hash
	hs = append(hs, h1)
	hs = append(hs, h2)

	root := ComputeMerkleRootHash(hs)

	rootStr := reverseHash(root).ToHexString()
	expectedRootStr := "ed0733aaeb0c0ec41b1f368c537ec32bea711c789d6d44ea1452077300f401a1"
	if rootStr != expectedRootStr {
		fmt.Printf("fail, root: %s, expected: %s\n", rootStr, expectedRootStr)
	} else {
		fmt.Printf("suc, root: %s, expected: %s\n", rootStr, expectedRootStr)
	}

}

// https://btc.com/000000000000000000010f01933e19182a0d2f3f134bfd6559eff2395436e5ec
func TestMerkleRootHash_BTC_000000000000000000010f01933e19182a0d2f3f134bfd6559eff2395436e5ec(t *testing.T) {
	var hs []Hash

	hs = append(hs, bitcoinHashConvert("683e4b3a66feeb66b7bf4d18678bf5ff859f8ae697744aca3794b986d831297a"))
	hs = append(hs, bitcoinHashConvert("c19de3f8f74bdf32825bc9f2475c322e43af4635041551091c675cbdc748c48a"))
	hs = append(hs, bitcoinHashConvert("514d24e75fcb68a00240bcf98248544944edaa3df9a90533edd7c81e4010d236"))

	root := ComputeMerkleRootHash(hs)

	rootStr := reverseHash(root).ToHexString()
	expectedRootStr := "c03c955ae773acf0d052a15a8a95487d29d2c210449e3094007b286a3b4d50f3"
	if rootStr != expectedRootStr {
		fmt.Printf("fail, root: %s, expected: %s\n", rootStr, expectedRootStr)
	} else {
		fmt.Printf("suc, root: %s, expected: %s\n", rootStr, expectedRootStr)
	}

}

//https://btc.com/000000000000030de89e7729d5785c4730839b6e16ea9fb686a54818d3860a8d
func TestMerkleRootHash_BTC_000000000000030de89e7729d5785c4730839b6e16ea9fb686a54818d3860a8d(t *testing.T) {

	h1 := bitcoinHashConvert("338bbd00b893c384eb2b11e70f3875447297c5f20815499e787867df4538e48d")
	h2 := bitcoinHashConvert("1ad1138c6064dd17d0a4d12016d629c18f15fc9d1472412945f9c91a696689c7")
	h3 := bitcoinHashConvert("c77834d14d66729014b06fcf45c5f82af4bdd9d816e787f01bfa525cfa147014")
	h4 := bitcoinHashConvert("bb3d83398d7517fe643b2421d795e73c342b6a478ef53acdaab35dbdffbbcdb5")
	h5 := bitcoinHashConvert("38d563caf0e9ed103515cab09e40e49da0ccb8c0765ce304f9556e5bc62e8ff5")
	h6 := bitcoinHashConvert("8fc0507359d0122fa14b5887034d857bd69c8bc0e74c8dd428c2fc098595c285")
	h7 := bitcoinHashConvert("9db9fe6d011c1c7e997418aeec78ccb659648cfc915b2ff1154cabb41359ac70")
	h8 := bitcoinHashConvert("3c72fdb7e38e4437faa9e5789df6b51505de014b062361ef47578244d5025628")

	var hs []Hash
	hs = append(hs, h1)
	hs = append(hs, h2)
	hs = append(hs, h3)
	hs = append(hs, h4)
	hs = append(hs, h5)
	hs = append(hs, h6)
	hs = append(hs, h7)
	hs = append(hs, h8)

	// compute merkle root hash
	root := ComputeMerkleRootHash(hs)

	rootStr := reverseHash(root).ToHexString()
	expectedRootStr := "acb5aeb11e2a607e610b90f2722cf68aec719af2a2fd6a6af179764e90169af4"
	if rootStr != expectedRootStr {
		fmt.Printf("fail, root: %s, expected: %s\n", rootStr, expectedRootStr)
	} else {
		fmt.Printf("suc, root: %s, expected: %s\n", rootStr, expectedRootStr)
	}

}

// https://btc.com/000000000000000000165f317acde835bc398f5186e987775afff9c43824acca
func TestMerkleRootHash_BTC_000000000000000000165f317acde835bc398f5186e987775afff9c43824acca(t *testing.T) {
	var hs []Hash

	hs = append(hs, bitcoinHashConvert("a83b0582908acb22486275260d3ca991286c50fcafa76e0a8c1137548ae1d0a1"))
	hs = append(hs, bitcoinHashConvert("517a29a5f480c05e5cb9af348b7d272f8936b1d1497d0fd211db3adec9f51d74"))
	hs = append(hs, bitcoinHashConvert("663460bd6e649dd4b44f936965d443b08f7d57c92ad1602f5f597c121ae2c894"))
	hs = append(hs, bitcoinHashConvert("8a1e01a40281a70ea9e5f57c39541524852121bdccededb0112603a8661dfc31"))
	hs = append(hs, bitcoinHashConvert("5c4c917f63d19451248ad16235403db2886d107172bbd8110069e6dad0cc235e"))
	hs = append(hs, bitcoinHashConvert("f1e9902bf4b829ef32c4cdb59e6870f1c7bcfd2d26c954342db4bfc517b65fc4"))
	hs = append(hs, bitcoinHashConvert("5d84c6b9b3a1ddaaf4fbbe6e1ef444cf8ec1bbf1282d4af77937fac4e52381fe"))
	hs = append(hs, bitcoinHashConvert("34ac7eb254d6f10f22a375fb44c982c3de4fb7cc2ed92d020529d06ebb81db39"))
	hs = append(hs, bitcoinHashConvert("a0824b281baabcf010ae4957d3a3280dd708c7f8e698588a2e9b5402a0cafc56"))
	hs = append(hs, bitcoinHashConvert("4615467e6cb481933d6aeda17cc9ac80c20b2e64cf03dedbe48c6dbe7c31f3c3"))
	hs = append(hs, bitcoinHashConvert("645a88620e10180b13c66f909ac34ab48df3c18c8367ee9f761a12035591b537"))
	hs = append(hs, bitcoinHashConvert("16769fb8f7364fad44c22d23cbca4acc3675b67613f3977c8a48a4f2edaa6a8d"))
	hs = append(hs, bitcoinHashConvert("a65fa2ef4db8f925c576ce71ecab9833f229f1cc6e8102e8eb93609bd413f1f0"))
	hs = append(hs, bitcoinHashConvert("5457ddb54b370413b7d7b7654b5adcff40d9d679525ed4100b11ae0ba5b705af"))
	hs = append(hs, bitcoinHashConvert("76685bc2ce2dcb13390d50097671e8eaef927f0694577cad3dd51e89c5b2e332"))
	hs = append(hs, bitcoinHashConvert("f4b60176f3a49fa8af6bd6c4e2c7d60928cbe08db7eb449888b53d6c44626b30"))
	hs = append(hs, bitcoinHashConvert("4978ce9c3211298acaada46a66b161d1c16476790e8dd3cb973add1d57830f42"))

	root := ComputeMerkleRootHash(hs)

	rootStr := reverseHash(root).ToHexString()
	expectedRootStr := "8361ee29a7f28a80d06c475f2160a2ee5c52fa6eb8de93098672145dd8a05ac3"
	if rootStr != expectedRootStr {
		fmt.Printf("fail, root: %s, expected: %s\n", rootStr, expectedRootStr)
	} else {
		fmt.Printf("suc, root: %s, expected: %s\n", rootStr, expectedRootStr)
	}
}

func reverse(a []byte) []byte {
	b := make([]byte, len(a))
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		b[i], b[opp] = a[opp], a[i]
	}
	return b
}

func reverseHash(h Hash) Hash {
	return BytesToHash(reverse(h[:]))
}

func bitcoinHashConvert(hashstr string) Hash {
	h := HexToHash(hashstr)
	h = reverseHash(h)
	return h
}
