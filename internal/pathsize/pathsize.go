package pathsize

import "os"

func GetSize(path string) int64 {
    info, err := os.Lstat(path)
    if err != nil {
        return 0
    }

    if !info.IsDir() {
        return info.Size()
    }

    entries, err := os.ReadDir(path)
    if err != nil {
        return 0
    }

    var total int64 = 0

    for _, entry := range entries {
        if entry.IsDir() {
            continue
        }

        fileInfo, err := entry.Info()
        if err != nil {
            continue
        }

        total += fileInfo.Size()
    }

    return total
}



