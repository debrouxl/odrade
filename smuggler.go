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

// TODO
func SmugglerMapHarvestersPrice(price uint16, write bool) uint16 {
	return price
}

func SmugglerGetHarvestersPrice(data *DuneMetadata, i uint) uint16 {
	return SmugglerMapHarvestersPrice(uint16((*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+9])*10, false)
}

func SmugglerSetHarvestersPrice(data *DuneMetadata, i uint, value uint16) {
	(*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+9] = byte(SmugglerMapHarvestersPrice(value, true) / 10)
}

// TODO
func SmugglerMapOrnithoptersPrice(price uint16, write bool) uint16 {
	return price
}

func SmugglerGetOrnithoptersPrice(data *DuneMetadata, i uint) uint16 {
	return SmugglerMapOrnithoptersPrice(uint16((*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+10])*10, false)
}

func SmugglerSetOrnithoptersPrice(data *DuneMetadata, i uint, value uint16) {
	(*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+10] = byte(SmugglerMapOrnithoptersPrice(value, true) / 10)
}

// TODO
func SmugglerMapKrysKnivesPrice(price uint16, write bool) uint16 {
	return price
}

func SmugglerGetKrysKnivesPrice(data *DuneMetadata, i uint) uint16 {
	return SmugglerMapKrysKnivesPrice(uint16((*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+11])*10, false)
}

func SmugglerSetKrysKnivesPrice(data *DuneMetadata, i uint, value uint16) {
	(*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+11] = byte(SmugglerMapKrysKnivesPrice(value, true) / 10)
}

// TODO
func SmugglerMapLaserGunsPrice(price uint16, write bool) uint16 {
	return price
}

func SmugglerGetLaserGunsPrice(data *DuneMetadata, i uint) uint16 {
	return SmugglerMapLaserGunsPrice(uint16((*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+12])*10, false)
}

func SmugglerSetLaserGunsPrice(data *DuneMetadata, i uint, value uint16) {
	(*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+12] = byte(SmugglerMapLaserGunsPrice(value, true) / 10)
}

// TODO
func SmugglerMapWeirdingModulesPrice(price uint16, write bool) uint16 {
	return price
}

func SmugglerGetWeirdingModulesPrice(data *DuneMetadata, i uint) uint16 {
	return SmugglerMapWeirdingModulesPrice(uint16((*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+13])*10, false)
}

func SmugglerSetWeirdingModulesPrice(data *DuneMetadata, i uint, value uint16) {
	(*((*data).current))[(data.smugglersOffset+(i-1)*(SMUGGLER_SIZE+SMUGGLER_PADDING))+13] = byte(SmugglerMapWeirdingModulesPrice(value, true) / 10)
}

func SmugglerSetSpecialDescription(data *DuneMetadata, i uint, description string) {
	((*data).specialSmugglerDescriptions)[i] = description
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

		// Obtain original smuggler data.
		(*data).current = &((*data).uncompressed)
		region1 := SmugglerGetRegion(data, i)
		willingnessToHaggle1 := SmugglerGetWillingnessToHaggle(data, i)
		fieldC1 := SmugglerGetFieldC(data, i)
		fieldD1 := SmugglerGetFieldD(data, i)
		harvesters1 := SmugglerGetHarvesters(data, i)
		ornithopters1 := SmugglerGetOrnithopters(data, i)
		krysKnives1 := SmugglerGetKrysKnives(data, i)
		laserGuns1 := SmugglerGetLaserGuns(data, i)
		weirdingModules1 := SmugglerGetWeirdingModules(data, i)
		harvestersPrice1 := SmugglerGetHarvestersPrice(data, i)
		ornithoptersPrice1 := SmugglerGetOrnithoptersPrice(data, i)
		krysKnivesPrice1 := SmugglerGetKrysKnivesPrice(data, i)
		laserGunsPrice1 := SmugglerGetLaserGunsPrice(data, i)
		weirdingModulesPrice1 := SmugglerGetWeirdingModulesPrice(data, i)

		// Obtain modified smuggler data.
		(*data).current = &((*data).modified)
		region2 := SmugglerGetRegion(data, i)
		willingnessToHaggle2 := SmugglerGetWillingnessToHaggle(data, i)
		fieldC2 := SmugglerGetFieldC(data, i)
		fieldD2 := SmugglerGetFieldD(data, i)
		harvesters2 := SmugglerGetHarvesters(data, i)
		ornithopters2 := SmugglerGetOrnithopters(data, i)
		krysKnives2 := SmugglerGetKrysKnives(data, i)
		laserGuns2 := SmugglerGetLaserGuns(data, i)
		weirdingModules2 := SmugglerGetWeirdingModules(data, i)
		harvestersPrice2 := SmugglerGetHarvestersPrice(data, i)
		ornithoptersPrice2 := SmugglerGetOrnithoptersPrice(data, i)
		krysKnivesPrice2 := SmugglerGetKrysKnivesPrice(data, i)
		laserGunsPrice2 := SmugglerGetLaserGunsPrice(data, i)
		weirdingModulesPrice2 := SmugglerGetWeirdingModulesPrice(data, i)

		if (region1 != region2) || (willingnessToHaggle1 != willingnessToHaggle2) || (fieldC1 != fieldC2) || (fieldD1 != fieldD2) || (harvesters1 != harvesters2) || (ornithopters1 != ornithopters2) || (krysKnives1 != krysKnives2) || (laserGuns1 != laserGuns2) || (weirdingModules1 != weirdingModules2) || (harvestersPrice1 != harvestersPrice2) || (ornithoptersPrice1 != ornithoptersPrice2) || (krysKnivesPrice1 != krysKnivesPrice2) || (laserGunsPrice1 != laserGunsPrice2) || (weirdingModulesPrice1 != weirdingModulesPrice2) {
			changelog = "Smuggler at " + LocationGetNameStr(region2, 0x0A) + ":"

			// Trim final ","
			if changelog[len(changelog)-1] == ',' {
				changelog = changelog[0 : len(changelog)-1]
				changelog += "."
			}

			specialDescription := data.specialSmugglerDescriptions[i]
			if specialDescription != "" {
				changelog += specialDescription
			}
		}

		return changelog, nil
	} else {
		return "", errors.New("Improper smuggler ID")
	}
}

func SmugglerDiffAllProduceChangelogEntries(data *DuneMetadata) (string, error) {
	changelog := `SMUGGLER ADJUSTMENTS:

`
	for i := uint(SMUGGLER_MIN_ID); i <= SMUGGLER_MAX_ID; i++ {
		changelogSmuggler, err := SmugglerDiffProduceChangelogEntry(data, i)
		if err != nil {
			return "", err
		}
		if changelogSmuggler != "" {
			changelog += changelogSmuggler + "\n"
		}
	}

	if changelog[len(changelog)-1] == '\n' && changelog[len(changelog)-2] != '\n' {
		changelog += "\n"
	}
	return changelog, nil
}
