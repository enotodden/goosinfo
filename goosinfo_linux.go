package goosinfo

import (
    "fmt"
    "syscall"
)

func UnameCharsToString(ca [65]int8) string {
    s := make([]byte, len(ca))
    var lens int
    for ; lens < len(ca); lens++ {
        if ca[lens] == 0 {
            break
        }
        s[lens] = uint8(ca[lens])
    }
    return string(s[0:lens])
}

func GetOSVersion() (string, error) {
    //FIXME: Read /etc/os-release or similar and then fall back on kernel
    //       version
    uts := syscall.Utsname{}
    err := syscall.Uname(&uts)
    if err != nil {
        return "", fmt.Errorf("Could not get Linux version")
    }
    return UnameCharsToString(uts.Release), nil
}
