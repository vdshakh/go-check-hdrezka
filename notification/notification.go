package notification

import (
	"fmt"
	"github.com/0xAX/notificator"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/go-check-hdrezka/request"
	"os"
	"time"
)

const soundFile = "notification/media/sound.mp3"

func PlaySound() {
	sound, err := os.Open(soundFile)
	if err != nil {
		fmt.Errorf("unable to open sound file: %v", err)
		os.Exit(1)
	}

	streamer, format, err := mp3.Decode(sound)
	if err != nil {
		fmt.Errorf("error decoding mp3 file: %v", err)
		os.Exit(1)
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)
}

func SendPush(userRequest request.Req) {
	notify := notificator.New(notificator.Options{
		AppName: "HDRezka Updates",
	})

	if err := notify.Push("Updates are here!", userRequest.Title+" series in "+userRequest.Voiceover+" voiceover is released", "", notificator.UR_NORMAL); err != nil {
		fmt.Errorf("error pushing notification: %v", err)
		os.Exit(1)
	}
}
