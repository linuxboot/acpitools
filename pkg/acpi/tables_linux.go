// Copyright 2019 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build linux

package acpi

import (
	"path/filepath"
)

var (
	// Methods is a map of string to methods.
	Methods = map[string]TableMethod{
		"files": RawTablesFromSys,
	}
)

// MethodNames returns the list of supported MethodNames.
func MethodNames() string {
	return "files"
}

// RawTablesFromSys returns an array of Raw tables, for all ACPI tables
// available in /sys.
func RawTablesFromSys() ([]Table, error) {
	n, err := filepath.Glob("/sys/firmware/acpi/tables/[A-Z]*")
	if err != nil {
		return nil, err
	}

	var tabs []Table
	for _, t := range n {
		r, err := RawFromName(t)
		if err != nil {
			return nil, err
		}
		tabs = append(tabs, r...)
	}
	return tabs, nil
}
