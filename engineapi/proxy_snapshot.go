package engineapi

import (
	longhorn "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta2"
)

func (p *Proxy) SnapshotCreate(e *longhorn.Engine, name string, labels map[string]string) (string, error) {
	return p.grpcClient.VolumeSnapshot(p.DirectToURL(e), name, labels)
}

func (p *Proxy) SnapshotList(e *longhorn.Engine) (snapshots map[string]*longhorn.Snapshot, err error) {
	recv, err := p.grpcClient.SnapshotList(p.DirectToURL(e))
	if err != nil {
		return nil, err
	}

	snapshots = map[string]*longhorn.Snapshot{}
	for k, v := range recv {
		snapshots[k] = (*longhorn.Snapshot)(v)
	}
	return snapshots, nil
}

func (p *Proxy) SnapshotGet(e *longhorn.Engine, name string) (snapshot *longhorn.Snapshot, err error) {
	recv, err := p.SnapshotList(e)
	if err != nil {
		return nil, err
	}

	return recv[name], nil
}

func (p *Proxy) SnapshotClone(e *longhorn.Engine, name, fromController string) (err error) {
	return p.grpcClient.SnapshotClone(p.DirectToURL(e), name, fromController)
}

func (p *Proxy) SnapshotRevert(e *longhorn.Engine, name string) (err error) {
	return p.grpcClient.SnapshotRevert(p.DirectToURL(e), name)
}

func (p *Proxy) SnapshotPurge(e *longhorn.Engine) (err error) {
	return p.grpcClient.SnapshotPurge(p.DirectToURL(e), true)
}

func (p *Proxy) SnapshotPurgeStatus(e *longhorn.Engine) (status map[string]*longhorn.PurgeStatus, err error) {
	recv, err := p.grpcClient.SnapshotPurgeStatus(p.DirectToURL(e))
	if err != nil {
		return nil, err
	}

	status = map[string]*longhorn.PurgeStatus{}
	for k, v := range recv {
		status[k] = (*longhorn.PurgeStatus)(v)
	}
	return status, nil
}

func (p *Proxy) SnapshotDelete(e *longhorn.Engine, name string) (err error) {
	return p.grpcClient.SnapshotRemove(p.DirectToURL(e), []string{name})
}
