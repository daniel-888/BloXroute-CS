package cmd

import (
	"os"
	"os/signal"
	"strings"
	"syscall"
	"fmt"

	client "github.com/daniel-888/BloXroute-CS/client/app"
	"github.com/daniel-888/BloXroute-CS/models"
	"github.com/iamolegga/enviper"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func clientCmdAddItem() *cobra.Command {
	var clientCmd = &cobra.Command{
		Use:   "AddItem",
		Short: "Client application",
		Run: func(cmd *cobra.Command, args []string) {
			log.SetOutput(os.Stdout)

			configuration, err := getClientConfiguration("AddItem")
			if err != nil {
				log.Errorf("Cannot read configuration: %v", err)
				return
			}

			startClientApp(configuration)
		},
	}

	return clientCmd
}

func clientCmdRemoveItem() *cobra.Command {
	var clientCmd = &cobra.Command{
		Use:   "RemoveItem",
		Short: "Client application",
		Run: func(cmd *cobra.Command, args []string) {
			log.SetOutput(os.Stdout)

			configuration, err := getClientConfiguration("RemoveItem")
			if err != nil {
				log.Errorf("Cannot read configuration: %v", err)
				return
			}

			startClientApp(configuration)
		},
	}

	return clientCmd
}

func clientCmdGetItem() *cobra.Command {
	var clientCmd = &cobra.Command{
		Use:   "GetItem",
		Short: "Client application",
		Run: func(cmd *cobra.Command, args []string) {
			log.SetOutput(os.Stdout)

			configuration, err := getClientConfiguration("GetItem")
			if err != nil {
				log.Errorf("Cannot read configuration: %v", err)
				return
			}

			startClientApp(configuration)
		},
	}

	return clientCmd
}

func clientCmdGetAllItems() *cobra.Command {
	var clientCmd = &cobra.Command{
		Use:   "GetAllItems",
		Short: "Client application",
		Run: func(cmd *cobra.Command, args []string) {
			log.SetOutput(os.Stdout)

			configuration, err := getClientConfiguration("GetAllItems")
			if err != nil {
				log.Errorf("Cannot read configuration: %v", err)
				return
			}

			startClientApp(configuration)
		},
	}

	return clientCmd
}

func getClientConfiguration(command string) (client.Configurations, error) {
	e := enviper.New(viper.New())

	var pwd string
	var err error
	if pwd, err = os.Getwd(); err != nil {
		log.Fatal("unable to get current working directory: ", err)
	}

	e.AddConfigPath(pwd)
	e.SetConfigName(".config.client")

	// enable viper to handle env values for nested structs
	e.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// defaults to ENV variable values
	e.AutomaticEnv()

	type RabbitMQConfig struct {
		RabbitMQConfig client.RabbitMQConfig
	}

	var rabbitmqconfig RabbitMQConfig
	if err := e.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatal("Error reading config file: ", err)
		}
	}

	err = e.Unmarshal(&rabbitmqconfig)
	if err != nil {
		log.Errorf("Unable to decode into struct, %v", err)
		return client.Configurations{}, err
	}

	var configuration client.Configurations
	configuration.CommandType = command
	configuration.RabbitMQConfig = rabbitmqconfig.RabbitMQConfig
	return configuration, nil
}

func startClientApp(configuration client.Configurations) {
	c := client.New(configuration)
	app := client.NewApp(c)

	terminate := make(chan os.Signal)
	signal.Notify(terminate, syscall.SIGTERM, syscall.SIGINT)

	commandType, ok := models.CommandType_value[configuration.CommandType]
	fmt.Println("======================================\n", configuration.CommandType, "is ", commandType)
	log.Errorf("======================================\n %s is %d", configuration.CommandType, commandType)
	if !ok {
		log.Errorf("Command not found: %s", configuration.CommandType)
		return
	}
	go func() {
		err := c.InitClient()
		if err != nil {
			log.Errorf("Cannot initialize client: %v", err)
			return
		}

		err = app.Start(models.CommandType(commandType))
		if err != nil {
			log.Errorf("Error happened for client: %v", err)
		}
	}()

	<-terminate
	log.Info("Terminating application")

	err := c.Cleanup()
	if err != nil {
		log.Errorf("Error happened when cleaning up client: %v", err)
		return
	}
}
