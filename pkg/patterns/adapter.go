package patterns

import (
	"fmt"
)

type IProcess interface {
	process()
}

type Adapter struct {
	adaptee Adaptee
}

func (adapter Adapter) process() {
	fmt.Println("Adapter process")
	adapter.adaptee.convert()
}
type Adaptee struct {
	adapterType int
}

func (adapter Adaptee) convert() {
	fmt.Println("Adaptee convert")
}

func AdapterFun() {
	var processor IProcess = Adapter{adaptee:Adaptee{1}}
	processor.process()
}
		//////////////////////// 
		// MediaPlayer example //
		////////////////////////
		
// MediaPlayer is the target interface for playing audio
type MediaPlayer interface {
	PlayAudio(fileName string)
}

// AudioPlayer is a concrete implementation of MediaPlayer
type AudioPlayer struct{}

func (a *AudioPlayer) PlayAudio(fileName string) {
    fmt.Printf("Playing audio file: %s\n", fileName)
}

// VideoPlayer is the adaptee that plays video files
type VideoPlayer struct{}

func(v *VideoPlayer) PlayVideo(fileName string) {
	fmt.Printf("Playing video file : %s\n",fileName)
}

// VideoAdapter adapts a VideoPlayer to a MediaPlayer
type VideoAdapter struct {
	VideoPlayer *VideoPlayer
}
func (v *VideoAdapter) PlayAudio(fileName string) {
	// Adapter adapts the PlayAudio method to call PplayVideo instead
	v.VideoPlayer.PlayVideo(fileName)
}

func MediaPlayerFunc() {
	// Audio player directly using the MediaPlayer inerface
	audioPlayer := new(AudioPlayer)
	audioPlayer.PlayAudio("song.mp3")

	// Use the adapter to play video files as audio fils
	videoPlayer := new(VideoPlayer)
	videoAdapter := &VideoAdapter{VideoPlayer:videoPlayer}

	// The client sees the PlayAudio method ,but it's actually playing a video
	videoAdapter.PlayAudio("video.mp4")
}