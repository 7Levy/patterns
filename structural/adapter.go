package structural

import "fmt"

type Player interface {
	PlayMusic()
}

//音乐播放器
type MusicPlayer struct {
	Src string
}

func (p MusicPlayer) PlayMusic() {
	fmt.Println("Play music:" + p.Src)
}

func MusicDemo() {
	var player Player = MusicPlayer{Src: "music.mp3"}
	play(player)
}

//游戏声音
type GameSoundPlayer struct {
	Src string
}

func (p GameSoundPlayer) PlaySound() {
	fmt.Println("play sound: " + p.Src)
}

//游戏声音适配器
type GameSoundAdapter struct {
	SoundPlayer GameSoundPlayer
}

func (p GameSoundAdapter) PlayMusic() {
	p.SoundPlayer.PlaySound()
}

func GameSoundDemo() {
	gameSound := GameSoundPlayer{Src: "game.mid"}
	gameAdapter := GameSoundAdapter{SoundPlayer: gameSound}
	play(gameAdapter)
}

func play(player Player) {
	player.PlayMusic()
}
