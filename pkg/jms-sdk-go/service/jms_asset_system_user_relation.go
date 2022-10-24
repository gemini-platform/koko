package service

import (
	"fmt"
	"time"

	"github.com/jumpserver/koko/pkg/logger"
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

func (s *JMService) CreateAssetSystemUserAssetRelation(asset, systemUser string) (relation *model.SystemUserAssetRelation, err error) {
	params := map[string]string{
		"asset":      asset,
		"systemuser": systemUser,
	}

	resp, err := s.authClient.Post("/api/v1/assets/system-users-assets-relations/", params, relation)
	if err != nil {
		logger.Errorf("failed to get asset node, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	logger.Debugf("update asset system user and asset relation success, relation:%v, resp:%v\n", relation, resp)
	return relation, nil
}

func (s *JMService) DeleteAssetSystemUserAssetRelation(id string) (err error) {
	resp, err := s.authClient.Delete(fmt.Sprintf("/api/v1/assets/system-users-assets-relations/%s/", id), nil)
	if err != nil {
		logger.Errorf("failed to delete asset system user and node relation, err:%v, resp:%v\n", err, resp)
		return err
	}

	logger.Debugf("delete asset system user and asset relation success, id:%s, resp:%v\n", id, resp)
	return nil
}

func (s *JMService) CreateAssetSystemUserNodeRelation(systemUserID, nodeID string) (relation *model.SystemUserNodeRelation, err error) {
	params := map[string]string{
		"systemuser_id": systemUserID,
		"node_id":       nodeID,
	}
	resp, err := s.authClient.Post("/api/v1/assets/system-users-node-relations/", params, &relation)
	if err != nil {
		logger.Errorf("failed to create asset system user and node relation, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	logger.Debugf("create asset system user and node relation success, relation:%v, resp:%v\n", relation, resp)
	return relation, nil
}

func (s *JMService) GetAssetSystemUserNodeRelation(id int) (relation *model.SystemUserNodeRelation, err error) {
	resp, err := s.authClient.Get(fmt.Sprintf("/api/v1/assets/system-users-node-relations/%d/", id), relation)
	if err != nil {
		logger.Errorf("failed to get asset system user and node relation, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	logger.Debugf("get asset system user and node relation success, relation:%v, resp:%v\n", relation, resp)
	return relation, nil
}

func (s *JMService) UpdateAssetSystemUserNodeRelation(id int, systemUserID, nodeID string) (relation *model.SystemUserNodeRelation, err error) {
	params := map[string]string{
		"systemuser_id": systemUserID,
		"node_id":       nodeID,
	}
	resp, err := s.authClient.Put(fmt.Sprintf("/api/v1/assets/system-users-node-relations/%d/", id), params, &relation)
	if err != nil {
		logger.Errorf("failed to update asset system user and node relation, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	logger.Debugf("update asset system user and node relation success, relation:%v, resp:%v\n", relation, resp)
	return relation, nil
}

func (s *JMService) DeleteAssetSystemUserNodeRelation(id int) (err error) {
	resp, err := s.authClient.Delete(fmt.Sprintf("/api/v1/assets/system-users-node-relations/%d/", id), nil)
	if err != nil {
		logger.Errorf("failed to delete asset system user and node relation, err:%v, resp:%v\n", err, resp)
		return err
	}

	logger.Debugf("delete asset system user and node relation success, id:%d, resp:%v\n", id, resp)
	return nil
}
