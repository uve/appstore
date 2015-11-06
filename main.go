package main
import (
	"fmt"
    "time"
)


func (track *Track) initialize() (error) {
    err := track.setPath("tracks.csv")
    if err != nil {
        return err
    }

    err = track.setLogs("search.log")
    if err != nil {
        return err
    }

    db, err := connectBigQueryDB()
    if err != nil {
        return err
    }
    err = track.setDatabase(db)
    if err != nil {
        return err
    }
    err = track.load()
    if err != nil {
        return err
    }
    return nil
}

func main() {
    time.Local = time.UTC

    var track Track
    err := track.initialize()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Total apps: ", track.size())

    for { // бесконечный цикл
        err = track.getApps()
        if err != nil {
        }
    }
}
