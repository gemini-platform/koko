package model

type Platform struct {
	Name     string                 `json:"name"`
	BaseOs   string                 `json:"base"`
	Charset  string                 `json:"charset"`
	MetaData map[string]interface{} `json:"meta"`
}

const (
	PlatformLinux       = "Linux"
	PlatformUnix        = "Unix"
	PlatformMacOS       = "MacOS"
	PlatformBSD         = "BSD"
	PlatformWindows     = "Windows"
	PlatformWindows2016 = "Windows2016"
	PlatformOther       = "Other"
	PlatformWindowsRDP  = "Windows-RDP"
	PlatformWindowsTLS  = "Windows-TLS"
	PlatformAIX         = "AIX"
)