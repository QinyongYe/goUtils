package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"encoding/hex"
)

func main() {
	var root = &cobra.Command{
		Use: "mstrguid",
//		Run: func(cmd * cobra.Command, args []string) {
//			mstr2Guids(args)
//		},
	}

	var sub = &cobra.Command{
		Use: "c",
		Run: func(cmd * cobra.Command, args []string) {
			mstr2Guids(args)
		},
	}
	root.AddCommand(sub)
	root.Execute()
}

func mstr2Guids(ids []string) {
	for _, id := range ids {
	  mstr2Guid(id)

	}
}

type Guid struct {
	data1 uint32
	data2 uint16
	data3 uint16
	data4 [8]byte
}

func mstr2Guid(id string) {
	if(len(id) != 32) {
		panic(fmt.Errorf("%v is not a valid mstr guid", id))
	}

	bytes, err := hex.DecodeString(id)
	if(err != nil) {
		panic(err)
	}

	var guid = Guid{}
	guid.data1 = (uint32(bytes[0]) << 24) | (uint32(bytes[1]) << 16) | (uint32(bytes[2]) << 8) | (uint32(bytes[3]))
	guid.data2 = (uint16(bytes[6]) << 8) | (uint16(bytes[7]))
	guid.data3 = (uint16(bytes[4]) << 8) | (uint16(bytes[5]))
	guid.data4[0] = bytes[8+3]
	guid.data4[1] = bytes[8+2]
	guid.data4[2] = bytes[8+1]
	guid.data4[3] = bytes[8+0]
	guid.data4[4] = bytes[8+7]
	guid.data4[5] = bytes[8+6]
	guid.data4[6] = bytes[8+5]
	guid.data4[7] = bytes[8+4]
	fmt.Printf("%x", guid)
}
