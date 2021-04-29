// Copyright (C) 2021 Lionel Debroux
// SPDX-License-Identifier: GPL-2.0
//
// Sources of information:
// * https://forum.dune2k.com/topic/20497-dune-cheats/
// * hugslab
// * Dmitri Fatkin

package main

// For testing purposes, primarily round-trip decompression + compression sanity check: a modification function which does nothing.
func ModifyTroopAndLocationData(data *DuneMetadata) error {
	/*LocationPrint(data, mode, 1)
	LocationPrint(data, mode, 2)
	TroopPrint(data, mode, 1)
	TroopPrint(data, mode, 67)*/
	return nil
}

func ModifyNPCAndSmugglerData(data *DuneMetadata) error {
	return nil
}

func ChangelogPrologue() (string, error) {
	return "", nil
}

func ChangelogEpilogue() (string, error) {
	return "", nil
}
