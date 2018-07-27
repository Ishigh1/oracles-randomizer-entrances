package main

var generatedPoints = map[string]Point{
	"beach":                     Point{Parents: []string{"beach 1", "beach 2", "beach 3", "beach 4", "beach 5", "beach 6", "beach 7"}, Type: 2},
	"bridge":                    Point{Parents: []string{"bridge 1", "bridge 2", "bridge 3"}, Type: 2},
	"d2 arrow room":             Point{Parents: []string{"d2 arrow room 1", "d2 arrow room 2"}, Type: 2},
	"d2 blade key chest":        Point{Parents: []string{"d2 blade key chest 1", "d2 blade key chest 2"}, Type: 2},
	"d2 compass chest":          Point{Parents: []string{"d2 compass chest 1", "d2 compass chest 2"}, Type: 2},
	"d3 basement A in":          Point{Parents: []string{"d3 basement A in 1", "d3 basement A in 2"}, Type: 2},
	"d3 basement A out":         Point{Parents: []string{"d3 basement A out 1", "d3 basement A out 2"}, Type: 2},
	"d3 basement B in":          Point{Parents: []string{"d3 basement B in 1", "d3 basement B in 2"}, Type: 2},
	"d3 basement B out":         Point{Parents: []string{"d3 basement B out 1", "d3 basement B out 2"}, Type: 2},
	"d3 feather stairs":         Point{Parents: []string{"d3 feather stairs 1", "d3 feather stairs 2", "d3 feather stairs 3"}, Type: 2},
	"d3 mimic stairs":           Point{Parents: []string{"d3 mimic stairs 1", "d3 mimic stairs 2"}, Type: 2},
	"d3 trampoline stairs":      Point{Parents: []string{"d3 trampoline stairs 1", "d3 trampoline stairs 2"}, Type: 2},
	"d5 large rupee chest":      Point{Parents: []string{"d5 large rupee chest 1", "d5 large rupee chest 2"}, Type: 2},
	"d5 stairs B out":           Point{Parents: []string{"d5 stairs B out 1", "d5 stairs B out 2"}, Type: 2},
	"d6 map chest":              Point{Parents: []string{"d6 map chest 1", "d6 map chest 2"}, Type: 2},
	"d7 armos room":             Point{Parents: []string{"d7 armos room 1", "d7 armos room 2"}, Type: 2},
	"d7 cross bridge":           Point{Parents: []string{"d7 cross bridge 1", "d7 cross bridge 2", "d7 cross bridge 3"}, Type: 2},
	"d8 HSS stairs":             Point{Parents: []string{"d8 HSS stairs 1", "d8 HSS stairs 2"}, Type: 2},
	"d8 cross bridge A":         Point{Parents: []string{"d8 cross bridge A 1", "d8 cross bridge A 2"}, Type: 2},
	"d8 ice puzzle room":        Point{Parents: []string{"d8 ice puzzle room 1", "d8 ice puzzle room 2"}, Type: 2},
	"d8 portal":                 Point{Parents: []string{"d8 portal 1", "d8 portal 2"}, Type: 2},
	"dragon key cross":          Point{Parents: []string{"dragon key cross 1", "dragon key cross 2"}, Type: 2},
	"enter poe sisters":         Point{Parents: []string{"enter poe sisters 1", "enter poe sisters 2"}, Type: 2},
	"furnace":                   Point{Parents: []string{"furnace 1", "furnace 2", "furnace 3"}, Type: 2},
	"ghastly stump":             Point{Parents: []string{"ghastly stump 1", "ghastly stump 2", "ghastly stump 3", "ghastly stump 4", "ghastly stump 5"}, Type: 2},
	"goron mountain":            Point{Parents: []string{"goron mountain 1", "goron mountain 2", "goron mountain 3", "goron mountain 4"}, Type: 2},
	"graveyard":                 Point{Parents: []string{"graveyard 1", "graveyard 2"}, Type: 2},
	"hide and seek":             Point{Parents: []string{"hide and seek 1", "hide and seek 2", "hide and seek 3", "hide and seek 4"}, Type: 2},
	"horon village":             Point{Parents: []string{"horon village 1", "horon village 2", "horon village 3", "horon village 4", "horon village 5", "horon village 6"}, Type: 2},
	"kill darknut (across pit)": Point{Parents: []string{"kill darknut (across pit) 1", "kill darknut (across pit) 2"}, Type: 2},
	"lake portal":               Point{Parents: []string{"lake portal 1", "lake portal 2", "lake portal 3"}, Type: 2},
	"mount cucco":               Point{Parents: []string{"mount cucco 1", "mount cucco 2", "mount cucco 3"}, Type: 2},
	"mountain portal":           Point{Parents: []string{"mountain portal 1", "mountain portal 2"}, Type: 2},
	"mystery tree":              Point{Parents: []string{"mystery tree 1", "mystery tree 2", "mystery tree 3", "mystery tree 4"}, Type: 2},
	"natzu":                     Point{Parents: []string{"natzu 1", "natzu 2", "natzu 3"}, Type: 2},
	"open floodgate":            Point{Parents: []string{"open floodgate 1", "open floodgate 2", "open floodgate 3"}, Type: 2},
	"pegasus tree":              Point{Parents: []string{"pegasus tree 1", "pegasus tree 2", "pegasus tree 3"}, Type: 2},
	"pirate house":              Point{Parents: []string{"pirate house 1", "pirate house 2", "pirate house 3"}, Type: 2},
	"post-d2 stump":             Point{Parents: []string{"post-d2 stump 1", "post-d2 stump 2", "post-d2 stump 3", "post-d2 stump 4"}, Type: 2},
	"remains portal":            Point{Parents: []string{"remains portal 1", "remains portal 2", "remains portal 3", "remains portal 4"}, Type: 2},
	"ricky pen":                 Point{Parents: []string{"ricky pen 1", "ricky pen 2", "ricky pen 3"}, Type: 2},
	"sokra stump":               Point{Parents: []string{"sokra stump 1", "sokra stump 2", "sokra stump 3", "sokra stump 4"}, Type: 2},
	"spool swamp":               Point{Parents: []string{"spool swamp 1", "spool swamp 2", "spool swamp 3"}, Type: 2},
	"sunken city":               Point{Parents: []string{"sunken city 1", "sunken city 2", "sunken city 3"}, Type: 2},
	"swamp portal":              Point{Parents: []string{"swamp portal 1", "swamp portal 2", "swamp portal 3", "swamp portal 4"}, Type: 2},
	"temple":                    Point{Parents: []string{"temple 1", "temple 2", "temple 3", "temple 4", "temple 5"}, Type: 2},
	"temple remains":            Point{Parents: []string{"temple remains 1", "temple remains 2", "temple remains 3"}, Type: 2},
	"village portal":            Point{Parents: []string{"village portal 1", "village portal 2", "village portal 3"}, Type: 2},
}
