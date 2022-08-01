package service

import (
	"fmt"

	"github.com/jumpserver/koko/pkg/jms-sdk-go/model"
)

func (s *JMService) CreateAssetSystemUser(name, username, privateKey, protocol string) (user model.SystemUser, err error) {
	params := map[string]string{
		"name":        name,
		"username":    username,
		"private_key": privateKey,
		"protocol":    protocol,
	}
	_, err = s.authClient.Post("/api/v1/assets/system-users/", params, &user)
	return
}

func (s *JMService) GetAssetSystemUser(id string) (user model.SystemUser, err error) {
	_, err = s.authClient.Get(fmt.Sprintf("/api/v1/assets/system-users/%s/", id), &user)
	return
}

func (s *JMService) UpdateAssetSystemUser(id, name, username, privateKey, protocol string) (user model.SystemUser, err error) {
	params := map[string]string{
		"name":        name,
		"username":    username,
		"private_key": privateKey,
		"protocol":    protocol,
	}
	_, err = s.authClient.Put(fmt.Sprintf("/api/v1/assets/system-users/%s/", id), params, &user)
	return
}

func (s *JMService) DeleteAssetSystemUser(id string) (err error) {
	_, err = s.authClient.Delete(fmt.Sprintf("/api/v1/assets/system-users/%s/", id), nil)
	return
}

func (s *JMService) CreateAssetSystemUserNodeRelation(systemUserID, nodeID string) (relation model.SystemUserNodeRelation, err error) {
	params := map[string]string{
		"systemuser_id": systemUserID,
		"node_id":       nodeID,
	}
	_, err = s.authClient.Post("/api/v1/assets/system-users-node-relations/", params, &relation)
	return
}

func (s *JMService) GetAssetSystemUserNodeRelation(id int) (relation model.SystemUserNodeRelation, err error) {
	_, err = s.authClient.Get(fmt.Sprintf("/api/v1/assets/system-users-node-relations/%d/", id), &relation)
	return
}

func (s *JMService) UpdateAssetSystemUserNodeRelation(id int, systemUserID, nodeID string) (relation model.SystemUserNodeRelation, err error) {
	params := map[string]string{
		"systemuser_id": systemUserID,
		"node_id":       nodeID,
	}
	_, err = s.authClient.Put(fmt.Sprintf("/api/v1/assets/system-users-node-relations/%d/", id), params, &relation)
	return
}

func (s *JMService) DeleteAssetSystemUserNodeRelation(id int) (err error) {
	_, err = s.authClient.Delete(fmt.Sprintf("/api/v1/assets/system-users-node-relations/%d/", id), nil)
	return
}
