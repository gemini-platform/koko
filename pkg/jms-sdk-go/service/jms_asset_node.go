package service

import (
	"errors"
	"fmt"

	"github.com/jumpserver/koko/pkg/logger"
	"github.com/jumpserver/koko/pkg/jms-sdk-go/model"
)

var ErrAssetNodeNotFound = errors.New("asset node found")

type AssetNodeInfo struct {
	Count    int          `json:"count"`
	Next     string       `json:"next"`
	Previous string       `json:"previous"`
	Results  []model.Node `json:"results"`
}

func (s *JMService) ListAssetNode() (nodes []model.Node, err error) {
	resp, err := s.authClient.Get("/api/v1/assets/nodes/", &nodes)
	if err != nil {
		logger.Errorf("failed to list asset node, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	logger.Debugf("list asset node success, nodes:%v, resp:%v\n", nodes, resp)
	return nodes, nil
}

func (s *JMService) CreateAssetNode(name string) (node *model.Node, err error) {
	params := map[string]string{
		"value": name,
	}

	resp, err := s.authClient.Post("/api/v1/assets/nodes/", params, &node)
	if err != nil {
		logger.Errorf("failed to create asset node, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	logger.Debugf("create asset node success, node:%v, resp:%v\n", node, resp)
	return node, nil
}

func (s *JMService) GetAssetNodeByName(name string) (node *model.Node, err error) {
	params := map[string]string{
		"value": name,
	}

	var nodes []model.Node
	resp, err := s.authClient.Get("/api/v1/assets/nodes/", &nodes, params)
	if err != nil {
		logger.Errorf("failed to list asset node by name, err:%v, resp:%v\n", err, resp)
		return node, err
	}

	if len(nodes) == 0 {
		logger.Errorf("asset node can not fond by name, err:%v, resp:%v\n", err, resp)
		return node, ErrAssetNodeNotFound
	} else {
		node = &nodes[0]
	}

	logger.Debugf("get asset node by name success, node:%v, resp:%v\n", node, resp)
	return
}

func (s *JMService) GetAssetNode(id string) (node *model.Node, err error) {
	resp, err := s.authClient.Get(fmt.Sprintf("/api/v1/assets/nodes/%s/", id), node)
	if err != nil {
		logger.Errorf("failed to get asset node, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	logger.Debugf("get asset node success, node:%v, resp:%v\n", node, resp)
	return node, nil
}

func (s *JMService) UpdateAssetNode(id, name string) (node *model.Node, err error) {
	params := map[string]string{
		"value": name,
	}
	resp, err := s.authClient.Put(fmt.Sprintf("/api/v1/assets/nodes/%s/", id), params, &node)
	if err != nil {
		logger.Errorf("failed to update asset node, err:%v, resp:%v\n", err, resp)
		return nil, err
	}

	logger.Debugf("update asset node success, node:%v, resp:%v\n", node, resp)
	return node, nil
}

func (s *JMService) DeleteAssetNode(id string) (err error) {
	resp, err := s.authClient.Delete(fmt.Sprintf("/api/v1/assets/nodes/%s/", id), nil)
	if err != nil {
		logger.Errorf("failed to get asset node, err:%v, resp:%v\n", err, resp)
		return err
	}

	logger.Debugf("delete asset node success, id:%s, resp:%v\n", id, resp)
	return nil
}

func (s *JMService) AssetNodeAddAsset(id string, assetIDs []string) (err error) {
	params := map[string][]string{
		"assets": assetIDs,
	}

	resp, err := s.authClient.Put(fmt.Sprintf("/api/v1/assets/nodes/%s/assets/add/", id), params, nil)
	if err != nil {
		logger.Errorf("failed to add asset to asset node, err:%v, resp:%v\n", err, resp)
		return err
	}

	logger.Debugf("add asset to asset node success, id:%s, assets: %v resp:%v\n", id, assetIDs, resp)
	return nil
}

func (s *JMService) AssetNodeRemoveAsset(id string, assetIDs []string) (err error) {
	params := map[string][]string{
		"assets": assetIDs,
	}

	resp, err := s.authClient.Put(fmt.Sprintf("/api/v1/assets/nodes/%s/assets/remove/", id), params, nil)
	if err != nil {
		logger.Errorf("failed to remove asset from asset node, err:%v, resp:%v\n", err, resp)
		return err
	}

	logger.Debugf("remove asset to asset node success, id:%s, assets: %v resp:%v\n", id, assetIDs, resp)
	return nil
}
