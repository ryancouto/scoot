package main

import ()

func main() {
	bundleAddr := flag.String("bundle_addr", "localhost:11100", "addr to serve bundles on")
	obsAddr := flag.String("obs_addr", "localhost:11101", "addr to serve observability on")
	var configFlag = flag.String("config", "local.local", "Bundle Server Config (either a filename like local.local or JSON text")
	flag.Parse()

	configText, err := jsonconfig.GetConfigText(*configFlag, nil)
	if err != nil {
		log.Fatal(err)
	}

	bag, schema := server.Defaults()
	bag.PutMany(
		func() server.Addr { return *bundleAddr },
		func(s stats.StatsReceiver) *endpoints.TwitterServer {
			return endpoints.NewTwitterServer(*obsAddr, s)
		},
	)

	server.RunServer(bag, schema, configText)
}
