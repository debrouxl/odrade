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
	"errors"
	"fmt"
)

func LocationGetNameStr(first byte, second byte) string {
	var str string

	switch first {
	case 0x01:
		str = "Arrakeen"
	case 0x02:
		str = "Carthag"
	case 0x03:
		str = "Tuono"
	case 0x04:
		str = "Habbanya"
	case 0x05:
		str = "Oxtyn"
	case 0x06:
		str = "Tsympo"
	case 0x07:
		str = "Bledan"
	case 0x08:
		str = "Ergsun"
	case 0x09:
		str = "Haga"
	case 0x0a:
		str = "Cielago"
	case 0x0b:
		str = "Sihaya"
	case 0x0c:
		str = "Celimyn"
	}

	switch second {
	case 0x01:
		str += " (Atreides)"
	case 0x02:
		str += " (Harkonnen)"
	case 0x03:
		str += "-Tabr"
	case 0x04:
		str += "-Timin"
	case 0x05:
		str += "-Tuek"
	case 0x06:
		str += "-Harg"
	case 0x07:
		str += "-Clam"
	case 0x08:
		str += "-Tsymyn"
	case 0x09:
		str += "-Siet"
	case 0x0a:
		str += "-Pyons"
	case 0x0b:
		str += "-Pyort"
	}

	return str
}

func LocationGetName(data *DuneMetadata, i uint) (byte, byte) {
	return (*((*data).current))[(data.locationsOffset+(i-1)*LOCATION_SIZE)+0], (*((*data).current))[(data.locationsOffset+(i-1)*LOCATION_SIZE)+1]
}

func LocationGetCoordinates(data *DuneMetadata, i uint) []byte {
	return (*((*data).current))[(data.locationsOffset+(i-1)*LOCATION_SIZE)+2 : (data.locationsOffset+(i-1)*LOCATION_SIZE)+8]
}

func LocationGetAppearance(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.locationsOffset+(i-1)*LOCATION_SIZE)+8]
}

func LocationSetAppearance(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.locationsOffset+(i-1)*LOCATION_SIZE)+8] = value
}

func LocationGetTroopID(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.locationsOffset+(i-1)*LOCATION_SIZE)+9]
}

func LocationSetTroopID(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.locationsOffset+(i-1)*LOCATION_SIZE)+9] = value
}

func LocationGetStatus(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.locationsOffset+(i-1)*LOCATION_SIZE)+10]
}

func LocationSetStatus(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.locationsOffset+(i-1)*LOCATION_SIZE)+10] = value
}

func LocationGetStage(data *DuneMetadata, i uint) []byte {
	return (*((*data).current))[(data.locationsOffset+(i-1)*LOCATION_SIZE)+11 : (data.locationsOffset+(i-1)*LOCATION_SIZE)+16]
}

func LocationSetStage(data *DuneMetadata, i uint, stage []byte) {
	copy((*((*data).current))[(data.locationsOffset+(i-1)*LOCATION_SIZE)+11:(data.locationsOffset+(i-1)*LOCATION_SIZE)+16], stage[0:5])
}

func LocationGetFieldID(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.locationsOffset+(i-1)*LOCATION_SIZE)+16]
}

func LocationGetSpiceAmount(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.locationsOffset+(i-1)*LOCATION_SIZE)+17]
}

func LocationSetSpiceAmount(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.locationsOffset+(i-1)*LOCATION_SIZE)+17] = value
}

func LocationGetSpiceDensity(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.locationsOffset+(i-1)*LOCATION_SIZE)+18]
}

func LocationSetSpiceDensity(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.locationsOffset+(i-1)*LOCATION_SIZE)+18] = value
}

func LocationGetFieldJ(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.locationsOffset+(i-1)*LOCATION_SIZE)+19]
}

func LocationGetHarvesters(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.locationsOffset+(i-1)*LOCATION_SIZE)+20]
}

func LocationSetHarvesters(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.locationsOffset+(i-1)*LOCATION_SIZE)+20] = value
}

func LocationGetOrnithopters(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.locationsOffset+(i-1)*LOCATION_SIZE)+21]
}

func LocationSetOrnithopters(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.locationsOffset+(i-1)*LOCATION_SIZE)+21] = value
}

func LocationGetKrysKnives(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.locationsOffset+(i-1)*LOCATION_SIZE)+22]
}

func LocationSetKrysKnives(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.locationsOffset+(i-1)*LOCATION_SIZE)+22] = value
}

func LocationGetLaserGuns(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.locationsOffset+(i-1)*LOCATION_SIZE)+23]
}

func LocationSetLaserGuns(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.locationsOffset+(i-1)*LOCATION_SIZE)+23] = value
}

func LocationGetWeirdingModules(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.locationsOffset+(i-1)*LOCATION_SIZE)+24]
}

func LocationSetWeirdingModules(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.locationsOffset+(i-1)*LOCATION_SIZE)+24] = value
}

func LocationGetAtomics(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.locationsOffset+(i-1)*LOCATION_SIZE)+25]
}

func LocationSetAtomics(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.locationsOffset+(i-1)*LOCATION_SIZE)+25] = value
}

func LocationGetBulbs(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.locationsOffset+(i-1)*LOCATION_SIZE)+26]
}

func LocationGetWater(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.locationsOffset+(i-1)*LOCATION_SIZE)+27]
}

func LocationSetWater(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.locationsOffset+(i-1)*LOCATION_SIZE)+27] = value
}

func LocationSetWeapons(data *DuneMetadata, i uint, krysKnives byte, laserGuns byte, weirdingModules byte, atomics byte) {
	LocationSetKrysKnives(data, i, krysKnives)
	LocationSetLaserGuns(data, i, laserGuns)
	LocationSetWeirdingModules(data, i, weirdingModules)
	LocationSetAtomics(data, i, atomics)
}

func LocationSetSpecialDescription(data *DuneMetadata, i uint, description string) {
	((*data).specialLocationDescriptions)[i] = description
}

func LocationPrint(data *DuneMetadata, i uint) {
	if i >= LOCATION_MIN_ID && i <= LOCATION_MAX_ID {
		namefirst, namesecond := LocationGetName(data, i)
		namestr := LocationGetNameStr(namefirst, namesecond)
		coordinates := LocationGetCoordinates(data, i)
		appearance := LocationGetAppearance(data, i)
		troopID := LocationGetTroopID(data, i)
		status := LocationGetStatus(data, i)
		stage := LocationGetStage(data, i)
		spiceFieldID := LocationGetFieldID(data, i)
		spiceAmount := LocationGetSpiceAmount(data, i)
		spiceDensity := LocationGetSpiceDensity(data, i)
		fieldJ := LocationGetFieldJ(data, i)
		harvesters := LocationGetHarvesters(data, i)
		ornithopters := LocationGetOrnithopters(data, i)
		krysKnives := LocationGetKrysKnives(data, i)
		laserGuns := LocationGetLaserGuns(data, i)
		weirdingModules := LocationGetWeirdingModules(data, i)
		atomics := LocationGetAtomics(data, i)
		bulbs := LocationGetBulbs(data, i)
		water := LocationGetWater(data, i)
		fmt.Printf("Location %d (%02x) Name: %s Coordinates: %x\n", i, i, namestr, coordinates)
		fmt.Printf("\tAppearance: %02x Troop ID: %02x Status: %02x Stage: %x\n\tSpice field ID: %02x Spice amount: %02x Spice density: %02x J: %02x\n", appearance, troopID, status, stage, spiceFieldID, spiceAmount, spiceDensity, fieldJ)
		fmt.Printf("\tHarvesters: %02x Ornithopters: %02x Krys knives: %02x Laser guns: %02x Weirding modules: %02x Atomics: %02x Bulbs: %02x Water: %02x\n", harvesters, ornithopters, krysKnives, laserGuns, weirdingModules, atomics, bulbs, water)
	}
}

func LocationDiffProduceChangelogEntry(data *DuneMetadata, i uint) (string, error) {
	if i >= LOCATION_MIN_ID && i <= LOCATION_MAX_ID {
		changelog := ""
		// Obtain original location data.
		(*data).current = &((*data).uncompressed)
		namefirst1, namesecond1 := LocationGetName(data, i)
		coordinates1 := LocationGetCoordinates(data, i)
		appearance1 := LocationGetAppearance(data, i)
		troopID1 := LocationGetTroopID(data, i)
		status1 := LocationGetStatus(data, i)
		stage1 := LocationGetStage(data, i)
		spiceFieldID1 := LocationGetFieldID(data, i)
		spiceAmount1 := LocationGetSpiceAmount(data, i)
		spiceDensity1 := LocationGetSpiceDensity(data, i)
		fieldJ1 := LocationGetFieldJ(data, i)
		harvesters1 := LocationGetHarvesters(data, i)
		ornithopters1 := LocationGetOrnithopters(data, i)
		krysKnives1 := LocationGetKrysKnives(data, i)
		laserGuns1 := LocationGetLaserGuns(data, i)
		weirdingModules1 := LocationGetWeirdingModules(data, i)
		atomics1 := LocationGetAtomics(data, i)
		bulbs1 := LocationGetBulbs(data, i)
		water1 := LocationGetWater(data, i)

		// Obtain modified location data.
		(*data).current = &((*data).modified)
		namefirst2, namesecond2 := LocationGetName(data, i)
		coordinates2 := LocationGetCoordinates(data, i)
		appearance2 := LocationGetAppearance(data, i)
		troopID2 := LocationGetTroopID(data, i)
		status2 := LocationGetStatus(data, i)
		stage2 := LocationGetStage(data, i)
		spiceFieldID2 := LocationGetFieldID(data, i)
		spiceAmount2 := LocationGetSpiceAmount(data, i)
		spiceDensity2 := LocationGetSpiceDensity(data, i)
		fieldJ2 := LocationGetFieldJ(data, i)
		harvesters2 := LocationGetHarvesters(data, i)
		ornithopters2 := LocationGetOrnithopters(data, i)
		krysKnives2 := LocationGetKrysKnives(data, i)
		laserGuns2 := LocationGetLaserGuns(data, i)
		weirdingModules2 := LocationGetWeirdingModules(data, i)
		atomics2 := LocationGetAtomics(data, i)
		bulbs2 := LocationGetBulbs(data, i)
		water2 := LocationGetWater(data, i)

		// If there are no modifications, bail out early.
		if (namefirst1 != namefirst2) || (namesecond1 != namesecond2) || !bytes.Equal(coordinates1, coordinates2) || (appearance1 != appearance2) || (troopID1 != troopID2) || (status1 != status2) || !bytes.Equal(stage1, stage2) || (spiceFieldID1 != spiceFieldID2) || (spiceAmount1 != spiceAmount2) || (spiceDensity1 != spiceDensity2) || (fieldJ1 != fieldJ2) || (harvesters1 != harvesters2) || (ornithopters1 != ornithopters2) || (krysKnives1 != krysKnives2) || (laserGuns1 != laserGuns2) || (weirdingModules1 != weirdingModules2) || (atomics1 != atomics2) || (bulbs1 != bulbs2) || (water1 != water2) {

			changelog = LocationGetNameStr(namefirst2, namesecond2) + ":"

			if troopID1 != troopID2 {
				changelog += fmt.Sprintf(" switched housed troop ID from %d (%X) to %d (%X),", troopID1, troopID1, troopID2, troopID2)
			}

			if appearance1 != appearance2 {
				changelog += fmt.Sprintf(" changed appearance type from %d (%X) to %d (%X),", appearance1, appearance1, appearance2, appearance2)
			}

			if status1 != status2 {
				changelog += fmt.Sprintf(" changed status from %d (%X) to %d (%X),", status1, status1, status2, status2)
			}

			if spiceDensity1 < spiceDensity2 {
				changelog += fmt.Sprintf(" scaled up the spice density from %d (%X) to %d (%X),", spiceDensity1, spiceDensity1, spiceDensity2, spiceDensity2)
			} else if spiceDensity1 > spiceDensity2 {
				changelog += fmt.Sprintf(" scaled down the spice density from %d (%X) to %d (%X),", spiceDensity1, spiceDensity1, spiceDensity2, spiceDensity2)
			}

			if spiceAmount1 < spiceAmount2 {
				changelog += fmt.Sprintf(" raised the amount of spice from %d (%X) to %d (%X),", spiceAmount1, spiceAmount1, spiceAmount2, spiceAmount2)
			} else if spiceAmount1 > spiceAmount2 {
				changelog += fmt.Sprintf(" lowered the amount of spice from %d (%X) to %d (%X),", spiceAmount1, spiceAmount1, spiceAmount2, spiceAmount2)
			}

			if harvesters1 < harvesters2 {
				changelog += fmt.Sprintf(" increased the number of harvesters from %d (%X) to %d (%X),", harvesters1, harvesters1, harvesters2, harvesters2)
			} else if harvesters1 > harvesters2 {
				changelog += fmt.Sprintf(" decreased the number of harvesters from %d (%X) to %d (%X),", harvesters1, harvesters1, harvesters2, harvesters2)
			}

			if ornithopters1 < ornithopters2 {
				changelog += fmt.Sprintf(" increased the number of ornithopters from %d (%X) to %d (%X),", ornithopters1, ornithopters1, ornithopters2, ornithopters2)
			} else if ornithopters1 > ornithopters2 {
				changelog += fmt.Sprintf(" decreased the number of ornithopters from %d (%X) to %d (%X),", ornithopters1, ornithopters1, ornithopters2, ornithopters2)
			}

			if krysKnives1 < krysKnives2 {
				changelog += fmt.Sprintf(" increased the number of krys knives from %d (%X) to %d (%X),", krysKnives1, krysKnives1, krysKnives2, krysKnives2)
			} else if krysKnives1 > krysKnives2 {
				changelog += fmt.Sprintf(" decreased the number of krys knives from %d (%X) to %d (%X),", krysKnives1, krysKnives1, krysKnives2, krysKnives2)
			}

			if laserGuns1 < laserGuns2 {
				changelog += fmt.Sprintf(" increased the number of laser guns from %d (%X) to %d (%X),", laserGuns1, laserGuns1, laserGuns2, laserGuns2)
			} else if laserGuns1 > laserGuns2 {
				changelog += fmt.Sprintf(" decreased the number of laser guns from %d (%X) to %d (%X),", laserGuns1, laserGuns1, laserGuns2, laserGuns2)
			}

			if weirdingModules1 < weirdingModules2 {
				changelog += fmt.Sprintf(" increased the number of weirding modules from %d (%X) to %d (%X),", weirdingModules1, weirdingModules1, weirdingModules2, weirdingModules2)
			} else if weirdingModules1 > weirdingModules2 {
				changelog += fmt.Sprintf(" decreased the number of weirding modules from %d (%X) to %d (%X),", weirdingModules1, weirdingModules1, weirdingModules2, weirdingModules2)
			}

			if atomics1 < atomics2 {
				changelog += fmt.Sprintf(" increased the number of atomics from %d (%X) to %d (%X),", atomics1, atomics1, atomics2, atomics2)
			} else if atomics1 > atomics2 {
				changelog += fmt.Sprintf(" decreased the number of atomics from %d (%X) to %d (%X),", atomics1, atomics1, atomics2, atomics2)
			}

			if bulbs1 < bulbs2 {
				changelog += fmt.Sprintf(" increased the number of bulbs from %d (%X) to %d (%X),", bulbs1, bulbs1, bulbs2, bulbs2)
			} else if bulbs1 > bulbs2 {
				changelog += fmt.Sprintf(" decreased the number of bulbs from %d (%X) to %d (%X),", bulbs1, bulbs1, bulbs2, bulbs2)
			}

			if water1 < water2 {
				changelog += fmt.Sprintf(" expanded the amount of water from %d (%X) to %d (%X),", water1, water1, water2, water2)
			} else if water1 > water2 {
				changelog += fmt.Sprintf(" drained the amount of water from %d (%X) to %d (%X),", water1, water1, water2, water2)
			}

			if !bytes.Equal(stage1, stage2) {
				changelog += fmt.Sprintf(" changed game stage for finding it from %d (%X) to %d (%X),", stage1[0], stage1[0], stage2[0], stage2[0])
			}

			// Trim final ","
			if changelog[len(changelog)-1] == ',' {
				changelog = changelog[0 : len(changelog)-1]
				changelog += "."
			}

			specialDescription := data.specialLocationDescriptions[i]
			if specialDescription != "" {
				changelog += specialDescription
			}
		}

		return changelog, nil
	} else {
		return "", errors.New("Improper location ID")
	}
}

func LocationDiffAllProduceChangelogEntries(data *DuneMetadata) (string, error) {
	changelog := `LOCATION ADJUSTMENTS:

`
	for i := uint(LOCATION_MIN_ID); i <= LOCATION_MAX_ID; i++ {
		changelogLocation, err := LocationDiffProduceChangelogEntry(data, i)
		if err != nil {
			return "", err
		}
		if changelogLocation != "" {
			changelog += changelogLocation + "\n"
		}
	}

	if changelog[len(changelog)-1] == '\n' && changelog[len(changelog)-2] != '\n' {
		changelog += "\n"
	}
	return changelog, nil
}
