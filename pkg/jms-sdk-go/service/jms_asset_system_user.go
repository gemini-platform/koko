package service

import (
	"fmt"

	"github.com/jumpserver/koko/pkg/jms-sdk-go/model"
)

func (s *JMService) CreateAssetSystemUser(name, username string) (user model.SystemUser, err error) {
	params := map[string]interface{}{
		"name":              name,
		"username":          username,
		"auto_generate_key": true,
		"protocol":          model.ProtocolSSH,
	}
	_, err = s.authClient.Post("/api/v1/assets/system-users/", params, &user)
	return
}

func (s *JMService) GetAssetSystemUser(id string) (user model.SystemUser, err error) {
	_, err = s.authClient.Get(fmt.Sprintf("/api/v1/assets/system-users/%s/", id), &user)
	return
}

func (s *JMService) UpdateAssetSystemUser(id, name, username, privateKey, publicKey string) (user model.SystemUser, err error) {
	params := map[string]string{
		"name":        name,
		"username":    username,
		"private_key": privateKey,
		"public_key":  publicKey,
		"protocol":    model.ProtocolSSH,
	}
	_, err = s.authClient.Put(fmt.Sprintf("/api/v1/assets/system-users/%s/", id), params, &user)
	return
}

func (s *JMService) DeleteAssetSystemUser(id string) (err error) {
	_, err = s.authClient.Delete(fmt.Sprintf("/api/v1/assets/system-users/%s/", id), nil)
	return
}
