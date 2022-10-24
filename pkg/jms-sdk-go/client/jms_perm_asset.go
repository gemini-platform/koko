package client

import (
	"fmt"
	"errors"

	"github.com/jumpserver/koko/pkg/jms-sdk-go/model"
	"github.com/jumpserver/koko/pkg/jms-sdk-go/service"
	"github.com/jumpserver/koko/pkg/logger"
)

func (s *Service) GetOrCreateAssetPermission(m map[string]interface{}) (perm *model.AssetPermission, err error) {
	perm, err = s.GetAssetPermissionByName(m["name"].(string))
	if err != nil {
		if errors.Is(err, service.ErrAssetPermissionNotFound) {
			return s.CreateAssetPermission(m)
		}

		logger.Errorf("failed to get asset permission, err:%v\n", err)
		return nil, err
	}

	logger.Debugf("target asset permissionfound, perm:%v\n", perm)
	return perm, err
}

// CreateJmsAssetPerm
// 系统用户可以放在有资产的时候，尝试创建，如果创建不成功，就重试
func (s *Service) CreateAssetPerm(userID, username string) (*model.AssetPermission, error) {
	systemUser, err := s.GetOrCreateAssetSystemUser(fmt.Sprintf("%s-ssh", username))
	if err != nil {
		logger.Errorf("failed to create system user, err:%v\n", err)
		return nil, err
	}

	assetNode, err := s.GetOrCreateAssetNode(username)
	if err != nil {
		logger.Errorf("failed to get asset node, err:%v\n", err)
		return nil, err
	}

	assetPerm, err := s.GetOrCreateAssetPermission(map[string]interface{}{
		"name":         fmt.Sprintf("%s-perm", username),
		"actions":      []string{model.ActionALL},
		"users":        []string{userID},
		"system_users": []string{systemUser.ID},
		"nodes":        []string{assetNode.ID},
	})
	if err != nil {
		logger.Errorf("failed to create asset permission, err:%v\n", err)
		return nil, err
	}

	logger.Debugf("create asset perm success, perm:%v\n", assetPerm)
	return assetPerm, nil
}