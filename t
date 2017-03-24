diff --git a/binaries/apiserver/main.go b/binaries/apiserver/main.go
index 4b4193e..b06f2f6 100644
--- a/binaries/apiserver/main.go
+++ b/binaries/apiserver/main.go
@@ -3,7 +3,7 @@ package main
 import (
 	"flag"
 	"fmt"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"net/http"
 	"time"
 
@@ -25,7 +25,7 @@ func main() {
 	configFlag := flag.String("config", "{}", "API Server Config (either a filename like local.local or JSON text")
 	flag.Parse()
 
-	log.SetFlags(log.LstdFlags | log.LUTC | log.Lshortfile)
+	// log.SetFlags(log.LstdFlags | log.LUTC | log.Lshortfile)
 
 	// The same config will be used for both bundlestore and frontend (TODO: frontend).
 	asset := func(s string) ([]byte, error) {
diff --git a/binaries/daemon/main.go b/binaries/daemon/main.go
index c78ee42..38076d4 100644
--- a/binaries/daemon/main.go
+++ b/binaries/daemon/main.go
@@ -2,7 +2,7 @@ package main
 
 import (
 	"flag"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"time"
 
 	"github.com/scootdev/scoot/daemon/server"
diff --git a/binaries/minfs/main.go b/binaries/minfs/main.go
index e70b618..2af92c3 100644
--- a/binaries/minfs/main.go
+++ b/binaries/minfs/main.go
@@ -1,7 +1,7 @@
 package main
 
 import (
-	"log"
+	log "github.com/Sirupsen/logrus"
 
 	"github.com/scootdev/scoot/fs/minfuse"
 )
diff --git a/binaries/recoverytest/main.go b/binaries/recoverytest/main.go
index 60189bd..b955d34 100644
--- a/binaries/recoverytest/main.go
+++ b/binaries/recoverytest/main.go
@@ -1,7 +1,7 @@
 package main
 
 import (
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"sync"
 	"time"
 
diff --git a/binaries/scheduler/main.go b/binaries/scheduler/main.go
index 45e6921..8ea28fe 100644
--- a/binaries/scheduler/main.go
+++ b/binaries/scheduler/main.go
@@ -4,7 +4,7 @@ package main
 
 import (
 	"flag"
-	"log"
+	log "github.com/Sirupsen/logrus"
 
 	"github.com/apache/thrift/lib/go/thrift"
 
@@ -26,7 +26,7 @@ var configFlag = flag.String("config", "local.memory", "Scheduler Config (either
 func main() {
 	flag.Parse()
 
-	log.SetFlags(log.LstdFlags | log.LUTC | log.Lshortfile)
+	// log.SetFlags(log.LstdFlags | log.LUTC | log.Lshortfile)
 
 	configText, err := jsonconfig.GetConfigText(*configFlag, config.Asset)
 	if err != nil {
diff --git a/binaries/scoot-snapshot-db/main.go b/binaries/scoot-snapshot-db/main.go
index 7e141a9..eaa3b96 100644
--- a/binaries/scoot-snapshot-db/main.go
+++ b/binaries/scoot-snapshot-db/main.go
@@ -2,7 +2,7 @@ package main
 
 import (
 	"fmt"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"os"
 
 	"github.com/spf13/cobra"
@@ -18,7 +18,7 @@ import (
 )
 
 func main() {
-	log.SetFlags(log.LstdFlags | log.LUTC | log.Lshortfile)
+	// log.SetFlags(log.LstdFlags | log.LUTC | log.Lshortfile)
 
 	inj := &injector{}
 	cmd := cli.MakeDBCLI(inj)
diff --git a/binaries/scootapi/main.go b/binaries/scootapi/main.go
index d4d1e77..033cabf 100644
--- a/binaries/scootapi/main.go
+++ b/binaries/scootapi/main.go
@@ -1,7 +1,7 @@
 package main
 
 import (
-	"log"
+	log "github.com/Sirupsen/logrus"
 
 	"github.com/apache/thrift/lib/go/thrift"
 	"github.com/scootdev/scoot/common/dialer"
@@ -18,7 +18,7 @@ import (
 //		--addr [<host:port> of cloud server]
 
 func main() {
-	log.SetFlags(log.LstdFlags | log.LUTC | log.Lshortfile)
+	// log.SetFlags(log.LstdFlags | log.LUTC | log.Lshortfile)
 
 	transportFactory := thrift.NewTTransportFactory()
 	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
diff --git a/binaries/setup-cloud-scoot/main.go b/binaries/setup-cloud-scoot/main.go
index 3088eb4..ba52f81 100644
--- a/binaries/setup-cloud-scoot/main.go
+++ b/binaries/setup-cloud-scoot/main.go
@@ -2,7 +2,7 @@ package main
 
 import (
 	"flag"
-	"log"
+	log "github.com/Sirupsen/logrus"
 
 	"github.com/scootdev/scoot/os/temp"
 	"github.com/scootdev/scoot/scootapi/setup"
@@ -17,7 +17,7 @@ func main() {
 	apiserversFlag := flag.Int("apiservers", setup.DefaultApiServerCount, "number of apiservers to use")
 	flag.Parse()
 
-	log.SetFlags(log.LstdFlags | log.LUTC | log.Lshortfile)
+	// log.SetFlags(log.LstdFlags | log.LUTC | log.Lshortfile)
 
 	tmp, err := temp.NewTempDir("", "setup-cloud-scoot-")
 	if err != nil {
diff --git a/binaries/workercl/main.go b/binaries/workercl/main.go
index b033e13..efb9ade 100644
--- a/binaries/workercl/main.go
+++ b/binaries/workercl/main.go
@@ -1,7 +1,7 @@
 package main
 
 import (
-	"log"
+	log "github.com/Sirupsen/logrus"
 
 	"github.com/apache/thrift/lib/go/thrift"
 	"github.com/scootdev/scoot/common/dialer"
diff --git a/binaries/workerserver/main.go b/binaries/workerserver/main.go
index e13a41c..9a9d3c7 100644
--- a/binaries/workerserver/main.go
+++ b/binaries/workerserver/main.go
@@ -4,7 +4,7 @@ package main
 
 import (
 	"flag"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"math/rand"
 	"net/http"
 	"strings"
@@ -38,7 +38,7 @@ var storeHandle = flag.String("bundlestore", "", "Abs file path or an http 'host
 func main() {
 	flag.Parse()
 
-	log.SetFlags(log.LstdFlags | log.LUTC | log.Lshortfile)
+	// log.SetFlags(log.LstdFlags | log.LUTC | log.Lshortfile)
 
 	configText, err := jsonconfig.GetConfigText(*configFlag, config.Asset)
 	if err != nil {
diff --git a/common/dialer/dialer.go b/common/dialer/dialer.go
index a65e911..9a96157 100644
--- a/common/dialer/dialer.go
+++ b/common/dialer/dialer.go
@@ -4,7 +4,7 @@ package dialer
 
 import (
 	"fmt"
-	"log"
+	log "github.com/Sirupsen/logrus"
 
 	"github.com/apache/thrift/lib/go/thrift"
 )
diff --git a/common/endpoints/endpoints.go b/common/endpoints/endpoints.go
index 79402fa..4affb72 100644
--- a/common/endpoints/endpoints.go
+++ b/common/endpoints/endpoints.go
@@ -6,7 +6,7 @@ import (
 	"bytes"
 	"fmt"
 	"io"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"net/http"
 	"time"
 
diff --git a/common/endpoints/setup.go b/common/endpoints/setup.go
index 779a841..35df108 100644
--- a/common/endpoints/setup.go
+++ b/common/endpoints/setup.go
@@ -1,7 +1,7 @@
 package endpoints
 
 import (
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"net/http"
 	"time"
 
diff --git a/common/stats/stats.go b/common/stats/stats.go
index a762a74..867875e 100644
--- a/common/stats/stats.go
+++ b/common/stats/stats.go
@@ -18,7 +18,7 @@ package stats
 
 import (
 	"encoding/json"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"strings"
 	"time"
 
diff --git a/config/jsonconfig/config.go b/config/jsonconfig/config.go
index 1918bea..4e1ba8d 100644
--- a/config/jsonconfig/config.go
+++ b/config/jsonconfig/config.go
@@ -3,7 +3,7 @@ package jsonconfig
 import (
 	"encoding/json"
 	"fmt"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"path"
 	"regexp"
 
diff --git a/fs/min/serve.go b/fs/min/serve.go
index 58c385c..bba61aa 100644
--- a/fs/min/serve.go
+++ b/fs/min/serve.go
@@ -2,7 +2,7 @@ package min
 
 import (
 	"fmt"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"runtime"
 	"time"
 
diff --git a/fs/minfuse/fs.go b/fs/minfuse/fs.go
index 211f1d6..b9c14fd 100644
--- a/fs/minfuse/fs.go
+++ b/fs/minfuse/fs.go
@@ -2,7 +2,7 @@ package minfuse
 
 import (
 	"io"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"os"
 	"syscall"
 	"time"
diff --git a/fs/minfuse/runfs.go b/fs/minfuse/runfs.go
index 6fa21a1..e5f9895 100644
--- a/fs/minfuse/runfs.go
+++ b/fs/minfuse/runfs.go
@@ -3,7 +3,7 @@ package minfuse
 import (
 	"errors"
 	"flag"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"net/http"
 	"os"
 	"os/signal"
@@ -29,7 +29,7 @@ type Options struct {
 }
 
 func SetupLog() {
-	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
+	// log.SetFlags(log.LstdFlags | log.Lmicroseconds)
 }
 
 func InitFlags() (*Options, error) {
diff --git a/fuse/connection.go b/fuse/connection.go
index c71e8ac..908cbc6 100644
--- a/fuse/connection.go
+++ b/fuse/connection.go
@@ -6,7 +6,7 @@ import (
 	"io"
 	"os"
 
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"syscall"
 	"unsafe"
 )
diff --git a/fuse/fuse.go b/fuse/fuse.go
index 43ee5a4..0fb940f 100644
--- a/fuse/fuse.go
+++ b/fuse/fuse.go
@@ -66,7 +66,7 @@ package fuse // import "github.com/scootdev/scoot/fuse"
 
 import (
 	"fmt"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"os"
 	"syscall"
 	"time"
diff --git a/fuse/mount.go b/fuse/mount.go
index 8054e90..2b465e2 100644
--- a/fuse/mount.go
+++ b/fuse/mount.go
@@ -4,7 +4,7 @@ import (
 	"bufio"
 	"errors"
 	"io"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"sync"
 )
 
diff --git a/fuse/mount_darwin.go b/fuse/mount_darwin.go
index d40aa19..268fb33 100644
--- a/fuse/mount_darwin.go
+++ b/fuse/mount_darwin.go
@@ -3,7 +3,7 @@ package fuse
 import (
 	"errors"
 	"fmt"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"os"
 	"os/exec"
 	"path"
diff --git a/fuse/mount_freebsd.go b/fuse/mount_freebsd.go
index 70bb410..3a184ac 100644
--- a/fuse/mount_freebsd.go
+++ b/fuse/mount_freebsd.go
@@ -2,7 +2,7 @@ package fuse
 
 import (
 	"fmt"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"os"
 	"os/exec"
 	"strings"
diff --git a/fuse/mount_linux.go b/fuse/mount_linux.go
index 197d104..5419b4b 100644
--- a/fuse/mount_linux.go
+++ b/fuse/mount_linux.go
@@ -2,7 +2,7 @@ package fuse
 
 import (
 	"fmt"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"net"
 	"os"
 	"os/exec"
diff --git a/ice/eval.go b/ice/eval.go
index 9e495fb..b68c2bb 100644
--- a/ice/eval.go
+++ b/ice/eval.go
@@ -4,7 +4,7 @@ import (
 	"bytes"
 	"fmt"
 	"github.com/davecgh/go-spew/spew"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"reflect"
 	"runtime"
 	"runtime/debug"
diff --git a/runner/runners/chaos.go b/runner/runners/chaos.go
index add0975..6f26bcd 100644
--- a/runner/runners/chaos.go
+++ b/runner/runners/chaos.go
@@ -7,7 +7,7 @@ import (
 
 	"github.com/davecgh/go-spew/spew"
 	"github.com/scootdev/scoot/runner"
-	"log"
+	log "github.com/Sirupsen/logrus"
 )
 
 // chaos.go: impl that introduces errors (for testing)
diff --git a/runner/runners/invoke.go b/runner/runners/invoke.go
index 8c1b8f3..0ec36f0 100644
--- a/runner/runners/invoke.go
+++ b/runner/runners/invoke.go
@@ -3,7 +3,7 @@ package runners
 import (
 	"fmt"
 	"io"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"os"
 	"path/filepath"
 	"time"
diff --git a/runner/runners/polling.go b/runner/runners/polling.go
index c750112..45222b2 100644
--- a/runner/runners/polling.go
+++ b/runner/runners/polling.go
@@ -5,7 +5,7 @@ import (
 
 	"github.com/davecgh/go-spew/spew"
 	"github.com/scootdev/scoot/runner"
-	"log"
+	log "github.com/Sirupsen/logrus"
 )
 
 // polling.go: turns a StatusQueryNower into a StatusQuerier by polling
diff --git a/runner/runners/queue.go b/runner/runners/queue.go
index 3f8a118..0f1efcf 100644
--- a/runner/runners/queue.go
+++ b/runner/runners/queue.go
@@ -9,7 +9,7 @@ import (
 	"github.com/scootdev/scoot/runner"
 	"github.com/scootdev/scoot/runner/execer"
 	"github.com/scootdev/scoot/snapshot"
-	"log"
+	log "github.com/Sirupsen/logrus"
 )
 
 const QueueFullMsg = "No resources available. Please try later."
diff --git a/runner/runners/status_manager.go b/runner/runners/status_manager.go
index 97a08c7..022c808 100644
--- a/runner/runners/status_manager.go
+++ b/runner/runners/status_manager.go
@@ -7,7 +7,7 @@ import (
 	"time"
 
 	"github.com/scootdev/scoot/runner"
-	"log"
+	log "github.com/Sirupsen/logrus"
 )
 
 const UnknownRunIDMsg = "unknown run id %v"
diff --git a/sched/scheduler/cluster_state.go b/sched/scheduler/cluster_state.go
index c436d16..b67814b 100644
--- a/sched/scheduler/cluster_state.go
+++ b/sched/scheduler/cluster_state.go
@@ -2,7 +2,7 @@ package scheduler
 
 import (
 	"github.com/scootdev/scoot/cloud/cluster"
-	"log"
+	log "github.com/Sirupsen/logrus"
 )
 
 const noTask = ""
diff --git a/sched/scheduler/recover_jobs.go b/sched/scheduler/recover_jobs.go
index 08643a2..e51dd17 100644
--- a/sched/scheduler/recover_jobs.go
+++ b/sched/scheduler/recover_jobs.go
@@ -1,7 +1,7 @@
 package scheduler
 
 import (
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"math"
 	"sync"
 	"time"
diff --git a/sched/scheduler/stateful_scheduler.go b/sched/scheduler/stateful_scheduler.go
index 0bd1677..8062c20 100644
--- a/sched/scheduler/stateful_scheduler.go
+++ b/sched/scheduler/stateful_scheduler.go
@@ -1,7 +1,7 @@
 package scheduler
 
 import (
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"strings"
 	"time"
 
diff --git a/sched/scheduler/stateful_scheduler_test.go b/sched/scheduler/stateful_scheduler_test.go
index 39015d1..4ec5010 100644
--- a/sched/scheduler/stateful_scheduler_test.go
+++ b/sched/scheduler/stateful_scheduler_test.go
@@ -3,7 +3,7 @@ package scheduler
 import (
 	"errors"
 	"fmt"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"testing"
 	"time"
 
diff --git a/sched/scheduler/task_runner.go b/sched/scheduler/task_runner.go
index 4e05f5e..c71d540 100644
--- a/sched/scheduler/task_runner.go
+++ b/sched/scheduler/task_runner.go
@@ -2,7 +2,7 @@ package scheduler
 
 import (
 	"fmt"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"strings"
 	"time"
 
diff --git a/scootapi/client/get_status_cmd.go b/scootapi/client/get_status_cmd.go
index daf0bd2..3a61f7f 100644
--- a/scootapi/client/get_status_cmd.go
+++ b/scootapi/client/get_status_cmd.go
@@ -6,7 +6,7 @@ import (
 	"fmt"
 	"github.com/scootdev/scoot/scootapi/gen-go/scoot"
 	"github.com/spf13/cobra"
-	"log"
+	log "github.com/Sirupsen/logrus"
 )
 
 type getStatusCmd struct {
diff --git a/scootapi/client/run_job_cmd.go b/scootapi/client/run_job_cmd.go
index f96da47..986f85e 100644
--- a/scootapi/client/run_job_cmd.go
+++ b/scootapi/client/run_job_cmd.go
@@ -8,7 +8,7 @@ import (
 	"github.com/scootdev/scoot/scootapi/gen-go/scoot"
 	"github.com/spf13/cobra"
 	"io/ioutil"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"os"
 )
 
diff --git a/scootapi/client/smoke_test_cmd.go b/scootapi/client/smoke_test_cmd.go
index 260b636..e27c378 100644
--- a/scootapi/client/smoke_test_cmd.go
+++ b/scootapi/client/smoke_test_cmd.go
@@ -12,7 +12,7 @@ import (
 	"github.com/scootdev/scoot/scootapi/gen-go/scoot"
 	"github.com/scootdev/scoot/tests/testhelpers"
 	"github.com/spf13/cobra"
-	"log"
+	log "github.com/Sirupsen/logrus"
 )
 
 type smokeTestCmd struct {
diff --git a/scootapi/client/watch_job_cmd.go b/scootapi/client/watch_job_cmd.go
index 7537e55..1960f38 100644
--- a/scootapi/client/watch_job_cmd.go
+++ b/scootapi/client/watch_job_cmd.go
@@ -2,7 +2,7 @@ package client
 
 import (
 	"fmt"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"time"
 
 	"github.com/pkg/errors"
diff --git a/scootapi/server/setup.go b/scootapi/server/setup.go
index 69a10d1..94a346a 100644
--- a/scootapi/server/setup.go
+++ b/scootapi/server/setup.go
@@ -1,7 +1,7 @@
 package server
 
 import (
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"time"
 
 	"github.com/apache/thrift/lib/go/thrift"
diff --git a/scootapi/setup/api.go b/scootapi/setup/api.go
index 40dff77..fe36a80 100644
--- a/scootapi/setup/api.go
+++ b/scootapi/setup/api.go
@@ -2,7 +2,7 @@ package setup
 
 import (
 	"fmt"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"os"
 
 	"github.com/scootdev/scoot/os/temp"
diff --git a/scootapi/setup/cmds.go b/scootapi/setup/cmds.go
index 030f157..72c49d5 100644
--- a/scootapi/setup/cmds.go
+++ b/scootapi/setup/cmds.go
@@ -2,7 +2,7 @@ package setup
 
 import (
 	"fmt"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"os"
 	"os/exec"
 	"os/signal"
diff --git a/scootapi/setup/ports.go b/scootapi/setup/ports.go
index 42af7b3..117c1da 100644
--- a/scootapi/setup/ports.go
+++ b/scootapi/setup/ports.go
@@ -2,7 +2,7 @@ package setup
 
 import (
 	"fmt"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"os/exec"
 	"strconv"
 	"time"
diff --git a/scootapi/setup/sched.go b/scootapi/setup/sched.go
index 264324d..109b3f1 100644
--- a/scootapi/setup/sched.go
+++ b/scootapi/setup/sched.go
@@ -1,7 +1,7 @@
 package setup
 
 import (
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"strconv"
 	"strings"
 
diff --git a/scootapi/setup/workers.go b/scootapi/setup/workers.go
index 8cdb1a4..b0db212 100644
--- a/scootapi/setup/workers.go
+++ b/scootapi/setup/workers.go
@@ -2,7 +2,7 @@ package setup
 
 import (
 	"fmt"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"strconv"
 
 	"github.com/scootdev/scoot/scootapi"
diff --git a/snapshot/bundlestore/file_store.go b/snapshot/bundlestore/file_store.go
index 2b3541f..0eb93a7 100644
--- a/snapshot/bundlestore/file_store.go
+++ b/snapshot/bundlestore/file_store.go
@@ -3,7 +3,7 @@ package bundlestore
 import (
 	"errors"
 	"io"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"os"
 	"path/filepath"
 	"strings"
diff --git a/snapshot/bundlestore/groupcache.go b/snapshot/bundlestore/groupcache.go
index cd21328..4aafdf7 100644
--- a/snapshot/bundlestore/groupcache.go
+++ b/snapshot/bundlestore/groupcache.go
@@ -4,7 +4,7 @@ import (
 	"bytes"
 	"io"
 	"io/ioutil"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"net/http"
 	"time"
 
diff --git a/snapshot/bundlestore/http_store.go b/snapshot/bundlestore/http_store.go
index 62d9aba..6424eb5 100644
--- a/snapshot/bundlestore/http_store.go
+++ b/snapshot/bundlestore/http_store.go
@@ -5,7 +5,7 @@ import (
 	"fmt"
 	"io"
 	"io/ioutil"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"net/http"
 	"os"
 	"strings"
diff --git a/snapshot/bundlestore/server.go b/snapshot/bundlestore/server.go
index 2c56cd1..3d24cae 100644
--- a/snapshot/bundlestore/server.go
+++ b/snapshot/bundlestore/server.go
@@ -3,7 +3,7 @@ package bundlestore
 import (
 	"fmt"
 	"io"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"net/http"
 	"regexp"
 	"strings"
diff --git a/snapshot/file_backed.go b/snapshot/file_backed.go
index 0cf2520..e0557dd 100644
--- a/snapshot/file_backed.go
+++ b/snapshot/file_backed.go
@@ -2,7 +2,7 @@ package snapshot
 
 import (
 	"io/ioutil"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"os"
 
 	"github.com/scootdev/scoot/fs/perf"
diff --git a/snapshot/git/gitdb/bundlestore.go b/snapshot/git/gitdb/bundlestore.go
index f42e514..3d7ae9f 100644
--- a/snapshot/git/gitdb/bundlestore.go
+++ b/snapshot/git/gitdb/bundlestore.go
@@ -4,7 +4,7 @@ import (
 	"errors"
 	"fmt"
 	"io"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"os"
 	"os/exec"
 	"path"
diff --git a/snapshot/git/gitdb/db_test.go b/snapshot/git/gitdb/db_test.go
index 362fb4b..d04d389 100644
--- a/snapshot/git/gitdb/db_test.go
+++ b/snapshot/git/gitdb/db_test.go
@@ -4,7 +4,7 @@ import (
 	"flag"
 	"fmt"
 	"io/ioutil"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"os"
 	"os/exec"
 	"path/filepath"
diff --git a/snapshot/git/gitfiler/cloner.go b/snapshot/git/gitfiler/cloner.go
index 66939f1..421ebdb 100644
--- a/snapshot/git/gitfiler/cloner.go
+++ b/snapshot/git/gitfiler/cloner.go
@@ -2,7 +2,7 @@ package gitfiler
 
 import (
 	"fmt"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"os/exec"
 
 	"github.com/scootdev/scoot/common/stats"
diff --git a/snapshot/git/repo/repo.go b/snapshot/git/repo/repo.go
index adf7774..aaeb845 100644
--- a/snapshot/git/repo/repo.go
+++ b/snapshot/git/repo/repo.go
@@ -5,7 +5,7 @@ package repo
 
 import (
 	"fmt"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"os"
 	"os/exec"
 	"strings"
diff --git a/snapshot/utils/checkout/checkout.go b/snapshot/utils/checkout/checkout.go
index 732b89e..54cec0f 100644
--- a/snapshot/utils/checkout/checkout.go
+++ b/snapshot/utils/checkout/checkout.go
@@ -9,7 +9,7 @@ import (
 	"fmt"
 	"io"
 	"io/ioutil"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"os"
 	"path"
 
diff --git a/snapshot/utils/countfiles/countfiles.go b/snapshot/utils/countfiles/countfiles.go
index 4202a6b..d89c1d1 100644
--- a/snapshot/utils/countfiles/countfiles.go
+++ b/snapshot/utils/countfiles/countfiles.go
@@ -7,7 +7,7 @@ package main
 import (
 	"flag"
 	"fmt"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"os"
 	"path"
 
diff --git a/tests/testhelpers/clusterHelpers.go b/tests/testhelpers/clusterHelpers.go
index 38962e2..31fd482 100644
--- a/tests/testhelpers/clusterHelpers.go
+++ b/tests/testhelpers/clusterHelpers.go
@@ -1,7 +1,7 @@
 package testhelpers
 
 import (
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"time"
 
 	"github.com/scootdev/scoot/os/temp"
diff --git a/tests/testhelpers/jobHelpers.go b/tests/testhelpers/jobHelpers.go
index 6e5bb9d..7d33251 100644
--- a/tests/testhelpers/jobHelpers.go
+++ b/tests/testhelpers/jobHelpers.go
@@ -2,7 +2,7 @@ package testhelpers
 
 import (
 	"fmt"
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"math/rand"
 	"sort"
 	"time"
diff --git a/workerapi/client/commands.go b/workerapi/client/commands.go
index b477455..d4d5e7e 100644
--- a/workerapi/client/commands.go
+++ b/workerapi/client/commands.go
@@ -1,7 +1,7 @@
 package client
 
 import (
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"time"
 
 	"github.com/scootdev/scoot/runner"
diff --git a/workerapi/server/server.go b/workerapi/server/server.go
index 2885dad..65056ab 100644
--- a/workerapi/server/server.go
+++ b/workerapi/server/server.go
@@ -3,7 +3,7 @@
 package server
 
 import (
-	"log"
+	log "github.com/Sirupsen/logrus"
 	"sync"
 	"time"
 
diff --git a/workerapi/server/setup.go b/workerapi/server/setup.go
index d5bf6b3..4718fb8 100644
--- a/workerapi/server/setup.go
+++ b/workerapi/server/setup.go
@@ -1,7 +1,7 @@
 package server
 
 import (
-	"log"
+	log "github.com/Sirupsen/logrus"
 
 	"github.com/apache/thrift/lib/go/thrift"
 	"github.com/scootdev/scoot/common/endpoints"
