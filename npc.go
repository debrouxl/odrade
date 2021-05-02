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

func NPCGetSpriteIdentificator(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.npcsOffset+(i-1)*(NPC_SIZE+NPC_PADDING))+0]
}

func NPCSetSpriteIdentificator(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.npcsOffset+(i-1)*(NPC_SIZE+NPC_PADDING))+0] = value
}

func NPCGetFieldB(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.npcsOffset+(i-1)*(NPC_SIZE+NPC_PADDING))+1]
}

func NPCSetFieldB(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.npcsOffset+(i-1)*(NPC_SIZE+NPC_PADDING))+1] = value
}

func NPCGetRoomLocation(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.npcsOffset+(i-1)*(NPC_SIZE+NPC_PADDING))+2]
}

func NPCSetRoomLocation(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.npcsOffset+(i-1)*(NPC_SIZE+NPC_PADDING))+2] = value
}

func NPCGetTypeOfPlace(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.npcsOffset+(i-1)*(NPC_SIZE+NPC_PADDING))+3]
}

func NPCSetTypeOfPlace(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.npcsOffset+(i-1)*(NPC_SIZE+NPC_PADDING))+3] = value
}

func NPCGetFieldE(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.npcsOffset+(i-1)*(NPC_SIZE+NPC_PADDING))+4]
}

func NPCSetFieldE(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.npcsOffset+(i-1)*(NPC_SIZE+NPC_PADDING))+4] = value
}

func NPCGetExactPlace(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.npcsOffset+(i-1)*(NPC_SIZE+NPC_PADDING))+5]
}

func NPCSetExactPlace(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.npcsOffset+(i-1)*(NPC_SIZE+NPC_PADDING))+5] = value
}

func NPCGetForDialogue(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.npcsOffset+(i-1)*(NPC_SIZE+NPC_PADDING))+6]
}

func NPCSetForDialogue(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.npcsOffset+(i-1)*(NPC_SIZE+NPC_PADDING))+6] = value
}

func NPCGetFieldH(data *DuneMetadata, i uint) byte {
	return (*((*data).current))[(data.npcsOffset+(i-1)*(NPC_SIZE+NPC_PADDING))+7]
}

func NPCSetFieldH(data *DuneMetadata, i uint, value byte) {
	(*((*data).current))[(data.npcsOffset+(i-1)*(NPC_SIZE+NPC_PADDING))+7] = value
}

func NPCPrint(data *DuneMetadata, i uint) {
	if i >= NPC_MIN_ID && i <= NPC_MAX_ID {
		spriteIdentificator := NPCGetSpriteIdentificator(data, i)
		fieldB := NPCGetFieldB(data, i)
		roomLocation := NPCGetRoomLocation(data, i)
		typeOfPlace := NPCGetTypeOfPlace(data, i)
		fieldE := NPCGetFieldE(data, i)
		exactPlace := NPCGetExactPlace(data, i)
		forDialogue := NPCGetForDialogue(data, i)
		fieldH := NPCGetFieldH(data, i)
		fmt.Printf("NPC %02x Sprite identificator: %x B: %x\n", i, spriteIdentificator, fieldB)
		fmt.Printf("\tRoom location: %02x Type of place: %02x E: %02x Exact place: %02x For dialogue: %02x H: %02x\n", roomLocation, typeOfPlace, fieldE, exactPlace, forDialogue, fieldH)
	}
}

func NPCDiffProduceChangelogEntry(data *DuneMetadata, i uint) (string, error) {
	if i >= NPC_MIN_ID && i <= NPC_MAX_ID {
		changelog := ""

		return changelog, nil
	} else {
		return "", errors.New("Improper NPC ID")
	}
}

func NPCDiffAllProduceChangelogEntries(data *DuneMetadata) (string, error) {
	changelog := "TODO"
	for i := uint(NPC_MIN_ID); i <= NPC_MAX_ID; i++ {
		changelogNPC, err := NPCDiffProduceChangelogEntry(data, i)
		if err != nil {
			return "", err
		}
		changelog = changelog + changelogNPC
	}

	return changelog, nil
}
