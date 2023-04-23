package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	BaseURL  string    `mapstructure:"base_url"`
	OS       OS        `mapstructure:"os"`
	Clusters []Cluster `mapstructure:"clusters"`
}

type Cluster struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
	Vip     string `mapstructure:"vip"`
	Nodes   []Node `mapstructure:"nodes"`
}

type OS struct {
	Token          string   `mapstructure:"token"`
	Password       string   `mapstructure:"password"`
	NtpServers     []string `mapstructure:"ntp_servers"`
	DNSNameservers []string `mapstructure:"dns_nameservers"`
	Gateway        string   `mapstructure:"gateway"`
	SubnetMask     string   `mapstructure:"subnet_mask"`
}

type Node struct {
	Hostname string `mapstructure:"hostname"`
	IP       string `mapstructure:"ip"`
	Mac      string `mapstructure:"mac"`
	Nic      string `mapstructure:"nic"`
	Device   string `mapstructure:"device"`
	Mode     string `mapstructure:"mode"`
}

type Value struct {
	BaseURL string  `mapstructure:"base_url"`
	OS      OS      `mapstructure:"os"`
	Cluster Cluster `mapstructure:"cluster"`
	Node    Node    `mapstructure:"node"`
}

var (
	Conf Config
)

func dynamicConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(&Conf); err != nil {
			panic(fmt.Errorf("unable to decode into struct, %w", err))
		}
	})
}

func Load() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/harvester-netboot/")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unable to decode into struct, %w", err))
	}
	go dynamicConfig()
}
