package windows

func SendShortNoteOnMessage(device MidiOutDevice, note int, velocity int, channel int) error {
	returnCode, _, err := procMidiOutShortMsg.Call(uintptr(device.Handle), 0x00403C90)
	if err != nil && returnCode != 0 {
		return err
	}
	return nil
}

func SendShortNoteOffMessage(device MidiOutDevice, note int, velocity int, channel int) error {
	returnCode, _, err := procMidiOutShortMsg.Call(uintptr(device.Handle), 0x00003C90)
	if err != nil && returnCode != 0 {
		return err
	}
	return nil
}
