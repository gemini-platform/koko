package service

import (
	"fmt"

	"github.com/jumpserver/koko/pkg/jms-sdk-go/model"
)

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

func (s *JMService) CreateAsset(hostname, ip string, port int, platform, domain string) (asset model.Asset, err error) {
	params := map[string]interface{}{
		"hostname": hostname,
		"ip":       ip,
		"port":     port,
		"platform": platform,
		"domain":   domain,
	}

	_, err = s.authClient.Post("/api/v1/assets/assets/", params, &asset)
	return
}

func (s *JMService) GetAsset(id string) (asset model.Asset, err error) {
	_, err = s.authClient.Get(fmt.Sprintf("/api/v1/assets/assets/%s", id), &asset)
	return
}

func (s *JMService) UpdateAsset(id, hostname, ip string, port int, platform, domain string) (asset model.Asset, err error) {
	params := map[string]interface{}{
		"hostname": hostname,
		"ip":       ip,
		"port":     port,
		"platform": platform,
		"domain":   domain,
	}

	_, err = s.authClient.Put(fmt.Sprintf("/api/v1/assets/assets/%s/", id), params, &asset)
	return
}

func (s *JMService) DeleteAsset(id string) (err error) {
	_, err = s.authClient.Delete(fmt.Sprintf("/api/v1/assets/assets/%s/", id), nil)
	return
}
