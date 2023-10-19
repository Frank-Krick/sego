package main

import (
	"fmt"
	"github.com/Frank-Krick/sego/midi/windows"
	"time"
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

	fmt.Println("Opening midi out device", devices[2].Name)
	var device windows.MidiOutDevice
	device, err = windows.OpenMidiOutDevice(devices[2])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Sending note on message")
	err = windows.SendShortNoteOnMessage(device, 0, 0, 0)
	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(10 * time.Second)

	fmt.Println("Sending note off message")
	err = windows.SendShortNoteOffMessage(device, 0, 0, 0)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Closing midi out device")
	err = windows.CloseMidiOutDevice(device)
	if err != nil {
		fmt.Println(err)
	}
}
