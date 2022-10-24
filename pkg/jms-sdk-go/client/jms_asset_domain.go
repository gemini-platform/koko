package client

import (
	"github.com/jumpserver/koko/pkg/logger"
	"github.com/jumpserver/koko/pkg/jms-sdk-go/model"
)

func (s *Service) GetDomain() *model.Domain {
	if s.domain == nil {
		logger.Fatal("domain is nil, please sync domain first.")
	}

	return s.domain
}

// SyncDomain 同步自己所在集群的 domain 和 gateway
func (s *Service) SyncDomain(getGateways func() []model.Gateway) {
	domains, err := s.ListAssetDomain()
	if err != nil {
		logger.Fatal("failed to list asset domain.")
	}

	logger.Debugf("list asset domain success, domains:%v\n", domains)
	for _, domain := range domains {
		if domain.Name == s.clusterId {
			logger.Debugf("current asset domain found, domain: %v, cluster_id:%v\n", domain, s.clusterId)
			s.domain = &domain
			break
		}
	}

	if s.domain == nil {
		logger.Debugf("current asset domain not found, create it, cluster_id:%v\n", s.clusterId)
		d, err := s.CreateAssetDomain(s.clusterId)
		if err != nil {
			logger.Fatal("failed to create asset domain, err: ", err.Error())
		}
		d.Gateways = make([]model.Gateway, 0)
		s.domain = d
		logger.Debugf("create asset domain success, domain: %v, cluster_id:%v\n", d, s.clusterId)
	}

	gateways := getGateways()
	logger.Debugf("initialize asset gateway, gateways:%v, cluster_id:%v\n", gateways, s.clusterId)

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
		g, err := s.CreateAssetGateway(
			gateway.Name,
			gateway.IP,
			gateway.Port,
			s.domain.ID,
			gateway.Username,
			gateway.PrivateKey,
			gateway.PublicKey,
		)
		if err != nil {
			logger.Warn("failed to create asset gateway, err: ", err.Error())
			continue
		}
		logger.Debugf("create gateway success, gateway:%v, cluster_id:%v\n", g, s.clusterId)
		s.domain.Gateways = append(s.domain.Gateways, *g)
	}
	logger.Debugf("set up domain success, domain:%v, cluster_id:%v\n", s.domain, s.clusterId)
}
