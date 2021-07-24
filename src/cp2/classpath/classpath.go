package classpath

import (
  "path/filepath"
  "os"
)

type Classpath struct {
  // 最底层的类加载器
  bootClasspath Entry
  // 外挂类加载器
  extClasspath Entry
  // 用户类加载器
  userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath  {
  cp := &Classpath{}
  cp.parseBootAndExtClasspath(jreOption)
  cp.parseUserClasspath(cpOption)
  return cp
}

// 搜索class file
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error)  {
  className = className + ".class"
  if data, entry, err := self.bootClasspath.readClass(className); err == nil {
    return data, entry, err
  }
  if data, entry, err := self.extClasspath.readClass(className); err == nil {
    return data, entry, err
  }
  return self.userClasspath.readClass(className)
}

// 字符串
func (self *Classpath) String() string {
  return self.userClasspath.String()
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string)  {
  jreDir := getJreDir(jreOption)

  // jre/lib/*
  jreLibPath := filepath.Join(jreDir, "lib", "*")
  self.bootClasspath = newWildcardEntry(jreLibPath)

  // jre/lib/ext/*
  jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
  self.extClasspath = newWildcardEntry(jreExtPath)
}

func (self *Classpath) parseUserClasspath(cpOption string)  {
  if cpOption == "" {
    cpOption = "."
  }
  self.userClasspath = newEntry(cpOption)
}

// 获取jre路径
func getJreDir(jreOption string) string  {
  if jreOption != "" && exists(jreOption) {
    return jreOption
  }
  if exists("./jre") {
    return "./jre"
  }
  if jh := os.Getenv("JAVA_HOME"); jh != "" {
    return filepath.Join(jh, "jre")
  }
  panic("Can not find jre folder!")
}

// 判断路径是否存在
func exists(path string) bool {
  if _, err := os.Stat(path); err != nil {
    if os.IsNotExist(err) {
      return false
    }
  }
  return true
}
