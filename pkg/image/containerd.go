package image

import (
	"context"
	"fmt"
	"github.com/containerd/containerd"
)

func ContainerdPull(client *containerd.Client, Context context.Context, imageName string) (err error) {
	opts := []containerd.RemoteOpt{
		containerd.WithPullUnpack,
	}
	_, err = client.Pull(Context, imageName, opts...)
	if err != nil {
		return fmt.Errorf("ctr image pull %s failed: %s", imageName, err)
	}
	return nil
}
