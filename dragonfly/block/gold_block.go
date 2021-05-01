package block

import (
	"github.com/df-mc/dragonfly/dragonfly/block/instrument"
	"github.com/df-mc/dragonfly/dragonfly/item"
	"github.com/df-mc/dragonfly/dragonfly/item/tool"
)

// GoldBlock is a precious metal block crafted from 9 gold ingots.
type GoldBlock struct {
	solid
}

// Instrument ...
func (g GoldBlock) Instrument() instrument.Instrument {
	return instrument.Bell()
}

// BreakInfo ...
func (g GoldBlock) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness: 5,
		Harvestable: func(t tool.Tool) bool {
			return t.ToolType() == tool.TypePickaxe && t.HarvestLevel() >= tool.TierIron.HarvestLevel
		},
		Effective: pickaxeEffective,
		Drops:     simpleDrops(item.NewStack(g, 1)),
	}
}

// PowersBeacon ...
func (GoldBlock) PowersBeacon() bool {
	return true
}

// EncodeItem ...
func (GoldBlock) EncodeItem() (id int32, name string, meta int16) {
	return 41, "minecraft:gold_block", 0
}

// EncodeBlock ...
func (GoldBlock) EncodeBlock() (string, map[string]interface{}) {
	return "minecraft:gold_block", nil
}
