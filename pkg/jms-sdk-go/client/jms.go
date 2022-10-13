package client

import (
	"github.com/jumpserver/koko/pkg/jms-sdk-go/model"
	"github.com/jumpserver/koko/pkg/jms-sdk-go/service"
	"github.com/jumpserver/koko/pkg/logger"
)

type Service struct {
	*service.JMService
	clusterId string
	domain    *model.Domain
	rootNode  *model.Node
}

var svc *Service

func Initial(host, accessKeyID, accessKeySecret, clusterID, proxy string) {
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

	if proxy != "" {
		s.SetProxy(proxy)
	}

	svc = &Service{s, clusterID, nil, nil}
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
	}

	if s.rootNode == nil {
		logger.Fatal("failed to find root node.")
	}
}

func (s *Service) GetOrCreateAssetNode(name string) (*model.Node, error) {
	nodes, err := s.ListAssetNode()
	if err != nil {
		logger.Error("failed to list asset node, err: ", err.Error())
		return nil, err
	}
	for _, node := range nodes {
		if node.Name == name {
			return &node, nil
		}
	}

	// 不存在就创建一个新的
	newNode, err := s.CreateAssetNode(name)
	if err != nil {
		logger.Error("failed to create asset node, err: ", err.Error())
		return nil, err
	}

	return &newNode, nil
}

// SyncDomain 同步自己所在集群的 domain 和 gateway
func (s *Service) SyncDomain(getGateways func() []model.Gateway) {
	domains, err := s.ListAssetDomain()
	if err != nil {
		logger.Fatal("failed to list asset domain.")
	}

	logger.Info("list asset domain success, domains:", domains)
	for _, domain := range domains {
		if domain.Name == s.clusterId {
			logger.Info("current asset domain found, domain", domain, "cluster_id", s.clusterId)
			s.domain = &domain
			break
		}
	}

	if s.domain == nil {
		logger.Info("current asset domain not found, create it, cluster_id:", s.clusterId)
		d, err := s.CreateAssetDomain(s.clusterId)
		if err != nil {
			logger.Fatal("failed to create asset domain, err: ", err.Error())
		}
		d.Gateways = make([]model.Gateway, 0)
		s.domain = &d
		logger.Info("create asset domain success, domain:", d, "cluster_id:", s.clusterId)
	}

	gateways := getGateways()
	logger.Info("initialize asset gateway, gateways:", gateways, "cluster_id:", s.clusterId)

	toBeCreateGateways := make([]model.Gateway, 0)

filterGatewayLoop:
	for _, gateway := range gateways {
		for _, existGateway := range s.domain.Gateways {
			if existGateway.IP == gateway.IP {
				continue filterGatewayLoop
			}
		}

		toBeCreateGateways = append(toBeCreateGateways, gateway)
	}

	for _, gateway := range toBeCreateGateways {
		g, err := s.CreateAssetGateway(gateway.Name, gateway.IP, gateway.Port, s.domain.ID)
		if err != nil {
			logger.Warn("failed to create asset gateway, err: ", err.Error())
			continue
		}
		logger.Info("create gateway success, gateway:", g, "cluster_id:", s.clusterId)
		s.domain.Gateways = append(s.domain.Gateways, g)
	}
	logger.Info("set up domain success, domain:", s.domain, "cluster_id:", s.clusterId)
}
