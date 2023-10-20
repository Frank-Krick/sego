package windows

import (
	"bytes"
	"encoding/binary"
	"github.com/samber/lo"
	"golang.org/x/sys/windows"
	"unsafe"
)

type MidiOutCapabilities struct {
	ManufacturerId    uint16
	ProductIdentifier uint16
	Version           uint32
	Name              [32]byte
	Technology        uint16
	Voices            uint16
	Notes             uint16
	ChannelMask       uint16
	Support           uint32
}

var sizeMidiOutCapabilities = unsafe.Sizeof(MidiOutCapabilities{})

type MidiOutDeviceDescription struct {
	Name        string
	DeviceIndex int
}

func ListMidiOutDevices() ([]MidiOutDeviceDescription, error) {
	winNumberOfDevices, _, _ := procMidiOutGetNumDevs.Call()
	numberOfDevices := int(winNumberOfDevices)
	midiOutCapabilities := make([]MidiOutCapabilities, numberOfDevices)
	for deviceIndex := 0; deviceIndex < numberOfDevices; deviceIndex++ {
		var plainData [64]byte
		returnCode, _, err := procMidiOutGetDevCaps.Call(uintptr(deviceIndex), uintptr(unsafe.Pointer(&plainData[0])), sizeMidiOutCapabilities)
		if err != nil && returnCode != 0 {
			return nil, err
		}

		err = binary.Read(bytes.NewBuffer(plainData[:]), binary.NativeEndian, &midiOutCapabilities[deviceIndex])
		if err != nil {
			return nil, err
		}
	}

	return lo.Map(midiOutCapabilities[:], func(descriptor MidiOutCapabilities, index int) MidiOutDeviceDescription {
		return MidiOutDeviceDescription{DeviceIndex: index, Name: windows.ByteSliceToString(descriptor.Name[:])}
	}), nil
}

type MidiOutDevice struct {
	Handle uint
}

func OpenMidiOutDevice(deviceInfo MidiOutDeviceDescription) (MidiOutDevice, error) {
	midiOutDevice := MidiOutDevice{}
	returnCode, _, err := procMidiOutOpen.Call(uintptr(unsafe.Pointer(&midiOutDevice.Handle)), uintptr(deviceInfo.DeviceIndex), uintptr(0), 0)
	if err != nil && returnCode != 0 {
		return MidiOutDevice{}, err
	}

	return midiOutDevice, nil
}

func CloseMidiOutDevice(device MidiOutDevice) error {
	returnCode, _, err := procMidiOutClose.Call(uintptr(device.Handle))
	if err != nil && returnCode != 0 {
		return err
	}
	return nil
}
