package runners

import (
	"math/rand"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/scootdev/scoot/runner"
	log "github.com/Sirupsen/logrus"
)

// chaos.go: impl that introduces errors (for testing)

// Creates a new Chaos Runner
func NewChaosRunner(delegate runner.Service) *ChaosRunner {
	runner := &ChaosRunner{del: delegate}
	log.Println("************** runner definition")
	runnerDesc := spew.Sdump(runner)
	log.Println(runnerDesc)
	return runner
}

// ChaosRunner implements Runner by calling to a delegate runner in the happy path,
// but delaying a random time between 0 and MaxDelay, or returning an error
type ChaosRunner struct {
	del      runner.Service
	mu       sync.Mutex
	maxDelay time.Duration
	err      error
}

// Chaos Controls

// SetError sets the error to return on calls to the Runner
func (r *ChaosRunner) SetError(err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.err = err
}

// SetDelay sets the max delay to delay; the actual delay will be randomly selected up to delay
func (r *ChaosRunner) SetDelay(delay time.Duration) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.maxDelay = delay
}

// Chaos implementations
func (r *ChaosRunner) delay() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.maxDelay != 0 {
		time.Sleep(time.Duration(rand.Int63n(int64(r.maxDelay))))
	}
	return r.err
}

// Implement Runner
func (r *ChaosRunner) Run(cmd *runner.Command) (runner.RunStatus, error) {
	err := r.delay()
	if err != nil {
		return runner.RunStatus{}, err
	}
	return r.del.Run(cmd)
}

func (r *ChaosRunner) Query(q runner.Query, w runner.Wait) ([]runner.RunStatus, error) {
	err := r.delay()
	if err != nil {
		return nil, err
	}
	return r.del.Query(q, w)
}

func (r *ChaosRunner) QueryNow(q runner.Query) ([]runner.RunStatus, error) {
	err := r.delay()
	if err != nil {
		return nil, err
	}
	return r.del.QueryNow(q)
}

func (r *ChaosRunner) Status(run runner.RunID) (runner.RunStatus, error) {
	err := r.delay()
	if err != nil {
		return runner.RunStatus{}, err
	}
	return r.del.Status(run)
}

func (r *ChaosRunner) StatusAll() ([]runner.RunStatus, error) {
	err := r.delay()
	if err != nil {
		return nil, err
	}
	return r.del.StatusAll()
}

func (r *ChaosRunner) Abort(run runner.RunID) (runner.RunStatus, error) {
	err := r.delay()
	if err != nil {
		return runner.RunStatus{}, err
	}
	return r.del.Abort(run)
}

func (r *ChaosRunner) Erase(run runner.RunID) error {
	err := r.delay()
	if err != nil {
		return err
	}
	return r.Erase(run)
}
