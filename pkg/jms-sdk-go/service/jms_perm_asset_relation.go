package service

import (
	"errors"
	"fmt"

	"github.com/jumpserver/koko/pkg/logger"
)

var ErrAssetPermissionAssetRelationNotFound = errors.New("asset permission asset relation not found")

type AssetPermissionAssetRelation struct {
	ID                     int    `json:"id"`
	Asset                  string `json:"asset"`
	AssetDisplay           string `json:"asset_display"`
	AssetPermission        string `json:"asset_permission"`
	AssetPermissionDisplay string `json:"asset_permission_display"`
}

func (s *JMService) CreateAssetPermissionAssetRelation(assetPermission, asset string) (relation *AssetPermissionAssetRelation, err error) {
	params := map[string]string{
		"asset":           asset,
		"assetpermission": assetPermission,
	}

	_, err = s.authClient.Post("/api/v1/perms/asset-permissions-assets-relations/", params, &relation)
	return
}

func (s *JMService) GetAssetPermissionAssetRelationByID(assetPermission, asset string) (relation *AssetPermissionAssetRelation, err error) {
	params := map[string]string{
		"asset":           asset,
		"assetpermission": assetPermission,
	}

	var relations []AssetPermissionAssetRelation
	resp, err := s.authClient.Get("/api/v1/perms/asset-permissions-assets-relations/", &relations, params)
	if err != nil {
		logger.Errorf("failed to get asset permission and asset relation by id, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	if len(relations) < 1 {
		return nil, ErrAssetPermissionNotFound
	} else {
		relation = &relations[0]
	}

	return relation, nil
}

func (s *JMService) DeleteAssetPermissionAssetRelation(id string) (err error) {
	_, err = s.authClient.Delete(fmt.Sprintf("/api/v1/perms/asset-permissions-assets-relations/%s/", id), nil)
	return
}

type AssetPermissionNodeRelation struct {
	ID                     int    `json:"id"`
	Node                   string `json:"node"`
	NodeDisplay            string `json:"node_display"`
	AssetPermission        string `json:"assetpermission"`
	AssetPermissionDisplay string `json:"assetpermission_display"`
}

func (s *JMService) CreateAssetPermissionNodeRelation(assetPermission, node string) (relation *AssetPermissionNodeRelation, err error) {
	params := map[string]string{
		"node":            node,
		"assetpermission": assetPermission,
	}

	_, err = s.authClient.Post("/api/v1/perms/asset-permissions-nodes-relations/", params, &relation)
	return
}

func (s *JMService) DeleteAssetPermissionNodeRelation(id string) (err error) {
	_, err = s.authClient.Delete(fmt.Sprintf("/api/v1/perms/asset-permissions-nodes-relations/%s/", id), nil)
	return
}

type AssetPermissionSystemUserRelation struct {
	ID                     int    `json:"id"`
	SystemUser             string `json:"systemuser"`
	SystemUserDisplay      string `json:"systemuser_display"`
	AssetPermission        string `json:"assetpermission"`
	AssetPermissionDisplay string `json:"assetpermission_display"`
}

func (s *JMService) CreateAssetPermissionSystemUserRelation(assetPermission, systemUser string) (relation *AssetPermissionSystemUserRelation, err error) {
	params := map[string]string{
		"systemuser":      systemUser,
		"assetpermission": assetPermission,
	}

	_, err = s.authClient.Post("/api/v1/perms/asset-permissions-system-users-relations/", params, &relation)
	return
}

func (s *JMService) DeleteAssetPermissionSystemUserRelation(id string) (err error) {
	_, err = s.authClient.Delete(fmt.Sprintf("/api/v1/perms/asset-permissions-system-users-relations/%s/", id), nil)
	return
}

type AssetPermissionUserRelation struct {
	ID                     int    `json:"id"`
	User                   string `json:"user"`
	UserDisplay            string `json:"user_display"`
	AssetPermission        string `json:"assetpermission"`
	AssetPermissionDisplay string `json:"assetpermission_display"`
}

func (s *JMService) CreateAssetPermissionUserRelation(assetPermission, systemUser string) (relation *AssetPermissionUserRelation, err error) {
	params := map[string]string{
		"systemuser":      systemUser,
		"assetpermission": assetPermission,
	}

	_, err = s.authClient.Post("/api/v1/perms/asset-permissions-users-relations/", params, &relation)
	return
}

func (s *JMService) DeleteAssetPermissionUserRelation(id string) (err error) {
	_, err = s.authClient.Delete(fmt.Sprintf("/api/v1/perms/asset-permissions-users-relations/%s/", id), nil)
	return
}