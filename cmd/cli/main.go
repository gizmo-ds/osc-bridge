package main

import (
	"os"
	"osc-bridge/internal/config"
	"osc-bridge/internal/hubsub"
	osc2 "osc-bridge/internal/osc"

	"github.com/hypebeast/go-osc/osc"
	"github.com/urfave/cli/v2"
)

var AppVersion = "v0.1.0"

func main() {
	hub := hubsub.NewHub[osc.Message, config.Bridge](10)
	app := &cli.App{
		Name:    "osc-bridge",
		Usage:   "OSC forwarding bridge",
		Version: AppVersion,
		Suggest: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "addr",
				Value: "127.0.0.1:9001",
			},
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "config.toml",
			},
		},
		Action: func(ctx *cli.Context) error {
			conf, err := config.Load(ctx.String("config"))
			if err != nil {
				return err
			}

			for _, b := range conf.Bridges {
				hub.Sub(b, func(msg *osc.Message, conf config.Bridge) {
					_ = osc2.Send(conf.Addr, msg)
				})
			}

			addr := ctx.String("addr")
			if !ctx.IsSet("addr") && conf.Addr != "" {
				addr = conf.Addr
			}
			server := osc2.NewServer(addr, hub)
			return server.Run()
		},
	}
	_ = app.Run(os.Args)
}
