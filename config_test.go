package main

import (
	"reflect"
	"testing"

	"github.com/kelseyhightower/confd/log"
)

func TestInitConfigDefaultConfig(t *testing.T) {
	log.SetLevel("warn")
	want := Config{
		Backend: []Backend{
			Backend{
				ID:           "default",
				Type:         "etcd",
				BackendNodes: []string{"http://127.0.0.1:4001"},
				ClientCaKeys: "",
				ClientCert:   "",
				ClientKey:    "",
				Prefix:       "",
				SRVDomain:    "",
				Scheme:       "http",
				Table:        "",
			},
		},
		ConfDir:      "/etc/confd",
		Noop:         false,
		Interval:     600,
	}

  defaultConfigFile = "" // stops any config on the host being picked up

	if err := initConfig(); err != nil {
		t.Errorf(err.Error())
	}
	if !reflect.DeepEqual(want, config) {
		t.Errorf("initConfig() = %v, want %v", config, want)
	}
}

func TestInitConfigMultiBackendConfig(t *testing.T) {
	log.SetLevel("warn")
	want := Config{
		Backend: []Backend{
			Backend{
				ID:           "default",
				Type:         "env",
				BackendNodes: []string{"http://127.0.0.1:4001"},
				ClientCaKeys: "",
				ClientCert:   "/etc/confd/ssl/client.crt",
				ClientKey:    "/etc/confd/ssl/client.key",
				Prefix:       "/production",
				SRVDomain:    "etcd.example.com",
				SRVRecord:    "_env._tcp.etcd.example.com.",
				Scheme:       "https",
				Table:        "",
			},
			Backend{
				ID:           "env_vars",
				Type:         "env",
				Prefix:       "/production",
			},
		},
		ConfDir:      "/etc/confd",
		Noop:         false,
		Interval:     600,
		LogLevel:     "debug",
	}

	configFile = "test_fixture/multi_backend_config.toml"

	if err := initConfig(); err != nil {
		t.Errorf(err.Error())
	}
	if !reflect.DeepEqual(want, config) {
		t.Errorf("initConfig() = %v, want %v", config, want)
	}
}
