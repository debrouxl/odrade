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

func NPCGetNameStr(i uint) string {
	var str string

	switch i {
	case 0x01:
		str = "Duke Leto Atreides"
	case 0x02:
		str = "Jessica Atreides"
	case 0x03:
		str = "Thufir Hawat"
	case 0x04:
		str = "Duncan Idaho"
	case 0x05:
		str = "Gurney Halleck"
	case 0x06:
		str = "Stilgar"
	case 0x07:
		str = "Liet Kynes"
	case 0x08:
		str = "Chani"
	case 0x09:
		str = "Harah"
	case 0x0a:
		str = "Baron Vladimir Harkonnen"
	case 0x0b:
		str = "Feyd-Rautha Harkonnen"
	case 0x0c:
		str = "Emperor Shaddam IV"
	case 0x0d:
		str = "Harkonnen Captains"
	case 0x0e:
		str = "Smugglers"
	case 0x0f:
		str = "The Fremen"
	case 0x10:
		str = "The Fremen"
	}

	return str
}

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

func NPCSetSpecialDescription(data *DuneMetadata, i uint, description string) {
	((*data).specialNPCDescriptions)[i] = description
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
		fmt.Printf("NPC %d (%02x) Sprite identificator: %x B: %x\n", i, i, spriteIdentificator, fieldB)
		fmt.Printf("\tRoom location: %02x Type of place: %02x E: %02x Exact place: %02x For dialogue: %02x H: %02x\n", roomLocation, typeOfPlace, fieldE, exactPlace, forDialogue, fieldH)
	}
}

func NPCDiffProduceChangelogEntry(data *DuneMetadata, i uint) (string, error) {
	if i >= NPC_MIN_ID && i <= NPC_MAX_ID {
		changelog := ""

		// Obtain original NPC data.
		(*data).current = &((*data).uncompressed)
		spriteIdentificator1 := NPCGetSpriteIdentificator(data, i)
		fieldB1 := NPCGetFieldB(data, i)
		roomLocation1 := NPCGetRoomLocation(data, i)
		typeOfPlace1 := NPCGetTypeOfPlace(data, i)
		fieldE1 := NPCGetFieldE(data, i)
		exactPlace1 := NPCGetExactPlace(data, i)
		forDialogue1 := NPCGetForDialogue(data, i)
		fieldH1 := NPCGetFieldH(data, i)

		// Obtain modified NPC data.
		(*data).current = &((*data).modified)
		spriteIdentificator2 := NPCGetSpriteIdentificator(data, i)
		fieldB2 := NPCGetFieldB(data, i)
		roomLocation2 := NPCGetRoomLocation(data, i)
		typeOfPlace2 := NPCGetTypeOfPlace(data, i)
		fieldE2 := NPCGetFieldE(data, i)
		exactPlace2 := NPCGetExactPlace(data, i)
		forDialogue2 := NPCGetForDialogue(data, i)
		fieldH2 := NPCGetFieldH(data, i)

		if (spriteIdentificator1 != spriteIdentificator2) || (fieldB1 != fieldB2) || (roomLocation1 != roomLocation2) || (typeOfPlace1 != typeOfPlace2) || (fieldE1 != fieldE2) || (exactPlace1 != exactPlace2) || (forDialogue1 != forDialogue2) || (fieldH1 != fieldH2) {
			changelog = NPCGetNameStr(i) + ":"

			if spriteIdentificator1 != spriteIdentificator2 {
				changelog += fmt.Sprintf(" sprite identificator changed from %d (%X) to %d (%X),", spriteIdentificator1, spriteIdentificator1, spriteIdentificator2, spriteIdentificator2)
			}
			if fieldB1 != fieldB2 {
				changelog += fmt.Sprintf(" B changed from %d (%X) to %d (%X),", fieldB1, fieldB1, fieldB2, fieldB2)
			}
			if roomLocation1 != roomLocation2 {
				changelog += fmt.Sprintf(" room location changed from %d (%X) to %d (%X),", roomLocation1, roomLocation1, roomLocation2, roomLocation2)
			}
			if typeOfPlace1 != typeOfPlace2 {
				changelog += fmt.Sprintf(" type of place changed from %d (%X) to %d (%X),", typeOfPlace1, typeOfPlace1, typeOfPlace2, typeOfPlace2)
			}
			if fieldE1 != fieldE2 {
				changelog += fmt.Sprintf(" E changed from %d (%X) to %d (%X),", fieldE1, fieldE1, fieldE2, fieldE2)
			}
			if exactPlace1 != exactPlace2 {
				changelog += fmt.Sprintf(" exact place changed from %d (%X) to %d (%X),", exactPlace1, exactPlace1, exactPlace2, exactPlace2)
			}
			if forDialogue1 != forDialogue2 {
				changelog += fmt.Sprintf(" for dialogue changed from %d (%X) to %d (%X),", forDialogue1, forDialogue1, forDialogue2, forDialogue2)
			}
			if fieldH1 != fieldH2 {
				changelog += fmt.Sprintf(" H changed from from %d (%X) to %d (%X),", fieldH1, fieldH1, fieldH2, fieldH2)
			}

			// Trim final ","
			if changelog[len(changelog)-1] == ',' {
				changelog = changelog[0 : len(changelog)-1]
				changelog += "."
			}

			specialDescription := data.specialNPCDescriptions[i]
			if specialDescription != "" {
				changelog += specialDescription
			}
		}

		return changelog, nil
	} else {
		return "", errors.New("Improper NPC ID")
	}
}

func NPCDiffAllProduceChangelogEntries(data *DuneMetadata) (string, error) {
	changelog := `NPC ADJUSTMENTS:

`
	for i := uint(NPC_MIN_ID); i <= NPC_MAX_ID; i++ {
		changelogNPC, err := NPCDiffProduceChangelogEntry(data, i)
		if err != nil {
			return "", err
		}
		if changelogNPC != "" {
			changelog += changelogNPC + "\n"
		}
	}

	if changelog[len(changelog)-1] == '\n' && changelog[len(changelog)-2] != '\n' {
		changelog += "\n"
	}
	return changelog, nil
}
