package service

import (
	"errors"
	"fmt"

	"github.com/jumpserver/koko/pkg/logger"
	"github.com/jumpserver/koko/pkg/jms-sdk-go/model"
)

var ErrAssetSystemUserNotFound = errors.New("asset system user not found")

func (s *JMService) CreateAssetSystemUser(name, username string) (user *model.SystemUser, err error) {
	params := map[string]interface{}{
		"name":              name,
		"username":          username,
		"auto_generate_key": true,
		"protocol":          model.ProtocolSSH,
	}

	resp, err := s.authClient.Post("/api/v1/assets/system-users/", params, &user)
	if err != nil {
		logger.Errorf("failed to create asset system user, err:%v, resp: %v\n", err, resp)
		return nil, err
	}

	logger.Debugf("create asset system user success, user:%v, resp:%v\n", user, resp)
	return user, nil
}

func (s *JMService) GetAssetSystemUserByName(name string) (user *model.SystemUser, err error) {
	params := map[string]string{
		"name": name,
	}
	users := make([]model.SystemUser, 0)
	resp, err := s.authClient.Get("/api/v1/assets/system-users/", &users, params)
	if err != nil {
		logger.Errorf("failed to list asset system user by name, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	if len(users) == 0 {
		logger.Errorf("asset system user can not found by name, err:%v, resp:%v\n", err, resp)
		return user, ErrAssetSystemUserNotFound
	} else {
		user = &users[0]
	}

	logger.Debugf("get asset system user by name success, user:%v, resp:%v\n", user, resp)
	return user, nil
}

func (s *JMService) GetAssetSystemUser(id string) (user *model.SystemUser, err error) {
	resp, err := s.authClient.Get(fmt.Sprintf("/api/v1/assets/system-users/%s/", id), user)
	if err != nil {
		logger.Errorf("get asset system user failed, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	logger.Debugf("get asset system user success, user:%v, resp:%v\n", user, resp)
	return user, nil
}

func (s *JMService) UpdateAssetSystemUser(id, name, username, privateKey, publicKey string) (user *model.SystemUser, err error) {
	params := map[string]string{
		"name":        name,
		"username":    username,
		"private_key": privateKey,
		"public_key":  publicKey,
		"protocol":    model.ProtocolSSH,
	}
	resp, err := s.authClient.Put(fmt.Sprintf("/api/v1/assets/system-users/%s/", id), params, &user)
	if err != nil {
		logger.Errorf("faiiled to update asset system user, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	logger.Debugf("update asset system user success, user:%v, resp:%v\n", user, resp)
	return user, nil
}

func (s *JMService) DeleteAssetSystemUser(id string) (err error) {
	resp, err := s.authClient.Delete(fmt.Sprintf("/api/v1/assets/system-users/%s/", id), nil)
	if err != nil {
		logger.Errorf("failed to delete asset system user, err:%v, resp:%v\n", err, resp)
		return err
	}

	logger.Debugf("delete asset system user success, id:%v, resp:%v\n", id, resp)
	return nil
}

func (s *JMService) GetAssetSystemUserAuthInfo(id string) (info *model.SystemUserAuthInfo, err error) {
	resp, err := s.authClient.Get(fmt.Sprintf("/api/v1/assets/system-users/%s/auth-info/", id), info)
	if err != nil {
		logger.Errorf("failed to get asset system user auth info, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	logger.Debugf("get asset system user auth info success, info:%v, resp:%v\n", info, resp)
	return info, nil
}