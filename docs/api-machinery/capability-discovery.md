# Capability Discovery

## Table of Contents

* [Capability Discovery](#capability-discovery)
  * [Discovery Go Package](#discovery-go-package)
    * [Building a ClusterQueryClient](#building-a-clusterqueryclient)
    * [Building and Executing Queries](#building-and-executing-queries)
  * [Executing Pre-defined TKG queries](#executing-pre-defined-tkg-queries)
  * [Capability CRD](#capability-crd)
    * [Example Capability Custom Resource](#example-capability-custom-resource)

------------------------

The capability discovery Go package `pkg/v1/sdk/capabilities/discovery`, along with the `Capability` CRD offer the
ability to query a cluster's capabilities. A "capability" is defined as anything a Kubernetes cluster can do or have,
such as objects and the API surface area. Capability discovery can be used to answer questions such
as `Is this a TKG cluster?`, `Does this cluster have a resource Foo?` etc.

## Discovery Go Package

The [`pkg/v1/sdk/capabilities/discovery`](https://github.com/vmware-tanzu/tanzu-framework/tree/main/pkg/v1/sdk/capabilities/discovery)
provides methods to query a Kubernetes cluster for the state of its API surface.

`ClusterQueryClient` allows clients to build queries to inspect a cluster and evaluate results.

The sections below illustrate how to build a client and query for APIs and objects.

### Building a ClusterQueryClient

Use the constructor(s) from `discovery` package to get a query client.

```go
import (
    "sigs.k8s.io/controller-runtime/pkg/client/config"
    "github.com/vmware-tanzu/tanzu-framework/pkg/v1/sdk/capabilities/discovery"
)

cfg := config.GetConfig()

clusterQueryClient, err := discovery.NewClusterQueryClientForConfig(cfg)
if err != nil {
    log.Error(err)
}
```

### Building and Executing Queries

Use `Group`, `Object` and `Schema` functions in the `discovery` package to build queries and execute them.

```go
import "github.com/vmware-tanzu/tanzu-framework/pkg/v1/sdk/capabilities/discovery"

// Define objects to query.
var pod = corev1.ObjectReference{
    Kind:       "Pod",
    Name:       "testpod",
    Namespace:  "testns",
    APIVersion: "v1",
}

var testAnnotations = map[string]string{
    "cluster.x-k8s.io/provider": "infrastructure-fake",
}

// Define queries.
var testObject = Object("podObj", &pod).WithAnnotations(testAnnotations)

var testGVR = Group("podResource", testapigroup.SchemeGroupVersion.Group).WithVersions("v1").WithResource("pods")

// Build query client.
c := clusterQueryClient.Query(testObject, testGVR)

// Execute returns combined result of all queries.
found, err := c.Execute()
if err != nil {
    log.Error(err)
}

if found {
    log.Info("Queries successful")
}

// Inspect granular results of each query using the Results method (should be called after Execute).
if result := c.Results().ForQuery("podResource"); result != nil {
    if result.Found {
        log.Info("Pod resource found")
    } else {
        log.Infof("Pod resource not found. Reason: %s", result.NotFoundReason)
    }
}
```

## Executing Pre-defined TKG queries

The `pkg/v1/sdk/capabilities/discovery/tkg` package builds on top of the generic discovery package and exposes
pre-defined queries to determine a TKG cluster's capabilities.

Some examples are shown below.

```go
import tkgdiscovery "github.com/vmware-tanzu/tanzu-framework/pkg/v1/sdk/capabilities/discovery/tkg"

c, err := tkgdiscovery.NewDiscoveryClientForConfig(cfg)
if err != nil {
    log.Fatal(err)
}

if c.IsTKGm() {
    log.Info("This is a TKGm cluster")
}

if c.IsManagementCluster() {
    log.Info("Management cluster")
}

if c.IsWorkloadCluster() {
    log.Info("Workload cluster")
}

if c.HasCloudProvider(ctx, tkgdiscovery.CloudProviderVsphere) {
    log.Info("Cluster has vSphere cloud provider")
}
```

## Capability CRD

Every TKG cluster starting from v1.4.0 includes a `Capability` CRD and an associated controller. Like the Go package
described above, a `Capability` CR can be used to craft queries to inspect a cluster's state and store the results the
CR's `status` field. `Capability` CRD's specification allows for different types of queries to inspect a cluster.

The full API can be found in [apis/run/v1alpha1/capability_types.go](../../apis/run/v1alpha1/capability_types.go)

### Example Capability Custom Resource

The following custom resource checks if the cluster is a TKG cluster which supports feature gating
abilities, and if it has NSX networking capabilities.

```yaml
apiVersion: run.tanzu.vmware.com/v1alpha1
kind: Capability
metadata:
  name: tkg-capabilities
spec:
  queries:
    - name: "tanzu-cluster-with-feature-gating"
      groupVersionResources:
        - name: "tanzu-resource"
          group: "run.tanzu.vmware.com"
          versions:
            - v1alpha1
          resource: "tanzukubernetesreleases"
        - name: "featuregate-resource"
          group: "config.tanzu.vmware.com"
          versions:
            - v1alpha1
          resource: "featuregates"
    - name: "nsx-support"
      objects:
        - name: "nsx-namespace"
          objectReference:
            kind: "Namespace"
            name: "vmware-system-nsx"
            apiVersion: "v1"
```

The capabilities controller:

1. Watches `Capability` resources that are created or updated.
1. Executes queries specified in the spec.
1. Writes the results to the status field of the resource.

After reconciliation, results can be inspected by looking at the status field. Results are grouped by GVK, Object and
Partial Schema queries, and provide a predictable data structure for consumers to parse. They can be accessed by the
paths `status.results.groupVersionResources`, `status.results.objects` and `status.results.partialSchemas` respectively.

An example of query results is shown below.

```yaml
apiVersion: run.tanzu.vmware.com/v1alpha1
kind: Capability
metadata:
  name: tkg-capabilities
spec:
  # Omitted
status:
  results:
  - groupVersionResources:
    - found: true
      name: tanzu-resource
    - found: true
      name: featuregate-resource
    name: tanzu-cluster-with-feature-gating
  - name: nsx-support
    objects:
    - found: false
      name: nsx-namespace
```
