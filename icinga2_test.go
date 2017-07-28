package main

import (
	"reflect"
	"testing"
)

var (
	i Icinga2
)

const (
	HOST_NAME    = "hostname.test.com"
	SERVICE_NAME = "vrf-aws-001-london"
)

func init() {
	i = Icinga2{
		Host:     "localhost",
		Username: "root",
		Password: "icinga",
	}
}

func TestConstructFilter(t *testing.T) {
	expectedFilterOneHost := Filter{
		Filter: "match(\"*" + HOST_NAME + "*\", host.display_name)",
		Attrs: []string{
			"display_name",
			"name",
			"last_check",
			"state",
			"check_command",
		},
	}
	resultFilterOneHost := i.constructFilter(HOST_NAME, HOSTS, false)
	if !reflect.DeepEqual(expectedFilterOneHost, resultFilterOneHost) {
		t.Error(
			"For 'check one host status'",
			"expected", expectedFilterOneHost,
			"got", resultFilterOneHost,
		)
	}

	expectedFilterOneService := Filter{
		Filter: "match(\"*" + SERVICE_NAME + "*\", service.display_name)",
		Attrs: []string{
			"display_name",
			"name",
			"last_check",
			"state",
			"check_command",
		},
	}
	resultFilterOneService := i.constructFilter(SERVICE_NAME, SERVICES, false)
	if !reflect.DeepEqual(expectedFilterOneService, resultFilterOneService) {
		t.Error(
			"For 'check one service status'",
			"expected", expectedFilterOneService,
			"got", resultFilterOneService,
		)
	}

	expectedFilterAllHosts := Filter{
		Attrs: []string{
			"display_name",
			"name",
			"last_check",
			"state",
			"check_command",
		},
	}
	resultFilterAllHosts := i.constructFilter("", HOSTS, true)
	if !reflect.DeepEqual(expectedFilterAllHosts, resultFilterAllHosts) {
		t.Error(
			"For 'check all hosts status'",
			"expected", expectedFilterAllHosts,
			"got", resultFilterAllHosts,
		)
	}

	expectedFilterAllServices := Filter{
		Attrs: []string{
			"display_name",
			"name",
			"last_check",
			"state",
			"check_command",
		},
	}
	resultFilterAllServices := i.constructFilter("", SERVICES, true)
	if !reflect.DeepEqual(expectedFilterAllServices, resultFilterAllServices) {
		t.Error(
			"For 'check all hosts status'",
			"expected", expectedFilterAllServices,
			"got", resultFilterAllServices,
		)
	}
}

func TestCheck(t *testing.T) {
	_, err := i.check("", HOSTS, true)
	if err != nil {
		t.Error("Cannot get list of hosts")
	}

	_, err = i.check("", SERVICES, true)
	if err != nil {
		t.Error("Cannot get list of services")
	}
}
