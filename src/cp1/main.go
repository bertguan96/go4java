package main

import (
  "cp1/cmd" // 导入自定义模块
  "fmt"
)

func main()  {
 cmd := cmd.ParseCmd()
 if cmd.VersionFlag {
   fmt.Println("version 0.0.1")
 } else {
   startJVM(cmd)
 }
}

func startJVM(cmd *cmd.Cmd)  {
 fmt.Printf("classpath: %s class: %s args: %v\n", cmd.CpOption, cmd.Class, cmd.Args)
}
