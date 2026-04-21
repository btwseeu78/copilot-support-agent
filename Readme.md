# Helm Chart Override Reference

This document lists all configurable options from the upstream `values.yaml` for each dependency that you can override in your local `values.yaml`.

---

## Dependency: `velero` (upstream v12.0.0)

Override keys must be nested under `velero:` in your local `values.yaml`.

### Namespace

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `velero.namespace.labels` | map | `{}` | Labels to add to the Velero installation namespace. |

### Image

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `velero.image.repository` | string | `docker.io/velero/velero` | Container image repository. |
| `velero.image.tag` | string | `v1.18.0` | Container image tag. |
| `velero.image.digest` | string | `""` | Image digest (takes precedence over tag if set). |
| `velero.image.pullPolicy` | string | `IfNotPresent` | Image pull policy. |
| `velero.image.imagePullSecrets` | list | `[]` | Image pull secrets. |

### Deployment

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `velero.nameOverride` | string | `""` | Override chart name. |
| `velero.fullnameOverride` | string | `""` | Override full chart name. |
| `velero.annotations` | map | `{}` | Annotations for the Velero deployment. |
| `velero.secretAnnotations` | map | `{}` | Annotations for the Velero secret. |
| `velero.labels` | map | `{}` | Labels for the Velero deployment. |
| `velero.podAnnotations` | map | `{}` | Pod template annotations. |
| `velero.podLabels` | map | `{}` | Additional pod labels. |
| `velero.resources` | map | `{}` | Resource requests/limits for the Velero pod. |
| `velero.resizePolicy` | list | `[]` | Container resize policy. |
| `velero.hostAliases` | list | `[]` | Host aliases for the Velero pod. |
| `velero.upgradeJobResources` | map | `{}` | Resource requests/limits for the upgradeCRDs job pod. |
| `velero.dnsPolicy` | string | `ClusterFirst` | DNS policy for the Velero deployment. |
| `velero.initContainers` | list | `[]` | Init containers (at least one plugin image required). |
| `velero.podSecurityContext` | map | `{}` | Pod-level security context. |
| `velero.containerSecurityContext` | map | `{}` | Container-level security context. |
| `velero.lifecycle` | map | `{}` | Container lifecycle hooks. |
| `velero.priorityClassName` | string | `""` | Pod priority class name. |
| `velero.runtimeClassName` | string | `""` | Pod runtime class name. |
| `velero.terminationGracePeriodSeconds` | int | `3600` | Graceful termination period (seconds). |
| `velero.livenessProbe` | map | (see values) | Liveness probe configuration. |
| `velero.readinessProbe` | map | (see values) | Readiness probe configuration. |
| `velero.tolerations` | list | `[]` | Tolerations for the Velero deployment. |
| `velero.affinity` | map | `{}` | Affinity rules for the Velero deployment. |
| `velero.nodeSelector` | map | `{}` | Node selector for the Velero deployment. |
| `velero.dnsConfig` | map | `{}` | DNS configuration. |
| `velero.extraVolumes` | list | `[]` | Extra volumes for the Velero deployment. |
| `velero.extraVolumeMounts` | list | `[]` | Extra volume mounts for the Velero deployment. |
| `velero.extraObjects` | list | `[]` | Extra Kubernetes manifests to deploy. |

### Upgrade CRDs Job

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `velero.upgradeCRDs` | bool | `true` | Enable CRD upgrade job. |
| `velero.cleanUpCRDs` | bool | `false` | Enable CRD cleanup job (destructive on production). |
| `velero.upgradeCRDsJob.extraVolumes` | list | `[]` | Extra volumes for the upgrade CRDs job. |
| `velero.upgradeCRDsJob.extraVolumeMounts` | list | `[]` | Extra volume mounts for the upgrade CRDs job. |
| `velero.upgradeCRDsJob.extraEnvVars` | list | `[]` | Extra env vars for the upgrade CRDs job. |
| `velero.upgradeCRDsJob.automountServiceAccountToken` | bool | `true` | Automount service account token in the upgrade job. |

### Metrics

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `velero.metrics.enabled` | bool | `true` | Enable Prometheus metrics. |
| `velero.metrics.scrapeInterval` | string | `30s` | Metrics scrape interval. |
| `velero.metrics.scrapeTimeout` | string | `10s` | Metrics scrape timeout. |
| `velero.metrics.service.annotations` | map | `{}` | Annotations for the metrics service. |
| `velero.metrics.service.type` | string | `ClusterIP` | Metrics service type. |
| `velero.metrics.service.labels` | map | `{}` | Labels for the metrics service. |
| `velero.metrics.service.nodePort` | int | `null` | Node port for the metrics service. |
| `velero.metrics.service.externalTrafficPolicy` | string | `""` | External traffic policy. |
| `velero.metrics.service.internalTrafficPolicy` | string | `""` | Internal traffic policy. |
| `velero.metrics.service.ipFamilyPolicy` | string | `""` | IP family policy (dual-stack). |
| `velero.metrics.service.ipFamilies` | list | `[]` | IP families for dual-stack. |
| `velero.metrics.podAnnotations` | map | (Prometheus defaults) | Pod annotations for Prometheus scraping. |
| `velero.metrics.serviceMonitor.autodetect` | bool | `true` | Auto-detect ServiceMonitor CRD. |
| `velero.metrics.serviceMonitor.enabled` | bool | `false` | Enable ServiceMonitor. |
| `velero.metrics.serviceMonitor.annotations` | map | `{}` | ServiceMonitor annotations. |
| `velero.metrics.serviceMonitor.additionalLabels` | map | `{}` | Additional ServiceMonitor labels. |
| `velero.metrics.nodeAgentPodMonitor.autodetect` | bool | `true` | Auto-detect PodMonitor CRD for node-agent. |
| `velero.metrics.nodeAgentPodMonitor.enabled` | bool | `false` | Enable PodMonitor for node-agent. |
| `velero.metrics.nodeAgentPodMonitor.annotations` | map | `{}` | PodMonitor annotations. |
| `velero.metrics.nodeAgentPodMonitor.additionalLabels` | map | `{}` | Additional PodMonitor labels. |
| `velero.metrics.prometheusRule.autodetect` | bool | `true` | Auto-detect PrometheusRule CRD. |
| `velero.metrics.prometheusRule.enabled` | bool | `false` | Enable PrometheusRule. |
| `velero.metrics.prometheusRule.additionalLabels` | map | `{}` | Additional PrometheusRule labels. |
| `velero.metrics.prometheusRule.spec` | list | `[]` | PrometheusRule alert spec. |

### kubectl (CRD Jobs)

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `velero.kubectl.image.repository` | string | `registry.k8s.io/kubectl` | kubectl image repository. |
| `velero.kubectl.image.digest` | string | `""` | kubectl image digest. |
| `velero.kubectl.containerSecurityContext` | map | `{}` | Security context for kubectl container. |
| `velero.kubectl.resources` | map | `{}` | Resources for the upgrade/cleanup job. |
| `velero.kubectl.annotations` | map | `{}` | Annotations for the upgrade/cleanup job. |
| `velero.kubectl.labels` | map | `{}` | Labels for the upgrade/cleanup job. |
| `velero.kubectl.extraVolumes` | list | `[]` | Extra volumes for the upgrade/cleanup job. |
| `velero.kubectl.extraVolumeMounts` | list | `[]` | Extra volume mounts for the upgrade/cleanup job. |

### Configuration (Server Flags & Storage Locations)

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `velero.configuration.backupStorageLocation` | list | (see values) | BackupStorageLocation definitions. |
| `velero.configuration.volumeSnapshotLocation` | list | (see values) | VolumeSnapshotLocation definitions. |
| `velero.configuration.uploaderType` | string | `""` | Uploader type (default: kopia). |
| `velero.configuration.backupSyncPeriod` | string | `""` | Backup sync period (default: 1m). |
| `velero.configuration.fsBackupTimeout` | string | `""` | File system backup timeout (default: 4h). |
| `velero.configuration.clientBurst` | int | `""` | Client burst (default: 30). |
| `velero.configuration.clientPageSize` | int | `""` | Client page size (default: 500). |
| `velero.configuration.clientQPS` | float | `""` | Client QPS (default: 20.0). |
| `velero.configuration.defaultBackupStorageLocation` | string | `""` | Default backup storage location name. |
| `velero.configuration.defaultItemOperationTimeout` | string | `""` | Default item operation timeout (default: 4h). |
| `velero.configuration.defaultBackupTTL` | string | `""` | Default backup TTL (default: 72h). |
| `velero.configuration.defaultVolumeSnapshotLocations` | string | `""` | Default volume snapshot location. |
| `velero.configuration.disableControllers` | string | `""` | Comma-separated list of controllers to disable. |
| `velero.configuration.disableInformerCache` | bool | `false` | Disable informer cache. |
| `velero.configuration.garbageCollectionFrequency` | string | `""` | Garbage collection frequency (default: 1h). |
| `velero.configuration.itemBlockWorkerCount` | int | `""` | Item block worker count (default: 1). |
| `velero.configuration.logFormat` | string | `""` | Log format: `text` or `json`. |
| `velero.configuration.logLevel` | string | `""` | Log level: `info`, `debug`, `warning`, `error`, etc. |
| `velero.configuration.metricsAddress` | string | `""` | Prometheus metrics address (default: `:8085`). |
| `velero.configuration.pluginDir` | string | `""` | Plugin directory (default: `/plugins`). |
| `velero.configuration.profilerAddress` | string | `""` | pprof profiler address. |
| `velero.configuration.restoreOnlyMode` | bool | `false` | Enable restore-only mode. |
| `velero.configuration.restoreResourcePriorities` | string | `""` | Restore resource priority order. |
| `velero.configuration.storeValidationFrequency` | string | `""` | Store validation frequency (default: 1m). |
| `velero.configuration.terminatingResourceTimeout` | string | `""` | Terminating resource timeout (default: 10m). |
| `velero.configuration.defaultSnapshotMoveData` | bool | `false` | Move data by default for all snapshots. |
| `velero.configuration.features` | string | `""` | Comma-separated Velero feature flags (e.g. `EnableCSI`). |
| `velero.configuration.dataMoverPrepareTimeout` | string | `""` | CSI snapshot provisioning timeout (default: 30m). |
| `velero.configuration.repositoryMaintenanceJob.repositoryConfigData.name` | string | `velero-repo-maintenance` | ConfigMap name for repo maintenance config. |
| `velero.configuration.repositoryMaintenanceJob.repositoryConfigData.global.keepLatestMaintenanceJobs` | int | `3` | Number of latest maintenance jobs to keep globally. |
| `velero.configuration.repositoryMaintenanceJob.repositoryConfigData.repositories` | map | `{}` | Per-repository maintenance configuration. |
| `velero.configuration.namespace` | string | `""` | Velero server namespace (default: `velero`). |
| `velero.configuration.extraArgs` | list | `[]` | Extra CLI args passed to `velero server`. |
| `velero.configuration.extraEnvVars` | list | `[]` | Extra environment variables for the Velero server. |
| `velero.configuration.defaultVolumesToFsBackup` | bool | `false` | Back up all pod volumes via file system by default. |
| `velero.configuration.defaultRepoMaintainFrequency` | string | `""` | Default repository maintenance frequency. |

### RBAC

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `velero.rbac.create` | bool | `true` | Create Velero Role and RoleBinding. |
| `velero.rbac.clusterAdministrator` | bool | `true` | Create ClusterRoleBinding with admin permissions. |
| `velero.rbac.clusterAdministratorName` | string | `cluster-admin` | ClusterRole name. |

### Service Account

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `velero.serviceAccount.server.create` | bool | `true` | Create the server service account. |
| `velero.serviceAccount.server.name` | string | `""` | Name of the service account. |
| `velero.serviceAccount.server.annotations` | map | `{}` | Annotations for the service account. |
| `velero.serviceAccount.server.labels` | map | `{}` | Labels for the service account. |
| `velero.serviceAccount.server.imagePullSecrets` | list | `[]` | Image pull secrets for the service account. |
| `velero.serviceAccount.server.automountServiceAccountToken` | bool | `true` | Automount service account token. |

### Credentials

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `velero.credentials.useSecret` | bool | `true` | Use a secret for cloud provider credentials. |
| `velero.credentials.name` | string | `""` | Name of the secret to create. |
| `velero.credentials.existingSecret` | string | `""` | Name of a pre-existing secret. |
| `velero.credentials.secretContents` | map | `{}` | Secret data (key: `cloud`). |
| `velero.credentials.extraEnvVars` | map | `{}` | Extra env vars stored as secret data. |
| `velero.credentials.extraSecretRef` | string | `""` | Pre-existing secret to load env vars from. |

### Backup & Snapshots

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `velero.backupsEnabled` | bool | `true` | Create BackupStorageLocation CRD. |
| `velero.snapshotsEnabled` | bool | `true` | Create VolumeSnapshotLocation CRD. |
| `velero.schedules` | map | `{}` | Backup schedules to create. |
| `velero.configMaps` | map | `{}` | Velero ConfigMaps to create. |

### Node Agent (DaemonSet)

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `velero.deployNodeAgent` | bool | `false` | Deploy the node-agent DaemonSet. |
| `velero.nodeAgent.disableHostPath` | bool | `false` | Disable host path volumes for node-agent. |
| `velero.nodeAgent.podVolumePath` | string | `/var/lib/kubelet/pods` | Host path for pod volumes. |
| `velero.nodeAgent.pluginVolumePath` | string | `/var/lib/kubelet/plugins` | Host path for plugin volumes. |
| `velero.nodeAgent.priorityClassName` | string | `""` | Priority class name for node-agent pods. |
| `velero.nodeAgent.runtimeClassName` | string | `""` | Runtime class name for node-agent pods. |
| `velero.nodeAgent.resources` | map | `{}` | Resource requests/limits for node-agent. |
| `velero.nodeAgent.resizePolicy` | list | `[]` | Container resize policy for node-agent. |
| `velero.nodeAgent.tolerations` | list | `[]` | Tolerations for node-agent DaemonSet. |
| `velero.nodeAgent.annotations` | map | `{}` | Annotations for node-agent DaemonSet. |
| `velero.nodeAgent.labels` | map | `{}` | Labels for node-agent DaemonSet. |
| `velero.nodeAgent.podLabels` | map | `{}` | Additional pod labels for node-agent. |
| `velero.nodeAgent.useScratchEmptyDir` | bool | `true` | Use emptyDir for `/scratch`. |
| `velero.nodeAgent.extraVolumes` | list | `[]` | Extra volumes for node-agent. |
| `velero.nodeAgent.extraVolumeMounts` | list | `[]` | Extra volume mounts for node-agent. |
| `velero.nodeAgent.extraEnvVars` | list | `[]` | Extra env vars for node-agent. |
| `velero.nodeAgent.extraArgs` | list | `[]` | Extra CLI args for node-agent. |
| `velero.nodeAgent.dnsPolicy` | string | `ClusterFirst` | DNS policy for node-agent. |
| `velero.nodeAgent.hostAliases` | list | `[]` | Host aliases for node-agent. |
| `velero.nodeAgent.podSecurityContext` | map | `{runAsUser: 0}` | Pod security context for node-agent. |
| `velero.nodeAgent.containerSecurityContext` | map | `{}` | Container security context for node-agent. |
| `velero.nodeAgent.lifecycle` | map | `{}` | Lifecycle hooks for node-agent. |
| `velero.nodeAgent.nodeSelector` | map | `{}` | Node selector for node-agent. |
| `velero.nodeAgent.affinity` | map | `{}` | Affinity for node-agent. |
| `velero.nodeAgent.dnsConfig` | map | `{}` | DNS configuration for node-agent. |
| `velero.nodeAgent.updateStrategy` | map | `{}` | Update strategy for node-agent DaemonSet. |

---

## Dependency: `gateway-helm` (upstream v1.7.1)

Override keys must be nested under `gateway-helm:` in your local `values.yaml`.

> ⚠️ **Important**: The local `values.yaml` currently uses incorrect key paths. See [STRUCTURAL_DIFF_REPORT.md](./STRUCTURAL_DIFF_REPORT.md) for details.

### Global Image Overrides

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `gateway-helm.global.imageRegistry` | string | `""` | Global registry override for all images. |
| `gateway-helm.global.imagePullSecrets` | list | `[]` | Global image pull secrets. |
| `gateway-helm.global.images.envoyGateway.image` | string | `docker.io/envoyproxy/gateway:v1.7.1` | Full envoy-gateway image (registry/repo:tag). |
| `gateway-helm.global.images.envoyGateway.pullPolicy` | string | `IfNotPresent` | Pull policy for envoy-gateway image. |
| `gateway-helm.global.images.envoyGateway.pullSecrets` | list | `[]` | Pull secrets for envoy-gateway image. |
| `gateway-helm.global.images.ratelimit.image` | string | `docker.io/envoyproxy/ratelimit:c8765e89` | Full ratelimit image (registry/repo:tag). |
| `gateway-helm.global.images.ratelimit.pullPolicy` | string | `IfNotPresent` | Pull policy for ratelimit image. |
| `gateway-helm.global.images.ratelimit.pullSecrets` | list | `[]` | Pull secrets for ratelimit image. |

### Pod Disruption Budget

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `gateway-helm.podDisruptionBudget.minAvailable` | int | `0` | Minimum available pods. |
| `gateway-helm.podDisruptionBudget.maxUnavailable` | int | — | Maximum unavailable pods. |

### Deployment

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `gateway-helm.deployment.annotations` | map | `{}` | Deployment annotations. |
| `gateway-helm.deployment.envoyGateway.image.repository` | string | `""` | Per-component image repository override. |
| `gateway-helm.deployment.envoyGateway.image.tag` | string | `""` | Per-component image tag override. |
| `gateway-helm.deployment.envoyGateway.imagePullPolicy` | string | `""` | Per-component image pull policy. |
| `gateway-helm.deployment.envoyGateway.imagePullSecrets` | list | `[]` | Per-component image pull secrets. |
| `gateway-helm.deployment.envoyGateway.resources.limits.memory` | string | `1024Mi` | Memory limit for envoy-gateway. |
| `gateway-helm.deployment.envoyGateway.resources.requests.cpu` | string | `100m` | CPU request for envoy-gateway. |
| `gateway-helm.deployment.envoyGateway.resources.requests.memory` | string | `256Mi` | Memory request for envoy-gateway. |
| `gateway-helm.deployment.envoyGateway.securityContext` | map | (see values) | Security context for envoy-gateway container. |
| `gateway-helm.deployment.ports` | list | (grpc/ratelimit/wasm/metrics) | Service ports exposed by the deployment. |
| `gateway-helm.deployment.priorityClassName` | string | `null` | Priority class name. |
| `gateway-helm.deployment.replicas` | int | `1` | Number of replicas. |
| `gateway-helm.deployment.pod.affinity` | map | `{}` | Pod affinity rules. |
| `gateway-helm.deployment.pod.annotations` | map | (Prometheus defaults) | Pod annotations. |
| `gateway-helm.deployment.pod.labels` | map | `{}` | Pod labels. |
| `gateway-helm.deployment.pod.topologySpreadConstraints` | list | `[]` | Topology spread constraints. |
| `gateway-helm.deployment.pod.tolerations` | list | `[]` | Tolerations for envoy-gateway pods. |
| `gateway-helm.deployment.pod.nodeSelector` | map | `{}` | Node selector for envoy-gateway pods. |

### Service

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `gateway-helm.service.trafficDistribution` | string | `""` | Traffic distribution policy (e.g. `PreferClose`). |
| `gateway-helm.service.annotations` | map | `{}` | Service annotations. |
| `gateway-helm.service.type` | string | `ClusterIP` | Service type. |

### HPA

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `gateway-helm.hpa.enabled` | bool | `false` | Enable Horizontal Pod Autoscaler. |
| `gateway-helm.hpa.minReplicas` | int | `1` | HPA minimum replicas. |
| `gateway-helm.hpa.maxReplicas` | int | `1` | HPA maximum replicas. |
| `gateway-helm.hpa.metrics` | list | `[]` | HPA metrics. |
| `gateway-helm.hpa.behavior` | map | `{}` | HPA scaling behavior. |

### Config

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `gateway-helm.config.envoyGateway.gateway.controllerName` | string | `gateway.envoyproxy.io/gatewayclass-controller` | Gateway controller name. |
| `gateway-helm.config.envoyGateway.provider.type` | string | `Kubernetes` | Provider type. |
| `gateway-helm.config.envoyGateway.logging.level.default` | string | `info` | Default log level. |
| `gateway-helm.config.envoyGateway.extensionApis` | map | `{}` | Extension API configuration. |

### Certgen

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `gateway-helm.certgen.job.annotations` | map | `{}` | Annotations for certgen job. |
| `gateway-helm.certgen.job.args` | list | `[]` | Extra args for certgen job. |
| `gateway-helm.certgen.job.pod.annotations` | map | `{}` | Annotations for certgen job pod. |
| `gateway-helm.certgen.job.pod.labels` | map | `{}` | Labels for certgen job pod. |
| `gateway-helm.certgen.job.resources` | map | `{}` | Resources for certgen job. |
| `gateway-helm.certgen.job.affinity` | map | `{}` | Affinity for certgen job. |
| `gateway-helm.certgen.job.tolerations` | list | `[]` | Tolerations for certgen job. |
| `gateway-helm.certgen.job.nodeSelector` | map | `{}` | Node selector for certgen job. |
| `gateway-helm.certgen.job.ttlSecondsAfterFinished` | int | `30` | TTL after job finishes. |
| `gateway-helm.certgen.job.securityContext` | map | (see values) | Security context for certgen container. |
| `gateway-helm.certgen.rbac.annotations` | map | `{}` | Annotations for certgen RBAC resources. |
| `gateway-helm.certgen.rbac.labels` | map | `{}` | Labels for certgen RBAC resources. |

### Misc

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `gateway-helm.createNamespace` | bool | `false` | Create the namespace during install. |
| `gateway-helm.kubernetesClusterDomain` | string | `cluster.local` | Kubernetes cluster domain. |
| `gateway-helm.topologyInjector.enabled` | bool | `true` | Enable topology injector. |
| `gateway-helm.topologyInjector.annotations` | map | `{}` | Annotations for topology injector. |
