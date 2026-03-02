package utils

import "os"

func IsTextFile(path string) bool {
    f, err := os.Open(path)
    if err != nil {
        return false
    }
    defer f.Close()

    buf := make([]byte, 8000)
    n, _ := f.Read(buf)

    for _, b := range buf[:n] {
        if b == 0 {
            return false
        }
    }

    return true
}
