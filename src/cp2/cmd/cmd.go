package cmd

import (
  "flag"
  "fmt"
  "os"
)

// cmd集合
type Cmd struct {
  HelpFlag    bool
  VersionFlag bool
  CpOption    string
  XjreOption string
  Class string
  Args  []string
}

func ParseCmd()  *Cmd {
  cmd := &Cmd{}
  flag.Usage = printUsage
  flag.BoolVar(&cmd.HelpFlag, "help", false, "print help message")
  flag.BoolVar(&cmd.HelpFlag, "?", false, "print help message")
  flag.BoolVar(&cmd.VersionFlag, "verson", false, "print version ")
  flag.StringVar(&cmd.CpOption, "classpath", "", "classpath")
  flag.StringVar(&cmd.CpOption, "cp", "", "classpath")
  flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
  flag.Parse()

  args := flag.Args()
  if len(args) >0 {
    cmd.Class = args[0]
    cmd.Args = args[1:]
  }

  return cmd
}

func printUsage() {
  fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
