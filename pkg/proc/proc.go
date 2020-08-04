package proc

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	idleField        = 3
	procStatLineSize = 256
)

func CPU(re io.Reader) (float64, error) {
	var line, _, errLine = bufio.NewReaderSize(re, procStatLineSize).ReadLine()
	if errLine != nil {
		return -1, fmt.Errorf("scanning stat info: %w", errLine)
	}
	line = bytes.TrimPrefix(line, []byte("cpu  "))
	var total, idle uint64
	for i, field := range strings.Fields(string(line)) {
		var p, _ = strconv.ParseUint(field, 10, 64)
		total += p
		if i == idleField {
			idle = p
		}
	}
	return 1 - float64(idle)/float64(total), nil
}

func Mem(re io.Reader) (float64, error) {
	var total, free, available uint64
	var _, errTotal = fmt.Fscanf(re, "MemTotal:        %d kB\n", &total)
	if errTotal != nil {
		return -1, fmt.Errorf("scanning mem info: MemTotal: %w", errTotal)
	}
	var _, errFree = fmt.Fscanf(re, "MemFree:        %d kB\n", &free)
	if errFree != nil {
		return -1, fmt.Errorf("scanning mem info: MemFree: %w", errFree)
	}
	var _, errMemAvailable = fmt.Fscanf(re, "MemAvailable:        %d kB\n", &available)
	if errFree != nil {
		return -1, fmt.Errorf("scanning mem info: MemFrMemAvailable: %w", errMemAvailable)
	}
	return 1 - float64(free+available)/float64(total), nil
}
