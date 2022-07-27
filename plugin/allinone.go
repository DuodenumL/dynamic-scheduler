package plugin

import (
	"context"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

const Name = "AllInOne"

type AllInOnePlugin struct {
	handle framework.Handle
}

func NewAllInOnePlugin(conf runtime.Object, handle framework.Handle) (framework.Plugin, error) {
	return &AllInOnePlugin{
		handle: handle,
	}, nil
}

func (a *AllInOnePlugin) Name() string {
	return Name
}

func (a *AllInOnePlugin) Less(info *framework.QueuedPodInfo, info2 *framework.QueuedPodInfo) bool {
	klog.V(4).Infof("less: %+v, %+v", info, info2)
	return info.Pod.Name < info2.Pod.Name
}

func (a *AllInOnePlugin) PreFilter(ctx context.Context, state *framework.CycleState, p *v1.Pod) *framework.Status {
	klog.V(4).Infof("prefilter: %+v", p)
	return nil
}

func (a *AllInOnePlugin) PreFilterExtensions() framework.PreFilterExtensions {
	klog.V(4).Infof("prefilter extensions: %+v", a)
	return nil
}

func (a *AllInOnePlugin) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	klog.V(4).Infof("filter: %+v, %+v", pod.Name, nodeInfo)
	return nil
}

func (a *AllInOnePlugin) PostFilter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, filteredNodeStatusMap framework.NodeToStatusMap) (*framework.PostFilterResult, *framework.Status) {
	klog.V(4).Infof("postfilter: %+v, %+v", pod.Name, filteredNodeStatusMap)
	return nil, nil
}

func (a *AllInOnePlugin) PreScore(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodes []*v1.Node) *framework.Status {
	klog.V(4).Infof("prescore: %+v, %+v", pod.Name, nodes)
	return nil
}

func (a *AllInOnePlugin) Score(ctx context.Context, state *framework.CycleState, p *v1.Pod, nodeName string) (int64, *framework.Status) {
	klog.V(4).Infof("score: %+v, %+v", p.Name, nodeName)
	return int64(len(nodeName)), nil
}

func (a *AllInOnePlugin) ScoreExtensions() framework.ScoreExtensions {
	return a
}

func (a *AllInOnePlugin) NormalizeScore(ctx context.Context, state *framework.CycleState, p *v1.Pod, scores framework.NodeScoreList) *framework.Status {
	klog.V(4).Infof("normalize score: %+v, %+v", p.Name, scores)
	maxScore := Max[int64](0, Map(scores, func(score framework.NodeScore) int64 { return score.Score })...)
	for i := range scores {
		scores[i].Score = int64(float64(scores[i].Score) / float64(maxScore) * float64(framework.MaxNodeScore))
	}
	return nil
}

func (a *AllInOnePlugin) PreBind(ctx context.Context, state *framework.CycleState, p *v1.Pod, nodeName string) *framework.Status {
	klog.V(4).Infof("prebind: %+v, %+v", p.Name, nodeName)
	return nil
}

func (a *AllInOnePlugin) Bind(ctx context.Context, state *framework.CycleState, p *v1.Pod, nodeName string) *framework.Status {
	klog.V(4).Infof("bind: %+v, %+v", p.Name, nodeName)
	return nil
}

func (a *AllInOnePlugin) PostBind(ctx context.Context, state *framework.CycleState, p *v1.Pod, nodeName string) {
	klog.V(4).Infof("postbind: %+v, %+v", p.Name, nodeName)
}

func (a *AllInOnePlugin) Reserve(ctx context.Context, state *framework.CycleState, p *v1.Pod, nodeName string) *framework.Status {
	klog.V(4).Infof("reserve: %+v, %+v", p.Name, nodeName)
	return nil
}

func (a *AllInOnePlugin) Permit(ctx context.Context, state *framework.CycleState, p *v1.Pod, nodeName string) (*framework.Status, time.Duration) {
	klog.V(4).Infof("permit: %+v, %+v", p.Name, nodeName)
	return nil, 0
}

func (a *AllInOnePlugin) Unreserve(ctx context.Context, state *framework.CycleState, p *v1.Pod, nodeName string) {
	klog.V(4).Infof("unreserve: %+v, %+v", p.Name, nodeName)
}
