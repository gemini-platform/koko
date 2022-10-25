package service

import (
	"fmt"

	"github.com/jumpserver/koko/pkg/logger"
	"github.com/jumpserver/koko/pkg/jms-sdk-go/model"
)

func (s *JMService) CreateAssetGateway(name, ip string, port int, domain, username, privateKey, publicKey string) (gateway *model.Gateway, err error) {
	params := map[string]interface{}{
		"name":        name,
		"ip":          ip,
		"port":        port,
		"protocol":    model.ProtocolSSH,
		"domain":      domain,
		"username":    username,
		"private_key": privateKey,
		"public_key":  publicKey,
	}
	resp, err := s.authClient.Post("/api/v1/assets/gateways/", params, &gateway)
	if err != nil {
		logger.Errorf("failed to create asset gateway, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	logger.Debugf("create asset gateway success, gateway:%v, resp:%v\n", gateway, resp)
	return gateway, nil
}

func (s *JMService) GetAssetGateway(id string) (gateway *model.Gateway, err error) {
	resp, err := s.authClient.Get(fmt.Sprintf("/api/v1/assets/gateways/%s/", id), gateway)
	if err != nil {
		logger.Errorf("failed to get asset gateway, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	logger.Debugf("get asset gateway success, gateway:%v, resp:%v\n", gateway, resp)
	return gateway, nil
}

func (s *JMService) UpdateAssetGateway(id, name, ip string, port int, domain, username, privateKey, publicKey string) (gateway *model.Gateway, err error) {
	params := map[string]interface{}{
		"name":        name,
		"ip":          ip,
		"port":        port,
		"protocol":    model.ProtocolSSH,
		"domain":      domain,
		"username":    username,
		"private_key": privateKey,
		"public_key":  publicKey,
	}
	resp, err := s.authClient.Put(fmt.Sprintf("/api/v1/assets/gateways/%s/", id), params, &gateway)
	if err != nil {
		logger.Errorf("failed to update asset gateway, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	logger.Debugf("update asset gateway success, gateway:%v, resp:%v\n", gateway, resp)
	return gateway, nil
}

func (s *JMService) DeleteAssetGateway(id string) (err error) {
	resp, err := s.authClient.Delete(fmt.Sprintf("/api/v1/assets/gateways/%s/", id), nil)
	if err != nil {
		logger.Errorf("failed to delete asset gateway, err:%v, resp:%v\n", err, resp)
		return err
	}

	logger.Debugf("delete asset gateway success, gateway:%s, resp:%v\n", id, resp)
	return nil
}
