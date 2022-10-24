package client

import (
	"errors"

	"github.com/jumpserver/koko/pkg/jms-sdk-go/model"
	"github.com/jumpserver/koko/pkg/jms-sdk-go/service"
	"github.com/jumpserver/koko/pkg/logger"
)


func (s *Service) GetRootNode() *model.Node {
	if s.rootNode == nil {
		logger.Fatal("root node is nil, please sync asset node first.")
	}

	return s.rootNode
}

// SyncDomain 同步 node 信息
func (s *Service) SyncAssetNode() {
	nodes, err := s.ListAssetNode()
	if err != nil {
		logger.Fatal("failed to list asset node, err: ", err.Error())
		return
	}

	for _, node := range nodes {
		if node.Value == "Default" && node.FullValue == "/Default" {
			logger.Debugf("root node found, node: %v\n", node)
			s.rootNode = &node
		}
	}

	if s.rootNode == nil {
		logger.Fatal("failed to find root node.")
	}
}

func (s *Service) GetOrCreateAssetNode(name string) (*model.Node, error) {
	oldNode, err := s.GetAssetNodeByName(name)
	if err != nil {
		if errors.Is(err, service.ErrAssetNodeNotFound) {
			// 不存在就创建一个新的
			return s.CreateAssetNode(name)
		}

		logger.Errorf("failed to get asset node, err: %v\n", err)
		return nil, err
	}

	logger.Debugf("target asset node found, node:%v\n", oldNode)
	return oldNode, nil
}
