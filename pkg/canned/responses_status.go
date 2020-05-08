package canned

import (
	"github.com/litesoft-go/mockvoltdb/pkg/utils"
)

//	{
//		"clusterState": "RUNNING",
//		"hostId": 2,
//		"nodeState": "UP",
//		"pid": 653,
//		"shutdownPending": false,
//		"startAction": "RECOVER",
//		"startupProgress": "complete"
//	}
func PortResonses_status() (r *utils.Responder) {
	r = &utils.Responder{}
	r.AddPath("/status").
		Add(`{
	"clusterState": "RUNNING",
	"nodeState": "UP",
	"shutdownPending": false
}`).
		NoMore()
	return
}
