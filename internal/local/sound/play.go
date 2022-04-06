package sound

import (
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"

	"github.com/maxb-odessa/sconf"
	"github.com/maxb-odessa/slog"
)

type soundObj struct {
	confName string
	samples  *beep.Resampler
}

var soundMap = map[int]*soundObj{
	ALARM: {confName: "file alarm", samples: nil},
	BEEP:  {confName: "file beep", samples: nil},
	CLICK: {confName: "file click", samples: nil},
	DROP:  {confName: "file drop", samples: nil},
	NOTE:  {confName: "file note", samples: nil},
	TING:  {confName: "file ting", samples: nil},
	TONE:  {confName: "file tone", samples: nil},
	TWIT:  {confName: "file twit", samples: nil},
	WARN:  {confName: "file warn", samples: nil},
	WARP:  {confName: "file warp", samples: nil},
}

var playSampleRate beep.SampleRate = 44100

func (h *handler) init() error {

	// init speaker with local samplerate
	speaker.Init(playSampleRate, playSampleRate.N(time.Second/5))

	// load all sounds
	for _, so := range soundMap {

		fname, err := sconf.Str("local sound", so.confName)
		if err != nil {
			return err
		}

		fh, err := os.Open(os.ExpandEnv(fname))
		if err != nil {
			return err
		}

		if streamer, format, err := wav.Decode(fh); err != nil {
			return err
		} else {
			so.samples = beep.Resample(5, format.SampleRate, playSampleRate, streamer)
			slog.Debug(1, "%s: loaded '%s' from '%s'", h.endpoint, so.confName, fname)
		}

	}

	return nil
}

func (h *handler) play(s *Track) {

	so, ok := soundMap[s.Id]
	if !ok {
		slog.Err("%s: sound id %d not defined", h.endpoint, s.Id)
		return
	}

	// TODO: play in go, play N times, start/stop

	//		done := make(chan bool)
	speaker.Play(beep.Seq(so.samples, beep.Callback(func() {
		//			done <- true
	})))

	//		<-done
}
