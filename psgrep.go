package main

import (
        "fmt"
        "os/exec"
        "sort"
        "strconv"
        "strings"
)

type Process struct {
        PID       int
        CPU       float64
        Memory    float64
        Command   string
}

func main() {
        // Get the list of processes
        ps, _ := exec.Command("ps", "aux").Output()
        lines := strings.Split(string(ps), "\n")

        // Parse the processes
        var processes []*Process
        for _, line := range lines[1:] {
                if line == "" {
                        continue
                }
                fields := strings.Fields(line)
                pid, _ := strconv.Atoi(fields[1])
                cpu, _ := strconv.ParseFloat(fields[2], 64)
                memory, _ := strconv.ParseFloat(fields[3], 64)
                command := fields[10]
                processes = append(processes, &Process{
                        PID:       pid,
                        CPU:       cpu,
                        Memory:    memory,
                        Command:   command,
                })
        }

        // Sort the processes by CPU usage
        sort.Slice(processes, func(i, j int) bool {
                return processes[i].CPU > processes[j].CPU
        })

        // Display the top 5 processes by CPU usage
        for i, p := range processes[:5] {
                fmt.Printf("%d. %d %.1f%% %.1fMB %s\n", i+1, p.PID, p.CPU, p.Memory, p.Command)
        }
}
