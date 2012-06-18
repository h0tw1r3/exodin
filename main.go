// Copyright 2012 Jeffrey Clark <h0tw1r3@gmail.com>. All rights reserved.
// License GPLv3+: GNU GPL version 3 or later <http://gnu.org/licenses.gpl.html>.
// This is free software: you are free to change and redistribute it.
// There is NO WARRANTY, to the extent permitted by law.

package main

import (
	"flag"
	"log"
	"os"
	"io"
	"fmt"
	"exodin/tar"
)

const (
	APP_VERSION = "0.1"
	APP_NAME = "exodin"
	APP_AUTHOR = "Jeffrey Clark <h0tw1r3@gmail.com>"
	APP_LEGAL = "Copyright (C) 2012 Jeffrey Clark\nLicense GPLv3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>.\nThis is free software: you are free to change and redistribute it.\nThere is NO WARRANTY, to the extent permitted by law.\n"
)

type cmdFlag struct {
	v *bool
	x *bool
	filename string
}

func main() {
	curFlags := new(cmdFlag)
	if ! procFlags(curFlags) {
		os.Exit(64)
	}
	if (*curFlags.v) {
		fmt.Printf("%s %s\n%s\nWritten by %s\n", APP_NAME, APP_VERSION, APP_LEGAL, APP_AUTHOR)
		os.Exit(0)
	}
	file, err := os.Open(curFlags.filename)
	if (err == nil) {
		r := io.ReadSeeker(file)
		tr := tar.NewReader(r)
		for {
			hdr, err := tr.Next()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
				break
			}
			fmt.Printf("%s %s/%s %9d %s %s\n", os.FileMode(hdr.Mode), hdr.Uname, hdr.Gname, hdr.Size, hdr.ModTime.Format("2006-01-02 15:04"), hdr.Name)
			if *curFlags.x {
				outfile, err := os.Create(hdr.Name)
				if err != nil {
					log.Fatal("Error creating file: ", err)
				}
				io.Copy(outfile, tr)
				outfile.Close()
			}
		}
	} else {
		fmt.Printf("Error: File not found - %s\n", curFlags.filename)
		os.Exit(74)
	}
}

func procFlags(f *cmdFlag) bool {
	f.v = flag.Bool("v", false, "")
	f.x = flag.Bool("x", false, "")
	flag.Usage = func() {
		fmt.Printf("Usage:\t%s [FLAGS] FILENAME\nFlags:\n\t-x\textract files\n\t-v\toutput version and exit\n",os.Args[0])
		fmt.Printf("\n%s", APP_LEGAL)
	}
	flag.Parse()
	if len(os.Args)==1 {
		flag.Usage()
		return false
	}
	if !*f.v && flag.NArg()!=1 {
		fmt.Printf("Error: unexpected command line argument(s)\n")
		flag.Usage()
		return false
	} else {
		f.filename = flag.Arg(0)
	}
	return true
}
