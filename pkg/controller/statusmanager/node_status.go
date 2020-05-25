package statusmanager

import (
	"fmt"
	"strings"
)

type NodeStatus struct {
	Address string
	Success bool
	Reason  string
}

func (status *StatusManager) SetFromNodes(cachedNodeSet map[string]*NodeStatus) {
	status.Lock()
	defer status.Unlock()
	messages := []string{}
	allProcessesDone := true
	for nodeName, _ := range(cachedNodeSet) {
		if !cachedNodeSet[nodeName].Success {
			messages = append(messages, cachedNodeSet[nodeName].Reason)
			allProcessesDone = false
		}
	}
	if allProcessesDone {
		status.setNotDegraded(ClusterNode)
	} else {
		message := strings.Join(messages, "\n")
		log.Info(fmt.Sprintf("Setting degraded status %s", message))
		status.setDegraded(ClusterNode, "ProcessClusterNodeError", message)
	}
}
