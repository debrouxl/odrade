// Copyright (C) 2021 Lionel Debroux
// SPDX-License-Identifier: GPL-2.0
//
// Sources of information:
// * https://forum.dune2k.com/topic/20497-dune-cheats/ : especially John2022's large post
// * hugslab
// * Dmitri Fatkin

package main

import (
	"errors"
	"fmt"
)

func SmugglerGetRegion(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+0]
}

func SmugglerSetRegion(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+0] = value
}

func SmugglerGetWillingnessToHaggle(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+1]
}

func SmugglerSetWillingnessToHaggle(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+1] = value
}

func SmugglerGetFieldC(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+2]
}

func SmugglerSetFieldC(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+2] = value
}

func SmugglerGetFieldD(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+3]
}

func SmugglerSetFieldD(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+3] = value
}

func SmugglerGetHarvesters(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+4]
}

func SmugglerSetHarvesters(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+4] = value
}

func SmugglerGetOrnithopters(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+5]
}

func SmugglerSetOrnithopters(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+5] = value
}

func SmugglerGetKrysKnives(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+6]
}

func SmugglerSetKrysKnives(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+6] = value
}

func SmugglerGetLaserGuns(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+7]
}

func SmugglerSetLaserGuns(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+7] = value
}

func SmugglerGetWeirdingModules(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+8]
}

func SmugglerSetWeirdingModules(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+8] = value
}

// TODO massage the input and output values like the game engine does.
func SmugglerGetHarvestersPrice(data *DuneMetadata, i uint) uint16 {
	return uint16((*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+9]) * 10
}

func SmugglerSetHarvestersPrice(data *DuneMetadata, i uint, value uint16) {
	(*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+9] = byte(value / 10)
}

func SmugglerGetOrnithoptersPrice(data *DuneMetadata, i uint) uint16 {
	return uint16((*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+10]) * 10
}

func SmugglerSetOrnithoptersPrice(data *DuneMetadata, i uint, value uint16) {
	(*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+10] = byte(value / 10)
}

func SmugglerGetKrysKnivesPrice(data *DuneMetadata, i uint) uint16 {
	return uint16((*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+11]) * 10
}

func SmugglerSetKrysKnivesPrice(data *DuneMetadata, i uint, value uint16) {
	(*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+11] = byte(value / 10)
}

func SmugglerGetLaserGunsPrice(data *DuneMetadata, i uint) uint16 {
	return uint16((*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+12]) * 10
}

func SmugglerSetLaserGunsPrice(data *DuneMetadata, i uint, value uint16) {
	(*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+12] = byte(value / 10)
}

func SmugglerGetWeirdingModulesPrice(data *DuneMetadata, i uint) uint16 {
	return uint16((*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+13]) * 10
}

func SmugglerSetWeirdingModulesPrice(data *DuneMetadata, i uint, value uint16) {
	(*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+13] = byte(value / 10)
}

func SmugglerPrint(data *DuneMetadata, i uint) {
	if i >= SMUGGLER_MIN_ID && i <= SMUGGLER_MAX_ID {
		region := SmugglerGetRegion(data, i)
		namestr := LocationGetNameStr(region, 0x0A)
		willingnessToHaggle := SmugglerGetWillingnessToHaggle(data, i)
		fieldC := SmugglerGetFieldC(data, i)
		fieldD := SmugglerGetFieldD(data, i)
		harvesters := SmugglerGetHarvesters(data, i)
		ornithopters := SmugglerGetOrnithopters(data, i)
		krysKnives := SmugglerGetKrysKnives(data, i)
		laserGuns := SmugglerGetLaserGuns(data, i)
		weirdingModules := SmugglerGetWeirdingModules(data, i)
		harvestersPrice := SmugglerGetHarvestersPrice(data, i)
		ornithoptersPrice := SmugglerGetOrnithoptersPrice(data, i)
		krysKnivesPrice := SmugglerGetKrysKnivesPrice(data, i)
		laserGunsPrice := SmugglerGetLaserGunsPrice(data, i)
		weirdingModulesPrice := SmugglerGetWeirdingModulesPrice(data, i)
		fmt.Printf("Smuggler %02x Location: %s Willingness to haggle: %x C: %02x D: %02x\n", i, namestr, willingnessToHaggle, fieldC, fieldD)
		fmt.Printf("\tHarvesters: %02x Ornithopters: %02x Krys knives: %02x Laser guns: %02x Weirding modules: %02x\n", harvesters, ornithopters, krysKnives, laserGuns, weirdingModules)
		fmt.Printf("\tHarvesters price: %d Ornithopters price: %d Krys knives price: %d Laser guns price: %d Weirding modules price: %d\n", harvestersPrice, ornithoptersPrice, krysKnivesPrice, laserGunsPrice, weirdingModulesPrice)
		println("NOTE: prices not adjusted yet")
	}
}

func SmugglerDiffProduceChangelogEntry(data *DuneMetadata, i uint) (string, error) {
	if i >= SMUGGLER_MIN_ID && i <= SMUGGLER_MAX_ID {
		changelog := ""

		return changelog, nil
	} else {
		return "", errors.New("Improper smuggler ID")
	}
}

func SmugglerDiffAllProduceChangelogEntries(data *DuneMetadata) (string, error) {
	changelog := "TODO"
	for i := uint(SMUGGLER_MIN_ID); i <= SMUGGLER_MAX_ID; i++ {
		changelogSmuggler, err := SmugglerDiffProduceChangelogEntry(data, i)
		if err != nil {
			return "", err
		}
		changelog = changelog + changelogSmuggler
	}

	return changelog, nil
}
