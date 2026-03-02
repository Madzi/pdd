package model

type Puzzle struct {
    Type        string `xml:"type"`
    Timestamp   string `xml:"timestamp"`
    Description string `xml:"description"`
    File        string `xml:"file"`
    Status      string `xml:"status"`
}

type Puzzles struct {
    Items []Puzzle `xml:"puzzle"`
}
