package service

import (
	"fmt"
	"time"

	"github.com/jumpserver/koko/pkg/jms-sdk-go/model"
)

type SystemUserAssetRelation struct {
	ID                string    `json:"id"`
	Asset             string    `json:"asset"`
	AssetDisplay      string    `json:"asset_display"`
	SystemUser        string    `json:"systemuser"`
	SystemUserDisplay string    `json:"systemuser_display"`
	Connectivity      string    `json:"connectivity"`
	DateVerified      time.Time `json:"date_verified"`
	OrgID             string    `json:"org_id"`
	OrgName           string    `json:"org_name"`
}

func (s *JMService) CreateAssetSystemUserAssetRelation(asset, systemUser string) (relation model.SystemUserAssetRelation, err error) {
	params := map[string]string{
		"asset":      asset,
		"systemuser": systemUser,
	}

	_, err = s.authClient.Post("/api/v1/assets/system-users-assets-relations/", params, &relation)
	return
}

func (s *JMService) DeleteAssetSystemUserAssetRelation(id string) (err error) {
	_, err = s.authClient.Delete(fmt.Sprintf("/api/v1/assets/system-users-assets-relations/%s/", id), nil)
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
