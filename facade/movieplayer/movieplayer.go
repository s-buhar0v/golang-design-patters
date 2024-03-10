package movieplayer

import "github.com/s-buhar0v/golang-design-patters/facade/devices"

type MoviePlayer struct {
	tv          devices.TV
	videoPlayer devices.VideoPlayer
}

func NewMoviePlayer() *MoviePlayer {
	return &MoviePlayer{}
}

func (m *MoviePlayer) SetDevices(
	tv devices.TV,
	videoPlayer devices.VideoPlayer,
) {
	m.tv = tv
	m.videoPlayer = videoPlayer
}

func (m *MoviePlayer) Start(name string) {
	m.tv.On()
	m.videoPlayer.On()
	m.videoPlayer.InsertDisk(name)
	m.videoPlayer.Play(name)
}

func (m *MoviePlayer) Stop(name string) {
	m.videoPlayer.Stop(name)
	m.videoPlayer.ExtractDisk(name)
	m.videoPlayer.Off()
	m.tv.Off()
}
