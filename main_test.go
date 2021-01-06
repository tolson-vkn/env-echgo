package main

import (
	"log"
	"os"
	"reflect"
	"testing"
)

func TestRegistration(t *testing.T) {
	err := os.Setenv("ECHGO_POD_NAME", "env-echgo-5fbc57dbc4-gmsk8")
	if err != nil {
		log.Fatalf("err: %s\n", err)
	}

	expected := make(map[string]string)
	expected["pod_name"] = "env-echgo-5fbc57dbc4-gmsk8"

	registered := registerEchgos()
	if !reflect.DeepEqual(expected, registered) {
		t.Error("ECHGO_POD_NAME was not parsed")
	}

	err = os.Unsetenv("ECHGO_POD_NAME")
	if err != nil {
		log.Fatalf("err: %s\n", err)
	}
}

func TestIgnoreMessage(t *testing.T) {
	err := os.Setenv("ECHGO_MESSAGE", "This shouldn't be seen!")
	if err != nil {
		log.Fatalf("err: %s\n", err)
	}

	expected := make(map[string]string)

	registered := registerEchgos()
	if !reflect.DeepEqual(expected, registered) {
		t.Error("ECHGO_MESSAGE Should be skipped.")
	}

	err = os.Unsetenv("ECHGO_MESSAGE")
	if err != nil {
		log.Fatalf("err: %s\n", err)
	}
}

func TestEchgoEchgo(t *testing.T) {
	err := os.Setenv("ECHGO_ECHGO_WHAT", "This shouldn't work, why would you do it?")
	if err != nil {
		log.Fatalf("err: %s\n", err)
	}

	expected := make(map[string]string)

	registered := registerEchgos()
	if !reflect.DeepEqual(expected, registered) {
		t.Error("ECHGO_ECHGO_WHAT Should be skipped.")
	}

	err = os.Unsetenv("ECHGO_ECHGO_WHAT")
	if err != nil {
		log.Fatalf("err: %s\n", err)
	}
}
