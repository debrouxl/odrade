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
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	FILE_FORMAT_NONE   = 0
	FILE_FORMAT_DUNE21 = 1
	FILE_FORMAT_DUNE23 = 2
	FILE_FORMAT_DUNE24 = 3
	FILE_FORMAT_DUNE37 = 4
)

const (
	RUN_MODE_MODIFY = 0
	RUN_MODE_PRINT  = 1
)

const RLE_BYTE = 0xF7

const LOCATION_MIN_ID = 1
const LOCATION_MAX_ID = 70
const LOCATION_SIZE = 28

const LOCATION_NAME1_MIN_ID = 0x1
const LOCATION_NAME1_MAX_ID = 0xC
const LOCATION_NAME2_MIN_ID = 0x1
const LOCATION_NAME2_MAX_ID = 0xB
const LOCATION_APPEARANCE_MIN_ID = 0x0
const LOCATION_APPEARANCE_MAX_ID = 0x30
const LOCATION_FIELD_MIN_ID = 0x1
const LOCATION_FIELD_MAX_ID = 0x4C
const LOCATION_STAGE_MAX = 0x68

const (
	LOCATION_STATUS_VEGETATION   = 0x01
	LOCATION_STATUS_IN_BATTLE    = 0x02
	LOCATION_STATUS_INVENTORY    = 0x10
	LOCATION_STATUS_WINDTRAP     = 0x20
	LOCATION_STATUS_PROSPECTED   = 0x40
	LOCATION_STATUS_UNDISCOVERED = 0x80
)

const TROOP_MIN_ID = 1
const TROOP_MAX_ID = 68
const TROOP_SIZE = 27

const TROOP_POSITION_TOP_LOCATION_FIRST = 9

const TROOP_OCCUPATION_NOT_HIRED_MASK = 0x80
const TROOP_OCCUPATION_FREMEN_MINING_SPICE = 0x0
const TROOP_OCCUPATION_FREMEN_WAITING_FOR_ORDERS = 0x2
const TROOP_OCCUPATION_HARKONNEN_MINING_SPICE = 0xC
const TROOP_OCCUPATION_HARKONNEN_PROSPECTING = 0xD
const TROOP_OCCUPATION_HARKONNEN_WAITING_FOR_ORDERS = 0xE
const TROOP_OCCUPATION_HARKONNEN_SPICE_MINERS_SEARCHING_FOR_EQUIPMENT = 0xF
const TROOP_OCCUPATION_HARKONNEN_MILITARIES_SEARCHING_FOR_EQUIPMENT = 0x1F
const TROOP_OCCUPATION_SPECIAL_FIRST = 0x10
const TROOP_OCCUPATION_MOVING_FIRST = 0x40
const TROOP_OCCUPATION_WAS_SLAVED_FIRST = 0xA0

const (
	TROOP_EQUIPMENT_BULB             = 0x02
	TROOP_EQUIPMENT_ATOMICS          = 0x04
	TROOP_EQUIPMENT_WEIRDING_MODULES = 0x08
	TROOP_EQUIPMENT_LASER_GUNS       = 0x10
	TROOP_EQUIPMENT_KRYS_KNIVES      = 0x20
	TROOP_EQUIPMENT_ORNITHOPTER      = 0x40
	TROOP_EQUIPMENT_HARVESTER        = 0x80
)

const NPC_MIN_ID = 1
const NPC_MAX_ID = 16
const NPC_SIZE = 8
const NPC_PADDING = 8

const SMUGGLER_MIN_ID = 1
const SMUGGLER_MAX_ID = 6
const SMUGGLER_SIZE = 14
const SMUGGLER_PADDING = 3

type DuneMetadata struct {
	format                      uint
	indata                      []byte
	uncompressed                []byte
	modified                    []byte
	compressed                  []byte
	current                     *[]byte
	changelog                   string
	dialoguesOffset             uint
	timeCountersOffset          uint
	locationsOffset             uint
	troopsOffset                uint
	npcsOffset                  uint
	smugglersOffset             uint
	shuffledTroops              map[uint]uint
	coordinatesMap              map[string]uint
	specialTroopDescriptions    map[uint]string
	specialLocationDescriptions map[uint]string
	specialNPCDescriptions      map[uint]string
	specialSmugglerDescriptions map[uint]string
	housedTroopsInLocations     [][]byte
	locationForTroops           [TROOP_MAX_ID+1]byte
	strictChecks                bool
}

func performInitialSanityCheck(data *DuneMetadata) bool {
	// We know how large a valid DUNE??S0.SAV file can be.
	if !((*data).strictChecks) || (len((*data).indata) > 9500 && len((*data).indata) < 10000) {
		if ((*data).indata)[2] == RLE_BYTE && ((*data).indata)[len((*data).indata)-1] != RLE_BYTE {
			var pos0 uint16 = uint16(((*data).indata)[4]) + 255*uint16(((*data).indata)[5])
			var expectedPos0 uint16 = uint16(len((*data).indata))
			switch data.format {
			case FILE_FORMAT_DUNE21:
				expectedPos0 = expectedPos0 - 39
			case FILE_FORMAT_DUNE23:
				expectedPos0 = expectedPos0 - 39
			case FILE_FORMAT_DUNE24:
				expectedPos0 = expectedPos0 - 39
			case FILE_FORMAT_DUNE37:
				expectedPos0 = expectedPos0 - 40
			}
			if !((*data).strictChecks) || pos0 == expectedPos0 {
				return true
			} else {
				fmt.Printf("Error: embedded length bytes %d do not match the expected value %d\n", pos0, expectedPos0)
			}
		} else {
			println("Warning: expected 0xF7 as RLE byte")
		}
	} else {
		println("File is too small or too large to be a valid DUNE??S0.SAV file")
	}
	return false
}

func performFinalComputationsAndSanityCheck(data *DuneMetadata) bool {
	ret := true

	// Map for associating 4-byte troop coordinates to location IDs.
	data.coordinatesMap = make(map[string]uint)

	// First pass, location checking part.
	for i := uint(LOCATION_MIN_ID); i <= LOCATION_MAX_ID; i++ {
		name1, name2 := LocationGetName(data, i)
		if name1 < LOCATION_NAME1_MIN_ID || name1 > LOCATION_NAME1_MAX_ID || name2 < LOCATION_NAME2_MIN_ID || name2 > LOCATION_NAME2_MAX_ID {
			fmt.Printf("Error: location #%d (%X) has invalid name\n", i, i)
			ret = false
		}

		coordinates := hex.EncodeToString(LocationGetCoordinates(data, i)[0:4])
		data.coordinatesMap[coordinates] = i

		appearance := LocationGetAppearance(data, i)
		if name2 == 0x0A && appearance != 0x21 {
			fmt.Printf("Error: location #%d (%X) has village name (-Pyons), but the appearance does not match the expected value\n", i, i)
			ret = false
		}
		if appearance > LOCATION_APPEARANCE_MAX_ID {
			fmt.Printf("Error: location #%d (%X) has improper appearance\n", i, i)
			ret = false
		}

		troopID := LocationGetTroopID(data, i)
		if troopID > TROOP_MAX_ID { // A location can have a troop ID of zero !
			fmt.Printf("Error: location #%d (%X) has invalid troop ID\n", i, i)
			ret = false
		}

		// Build the data structures which enable us to perform more advanced checks on troops and locations.
		data.housedTroopsInLocations[i][0] = troopID
		if troopID != 0 {
			locationForThisTroop := data.locationForTroops[troopID]
			if locationForThisTroop != 0 {
				// No troop should appear twice in the list of housed troops.
				// This check should also prevent loops in the troop chain (which send the game into an infinite loop - checked).
				fmt.Printf("Error: troop ID #%d (%X) declared housed at location #%d (%X) was already housed at location #%d (%X)\n", troopID, troopID, i, i, locationForThisTroop, locationForThisTroop)
				ret = false
			}
			data.locationForTroops[troopID] = byte(i)
			coordinates2 := hex.EncodeToString(TroopGetCoordinates(data, uint(troopID)))
			if coordinates != coordinates2 {
				fmt.Printf("Error: housed troop ID #%d (%X) coordinates do not match location #%d (%X) coordinates\n", troopID, troopID, i, i)
				ret = false
			}
		}

		status := LocationGetStatus(data, i)
		if status & 0xC != 0 {
			fmt.Printf("Warning: location #%d (%X) has undocumented status flags\n", i, i)
		}

		stage := LocationGetStage(data, i)
		if stage[0] > LOCATION_STAGE_MAX && stage[0] != 0xFF {
			fmt.Printf("Warning: location #%d (%X) has weird stage value, chances are that it won't be discoverable through travel\n", i, i)
		}

		fieldID := LocationGetFieldID(data, i)
		if fieldID < LOCATION_FIELD_MIN_ID || fieldID > LOCATION_FIELD_MAX_ID {
			fmt.Printf("Error: location #%d (%X) has invalid spice field ID\n", i, i)
			ret = false
		}

		// No possible sanity check on spice amount, spice density.

		fieldJ := LocationGetFieldJ(data, i)
		if fieldJ != 0x00 {
			fmt.Printf("Warning: location #%d (%X) has weird field J\n", i, i)
		}

		// No possible / desirable sanity check on harvesters, ornithopters, krys knives, laser guns, weirding modules, atomics, bulbs, water.
	}

	// First pass, troop checking part.
	positionsMask := make([]uint16, LOCATION_MAX_ID+1)
	eightzeros := make([]byte, 8)
	for i := uint(TROOP_MIN_ID); i <= TROOP_MAX_ID; i++ {
		troopID := TroopGetTroopID(data, i)
		if i == 68 && troopID == 0 {
			continue // Assume vanilla Dune and skip this entry.
		}
		if troopID != byte(i) {
			fmt.Printf("Error: troop #%d (%X) has invalid troop ID\n", i, i)
			ret = false
		}

		nextTroopID := TroopGetNextTroopID(data, i)
		if nextTroopID > TROOP_MAX_ID { // A next troop ID of zero is valid !
			fmt.Printf("Error: troop #%d (%X) has invalid next troop ID\n", i, i)
			ret = false
		}
		locationForThisTroop := data.locationForTroops[i]
		locationForNextTroop := data.locationForTroops[nextTroopID]
		if nextTroopID != 0 && locationForNextTroop != 0 {
			fmt.Printf("Error: troop #%d (%X) is already part of another location\n", nextTroopID, nextTroopID)
			ret = false
		}
		data.locationForTroops[nextTroopID] = locationForThisTroop
		// The location for the first troop was already filled in above when indexing the locations.
		if nextTroopID != 0 {
			data.housedTroopsInLocations[locationForThisTroop] = append(data.housedTroopsInLocations[locationForThisTroop], nextTroopID)
		}

		troopPosition := TroopGetPosition(data, i)
		troopOccupation := TroopGetOccupation(data, i)
		fieldG := TroopGetFieldG(data, i)
		if troopOccupation < TROOP_OCCUPATION_SPECIAL_FIRST || (troopOccupation >= TROOP_OCCUPATION_NOT_HIRED_MASK && troopOccupation < TROOP_OCCUPATION_NOT_HIRED_MASK + TROOP_OCCUPATION_SPECIAL_FIRST) { // 0x00-0xF, 0x80-0x8F
			troopOccupation &= 0xFF ^ TROOP_OCCUPATION_NOT_HIRED_MASK
			if troopOccupation == TROOP_OCCUPATION_HARKONNEN_MINING_SPICE {
				if troopPosition < TROOP_POSITION_TOP_LOCATION_FIRST && i != 67 { // Arrakeen (Harkonnen) palace troop
					fmt.Printf("Error: troop #%d (%X) is Harkonnen but has improper position\n", i, i)
					ret = false
				}
				if bytes.Compare(fieldG, eightzeros) == 0 {
					fmt.Printf("Error: troop #%d (%X) is Harkonnen but has all zeros in field G\n", i, i)
					ret = false
				}
			} else if troopOccupation == TROOP_OCCUPATION_HARKONNEN_PROSPECTING || troopOccupation == TROOP_OCCUPATION_HARKONNEN_WAITING_FOR_ORDERS || troopOccupation == TROOP_OCCUPATION_HARKONNEN_SPICE_MINERS_SEARCHING_FOR_EQUIPMENT || troopOccupation == TROOP_OCCUPATION_HARKONNEN_MILITARIES_SEARCHING_FOR_EQUIPMENT {
				fmt.Printf("Error: troop #%d (%X) is Harkonnen but has weird occupation\n", i, i)
				ret = false
			} else if troopOccupation == TROOP_OCCUPATION_FREMEN_MINING_SPICE || troopOccupation == TROOP_OCCUPATION_FREMEN_WAITING_FOR_ORDERS {
				if troopPosition >= TROOP_POSITION_TOP_LOCATION_FIRST {
					fmt.Printf("Error: troop #%d (%X) is Fremen but has improper position\n", i, i)
					ret = false
				}
				if bytes.Compare(fieldG, eightzeros) != 0 {
					fmt.Printf("Error: troop #%d (%X) is Fremen but has all some nonzero bytes in field G\n", i, i)
					ret = false
				}
			} else {
				fmt.Printf("Warning: troop #%d (%X) is Fremen but has weird occupation\n", i, i)
			}
		} else {
			// 0x10-0x1F, 0x20-0x2F, 0x30-0x3F: special occupations (job finished, captured, etc.)
			// 0x40-0x7F: troop moving.
			// 0x90-0x9F: "not yet hired" version of 0x10-0x1F.
			// 0xA0-0xFF: Fremen complaining about slaving.
			fmt.Printf("Error: troop #%d (%X) has weird occupation\n", i, i)
			ret = false
		}

		// No possible check on dissatisfaction ?

		speech := TroopGetSpeech(data, i)
		if speech != 0 {
			fmt.Printf("Note: troop #%d (%X) has nonzero speech\n", i, i)
		}

		fieldJ := TroopGetFieldJ(data, i)
		if fieldJ != 0x00 {
			fmt.Printf("Warning: troop #%d (%X) has weird field J\n", i, i)
		}

		motivation := TroopGetMotivation(data, i)
		if motivation > 254 {
			fmt.Printf("Warning: troop #%d (%X) has excessive motivation\n", i, i)
		}

		spiceMiningSkill := TroopGetSpiceMiningSkill(data, i)
		if spiceMiningSkill > 100 {
			fmt.Printf("Warning: troop #%d (%X) has excessive spice mining skill\n", i, i)
		}

		armySkill := TroopGetArmySkill(data, i)
		if armySkill > 207 {
			fmt.Printf("Warning: troop #%d (%X) has excessive army skill, which can cause undesirable effects in battles unless it carries only Atomics\n", i, i)
		}

		ecologySkill := TroopGetEcologySkill(data, i)
		if ecologySkill > 100 {
			fmt.Printf("Warning: troop #%d (%X) has excessive ecology skill\n", i, i)
		}

		// No possible check on troop equipment ?

		population := TroopGetPopulation(data, i)
		if population < 300 && i != 67 { // Arrakeen (Harkonnen) palace troop
			fmt.Printf("Warning: troop #%d (%X) has very low population\n", i, i)
		}

		coordinates := hex.EncodeToString(TroopGetCoordinates(data, i))
		locationIdx := data.coordinatesMap[coordinates]
		if locationIdx == 0 {
			fmt.Printf("Warning: troop #%d (%X) has coordinates not referenced in the locations array\n", i, i)
			//ret = false
		}
		if positionsMask[locationIdx]&(1<<troopPosition) != 0 {
			fmt.Printf("Error: troop #%d (%X) has the same position in location as another troop\n", i, i)
			ret = false
		}
		positionsMask[locationIdx] |= (1 << troopPosition)
	}

	// Second pass.
	/*println("locationForTroops")
	for i := uint(TROOP_MIN_ID); i <= TROOP_MAX_ID; i++ {
		fmt.Printf("Troop #%d (%X): %x\n", i, i, data.locationForTroops[i])
	}*/
	//println("housedTroopsInLocations")
	for i := uint(LOCATION_MIN_ID); i <= LOCATION_MAX_ID; i++ {
		//fmt.Printf("Location #%d (%X) (len %d): %x\n", i, i, len(data.housedTroopsInLocations[i]), data.housedTroopsInLocations[i])

		// Complain if there are any locations where there are both Fremen and Harkonnen troops;
		// TODO All troops at the same location should have the same value for field E, shouldn't they ?
		// TODO All troops at the same location should have the same coordinates.
		// TODO Check the index and the list wrt. the position: first troop should be at 1 or 9, second troop at 2 or A, etc.
		foundFremenTroops := false
		foundHarkonnenTroops := false
		position := 1
		for j := 0; j < len(data.housedTroopsInLocations[i]); j++ {
			troopOccupation := TroopGetOccupation(data, uint(data.housedTroopsInLocations[i][j])) & (0xFF ^ TROOP_OCCUPATION_NOT_HIRED_MASK);
			if troopOccupation == TROOP_OCCUPATION_HARKONNEN_MINING_SPICE {
				foundHarkonnenTroops = true;
			} else {
				// Other values were filtered previously.
				foundFremenTroops = true
			}
			position++
		}
		if foundFremenTroops && foundHarkonnenTroops {
			fmt.Printf("Location #%d (%X) has a mix of Fremen and Harkonnen troops, that shouldn't happen\n", i, i)
			ret = false;
		}

	}


	return ret
}

func GetOffsets(mode uint) (uint, uint, uint, uint, uint, uint) {
	var dialoguesOffset uint = 0x3339
	var timeCountersOffset uint
	var locationsOffset uint
	var troopsOffset uint
	var npcsOffset uint
	var smugglersOffset uint
	switch mode {
	case FILE_FORMAT_DUNE21:
		timeCountersOffset = 0x43BB
		locationsOffset = 0x44BB
		troopsOffset = 0x4C65
		npcsOffset = 0x5391
		smugglersOffset = 0x5493
	case FILE_FORMAT_DUNE23:
		timeCountersOffset = 0x43CF
		locationsOffset = 0x44CF
		troopsOffset = 0x4C79
		npcsOffset = 0x53A5
		smugglersOffset = 0x54A7
	case FILE_FORMAT_DUNE24:
		timeCountersOffset = 0x43D9
		locationsOffset = 0x44D9
		troopsOffset = 0x4C83
		npcsOffset = 0x53AF
		smugglersOffset = 0x54B1
	case FILE_FORMAT_DUNE37:
		timeCountersOffset = 0x441F
		locationsOffset = 0x451F
		troopsOffset = 0x4CC9
		npcsOffset = 0x53F5
		smugglersOffset = 0x54F7
	}
	return dialoguesOffset, timeCountersOffset, locationsOffset, troopsOffset, npcsOffset, smugglersOffset
}

func checkSupportedVersion(data *DuneMetadata) error {
	// Compute the hash of everything but the time counters (sub-period counter + period counter).
	hasher := sha512.New()
	hasher.Write(((*data).uncompressed)[:data.timeCountersOffset])
	hasher.Write(((*data).uncompressed)[data.timeCountersOffset+4:])
	value := hasher.Sum(nil)
	valueStr := hex.EncodeToString(value[:])

	var knownvalue string
	switch data.format {
	case FILE_FORMAT_DUNE21:
		knownvalue = "c1085ed018e71443be3d2cb4337f52b8599e3823fb4fb7e0219eac3058a1d7f814e99eb048248cab9ff4b10e85f3e03ef1b37c93037d4b14b5a8a7a35212e41c"
	case FILE_FORMAT_DUNE23:
		knownvalue = "55b97b61b6cf0764974a9eb5463ee6d12a075c63acdeda59754ab8e56236bba9c12d3adc40757678f8faa84e933cfaa34b6552a646c289fd6edcc587996b5eff"
	case FILE_FORMAT_DUNE24:
		knownvalue = "b8d0657c9f4f4fa7c15d15ec92111bba3a2de68630d7a39bb562603a7a1eedf171c81968cac61e63ce70cfcbd7cb7444ec0d417dd3d0a54a000b19d2619a73ba"
	case FILE_FORMAT_DUNE37:
		knownvalue = "4f1ba16132e9f9f374954d1084ca89f9dbc78221fa4399ec31493938edc53a004f73f49543fc7c1585124a9d7b090fc6fd93bb5e9093f847dd4a4e809aa9b1b8"
	}

	//println(valueStr)
	if valueStr != knownvalue {
		return errors.New("Unexpected contents for the input file")
	} else {
		println("File content matches the expected one, good") // Well, almost - but in 2021, no collisions are publicly known for the SHA-2 family :)
		return nil
	}
}

func RLEDecompress(indata *[]byte) ([]byte, error) {
	// Temporarily modify RLE byte value, and expand the array by two dummy bytes so that indata[x+2] falls within array bounds.
	//println(len(*indata))
	(*indata)[2] = 0xF6
	(*indata) = append((*indata), (*indata)[len((*indata))-1]+1)
	(*indata) = append((*indata), (*indata)[len((*indata))-1]+1)

	// Even if the file were almost entirely made of F7 FF xx sequences, the size of the output would be less than 9 MB, which isn't really excessive on remotely modern computers.
	outdata := make([]byte, 0)
	for i := 0; i < len((*indata))-2; {
		if (*indata)[i] != RLE_BYTE {
			outdata = append(outdata, (*indata)[i])
			i = i + 1
		} else {
			for j := 0; j < int((*indata)[i+1]); j++ {
				outdata = append(outdata, (*indata)[i+2])
			}
			i = i + 3
		}
	}

	if len(outdata) < 22000 || len(outdata) >= 22200 {
		return outdata, errors.New("Uncompressed data is unexpectedly small or large")
	}

	return outdata, nil
}

func RLECompress(indata *[]byte) ([]byte, error) {
	// Temporarily modify RLE byte value, and expand the array by two dummy bytes so that indata[x+2] falls within array bounds.
	*indata = append(*indata, (*indata)[len(*indata)-1]+1)
	*indata = append(*indata, (*indata)[len(*indata)-1]+1)

	outdata := make([]byte, 0)
	for i := 0; i < len(*indata)-2; {
		b1 := (*indata)[i]
		b2 := (*indata)[i+1]
		b3 := (*indata)[i+2]
		// RLE byte ?
		if b1 == RLE_BYTE {
			// Yup, replace it.
			outdata = append(outdata, RLE_BYTE)
			outdata = append(outdata, 0x01)
			outdata = append(outdata, RLE_BYTE)
			i = i + 1
		} else {
			if b1 == b2 && b2 == b3 {
				// Can compress this sequence of at least three consecutive identical bytes.
				c := uint(3)
				j := i + 3
				for {
					if c > 255 {
						// Cap the current RLE sequence to 255 bytes.
						outdata = append(outdata, RLE_BYTE)
						outdata = append(outdata, 0xFF)
						outdata = append(outdata, b1)
						i = i + 255
						break
					}
					if (*indata)[j] == b1 {
						c++
						j++
					} else {
						// Found the end of the sequence
						outdata = append(outdata, RLE_BYTE)
						outdata = append(outdata, byte(c))
						outdata = append(outdata, b1)
						i = j
						break
					}
				}
			} else {
				// Output the first byte as is, introducing a RLE sequence would waste space.
				outdata = append(outdata, b1)
				i = i + 1
			}
		}
	}

	return outdata, nil
}

func usage() {
	println("ODRADE 0.1")
	println("Usage: odrade modify <input file> - apply modifications to the given input file")
	println("       odrade print <input file> - print data contained in file")
	os.Exit(1)
}

func main() {
	var data DuneMetadata
	var err error

	if len(os.Args) < 3 {
		usage()
	}

	var runMode uint
	if os.Args[1] == "modify" {
		runMode = RUN_MODE_MODIFY
	} else if os.Args[1] == "print" {
		runMode = RUN_MODE_PRINT
	} else {
		usage()
	}

	var fileFormat uint = FILE_FORMAT_NONE
	if strings.Contains(os.Args[2], "DUNE21S") {
		println("Probably targeting DUNE21")
		fileFormat = FILE_FORMAT_DUNE21
	} else if strings.Contains(os.Args[2], "DUNE23S") {
		println("Probably targeting DUNE23")
		fileFormat = FILE_FORMAT_DUNE23
	} else if strings.Contains(os.Args[2], "DUNE24S") {
		println("Probably targeting DUNE24")
		fileFormat = FILE_FORMAT_DUNE24
	} else if strings.Contains(os.Args[2], "DUNE37S") {
		println("Probably targeting DUNE37 (CD version)")
		fileFormat = FILE_FORMAT_DUNE37
	}
	if fileFormat == FILE_FORMAT_NONE {
		println("Error: file type / version not recognized from its name")
		os.Exit(1)
	} /*else {
		fmt.Printf("File type is probably %d\n", mode)
	}*/

	data.format = fileFormat
	data.strictChecks = true

	// If users want to DoS their computers by feeding a large file into this program... too bad for them ?
	data.indata, err = os.ReadFile(os.Args[2])
	if err != nil {
		println("Error: failed to read input file")
		os.Exit(1)
	}

	if !performInitialSanityCheck(&data) {
		println("Error: input file validation failed")
		os.Exit(1)
	}

	data.uncompressed, err = RLEDecompress(&(data.indata))
	if err != nil {
		println("Error: decompression failed")
		os.Exit(1)
	}
	//println(len(data.uncompressed))

	err = os.WriteFile(os.Args[2]+"_uncompressed", data.uncompressed, 0644)
	if err != nil {
		println("Error: writing uncompressed file failed")
		os.Exit(1)
	}

	data.dialoguesOffset, data.timeCountersOffset, data.locationsOffset, data.troopsOffset, data.npcsOffset, data.smugglersOffset = GetOffsets(data.format)

	err = checkSupportedVersion(&data)
	if err != nil {
		println("Warning: file format is not supported, you're on your own")
	}

	if runMode == RUN_MODE_MODIFY {
		err = SetConfig(&data)
		if err != nil {
			println("Error: could not set config")
			os.Exit(1)
		}

		// Make a copy of the uncompressed input data, so that we can perform a diff later to (eventually) produce the changelog.
		data.modified = make([]byte, len(data.uncompressed))
		copy(data.modified[:], data.uncompressed[:])

		data.current = &data.modified

		// Initialize remaining variables.
		data.shuffledTroops = make(map[uint]uint)
		for i := uint(TROOP_MIN_ID); i < TROOP_MAX_ID; i++ {
			(data.shuffledTroops)[i] = i
		}
		data.specialTroopDescriptions = make(map[uint]string)
		data.specialLocationDescriptions = make(map[uint]string)
		data.specialNPCDescriptions = make(map[uint]string)
		data.specialSmugglerDescriptions = make(map[uint]string)
		data.housedTroopsInLocations = make([][]byte, LOCATION_MAX_ID+1)
		for i := uint(LOCATION_MIN_ID); i <= LOCATION_MAX_ID; i++ {
			data.housedTroopsInLocations[i] = make([]byte, 1)
		}

		// Modify data inside a function located into another file - that's how the set of modifications can be selected.
		err = ModifyTroopAndLocationData(&data)
		if err != nil {
			println("Error: modification 1 failed")
			os.Exit(1)
		}
		//println(len(data.modified))

		err = ModifyNPCAndSmugglerData(&data)
		if err != nil {
			println("Error: modification 2 failed")
			os.Exit(1)
		}
		//println(len(data.modified))

		err = os.WriteFile(os.Args[2]+"_modified", data.modified, 0644)
		if err != nil {
			println("Error: writing modified, uncompressed file failed")
			os.Exit(1)
		}

		if !performFinalComputationsAndSanityCheck(&data) {
			println("Error: final computations and sanity check failed")
			os.Exit(1)
		}

		// Changelog generation... here we go :)
		data.changelog, err = ChangelogPrologue()

		changelogTroops, err := TroopDiffAllProduceChangelogEntries(&data)
		if err != nil {
			println("Error: changelog generation for troops failed")
			os.Exit(1)
		}
		data.changelog += changelogTroops
		println(changelogTroops)

		changelogLocations, err := LocationDiffAllProduceChangelogEntries(&data)
		if err != nil {
			println("Error: changelog generation for locations failed")
			os.Exit(1)
		}
		data.changelog += changelogLocations
		println(changelogLocations)

		changelogNPCs, err := NPCDiffAllProduceChangelogEntries(&data)
		if err != nil {
			println("Error: changelog generation for NPCs failed")
			os.Exit(1)
		}
		data.changelog += changelogNPCs
		println(changelogNPCs)

		changelogSmugglers, err := SmugglerDiffAllProduceChangelogEntries(&data)
		if err != nil {
			println("Error: changelog generation for smugglers failed")
			os.Exit(1)
		}
		data.changelog += changelogSmugglers
		println(changelogSmugglers)

		epilogue, err := ChangelogEpilogue()
		if err != nil {
			println("Error: changelog epilogue generation failed")
			os.Exit(1)
		}
		data.changelog += epilogue

		// Perform RLE compression.
		data.compressed, err = RLECompress(&data.modified)
		if err != nil {
			println("Error: compression failed")
			os.Exit(1)
		}
		//println(len(data.compressed))

		// Update the two bytes at offsets 4-5 according to the compressed size
		pos0 := uint16(len(data.compressed))
		switch fileFormat {
		case FILE_FORMAT_DUNE21, FILE_FORMAT_DUNE23, FILE_FORMAT_DUNE24:
			pos0 = pos0 - 39
		case FILE_FORMAT_DUNE37:
			pos0 = pos0 - 40
		}
		data.compressed[4] = byte(pos0 % 255)
		data.compressed[5] = byte(pos0 / 255)

		// Restore RLE byte value.
		data.compressed[2] = RLE_BYTE

		// Write compressed output file.
		err = os.WriteFile(os.Args[2]+"_compressed", data.compressed, 0644)
		if err != nil {
			println("Error: writing modified, compressed file failed")
			os.Exit(1)
		}
	} else if runMode == RUN_MODE_PRINT {
		data.current = &data.uncompressed

		for i := uint(LOCATION_MIN_ID); i <= LOCATION_MAX_ID; i++ {
			LocationPrint(&data, i)
		}
		for i := uint(TROOP_MIN_ID); i <= TROOP_MAX_ID; i++ {
			TroopPrint(&data, i)
		}

		for i := uint(NPC_MIN_ID); i <= NPC_MAX_ID; i++ {
			NPCPrint(&data, i)
		}
		for i := uint(SMUGGLER_MIN_ID); i <= SMUGGLER_MAX_ID; i++ {
			SmugglerPrint(&data, i)
		}
	}

	os.Exit(0)
}
