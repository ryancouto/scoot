package bundlestore

import (
	"log"
	"time"

	"github.com/scootdev/scoot/common/endpoints"
	"github.com/scootdev/scoot/common/stats"
	"github.com/scootdev/scoot/config/jsonconfig"
	"github.com/scootdev/scoot/ice"
	"github.com/scootdev/scoot/os/temp"
)

type servers struct {
	bs  *Server
	obs *endpoints.TwitterServer
}

func makeServers(bs *Server, obs *endpoints.TwitterServer) servers {
	return servers{bs, obs}
}

func Defaults() (*ice.MagicBag, jsonconfig.Schema) {
	bag := ice.NewMagicBag()
	bag.PutMany(
		func() endpoints.StatScope { return "bundleserver" },
		func(scope endpoints.StatScope) stats.StatsReceiver {
			return endpoints.MakeStatsReceiver(scope).Precision(time.Millisecond)
		},
		func() (*temp.TempDir, error) { return temp.TempDirDefault() },
		func(p Processor, s Store, a Addr) *Server {
			return MakeServer(p, s, a)
		},
		func(tmpDir *temp.TempDir) (Store, error) {
			return MakeFileStore(tmpDir)
		},
		func() Processor {
			return &IdentityProcessor{}
		},
		makeServers,
	)

	schema := jsonconfig.Schema(make(map[string]jsonconfig.Implementations))
	return bag, schema
}

// Starts the Server based on the MagicBag and config schema provided
// this method blocks until the server completes running or an
// exception occurs.
func RunServer(
	bag *ice.MagicBag,
	schema jsonconfig.Schema,
	config []byte) {

	// Parse Config
	mod, err := schema.Parse(config)
	if err != nil {
		log.Fatal("Error configuring Worker: ", err)
	}

	// Initialize Objects Based on Config Settings
	bag.InstallModule(mod)

	var servers servers
	err = bag.Extract(&servers)
	if err != nil {
		log.Fatal("Error injecting servers", err)
	}

	errCh := make(chan error)
	go func() {
		errCh <- servers.bs.Serve()
	}()
	go func() {
		errCh <- servers.obs.Serve()
	}()
	log.Fatal("Error serving: ", <-errCh)
}
