package service

import (
	"fmt"

	"github.com/jumpserver/koko/pkg/jms-sdk-go/model"
)

func (s *JMService) CreateAssetDomain(name string) (domain model.Domain, err error) {
	params := map[string]string{
		"name": name,
	}
	_, err = s.authClient.Post("/api/v1/assets/domains/", params, &domain)
	return
}

func (s *JMService) GetAssetDomain(id string) (domain model.Domain, err error) {
	_, err = s.authClient.Get(fmt.Sprintf("/api/v1/assets/domains/%s/", id), &domain)
	return
}

func (s *JMService) UpdateAssetDomain(id, name string) (domain model.Domain, err error) {
	params := map[string]string{
		"name": name,
	}
	_, err = s.authClient.Put(fmt.Sprintf("/api/v1/assets/domains/%s/", id), params, &domain)
	return
}

func (s *JMService) DeleteAssetDomain(id string) (err error) {
	_, err = s.authClient.Delete(fmt.Sprintf("/api/v1/assets/domains/%s/", id), nil)
	return
}
