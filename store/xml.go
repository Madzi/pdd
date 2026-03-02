package store

import (
    "encoding/xml"
    "io/ioutil"
    "os"
    "pdd/model"
)

const FileName = ".puzzles.xml"

func Load() (model.Puzzles, error) {
    var puzzles model.Puzzles

    if _, err := os.Stat(FileName); os.IsNotExist(err) {
        return puzzles, nil
    }

    data, err := ioutil.ReadFile(FileName)
    if err != nil {
        return puzzles, err
    }

    err = xml.Unmarshal(data, &puzzles)
    return puzzles, err
}

func Save(puzzles model.Puzzles) error {
    data, err := xml.MarshalIndent(puzzles, "", "  ")
    if err != nil {
        return err
    }
    return ioutil.WriteFile(FileName, data, 0644)
}
