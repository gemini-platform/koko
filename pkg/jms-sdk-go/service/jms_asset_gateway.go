package service

import (
	"fmt"

	"github.com/jumpserver/koko/pkg/jms-sdk-go/model"
)

func (s *JMService) CreateAssetGateway(name, ip string, port int, domain, username, privateKey, publicKey string) (gateway model.Gateway, err error) {
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
	_, err = s.authClient.Post("/api/v1/assets/gateways/", params, &gateway)
	return
}

func (s *JMService) GetAssetGateway(id string) (gateway model.Gateway, err error) {
	_, err = s.authClient.Get(fmt.Sprintf("/api/v1/assets/gateways/%s/", id), &gateway)
	return
}

func (s *JMService) UpdateAssetGateway(id, name, ip string, port int, domain, username, privateKey, publicKey string) (gateway model.Gateway, err error) {
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
	_, err = s.authClient.Put(fmt.Sprintf("/api/v1/assets/gateways/%s/", id), params, &gateway)
	return
}

func (s *JMService) DeleteAssetGateway(id string) (err error) {
	_, err = s.authClient.Delete(fmt.Sprintf("/api/v1/assets/gateways/%s/", id), nil)
	return
}
