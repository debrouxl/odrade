# Original Dune Restructuring And Diversification Extension

## Goal and remarks

This repository aims at providing a framework for modifying save files for the original Dune 1 game by Cryo Interactive. Five known versions distinguished by the names of the save files they create are handled: DUNE21, DUNE23, DUNE24 (floppy versions), DUNE37 and DUNE38 (CD versions).

The framework was originally designed for reimplementing later versions of Dmitri Fatkin's A Harder Path in a scriptable way, and as such, it aims at modifying DUNE{21,23,24,37,38}S0.SAV files. However, with a bit of elbow grease, it could be used for other save file modifications, at different game points. Yes, you *could* be a dirty cheater increasing spice density and amount, raising Fremen troops' skill and equipment, lowering Harkonnen troops' skill levels and equipment, reducing equipment price at smugglers - completely spoiling the game balance, in other words.

However, note that:
* if you have to resort to making the game *easier* than vanilla is, it means that your game strategy is seriously suboptimal, possibly because you've missed some aspect of the game (I sure used to as well, in the 1990s!), and you need to explore the game more and rethink your strategy from the ground up. Indeed, in vanilla Dune, one can easily destroy Harkonnen fortresses before day 30, using only the Fremen troops from the northern part of the planet + the ones found after winning a battle in a Harkonnen fortress on army tasks, the other Fremen troops on spice mining or prospecting tasks (no troops planting bulbs), buying few weapons from smugglers, and not even using improved training for the army troops ;)
* improper save file modifications can yield improper NPC dialog state (e.g. when forcing the visibility of sietches in a way which does not match the current game stage) or even game crashes... you have been warned :)

## Capabilities

For the initial release of ODRADE, four types of game data can be modified:
* location data: 18 pieces of information about Fremen sietches and Harkonnen fortresses, e.g. coordinates, visibility, spice density and amount, equipment;
* troop data: 16 pieces of information about Fremen and Harkonnen troops, e.g. type and occupation, coordinates, skill levels, equipment;
* NPC data: 8 pieces of information about the main characters in the storyline or generic groups of characters, e.g. location, dialogue;
* smuggler data: 10 pieces of information about the respective smugglers, e.g. willingness to haggle over the price, item price, equipment.

Changelog generation is automated.

## Other notes

The framework is written in Golang, primarily because I needed to make practice in that language, but also because it's arguably relatively easy to program, and the toolchain supports combined build + invocation in a single command line without saving a binary: IOW, invoking `go run ...` isn't really different from invoking `python3 ...`  or another script interpreter.

## Authors

* Dmitri Fatkin: A Harder Path save game modifications, *Original Dune Restructuring And Diversification Extension* acronym;
* Lionel Debroux: ODRADE code, A Harder Path modding scripts.

License: GPLv2.

# Usage

## Modding program

The framework uses a callback-based design: any modding program needs to implement several callbacks. See `verbatim.go`, the neutral modder, for an up to date list.

Sample invocations:

```
for file in DUNE21S0.SAV DUNE23S0.SAV DUNE24S0.SAV DUNE37S0.SAV DUNE38S0.SAV; do
    (path to `go`, /usr/lib/go-1.16/bin/go for me) run odrade.go verbatim.go location.go troop.go npc.go smuggler.go print $file; # just print out the contents of the original save file.
    (path to `go`                                ) run odrade.go ahp20c.go location.go troop.go npc.go smuggler.go modify $file; # apply modder A Harder Path version 20c onto the original save file. 
    (path to `go`                                ) run odrade.go (yourmodder.go) location.go troop.go npc.go smuggler.go modify $file;
done
```

## Feeding modified save files into your Dune 1 install

1. start the game, skip the intro with ESC key;
2. replace the `DUNE??S0.SAV` file inside your Dune 1 save file directory by the `DUNE??S0_compressed.SAV` file produced by the modder;
3. go to the mirror room (next room up from where the story began);
4. look at the mirror;
5. restart the game.
