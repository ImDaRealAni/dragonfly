package item

// Coal is an item used as fuel & crafting torches.
type Coal struct{} //TODO: Fuel

// EncodeItem ...
func (Coal) EncodeItem() (id int32, name string, meta int16) {
	return 263, "minecraft:coal", 0
}
