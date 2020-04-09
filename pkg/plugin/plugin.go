package plugin

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
	"k8s.io/kubernetes/pkg/scheduler/nodeinfo"
)

const Name = "TopologyAffinity"

type TopologyAffinityPlugin struct {
}

// New initializes a new plugin and returns it.
func New(_ *runtime.Unknown, _ framework.FrameworkHandle) (framework.Plugin, error) {
	return &TopologyAffinityPlugin{}, nil
}

// Ensure we implement the FilterPluin interface
var _ framework.FilterPlugin = &TopologyAffinityPlugin{}

// Ensure we implement the UnreservePlugin interface
var _ framework.UnreservePlugin = &TopologyAffinityPlugin{}

func (t *TopologyAffinityPlugin) Name() string {
	return Name
}

func (t *TopologyAffinityPlugin) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeInfo *nodeinfo.NodeInfo) *framework.Status {
	klog.Infof("TopologyAffinityPlugin: Filter was called for pod %s and node %+v", pod.Name, nodeInfo.Node().Name)
	// nil is considered Success
	return nil
}

func (t *TopologyAffinityPlugin) Unreserve(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) {
	klog.Infof("TopologyAffinityPlugin: Unreserve was called for pod %s and node %+v", pod.Name, nodeName)
	klog.Infof("TopologyAffinityPlugin: pod phase: %s", pod.Status.Phase)
}
