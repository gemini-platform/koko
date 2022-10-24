package client

import (
	"errors"

	"github.com/jumpserver/koko/pkg/jms-sdk-go/model"
	"github.com/jumpserver/koko/pkg/jms-sdk-go/service"
	"github.com/jumpserver/koko/pkg/logger"
)

func (s *Service) GetOrCreateAssetSystemUser(name string) (*model.SystemUser, error) {
	user, err := s.GetAssetSystemUserByName(name)
	if err != nil {
		if errors.Is(err, service.ErrAssetSystemUserNotFound) {
			return s.CreateAssetSystemUser(name, "root")
		}

		logger.Errorf("failed to get asset system user, err:%v\n", err)
		return nil, err
	}

	logger.Debugf("target asset system user found, user:%v\n", user)
	return user, err
}