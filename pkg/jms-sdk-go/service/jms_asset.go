package service

import (
	"errors"
	"fmt"

	"github.com/jumpserver/koko/pkg/jms-sdk-go/model"
	"github.com/jumpserver/koko/pkg/logger"
)

var ErrAssetNotFound = errors.New("asset not found")

func (s *JMService) GetAssetById(assetId string) (asset model.Asset, err error) {
	url := fmt.Sprintf(AssetDetailURL, assetId)
	_, err = s.authClient.Get(url, &asset)
	return
}

func (s *JMService) GetAssetPlatform(assetId string) (platform model.Platform, err error) {
	url := fmt.Sprintf(AssetPlatFormURL, assetId)
	_, err = s.authClient.Get(url, &platform)
	return
}

func (s *JMService) GetDomainGateways(domainId string) (domain model.Domain, err error) {
	Url := fmt.Sprintf(DomainDetailWithGateways, domainId)
	_, err = s.authClient.Get(Url, &domain)
	return
}

func (s *JMService) GetSystemUserById(systemUserId string) (sysUser model.SystemUser, err error) {
	url := fmt.Sprintf(SystemUserDetailURL, systemUserId)
	_, err = s.authClient.Get(url, &sysUser)
	return
}

func (s *JMService) GetSystemUserAuthById(systemUserId, assetId, userId,
	username string) (sysUser model.SystemUserAuthInfo, err error) {
	url := fmt.Sprintf(SystemUserAuthURL, systemUserId)
	if assetId != "" {
		url = fmt.Sprintf(SystemUserAssetAuthURL, systemUserId, assetId)
	}
	params := map[string]string{
		"username": username,
		"user_id":  userId,
	}
	_, err = s.authClient.Get(url, &sysUser, params)
	return
}

func (s *JMService) CreateAsset(hostname, ip string, port int, platform, domain string, nodes []string) (asset *model.Asset, err error) {
	params := map[string]interface{}{
		"hostname": hostname,
		"ip":       ip,
		"port":     port,
		"platform": platform,
		"domain":   domain,
		"nodes":    nodes,
	}

	resp, err := s.authClient.Post("/api/v1/assets/assets/", params, &asset)
	if err != nil {
		logger.Errorf("failed to create asset, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	logger.Debugf("create asset success, asset:%v, resp:%v\n", asset, resp)
	return
}

func (s *JMService) GetAsset(id string) (asset *model.Asset, err error) {
	resp, err := s.authClient.Get(fmt.Sprintf("/api/v1/assets/assets/%s/", id), &asset)
	if err != nil {
		logger.Errorf("failed to get asset, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	logger.Debugf("get asset success, asset:%v, resp:%v\n", asset, resp)
	return
}

func (s *JMService) UpdateAsset(id, hostname, ip string, port int, platform, domain string, nodes []string) (asset *model.Asset, err error) {
	params := map[string]interface{}{
		"hostname": hostname,
		"ip":       ip,
		"port":     port,
		"platform": platform,
		"domain":   domain,
		"nodes":    nodes,
	}

	resp, err := s.authClient.Put(fmt.Sprintf("/api/v1/assets/assets/%s/", id), params, &asset)
	if err != nil {
		logger.Errorf("failed to update asset, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	logger.Debugf("update asset success, asset:%v, resp:%v\n", asset, resp)
	return
}

func (s *JMService) UpdateAssetName(id, hostname string) (asset *model.Asset, err error) {
	params := map[string]interface{}{
		"hostname": hostname,
	}

	resp, err := s.authClient.Patch(fmt.Sprintf("/api/v1/assets/assets/%s/", id), params, &asset)
	if err != nil {
		logger.Errorf("failed to update asset name, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	logger.Debugf("update asset name success, asset:%v, resp:%v\n", asset, resp)
	return
}

func (s *JMService) ActiveAsset(id string, active bool) (asset *model.Asset, err error) {
	params := map[string]interface{}{
		"is_active": active,
	}

	resp, err := s.authClient.Patch(fmt.Sprintf("/api/v1/assets/assets/%s/", id), params, asset)
	if err != nil {
		logger.Errorf("failed to update asset active, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	logger.Debugf("update asset active success, id:%v, active: %v, resp:%v\n", id, active, resp)
	return
}

func (s *JMService) DeleteAsset(id string) (err error) {
	resp, err := s.authClient.Delete(fmt.Sprintf("/api/v1/assets/assets/%s/", id), nil)
	if err != nil {
		logger.Errorf("failed to delete asset, err:%v, resp:%v\n", err, resp)
		return err
	}

	logger.Debugf("delete asset success, id:%v, resp:%v\n", id, resp)
	return
}

func (s *JMService) GetAssetByHostname(hostname string) (asset *model.Asset, err error) {
	params := map[string]string{
		"hostname": hostname,
	}

	var assets []model.Asset
	resp, err := s.authClient.Get("/api/v1/assets/assets/", &assets, params)
	if err != nil {
		logger.Errorf("failed to list asset by name, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	if len(assets) == 0 {
		logger.Errorf("asset can not fond by name, err:%v, resp:%v\n", err, resp)
		return asset, ErrAssetNotFound
	} else {
		asset = &assets[0]
	}

	logger.Debugf("get asset success, asset:%v, resp:%v\n", asset, resp)
	return
}

func (s *JMService) GetOrCreateAsset(hostname, ip string, port int, platform, domain string, nodes []string) (*model.Asset, error) {
	asset, err := s.GetAssetByHostname(hostname)
	if err != nil {
		if errors.Is(err, ErrAssetNotFound) {
			return s.CreateAsset(hostname, ip, port, platform, domain, nodes)
		}

		logger.Errorf("failed to get asset: %v\n", err)
		return nil, err
	}

	logger.Debugf("target asset found, asset:%v\n", asset)
	return asset, nil
}

