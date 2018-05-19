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
)

type MerkleHashTree struct {
	Root  *MTNode
}

type MTNode struct {
	Hash  Hash
}

func dSha256(h1 Hash, h2 Hash) Hash {
	var data []byte
	data = append(data, h1.Bytes()...)
	data = append(data, h2.Bytes()...)
	t1 := Sha256(data)
	t2 := Sha256(t1[:])
	return t2
}

func CreateMerkleTree(hs []Hash) *MerkleHashTree {
	if len(hs) == 0 {
		return nil
	}

	nodes := createLeafNodes(hs)
	for len(nodes) > 1 {
		nodes = createNextLevel(nodes)
	}
	mt := &MerkleHashTree{
		Root:  nodes[0],
	}
	return mt
}

func createLeafNodes(hs []Hash) []*MTNode {
	var nodes []*MTNode
	for _, h := range hs {
		node := &MTNode{
			Hash: h,
		}
		nodes = append(nodes, node)
	}
	if (len(hs) % 2 == 1) {
		node := &MTNode{
			Hash: hs[len(hs)-1],
		}
		nodes = append(nodes, node)
	}
	return nodes
}

func createNextLevel(nodes []*MTNode) []*MTNode {
	var nlNodes []*MTNode
	for i := 0; i < len(nodes)/2; i++ {
		hash := dSha256(nodes[i*2].Hash, nodes[i*2+1].Hash)
		node := &MTNode{Hash:  hash}
		nlNodes = append(nlNodes, node)
	}
	if (len(nodes) % 2 == 1) {
		hash := dSha256(nodes[len(nodes)-1].Hash, nodes[len(nodes)-1].Hash)
		node := &MTNode{Hash:  hash}
		nlNodes = append(nlNodes, node)
	}
	return nlNodes
}

func ComputeMerkleRootHash(hs []Hash) Hash {
	if len(hs) == 0 {
		return Hash{}
	}
	if len(hs) == 1 {
		return hs[0]
	}

	tree := CreateMerkleTree(hs)
	if tree != nil {
		return tree.Root.Hash
	}

	return Hash{}
}
