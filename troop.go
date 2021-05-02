// Copyright (C) 2021 Lionel Debroux
// SPDX-License-Identifier: GPL-2.0
//
// Sources of information:
// * https://forum.dune2k.com/topic/20497-dune-cheats/ : especially John2022's large post
// * hugslab
// * Dmitri Fatkin

package main

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
)

func TroopGetTroopID(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+0]
}

func TroopSetTroopID(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+0] = value
}

func TroopGetNextTroopID(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+1]
}

func TroopSetNextTroopID(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+1] = value
}

func TroopGetPosition(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+2]
}

func TroopSetPosition(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+2] = value
}

func TroopGetOccupation(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+3]
}

func TroopSetOccupation(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+3] = value
}

func TroopGetFieldE(data *DuneMetadata, i uint) uint16 {
	return uint16((*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+4]) | uint16((*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+5])*256
}

func TroopSetFieldE(data *DuneMetadata, i uint, e uint16) {
	(*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+4] = byte(e)
	(*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+5] = byte(e >> 8)
}

func TroopGetCoordinates(data *DuneMetadata, i uint) []byte {
	return (*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+6 : (data.troopsOffset+(i-1)*TROOP_SIZE)+10]
}

func TroopSetCoordinates(data *DuneMetadata, i uint, coordinates []byte) {
	copy((*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+6:(data.troopsOffset+(i-1)*TROOP_SIZE)+10], coordinates[0:4])
}

func TroopGetFieldG(data *DuneMetadata, i uint) []byte {
	return (*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+10 : (data.troopsOffset+(i-1)*TROOP_SIZE)+18]
}

func TroopSetFieldG(data *DuneMetadata, i uint, g []byte) {
	copy((*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+10:(data.troopsOffset+(i-1)*TROOP_SIZE)+18], g[0:8])
}

func TroopGetDissatisfaction(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+18]
}

func TroopSetDissatisfaction(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+18] = value
}

func TroopGetSpeech(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+19]
}

func TroopSetSpeech(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+19] = value
}

func TroopGetFieldJ(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+20]
}

func TroopSetFieldJ(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+20] = value
}

func TroopGetMotivation(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+21]
}

func TroopSetMotivation(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+21] = value
}

func TroopGetSpiceMiningSkill(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+22]
}

func TroopSetSpiceMiningSkill(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+22] = value
}

func TroopGetArmySkill(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+23]
}

func TroopSetArmySkill(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+23] = value
}

func TroopGetEcologySkill(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+24]
}

func TroopSetEcologySkill(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+24] = value
}

func TroopGetEquipment(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+25]
}

func TroopSetEquipment(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+25] = value
}

func TroopGetPopulation(data *DuneMetadata, i uint) uint16 {
	return uint16((*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+26]) * 10
}

func TroopSetPopulation(data *DuneMetadata, i uint, value uint16) {
	(*((*data).current))[(data.troopsOffset+(i-1)*TROOP_SIZE)+26] = byte((value / 10) & 0xFF)
}

func TroopSetPopulationMotivationArmySkill(data *DuneMetadata, i uint, population uint16, motivation byte, armySkill byte) {
	TroopSetPopulation(data, i, population)
	TroopSetMotivation(data, i, motivation)
	TroopSetArmySkill(data, i, armySkill)
}

func TroopSetPopulationMotivationArmySkillEquipment(data *DuneMetadata, i uint, population uint16, motivation byte, armySkill byte, equipment byte) {
	TroopSetPopulationMotivationArmySkill(data, i, population, motivation, armySkill)
	TroopSetEquipment(data, i, equipment)
}

func TroopSetSpecialDescription(data *DuneMetadata, i uint, description string) {
	((*data).specialTroopDescriptions)[i] = description
}

func TroopSwapPair(data *DuneMetadata, i1 uint, i2 uint) {
	tmp := make([]byte, TROOP_SIZE)
	copy(tmp[:], (*((*data).current))[(data.troopsOffset+(i1-1)*TROOP_SIZE):data.troopsOffset+i1*TROOP_SIZE])
	copy((*((*data).current))[(data.troopsOffset+(i1-1)*TROOP_SIZE):data.troopsOffset+i1*TROOP_SIZE], (*((*data).current))[(data.troopsOffset+(i2-1)*TROOP_SIZE):data.troopsOffset+i2*TROOP_SIZE])
	copy((*((*data).current))[(data.troopsOffset+(i2-1)*TROOP_SIZE):data.troopsOffset+i2*TROOP_SIZE], tmp[:])
	((*data).shuffledTroops)[i1] = i2
	((*data).shuffledTroops)[i2] = i1
}

// Move troop i1 to i4, i2 to i5, i3 to i6
func TroopShuffleTrio(data *DuneMetadata, i1 uint, i2 uint, i3 uint, i4 uint, i5 uint, i6 uint) {
	tmp1 := make([]byte, TROOP_SIZE)
	tmp2 := make([]byte, TROOP_SIZE)
	tmp3 := make([]byte, TROOP_SIZE)
	copy(tmp1[:], (*((*data).current))[(data.troopsOffset+(i1-1)*TROOP_SIZE):data.troopsOffset+i1*TROOP_SIZE])
	copy(tmp2[:], (*((*data).current))[(data.troopsOffset+(i2-1)*TROOP_SIZE):data.troopsOffset+i2*TROOP_SIZE])
	copy(tmp3[:], (*((*data).current))[(data.troopsOffset+(i3-1)*TROOP_SIZE):data.troopsOffset+i3*TROOP_SIZE])
	copy((*((*data).current))[(data.troopsOffset+(i4-1)*TROOP_SIZE):data.troopsOffset+i4*TROOP_SIZE], tmp1[:])
	copy((*((*data).current))[(data.troopsOffset+(i5-1)*TROOP_SIZE):data.troopsOffset+i5*TROOP_SIZE], tmp2[:])
	copy((*((*data).current))[(data.troopsOffset+(i6-1)*TROOP_SIZE):data.troopsOffset+i6*TROOP_SIZE], tmp3[:])
	((*data).shuffledTroops)[i1] = i4
	((*data).shuffledTroops)[i2] = i5
	((*data).shuffledTroops)[i3] = i6
	((*data).shuffledTroops)[i4] = i1
	((*data).shuffledTroops)[i5] = i2
	((*data).shuffledTroops)[i6] = i3
}

func TroopPrint(data *DuneMetadata, i uint) {
	// Number troops from 1
	if i >= TROOP_MIN_ID && i <= TROOP_MAX_ID {
		troopID := TroopGetTroopID(data, i)
		nextTroopID := TroopGetNextTroopID(data, i)
		position := TroopGetPosition(data, i)
		occupation := TroopGetOccupation(data, i)
		fieldE := TroopGetFieldE(data, i)
		coordinates := TroopGetCoordinates(data, i)
		fieldG := TroopGetFieldG(data, i)
		dissatisfaction := TroopGetDissatisfaction(data, i)
		speech := TroopGetSpeech(data, i)
		fieldJ := TroopGetFieldJ(data, i)
		motivation := TroopGetMotivation(data, i)
		spiceSkill := TroopGetSpiceMiningSkill(data, i)
		armySkill := TroopGetArmySkill(data, i)
		ecologySkill := TroopGetEcologySkill(data, i)
		equipment := TroopGetEquipment(data, i)
		population := TroopGetPopulation(data, i)
		fmt.Printf("Troop %02x Troop ID: %02x Next Troop ID: %02x Position: %02x Occupation: %02x\n", i, troopID, nextTroopID, position, occupation)
		fmt.Printf("\tE: %04x Coordinates: %x G: %x Dissatisfaction: %02x Speech: %02x J: %02x\n", fieldE, coordinates, fieldG, dissatisfaction, speech, fieldJ)
		fmt.Printf("\tMotivation: %02x Spice skill: %02x Army skill: %02x Ecology skill: %02x Equipment: %02x Population: %02d\n", motivation, spiceSkill, armySkill, ecologySkill, equipment, population)
	}
}

func TroopIsFremen(occupation byte) bool {
	return (occupation&0xF < 0xC) || (occupation > 0xA0)
}

func TroopDiffProduceChangelogEntry(data *DuneMetadata, i uint) (string, error) {
	//1. Modifications for Harkonnen troop 67 at Harkonnen Palace: population increased from 100 to 2500, motivation raised from 8 to 68,
	//army skill improved from 95 to 206, equipment now includes Atomic Weapons. Positioning rearranged from "on top" to "the bottom" of the palace
	//to affect possibly-liberated group.

	if i >= TROOP_MIN_ID && i <= TROOP_MAX_ID {
		changelog := ""
		// Obtain original troop data.
		(*data).current = &((*data).uncompressed)
		troopID1 := TroopGetTroopID(data, i)
		nextTroopID1 := TroopGetNextTroopID(data, i)
		position1 := TroopGetPosition(data, i)
		occupation1 := TroopGetOccupation(data, i)
		fieldE1 := TroopGetFieldE(data, i)
		coordinates1 := TroopGetCoordinates(data, i)
		fieldG1 := TroopGetFieldG(data, i)
		dissatisfaction1 := TroopGetDissatisfaction(data, i)
		speech1 := TroopGetSpeech(data, i)
		fieldJ1 := TroopGetFieldJ(data, i)
		motivation1 := TroopGetMotivation(data, i)
		spiceSkill1 := TroopGetSpiceMiningSkill(data, i)
		armySkill1 := TroopGetArmySkill(data, i)
		ecologySkill1 := TroopGetEcologySkill(data, i)
		equipment1 := TroopGetEquipment(data, i)
		population1 := TroopGetPopulation(data, i)

		// Obtain modified troop data.

		// TODO handle troop swaps and shuffles better ?

		(*data).current = &((*data).modified)
		troopID2 := TroopGetTroopID(data, i)
		nextTroopID2 := TroopGetNextTroopID(data, i)
		position2 := TroopGetPosition(data, i)
		occupation2 := TroopGetOccupation(data, i)
		fieldE2 := TroopGetFieldE(data, i)
		coordinates2 := TroopGetCoordinates(data, i)
		fieldG2 := TroopGetFieldG(data, i)
		dissatisfaction2 := TroopGetDissatisfaction(data, i)
		speech2 := TroopGetSpeech(data, i)
		fieldJ2 := TroopGetFieldJ(data, i)
		motivation2 := TroopGetMotivation(data, i)
		spiceSkill2 := TroopGetSpiceMiningSkill(data, i)
		armySkill2 := TroopGetArmySkill(data, i)
		ecologySkill2 := TroopGetEcologySkill(data, i)
		equipment2 := TroopGetEquipment(data, i)
		population2 := TroopGetPopulation(data, i)

		// If there are no modifications, bail out early.
		if (troopID1 != troopID2) || (nextTroopID1 != nextTroopID2) || (position1 != position2) || (occupation1 != occupation2) || (fieldE1 != fieldE2) || !bytes.Equal(coordinates1, coordinates2) || !bytes.Equal(fieldG1, fieldG2) || (dissatisfaction1 != dissatisfaction2) || (speech1 != speech2) || (fieldJ1 != fieldJ2) || (motivation1 != motivation2) || (spiceSkill1 != spiceSkill2) || (armySkill1 != armySkill2) || (ecologySkill1 != ecologySkill2) || (equipment1 != equipment2) || (population1 != population2) {
			var troopKind string
			if TroopIsFremen(occupation2) {
				troopKind = "Fremen"
			} else {
				troopKind = "Harkonnen"
			}
			locationID := (*data).coordinatesMap[hex.EncodeToString(coordinates2)]
			namefirst, namesecond := LocationGetName(data, locationID)

			//specialTroop68 := false
			movedTroop := ""
			troopID := byte(((*data).shuffledTroops)[uint(troopID2)])
			troopIDSetStr := ""
			if troopID2 != troopID {
				if troopID != 0 {
					movedTroop = fmt.Sprintf(" (former %d)", troopID)
					troopIDSetStr = fmt.Sprintf(" identifier of the troop changed to %d,", troopID2)
				} else {
					//specialTroop68 = true
					movedTroop = " (new troop, special)"
					troopIDSetStr = fmt.Sprintf(" identifier of the troop set to %d,", troopID2)
				}
			}
			changelog = fmt.Sprintf("Modifications for %s troop %d%s at %s:%s", troopKind, i, movedTroop, LocationGetNameStr(namefirst, namesecond), troopIDSetStr)

			// TODO handle troop 68 more specially.

			if population1 < population2 {
				changelog = changelog + fmt.Sprintf(" population increased from %d to %d,", population1, population2)
			} else if population1 > population2 {
				changelog = changelog + fmt.Sprintf(" population decreased from %d to %d,", population1, population2)
			}

			if motivation1 < motivation2 {
				changelog = changelog + fmt.Sprintf(" motivation raised from %d to %d,", motivation1, motivation2)
			} else if motivation1 > motivation2 {
				changelog = changelog + fmt.Sprintf(" motivation lowered from %d to %d,", motivation1, motivation2)
			}

			if dissatisfaction1 < dissatisfaction2 {
				changelog = changelog + fmt.Sprintf(" dissatisfaction increased from %d to %d,", dissatisfaction1, dissatisfaction2)
			} else if dissatisfaction1 > dissatisfaction2 {
				changelog = changelog + fmt.Sprintf(" dissatisfaction decreased from %d to %d,", dissatisfaction1, dissatisfaction2)
			}

			if spiceSkill1 < spiceSkill2 {
				changelog = changelog + fmt.Sprintf(" spice mining skill improved from %d to %d,", spiceSkill1, spiceSkill2)
			} else if spiceSkill1 > spiceSkill2 {
				changelog = changelog + fmt.Sprintf(" spice mining skill worsened from %d to %d,", spiceSkill1, spiceSkill2)
			}

			if armySkill1 < armySkill2 {
				changelog = changelog + fmt.Sprintf(" army skill improved from %d to %d,", armySkill1, armySkill2)
			} else if armySkill1 > armySkill2 {
				changelog = changelog + fmt.Sprintf(" army skill worsened from %d to %d,", armySkill1, armySkill2)
			}

			if ecologySkill1 < ecologySkill2 {
				changelog = changelog + fmt.Sprintf(" ecology skill improved from %d to %d,", ecologySkill1, ecologySkill2)
			} else if ecologySkill1 > ecologySkill2 {
				changelog = changelog + fmt.Sprintf(" ecology skill worsened from %d to %d,", ecologySkill1, ecologySkill2)
			}

			// TODO equipment
			// TODO position
			// TODO etc.
			// TODO special description.

			// Trim final ","
			if changelog[len(changelog)-1] == ',' {
				changelog = changelog[0 : len(changelog)-1]
				changelog += "."
			}
		}

		return changelog, nil
	} else {
		return "", errors.New("Improper troop ID")
	}
}

func TroopDiffAllProduceChangelogEntries(data *DuneMetadata) (string, error) {
	changelogHarkonnen := ""
	changelogFremen := ""
	for i := uint(TROOP_MIN_ID); i <= TROOP_MAX_ID; i++ {
		changelogTroop, err := TroopDiffProduceChangelogEntry(data, i)
		if err != nil {
			return "", err
		}
		if TroopIsFremen(TroopGetOccupation(data, i)) {
			if changelogTroop != "" {
				changelogFremen = changelogFremen + "\n" + changelogTroop + "\n"
			}
		} else {
			if changelogTroop != "" {
				changelogHarkonnen = changelogHarkonnen + "\n" + changelogTroop + "\n"
			}
		}
	}

	return changelogHarkonnen + changelogFremen, nil
}
