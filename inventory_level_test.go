package goshopify

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func inventoryLevelTest(t *testing.T, level *InventoryLevel) {
	if level == nil {
		t.Errorf("InventoryLevel is nil")
		return
	}

	expectedInt := int64(808950810)
	if level.InventoryItemID != expectedInt {
		t.Errorf("InventoryLevel.InventoryItemID returned %+v, expected %+v", level.InventoryItemID, expectedInt)
	}

	expectedInt = 39072856
	if level.LocationID != expectedInt {
		t.Errorf("InventoryLevel.LocationID returned %+v, expected %+v", level.LocationID, expectedInt)
	}

	expectedInt = 6
	if level.Available != expectedInt {
		t.Errorf("InventoryLevel.Available returned %+v, expected %+v", level.Available, expectedInt)
	}
}

func TestInventoryLevelAdjust(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf(
		"https://fooshop.myshopify.com/%s/inventory_levels/adjust.json",
		client.pathPrefix),
		httpmock.NewBytesResponder(200, loadFixture("inventory_level.json")))

	param := InventoryLevelAdjustParam{
		AvailableAdjustment: -1,
		InventoryItemID:     808950810,
		LocationID:          39072856,
	}

	adjustedLevel, err := client.InventoryLevel.Adjust(param)
	if err != nil {
		t.Errorf("InventoryLevel.Adjust returned error: %v", err)
	}

	inventoryLevelTest(t, adjustedLevel)
}
