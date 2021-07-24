package main

import (
  "cp2/classpath"
  "cp2/cmd" // 导入自定义模块
  "fmt"
  "strings"
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
 cp := classpath.Parse(cmd.XjreOption, cmd.CpOption)
 fmt.Printf("classpath:%v class:%v args:%v\n", cp, cmd.Class, cmd.Args)

 className := strings.Replace(cmd.Class, ".", "/", -1)
 classData, _, err := cp.ReadClass(className)
 if err != nil {
   fmt.Printf("Could not find or load main class %s\n", cmd.Class)
   return
 }

 fmt.Printf("class data:%v\n", classData)
}
