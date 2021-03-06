module github.com/AnatolyRugalev/kube-commander

go 1.14

require (
	github.com/atotto/clipboard v0.1.2
	github.com/creack/pty v1.1.11
	github.com/fsnotify/fsnotify v1.4.7
	github.com/gdamore/tcell v1.4.0
	github.com/golang/protobuf v1.4.1
	github.com/googleapis/gnostic v0.2.0 // indirect
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/kr/text v0.2.0
	github.com/mattn/go-runewidth v0.0.9
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v0.0.7
	go.uber.org/atomic v1.4.0
	golang.org/x/crypto v0.0.0-20200220183623-bac4c82f6975
	google.golang.org/appengine v1.6.1 // indirect
	google.golang.org/protobuf v1.25.0
	k8s.io/api v0.18.3
	k8s.io/apimachinery v0.18.3
	k8s.io/cli-runtime v0.18.3
	k8s.io/client-go v0.18.3
	k8s.io/klog v1.0.0
	k8s.io/kubectl v0.18.3
	sigs.k8s.io/yaml v1.2.0
)

replace vbom.ml/util => github.com/fvbommel/util v0.0.3-0.20200828113648-063df19a25ff
