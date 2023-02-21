package goshopify

import (
	"fmt"
	"time"
)

const inventoryLevelsBasePath = "inventory_levels"

type InventoryLevelService interface {
	Adjust(param InventoryLevelAdjustParam) (*InventoryLevel, error)
}

type InventoryLevelServiceOp struct {
	client *Client
}

type InventoryLevel struct {
	InventoryItemID   int64      `json:"inventory_item_id,omitempty"`
	LocationID        int64      `json:"location_id,omitempty"`
	Available         int64      `json:"available,omitempty"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
	AdminGraphqlAPIID string     `json:"admin_graphql_api_id,omitempty"`
}

type InventoryLevelResource struct {
	InventoryLevel *InventoryLevel `json:"inventory_level"`
}

type InventoryLevelAdjustParam struct {
	AvailableAdjustment int64 `json:"available_adjustment,omitempty"`
	InventoryItemID     int64 `json:"inventory_item_id,omitempty"`
	LocationID          int64 `json:"location_id,omitempty"`
}

func (s *InventoryLevelServiceOp) Adjust(param InventoryLevelAdjustParam) (*InventoryLevel, error) {
	path := fmt.Sprintf("%s/adjust.json", inventoryLevelsBasePath)
	resource := new(InventoryLevelResource)
	err := s.client.Post(path, param, resource)
	return resource.InventoryLevel, err
}
