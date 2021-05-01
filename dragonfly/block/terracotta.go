package block

import (
	"github.com/df-mc/dragonfly/dragonfly/item"
)

// Terracotta is a block formed from clay, with a hardness and blast resistance comparable to stone. For colouring it,
// take a look at StainedTerracotta.
type Terracotta struct {
	solid
	bassDrum
}

// BreakInfo ...
func (t Terracotta) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness:    1.25,
		Harvestable: pickaxeHarvestable,
		Effective:   pickaxeEffective,
		Drops:       simpleDrops(item.NewStack(t, 1)),
	}
}

// EncodeItem ...
func (Terracotta) EncodeItem() (id int32, name string, meta int16) {
	return 172, "minecraft:hardened_clay", meta
}

// EncodeBlock ...
func (Terracotta) EncodeBlock() (string, map[string]interface{}) {
	return "minecraft:hardened_clay", nil
}
