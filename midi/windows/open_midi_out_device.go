package windows

import "unsafe"

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
