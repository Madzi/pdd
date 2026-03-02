package cmd

import (
    "os"
)

func Execute() {
    if len(os.Args) < 2 {
        println("Usage: pdd <scan|list>")
        return
    }

    switch os.Args[1] {
    case "scan":
        Scan()
    case "list":
        List()
    default:
        println("Unknown command")
    }
}
