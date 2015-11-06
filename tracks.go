package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "errors"
)

type Track struct {
    path string
    logs string
    values []string
    db *bqDataset
    count int
}

func (track *Track) setPath(path string) error {
    if len(path) == 0 {
        return errors.New("tracks path can't be empty")
    }
    track.path = path
    return nil
}

func (track *Track) getPath() string {
    return track.path
} 

func (track *Track) setLogs(path string) error {
    if len(path) == 0 {
        return errors.New("tracks path can't be empty")
    }
    track.logs = path
    return nil
} 

func (track *Track) getLogs() string {
    return track.logs
} 

func (track *Track) setDatabase(db *bqDataset) error {
    track.db = db
    return nil
} 

func (track *Track) load() error {
    if len(track.path) == 0 {
        return errors.New("tracks path is not defined")
    }

    lines, err := readLines(track.path)
    if err != nil {
        return err
    }
    track.values = lines

    track.setSize(len(lines))

    return nil
}

func (track *Track) add(new_tracks []string) error {
    track.values = append(track.values, new_tracks...)

    return track.save(new_tracks)
}

func (track *Track) save(new_tracks []string) error {
    if len(track.path) == 0 {
        return errors.New("tracks path is not defined")
    }
    return writeLines(new_tracks, track.getPath())
}

func (track *Track) ids() []string {
    return track.values
}

func (track *Track) size() int {
    return track.count
}

func (track *Track) setSize(newSize int) {
    track.count = newSize 
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
    file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0666)
    if err != nil {
        return err
    }
    defer file.Close()

    w := bufio.NewWriter(file)
    for _, line := range lines {
        fmt.Fprintln(w, line)
    }
    return w.Flush()
}

func InSlice(a int, list []string) bool {
    for _, b := range list {
        if b == strconv.Itoa(a) {
            return true
        }
    }
    return false
}

/*
func copyTracks() {
    lines, err := readLines("tracks.csv")
    if err != nil {
        log.Fatalf("readLines: %s", err)
    }
    for i, line := range lines {
        fmt.Println(i, line)
    }

    if err := writeLines(lines, "foo.out.txt"); err != nil {
        log.Fatalf("writeLines: %s", err)
    }
}
*/