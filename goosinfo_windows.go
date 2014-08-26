package goosinfo

import (
    "fmt"
    "syscall"
)

func GetOSVersion() (string, error) {
    dll := syscall.MustLoadDLL("kernel32.dll")
    p := dll.MustFindProc("GetVersion")
    v, _, _ := p.Call()
    return fmt.Sprintf("%d.%d (%d)", byte(v), uint8(v>>8), uint16(v>>16)), nil
}
