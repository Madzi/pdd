package scanner

import (
    "io/fs"
    "path/filepath"
)

func ScanFiles(root string, fn func(path string) error) error {
    return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            return err
        }

        if d.IsDir() {
            if d.Name() == ".git" || d.Name() == "node_modules" {
                return filepath.SkipDir
            }
            return nil
        }

        return fn(path)
    })
}
