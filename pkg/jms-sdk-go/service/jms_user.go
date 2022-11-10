package service

import (
	"errors"
	"fmt"

	"github.com/jumpserver/koko/pkg/jms-sdk-go/model"
	"github.com/jumpserver/koko/pkg/logger"
)

var ErrUserNotFound = errors.New("user not found")

func (s *JMService) CheckUserCookie(cookies map[string]string) (user *model.User, err error) {
	client := s.authClient.Clone()
	for k, v := range cookies {
		client.SetCookie(k, v)
	}
	_, err = client.Get(UserProfileURL, &user)
	return
}

func (s *JMService) CreateUser(m map[string]interface{}) (user *model.User, err error) {
	resp, err := s.authClient.Post("/api/v1/users/users/", m, &user)
	if err != nil {
		logger.Errorf("failed to create user, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	logger.Debugf("create user success, user:%v, resp:%v\n", user, resp)
	return
}

func (s *JMService) GetUserByUserName(username string) (user *model.User, err error) {
	params := map[string]string{
		"username": username,
	}
	var users []model.User
	resp, err := s.authClient.Get("/api/v1/users/users/", &users, params)
	if err != nil {
		logger.Errorf("failed to get user by name, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	if len(users) < 1 {
		return user, ErrUserNotFound
	} else {
		user = &users[0]
	}

	logger.Debugf("get user by name success, user:%v, resp:%v\n", user, resp)
	return
}

func (s *JMService) GetUser(id string) (user model.User, err error) {
	_, err = s.authClient.Get(fmt.Sprintf("/api/v1/users/users/%s/", id), &user)
	return
}

func (s *JMService) UpdateUser(id string, m map[string]interface{}) (user *model.User, err error) {
	resp, err := s.authClient.Patch(fmt.Sprintf("/api/v1/users/users/%s/", id), m, &user)
	if err != nil {
		logger.Errorf("failed to update user, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	logger.Debugf("update user success, user:%v, resp:%v\n", user, resp)
	return
}

func (s *JMService) DeleteUser(id string) (err error) {
	resp, err := s.authClient.Delete(fmt.Sprintf("/api/v1/users/users/%s/", id), nil)
	if err != nil {
		logger.Errorf("failed to delete user, err:%v, resp:%v\n", err, resp)
		return err
	}

	logger.Debugf("delete user success, id:%s, resp:%v\n", id, resp)
	return
}

func (s *JMService) UpdateUserPassword(id, password string) (err error) {
	params := map[string]interface{}{
		"set_password": true,
		"password":     password,
	}
	resp, err := s.authClient.Patch(fmt.Sprintf("/api/v1/users/users/%s/", id), params, nil)
	if err != nil {
		logger.Errorf("failed to update user password, err:%v, resp:%v\n", err, resp)
		return err
	}

	logger.Debugf("update user password success, id:%s, resp:%v\n", id, resp)
	return
}

func (s *JMService) UpdateUserPublicKey(id, publicKey string) (err error) {
	params := map[string]interface{}{
		"set_public_key": true,
		"public_key":     publicKey,
	}
	resp, err := s.authClient.Patch(fmt.Sprintf("/api/v1/users/users/%s/", id), params, nil)
	if err != nil {
		logger.Errorf("failed to update user public key, err:%v, resp:%v\n", err, resp)
		return err
	}

	logger.Debugf("update user public_key success, id:%s, resp:%v\n", id, resp)
	return
}
