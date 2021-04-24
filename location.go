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

func LocationGetName(data *[]byte, locationsOffset uint, i uint) (byte, byte) {
	return (*data)[(locationsOffset+(i-1)*LOCATION_SIZE)+0], (*data)[(locationsOffset+(i-1)*LOCATION_SIZE)+1]
}

func LocationGetCoordinates(data *[]byte, locationsOffset uint, i uint) []byte {
	return (*data)[(locationsOffset+(i-1)*LOCATION_SIZE)+2 : (locationsOffset+(i-1)*LOCATION_SIZE)+8]
}

func LocationGetAppearance(data *[]byte, locationsOffset uint, i uint) byte {
	return (*data)[(locationsOffset+(i-1)*LOCATION_SIZE)+8]
}

func LocationSetAppearance(data *[]byte, locationsOffset uint, i uint, value byte) {
	(*data)[(locationsOffset+(i-1)*LOCATION_SIZE)+8] = value
}

func LocationGetTroopID(data *[]byte, locationsOffset uint, i uint) byte {
	return (*data)[(locationsOffset+(i-1)*LOCATION_SIZE)+9]
}

func LocationSetTroopID(data *[]byte, locationsOffset uint, i uint, value byte) {
	(*data)[(locationsOffset+(i-1)*LOCATION_SIZE)+9] = value
}

func LocationGetStatus(data *[]byte, locationsOffset uint, i uint) byte {
	return (*data)[(locationsOffset+(i-1)*LOCATION_SIZE)+10]
}

func LocationSetStatus(data *[]byte, locationsOffset uint, i uint, value byte) {
	(*data)[(locationsOffset+(i-1)*LOCATION_SIZE)+10] = value
}

func LocationGetStage(data *[]byte, locationsOffset uint, i uint) []byte {
	return (*data)[(locationsOffset+(i-1)*LOCATION_SIZE)+11 : (locationsOffset+(i-1)*LOCATION_SIZE)+16]
}

func LocationGetFieldID(data *[]byte, locationsOffset uint, i uint) byte {
	return (*data)[(locationsOffset+(i-1)*LOCATION_SIZE)+16]
}

func LocationGetSpiceAmount(data *[]byte, locationsOffset uint, i uint) byte {
	return (*data)[(locationsOffset+(i-1)*LOCATION_SIZE)+17]
}

func LocationSetSpiceAmount(data *[]byte, locationsOffset uint, i uint, value byte) {
	(*data)[(locationsOffset+(i-1)*LOCATION_SIZE)+17] = value
}

func LocationGetSpiceDensity(data *[]byte, locationsOffset uint, i uint) byte {
	return (*data)[(locationsOffset+(i-1)*LOCATION_SIZE)+18]
}

func LocationSetSpiceDensity(data *[]byte, locationsOffset uint, i uint, value byte) {
	(*data)[(locationsOffset+(i-1)*LOCATION_SIZE)+18] = value
}

func LocationGetFieldJ(data *[]byte, locationsOffset uint, i uint) byte {
	return (*data)[(locationsOffset+(i-1)*LOCATION_SIZE)+19]
}

func LocationGetHarvesters(data *[]byte, locationsOffset uint, i uint) byte {
	return (*data)[(locationsOffset+(i-1)*LOCATION_SIZE)+20]
}

func LocationSetHarvesters(data *[]byte, locationsOffset uint, i uint, value byte) {
	(*data)[(locationsOffset+(i-1)*LOCATION_SIZE)+20] = value
}

func LocationGetOrnithopters(data *[]byte, locationsOffset uint, i uint) byte {
	return (*data)[(locationsOffset+(i-1)*LOCATION_SIZE)+21]
}

func LocationSetOrnithopters(data *[]byte, locationsOffset uint, i uint, value byte) {
	(*data)[(locationsOffset+(i-1)*LOCATION_SIZE)+21] = value
}

func LocationGetKrysKnives(data *[]byte, locationsOffset uint, i uint) byte {
	return (*data)[(locationsOffset+(i-1)*LOCATION_SIZE)+22]
}

func LocationSetKrysKnives(data *[]byte, locationsOffset uint, i uint, value byte) {
	(*data)[(locationsOffset+(i-1)*LOCATION_SIZE)+22] = value
}

func LocationGetLaserPistols(data *[]byte, locationsOffset uint, i uint) byte {
	return (*data)[(locationsOffset+(i-1)*LOCATION_SIZE)+23]
}

func LocationSetLaserPistols(data *[]byte, locationsOffset uint, i uint, value byte) {
	(*data)[(locationsOffset+(i-1)*LOCATION_SIZE)+23] = value
}

func LocationGetWeirdingModules(data *[]byte, locationsOffset uint, i uint) byte {
	return (*data)[(locationsOffset+(i-1)*LOCATION_SIZE)+24]
}

func LocationSetWeirdingModules(data *[]byte, locationsOffset uint, i uint, value byte) {
	(*data)[(locationsOffset+(i-1)*LOCATION_SIZE)+24] = value
}

func LocationGetAtomics(data *[]byte, locationsOffset uint, i uint) byte {
	return (*data)[(locationsOffset+(i-1)*LOCATION_SIZE)+25]
}

func LocationSetAtomics(data *[]byte, locationsOffset uint, i uint, value byte) {
	(*data)[(locationsOffset+(i-1)*LOCATION_SIZE)+25] = value
}

func LocationGetBulbs(data *[]byte, locationsOffset uint, i uint) byte {
	return (*data)[(locationsOffset+(i-1)*LOCATION_SIZE)+26]
}

func LocationGetWater(data *[]byte, locationsOffset uint, i uint) byte {
	return (*data)[(locationsOffset+(i-1)*LOCATION_SIZE)+27]
}

func LocationSetWater(data *[]byte, locationsOffset uint, i uint, value byte) {
	(*data)[(locationsOffset+(i-1)*LOCATION_SIZE)+27] = value
}

func LocationPrint(data *[]byte, mode uint, i uint) {
	// Number locations from 1
	if i > 0 && i <= LOCATION_MAX_ID {
		_, _, locationsOffset, _, _, _ := GetOffsets(mode)
		namefirst, namesecond := LocationGetName(data, locationsOffset, i)
		namestr := LocationGetNameStr(namefirst, namesecond)
		coordinates := LocationGetCoordinates(data, locationsOffset, i)
		appearance := LocationGetAppearance(data, locationsOffset, i)
		troopID := LocationGetTroopID(data, locationsOffset, i)
		status := LocationGetStatus(data, locationsOffset, i)
		stage := LocationGetStage(data, locationsOffset, i)
		spiceFieldID := LocationGetFieldID(data, locationsOffset, i)
		spiceAmount := LocationGetSpiceAmount(data, locationsOffset, i)
		spiceDensity := LocationGetSpiceDensity(data, locationsOffset, i)
		fieldJ := LocationGetFieldJ(data, locationsOffset, i)
		harvesters := LocationGetHarvesters(data, locationsOffset, i)
		ornithopters := LocationGetOrnithopters(data, locationsOffset, i)
		krysKnives := LocationGetKrysKnives(data, locationsOffset, i)
		laserPistols := LocationGetLaserPistols(data, locationsOffset, i)
		weirdingModules := LocationGetWeirdingModules(data, locationsOffset, i)
		atomics := LocationGetAtomics(data, locationsOffset, i)
		bulbs := LocationGetBulbs(data, locationsOffset, i)
		water := LocationGetWater(data, locationsOffset, i)
		fmt.Printf("Location %02x Name: %s Coordinates: %x\n", i, namestr, coordinates)
		fmt.Printf("\tAppearance: %02x Troop ID: %02x Status: %02x Stage: %x\n\tSpice field ID: %02x Spice amount: %02x Spice density: %02x J: %02x\n", appearance, troopID, status, stage, spiceFieldID, spiceAmount, spiceDensity, fieldJ)
		fmt.Printf("\tHarvesters: %02x Ornithopters: %02x Krys knives: %02x Laser pistols: %02x Weirding modules: %02x Atomics: %02x Bulbs: %02x Water: %02x\n", harvesters, ornithopters, krysKnives, laserPistols, weirdingModules, atomics, bulbs, water)
	}
}
