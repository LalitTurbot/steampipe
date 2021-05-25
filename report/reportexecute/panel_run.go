package reportexecute

import (
	"context"

	"github.com/turbot/steampipe/control/controlexecute"

	typehelpers "github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe/db"
	"github.com/turbot/steampipe/steampipeconfig/modconfig"
)

type PanelRunStatus uint32

const (
	PanelRunReady PanelRunStatus = 1 << iota
	PanelRunStarted
	PanelRunComplete
	PanelRunError
)

// PanelRun is a struct representing a  a panel run - will contain one or more result items (i.e. for one or more resources)
type PanelRun struct {
	Error error `json:"-"`

	// the result
	Text   string          `json:"text,omitempty"`
	Title  string          `json:"title,omitempty"`
	Width  int             `json:"width,omitempty"`
	Source string          `json:"source"`
	Data   [][]interface{} `json:"data,omitempty"`

	// children
	PanelRuns  []*PanelRun  `json:"panels,omitempty"`
	ReportRuns []*ReportRun `json:"reports,omitempty"`

	runStatus     PanelRunStatus
	executionTree *ReportExecutionTree
}

func NewPanelRun(panel *modconfig.Panel, executionTree *controlexecute.ExecutionTree) *PanelRun {
	return &PanelRun{
		Title: typehelpers.SafeString(panel.Title),
		// TODO OTHER STUFF
		runStatus: PanelRunReady,
	}
}

func (r *PanelRun) Start(ctx context.Context, client *db.Client) {

}

func (r *PanelRun) SetError(err error) {
	r.Error = err
	r.runStatus = PanelRunError
}
