package main

import (
	"github.com/colinmarc/hdfs"
)

func (jt *jobTracker) testLogsDir() error {
	client, err := hdfs.New(*namenodeAddress)
	defer client.Close()
	if err != nil {
		return err
	}

	_, err = client.ReadDir(*yarnLogDir)
	return err
}
