package service

import (
	"fmt"

	"github.com/jumpserver/koko/pkg/jms-sdk-go/model"
)

func (s *JMService) CheckUserCookie(cookies map[string]string) (user *model.User, err error) {
	client := s.authClient.Clone()
	for k, v := range cookies {
		client.SetCookie(k, v)
	}
	_, err = client.Get(UserProfileURL, &user)
	return
}

func (s *JMService) CreateUser(name, username, email string, systemRoles []string) (user model.User, err error) {
	params := map[string]interface{}{
		"name":         name,
		"username":     username,
		"email":        email,
		"system_roles": systemRoles,
	}

	_, err = s.authClient.Post("/api/v1/users/users/", params, &user)
	return
}

func (s *JMService) GetUser(id string) (user model.User, err error) {
	_, err = s.authClient.Get(fmt.Sprintf("/api/v1/users/users/%s/", id), &user)
	return
}

func (s *JMService) UpdateUser(id, name, email string, systemRoles []string) (user model.User, err error) {
	params := map[string]interface{}{
		"name":         name,
		"email":        email,
		"system_roles": systemRoles,
	}
	_, err = s.authClient.Put(fmt.Sprintf("/api/v1/users/users/%s/", id), params, &user)
	return
}

func (s *JMService) DeleteUser(id string) (err error) {
	_, err = s.authClient.Delete(fmt.Sprintf("/api/v1/users/users/%s/", id), nil)
	return
}

func (s *JMService) UpdateUserPassword(id, password string) (err error) {
	params := map[string]interface{}{
		"password": password,
	}
	_, err = s.authClient.Patch(fmt.Sprintf("/api/v1/users/users/%s/", id), params, nil)
	return
}

func (s *JMService) UpdateUserPublicKey(id, publicKey string) (err error) {
	params := map[string]interface{}{
		"public_key": publicKey,
	}
	_, err = s.authClient.Put(fmt.Sprintf("/api/v1/users/users/%s/", id), params, nil)
	return
}
