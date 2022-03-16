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
	var event string
	if e, ok := entry["event"]; ok {
		event = e.(string)
	}

	if _, ok := entry["timestamp"]; !ok {
		return
	}

	// something may go wrong in event handler...
	defer func() {
		if r := recover(); r != nil {
			slog.Warn("parse journal failed, recoverd in %v", r)
		}
	}()

	// find event handler and call it
	switch event {
	case "Docked":
		h.jDocked(entry)
	case "DockingCancelled":
		h.jDockingCancelled(entry)
	case "DockingGranted":
		h.jDockingGranted(entry)
	case "Fileheader":
		h.jFileheader(entry)
	case "FSDJump":
		h.jFSDJump(entry)
	case "FSDTarget":
		h.jFSDTarget(entry)
	case "FSSAllBodiesFound":
		h.jFSSAllBodiesFound(entry)
	case "FSSBodySignals":
		h.jFSSBodySignals(entry)
	case "FSSDiscoveryScan":
		h.jFSSDiscoveryScan(entry)
	case "FSSSignalDiscovered":
		h.jFSSSignalDiscovered(entry)
	case "LeaveBody":
		h.jLeaveBody(entry)
	case "Liftoff":
		h.jLiftoff(entry)
	case "MaterialCollected":
		h.jMaterialCollected(entry)
	case "NavRoute":
		h.jNavRoute(entry)
	case "SAAScanComplete":
		h.jSAAScanComplete(entry)
	case "SAASignalsFound":
		h.jSAASignalsFound(entry)
	case "Scan":
		h.jScan(entry)
	case "ScanBaryCentre":
		h.jScanBaryCentre(entry)
	case "Scanned":
		h.jScanned(entry)
	case "ScanOrganic":
		h.jScanOrganic(entry)
	case "Shutdown":
		h.jShutdown(entry)
	case "StartJump":
		h.jStartJump(entry)
	case "SupercruiseEntry":
		h.jSupercruiseEntry(entry)
	case "SupercruiseExit":
		h.jSupercruiseExit(entry)
	case "Touchdown":
		h.jTouchdown(entry)
	case "Undocked":
		h.jUndocked(entry)
	default:
		slog.Debug(5, "journal: ignoring unhandled event '%s'", event)
	}

}
