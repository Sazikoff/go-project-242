package pathsize

import "os"

func GetSize(path string) (int64, error) {
    info, err := os.Lstat(path)
    if err != nil {
        return 0, err
    }

    if !info.IsDir() {
        return info.Size(), nil
    }

    entries, err := os.ReadDir(path)
    if err != nil {
        return 0, err
    }

    var total int64

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

    return total, nil
}



