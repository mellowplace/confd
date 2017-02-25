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

	if err := initConfig(); err != nil {
		t.Errorf(err.Error())
	}
	if !reflect.DeepEqual(want, config) {
		t.Errorf("initConfig() = %v, want %v", config, want)
	}
}
