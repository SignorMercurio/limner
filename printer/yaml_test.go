package printer

import (
	"bytes"
	"testing"

	"github.com/SignorMercurio/limner/testutil"
)

func TestYamlPrinter(t *testing.T) {
	tests := []struct {
		name string
		src  string
		dst  string
	}{
		{
			name: "k describe po",
			src: `Name:         nginx-7848d4b86f-2pq9t
Namespace:    default
Priority:     0
Node:         docker-desktop/192.168.65.4
Start Time:   Wed, 13 Oct 2021 20:21:46 +0100
Labels:       app=nginx
              pod-template-hash=7848d4b86f
Annotations:  <none>
Status:       Running
IP:           10.1.0.145
IPs:
  IP:           10.1.0.145
Controlled By:  ReplicaSet/nginx-7848d4b86f
Containers:
  nginx:
    Container ID:   docker://8df934bb66ad23d3c6bba75b6ad39285aabbd282e4290b033dd8fced3bc7b26e
    Image:          nginx
    Image ID:       docker-pullable://nginx@sha256:644a70516a26004c97d0d85c7fe1d0c3a67ea8ab7ddf4aff193d9f301670cf36
    Port:           80/TCP
    Host Port:      0/TCP
    State:          Running
      Started:      Thu, 14 Oct 2021 19:20:55 +0100
    Last State:     Terminated
      Reason:       Error
      Exit Code:    255
      Started:      Wed, 13 Oct 2021 20:21:47 +0100
      Finished:     Thu, 14 Oct 2021 19:20:37 +0100
    Ready:          True
    Restart Count:  1
    Environment:    <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from kube-api-access-m9b7k (ro)
Conditions:
Type              Status
Initialized       True 
Ready             True 
ContainersReady   True 
PodScheduled      True 
Volumes:
  kube-api-access-m9b7k:
    Type:                    Projected (a volume that contains injected data from multiple sources)
    TokenExpirationSeconds:  3607
    ConfigMapName:           kube-root-ca.crt
    ConfigMapOptional:       <nil>
    DownwardAPI:             true
QoS Class:                   BestEffort
Node-Selectors:              <none>
Tolerations:                 node.kubernetes.io/not-ready:NoExecute op=Exists for 300s
                             node.kubernetes.io/unreachable:NoExecute op=Exists for 300s
Events:
Type    Reason          Age   From     Message
----    ------          ----  ----     -------
Normal  SandboxChanged  28m   kubelet  Pod sandbox changed, it will be killed and re-created.
Normal  Pulling         28m   kubelet  Pulling image "nginx"
Normal  Pulled          28m   kubelet  Successfully pulled image "nginx" in 2.62320171s
Normal  Created         28m   kubelet  Created container nginx
Normal  Started         28m   kubelet  Started container nginx
`,
			dst: `[31mName[0m:         [32mnginx-7848d4b86f-2pq9t[0m
[31mNamespace[0m:    [32mdefault[0m
[31mPriority[0m:     [33m0[0m
[31mNode[0m:         [32mdocker-desktop/192.168.65.4[0m
[31mStart Time[0m:   [32mWed, 13 Oct 2021 20:21:46 +0100[0m
[31mLabels[0m:       [32mapp=nginx[0m
              [32mpod-template-hash=7848d4b86f[0m
[31mAnnotations[0m:  [36m<none>[0m
[31mStatus[0m:       [32mRunning[0m
[31mIP[0m:           [32m10.1.0.145[0m
[31mIPs[0m:
  [31mIP[0m:           [32m10.1.0.145[0m
[31mControlled By[0m:  [32mReplicaSet/nginx-7848d4b86f[0m
[31mContainers[0m:
  [31mnginx[0m:
    [31mContainer ID[0m:   [32mdocker://8df934bb66ad23d3c6bba75b6ad39285aabbd282e4290b033dd8fced3bc7b26e[0m
    [31mImage[0m:          [32mnginx[0m
    [31mImage ID[0m:       [32mdocker-pullable://nginx@sha256:644a70516a26004c97d0d85c7fe1d0c3a67ea8ab7ddf4aff193d9f301670cf36[0m
    [31mPort[0m:           [32m80/TCP[0m
    [31mHost Port[0m:      [32m0/TCP[0m
    [31mState[0m:          [32mRunning[0m
      [31mStarted[0m:      [32mThu, 14 Oct 2021 19:20:55 +0100[0m
    [31mLast State[0m:     [32mTerminated[0m
      [31mReason[0m:       [32mError[0m
      [31mExit Code[0m:    [33m255[0m
      [31mStarted[0m:      [32mWed, 13 Oct 2021 20:21:47 +0100[0m
      [31mFinished[0m:     [32mThu, 14 Oct 2021 19:20:37 +0100[0m
    [31mReady[0m:          [33mTrue[0m
    [31mRestart Count[0m:  [33m1[0m
    [31mEnvironment[0m:    [36m<none>[0m
    [31mMounts[0m:
      [32m/var/run/secrets/kubernetes.io/serviceaccount from kube-api-access-m9b7k (ro)[0m
[31mConditions[0m:
[34mType              Status[0m
[32mInitialized[0m       [36mTrue [0m
[32mReady[0m             [36mTrue [0m
[32mContainersReady[0m   [36mTrue [0m
[32mPodScheduled[0m      [36mTrue [0m
[31mVolumes[0m:
  [31mkube-api-access-m9b7k[0m:
    [31mType[0m:                    [32mProjected (a volume that contains injected data from multiple sources)[0m
    [31mTokenExpirationSeconds[0m:  [33m3607[0m
    [31mConfigMapName[0m:           [32mkube-root-ca.crt[0m
    [31mConfigMapOptional[0m:       [36m<nil>[0m
    [31mDownwardAPI[0m:             [33mtrue[0m
[31mQoS Class[0m:                   [32mBestEffort[0m
[31mNode-Selectors[0m:              [36m<none>[0m
[31mTolerations[0m:                 [32mnode.kubernetes.io/not-ready:NoExecute op=Exists for 300s[0m
                             [32mnode.kubernetes.io/unreachable:NoExecute op=Exists for 300s[0m
[31mEvents[0m:
[34mType    Reason          Age   From     Message[0m
[37m----[0m    [36m------[0m          [37m----[0m  [36m----[0m     [37m-------[0m
[32mNormal[0m  [32mSandboxChanged[0m  [37m28m[0m   [36mkubelet[0m  [32mPod sandbox changed, it will be killed and re-created.[0m
[32mNormal[0m  [33mPulling[0m         [37m28m[0m   [36mkubelet[0m  [33mPulling image "nginx"[0m
[32mNormal[0m  [32mPulled[0m          [37m28m[0m   [36mkubelet[0m  [32mSuccessfully pulled image "nginx" in 2.62320171s[0m
[32mNormal[0m  [32mCreated[0m         [37m28m[0m   [36mkubelet[0m  [32mCreated container nginx[0m
[32mNormal[0m  [32mStarted[0m         [37m28m[0m   [36mkubelet[0m  [32mStarted container nginx[0m
[32m[0m
`,
		},
		{
			name: "strange cases",
			src: `key: {}
array1:
- item1: is a key
array2:
- item2 is a value
longStr: "hello
"world
"!"
quoted: "hello"
`,
			dst: `[31mkey[0m: {}
[31marray1[0m:
- [31mitem1[0m: [32mis a key[0m
[31marray2[0m:
- [32mitem2 is a value[0m
[31mlongStr[0m: [32mhello[0m
[32mworld[0m
"[32m![0m"
[31mquoted[0m: "[32mhello[0m"
[32m[0m
`,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var w bytes.Buffer
			p := &ColorPrinter{}
			p.Print(tt.src, &w)
			testutil.MustEqual(t, tt.dst, w.String())
			w.Reset()

			p = &ColorPrinter{Type: "yaml"}
			p.Print(tt.src, &w)
			testutil.MustEqual(t, tt.dst, w.String())
		})
	}
}
