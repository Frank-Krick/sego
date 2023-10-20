package windows

import (
	"bytes"
	"encoding/binary"
	"errors"
)

func (device MidiOutDevice) SendNoteOnMessage(note byte, velocity byte, channel int) error {
	if err := validateChannel(channel); err != nil {
		return err
	}

	messageBytes := [4]byte{0, velocity, note, getChannelByte(channel)}
	var message uint32
	if err := binary.Read(bytes.NewReader(messageBytes[:]), binary.BigEndian, &message); err != nil {
		return err
	}

	returnCode, _, err := procMidiOutShortMsg.Call(uintptr(device.Handle), uintptr(message))
	if err != nil && returnCode != 0 {
		return err
	}
	return nil
}

func (device MidiOutDevice) SendNoteOffMessage(note byte, channel int) error {
	return device.SendNoteOnMessage(note, 0, channel)
}

func validateChannel(channel int) error {
	if channel < 1 && channel > 16 {
		return errors.New("channel needs to be between 1 and 16 (inclusive)")
	}
	return nil
}

func getChannelByte(channel int) byte {
	switch channel {
	case 1:
		return 0x90
	case 2:
		return 0x91
	case 3:
		return 0x92
	case 4:
		return 0x93
	case 5:
		return 0x94
	case 6:
		return 0x95
	case 7:
		return 0x96
	case 8:
		return 0x97
	case 9:
		return 0x98
	case 10:
		return 0x99
	case 11:
		return 0x9a
	case 12:
		return 0x9b
	case 13:
		return 0x9c
	case 14:
		return 0x9d
	case 15:
		return 0x9e
	case 16:
		return 0x9f
	}
	return 0
}
