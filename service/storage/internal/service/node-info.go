package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/code/bottos/service/storage/proto"
)

func (c *StorageService) GetNodeInfos(ctx context.Context, request *storage.AllRequest, response *storage.NodeInfosResponse) error {

	if request == nil {
		response.Code = 0
		return errors.New("Missing storage request")
	}
	fmt.Println("GetNodeInfos")
	nodes, err := c.mgoRepo.CallGetNodeInfos()

	if nodes == nil || err != nil {
		response.Code = 0
		fmt.Println(err)
		return errors.New("Failed CallGetNodeInfos")

	}
	response.NodeList = []*storage.Node{}
	for _, node := range nodes {
		dbTag := &storage.Node{node.NodeId,
			node.NodeIP,
			node.NodePort}
		response.NodeList = append(response.NodeList, dbTag)
	}
	response.Code = 1
	return nil
}
