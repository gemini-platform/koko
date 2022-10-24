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