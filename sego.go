package main

import (
	"fmt"
	"github.com/Frank-Krick/sego/midi/windows"
)

func main() {
	fmt.Println("Reading all midi out devices")
	devices, err := windows.ListMidiOutDevices()
	if err != nil {
		return
	}

	for _, device := range devices {
		fmt.Println("Midi device:", device.Name)
	}

	fmt.Println("Opening midi out device")
	var device windows.MidiOutDevice
	device, err = windows.OpenMidiOutDevice()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Closing midi out device")
	err = windows.CloseMidiOutDevice(device)
	if err != nil {
		fmt.Println(err)
	}
}
