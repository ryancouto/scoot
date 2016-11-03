package client

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/scootdev/scoot/scootapi/gen-go/scoot"
	"github.com/spf13/cobra"
)

type getStatusCmd struct{}

func (c *getStatusCmd) registerFlags() *cobra.Command {
	return &cobra.Command{
		Use:   "get_job_status",
		Short: "GetJobStatus",
	}
}

func (c *getStatusCmd) run(cl *simpleCLIClient, cmd *cobra.Command, args []string) error {

	log.Println("Checking Status for Scoot Job", args)

	if len(args) == 0 {
		return errors.New("a job id must be provided")
	}

	err := cl.Dial()
	if err != nil {
		return err
	}

	jobId := args[0]

	status, err := cl.scootClient.GetStatus(jobId)

	if err != nil {
		switch err := err.(type) {
		case *scoot.InvalidRequest:
			return fmt.Errorf("Invalid Request: %v", err.GetMessage())
		case *scoot.ScootServerError:
			return fmt.Errorf("Error getting status: %v", err.Error())
		}
	}

	for _, runStatus := range status.GetTaskData() {
		if runStatus.GetStatus() > 2 { // 0 - UNKNOWN, 1 - PENDING, 2 - RUNNING
			c.saveStdOutAndErr(runStatus)
		}
	}

	fmt.Println("Job Status: ", status)

	return nil
}

func (c *getStatusCmd) saveStdOutAndErr(runStatus *scoot.RunStatus) {
	runID, out, err := runStatus.GetRunId(), runStatus.GetOutUri(), runStatus.GetErrUri()
	if _, err := os.Stat("~/scoot-std"); os.IsNotExist(err) {
		os.Mkdir("~/scoot-std", 0777)
	}
	if _, err2 := os.Stat("~/scoot-std/" + runID); os.IsNotExist(err2) {
		os.Mkdir("~/scoot-std"+runID, 0777)
	}
	c.saveStdStream(out, runID)
	c.saveStdStream(err, runID)
}

func (c *getStatusCmd) saveStdStream(uri, runID string) {
	uriSlice := strings.Split(uri, "://")
	scheme := uriSlice[0]
	hierPart := uriSlice[1]
	switch scheme {
	case "file":
		c.scpFile(hierPart, runID)
	default:
		log.Fatal("Error resolving URI protocol")
	}
}

func (c *getStatusCmd) scpFile(hierPart, runID string) {
	re := regexp.MustCompile("([^://?#]*)?")
	authority := re.FindString(hierPart)
	filePath := strings.Split(hierPart, authority)[1]
	scp := exec.Command("scp", authority+":"+filePath, "~/scoot-std/"+runID+"/")
	err := scp.Run()
	if err != nil {
		log.Fatal("Error securely copying file: %v", err)
	}
}
