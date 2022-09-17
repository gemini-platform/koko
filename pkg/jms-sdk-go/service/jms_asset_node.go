package service

import (
	"fmt"

	"github.com/jumpserver/koko/pkg/jms-sdk-go/model"
)

type AssetNodeInfo struct {
	Count    int          `json:"count"`
	Next     string       `json:"next"`
	Previous string       `json:"previous"`
	Results  []model.Node `json:"results"`
}

func (s *JMService) ListAssetNode() (nodes []model.Node, err error) {
	_, err = s.authClient.Get("/api/v1/assets/nodes/", &nodes)
	return
}

func (s *JMService) CreateAssetNode(name string) (node model.Node, err error) {
	params := map[string]string{
		"value": name,
	}
	_, err = s.authClient.Post("/api/v1/assets/nodes/", params, &node)
	return
}

func (s *JMService) GetAssetNode(id string) (node model.Node, err error) {
	_, err = s.authClient.Get(fmt.Sprintf("/api/v1/assets/nodes/%s/", id), &node)
	return
}

func (s *JMService) UpdateAssetNode(id, name string) (node model.Node, err error) {
	params := map[string]string{
		"value": name,
	}
	_, err = s.authClient.Put(fmt.Sprintf("/api/v1/assets/nodes/%s/", id), params, &node)
	return
}

func (s *JMService) DeleteAssetNode(id string) (err error) {
	_, err = s.authClient.Delete(fmt.Sprintf("/api/v1/assets/nodes/%s/", id), nil)
	return
}

func (s *JMService) AssetNodeAddAsset(id string, assetIDs []string) (err error) {
	params := map[string][]string{
		"assets": assetIDs,
	}

	_, err = s.authClient.Put(fmt.Sprintf("/api/v1/assets/nodes/%s/assets/add/", id), params, nil)
	return
}

func (s *JMService) AssetNodeRemoveAsset(id string, assetIDs []string) (err error) {
	params := map[string][]string{
		"assets": assetIDs,
	}

	_, err = s.authClient.Put(fmt.Sprintf("/api/v1/assets/nodes/%s/assets/remove/", id), params, nil)
	return
}
