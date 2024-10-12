package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)


var (
	path = flag.String("file", "measurements.txt", "Measurement file path")
	cpuprofile = flag.Bool("cpu", false, "write cpu profile to `file`")
	memprofile = flag.Bool("mem", false, "write memory profile to `file`")
)

func main() {
	flag.Parse()

	if *cpuprofile {
		f, err := os.Create("./profiles/cpu.prof")
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}



    file, err := os.Open(*path)

    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

	_ = solve(file)


	if *memprofile  {
		f, err := os.Create("./profiles/mem.prof" )
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close()
		runtime.GC()
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
