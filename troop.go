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

func TroopGetTroopID(data *DuneMetadata, i uint) byte {
	return ((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+0]
}

func TroopSetTroopID(data *DuneMetadata, i uint, value byte) {
	((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+0] = value
}

func TroopGetNextTroopID(data *DuneMetadata, i uint) byte {
	return ((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+1]
}

func TroopSetNextTroopID(data *DuneMetadata, i uint, value byte) {
	((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+1] = value
}

func TroopGetPosition(data *DuneMetadata, i uint) byte {
	return ((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+2]
}

func TroopSetPosition(data *DuneMetadata, i uint, value byte) {
	((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+2] = value
}

func TroopGetOccupation(data *DuneMetadata, i uint) byte {
	return ((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+3]
}

func TroopSetOccupation(data *DuneMetadata, i uint, value byte) {
	((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+3] = value
}

func TroopGetFieldE(data *DuneMetadata, i uint) uint16 {
	return uint16(((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+4]) | uint16(((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+5])*256
}

func TroopSetFieldE(data *DuneMetadata, i uint, e uint16) {
	((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+4] = byte(e)
	((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+5] = byte(e >> 8)
}

func TroopGetCoordinates(data *DuneMetadata, i uint) []byte {
	return ((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+6 : (data.troopsOffset+(i-1)*TROOP_SIZE)+10]
}

func TroopSetCoordinates(data *DuneMetadata, i uint, coordinates []byte) {
	copy(((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+6:(data.troopsOffset+(i-1)*TROOP_SIZE)+10], coordinates[0:4])
}

func TroopGetFieldG(data *DuneMetadata, i uint) []byte {
	return ((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+10 : (data.troopsOffset+(i-1)*TROOP_SIZE)+18]
}

func TroopSetFieldG(data *DuneMetadata, i uint, g []byte) {
	copy(((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+10:(data.troopsOffset+(i-1)*TROOP_SIZE)+18], g[0:8])
}

func TroopGetDissatisfaction(data *DuneMetadata, i uint) byte {
	return ((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+18]
}

func TroopSetDissatisfaction(data *DuneMetadata, i uint, value byte) {
	((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+18] = value
}

func TroopGetSpeech(data *DuneMetadata, i uint) byte {
	return ((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+19]
}

func TroopSetSpeech(data *DuneMetadata, i uint, value byte) {
	((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+19] = value
}

func TroopGetFieldJ(data *DuneMetadata, i uint) byte {
	return ((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+20]
}

func TroopSetFieldJ(data *DuneMetadata, i uint, value byte) {
	((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+20] = value
}

func TroopGetMotivation(data *DuneMetadata, i uint) byte {
	return ((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+21]
}

func TroopSetMotivation(data *DuneMetadata, i uint, value byte) {
	((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+21] = value
}

func TroopGetSpiceMiningSkill(data *DuneMetadata, i uint) byte {
	return ((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+22]
}

func TroopSetSpiceMiningSkill(data *DuneMetadata, i uint, value byte) {
	((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+22] = value
}

func TroopGetArmySkill(data *DuneMetadata, i uint) byte {
	return ((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+23]
}

func TroopSetArmySkill(data *DuneMetadata, i uint, value byte) {
	((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+23] = value
}

func TroopGetEcologySkill(data *DuneMetadata, i uint) byte {
	return ((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+24]
}

func TroopSetEcologySkill(data *DuneMetadata, i uint, value byte) {
	((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+24] = value
}

func TroopGetEquipment(data *DuneMetadata, i uint) byte {
	return ((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+25]
}

func TroopSetEquipment(data *DuneMetadata, i uint, value byte) {
	((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+25] = value
}

func TroopGetPopulation(data *DuneMetadata, i uint) uint16 {
	return uint16(((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+26]) * 10
}

func TroopSetPopulation(data *DuneMetadata, i uint, value uint16) {
	((*data).modified)[(data.troopsOffset+(i-1)*TROOP_SIZE)+26] = byte((value / 10) & 0xFF)
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

func TroopSwapPair(data *DuneMetadata, i1 uint, i2 uint) {
	tmp := make([]byte, TROOP_SIZE)
	copy(tmp[:], ((*data).modified)[(data.troopsOffset+(i1-1)*TROOP_SIZE):data.troopsOffset+i1*TROOP_SIZE])
	copy(((*data).modified)[(data.troopsOffset+(i1-1)*TROOP_SIZE):data.troopsOffset+i1*TROOP_SIZE], ((*data).modified)[(data.troopsOffset+(i2-1)*TROOP_SIZE):data.troopsOffset+i2*TROOP_SIZE])
	copy(((*data).modified)[(data.troopsOffset+(i2-1)*TROOP_SIZE):data.troopsOffset+i2*TROOP_SIZE], tmp[:])
}

// Move troop i1 to i4, i2 to i5, i3 to i6
func TroopShuffleTrio(data *DuneMetadata, i1 uint, i2 uint, i3 uint, i4 uint, i5 uint, i6 uint) {
	tmp1 := make([]byte, TROOP_SIZE)
	tmp2 := make([]byte, TROOP_SIZE)
	tmp3 := make([]byte, TROOP_SIZE)
	copy(tmp1[:], ((*data).modified)[(data.troopsOffset+(i1-1)*TROOP_SIZE):data.troopsOffset+i1*TROOP_SIZE])
	copy(tmp2[:], ((*data).modified)[(data.troopsOffset+(i2-1)*TROOP_SIZE):data.troopsOffset+i2*TROOP_SIZE])
	copy(tmp3[:], ((*data).modified)[(data.troopsOffset+(i3-1)*TROOP_SIZE):data.troopsOffset+i3*TROOP_SIZE])
	copy(((*data).modified)[(data.troopsOffset+(i4-1)*TROOP_SIZE):data.troopsOffset+i4*TROOP_SIZE], tmp1[:])
	copy(((*data).modified)[(data.troopsOffset+(i5-1)*TROOP_SIZE):data.troopsOffset+i5*TROOP_SIZE], tmp2[:])
	copy(((*data).modified)[(data.troopsOffset+(i6-1)*TROOP_SIZE):data.troopsOffset+i6*TROOP_SIZE], tmp3[:])
}

func TroopPrint(data *DuneMetadata, i uint) {
	// Number troops from 1
	if i > 0 && i <= TROOP_MAX_ID {
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
