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
		fmt.Printf("Troop %d (%X) Troop ID: %02X Next Troop ID: %02X Position: %02X Occupation: %02X\n", i, i, troopID, nextTroopID, position, occupation)
		fmt.Printf("\tE: %04X Coordinates: %X G: %X Dissatisfaction: %02X Speech: %02X J: %02X\n", fieldE, coordinates, fieldG, dissatisfaction, speech, fieldJ)
		fmt.Printf("\tMotivation: %02X Spice skill: %02X Army skill: %02X Ecology skill: %02X Equipment: %02X Population: %02d (%02X)\n", motivation, spiceSkill, armySkill, ecologySkill, equipment, population, population)
	}
}

func TroopIsFremen(occupation byte) bool {
	return (occupation&0xF < 0xC) || (occupation > 0xA0)
}

func TroopDiffProduceChangelogEntry(data *DuneMetadata, i uint) (string, error) {
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

		// TODO handle troop swaps and shuffles better.

		// Obtain modified troop data.
		j := i
		(*data).current = &((*data).modified)
		troopID2 := TroopGetTroopID(data, j)
		nextTroopID2 := TroopGetNextTroopID(data, j)
		occupation2 := TroopGetOccupation(data, j)
		/*if occupation1 != occupation2 {
			// Troop shuffled in such a way that it switched from Fremen to Harkonnen or from Harkonnen to Fremen.
			// We might get a better diff if we use the original ID instead... well, not so simple :)
			j = data.shuffledTroops[i]
		}*/
		position2 := TroopGetPosition(data, j)
		fieldE2 := TroopGetFieldE(data, j)
		coordinates2 := TroopGetCoordinates(data, j)
		fieldG2 := TroopGetFieldG(data, j)
		dissatisfaction2 := TroopGetDissatisfaction(data, j)
		speech2 := TroopGetSpeech(data, j)
		fieldJ2 := TroopGetFieldJ(data, j)
		motivation2 := TroopGetMotivation(data, j)
		spiceSkill2 := TroopGetSpiceMiningSkill(data, j)
		armySkill2 := TroopGetArmySkill(data, j)
		ecologySkill2 := TroopGetEcologySkill(data, j)
		equipment2 := TroopGetEquipment(data, j)
		population2 := TroopGetPopulation(data, j)

		// If there are no modifications, bail out early.
		if (troopID1 != troopID2) || (nextTroopID1 != nextTroopID2) || (position1 != position2) || (occupation1 != occupation2) || (fieldE1 != fieldE2) || !bytes.Equal(coordinates1, coordinates2) || !bytes.Equal(fieldG1, fieldG2) || (dissatisfaction1 != dissatisfaction2) || (speech1 != speech2) || (fieldJ1 != fieldJ2) || (motivation1 != motivation2) || (spiceSkill1 != spiceSkill2) || (armySkill1 != armySkill2) || (ecologySkill1 != ecologySkill2) || (equipment1 != equipment2) || (population1 != population2) {
			var troopKind string
			if TroopIsFremen(occupation2) {
				troopKind = "Fremen"
			} else {
				troopKind = "Harkonnen"
			}
			locationID1Str := (*data).coordinatesMap[hex.EncodeToString(coordinates1)]
			namefirst1, namesecond1 := LocationGetName(data, locationID1Str)
			locationID2Str := (*data).coordinatesMap[hex.EncodeToString(coordinates2)]
			namefirst2, namesecond2 := LocationGetName(data, locationID2Str)

			//specialTroop68 := false
			movedTroop := ""
			troopID := byte(((*data).shuffledTroops)[uint(troopID2)])
			troopIDSetStr := ""
			fromStr := ""
			if troopID2 != troopID {
				if troopID != 0 {
					movedTroop = fmt.Sprintf(" (former %d)", troopID)
					troopIDSetStr = fmt.Sprintf(" identifier of the troop changed to %d (%X),", troopID2, troopID2)
				} else {
					//specialTroop68 = true
					movedTroop = " (new troop, special)"
					troopIDSetStr = fmt.Sprintf(" identifier of the troop set to %d,", troopID2)
				}
			}
			if locationID1Str != locationID2Str && troopID2 == troopID { // Do not print the origin sietch if the troop was already marked as moved.
				fromStr = fmt.Sprintf(" (from %s)", LocationGetNameStr(namefirst1, namesecond1))
			}
			changelog = fmt.Sprintf("Modifications for %s troop %d (%X)%s%s at %s:%s", troopKind, i, i, movedTroop, fromStr, LocationGetNameStr(namefirst2, namesecond2), troopIDSetStr)

			// TODO handle troop 68 more specially.

			if occupation1 < occupation2 {
				changelog += " occupation changed to Harkonnen,"
			} else if occupation1 > occupation2 {
				changelog += " occupation changed to Fremen,"
			}

			if population1 < population2 {
				changelog += fmt.Sprintf(" population increased from %d (%X) to %d (%X),", population1, population1 / 10, population2, population2 / 10)
			} else if population1 > population2 {
				changelog += fmt.Sprintf(" population decreased from %d (%X) to %d (%X),", population1, population1 / 10, population2, population2 / 10)
			}

			if motivation1 < motivation2 {
				changelog += fmt.Sprintf(" motivation raised from %d (%X) to %d (%X),", motivation1, motivation1, motivation2, motivation2)
			} else if motivation1 > motivation2 {
				changelog += fmt.Sprintf(" motivation lowered from %d (%X) to %d (%X),", motivation1, motivation1, motivation2, motivation2)
			}

			if dissatisfaction1 < dissatisfaction2 {
				changelog += fmt.Sprintf(" dissatisfaction increased from %d (%X) to %d (%X),", dissatisfaction1, dissatisfaction1, dissatisfaction2, dissatisfaction2)
			} else if dissatisfaction1 > dissatisfaction2 {
				changelog += fmt.Sprintf(" dissatisfaction decreased from %d (%X) to %d (%X),", dissatisfaction1, dissatisfaction1, dissatisfaction2, dissatisfaction2)
			}

			if spiceSkill1 < spiceSkill2 {
				changelog += fmt.Sprintf(" spice mining skill improved from %d (%X) to %d (%X),", spiceSkill1, spiceSkill1, spiceSkill2, spiceSkill2)
			} else if spiceSkill1 > spiceSkill2 {
				changelog += fmt.Sprintf(" spice mining skill worsened from %d (%X) to %d (%X),", spiceSkill1, spiceSkill1, spiceSkill2, spiceSkill2)
			}

			if armySkill1 < armySkill2 {
				changelog += fmt.Sprintf(" army skill improved from %d (%X) to %d (%X),", armySkill1, armySkill1, armySkill2, armySkill2)
			} else if armySkill1 > armySkill2 {
				changelog += fmt.Sprintf(" army skill worsened from %d (%X) to %d (%X),", armySkill1, armySkill1, armySkill2, armySkill2)
			}

			if ecologySkill1 < ecologySkill2 {
				changelog += fmt.Sprintf(" ecology skill improved from %d (%X) to %d (%X),", ecologySkill1, ecologySkill1, ecologySkill2, ecologySkill2)
			} else if ecologySkill1 > ecologySkill2 {
				changelog += fmt.Sprintf(" ecology skill worsened from %d (%X) to %d (%X),", ecologySkill1, ecologySkill1, ecologySkill2, ecologySkill2)
			}

			if equipment1 != equipment2 {
				additionalEquipment := equipment2 & (0xFF ^ equipment1)
				removedEquipment := equipment1 & (0xFF ^ equipment2)
				if additionalEquipment != 0 {
					changelog += " equipment now includes"
					if additionalEquipment&TROOP_EQUIPMENT_BULB == TROOP_EQUIPMENT_BULB {
						changelog += " Bulbs,"
					}
					if additionalEquipment&TROOP_EQUIPMENT_ORNITHOPTER == TROOP_EQUIPMENT_ORNITHOPTER {
						changelog += " Ornithopter,"
					}
					if additionalEquipment&TROOP_EQUIPMENT_HARVESTER == TROOP_EQUIPMENT_HARVESTER {
						changelog += " Harvester,"
					}
					if additionalEquipment&TROOP_EQUIPMENT_KRYS_KNIVES == TROOP_EQUIPMENT_KRYS_KNIVES {
						changelog += " Krys knives,"
					}
					if additionalEquipment&TROOP_EQUIPMENT_LASER_GUNS == TROOP_EQUIPMENT_LASER_GUNS {
						changelog += " Laser guns,"
					}
					if additionalEquipment&TROOP_EQUIPMENT_WEIRDING_MODULES == TROOP_EQUIPMENT_WEIRDING_MODULES {
						changelog += " Weirding modules,"
					}
					if additionalEquipment&TROOP_EQUIPMENT_ATOMICS == TROOP_EQUIPMENT_ATOMICS {
						changelog += " Atomics,"
					}
				}
				if removedEquipment != 0 {
					changelog += " equipment no longer includes"
					if removedEquipment&TROOP_EQUIPMENT_BULB == TROOP_EQUIPMENT_BULB {
						changelog += " Bulbs,"
					}
					if removedEquipment&TROOP_EQUIPMENT_ORNITHOPTER == TROOP_EQUIPMENT_ORNITHOPTER {
						changelog += " Ornithopter,"
					}
					if removedEquipment&TROOP_EQUIPMENT_HARVESTER == TROOP_EQUIPMENT_HARVESTER {
						changelog += " Harvester,"
					}
					if removedEquipment&TROOP_EQUIPMENT_KRYS_KNIVES == TROOP_EQUIPMENT_KRYS_KNIVES {
						changelog += " Krys knives,"
					}
					if removedEquipment&TROOP_EQUIPMENT_LASER_GUNS == TROOP_EQUIPMENT_LASER_GUNS {
						changelog += " Laser guns,"
					}
					if removedEquipment&TROOP_EQUIPMENT_WEIRDING_MODULES == TROOP_EQUIPMENT_WEIRDING_MODULES {
						changelog += " Weirding modules,"
					}
					if removedEquipment&TROOP_EQUIPMENT_ATOMICS == TROOP_EQUIPMENT_ATOMICS {
						changelog += " Atomics,"
					}
				}
			}

			// Trim final ","
			if changelog[len(changelog)-1] == ',' {
				changelog = changelog[0 : len(changelog)-1]
				changelog += "."
			}

			specialDescription := data.specialTroopDescriptions[i]
			if specialDescription != "" {
				changelog += specialDescription
			}
		}

		return changelog, nil
	} else {
		return "", errors.New("Improper troop ID")
	}
}

func TroopDiffAllProduceChangelogEntries(data *DuneMetadata) (string, error) {
	changelog := `TROOP ADJUSTMENTS:

`
	changelogHarkonnen := ""
	changelogFremen := ""
	for i := uint(TROOP_MIN_ID); i <= TROOP_MAX_ID; i++ {
		changelogTroop, err := TroopDiffProduceChangelogEntry(data, i)
		if err != nil {
			return "", err
		}
		if TroopIsFremen(TroopGetOccupation(data, i)) {
			if changelogTroop != "" {
				changelogFremen += "\n" + changelogTroop + "\n"
			}
		} else {
			if changelogTroop != "" {
				changelogHarkonnen += "\n" + changelogTroop + "\n"
			}
		}
	}

	return changelog + changelogHarkonnen + changelogFremen, nil
}
