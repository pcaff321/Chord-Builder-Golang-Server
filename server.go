package main

import (
	"fmt"
	//"strconv"
	//"math/rand"
	"net/http"
	"strings"
)

var notes = [12]string {"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}

func findNoteIndex(note string) int{
	for i := 0; i < len(notes); i++ {
		if notes[i] == note{
		return i
		}
	}
	return -1
}

func minorChord(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	dirs := strings.Split(r.URL.String(), "/")



	if  ((len(dirs) >= 3) && (len(dirs[2]) > 0)){
	note := strings.ToUpper(string(dirs[2][0]))
	if len(dirs[2]) > 1 {
	note += "#"
	}
	noteIndex := findNoteIndex(note)
	if noteIndex == -1 {
	fmt.Fprintf(w, "ERROR: Please enter an acceptable note")
	return
	}
	fmt.Fprintf(w, "Note Chosen As Root: " + note + "\n")
	chordFormation := note + " " +  notes[(noteIndex + 3) % len(notes)] + " " + notes[(noteIndex + 7) % len(notes)]
	if (len(dirs) > 3) && (strings.ToUpper(dirs[3]) == "SEVENTH") { 
		chordFormation += " " + notes[(noteIndex + 10) % len(notes)]
		fmt.Fprintf(w, "Minor Seventh Chord: " + chordFormation)
		return
	}
	fmt.Fprintf(w, "Minor Chord: " + chordFormation)
	return
	}

	fmt.Fprintf(w, "ERROR: Please enter an acceptable note")
	return
}

func majorChord(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	dirs := strings.Split(r.URL.String(), "/")



	if  ((len(dirs) >= 3) && (len(dirs[2]) > 0)){
	note := strings.ToUpper(string(dirs[2][0]))
	if len(dirs[2]) > 1 {
	note += "#"
	}
	noteIndex := findNoteIndex(note)
	if noteIndex == -1 {
	fmt.Fprintf(w, "ERROR: Please enter an acceptable note")
	return
	}
	fmt.Fprintf(w, "Note Chosen As Root: " + note + "\n")
	chordFormation := note + " " +  notes[(noteIndex + 4) % len(notes)] + " " + notes[(noteIndex + 7) % len(notes)]
	if (len(dirs) > 3) && (strings.ToUpper(dirs[3]) == "SEVENTH") { 
		chordFormation += " " + notes[(noteIndex + 10) % len(notes)]
		fmt.Fprintf(w, "Major Seventh Chord: " + chordFormation)
		return
	}
	fmt.Fprintf(w, "Major Chord: " + chordFormation)
	return
	}

	fmt.Fprintf(w, "ERROR: Please enter an acceptable note")
	return
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome To Musical Chord Calculator \n")
	fmt.Fprintf(w, "Supported Chords: Major, Minor, Major Sevenths, Minor Sevenths \n")
	fmt.Fprintf(w, "Hint: Go to /minor/A to see the notes in Am! \n")
	fmt.Fprintf(w, "Hint: Go to /major/C to see the notes in C Major! \n")
	fmt.Fprintf(w, "Hint: To make it a seventh chord, add /seventh/ at the end of the URL \n")
	return
}



func main() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/minor/",  minorChord)
	http.HandleFunc("/major/",  majorChord)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}