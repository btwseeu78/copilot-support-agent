# Velero Helm Chart - Configuration Reference

## Overview
This document provides a complete reference of all configurable options available in the Velero Helm chart (v12.0.0). All values can be overridden in your `values.yaml` file.

---

## Table of Contents
1. [Namespace Configuration](#namespace-configuration)
2. [Container Image Settings](#container-image-settings)
3. [Deployment & Release Naming](#deployment--release-naming)
4. [Pod & Deployment Metadata](#pod--deployment-metadata)
5. [Resource Management](#resource-management)
6. [Pod Configuration](#pod-configuration)
7. [Pod Scheduling & Affinity](#pod-scheduling--affinity)
8. [Health Checks](#health-checks)
9. [Storage & Volumes](#storage--volumes)
10. [Metrics & Monitoring](#metrics--monitoring)
11. [CRD Upgrade Job](#crd-upgrade-job)

---

## Namespace Configuration

### namespace.labels
Kubernetes labels to apply to the Velero installation namespace. Useful for Pod Security Standards enforcement.

```yaml
namespace:
  labels: {}
    # pod-security.kubernetes.io/enforce: privileged
    # pod-security.kubernetes.io/enforce-version: latest
```

**Type:** `map`  
**Default:** `{}`

---

## Container Image Settings

### image.*
Configure the Velero container image used in the deployment and daemonset (if node-agent is enabled).

```yaml
image:
  repository: docker.io/velero/velero
  tag: v1.18.0
  # Optional: container image digest (takes precedence over tag)
  # digest: sha256:73266e6bd7e2fe3f65fb20677d36c4a8df76d417d9ba76bd9932e0d983a776ec
  pullPolicy: IfNotPresent
  # Kubernetes image pull secrets for private registries
  imagePullSecrets: []
    # - registrySecretName
```

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `image.repository` | string | `docker.io/velero/velero` | Container image repository |
| `image.tag` | string | `v1.18.0` | Container image tag |
| `image.digest` | string | `` | Container image digest (overrides tag) |
| `image.pullPolicy` | string | `IfNotPresent` | Image pull policy (`Always`, `IfNotPresent`, `Never`) |
| `image.imagePullSecrets` | list | `[]` | Image pull secrets for private registries |

---

## Deployment & Release Naming

### nameOverride
Overrides the chart name used in templates. If set, this is used instead of `.Chart.Name`.

```yaml
nameOverride: ""
```

**Type:** `string`  
**Default:** `""`

### fullnameOverride
Overrides the full release name. If set, this is used instead of the generated fullname.

```yaml
fullnameOverride: ""
```

**Type:** `string`  
**Default:** `""`

---

## Pod & Deployment Metadata

### annotations
Annotations to add to the Velero deployment. Useful for triggering rolling updates with reloader.

```yaml
annotations: {}
  # secret.reloader.stakater.com/reload: "<VELERO_SECRET_NAME>"
```

**Type:** `map`  
**Default:** `{}`

### secretAnnotations
Annotations to add to the Velero secret.

```yaml
secretAnnotations: {}
```

**Type:** `map`  
**Default:** `{}`

### labels
Labels to add to the Velero deployment.

```yaml
labels: {}
```

**Type:** `map`  
**Default:** `{}`

### podAnnotations
Annotations to add to the Velero deployment's pod template. Useful for IAM roles (kube2iam, kiam).

```yaml
podAnnotations: {}
  # iam.amazonaws.com/role: "arn:aws:iam::<AWS_ACCOUNT_ID>:role/<VELERO_ROLE_NAME>"
```

**Type:** `map`  
**Default:** `{}`

### podLabels
Additional labels for the Velero deployment's pod template.

```yaml
podLabels: {}
```

**Type:** `map`  
**Default:** `{}`

---

## Resource Management

### resources
CPU and memory requests and limits for the Velero deployment container.

```yaml
resources: {}
  # requests:
  #   cpu: 500m
  #   memory: 128Mi
  # limits:
  #   cpu: 1000m
  #   memory: 512Mi
```

| Key | Type | Default |
|-----|------|---------|
| `resources.requests.cpu` | string | `` |
| `resources.requests.memory` | string | `` |
| `resources.limits.cpu` | string | `` |
| `resources.limits.memory` | string | `` |

### resizePolicy
Container resize policy for the Velero deployment. Requires Kubernetes 1.27+.

```yaml
resizePolicy: []
  # - resourceName: cpu
  #   restartPolicy: NotRequired
  # - resourceName: memory
  #   restartPolicy: RestartContainer
```

**Type:** `list`  
**Default:** `[]`

### upgradeJobResources
CPU and memory requests and limits for the upgradeCRDs job pod.

```yaml
upgradeJobResources: {}
  # requests:
  #   cpu: 50m
  #   memory: 128Mi
  # limits:
  #   cpu: 100m
  #   memory: 256Mi
```

**Type:** `map`  
**Default:** `{}`

### hostAliases
Configure custom host aliases for the Velero pod.

```yaml
hostAliases: []
  # - ip: "127.0.0.1"
  #   hostnames:
  #     - "foo.local"
  #     - "bar.local"
```

**Type:** `list`  
**Default:** `[]`

---

## Pod Configuration

### dnsPolicy
DNS policy for the Velero pod.

```yaml
dnsPolicy: ClusterFirst
```

**Type:** `string`  
**Default:** `ClusterFirst`

### initContainers
Init containers to run before the main Velero container.

```yaml
initContainers: []
```

**Type:** `list`  
**Default:** `[]`

### podSecurityContext
Security context for the entire Velero pod.

```yaml
podSecurityContext: {}
  # runAsUser: 1000
  # runAsGroup: 3000
  # fsGroup: 2000
```

**Type:** `map`  
**Default:** `{}`

### containerSecurityContext
Security context for the Velero container.

```yaml
containerSecurityContext: {}
  # allowPrivilegeEscalation: false
  # readOnlyRootFilesystem: true
  # capabilities:
  #   drop:
  #     - ALL
```

**Type:** `map`  
**Default:** `{}`

### lifecycle
Container lifecycle hooks (postStart, preStop).

```yaml
lifecycle: {}
  # preStop:
  #   exec:
  #     command: ["/bin/sh", "-c", "sleep 15"]
```

**Type:** `map`  
**Default:** `{}`

### priorityClassName
Priority class name for pod scheduling priority.

```yaml
priorityClassName: ""
```

**Type:** `string`  
**Default:** `""`

### runtimeClassName
Runtime class for the pod.

```yaml
runtimeClassName: ""
```

**Type:** `string`  
**Default:** `""`

### terminationGracePeriodSeconds
Grace period in seconds for pod termination.

```yaml
terminationGracePeriodSeconds: 3600
```

**Type:** `integer`  
**Default:** `3600`

---

## Health Checks

### livenessProbe
Liveness probe configuration to determine if Velero is alive.

```yaml
livenessProbe:
  httpGet:
    path: /metrics
    port: http-monitoring
    scheme: HTTP
  initialDelaySeconds: 10
  periodSeconds: 30
  timeoutSeconds: 5
  successThreshold: 1
  failureThreshold: 5
```

**Type:** `map`

### readinessProbe
Readiness probe configuration to determine if Velero is ready to serve traffic.

```yaml
readinessProbe:
  httpGet:
    path: /metrics
    port: http-monitoring
    scheme: HTTP
  initialDelaySeconds: 10
  periodSeconds: 30
  timeoutSeconds: 5
  successThreshold: 1
  failureThreshold: 5
```

**Type:** `map`

---

## Pod Scheduling & Affinity

### tolerations
Node tolerations for the Velero pod, allowing scheduling on tainted nodes.

```yaml
tolerations: []
  # - key: "key1"
  #   operator: "Equal"
  #   value: "value1"
  #   effect: "NoSchedule"
```

**Type:** `list`  
**Default:** `[]`

### affinity
Pod affinity and anti-affinity rules for the Velero pod.

```yaml
affinity: {}
  # podAntiAffinity:
  #   requiredDuringSchedulingIgnoredDuringExecution:
  #     - labelSelector:
  #         matchExpressions:
  #           - key: app
  #             operator: In
  #             values:
  #               - velero
```

**Type:** `map`  
**Default:** `{}`

### nodeSelector
Node selector labels for scheduling Velero on specific nodes.

```yaml
nodeSelector: {}
  # workload-type: backup
```

**Type:** `map`  
**Default:** `{}`

### dnsConfig
DNS configuration for the Velero pod.

```yaml
dnsConfig: {}
  # nameservers:
  #   - 8.8.8.8
  # options:
  #   - name: ndots
  #     value: "2"
```

**Type:** `map`  
**Default:** `{}`

---

## Storage & Volumes

### extraVolumes
Additional volumes to attach to the Velero pod.

```yaml
extraVolumes: []
  # - name: extra-config
  #   configMap:
  #     name: extra-config-map
```

**Type:** `list`  
**Default:** `[]`

### extraVolumeMounts
Additional volume mounts in the Velero container.

```yaml
extraVolumeMounts: []
  # - name: extra-config
  #   mountPath: /etc/extra-config
  #   readOnly: true
```

**Type:** `list`  
**Default:** `[]`

### extraObjects
Additional Kubernetes objects to create (e.g., ConfigMaps, Secrets, NetworkPolicies).

```yaml
extraObjects: []
  # - apiVersion: v1
  #   kind: ConfigMap
  #   metadata:
  #     name: extra-config
  #   data:
  #     key: value
```

**Type:** `list`  
**Default:** `[]`

---

## Metrics & Monitoring

### metrics.enabled
Enable Prometheus metrics exposure.

```yaml
metrics:
  enabled: true
```

**Type:** `boolean`  
**Default:** `true`

### metrics.scrapeInterval
Prometheus scrape interval for Velero metrics.

```yaml
metrics:
  scrapeInterval: 30s
```

**Type:** `string`  
**Default:** `30s`

### metrics.scrapeTimeout
Prometheus scrape timeout for Velero metrics.

```yaml
metrics:
  scrapeTimeout: 10s
```

**Type:** `string`  
**Default:** `10s`

### metrics.service
Kubernetes service configuration for metrics endpoint.

```yaml
metrics:
  service:
    annotations: {}
    type: ClusterIP
    labels: {}
    nodePort: null
    externalTrafficPolicy: ""
    internalTrafficPolicy: ""
    ipFamilyPolicy: ""
    ipFamilies: []
```

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `metrics.service.type` | string | `ClusterIP` | Service type (`ClusterIP`, `NodePort`, `LoadBalancer`) |
| `metrics.service.nodePort` | integer | `null` | NodePort when type is NodePort |
| `metrics.service.externalTrafficPolicy` | string | `` | External traffic policy |
| `metrics.service.internalTrafficPolicy` | string | `` | Internal traffic policy |

### metrics.podAnnotations
Pod annotations for Prometheus scraping.

```yaml
metrics:
  podAnnotations:
    prometheus.io/scrape: "true"
    prometheus.io/port: "8085"
    prometheus.io/path: "/metrics"
```

**Type:** `map`

### metrics.serviceMonitor
Prometheus ServiceMonitor configuration for automated scraping.

```yaml
metrics:
  serviceMonitor:
    autodetect: true
    enabled: false
    annotations: {}
    additionalLabels: {}
```

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `metrics.serviceMonitor.autodetect` | boolean | `true` | Auto-detect CRD availability |
| `metrics.serviceMonitor.enabled` | boolean | `false` | Enable ServiceMonitor |
| `metrics.serviceMonitor.annotations` | map | `{}` | ServiceMonitor annotations |
| `metrics.serviceMonitor.additionalLabels` | map | `{}` | ServiceMonitor labels |

---

## CRD Upgrade Job

### upgradeCRDsJob.*
Configuration for the Kubernetes CRD upgrade job that runs during Helm installation/upgrade.

```yaml
upgradeCRDsJob:
  extraVolumes: []
  extraVolumeMounts: []
  extraEnvVars: []
  automountServiceAccountToken: true
```

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `upgradeCRDsJob.extraVolumes` | list | `[]` | Additional volumes for upgrade job |
| `upgradeCRDsJob.extraVolumeMounts` | list | `[]` | Additional volume mounts for upgrade job |
| `upgradeCRDsJob.extraEnvVars` | list | `[]` | Additional environment variables |
| `upgradeCRDsJob.automountServiceAccountToken` | boolean | `true` | Mount service account token |

---

## Example: Complete Override

```yaml
# Minimal: Override only image
image:
  repository: myregistry.azurecr.io/velero
  tag: v1.18.0
  pullPolicy: Always

---

# Complete: Override multiple sections
image:
  repository: docker.io/velero/velero
  tag: v1.18.0
  pullPolicy: IfNotPresent

resources:
  requests:
    cpu: 250m
    memory: 512Mi
  limits:
    cpu: 1000m
    memory: 1Gi

podAnnotations:
  iam.amazonaws.com/role: "arn:aws:iam::123456789012:role/velero-role"

nodeSelector:
  workload-type: backup

tolerations:
  - key: workload-type
    operator: Equal
    value: backup
    effect: NoSchedule

metrics:
  enabled: true
  serviceMonitor:
    enabled: true
    additionalLabels:
      prometheus: kube-prometheus
```

---

## References
- [Velero Documentation](https://velero.io/docs/)
- [Helm Charts Repository](https://github.com/vmware-tanzu/helm-charts)
- [Kubernetes Pod Configuration](https://kubernetes.io/docs/tasks/configure-pod-container/)

