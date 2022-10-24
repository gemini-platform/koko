package client

import (
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