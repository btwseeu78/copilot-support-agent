# Velero Support Agent — Helm Values Override Reference

This document describes all upstream `values.yaml` options available to override for the two Helm chart dependencies: **velero** (upstream v12.0.0) and **gateway-helm** (upstream v1.7.1).

All keys below should be scoped under their chart prefix in your local `values.yaml`, e.g.:
```yaml
velero:
  image:
    tag: v1.18.0
gateway-helm:
  deployment:
    replicas: 2
```

---

## velero (vmware-tanzu/velero v12.0.0)

### Namespace

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `velero.namespace.labels` | map | `{}` | Labels to apply to the Velero install namespace. Can be used to enforce Pod Security Standards. |

### Image

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `velero.image.repository` | string | `docker.io/velero/velero` | Container image repository. |
| `velero.image.tag` | string | `v1.18.0` | Container image tag. |
| `velero.image.digest` | string | `""` | Image digest; takes precedence over tag if set. |
| `velero.image.pullPolicy` | string | `IfNotPresent` | Image pull policy. |
| `velero.image.imagePullSecrets` | list | `[]` | List of image pull secret names. |

### Deployment

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `velero.nameOverride` | string | `""` | Override chart name. |
| `velero.fullnameOverride` | string | `""` | Override full release name. |
| `velero.annotations` | map | `{}` | Annotations on the Velero Deployment. |
| `velero.secretAnnotations` | map | `{}` | Annotations on the Velero Secret. |
| `velero.labels` | map | `{}` | Labels on the Velero Deployment. |
| `velero.podAnnotations` | map | `{}` | Annotations on the Velero pod template. |
| `velero.podLabels` | map | `{}` | Additional labels on the Velero pod template. |
| `velero.revisionHistoryLimit` | int | (k8s default: 10) | Number of old ReplicaSets to retain. |
| `velero.resources` | map | `{}` | CPU/memory resource requests and limits for the Velero container. |
| `velero.resizePolicy` | list | `[]` | Container resize policy entries (`resourceName`, `restartPolicy`). |
| `velero.hostAliases` | list | `[]` | Host aliases to inject into the pod's `/etc/hosts`. |
| `velero.upgradeJobResources` | map | `{}` | Resource requests/limits for the `upgradeCRDs` job pod. |
| `velero.dnsPolicy` | string | `ClusterFirst` | DNS policy for the Velero deployment pod. |
| `velero.initContainers` | list | `[]` | Init containers (used to install Velero plugins). |
| `velero.podSecurityContext` | map | `{}` | Pod-level security context (e.g. `fsGroup`). |
| `velero.containerSecurityContext` | map | `{}` | Container-level security context for the `velero` container. |
| `velero.lifecycle` | map | `{}` | Container lifecycle hooks. |
| `velero.priorityClassName` | string | `""` | Pod priority class name. |
| `velero.runtimeClassName` | string | `""` | Pod runtime class name. |
| `velero.terminationGracePeriodSeconds` | int | `3600` | Graceful termination period in seconds. |
| `velero.livenessProbe` | map | (HTTP `/metrics`) | Liveness probe configuration. |
| `velero.readinessProbe` | map | (HTTP `/metrics`) | Readiness probe configuration. |
| `velero.tolerations` | list | `[]` | Tolerations for the Velero deployment. |
| `velero.affinity` | map | `{}` | Affinity rules for the Velero deployment. |
| `velero.nodeSelector` | map | `{}` | Node selector for the Velero deployment. |
| `velero.dnsConfig` | map | `{}` | DNS config for the Velero deployment. |
| `velero.extraVolumes` | list | `[]` | Extra volumes for the Velero deployment. |
| `velero.extraVolumeMounts` | list | `[]` | Extra volume mounts for the Velero deployment. |
| `velero.extraObjects` | list | `[]` | Extra Kubernetes manifests to deploy alongside Velero. |

### Upgrade CRDs Job

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `velero.upgradeCRDs` | bool | `true` | Whether to run the CRD upgrade job on install/upgrade. |
| `velero.cleanUpCRDs` | bool | `false` | Whether to clean up CRDs on uninstall. **Destructive in production.** |
| `velero.upgradeCRDsJob.extraVolumes` | list | `[]` | Extra volumes for the upgrade CRDs job. |
| `velero.upgradeCRDsJob.extraVolumeMounts` | list | `[]` | Extra volume mounts for the upgrade CRDs job. |
| `velero.upgradeCRDsJob.extraEnvVars` | list | `[]` | Extra environment variables for the upgrade CRDs job. |
| `velero.upgradeCRDsJob.automountServiceAccountToken` | bool | `true` | Whether to automount the service account token in the upgrade job. |

### kubectl (used by upgrade/cleanup jobs)

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `velero.kubectl.image.repository` | string | `registry.k8s.io/kubectl` | kubectl image repository. |
| `velero.kubectl.image.tag` | string | (cluster k8s version) | kubectl image tag override. |
| `velero.kubectl.image.digest` | string | `""` | kubectl image digest override. |
| `velero.kubectl.containerSecurityContext` | map | `{}` | Security context for the kubectl container. |
| `velero.kubectl.resources` | map | `{}` | Resource requests/limits for kubectl containers. |
| `velero.kubectl.annotations` | map | `{}` | Annotations for the upgrade/cleanup job. |
| `velero.kubectl.labels` | map | `{}` | Labels for the upgrade/cleanup job. |
| `velero.kubectl.extraVolumes` | list | `[]` | Extra volumes for the upgrade/cleanup job. |
| `velero.kubectl.extraVolumeMounts` | list | `[]` | Extra volume mounts for the upgrade/cleanup job. |

### Metrics

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `velero.metrics.enabled` | bool | `true` | Enable Prometheus metrics. |
| `velero.metrics.scrapeInterval` | string | `30s` | Prometheus scrape interval. |
| `velero.metrics.scrapeTimeout` | string | `10s` | Prometheus scrape timeout. |
| `velero.metrics.service.annotations` | map | `{}` | Annotations for the metrics Service. |
| `velero.metrics.service.type` | string | `ClusterIP` | Service type for metrics. |
| `velero.metrics.service.labels` | map | `{}` | Labels for the metrics Service. |
| `velero.metrics.service.nodePort` | int | `null` | Node port (when type is NodePort). |
| `velero.metrics.service.externalTrafficPolicy` | string | `""` | External traffic policy (`Cluster` or `Local`). |
| `velero.metrics.service.internalTrafficPolicy` | string | `""` | Internal traffic policy (`Cluster` or `Local`). |
| `velero.metrics.service.ipFamilyPolicy` | string | `""` | IP family policy for dual-stack. |
| `velero.metrics.service.ipFamilies` | list | `[]` | IP families (`IPv4`, `IPv6`). |
| `velero.metrics.podAnnotations` | map | (prometheus scrape annotations) | Pod annotations for Prometheus scraping. |
| `velero.metrics.serviceMonitor.autodetect` | bool | `true` | Auto-detect if ServiceMonitor CRD is available. |
| `velero.metrics.serviceMonitor.enabled` | bool | `false` | Create a Prometheus ServiceMonitor. |
| `velero.metrics.serviceMonitor.annotations` | map | `{}` | Annotations for the ServiceMonitor. |
| `velero.metrics.serviceMonitor.additionalLabels` | map | `{}` | Additional labels for the ServiceMonitor. |
| `velero.metrics.nodeAgentPodMonitor.autodetect` | bool | `true` | Auto-detect if PodMonitor CRD is available. |
| `velero.metrics.nodeAgentPodMonitor.enabled` | bool | `false` | Create a PodMonitor for node-agent. |
| `velero.metrics.nodeAgentPodMonitor.annotations` | map | `{}` | Annotations for the PodMonitor. |
| `velero.metrics.nodeAgentPodMonitor.additionalLabels` | map | `{}` | Additional labels for the PodMonitor. |
| `velero.metrics.prometheusRule.autodetect` | bool | `true` | Auto-detect if PrometheusRule CRD is available. |
| `velero.metrics.prometheusRule.enabled` | bool | `false` | Create a PrometheusRule. |
| `velero.metrics.prometheusRule.additionalLabels` | map | `{}` | Additional labels for the PrometheusRule. |
| `velero.metrics.prometheusRule.spec` | list | `[]` | Prometheus alerting rules. |

### Configuration (server flags & storage locations)

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `velero.configuration.backupStorageLocation` | list | (one entry) | List of BackupStorageLocation configurations. |
| `velero.configuration.backupStorageLocation[].name` | string | `""` | BSL name (empty = "default"). |
| `velero.configuration.backupStorageLocation[].provider` | string | `""` | Storage provider (e.g. `aws`, `gcp`, `azure`). |
| `velero.configuration.backupStorageLocation[].bucket` | string | `""` | Bucket name. |
| `velero.configuration.backupStorageLocation[].caCert` | string | `""` | Base64-encoded CA bundle for TLS. |
| `velero.configuration.backupStorageLocation[].prefix` | string | `""` | Directory prefix within the bucket. |
| `velero.configuration.backupStorageLocation[].default` | bool | `false` | Whether this is the default BSL. |
| `velero.configuration.backupStorageLocation[].validationFrequency` | string | `""` | How often to validate the storage location. |
| `velero.configuration.backupStorageLocation[].accessMode` | string | `ReadWrite` | Access mode (`ReadWrite` or `ReadOnly`). |
| `velero.configuration.backupStorageLocation[].credential.name` | string | `""` | Secret name for credentials. |
| `velero.configuration.backupStorageLocation[].credential.key` | string | `""` | Key within the secret. |
| `velero.configuration.backupStorageLocation[].config` | map | `{}` | Provider-specific config (region, s3Url, etc.). |
| `velero.configuration.backupStorageLocation[].annotations` | map | `{}` | Annotations on the BSL resource. |
| `velero.configuration.volumeSnapshotLocation` | list | (one entry) | List of VolumeSnapshotLocation configurations. |
| `velero.configuration.volumeSnapshotLocation[].name` | string | `""` | VSL name. |
| `velero.configuration.volumeSnapshotLocation[].provider` | string | `""` | Snapshot provider. |
| `velero.configuration.volumeSnapshotLocation[].credential.name` | string | `""` | Secret name. |
| `velero.configuration.volumeSnapshotLocation[].credential.key` | string | `""` | Secret key. |
| `velero.configuration.volumeSnapshotLocation[].config` | map | `{}` | Provider-specific config. |
| `velero.configuration.volumeSnapshotLocation[].annotations` | map | `{}` | Annotations on the VSL resource. |
| `velero.configuration.uploaderType` | string | `kopia` | Uploader type (kopia or restic). |
| `velero.configuration.backupSyncPeriod` | string | `1m` | How often to sync backup objects from object storage. |
| `velero.configuration.fsBackupTimeout` | string | `4h` | Timeout for file-system backup operations. |
| `velero.configuration.clientBurst` | int | `30` | Maximum burst for k8s client. |
| `velero.configuration.clientPageSize` | int | `500` | Page size for k8s list calls. |
| `velero.configuration.clientQPS` | float | `20.0` | QPS for k8s client. |
| `velero.configuration.defaultBackupStorageLocation` | string | `default` | Name of the default BSL. |
| `velero.configuration.defaultItemOperationTimeout` | string | `4h` | Default timeout for item operations. |
| `velero.configuration.defaultBackupTTL` | string | `72h` | Default backup time-to-live. |
| `velero.configuration.defaultVolumeSnapshotLocations` | string | `""` | Default VSL name(s). |
| `velero.configuration.disableControllers` | string | `""` | Comma-separated list of controllers to disable. |
| `velero.configuration.disableInformerCache` | bool | `false` | Disable informer cache. |
| `velero.configuration.garbageCollectionFrequency` | string | `1h` | How often to run garbage collection. |
| `velero.configuration.itemBlockWorkerCount` | int | `1` | Number of workers for item block processing. |
| `velero.configuration.logFormat` | string | `text` | Log format (`text` or `json`). |
| `velero.configuration.logLevel` | string | `info` | Log level (`debug`, `info`, `warning`, `error`, `fatal`, `panic`). |
| `velero.configuration.metricsAddress` | string | `:8085` | Address to expose Prometheus metrics. |
| `velero.configuration.pluginDir` | string | `/plugins` | Directory containing Velero plugins. |
| `velero.configuration.profilerAddress` | string | `localhost:6060` | Address to expose pprof profiler. |
| `velero.configuration.restoreOnlyMode` | bool | `false` | Disable backup creation; allow restores only. |
| `velero.configuration.restoreResourcePriorities` | string | (long default list) | Comma-separated resource restore order. |
| `velero.configuration.storeValidationFrequency` | string | `1m` | How often to validate storage. |
| `velero.configuration.terminatingResourceTimeout` | string | `10m` | Timeout for terminating PVs and namespaces during restore. |
| `velero.configuration.defaultSnapshotMoveData` | bool | `false` | Move snapshot data by default. |
| `velero.configuration.features` | string | `""` | Comma-separated feature flags (e.g. `EnableCSI`). |
| `velero.configuration.dataMoverPrepareTimeout` | string | `30m` | Timeout for CSI snapshot volume provisioning. |
| `velero.configuration.namespace` | string | `velero` | Velero server namespace. |
| `velero.configuration.extraArgs` | list | `[]` | Additional CLI flags for `velero server`. |
| `velero.configuration.extraEnvVars` | list | `[]` | Extra environment variables for the Velero server. |
| `velero.configuration.defaultVolumesToFsBackup` | bool | `false` | Back up all pod volumes via file system by default. |
| `velero.configuration.defaultRepoMaintainFrequency` | string | `""` | How often repository maintenance runs. |

### Repository Maintenance Job

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `velero.configuration.repositoryMaintenanceJob.repositoryConfigData.name` | string | `velero-repo-maintenance` | Name of the ConfigMap for maintenance settings. |
| `velero.configuration.repositoryMaintenanceJob.repositoryConfigData.global.keepLatestMaintenanceJobs` | int | `3` | Number of latest maintenance jobs to retain globally. |
| `velero.configuration.repositoryMaintenanceJob.repositoryConfigData.repositories` | map | `{}` | Per-repository resource/job config. Keys: `{namespace}-{storageLocation}-{repositoryType}`. |

### RBAC

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `velero.rbac.create` | bool | `true` | Create Role and RoleBinding for Velero. |
| `velero.rbac.clusterAdministrator` | bool | `true` | Create ClusterRoleBinding with admin permissions. |
| `velero.rbac.clusterAdministratorName` | string | `cluster-admin` | Name of the ClusterRole to bind. |

### Service Account

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `velero.serviceAccount.server.create` | bool | `true` | Create a ServiceAccount for Velero. |
| `velero.serviceAccount.server.name` | string | `""` | Name of the ServiceAccount (auto-generated if empty). |
| `velero.serviceAccount.server.annotations` | map | `{}` | Annotations on the ServiceAccount. |
| `velero.serviceAccount.server.labels` | map | `{}` | Labels on the ServiceAccount. |
| `velero.serviceAccount.server.imagePullSecrets` | list | `[]` | Image pull secrets for the ServiceAccount. |
| `velero.serviceAccount.server.automountServiceAccountToken` | bool | `true` | Automount service account token. |

### Credentials

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `velero.credentials.useSecret` | bool | `true` | Create a secret for cloud provider credentials. |
| `velero.credentials.name` | string | `""` | Name of the secret to create. |
| `velero.credentials.existingSecret` | string | `""` | Name of a pre-existing secret to use. |
| `velero.credentials.secretContents` | map | `{}` | Secret data (key: `cloud`, value: credential file contents). |
| `velero.credentials.extraEnvVars` | map | `{}` | Extra key/value pairs stored in the secret as env vars. |
| `velero.credentials.extraSecretRef` | string | `""` | Name of an existing secret to load env vars from. |

### Backup / Snapshot Toggles

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `velero.backupsEnabled` | bool | `true` | Create a default BackupStorageLocation CRD. |
| `velero.snapshotsEnabled` | bool | `true` | Enable volume snapshot feature. |

### Node Agent (DaemonSet)

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `velero.deployNodeAgent` | bool | `false` | Deploy the node-agent DaemonSet. |
| `velero.nodeAgent.disableHostPath` | bool | `false` | Disable host path volumes for node-agent. |
| `velero.nodeAgent.podVolumePath` | string | `/var/lib/kubelet/pods` | Kubelet pods path on the host. |
| `velero.nodeAgent.pluginVolumePath` | string | `/var/lib/kubelet/plugins` | Kubelet plugins path on the host. |
| `velero.nodeAgent.priorityClassName` | string | `""` | Priority class for node-agent pods. |
| `velero.nodeAgent.runtimeClassName` | string | `""` | Runtime class for node-agent pods. |
| `velero.nodeAgent.resources` | map | `{}` | Resource requests/limits for node-agent containers. |
| `velero.nodeAgent.resizePolicy` | list | `[]` | Container resize policy for node-agent. |
| `velero.nodeAgent.tolerations` | list | `[]` | Tolerations for node-agent pods. |
| `velero.nodeAgent.annotations` | map | `{}` | Annotations on the node-agent DaemonSet. |
| `velero.nodeAgent.labels` | map | `{}` | Labels on the node-agent DaemonSet. |
| `velero.nodeAgent.podLabels` | map | `{}` | Additional labels on node-agent pod template. |
| `velero.nodeAgent.useScratchEmptyDir` | bool | `true` | Mount an emptyDir at `/scratch`. |
| `velero.nodeAgent.extraVolumes` | list | `[]` | Extra volumes for node-agent. |
| `velero.nodeAgent.extraVolumeMounts` | list | `[]` | Extra volume mounts for node-agent. |
| `velero.nodeAgent.extraEnvVars` | list | `[]` | Extra environment variables for node-agent. |
| `velero.nodeAgent.extraArgs` | list | `[]` | Extra CLI arguments for node-agent. |
| `velero.nodeAgent.dnsPolicy` | string | `ClusterFirst` | DNS policy for node-agent pods. |
| `velero.nodeAgent.hostAliases` | list | `[]` | Host aliases for node-agent pods. |
| `velero.nodeAgent.podSecurityContext` | map | `{runAsUser: 0}` | Pod-level security context for node-agent. |
| `velero.nodeAgent.containerSecurityContext` | map | `{}` | Container-level security context for node-agent. |
| `velero.nodeAgent.lifecycle` | map | `{}` | Lifecycle hooks for node-agent containers. |
| `velero.nodeAgent.nodeSelector` | map | `{}` | Node selector for node-agent. |
| `velero.nodeAgent.affinity` | map | `{}` | Affinity rules for node-agent. |
| `velero.nodeAgent.dnsConfig` | map | `{}` | DNS config for node-agent pods. |
| `velero.nodeAgent.updateStrategy` | map | `{}` | DaemonSet update strategy. |

### Schedules & ConfigMaps

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `velero.schedules` | map | `{}` | Backup schedule definitions. Each key is a schedule name. |
| `velero.configMaps` | map | `{}` | Velero ConfigMaps to create (e.g. `fs-restore-action-config`). |

---

## gateway-helm (envoyproxy/gateway-helm v1.7.1)

> ⚠️ **Note:** The local `values.yaml` currently overrides `global.image` and `global.service` — neither of these paths exists in upstream. See `STRUCTURAL_DIFF_REPORT.md` for details and correct key paths.

### Global

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `gateway-helm.global.imageRegistry` | string | `""` | Global override for image registry (applies to both envoyGateway and ratelimit images). |
| `gateway-helm.global.imagePullSecrets` | list | `[]` | Global image pull secrets (applies to both images). |
| `gateway-helm.global.images.envoyGateway.image` | string | `docker.io/envoyproxy/gateway:v1.7.1` | Full image reference for EnvoyGateway (registry+repo+tag). |
| `gateway-helm.global.images.envoyGateway.pullPolicy` | string | `IfNotPresent` | Pull policy for EnvoyGateway image. |
| `gateway-helm.global.images.envoyGateway.pullSecrets` | list | `[]` | Pull secrets for EnvoyGateway image. |
| `gateway-helm.global.images.ratelimit.image` | string | `docker.io/envoyproxy/ratelimit:c8765e89` | Full image reference for the ratelimit sidecar. |
| `gateway-helm.global.images.ratelimit.pullPolicy` | string | `IfNotPresent` | Pull policy for ratelimit image. |
| `gateway-helm.global.images.ratelimit.pullSecrets` | list | `[]` | Pull secrets for ratelimit image. |

### Pod Disruption Budget

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `gateway-helm.podDisruptionBudget.minAvailable` | int | `0` | Minimum available pods during disruption. |
| `gateway-helm.podDisruptionBudget.maxUnavailable` | int | _(not set)_ | Maximum unavailable pods during disruption. |

### Deployment

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `gateway-helm.deployment.annotations` | map | `{}` | Annotations on the Deployment. |
| `gateway-helm.deployment.envoyGateway.image.repository` | string | `""` | Repository override for EnvoyGateway image (must include registry if `global.imageRegistry` is also set). |
| `gateway-helm.deployment.envoyGateway.image.tag` | string | `""` | Tag override for EnvoyGateway image. |
| `gateway-helm.deployment.envoyGateway.imagePullPolicy` | string | `""` | Pull policy override for EnvoyGateway. |
| `gateway-helm.deployment.envoyGateway.imagePullSecrets` | list | `[]` | Pull secrets for EnvoyGateway deployment. |
| `gateway-helm.deployment.envoyGateway.resources.limits.memory` | string | `1024Mi` | Memory limit for EnvoyGateway container. |
| `gateway-helm.deployment.envoyGateway.resources.requests.cpu` | string | `100m` | CPU request for EnvoyGateway container. |
| `gateway-helm.deployment.envoyGateway.resources.requests.memory` | string | `256Mi` | Memory request for EnvoyGateway container. |
| `gateway-helm.deployment.envoyGateway.securityContext` | map | (restricted defaults) | Security context for EnvoyGateway container. |
| `gateway-helm.deployment.ports` | list | (grpc/ratelimit/wasm/metrics) | Service ports exposed by the deployment. |
| `gateway-helm.deployment.priorityClassName` | string | `null` | Priority class for the gateway pod. |
| `gateway-helm.deployment.replicas` | int | `1` | Number of EnvoyGateway replicas. |
| `gateway-helm.deployment.pod.affinity` | map | `{}` | Affinity rules for the gateway pod. |
| `gateway-helm.deployment.pod.annotations` | map | (prometheus annotations) | Annotations on the gateway pod. |
| `gateway-helm.deployment.pod.labels` | map | `{}` | Labels on the gateway pod. |
| `gateway-helm.deployment.pod.topologySpreadConstraints` | list | `[]` | Topology spread constraints. |
| `gateway-helm.deployment.pod.tolerations` | list | `[]` | Tolerations for the gateway pod. |
| `gateway-helm.deployment.pod.nodeSelector` | map | `{}` | Node selector for the gateway pod. |

### Service

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `gateway-helm.service.trafficDistribution` | string | `""` | Set to `PreferClose` to prefer topologically closest pods. |
| `gateway-helm.service.annotations` | map | `{}` | Annotations on the gateway Service. |
| `gateway-helm.service.type` | string | `ClusterIP` | Service type (`ClusterIP`, `LoadBalancer`, etc.). |

### Horizontal Pod Autoscaler

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `gateway-helm.hpa.enabled` | bool | `false` | Enable HPA for EnvoyGateway. |
| `gateway-helm.hpa.minReplicas` | int | `1` | Minimum replicas for HPA. |
| `gateway-helm.hpa.maxReplicas` | int | `1` | Maximum replicas for HPA. |
| `gateway-helm.hpa.metrics` | list | `[]` | HPA metric sources. |
| `gateway-helm.hpa.behavior` | map | `{}` | HPA scaling behavior. |

### Config

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `gateway-helm.config.envoyGateway.gateway.controllerName` | string | `gateway.envoyproxy.io/gatewayclass-controller` | GatewayClass controller name. |
| `gateway-helm.config.envoyGateway.provider.type` | string | `Kubernetes` | Provider type. |
| `gateway-helm.config.envoyGateway.logging.level.default` | string | `info` | Default log level. |
| `gateway-helm.config.envoyGateway.extensionApis` | map | `{}` | EnvoyGateway extension API configuration. |

### Misc

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `gateway-helm.createNamespace` | bool | `false` | Create the install namespace. |
| `gateway-helm.kubernetesClusterDomain` | string | `cluster.local` | Kubernetes cluster domain. |

### Certgen

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `gateway-helm.certgen.job.annotations` | map | `{}` | Annotations on the certgen Job. |
| `gateway-helm.certgen.job.args` | list | `[]` | Extra args for the certgen job. |
| `gateway-helm.certgen.job.pod.annotations` | map | `{}` | Annotations on the certgen pod. |
| `gateway-helm.certgen.job.pod.labels` | map | `{}` | Labels on the certgen pod. |
| `gateway-helm.certgen.job.resources` | map | `{}` | Resource requests/limits for the certgen job. |
| `gateway-helm.certgen.job.affinity` | map | `{}` | Affinity for the certgen job. |
| `gateway-helm.certgen.job.tolerations` | list | `[]` | Tolerations for the certgen job. |
| `gateway-helm.certgen.job.nodeSelector` | map | `{}` | Node selector for the certgen job. |
| `gateway-helm.certgen.job.ttlSecondsAfterFinished` | int | `30` | TTL for certgen job cleanup. |
| `gateway-helm.certgen.job.securityContext` | map | (restricted defaults) | Security context for the certgen container. |
| `gateway-helm.certgen.rbac.annotations` | map | `{}` | Annotations on certgen RBAC resources. |
| `gateway-helm.certgen.rbac.labels` | map | `{}` | Labels on certgen RBAC resources. |

### Topology Injector

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `gateway-helm.topologyInjector.enabled` | bool | `true` | Enable the topology injector webhook. |
| `gateway-helm.topologyInjector.annotations` | map | `{}` | Annotations on the topology injector resources. |
