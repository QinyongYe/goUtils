package main

import (
	"os"
	. "code.google.com/p/go.crypto/ssh"
	"io/ioutil"
	"github.com/spf13/cobra"
	"fmt"
	"strings"
)

var view = "~/qye_unix_view0/"
var projects string
var machines string
var oneflag bool

func main() {
	var cmdBuild = &cobra.Command{
		Use:   "build",
		Short: "build projects",
		Run: func(cmd *cobra.Command, args []string) {
			run(buildScripts)
		},
	}
	cmdBuild.Flags().StringVarP(&projects, "projects", "p", "Common/Search/FacebookSearchProvider,Common/Search/FacebookSearchManager", "the prjects to build")
	cmdBuild.Flags().BoolVarP(&oneflag, "one", "o", true, "only build the project, not dependant projects")
	var cmdClean = &cobra.Command{
		Use:   "clean",
		Short: "clean projects",
		Run: func(cmd *cobra.Command, args []string) {
			run(cleanScripts)
		},
	}
	cmdClean.Flags().StringVarP(&projects, "projects", "p", "Common/Search/FacebookSearchProvider,Common/Search/FacebookSearchManager", "the prjects to build")

	var buildNum string
	var cmdCopy = &cobra.Command{
		Use:   "copy",
		Short: "copy build",
		Run: func(cmd *cobra.Command, args []string) {
			for _, m := range strings.Split(machines, ",") {
				login(m, genExecuteFunc(copyScripts(buildNum)))
			}
		},
	}
	cmdCopy.Flags().StringVarP(&buildNum, "buildnum", "n", "9.4.0000.0097", "the build to copy")

	var rootCmd = &cobra.Command{Use: "remotebuild"}
	rootCmd.PersistentFlags().StringVarP(&machines, "machines", "m", "adric,earth10,RHEL5U8-TS7", "the machines to connect")
	rootCmd.AddCommand(cmdBuild, cmdClean, cmdCopy)
	rootCmd.Execute()
}

func genExecuteFunc(scripts []string) func(*Client) {
	return func(c *Client) {
		// Each ClientConn can support multiple interactive sessions,
		// represented by a Session.
		s, err := c.NewSession()
		if err != nil {
			panic("Failed to create session: " + err.Error())
		}
		defer s.Close()

		s.Stdout = os.Stdout
		s.Stderr = os.Stderr
		var cmd = ""
		for _, v := range scripts {
			cmd += v + ";"
		}
		fmt.Println("script: ", cmd)
		s.Run(cmd)
	}
}


var cleanScripts = func (project string) []string {
	return []string{
		// remove BuildControl.db, or it will complains in linux
				"rm " + view + "BuildScripts/BuildControl/BuildControl.db -f",
				"cd " + view + project, // cd project folder
			view + "BuildScripts/one.pl -one -notest", // one.pl
		"make -f Makefile`uname` clean", // make
	}
}

var copyScripts = func (buildNum string) []string {
	return []string{
				"cd /user4/Builds/" + buildNum + "/DEBUG/BIN",
			"perl copyto.pl " + view,
	}
}

var run = func(genScripts func(string) []string) {
	for _, m := range strings.Split(machines, ",") {
		fmt.Println("machine: ", m)
		for _, p := range strings.Split(projects, ",") {
			fmt.Println("project: ", p)
			var scripts = genScripts(p)
			login(m, genExecuteFunc(scripts))
		}
	}
}

var buildScripts = func (project string) []string {
	one := ""
	if oneflag {
		one = " -one "
	}

	return []string{
		// remove BuildControl.db, or it will complains in linux
				"rm " + view + "BuildScripts/BuildControl/BuildControl.db -f",
				"cd " + view + project, // cd project folder
			view + "BuildScripts/one.pl" + one + " -notest > /dev/null", // one.pl
		"make -j 4 -f Makefile`uname`", // make
	}
}

func login(machine string, cb func (*Client)) {
	// An SSH client is represented with a ClientConn. Currently only
	// the "password" authentication method is supported.
	//
	// To authenticate with the remote server you must pass at least one
	// implementation of AuthMethod via the Auth field in ClientConfig.
	keyfile, err := os.OpenFile("C:\\Users\\qye\\.ssh\\id_rsa", os.O_RDONLY, os.FileMode(0))
	if err != nil {
		panic("Failed to open file: " + err.Error())
	}
	defer keyfile.Close()

	keybytes, err := ioutil.ReadAll(keyfile)
	key, err := ParsePrivateKey(keybytes)
	if err != nil {
		panic("Failed to parse key: " + err.Error())
	}

	config := &ClientConfig{
		User: "qye",
		Auth: []AuthMethod{
			PublicKeys(key),
		},
	}
	client, err := Dial("tcp", machine+":22", config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}

	cb(client)
}

