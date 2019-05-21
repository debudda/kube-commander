# Kube Commander

## TUI

Kube Commander UI is based on [termui](https://github.com/gizak/termui).

## TODO

- [X] Configuration architecture
- [X] Basic kubernetes/client-go library integration
- [X] TUI library selection
- [X] Basic navigation implementation
- [ ] Determine initial release feature set
- [ ] Colorization and theming
- [ ] Implement features

## Feature Set

- [X] List namespaces
- [X] List pods of namespace
- [ ] List workloads of namespace (2nd level menu?)
- [X] List nodes
- [ ] Basic resources view screen
    - [ ] Pod
    - [ ] Namespace
    - [ ] Node
    - [ ] ReplicaSet
    - [ ] Deployment
    - [ ] StatefulSet
    - [ ] PV
    - [ ] PVC
- [ ] Resource deletion with confirmation
- [ ] Resource editing with external $EDITOR
- [ ] Screen refresh: hotkey + auto-refresh
- [ ] Shortcuts bottom bar (eg. Midnight Commander)
- [ ] Charts
    - [ ] Node capacity
    - [ ] metrics-server based charts

## Current questions

1. How to autotest?
2. Should we support custom resources navigation?
3. **CAN WE BEAT OFFICIAL KUBERNETES DASHBOARD** in terms of features, speed and usability?
