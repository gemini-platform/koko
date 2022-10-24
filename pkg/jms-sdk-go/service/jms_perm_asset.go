package service

import (
	"errors"
	"fmt"

	"github.com/jumpserver/koko/pkg/jms-sdk-go/model"
	"github.com/jumpserver/koko/pkg/logger"
)

var ErrAssetPermissionNotFound = errors.New("asset permission not found")

func (s *JMService) SearchPermAsset(userId, key string) (res model.AssetList, err error) {
	Url := fmt.Sprintf(UserPermsAssetsURL, userId)
	payload := map[string]string{"search": key}
	_, err = s.authClient.Get(Url, &res, payload)
	return
}

func (s *JMService) GetSystemUsersByUserIdAndAssetId(userId, assetId string) (sysUsers []model.SystemUser, err error) {
	Url := fmt.Sprintf(UserPermsAssetSystemUsersURL, userId, assetId)
	_, err = s.authClient.Get(Url, &sysUsers)
	return
}

func (s *JMService) GetAllUserPermsAssets(userId string) ([]map[string]interface{}, error) {
	var params model.PaginationParam
	res, err := s.GetUserPermsAssets(userId, params)
	if err != nil {
		return nil, err
	}
	return res.Data, nil
}

func (s *JMService) GetUserPermsAssets(userID string, params model.PaginationParam) (resp model.PaginationResponse, err error) {
	Url := fmt.Sprintf(UserPermsAssetsURL, userID)
	return s.getPaginationResult(Url, params)
}

func (s *JMService) RefreshUserAllPermsAssets(userId string) ([]map[string]interface{}, error) {
	var params model.PaginationParam
	params.Refresh = true
	res, err := s.GetUserPermsAssets(userId, params)
	if err != nil {
		return nil, err
	}
	return res.Data, nil
}

func (s *JMService) GetUserAssetByID(userId, assetId string) (assets []model.Asset, err error) {
	params := map[string]string{
		"id": assetId,
	}
	Url := fmt.Sprintf(UserPermsAssetsURL, userId)
	_, err = s.authClient.Get(Url, &assets, params)
	return
}

func (s *JMService) GetUserPermAssetsByIP(userId, assetIP string) (assets []model.Asset, err error) {
	params := map[string]string{
		"ip": assetIP,
	}
	reqUrl := fmt.Sprintf(UserPermsAssetsURL, userId)
	_, err = s.authClient.Get(reqUrl, &assets, params)
	return
}

func (s *JMService) CreateAssetPermission(params map[string]interface{}) (perm *model.AssetPermission, err error) {
	resp, err := s.authClient.Post("/api/v1/perms/asset-permissions/", params, &perm)
	if err != nil {
		logger.Errorf("failed to create asset permission, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	logger.Debugf("create asset permission success, perm:%v, resp:%v\n", perm, resp)
	return
}

func (s *JMService) GetAssetPermissionByName(name string) (perm *model.AssetPermission, err error) {
	params := map[string]string{
		"name": name,
	}
	var perms []model.AssetPermission
	resp, err := s.authClient.Get("/api/v1/perms/asset-permissions/", &perms, params)
	if err != nil {
		logger.Errorf("failed to get asset permission by name, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	if len(perms) < 1 {
		return perm, ErrAssetPermissionNotFound
	} else {
		perm = &perms[0]
	}

	logger.Debugf("get asset permission by name success, perm:%v, resp:%v\n", perm, resp)
	return
}

func (s *JMService) GetAssetPermission(id string) (perm *model.AssetPermission, err error) {
	resp, err := s.authClient.Get(fmt.Sprintf("/api/v1/perms/asset-permissions/%s/", id), &perm)
	if err != nil {
		logger.Errorf("failed to get asset permission, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	logger.Debugf("get asset permission success, perm:%v, resp:%v\n", perm, resp)
	return
}

func (s *JMService) UpdateAssetPermission(id, params map[string]interface{}) (perm *model.AssetPermission, err error) {
	resp, err := s.authClient.Put(fmt.Sprintf("/api/v1/perms/asset-permissions/%s/", id), params, &perm)
	if err != nil {
		logger.Errorf("failed to update asset permission, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	logger.Debugf("update asset permission success, perm:%v, resp:%v\n", perm, resp)
	return
}

func (s *JMService) DeleteAssetPermission(id string) (err error) {
	resp, err := s.authClient.Delete(fmt.Sprintf("/api/v1/perms/asset-permissions/%s/", id), nil)
	if err != nil {
		logger.Errorf("failed to delete asset permission, err:%v, resp:%v\n", err, resp)
		return err
	}

	logger.Debugf("delete asset permission success, id:%s, resp:%v\n", id, resp)
	return
}
