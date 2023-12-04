# checkout-go

Repro for https://github.com/etcd-io/etcd/issues/16948

Trying to generate an error like: https://github.com/containerd/containerd/actions/runs/7067476884/job/19240932680?pr=9456

```
Setup go version spec 1.21.3
Attempting to download 1.21.3...
matching 1.21.3...
Not found in manifest.  Falling back to download directly from Go
```

Followed by:

```
/bin/sh: 1: go: not found
```
