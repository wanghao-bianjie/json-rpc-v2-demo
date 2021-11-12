package cmd

import (
	"github.com/spf13/cobra"
	"io/ioutil"
	"myjsonrpcv2/internal/app"
	"myjsonrpcv2/internal/app/conf"
)

var (
	localConfig string
	startCmd    = &cobra.Command{
		Use:   "start",
		Short: "Start Supper Wallet Server.",
		Run: func(cmd *cobra.Command, args []string) {
			online()
		},
	}
	testCmd = &cobra.Command{ // test
		Use:   "test",
		Short: "Start Test Supper Wallet Server.",
		Run: func(cmd *cobra.Command, args []string) {
			test()
		},
	}
)

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.AddCommand(testCmd)
	testCmd.Flags().StringVarP(&localConfig, "CONFIG", "c", "", "conf path: /opt/local.toml")
}

func test() {
	//data, err := ioutil.ReadFile(localConfig)
	data, err := ioutil.ReadFile("configs/cfg.toml")
	if err != nil {
		panic(err)
	}
	config, err := conf.ReadConfig(data)
	if err != nil {
		panic(err)
	}
	run(config)
}

func online() {
	//var config *conf.Config
	//
	//zkConn, err := initialization.NewZkConn()
	//if err != nil {
	//	panic(err)
	//}
	//
	//configPath := "/visualization/config"
	//if v, ok := os.LookupEnv(constant.EnvNameZkConfigPath); ok {
	//	configPath = v
	//}
	//data, _, err := zkConn.Get(configPath)
	//if err != nil {
	//	panic(err)
	//}
	//config, err = conf.ReadConfig(data)
	//if err != nil {
	//	panic(err)
	//}
	//
	//run(config)
}

func run(cfg *conf.Config) {
	app.Serve(cfg)
}
