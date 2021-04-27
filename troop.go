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

func TroopGetTroopID(data *[]byte, troopsOffset uint, i uint) byte {
	return (*data)[(troopsOffset+(i-1)*TROOP_SIZE)+0]
}

func TroopSetTroopID(data *[]byte, troopsOffset uint, i uint, value byte) {
	(*data)[(troopsOffset+(i-1)*TROOP_SIZE)+0] = value
}


func TroopGetNextTroopID(data *[]byte, troopsOffset uint, i uint) byte {
	return (*data)[(troopsOffset+(i-1)*TROOP_SIZE)+1]
}

func TroopSetNextTroopID(data *[]byte, troopsOffset uint, i uint, value byte) {
	(*data)[(troopsOffset+(i-1)*TROOP_SIZE)+1] = value
}


func TroopGetPosition(data *[]byte, troopsOffset uint, i uint) byte {
	return (*data)[(troopsOffset+(i-1)*TROOP_SIZE)+2]
}

func TroopSetPosition(data *[]byte, troopsOffset uint, i uint, value byte) {
	(*data)[(troopsOffset+(i-1)*TROOP_SIZE)+2] = value
}


func TroopGetOccupation(data *[]byte, troopsOffset uint, i uint) byte {
	return (*data)[(troopsOffset+(i-1)*TROOP_SIZE)+3]
}

func TroopSetOccupation(data *[]byte, troopsOffset uint, i uint, value byte) {
	(*data)[(troopsOffset+(i-1)*TROOP_SIZE)+3] = value
}


func TroopGetFieldE(data *[]byte, troopsOffset uint, i uint) uint16 {
	return uint16((*data)[(troopsOffset+(i-1)*TROOP_SIZE)+4]) | uint16((*data)[(troopsOffset+(i-1)*TROOP_SIZE)+5])*256
}

func TroopSetFieldE(data *[]byte, troopsOffset uint, i uint, e uint16) {
	(*data)[(troopsOffset+(i-1)*TROOP_SIZE)+4] = byte(e)
	(*data)[(troopsOffset+(i-1)*TROOP_SIZE)+5] = byte(e >> 8)
}


func TroopGetCoordinates(data *[]byte, troopsOffset uint, i uint) []byte {
	return (*data)[(troopsOffset+(i-1)*TROOP_SIZE)+6 : (troopsOffset+(i-1)*TROOP_SIZE)+10]
}

func TroopSetCoordinates(data *[]byte, troopsOffset uint, i uint, coordinates []byte) {
	copy((*data)[(troopsOffset+(i-1)*TROOP_SIZE)+6:(troopsOffset+(i-1)*TROOP_SIZE)+10], coordinates[0:4])
}


func TroopGetFieldG(data *[]byte, troopsOffset uint, i uint) []byte {
	return (*data)[(troopsOffset+(i-1)*TROOP_SIZE)+10 : (troopsOffset+(i-1)*TROOP_SIZE)+18]
}

func TroopSetFieldG(data *[]byte, troopsOffset uint, i uint, g []byte) {
	copy((*data)[(troopsOffset+(i-1)*TROOP_SIZE)+10:(troopsOffset+(i-1)*TROOP_SIZE)+18], g[0:8])
}


func TroopGetDissatisfaction(data *[]byte, troopsOffset uint, i uint) byte {
	return (*data)[(troopsOffset+(i-1)*TROOP_SIZE)+18]
}

func TroopSetDissatisfaction(data *[]byte, troopsOffset uint, i uint, value byte) {
	(*data)[(troopsOffset+(i-1)*TROOP_SIZE)+18] = value
}


func TroopGetSpeech(data *[]byte, troopsOffset uint, i uint) byte {
	return (*data)[(troopsOffset+(i-1)*TROOP_SIZE)+19]
}

func TroopSetSpeech(data *[]byte, troopsOffset uint, i uint, value byte) {
	(*data)[(troopsOffset+(i-1)*TROOP_SIZE)+19] = value
}


func TroopGetFieldJ(data *[]byte, troopsOffset uint, i uint) byte {
	return (*data)[(troopsOffset+(i-1)*TROOP_SIZE)+20]
}

func TroopSetFieldJ(data *[]byte, troopsOffset uint, i uint, value byte) {
	(*data)[(troopsOffset+(i-1)*TROOP_SIZE)+20] = value
}


func TroopGetMotivation(data *[]byte, troopsOffset uint, i uint) byte {
	return (*data)[(troopsOffset+(i-1)*TROOP_SIZE)+21]
}

func TroopSetMotivation(data *[]byte, troopsOffset uint, i uint, value byte) {
	(*data)[(troopsOffset+(i-1)*TROOP_SIZE)+21] = value
}


func TroopGetSpiceMiningSkill(data *[]byte, troopsOffset uint, i uint) byte {
	return (*data)[(troopsOffset+(i-1)*TROOP_SIZE)+22]
}

func TroopSetSpiceMiningSkill(data *[]byte, troopsOffset uint, i uint, value byte) {
	(*data)[(troopsOffset+(i-1)*TROOP_SIZE)+22] = value
}


func TroopGetArmySkill(data *[]byte, troopsOffset uint, i uint) byte {
	return (*data)[(troopsOffset+(i-1)*TROOP_SIZE)+23]
}

func TroopSetArmySkill(data *[]byte, troopsOffset uint, i uint, value byte) {
	(*data)[(troopsOffset+(i-1)*TROOP_SIZE)+23] = value
}


func TroopGetEcologySkill(data *[]byte, troopsOffset uint, i uint) byte {
	return (*data)[(troopsOffset+(i-1)*TROOP_SIZE)+24]
}

func TroopSetEcologySkill(data *[]byte, troopsOffset uint, i uint, value byte) {
	(*data)[(troopsOffset+(i-1)*TROOP_SIZE)+24] = value
}


func TroopGetEquipment(data *[]byte, troopsOffset uint, i uint) byte {
	return (*data)[(troopsOffset+(i-1)*TROOP_SIZE)+25]
}

func TroopSetEquipment(data *[]byte, troopsOffset uint, i uint, value byte) {
	(*data)[(troopsOffset+(i-1)*TROOP_SIZE)+25] = value
}


func TroopGetPopulation(data *[]byte, troopsOffset uint, i uint) uint16 {
	return uint16((*data)[(troopsOffset+(i-1)*TROOP_SIZE)+26]) * 10
}

func TroopSetPopulation(data *[]byte, troopsOffset uint, i uint, value uint16) {
	(*data)[(troopsOffset+(i-1)*TROOP_SIZE)+26] = byte((value / 10) & 0xFF)
}

func TroopSetPopulationMotivationArmySkill(data *[]byte, troopsOffset uint, i uint, population uint16, motivation byte, armySkill byte) {
	TroopSetPopulation(data, troopsOffset, i, population)
	TroopSetMotivation(data, troopsOffset, i, motivation)
	TroopSetArmySkill(data, troopsOffset, i, armySkill)
}

func TroopSetPopulationMotivationArmySkillEquipment(data *[]byte, troopsOffset uint, i uint, population uint16, motivation byte, armySkill byte, equipment byte) {
	TroopSetPopulationMotivationArmySkill(data, troopsOffset, i, population, motivation, armySkill)
	TroopSetEquipment(data, troopsOffset, i, equipment)
}

func TroopSwapPair(data *[]byte, troopsOffset uint, i1 uint, i2 uint) {
	tmp := make([]byte, TROOP_SIZE)
	copy(tmp[:], (*data)[(troopsOffset+(i1-1)*TROOP_SIZE):troopsOffset+i1*TROOP_SIZE])
	copy((*data)[(troopsOffset+(i1-1)*TROOP_SIZE):troopsOffset+i1*TROOP_SIZE], (*data)[(troopsOffset+(i2-1)*TROOP_SIZE):troopsOffset+i2*TROOP_SIZE])
	copy((*data)[(troopsOffset+(i2-1)*TROOP_SIZE):troopsOffset+i2*TROOP_SIZE], tmp[:])
}

// Move troop i1 to i4, i2 to i5, i3 to i6
func TroopShuffleTrio(data *[]byte, troopsOffset uint, i1 uint, i2 uint, i3 uint, i4 uint, i5 uint, i6 uint) {
	tmp1 := make([]byte, TROOP_SIZE)
	tmp2 := make([]byte, TROOP_SIZE)
	tmp3 := make([]byte, TROOP_SIZE)
	copy(tmp1[:], (*data)[(troopsOffset+(i1-1)*TROOP_SIZE):troopsOffset+i1*TROOP_SIZE])
	copy(tmp2[:], (*data)[(troopsOffset+(i2-1)*TROOP_SIZE):troopsOffset+i2*TROOP_SIZE])
	copy(tmp3[:], (*data)[(troopsOffset+(i3-1)*TROOP_SIZE):troopsOffset+i3*TROOP_SIZE])
	copy((*data)[(troopsOffset+(i4-1)*TROOP_SIZE):troopsOffset+i4*TROOP_SIZE], tmp1[:])
	copy((*data)[(troopsOffset+(i5-1)*TROOP_SIZE):troopsOffset+i5*TROOP_SIZE], tmp2[:])
	copy((*data)[(troopsOffset+(i6-1)*TROOP_SIZE):troopsOffset+i6*TROOP_SIZE], tmp3[:])
}

func TroopPrint(data *[]byte, mode uint, i uint) {
	// Number troops from 1
	if i > 0 && i <= TROOP_MAX_ID {
		_, _, _, troopsOffset, _, _ := GetOffsets(mode)
		troopID := TroopGetTroopID(data, troopsOffset, i)
		nextTroopID := TroopGetNextTroopID(data, troopsOffset, i)
		position := TroopGetPosition(data, troopsOffset, i)
		occupation := TroopGetOccupation(data, troopsOffset, i)
		fieldE := TroopGetFieldE(data, troopsOffset, i)
		coordinates := TroopGetCoordinates(data, troopsOffset, i)
		fieldG := TroopGetFieldG(data, troopsOffset, i)
		dissatisfaction := TroopGetDissatisfaction(data, troopsOffset, i)
		speech := TroopGetSpeech(data, troopsOffset, i)
		fieldJ := TroopGetFieldJ(data, troopsOffset, i)
		motivation := TroopGetMotivation(data, troopsOffset, i)
		spiceSkill := TroopGetSpiceMiningSkill(data, troopsOffset, i)
		armySkill := TroopGetArmySkill(data, troopsOffset, i)
		ecologySkill := TroopGetEcologySkill(data, troopsOffset, i)
		equipment := TroopGetEquipment(data, troopsOffset, i)
		population := TroopGetPopulation(data, troopsOffset, i)
		fmt.Printf("Troop %02x Troop ID: %02x Next Troop ID: %02x Position: %02x Occupation: %02x\n", i, troopID, nextTroopID, position, occupation)
		fmt.Printf("\tE: %04x Coordinates: %x G: %x Dissatisfaction: %02x Speech: %02x J: %02x\n", fieldE, coordinates, fieldG, dissatisfaction, speech, fieldJ)
		fmt.Printf("\tMotivation: %02x Spice skill: %02x Army skill: %02x Ecology skill: %02x Equipment: %02x Population: %02d\n", motivation, spiceSkill, armySkill, ecologySkill, equipment, population)
	}
}
