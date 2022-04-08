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
	buffer   *beep.Buffer
}

var soundMap = map[int]*soundObj{
	ALARM: {confName: "file alarm"},
	BEEP:  {confName: "file beep"},
	CLICK: {confName: "file click"},
	DROP:  {confName: "file drop"},
	NOTE:  {confName: "file note"},
	TING:  {confName: "file ting"},
	TONE:  {confName: "file tone"},
	TWIT:  {confName: "file twit"},
	WARN:  {confName: "file warn"},
	WARP:  {confName: "file warp"},
}

var playSampleRate beep.SampleRate = 44100

func (h *handler) init() error {

	// init speaker with local samplerate
	speaker.Init(playSampleRate, playSampleRate.N(time.Second/10))

	// load all sounds
	for _, so := range soundMap {

		fname, err := sconf.Str("local sound", so.confName)
		if err != nil {
			// it's ok, this sound is nto configured
			continue
		}

		fh, err := os.Open(os.ExpandEnv(fname))
		if err != nil {
			slog.Warn("%s: error: %s", h.endpoint, err)
			continue
		}

		if streamer, format, err := wav.Decode(fh); err != nil {
			return err
		} else {
			so.buffer = beep.NewBuffer(format)
			so.buffer.Append(streamer)
			streamer.Close()
			slog.Debug(1, "%s: loaded '%s' from '%s'", h.endpoint, so.confName, fname)
		}

	}

	return nil
}

func (h *handler) play(s *Track) {

	//	speaker.Lock()
	//	defer speaker.Unlock()

	so, ok := soundMap[s.Id]
	if !ok {
		slog.Err("%s: sound id %d not defined", h.endpoint, s.Id)
		return
	}

	if so.buffer == nil {
		slog.Err("%s: sound id %d not loaded", h.endpoint, s.Id)
		return
	}

	// TODO: start/stop

	// repeat: <0, 0, 1 = repeat one time, > 1 = repeat N times
	rTimes := s.Repeat
	if rTimes < 2 {
		rTimes = 1
	}

	track := so.buffer.Streamer(0, so.buffer.Len())
	for i := 0; i < rTimes; i++ {
		slog.Debug(9, "%s PLAYING %s", h.endpoint, so.confName)
		speaker.Play(track)
	}

}
