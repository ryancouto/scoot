package main

import (
	"log"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/scootdev/scoot/os/temp"
	"github.com/scootdev/scoot/scootapi/setup"
)

func main() {

	tmp, err := temp.NewTempDir("", "recoveryTest")
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	clusterCmds := setup.NewSignalHandlingCmds(tmp)
	builder := setup.NewGoBuilder(clusterCmds)

	// TODO: actually read this from config
	// This needs to match the SagaLog config "Directory" in
	// ./binaries/scheduler/config/local.local
	sagalogDirName := "sagalog"
	defer cleanup(clusterCmds, sagalogDirName)

	// Setup Cluster
	go func() {
		strategies := map[string]setup.SchedulerStrategy{
			"local.local": setup.NewLocalLocal("", builder, clusterCmds),
		}
		setup.Main(clusterCmds, strategies, "local.local", []string{})
	}()

	// TODO: This is hacky, wait until we get a response from scheduler
	time.Sleep(2 * time.Second)

	log.Printf("Starting Smoke Test")
	// run smoke test
	wg.Add(1)
	go func() {
		defer wg.Done()
		smokeTestCmd, err := createCmd("$GOPATH/bin/scootapi", "run_smoke_test", "10")
		if err != nil {
			log.Fatal(err)
		}

		err = smokeTestCmd.Start()
		if err != nil {
			log.Fatal(err)
		}

		smokeTestCmd.Wait()
	}()

	// TODO: this is hacky we really want to just queue all the jobs then kill
	// scheduler and workers  Need to refactor smoketest methods to allow us
	// to queue a bunch of jobs, then return
	time.Sleep(1500 * time.Millisecond)
	log.Printf(
		`-------------------------------KILLING CLUSTER-------------------------------
                                     ________________
                            ____/ (  (    )   )  \___
                           /( (  (  )   _    ))  )   )\
                         ((     (   )(    )  )   (   )  )
                       ((/  ( _(   )   (   _) ) (  () )  )
                      ( (  ( (_)   ((    (   )  .((_ ) .  )_
                     ( (  )    (      (  )    )   ) . ) (   )
                    (  (   (  (   ) (  _  ( _) ).  ) . ) ) ( )
                    ( (  (   ) (  )   (  ))     ) _)(   )  )  )
                   ( (  ( \ ) (    (_  ( ) ( )  )   ) )  )) ( )
                    (  (   (  (   (_ ( ) ( _    )  ) (  )  )   )
                   ( (  ( (  (  )     (_  )  ) )  _)   ) _( ( )
                    ((  (   )(    (     _    )   _) _(_ (  (_ )
                     (_((__(_(__(( ( ( |  ) ) ) )_))__))_)___)
                     ((__)        \\||lll|l||///          \_))
                              (   /(/ (  )  ) )\   )
                            (    ( ( ( | | ) ) )\   )
                             (   /(| / ( )) ) ) )) )
                           (     ( ((((_(|)_)))))     )
                            (      ||\(|(|)|/||     )
                          (        |(||(||)||||        )
                            (     //|/l|||)|\\ \     )
                          (/ / //  /|//||||\\  \ \  \ _)
  -------------------------------------------------------------------------------`)
	clusterCmds.Kill()

	log.Printf("Reviving Cluster")
	//Startup Recovery Cluster
	go func() {
		strategies := map[string]setup.SchedulerStrategy{
			"local.local": setup.NewLocalLocal("", builder, clusterCmds),
		}
		setup.Main(clusterCmds, strategies, "local.local", []string{})
	}()

	wg.Wait()
}

// Cleansup any remaining processes and deletes the
// file sagalog
func cleanup(cmds *setup.Cmds, sagaLogDir string) {
	cmds.Kill()
	err := os.RemoveAll(sagaLogDir)
	if err != nil {
		log.Println("Error Cleaning Up SagaLogDir", err)
	}
}

// Run the given command by first evaluating any env vars and redirecting output to global stdout/stderr.
// If blocking, waits for the command to finish and returns any err, otherwise returns nil.
func createCmd(name string, args ...string) (*exec.Cmd, error) {
	for i, _ := range args {
		args[i] = os.ExpandEnv(args[i])
	}
	cmd := exec.Command(os.ExpandEnv(name), args...)
	cmd.Dir = os.ExpandEnv("$GOPATH/src/github.com/scootdev/scoot")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd, nil
}
