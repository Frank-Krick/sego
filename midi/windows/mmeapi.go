package windows

import (
	golang_windows "golang.org/x/sys/windows"
)

var libMultimedia = golang_windows.NewLazyDLL("Winmm.dll")
var procMidiOutGetNumDevs = libMultimedia.NewProc("midiOutGetNumDevs")
var procMidiOutGetDevCaps = libMultimedia.NewProc("midiOutGetDevCapsA")
var procMidiOutOpen = libMultimedia.NewProc("midiOutOpen")
var procMidiOutClose = libMultimedia.NewProc("midiOutClose")
var procMidiOutShortMsg = libMultimedia.NewProc("midiOutShortMsg")
