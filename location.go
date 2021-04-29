// Copyright (C) 2021 Lionel Debroux
// SPDX-License-Identifier: GPL-2.0
//
// Sources of information:
// * https://forum.dune2k.com/topic/20497-dune-cheats/
// * hugslab
// * Dmitri Fatkin

package main

import (
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
		str = str + " (Atreides)"
	case 0x02:
		str = str + " (Harkonnen)"
	case 0x03:
		str = str + "-Tabr"
	case 0x04:
		str = str + "-Timin"
	case 0x05:
		str = str + "-Tuek"
	case 0x06:
		str = str + "-Harg"
	case 0x07:
		str = str + "-Clam"
	case 0x08:
		str = str + "-Tsymyn"
	case 0x09:
		str = str + "-Siet"
	case 0x0a:
		str = str + "-Pyons"
	case 0x0b:
		str = str + "-Pyort"
	case 0x0c:
		str = str + "-Celimyn"
	}

	return str
}

func LocationGetName(data *DuneMetadata, i uint) (byte, byte) {
	return ((*data).modified)[(data.locationsOffset+(i-1)*LOCATION_SIZE)+0], ((*data).modified)[(data.locationsOffset+(i-1)*LOCATION_SIZE)+1]
}

func LocationGetCoordinates(data *DuneMetadata, i uint) []byte {
	return ((*data).modified)[(data.locationsOffset+(i-1)*LOCATION_SIZE)+2 : (data.locationsOffset+(i-1)*LOCATION_SIZE)+8]
}

func LocationGetAppearance(data *DuneMetadata, i uint) byte {
	return ((*data).modified)[(data.locationsOffset+(i-1)*LOCATION_SIZE)+8]
}

func LocationSetAppearance(data *DuneMetadata, i uint, value byte) {
	((*data).modified)[(data.locationsOffset+(i-1)*LOCATION_SIZE)+8] = value
}

func LocationGetTroopID(data *DuneMetadata, i uint) byte {
	return ((*data).modified)[(data.locationsOffset+(i-1)*LOCATION_SIZE)+9]
}

func LocationSetTroopID(data *DuneMetadata, i uint, value byte) {
	((*data).modified)[(data.locationsOffset+(i-1)*LOCATION_SIZE)+9] = value
}

func LocationGetStatus(data *DuneMetadata, i uint) byte {
	return ((*data).modified)[(data.locationsOffset+(i-1)*LOCATION_SIZE)+10]
}

func LocationSetStatus(data *DuneMetadata, i uint, value byte) {
	((*data).modified)[(data.locationsOffset+(i-1)*LOCATION_SIZE)+10] = value
}

func LocationGetStage(data *DuneMetadata, i uint) []byte {
	return ((*data).modified)[(data.locationsOffset+(i-1)*LOCATION_SIZE)+11 : (data.locationsOffset+(i-1)*LOCATION_SIZE)+16]
}

func LocationSetStage(data *DuneMetadata, i uint, stage []byte) {
	copy(((*data).modified)[(data.locationsOffset+(i-1)*LOCATION_SIZE)+11:(data.locationsOffset+(i-1)*LOCATION_SIZE)+16], stage[0:5])
}

func LocationGetFieldID(data *DuneMetadata, i uint) byte {
	return ((*data).modified)[(data.locationsOffset+(i-1)*LOCATION_SIZE)+16]
}

func LocationGetSpiceAmount(data *DuneMetadata, i uint) byte {
	return ((*data).modified)[(data.locationsOffset+(i-1)*LOCATION_SIZE)+17]
}

func LocationSetSpiceAmount(data *DuneMetadata, i uint, value byte) {
	((*data).modified)[(data.locationsOffset+(i-1)*LOCATION_SIZE)+17] = value
}

func LocationGetSpiceDensity(data *DuneMetadata, i uint) byte {
	return ((*data).modified)[(data.locationsOffset+(i-1)*LOCATION_SIZE)+18]
}

func LocationSetSpiceDensity(data *DuneMetadata, i uint, value byte) {
	((*data).modified)[(data.locationsOffset+(i-1)*LOCATION_SIZE)+18] = value
}

func LocationGetFieldJ(data *DuneMetadata, i uint) byte {
	return ((*data).modified)[(data.locationsOffset+(i-1)*LOCATION_SIZE)+19]
}

func LocationGetHarvesters(data *DuneMetadata, i uint) byte {
	return ((*data).modified)[(data.locationsOffset+(i-1)*LOCATION_SIZE)+20]
}

func LocationSetHarvesters(data *DuneMetadata, i uint, value byte) {
	((*data).modified)[(data.locationsOffset+(i-1)*LOCATION_SIZE)+20] = value
}

func LocationGetOrnithopters(data *DuneMetadata, i uint) byte {
	return ((*data).modified)[(data.locationsOffset+(i-1)*LOCATION_SIZE)+21]
}

func LocationSetOrnithopters(data *DuneMetadata, i uint, value byte) {
	((*data).modified)[(data.locationsOffset+(i-1)*LOCATION_SIZE)+21] = value
}

func LocationGetKrysKnives(data *DuneMetadata, i uint) byte {
	return ((*data).modified)[(data.locationsOffset+(i-1)*LOCATION_SIZE)+22]
}

func LocationSetKrysKnives(data *DuneMetadata, i uint, value byte) {
	((*data).modified)[(data.locationsOffset+(i-1)*LOCATION_SIZE)+22] = value
}

func LocationGetLaserGuns(data *DuneMetadata, i uint) byte {
	return ((*data).modified)[(data.locationsOffset+(i-1)*LOCATION_SIZE)+23]
}

func LocationSetLaserGuns(data *DuneMetadata, i uint, value byte) {
	((*data).modified)[(data.locationsOffset+(i-1)*LOCATION_SIZE)+23] = value
}

func LocationGetWeirdingModules(data *DuneMetadata, i uint) byte {
	return ((*data).modified)[(data.locationsOffset+(i-1)*LOCATION_SIZE)+24]
}

func LocationSetWeirdingModules(data *DuneMetadata, i uint, value byte) {
	((*data).modified)[(data.locationsOffset+(i-1)*LOCATION_SIZE)+24] = value
}

func LocationGetAtomics(data *DuneMetadata, i uint) byte {
	return ((*data).modified)[(data.locationsOffset+(i-1)*LOCATION_SIZE)+25]
}

func LocationSetAtomics(data *DuneMetadata, i uint, value byte) {
	((*data).modified)[(data.locationsOffset+(i-1)*LOCATION_SIZE)+25] = value
}

func LocationGetBulbs(data *DuneMetadata, i uint) byte {
	return ((*data).modified)[(data.locationsOffset+(i-1)*LOCATION_SIZE)+26]
}

func LocationGetWater(data *DuneMetadata, i uint) byte {
	return ((*data).modified)[(data.locationsOffset+(i-1)*LOCATION_SIZE)+27]
}

func LocationSetWater(data *DuneMetadata, i uint, value byte) {
	((*data).modified)[(data.locationsOffset+(i-1)*LOCATION_SIZE)+27] = value
}

func LocationSetWeapons(data *DuneMetadata, i uint, krysKnives byte, laserGuns byte, weirdingModules byte, atomics byte) {
	LocationSetKrysKnives(data, i, krysKnives)
	LocationSetLaserGuns(data, i, laserGuns)
	LocationSetWeirdingModules(data, i, weirdingModules)
	LocationSetAtomics(data, i, atomics)
}

func LocationPrint(data *DuneMetadata, i uint) {
	// Number locations from 1
	if i > 0 && i <= LOCATION_MAX_ID {
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
		fmt.Printf("Location %02x Name: %s Coordinates: %x\n", i, namestr, coordinates)
		fmt.Printf("\tAppearance: %02x Troop ID: %02x Status: %02x Stage: %x\n\tSpice field ID: %02x Spice amount: %02x Spice density: %02x J: %02x\n", appearance, troopID, status, stage, spiceFieldID, spiceAmount, spiceDensity, fieldJ)
		fmt.Printf("\tHarvesters: %02x Ornithopters: %02x Krys knives: %02x Laser guns: %02x Weirding modules: %02x Atomics: %02x Bulbs: %02x Water: %02x\n", harvesters, ornithopters, krysKnives, laserGuns, weirdingModules, atomics, bulbs, water)
	}
}
