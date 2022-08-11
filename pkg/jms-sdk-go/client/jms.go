package client

import (
	"github.com/jumpserver/koko/pkg/jms-sdk-go/model"
	"github.com/jumpserver/koko/pkg/jms-sdk-go/service"
	"github.com/jumpserver/koko/pkg/logger"
)

type Service struct {
	*service.JMService
	clusterId   string
	domain      *model.Domain
	rootNode    *model.Node
	currentNode *model.Node
}

var svc *Service

func Initial(host, accessKeyID, accessKeySecret, clusterID string) {
	var s *service.JMService
	var err error
	s, err = service.NewAuthJMService(
		service.JMSCoreHost(host),
		service.JMSAccessKey(accessKeyID, accessKeySecret),
	)
	if err != nil {
		logger.Fatal("failed to initialize jm service, err: ", err.Error())
		return
	}

	svc = &Service{s, clusterID, nil, nil, nil}
}

func GetService() *Service {
	return svc
}

func (s *Service) GetDomain() *model.Domain {
	if s.domain == nil {
		logger.Fatal("domain is nil, please sync domain first.")
	}

	return s.domain
}

func (s *Service) GetCurrentNode() *model.Node {
	if s.currentNode == nil {
		logger.Fatal("currnet node is nil, please sync asset node first.")
	}

	return s.currentNode
}

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
			s.rootNode = &node
		}

		if node.Value == s.clusterId {
			s.currentNode = &node
		}
	}

	if s.rootNode == nil {
		logger.Fatal("failed to find root node.")
	}

	if s.currentNode == nil {
		logger.Fatal("failed to find current node, please check cluster id.")
	}
}

// SyncDomain 同步自己所在集群的 domain 和 gateway
func (s *Service) SyncDomain() {
	d, err := s.GetAssetDomain(s.clusterId)
	if err != nil {
		d, err = s.CreateAssetDomain(s.clusterId)
		if err != nil {
			logger.Fatal("failed to create domain, err: ", err.Error())
			return
		}
	}

	s.domain = &d
}
