package block

import (
	"github.com/df-mc/dragonfly/dragonfly/item"
	"math/rand"
)

// NetherGoldOre is a variant of gold ore found exclusively in The Nether.
type NetherGoldOre struct {
	noNBT
	solid
}

// BreakInfo ...
func (n NetherGoldOre) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness:    3,
		Harvestable: pickaxeHarvestable,
		Effective:   pickaxeEffective,
		Drops:       simpleDrops(item.NewStack(item.GoldNugget{}, rand.Intn(4)+2)), //TODO: Silk Touch
		XPDrops:     XPDropRange{0, 1},
	}
}

// EncodeItem ...
func (n NetherGoldOre) EncodeItem() (id int32, meta int16) {
	return -288, 0
}

// EncodeBlock ...
func (n NetherGoldOre) EncodeBlock() (name string, properties map[string]interface{}) {
	return "minecraft:nether_gold_ore", nil
}

// Hash ...
func (n NetherGoldOre) Hash() uint64 {
	return hashNetherGoldOre
}
