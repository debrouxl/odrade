// Copyright (C) 2021 Lionel Debroux
// SPDX-License-Identifier: GPL-2.0
//
// Sources of information:
// * https://forum.dune2k.com/topic/20497-dune-cheats/
// * hugslab
// * Dmitri Fatkin

package main

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	MODE_NONE   = 0
	MODE_DUNE21 = 1
	MODE_DUNE23 = 2
	MODE_DUNE24 = 3
	MODE_DUNE37 = 4
)

const RLE_BYTE = 0xF7

const LOCATION_MAX_ID = 70
const LOCATION_SIZE = 28

const (
	LOCATION_STATUS_VEGETATION   = 0x01
	LOCATION_STATUS_IN_BATTLE    = 0x02
	LOCATION_STATUS_INVENTORY    = 0x10
	LOCATION_STATUS_WINDTRAP     = 0x20
	LOCATION_STATUS_PROSPECTED   = 0x40
	LOCATION_STATUS_UNDISCOVERED = 0x80
)

const TROOP_MAX_ID = 68
const TROOP_SIZE = 27

const TROOP_POSITION_TOP_LOCATION_FIRST = 9

const TROOP_OCCUPATION_NOT_HIRED_MASK = 0x80
const TROOP_OCCUPATION_HARKONNEN_MINING_SPICE = 0xC

const (
	TROOP_EQUIPMENT_BULB             = 0x02
	TROOP_EQUIPMENT_ATOMICS          = 0x04
	TROOP_EQUIPMENT_WEIRDING_MODULES = 0x08
	TROOP_EQUIPMENT_LASER_GUNS       = 0x10
	TROOP_EQUIPMENT_KRYS_KNIVES      = 0x20
	TROOP_EQUIPMENT_ORNITHOPTER      = 0x40
	TROOP_EQUIPMENT_HARVESTER        = 0x80
)

const NPC_MAX_ID = 15
const NPC_SIZE = 8
const NPC_PADDING = 8

const SMUGGLER_MAX_ID = 6
const SMUGGLER_SIZE = 14
const SMUGGLER_PADDING = 3

type DuneMetadata struct {
	mode               uint
	indata             []byte
	uncompressed       []byte
	modified           []byte
	compressed         []byte
	changelog          string
	shuffledTroops     map[uint]uint
	dialoguesOffset    uint
	timeCountersOffset uint
	locationsOffset    uint
	troopsOffset       uint
	npcsOffset         uint
	smugglersOffset    uint
}

func performInitialSanityCheck(data *DuneMetadata) bool {
	// We know how large a valid DUNE??S0.SAV file can be.
	if len((*data).indata) > 9500 && len((*data).indata) < 10000 {
		if ((*data).indata)[2] == RLE_BYTE && ((*data).indata)[len((*data).indata)-1] != RLE_BYTE {
			var pos0 uint16 = uint16(((*data).indata)[4]) + 255*uint16(((*data).indata)[5])
			var expectedPos0 uint16 = uint16(len((*data).indata))
			switch data.mode {
			case MODE_DUNE21:
				expectedPos0 = expectedPos0 - 39
			case MODE_DUNE23:
				expectedPos0 = expectedPos0 - 39
			case MODE_DUNE24:
				expectedPos0 = expectedPos0 - 39
			case MODE_DUNE37:
				expectedPos0 = expectedPos0 - 40
			}
			if pos0 == expectedPos0 {
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

func performFinalSanityCheck(data *DuneMetadata, mode uint) bool {
	ret := true

	for i := uint(1); i <= TROOP_MAX_ID; i++ {
		if TroopGetOccupation(data, i) == TROOP_OCCUPATION_HARKONNEN_MINING_SPICE|TROOP_OCCUPATION_NOT_HIRED_MASK && TroopGetPosition(data, i) < TROOP_POSITION_TOP_LOCATION_FIRST && i != 67 {
			fmt.Printf("Warning: troop #%d is Harkonnen but has improper position\n", i)
			ret = false
		}
		if TroopGetOccupation(data, i) != TROOP_OCCUPATION_HARKONNEN_MINING_SPICE|TROOP_OCCUPATION_NOT_HIRED_MASK && TroopGetPosition(data, i) >= TROOP_POSITION_TOP_LOCATION_FIRST {
			fmt.Printf("Warning: troop #%d is Fremen but has improper position\n", i)
			ret = false
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
	case MODE_DUNE21:
		timeCountersOffset = 0x43BB
		locationsOffset = 0x44BB
		troopsOffset = 0x4C65
		npcsOffset = 0x5391
		smugglersOffset = 0x5393
	case MODE_DUNE23:
		timeCountersOffset = 0x43CF
		locationsOffset = 0x44CF
		troopsOffset = 0x4C79
		npcsOffset = 0x53A5
		smugglersOffset = 0x54A7
	case MODE_DUNE24:
		timeCountersOffset = 0x43D9
		locationsOffset = 0x44D9
		troopsOffset = 0x4C83
		npcsOffset = 0x53AF
		smugglersOffset = 0x54B1
	case MODE_DUNE37:
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
	switch data.mode {
	case MODE_DUNE21:
		knownvalue = "c1085ed018e71443be3d2cb4337f52b8599e3823fb4fb7e0219eac3058a1d7f814e99eb048248cab9ff4b10e85f3e03ef1b37c93037d4b14b5a8a7a35212e41c"
	case MODE_DUNE23:
		knownvalue = "55b97b61b6cf0764974a9eb5463ee6d12a075c63acdeda59754ab8e56236bba9c12d3adc40757678f8faa84e933cfaa34b6552a646c289fd6edcc587996b5eff"
	case MODE_DUNE24:
		knownvalue = "b8d0657c9f4f4fa7c15d15ec92111bba3a2de68630d7a39bb562603a7a1eedf171c81968cac61e63ce70cfcbd7cb7444ec0d417dd3d0a54a000b19d2619a73ba"
	case MODE_DUNE37:
		knownvalue = "4f1ba16132e9f9f374954d1084ca89f9dbc78221fa4399ec31493938edc53a004f73f49543fc7c1585124a9d7b090fc6fd93bb5e9093f847dd4a4e809aa9b1b8"
	}

	println(valueStr)
	if valueStr != knownvalue {
		return errors.New("Unexpected contents for the input file")
	} else {
		println("File content matches the expected one, good") // Well, almost - but in 2021, no collisions are publicly known for the SHA-2 family :)
		return nil
	}
}

func RLEDecompress(indata *[]byte) ([]byte, error) {
	// Temporarily modify RLE byte value, and expand the array by two dummy bytes so that indata[x+2] falls within array bounds.
	println(len(*indata))
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

func main() {
	var data DuneMetadata
	var err error

	if len(os.Args) < 2 {
		println("Usage: dunes0mod <input file>")
		os.Exit(1)
	}

	var mode uint = MODE_NONE
	if strings.HasSuffix(os.Args[1], "DUNE21S0.SAV") {
		println("Probably targeting DUNE21")
		mode = MODE_DUNE21
	} else if strings.HasSuffix(os.Args[1], "DUNE23S0.SAV") {
		println("Probably targeting DUNE23")
		mode = MODE_DUNE23
	} else if strings.HasSuffix(os.Args[1], "DUNE24S0.SAV") {
		println("Probably targeting DUNE24")
		mode = MODE_DUNE24
	} else if strings.HasSuffix(os.Args[1], "DUNE37S0.SAV") {
		println("Probably targeting DUNE37 (CD version)")
		mode = MODE_DUNE37
	}
	if mode == MODE_NONE {
		println("File type / version not recognized from its name")
		os.Exit(1)
	} /*else {
		fmt.Printf("File type is probably %d\n", mode)
	}*/

	data.mode = mode

	// If users want to DoS their computers by feeding a large file into this program... too bad for them ?
	data.indata, err = os.ReadFile(os.Args[1])
	if err != nil {
		println("Failed to read input file")
		os.Exit(1)
	}

	if !performInitialSanityCheck(&data) {
		println("Input file validation failed")
		os.Exit(1)
	}

	data.uncompressed, err = RLEDecompress(&(data.indata))
	if err != nil {
		println("Decompression failed")
		os.Exit(1)
	}
	println(len(data.uncompressed))

	err = os.WriteFile(os.Args[1]+"_uncompressed", data.uncompressed, 0644)
	if err != nil {
		println("Writing uncompressed file failed")
		os.Exit(1)
	}

	err = checkSupportedVersion(&data)
	if err != nil {
		println("Warning: file format is not supported, you're on your own")
	}

	data.dialoguesOffset, data.timeCountersOffset, data.locationsOffset, data.troopsOffset, data.npcsOffset, data.smugglersOffset = GetOffsets(data.mode)

	// Make a copy of the uncompressed input data, so that we can perform a diff later to (eventually) produce the changelog.
	data.modified = make([]byte, len(data.uncompressed))
	copy(data.modified[:], data.uncompressed[:])

	// Modify data inside a function located into another file - that's how the set of modifications can be selected.
	err = ModifyTroopAndLocationData(&data)
	if err != nil {
		println("Modification 1 failed")
		os.Exit(1)
	}
	println(len(data.modified))

	err = ModifyNPCAndSmugglerData(&data)
	if err != nil {
		println("Modification 2 failed")
		os.Exit(1)
	}
	println(len(data.modified))

	err = os.WriteFile(os.Args[1]+"_modified", data.modified, 0644)
	if err != nil {
		println("Writing modified, uncompressed file failed")
		os.Exit(1)
	}

	if !performFinalSanityCheck(&data, mode) {
		println("Final sanity check failed")
		os.Exit(1)
	}

	data.changelog, err = ChangelogPrologue()

	// TODO for changelog generation:
	// troop diff: traversing locations in ascending order, print Harkonnen troops diff, then Fremen troops diff
	// It may be interesting to build up a locations map
	// Shuffle map.
	// locations diff: traversing locations in ascending order

	println(data.changelog[0:100])

	// Perform RLE compression.
	data.compressed, err = RLECompress(&data.modified)
	if err != nil {
		println("Compression failed")
		os.Exit(1)
	}
	println(len(data.compressed))

	// Update the two bytes at offsets 4-5 according to the compressed size
	pos0 := uint16(len(data.compressed))
	switch mode {
	case MODE_DUNE21, MODE_DUNE23, MODE_DUNE24:
		pos0 = pos0 - 39
	case MODE_DUNE37:
		pos0 = pos0 - 40
	}
	data.compressed[4] = byte(pos0 % 255)
	data.compressed[5] = byte(pos0 / 255)

	// Restore RLE byte value.
	data.compressed[2] = RLE_BYTE

	// Write compressed output file.
	err = os.WriteFile(os.Args[1]+"_compressed", data.compressed, 0644)
	if err != nil {
		println("Writing modified, compressed file failed")
		os.Exit(1)
	}

	os.Exit(0)
}
