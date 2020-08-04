package proc_test

import (
	"strings"
	"testing"

	"github.com/ninedraft/sshop/pkg/proc"
)

func TestCPU(test *testing.T) {
	const data = `cpu  2264588 534 474896 17698398 32236 125742 33127 0 0 0
cpu0 296434 100 69628 3907800 7259 72543 12087 0 0 0
cpu1 235324 66 46479 2004908 2393 9857 3861 0 0 0
cpu2 302543 92 69166 1942932 6160 13032 3672 0 0 0
cpu3 218713 46 50460 2023258 2666 4962 2782 0 0 0
cpu4 334544 71 72321 1923204 4607 8159 3407 0 0 0
cpu5 261727 50 46269 1995889 2387 4667 2049 0 0 0
cpu6 348116 63 73071 1910177 4461 7865 3271 0 0 0
cpu7 267184 44 47499 1990227 2300 4653 1994 0 0 0
intr 262625936 40 50517 0 0 0 0 40511974 598384 1 218 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1067000 43807 25184 13126 24683 14721 21916 16638 26190 14161 0 0 0 0 10 0 813 0 0 0 0 0 0 0 0 596 0 0 0 0 0 0 0 598384 0 0 0 0 0 0 0 984493 13121138 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
ctxt 440348557
btime 1596479897
processes 239713 
procs_running 1
procs_blocked 1
softirq 104648406 737 35850156 58 1267689 62852 14 50862 28162178 1108 39252752
`
	var p, err = proc.CPU(strings.NewReader(data))
	if err != nil {
		test.Fatal(err)
	}
	test.Logf("CPU usage: %f", p)
}

func TestMem(test *testing.T) {
	const data = `MemTotal:        6089604 kB
MemFree:         1747304 kB
MemAvailable:    2721344 kB
Buffers:          166328 kB
Cached:           972676 kB
SwapCached:            0 kB
Active:          3182556 kB
Inactive:         494604 kB
Active(anon):    2416312 kB
Inactive(anon):   124968 kB
Active(file):     766244 kB
Inactive(file):   369636 kB
Unevictable:          96 kB
Mlocked:              96 kB
SwapTotal:             0 kB
SwapFree:              0 kB
Dirty:               184 kB
Writeback:             0 kB
AnonPages:       2538260 kB
Mapped:           404012 kB
Shmem:            147424 kB
KReclaimable:     104432 kB
Slab:             253608 kB
SReclaimable:     104432 kB
SUnreclaim:       149176 kB
KernelStack:       12832 kB
PageTables:        28548 kB
NFS_Unstable:          0 kB
Bounce:                0 kB
WritebackTmp:          0 kB
CommitLimit:     3044800 kB
Committed_AS:    8091128 kB
VmallocTotal:   34359738367 kB
VmallocUsed:       35040 kB
VmallocChunk:          0 kB
Percpu:            15104 kB
HardwareCorrupted:     0 kB
AnonHugePages:         0 kB
ShmemHugePages:        0 kB
ShmemPmdMapped:        0 kB
FileHugePages:         0 kB
FilePmdMapped:         0 kB
HugePages_Total:       0
HugePages_Free:        0
HugePages_Rsvd:        0
HugePages_Surp:        0
Hugepagesize:       2048 kB
Hugetlb:               0 kB
DirectMap4k:      998420 kB
DirectMap2M:     5230592 kB
DirectMap1G:     1048576 kB
`
	var p, err = proc.Mem(strings.NewReader(data))
	if err != nil {
		test.Fatal(err)
	}
	test.Logf("MEM usage: %f", p)
}
