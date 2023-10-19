package windows

func CloseMidiOutDevice(device MidiOutDevice) error {
	returnCode, _, err := procMidiOutClose.Call(uintptr(device.Handle))
	if err != nil && returnCode != 0 {
		return err
	}
	return nil
}
