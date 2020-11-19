package main

import (
	"flag"
	"net/http"

	"./configuration"
	"./log"

	"./adapter"
	"./format"
	"./router"
	"./writer"
)

func main() {
	config := flag.String("config", "config.yaml", "The path for the config file")
	flag.Parse()

	l := log.NewJsonLogger(&log.LoggerConfiguration{})
	cfg, err := configuration.NewKoanfProvider(configuration.ProviderConfig{
		Logger: l,
		Source: *config,
	})
	if err != nil {
		l.Fatal(err.Error())
	}
	l.SetVerbosity(cfg.GetString("log.verbosity"))

	r := router.NewRouter(&router.RouterConfiguration{Logger: l})

	formatterCfg := &format.FormatterConfiguration{Logger: l, Config: cfg}
	writerCfg := &writer.WriterConfiguration{Logger: l, Config: cfg}
	a, err := adapter.NewAdapter(&adapter.AdapterConfiguration{
		Logger: l,
		Config: cfg,
		FormatterByName: func(name string) (format.Formatter, error) {
			return format.GetFormatter(formatterCfg, name)
		},
		WriterByName: func(name string) (writer.Writer, error) {
			return writer.GetWriter(writerCfg, name)
		},
	})
	if err != nil {
		l.Fatal(err.Error())
	}
	events := make(map[string]*adapter.EventConfiguration)
	cfg.Load("events", &events)
	for path, event := range events {
		c, err := a.AdaptEvent(event)
		if err != nil {
			l.Fatal(err.Error())
		}
		r.AddRoute(path, c)
	}

	http.Handle("/", r)
	port := cfg.GetString("port")
	l.Info("server listening on %s", port)
	l.Fatal(http.ListenAndServe(":"+port, nil).Error())
}
