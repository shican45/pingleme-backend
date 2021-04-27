//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package util

type InterfaceTypeErr struct {
	Name string
}

func (err *InterfaceTypeErr) Error() string {
	return "unknown interface " + err.Name + " ."
}
