package classpath

import (
  "os"
  "path/filepath"
  "strings"
)

func newWildcardEntry(path string) CompositeEntry {
  baseDir := string(path[len(path) - 1]) // remove *
  var compositeEntry []Entry
  // 找到jar文件，并返回跳过子目录
  walkFn := func(path string, info os.FileInfo, err error) error {
    if err != nil {
      return err
    }

    if info.IsDir() && path != baseDir {
      return filepath.SkipDir
    }
    if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
      jarEntry := newZipEntry(path)
      compositeEntry = append(compositeEntry, jarEntry)
    }
    return nil
  }
  filepath.Walk(baseDir, walkFn)
  return compositeEntry
}

