package file

import (
	"encoding/json"

	pb "github.com/maxb-odessa/gamenode/pkg/gamenodepb"
	"github.com/maxb-odessa/slog"
)

type state struct {
	currentStar string
}

func (h *handler) processJournalMsg(ev *pb.FileEvent) {

	line := ev.GetLine().GetLine()
	slog.Debug(9, "get line: '%+v", line)

	var entry map[string]interface{}

	if err := json.Unmarshal([]byte(line), &entry); err != nil {
		slog.Err("journal entry json unmarshal failed: %s", err)
		return
	}

	// every journal entry must have 'timestamp' and 'event' fields
	if _, ok := entry["timestamp"]; !ok {
		return
	}

	var event string
	if e, ok := entry["event"]; ok {
		event = e.(string)
	}

	// handler event
	h.handleEvent(event, entry)
}
