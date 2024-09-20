# GKE Autopilot

## Setting resource requests & limits

GKE Autopilot takes responsibility for managing worker nodes and pools; to ensure there is enough capacity, it uses the resource request & limit values defined in the Pod specs to determine the nodes (and their size) to provision. There are extensive rules which specify what the min, max and default resource values will be under a variety of use cases - a full write up of these can be found here [here](https://cloud.google.com/kubernetes-engine/docs/concepts/autopilot-resource-requests). There is a lot to consider, so we have written this page to provide a quick overview of the main points.

### Bursting vs non-bursting

Different rules apply depending on whether your Autopilot cluster supports bursting, basically any cluster creating using v1.30.2+ supports bursting; for earlier versions it's more complicated and [this page](https://cloud.google.com/kubernetes-engine/docs/how-to/pod-bursting-gke#availability-in-gke) provides an explanation. A cluster that supports bursting provides 2 major advantages over non-burtsing clusters:

1. Smaller resource request CPU and memory values can be specified
1. Resource limit values for CPU and memory can be different to the request values

Clearly both of these mean the ability to operate more efficient clusters.

### Generous defaults

If you don't specify resource request or limit values for a Pod, Autopilot will do it for you. These defaults are quite generous so you should **always** specify at least resource request values.

### Resource restrictions

Autopilot imposes min and max resource request values (which are slightly different for the different compute classes and Daemonsets). If you specify values outside of these ranges Autopilot will automatically adjust them. The different compute classes allow you some control over the underlying compute instance that your workload will run on, and each class has different default, max and min resource request values. For simplicity these notes only focus on the default class: General-purpose.

#### General-purpose class resource constraints

| Resource | Default | Min | Max |
|---|---|---|---|
| CPU | 0.5 | 50m (burst)<br>250m (non-burst) | 30 |
| Memory | 2GiB | 52MiB (burst)<br>512MiB (non-burst) | 110GiB |

**NOTE:** you can't specify just any CPU and memory values between min and max - they have to comply with a CPU:memory ratio. For the General-purpose class that ratio is between 1:1 and 1:6.5. Examples:

| CPU | Min memory | Max memory |
|---|---|---|
| 50m | 52MiB | 325MiB |
| 100m | 103MiB | 650MiB |
| 200m | 205MiB | 1.3GiB |
| 500m | 512MiB | 3.2GiB |

If you do specify incorrect values then the CPU value is adjusted accordingly, e.g. if your request is 200m CPU and 2GiB memory then Autopilot will adjust the CPU value to 308m.

### Resource limits

In clusters that _don't_ support bursting, request and and limit values will always be set to be the same (irrespective of any limits you've defined in the pod spec). The request values will be used if both request and limit values have been specified.

In clusters that _do_ support bursting, Autopilot applies different rules depending on which values have been set. A good rule of thumb: set the limit values to be the same or greater than the request values.