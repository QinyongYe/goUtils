package main

/**
	copy a build, unzip it, and register the components
 */
import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var dontcopy = false
	var dontunzip = false
	var dontregister = false
	var buildno string
	var srcdir string

	var root = &cobra.Command{
		Use:"cpb",
		Short:"copy build, unzip it, and register the components.",
		Run: func(cmd * cobra.Command, args []string) {
						doit(!dontcopy, !dontunzip, !dontregister, buildno, srcdir)
					},
	}

	root.Flags().BoolVar(&dontcopy, "dontcopy", "c", false, "don't copy the binary")
	root.Flags().BoolVar(&dontunzip, "dontunzip", "z", false, "don't unzip the binary")
	root.Flags().BoolVar(&dontregister, "dontregister", "c", false, "don't register the binary")
	root.Flags().StringVar(&buildno, "buildno", "b", "", "build number")
	root.Flags().StringVar(&srcdir, "srcdir", "\\\\was-cc2-tech\\cm_bld1", "the source directory of the build")

}

func doit(copy, unzip, register bool, buildno, srcdir string) {
	// remove old binaries
	error := os.RemoveAll("z:/bin")
	if error != nil {
		panic(error)
	}

	error = os.RemoveAll("z:/lib")
	if error != nil {
		panic(error)
	}

	// copy

	// unzip

	// register

}
