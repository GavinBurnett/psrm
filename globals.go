package main

const (
	PLATFORM = "Pi/Arm"

	PI_ARM_PLATFORM  = "Pi/Arm"
	LINUX_PLATFORM   = "Linux"
	WINDOWS_PLATFORM = "Windows"
	MACOS_PLATFORM   = "MacOS"

	SECTIONS = 10
)

type SECTION_TYPE int

const (
	RANDOM_SPACE      SECTION_TYPE = 1
	OVERWRITE_SECTION SECTION_TYPE = 2
)

var DEBUG bool = false
