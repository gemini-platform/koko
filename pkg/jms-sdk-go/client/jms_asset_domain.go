package client

import (
	"github.com/jumpserver/koko/pkg/jms-sdk-go/service"
	"errors"

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
	domain, err := s.GetOrCreateAssetDomain(s.clusterId)
	if err != nil {
		logger.Fatal("failed to get asset domain.")
	}

	s.domain = domain

	if len(domain.Gateways) > 0 {
		// 这里全部移除一遍，防止秘钥更新了不同步
		for _, gateway := range domain.Gateways {
			err = s.DeleteAssetDomain(gateway.ID)
			if err != nil {
				logger.Warnf("delete gateway failed, gateway:%v, cluster_id:%v\n", gateway, s.clusterId)
				continue
			}
		}
	}

	gateways := getGateways()
	logger.Debugf("initialize asset gateway, gateways:%v, cluster_id:%v\n", gateways, s.clusterId)

	for _, gateway := range gateways {
		g, err := s.CreateAssetGateway(
			gateway.Name,
			gateway.IP,
			gateway.Port,
			domain.ID,
			gateway.Username,
			gateway.PrivateKey,
			gateway.PublicKey,
		)
		if err != nil {
			logger.Warn("failed to create asset gateway, err: ", err.Error())
			continue
		}
		logger.Debugf("create gateway success, gateway:%v, cluster_id:%v\n", g, s.clusterId)
		domain.Gateways = append(domain.Gateways, *g)
	}
	logger.Debugf("set up domain success, domain:%v, cluster_id:%v\n", s.domain, s.clusterId)
}

func (s *Service) GetOrCreateAssetDomain(name string) (*model.Domain, error) {
	domain, err := s.GetAssetDomainByName(name)
	if err != nil {
		if errors.Is(err, service.ErrAssetDomainNotFound) {
			return s.CreateAssetDomain(name)
		}

		logger.Errorf("failed to get asset domain, err: %v\n", err)
		return nil, err
	}

	logger.Debugf("target asset domain found, domain:%v\n", domain)
	return domain, nil
}
