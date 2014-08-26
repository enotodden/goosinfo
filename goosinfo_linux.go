package goosinfo

import (
    "io/ioutil"
    "strings"
    "os"
    "fmt"
    "syscall"
)

type OSRelease struct {
    Name string
    PrettyName string
    Version string
    VersionID string
    ID string
}

func uname_chars_to_string(ca [65]int8) string {
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

func parse_os_release (contents string) (*OSRelease, error) {
    osr := OSRelease{}
    contents = strings.TrimSpace(contents)
    lines := strings.Split(contents, "\n")
    for _, line := range(lines) {
        spl := strings.Split(line, "=")
        k := strings.ToUpper(spl[0])
        v := strings.Trim(spl[1], "\"")
        if k == "NAME" {
            osr.Name = v
        } else if k == "PRETTY_NAME" {
            osr.PrettyName = v
        } else if k == "VERSION" {
            osr.Version = v
        } else if k == "VERSION_ID" {
            osr.VersionID = v
        } else if k == "ID" {
            osr.ID = v
        }
    }
    return &osr, nil
}

func read_os_release () (*OSRelease, error) {
    file, err := os.Open("/etc/os-release")
    if err != nil {
        return nil, err
    }
    c, err := ioutil.ReadAll(file)
    if err != nil {
        return nil, err
    }
    file.Close()
    return parse_os_release(string(c))
}

func get_kernel_version_uname() (string, error) {
    uts := syscall.Utsname{}
    err := syscall.Uname(&uts)
    if err != nil {
        return "", fmt.Errorf("Could not get Linux version")
    }
    y, _ := read_os_release()
    fmt.Println(y)
    return uname_chars_to_string(uts.Release), nil
}

func GetOSVersion() (string, error) {
    osr, err := read_os_release()
    if err != nil {
        kver, err := get_kernel_version_uname()
        if err != nil {
            return "", err
        }
        return kver, nil
    }
    return osr.Version, nil
}
