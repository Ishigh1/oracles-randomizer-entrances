package rom

func newAgesRomBanks() *romBanks {
	r := romBanks{
		endOfBank: make([]uint16, 0x40),
	}

	r.endOfBank[0x00] = 0x3ef8
	r.endOfBank[0x02] = 0x7e93
	r.endOfBank[0x03] = 0x7ebd
	r.endOfBank[0x04] = 0x7edb
	r.endOfBank[0x05] = 0x7d9d
	r.endOfBank[0x06] = 0x7a31
	r.endOfBank[0x09] = 0x7dee
	r.endOfBank[0x0a] = 0x7e09
	r.endOfBank[0x0b] = 0x7fa8
	r.endOfBank[0x0c] = 0x7f94
	r.endOfBank[0x10] = 0x7ef4
	r.endOfBank[0x11] = 0x7f73
	r.endOfBank[0x12] = 0x7e8f
	r.endOfBank[0x15] = 0x7bfb
	r.endOfBank[0x16] = 0x7e03
	r.endOfBank[0x3f] = 0x7d0a

	return &r
}

func initAgesEOB() {
	r := newAgesRomBanks()

	// bank 00

	// don't play any music if the -nomusic flag is given. because of this,
	// this *must* be the first function at the end of bank zero (for now).
	r.appendToBank(0x00, "no music func",
		"\x67\xfe\x47\x30\x03\x3e\x08\xc9\xf0\xb7\xc9")
	r.replace(0x00, 0x0c9a, "no music call",
		"\x67\xf0\xb7", "\x67\xf0\xb7") // modified only by SetNoMusic()

	// only increment the maku tree's state if on the maku tree screen, or if
	// all essences are obtained, set it to the value it would normally have at
	// that point in the game. this allows getting the maku tree's item as long
	// as you haven't collected all essences.
	makuStateCheck := r.appendToBank(0x00, "maku state check",
		"\xfa\x2d\xcc\xfe\x02\x30\x0e\xfa\x30\xcc\xfe\x38\x20\x07"+
			"\xfa\xe8\xc6\x3c\xfe\x11\xc9\xfa\xbf\xc6\x3c\x37\x20\x03"+
			"\x3e\x0e\xc9\xfa\xe8\xc6\xc9")
	r.replace(0x00, 0x3e56, "call maku state check",
		"\x3c\xfe\x11", "\xcd"+makuStateCheck)

	// return z iff the current group and room match c and b.
	compareRoom := r.appendToBank(0x00, "compare room",
		"\xfa\x2d\xcc\xb9\xc0\xfa\x30\xcc\xb8\xc9")

	// read 2 bytes from bank e at hl into bc.
	readWord := r.appendToBank(0x00, "read word",
		"\xfa\x97\xff\xf5\x7b\xea\x97\xff\xea\x22\x22"+ // switch bank
			"\x2a\x47\x7e\x4f"+ // read
			"\xf1\xea\x97\xff\xea\x22\x22\xc9") // switch back

	// bank 02

	// warp to ember tree if holding start when closing the map screen.
	treeWarp := r.appendToBank(0x02, "tree warp",
		"\xfa\x81\xc4\xe6\x08\x28\x1b"+ // close as normal if start not held
			"\xfa\x2d\xcc\xfe\x02\x38\x06"+ // check if indoors
			"\x3e\x5a\xcd\x98\x0c\xc9"+ // play error sound and ret
			"\x21\x47\xcc\x36\x80\x23\x36\x78\x2e\x4a\x36\x55"+ // set dest
			"\xcd\xbf\x5f\xc3\xba\x4f") // close + warp
	r.replaceMultiple([]Addr{{0x02, 0x6133}, {0x02, 0x618b}}, "tree warp jump",
		"\xc2\xba\x4f", "\xc4"+treeWarp)

	// warp to room under cursor if wearing developer ring.
	devWarp := r.appendToBank(0x02, "dev ring warp func",
		"\xfa\xcb\xc6\xfe\x40\x20\x12\xfa\x2d\xcc\xfe\x02\x30\x0b\xf6\x80"+
			"\xea\x47\xcc\xfa\xb6\xcb\xea\x48\xcc\x3e\x03\xcd\xad\x0c\xc9")
	r.replace(0x02, 0x5fcc, "dev ring warp call", "\xad\x0c", devWarp)

	// bank 03

	// allow skipping the capcom screen after one second by pressing start
	skipCapcom := r.appendToBank(0x03, "skip capcom func",
		"\xe5\xfa\xb3\xcb\xfe\x94\x30\x03\xcd\x86\x08\xe1\xcd\x37\x02\xc9")
	r.replace(0x03, 0x4d6c, "skip capcom call", "\x37\x02", skipCapcom)

	// set intro flags to skip opening
	skipOpening := r.appendToBank(0x03, "skip opening",
		"\xcd\xf9\x31\x3e\x0a\xcd\xf9\x31"+ // set global flags
			"\xe5\x21\x7a\xc7\xcb\xf6\x2e\x6a\xcb\xf6"+ // set room flags
			"\x2e\x59\xcb\xf6\x2e\x39\x36\xc8\xe1\xc9") // set more room flags
	r.replace(0x03, 0x6e97, "call skip opening",
		"\xc3\xf9\x31", "\xc3"+skipOpening)

	// bank 04

	// look up tiles in custom replacement table after loading a room. the
	// format is (group, room, bitmask, YX, tile ID), with ff ending the table.
	// if the bitmask AND the current room flags is nonzero, the replacement is
	// not made.
	tileReplaceTable := r.appendToBank(0x04, "tile replace table",
		"\x01\x48\x00\x45\xd7"+ // portal south of past maku tree
			"\x00\x39\x00\x63\xf0"+ // open chest on intro screen
			"\x00\x39\x20\x63\xf1"+ // closed chest on intro screen
			"\x00\x6b\x00\x42\x3a"+ // removed tree in yoll graveyard
			"\x00\x6b\x02\x42\xce"+ // not removed tree in yoll graveyard
			"\x03\x0f\x00\x66\xf9"+ // water in d6 past entrance
			"\x04\x1b\x01\x03\x78"+ // key door in D1
			"\xff")
	tileReplaceFunc := r.appendToBank(0x04, "tile replace body",
		"\xc5\xd5\xcd\x7d\x19\x5f\x21"+tileReplaceTable+"\xfa\x2d\xcc\x47"+
			"\xfa\x30\xcc\x4f"+ // load room flags, table addr, group, room
			"\x2a\xfe\xff\x28\x1b\xb8\x20\x12\x2a\xb9\x20\x0f"+
			"\x2a\xa3\x20\x0c"+ // compare group, room, flags
			"\xd5\x16\xcf\x2a\x5f\x2a\x12\xd1\x18\xe6"+ // replace
			"\x23\x23\x23\x23\x18\xe0\xd1\xc1\xcd\xef\x5f\xc9")
	r.replace(0x00, 0x38c0, "tile replace call",
		"\xcd\xef\x5f", "\xcd"+tileReplaceFunc)

	// bank 05

	// if wearing dev ring, jump over any tile like a ledge by pressing B with
	// no B item equipped.
	cliffLookup := r.appendToBank(0x05, "cliff lookup",
		"\xf5\xfa\xcb\xc6\xfe\x40\x20\x13"+ // check ring
			"\xfa\x88\xc6\xb7\x20\x0d"+ // check B item
			"\xfa\x81\xc4\xe6\x02\x28\x06"+ // check input
			"\xf1\xfa\x09\xd0\x37\xc9"+ // jump over ledge
			"\xf1\xc3\x1f\x1e") // jp to normal lookup
	r.replace(0x05, 0x6083, "call cliff lookup",
		"\xcd\x1f\x1e", "\xcd"+cliffLookup)

	// bank 06

	// burning the first tree in yoll graveyard should set room flag 1 so that
	// it can be gone for good.
	removeYollTree := r.appendToBank(0x06, "remove yoll tree",
		"\xf5\x01\x00\x6b\xcd"+compareRoom+"\x20\x05\x21\x6b\xc7\xcb\xce\xf1"+
			"\xc3\xeb\x16")
	r.replace(0x06, 0x483e, "call remove yoll tree",
		"\xcd\xeb\x16", "\xcd"+removeYollTree)

	// bank 09

	// set treasure ID 07 (rod of seasons) when buying the 150 rupee shop item,
	// so that the shop can check this specific ID.
	shopSetFakeID := r.appendToBank(0x09, "shop set fake ID",
		"\xfe\x0d\x20\x05\x21\x9a\xc6\xcb\xfe\x21\xf7\x44\xc9")
	r.replace(0x09, 0x4418, "call shop set fake ID",
		"\x21\xf7\x44", "\xcd"+shopSetFakeID)

	// set treasure ID 08 (magnet gloves) when getting item from south shore
	// dirt pile.
	digSetFakeID := r.appendToBank(0x09, "dirt set fake ID",
		"\xc5\x01\x00\x98\xcd"+compareRoom+"\xc1\xc0\xe5\x21\x9b\xc6\xcb\xc6"+
			"\xe1\xc9")
	// set flag for d6 past boss key whether you get it in past or present.
	setD6BossKey := r.appendToBank(0x09, "set d6 boss key",
		"\x7b\xfe\x31\xc0\xfa\x39\xcc\xfe\x06\xc0\xe5\x21\x83\xc6\xcb\xe6\xe1"+
			"\xc9")
	// refill all seeds when picking up a seed satchel.
	refillSeedSatchel := r.appendToBank(0x09, "refill seed satchel",
		"\x7b\xfe\x19\xc0\xd5\xe5\x21\xb4\xc6\x34\xcd\x0c\x18\x35\xe1\xd1\xc9")
	// give 20 seeds when picking up the seed shooter.
	fillSeedShooter := r.appendToBank(0x09, "fill seed shooter",
		"\x7b\xfe\x0f\xc0\xc5\x3e\x20\x0e\x20\xcd\x1c\x17\xc1\xc9")
	// this function checks all the above conditions when collecting an item.
	handleGetItem := r.appendToBank(0x09, "handle get item",
		"\x5f\xcd"+digSetFakeID+"\xcd"+setD6BossKey+"\xcd"+refillSeedSatchel+
			"\xcd"+fillSeedShooter+"\x7b\xc3\x1c\x17")
	r.replace(0x09, 0x4c4e, "call handle get item",
		"\xcd\x1c\x17", "\xcd"+handleGetItem)

	// bank 0a

	// set sub ID for south shore dig item.
	dirtSpawnItem := r.appendToBank(0x0a, "dirt spawn item",
		"\xcd\xd4\x27\xc0\xcd\x42\x22\xaf\xc9")
	r.replace(0x0a, 0x5e3e, "call dirt spawn item",
		"\xcd\xc5\x24", "\xcd"+dirtSpawnItem)

	// bank 0c

	// use custom script for soldier in deku forest with sub ID 0; they should
	// give an item in exchange for mystery seeds.
	soldierScriptAfter := r.appendToBank(0x0c, "soldier script after item",
		"\x97\x59\x08\x00")
	soldierScriptGive := r.appendToBank(0x0c, "soldier script give item",
		"\xeb\x9e\x98\x59\x0b\xb4\xbd\x00\x92\xe9\xcb\x02\xde\x00\x00\xb1\x20"+
			"\xc4"+soldierScriptAfter)
	soldierScriptCheck := r.appendToBank(0x0c, "soldier script check count",
		"\xb3\xbd\xff"+soldierScriptGive+"\x5d\xee")
	soldierScript := r.appendToBank(0x0c, "soldier script",
		"\xb0\x20"+soldierScriptAfter+"\xdf\x24"+soldierScriptCheck+"\x5d\xee")
	r.replace(0x09, 0x5207, "soldier script pointer", "\xee\x5d", soldierScript)

	// bank 11

	// allow collection of seeds with only shooter and no satchel
	checkSeedHarvest := r.appendToBank(0x11, "check seed harvest",
		"\xcd\x48\x17\xd8\x3e\x0f\xc3\x48\x17")
	r.replace(0x11, 0x4aba, "call check seed harvest",
		"\xcd\x48\x17", "\xcd"+checkSeedHarvest)

	// bank 15

	// don't equip sword for shooting galleries if player don't have it
	// (doesn't work anyway).
	shootingGalleryEquip := r.appendToBank(0x15, "shooting gallery equip",
		"\x3e\x05\xcd\x48\x17\x3e\x00\x22\xd0\x2b\x3e\x05\x22\xc9")
	r.replace(0x15, 0x50ae, "call shooting gallery equip",
		"\x3e\x05\x22", "\xcd"+shootingGalleryEquip)

	// always make "boomerang" second prize for target carts, checking room
	// flag 6 to track it.
	targetCartsItem := r.appendToBank(0x15, "target carts item",
		"\xcd\x7d\x19\xcb\x77\x3e\x04\xca\xbb\x66\xcd\x3e\x04\xc3\xa5\x66")
	r.replace(0x15, 0x66a2, "call target carts item",
		"\xcd\x3e\x04", "\xc3"+targetCartsItem)
	// set room flag 6 when "boomerang" is given in script.
	targetCartsFlag := r.appendToBank(0x0c, "target carts flag",
		"\xde\x06\x02\xb1\x40\xc1")
	r.replace(0x0c, 0x6e6e, "jump target carts flag",
		"\x88\x6e", targetCartsFlag)

	// bank 16

	// upgraded item data (one byte for old ID, one for new ID two for address):
	progItemAddrs := r.appendToBank(0x16, "progressive item addrs",
		"\x05\x05\xea\x54"+ // noble sword
			"\x0a\x0a\x12\x55"+ // long switch
			"\x16\x16\x52\x55"+ // power glove
			"\x19\x19\x76\x55"+ // satchel upgrade
			"\x25\x26\xca\x53"+ // tune of currents
			"\x26\x27\xce\x53"+ // tune of ages
			"\x2e\x4a\x5a\x54"+ // mermaid suit
			"\xff")
	// given a treasure ID in b, make hl = the start of the upgraded treasure
	// data + 1, if the treasure needs to be upgraded, and returns the new
	// treasure ID in b.
	getUpgradedTreasure := r.appendToBank(0x16, "get upgraded treasure",
		"\x78\xcd\x48\x17\x78\xd0"+ // check obtained
			"\xfe\x25\x20\x09\x3e\x26\x5f\xcd\x48\x17\x30\x01\x43"+ // harp
			"\xe5\x21"+progItemAddrs+"\x2a\xfe\xff\x28\x13"+ // search
			"\xb8\x20\x06\x2a\x47\x2a\x5e\x18\x05\x23\x23\x23\x18\xed"+
			"\xe1\x63\x6f\x23\xc9\xe1\xc9") // done
	// given a treasure ID in dx42, return hl = the start of the treasure data
	// + 1, accounting for progressive upgrades. also writes the new treasure
	// ID to d070, which is used to set the treasure obtained flag.
	upgradeTreasure := r.appendToBank(0x16, "upgrade treasure",
		"\x1e\x42\x1a\x47\xcd"+getUpgradedTreasure+"\x1e\x70\x78\x12\xc9")

	// load the address of a treasure's 4-byte data entry + 1 into hl, using b
	// as the ID and c as sub ID, accounting for progressive upgrades.
	getTreasureDataBody := r.appendToBank(0x16, "get treasure data body",
		"\x21\x32\x53\x78\x87\x87\xd7\xcb\x7e\x28\x04\x23\x2a\x66\x6f"+
			"\x79\x87\x87\xd7\x23\xc3"+getUpgradedTreasure)
	// just get item bc's sprite index in e.
	getItemSpriteIndexBody := r.appendToBank(0x16, "get item sprite index body",
		"\xcd"+getTreasureDataBody+"\x23\x23\x7e\x5f\xc9")
	getItemSpriteIndex := r.appendToBank(0x00, "get item sprite index",
		"\x1e\x16\x21"+getItemSpriteIndexBody+"\xc3\x8a\x00")

	// return collection mode in a and e, based on current room. call is in
	// bank 16, func is in bank 00, body is in bank 06.
	collectModeTable := r.appendToBank(0x06, "collect mode table",
		makeCollectModeTable())
	// make tree item falls or exists on floor depending on script position.
	collectMakuTreeFunc := r.appendToBank(0x06, "collect maku tree",
		"\xfa\x58\xd2\xfe\x84\x1e\x29\xc8\x1e\x0a\xc9")
	// target carts items appear with a poof if they're in the enclosure.
	collectTargetCartsFunc := r.appendToBank(0x06, "collect target carts",
		"\x1e\x4d\x1a\xfe\x78\x1e\x19\xc8\x1e\x0a\xc9")
	// big bang game items appear with a poof if they're above the goron.
	collectBigBangFunc := r.appendToBank(0x06, "collect big bang game",
		"\x1e\x4b\x1a\xfe\x38\x1e\x19\xc8\x1e\x0a\xc9")
	collectModeJumpTable := r.appendToBank(0x06, "collect mode jump table",
		collectMakuTreeFunc+collectTargetCartsFunc+collectBigBangFunc)
	collectModeLookupBody := r.appendToBank(0x06, "collect mode lookup body",
		"\xfa\x2d\xcc\x47\xfa\x30\xcc\x4f\x21"+collectModeTable+
			"\x2a\xfe\xff\xc8\xb8\x20\x14\x2a\xb9\x20\x11\x7e\x5f"+
			"\xfe\x80\xd8\x21"+collectModeJumpTable+"\xe6\x7f"+
			"\x87\xd7\x2a\x66\x6f\xe9\x23\x23\x18\xe1")
	collectModeLookup := r.appendToBank(0x00, "collect mode lookup",
		"\xc5\xd5\xe5\x1e\x06\x21"+collectModeLookupBody+"\xcd\x8a\x00\x7b"+
			"\xe1\xd1\xc1\xc9")
	// return treasure data address and collect mode modified as necessary,
	// given a treasure ID in dx42.
	modifyTreasure := r.appendToBank(0x16, "modify treasure",
		"\xcd"+upgradeTreasure+"\xcd"+collectModeLookup+"\x47\xcb\x37\xc9")
	r.replace(0x16, 0x4539, "call modify treasure",
		"\x47\xcb\x37", "\xcd"+modifyTreasure)

	// bank 3f

	// set hl to the address of the item sprite with ID a.
	getItemSpriteAddr := r.appendToBank(0x3f, "get item sprite addr",
		"\x21\xdb\x66\x5f\x87\xd7\x7b\xd7\xc9")

	// this should be jumped to at the end of every custom sprite function.
	endCustomSprite := r.appendToBank(0x3f, "end custom sprite",
		"\xf1\xc9")
	// make the deku forest soldier that gives the item red instead of blue.
	soldierSprite := r.appendToBank(0x3f, "soldier sprite", "\x4d\x00\x22")
	setSoldierSprite := r.appendToBank(0x3f, "set soldier sprite",
		"\x21"+soldierSprite+"\xc3"+endCustomSprite)
	setShopItemSprite := r.appendToBank(0x3f, "set shop item sprite",
		"\x1e\x09\x21\x11\x45\xc5\xcd"+readWord+"\xcd"+getItemSpriteIndex+
			"\x7b\xcd"+getItemSpriteAddr+"\xc1\xc3"+endCustomSprite)
	// table of ID, sub ID, jump address
	customSpriteTable := r.appendToBank(0x3f, "custom sprite table",
		"\x40\x00"+setSoldierSprite+
			"\x47\x0d"+setShopItemSprite+
			"\xff")
	// override the sprites loaded for certain ID / sub ID pairs.
	loadCustomSprite := r.appendToBank(0x3f, "load custom sprite",
		"\xcd\x37\x44"+
			"\xf5\xc5\xe5\x1e\x41\x1a\x47\x1c\x1a\x4f\x21"+customSpriteTable+
			"\x2a\xfe\xff\x28\x14\xb8\x20\x0c\x2a\xb9\x20\x09"+
			"\x2a\x47\x7e\xe1\x67\x68\xc1\xe9"+
			"\x23\x23\x23\x18\xe7\xe1\xc1\xc3"+endCustomSprite)
	r.replace(0x3f, 0x4356, "call load custom sprite",
		"\xcd\x37\x44", "\xcd"+loadCustomSprite)
}