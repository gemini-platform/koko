package client

import (
	"errors"

	"github.com/jumpserver/koko/pkg/jms-sdk-go/model"
	"github.com/jumpserver/koko/pkg/jms-sdk-go/service"
	"github.com/jumpserver/koko/pkg/logger"
)

func (s *Service) GetOrCreateUser(m map[string]interface{}) (user *model.User, err error) {
	user, err = s.GetUserByUserName(m["username"].(string))
	if err != nil {
		if errors.Is(err, service.ErrAssetPermissionNotFound) {
			return s.CreateUser(m)
		}

		logger.Errorf("failed to get user, err:%v\n", err)
		return nil, err
	}

	logger.Debugf("target user found, user:%v\n", user)
	return user, err
}