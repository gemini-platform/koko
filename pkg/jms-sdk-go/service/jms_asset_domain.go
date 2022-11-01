package service

import (
	"errors"
	"fmt"

	"github.com/jumpserver/koko/pkg/logger"
	"github.com/jumpserver/koko/pkg/jms-sdk-go/model"
)

var ErrAssetDomainNotFound = errors.New("asset domain found")

func (s *JMService) ListAssetDomain() (domains []model.Domain, err error) {
	resp, err := s.authClient.Get("/api/v1/assets/domains/", &domains)
	if err != nil {
		logger.Errorf("failed to list asset domain, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	logger.Debugf("list asset domain success, domains:%v, resp:%v\n", domains, resp)
	return domains, nil
}

func (s *JMService) CreateAssetDomain(name string) (domain *model.Domain, err error) {
	params := map[string]string{
		"name": name,
	}
	resp, err := s.authClient.Post("/api/v1/assets/domains/", params, &domain)
	if err != nil {
		logger.Errorf("failed to create asset domain, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	logger.Debugf("create asset domain success, domain:%v, resp:%v\n", domain, resp)
	return domain, nil
}

func (s *JMService) GetAssetDomainByName(name string) (domain *model.Domain, err error) {
	params := map[string]string{
		"name": name,
	}

	var domains []model.Domain
	resp, err := s.authClient.Get("/api/v1/assets/domains/", &domains, params)
	if err != nil {
		logger.Errorf("failed to list asset domain by name, err:%v, resp:%v\n", err, resp)
		return domain, err
	}

	if len(domains) == 0 {
		logger.Errorf("asset domain can not fond by name, err:%v, resp:%v\n", err, resp)
		return domain, ErrAssetDomainNotFound
	} else {
		domain = &domains[0]
	}

	logger.Debugf("get asset domain by name success, domain:%v, resp:%v\n", domain, resp)
	return
}

func (s *JMService) GetAssetDomain(id string) (domain *model.Domain, err error) {
	resp, err := s.authClient.Get(fmt.Sprintf("/api/v1/assets/domains/%s/", id), &domain)
	if err != nil {
		logger.Errorf("failed to get asset domain, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	logger.Debugf("get asset domain success, domain:%v, resp:%v\n", domain, resp)
	return domain, nil
}

func (s *JMService) UpdateAssetDomain(id, name string) (domain *model.Domain, err error) {
	params := map[string]string{
		"name": name,
	}
	resp, err := s.authClient.Put(fmt.Sprintf("/api/v1/assets/domains/%s/", id), params, &domain)
	if err != nil {
		logger.Errorf("failed to update asset domain, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	logger.Debugf("update asset domain success, domain:%v, resp:%v\n", domain, resp)
	return domain, nil
}

func (s *JMService) DeleteAssetDomain(id string) (err error) {
	resp, err := s.authClient.Delete(fmt.Sprintf("/api/v1/assets/domains/%s/", id), nil)
	if err != nil {
		logger.Errorf("failed to delete asset domain, err:%v, resp:%v\n", err, resp)
		return err
	}

	logger.Debugf("delete asset domain success, domain:%s, resp:%v\n", id, resp)
	return nil
}
