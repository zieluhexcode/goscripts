package main

import (
        "fmt"
        "os/exec"
        "strings"
)

func main() {
        // Get the list of connections
        netstat, _ := exec.Command("netstat", "-tanp").Output()
        lines := strings.Split(string(netstat), "\n")

        // Print the header
        fmt.Println("Proto  Local Address          Foreign Address           Process")

        // Parse the connections
        for _, line := range lines[2:] {
                if line == "" {
                        continue
                }
                fields := strings.Fields(line)
                proto := fields[0]
                localAddr := fields[3]
                foreignAddr := fields[4]
                process := fields[6]
                fmt.Printf("%-6s %-22s %-22s %s\n", proto, localAddr, foreignAddr, process)
        }
}
