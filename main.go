package main
import (
	"fmt"
)

func parse(track Track) []string  {

	var request AppRequest

	request.find()
	fmt.Println("Parsed apps: ", request.size())

    request.filter(track.ids())
    fmt.Println("New apps: ", request.size())

    request.save()

    return request.getTrackIds()
}

func main() {

	var track Track
	err := track.setPath("tracks.csv")
    if err != nil {
    	fmt.Println(err)
        return
    }

	err = track.load()
	if err != nil {
    	fmt.Println(err)
        return
    }
    fmt.Println("Total apps: ", track.size())

    new_tracks := parse(track)
    err = track.add(new_tracks)
	if err != nil {
    	fmt.Println(err)
        return
    }
}