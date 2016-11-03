package client

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

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
	runID, stdOut, stdErr := runStatus.GetRunId(), runStatus.GetOutUri(), runStatus.GetErrUri()
	homeDir := os.Getenv("HOME")
	runDir := fmt.Sprintf(runID+"_%s.", strings.Replace(time.Now().String(), " ", "_", -1))
	dir := filepath.Join(homeDir, "scoot-std", dir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir), 0777)
	}
	c.saveStdStream(stdOut, dir)
	c.saveStdStream(stdErr, dir)
}

func (c *getStatusCmd) saveStdStream(uri, dir string) {
	uriSlice := strings.Split(uri, "://")
	scheme := uriSlice[0]
	hierPart := uriSlice[1]
	switch scheme {
	case "file":
		c.scpFile(hierPart, dir)
	default:
		log.Fatal("Error resolving URI protocol")
	}
}

func (c *getStatusCmd) scpFile(hierPart, dir string) {
	re := regexp.MustCompile("([^://?#]*)?")
	authority := re.FindString(hierPart)
	filePath := strings.Split(hierPart, authority)[1]
	scp := exec.Command("scp", "-v", authority+":"+filePath, dir)
	err := scp.Run()
	if err != nil {
		log.Println("hier: ", hierPart)
		log.Println("auth: ", authority)
		log.Println("filepath: ", filePath)
		log.Fatal("Error securely copying file: ", err)
	}
}
