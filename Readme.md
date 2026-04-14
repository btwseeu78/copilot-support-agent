# Velero Helm Chart — Overridable Values Reference

This document lists all upstream `values.yaml` options available for override across the chart's dependencies.

---

## Dependency: `velero` (upstream chart v12.0.0)

### Namespace

| Key | Default | Description |
|-----|---------|-------------|
| `velero.namespace.labels` | `{}` | Labels to add to the Velero installation namespace. |

### Image

| Key | Default | Description |
|-----|---------|-------------|
| `velero.image.repository` | `docker.io/velero/velero` | Container image repository. |
| `velero.image.tag` | `v1.18.0` | Container image tag. |
| `velero.image.digest` | _(unset)_ | Image digest; takes precedence over tag if set. |
| `velero.image.pullPolicy` | `IfNotPresent` | Image pull policy. |
| `velero.image.imagePullSecrets` | `[]` | List of image pull secret names. |

### Deployment

| Key | Default | Description |
|-----|---------|-------------|
| `velero.nameOverride` | `""` | Override the chart name. |
| `velero.fullnameOverride` | `""` | Override the full release name. |
| `velero.annotations` | `{}` | Annotations for the Velero Deployment. |
| `velero.secretAnnotations` | `{}` | Annotations for the Velero Secret. |
| `velero.labels` | `{}` | Labels for the Velero Deployment. |
| `velero.podAnnotations` | `{}` | Annotations for Velero pods. |
| `velero.podLabels` | `{}` | Additional labels for Velero pods. |
| `velero.revisionHistoryLimit` | _(unset)_ | Number of old ReplicaSets to retain. |
| `velero.resources` | `{}` | CPU/memory requests and limits for Velero. |
| `velero.resizePolicy` | `[]` | Container resize policy. |
| `velero.hostAliases` | `[]` | Host aliases for Velero pods. |
| `velero.upgradeJobResources` | `{}` | Resources for the upgradeCRDs job pod. |
| `velero.upgradeCRDsJob.extraVolumes` | `[]` | Extra volumes for the Upgrade CRDs Job. |
| `velero.upgradeCRDsJob.extraVolumeMounts` | `[]` | Extra volume mounts for the Upgrade CRDs Job. |
| `velero.upgradeCRDsJob.extraEnvVars` | `[]` | Extra environment variables for the Upgrade CRDs Job. |
| `velero.upgradeCRDsJob.automountServiceAccountToken` | `true` | Automount API credential for the Upgrade CRDs Job service account. |
| `velero.dnsPolicy` | `ClusterFirst` | DNS policy for the Velero deployment. |
| `velero.initContainers` | _(unset)_ | Init containers (plugin provider images). |
| `velero.podSecurityContext` | `{}` | Pod-level security context. |
| `velero.containerSecurityContext` | `{}` | Container-level security context. |
| `velero.lifecycle` | `{}` | Container lifecycle hooks. |
| `velero.priorityClassName` | `""` | Pod priority class name. |
| `velero.runtimeClassName` | `""` | Pod runtime class name. |
| `velero.terminationGracePeriodSeconds` | `3600` | Graceful termination period in seconds. |
| `velero.livenessProbe` | _(see values)_ | Liveness probe configuration. |
| `velero.readinessProbe` | _(see values)_ | Readiness probe configuration. |
| `velero.tolerations` | `[]` | Tolerations for the Velero deployment. |
| `velero.affinity` | `{}` | Affinity rules for the Velero deployment. |
| `velero.nodeSelector` | `{}` | Node selector for the Velero deployment. |
| `velero.dnsConfig` | `{}` | DNS configuration for Velero pods. |
| `velero.extraVolumes` | `[]` | Extra volumes. |
| `velero.extraVolumeMounts` | `[]` | Extra volume mounts. |
| `velero.extraObjects` | `[]` | Extra Kubernetes manifests to deploy. |
| `velero.upgradeCRDs` | `true` | Run the CRD upgrade job on install/upgrade. |
| `velero.cleanUpCRDs` | `false` | Run the CRD cleanup job (for CI; destructive in production). |

### Metrics

| Key | Default | Description |
|-----|---------|-------------|
| `velero.metrics.enabled` | `true` | Enable Prometheus metrics. |
| `velero.metrics.scrapeInterval` | `30s` | Metrics scrape interval. |
| `velero.metrics.scrapeTimeout` | `10s` | Metrics scrape timeout. |
| `velero.metrics.service.annotations` | `{}` | Annotations for the metrics Service. |
| `velero.metrics.service.type` | `ClusterIP` | Metrics Service type. |
| `velero.metrics.service.labels` | `{}` | Labels for the metrics Service. |
| `velero.metrics.service.nodePort` | `null` | Node port (when type is NodePort). |
| `velero.metrics.service.externalTrafficPolicy` | `""` | External traffic policy. |
| `velero.metrics.service.internalTrafficPolicy` | `""` | Internal traffic policy. |
| `velero.metrics.service.ipFamilyPolicy` | `""` | IP family policy. |
| `velero.metrics.service.ipFamilies` | `[]` | Supported IP families. |
| `velero.metrics.podAnnotations` | _(prometheus scrape annotations)_ | Pod annotations for Prometheus scraping. |
| `velero.metrics.serviceMonitor.autodetect` | `true` | Auto-detect Prometheus Operator. |
| `velero.metrics.serviceMonitor.enabled` | `false` | Create a ServiceMonitor resource. |
| `velero.metrics.serviceMonitor.annotations` | `{}` | ServiceMonitor annotations. |
| `velero.metrics.serviceMonitor.additionalLabels` | `{}` | Additional ServiceMonitor labels. |
| `velero.metrics.nodeAgentPodMonitor.autodetect` | `true` | Auto-detect Prometheus Operator for node-agent. |
| `velero.metrics.nodeAgentPodMonitor.enabled` | `false` | Create a PodMonitor for node-agent. |
| `velero.metrics.nodeAgentPodMonitor.annotations` | `{}` | PodMonitor annotations. |
| `velero.metrics.nodeAgentPodMonitor.additionalLabels` | `{}` | Additional PodMonitor labels. |
| `velero.metrics.prometheusRule.autodetect` | `true` | Auto-detect Prometheus Operator for rules. |
| `velero.metrics.prometheusRule.enabled` | `false` | Create a PrometheusRule resource. |
| `velero.metrics.prometheusRule.additionalLabels` | `{}` | Additional PrometheusRule labels. |
| `velero.metrics.prometheusRule.spec` | `[]` | Alert rule definitions. |

### kubectl (CRD Job)

| Key | Default | Description |
|-----|---------|-------------|
| `velero.kubectl.image.repository` | `registry.k8s.io/kubectl` | kubectl image repository. |
| `velero.kubectl.image.tag` | _(cluster version)_ | kubectl image tag override. |
| `velero.kubectl.image.digest` | _(unset)_ | kubectl image digest override. |
| `velero.kubectl.containerSecurityContext` | `{}` | Container security context for kubectl jobs. |
| `velero.kubectl.resources` | `{}` | Resource requests/limits for upgrade/cleanup jobs. |
| `velero.kubectl.annotations` | `{}` | Annotations for upgrade/cleanup jobs. |
| `velero.kubectl.labels` | `{}` | Labels for upgrade/cleanup jobs. |
| `velero.kubectl.extraVolumes` | `[]` | Extra volumes for upgrade/cleanup jobs. |
| `velero.kubectl.extraVolumeMounts` | `[]` | Extra volume mounts for upgrade/cleanup jobs. |

### Configuration (Server Flags & Storage Locations)

| Key | Default | Description |
|-----|---------|-------------|
| `velero.configuration.backupStorageLocation` | _(see values)_ | List of BackupStorageLocation configurations. |
| `velero.configuration.volumeSnapshotLocation` | _(see values)_ | List of VolumeSnapshotLocation configurations. |
| `velero.configuration.uploaderType` | _(kopia)_ | Uploader type. |
| `velero.configuration.backupSyncPeriod` | _(1m)_ | Backup sync period. |
| `velero.configuration.fsBackupTimeout` | _(4h)_ | File system backup timeout. |
| `velero.configuration.clientBurst` | _(30)_ | Client burst setting. |
| `velero.configuration.clientPageSize` | _(500)_ | Client page size. |
| `velero.configuration.clientQPS` | _(20.0)_ | Client QPS. |
| `velero.configuration.defaultBackupStorageLocation` | _(default)_ | Default BSL name. |
| `velero.configuration.defaultItemOperationTimeout` | _(4h)_ | Default item operation timeout. |
| `velero.configuration.defaultBackupTTL` | _(72h)_ | Default backup TTL. |
| `velero.configuration.defaultVolumeSnapshotLocations` | _(unset)_ | Default VSL name. |
| `velero.configuration.disableControllers` | _(unset)_ | Comma-separated list of controllers to disable. |
| `velero.configuration.disableInformerCache` | `false` | Disable informer cache. |
| `velero.configuration.garbageCollectionFrequency` | _(1h)_ | Garbage collection frequency. |
| `velero.configuration.itemBlockWorkerCount` | _(1)_ | Item block worker count. |
| `velero.configuration.logFormat` | _(text)_ | Log format (`text` or `json`). |
| `velero.configuration.logLevel` | _(info)_ | Log level. |
| `velero.configuration.metricsAddress` | _(:8085)_ | Address to expose Prometheus metrics. |
| `velero.configuration.pluginDir` | _(/plugins)_ | Plugin directory. |
| `velero.configuration.profilerAddress` | _(localhost:6060)_ | pprof profiler address. |
| `velero.configuration.restoreOnlyMode` | _(false)_ | Enable restore-only mode. |
| `velero.configuration.restoreResourcePriorities` | _(see values)_ | Resource restore order. |
| `velero.configuration.storeValidationFrequency` | _(1m)_ | Storage validation frequency. |
| `velero.configuration.terminatingResourceTimeout` | _(10m)_ | Timeout for terminating PVs/namespaces on restore. |
| `velero.configuration.defaultSnapshotMoveData` | _(false)_ | Move data by default for all snapshots. |
| `velero.configuration.features` | _(unset)_ | Comma-separated feature flags (e.g., `EnableCSI`). |
| `velero.configuration.dataMoverPrepareTimeout` | _(30m)_ | Timeout for CSI snapshot volume provisioning. |
| `velero.configuration.repositoryMaintenanceJob.repositoryConfigData.name` | `velero-repo-maintenance` | ConfigMap name for maintenance job config. |
| `velero.configuration.repositoryMaintenanceJob.repositoryConfigData.global.keepLatestMaintenanceJobs` | `3` | Number of latest maintenance jobs to keep globally. |
| `velero.configuration.repositoryMaintenanceJob.repositoryConfigData.repositories` | `{}` | Per-repository maintenance job configuration. |
| `velero.configuration.namespace` | _(velero)_ | Server namespace. |
| `velero.configuration.extraArgs` | `[]` | Additional CLI arguments for `velero server`. |
| `velero.configuration.extraEnvVars` | `[]` | Additional environment variables. |
| `velero.configuration.defaultVolumesToFsBackup` | _(false)_ | Back up all pod volumes via file system backup by default. |
| `velero.configuration.defaultRepoMaintainFrequency` | _(unset)_ | Default repository maintenance frequency. |

### RBAC & Service Account

| Key | Default | Description |
|-----|---------|-------------|
| `velero.rbac.create` | `true` | Create Velero Role and RoleBinding. |
| `velero.rbac.clusterAdministrator` | `true` | Create ClusterRoleBinding for cluster-admin. |
| `velero.rbac.clusterAdministratorName` | `cluster-admin` | ClusterRole name. |
| `velero.serviceAccount.server.create` | `true` | Create the Velero ServiceAccount. |
| `velero.serviceAccount.server.name` | _(unset)_ | Override ServiceAccount name. |
| `velero.serviceAccount.server.annotations` | _(unset)_ | ServiceAccount annotations. |
| `velero.serviceAccount.server.labels` | _(unset)_ | ServiceAccount labels. |
| `velero.serviceAccount.server.imagePullSecrets` | `[]` | Image pull secrets for the ServiceAccount. |
| `velero.serviceAccount.server.automountServiceAccountToken` | `true` | Automount API token. |

### Credentials

| Key | Default | Description |
|-----|---------|-------------|
| `velero.credentials.useSecret` | `true` | Create a credentials secret. |
| `velero.credentials.name` | _(unset)_ | Override the secret name. |
| `velero.credentials.existingSecret` | _(unset)_ | Use a pre-existing secret. |
| `velero.credentials.secretContents` | `{}` | Secret data (`cloud` key with IAM credentials). |
| `velero.credentials.extraEnvVars` | `{}` | Extra env vars stored in the secret. |
| `velero.credentials.extraSecretRef` | `""` | Pre-existing secret for environment variables. |

### Feature Flags

| Key | Default | Description |
|-----|---------|-------------|
| `velero.backupsEnabled` | `true` | Create a default BackupStorageLocation CRD. |
| `velero.snapshotsEnabled` | `true` | Enable volume snapshot feature. |
| `velero.deployNodeAgent` | `false` | Deploy the node-agent DaemonSet. |

### Node Agent (DaemonSet)

| Key | Default | Description |
|-----|---------|-------------|
| `velero.nodeAgent.disableHostPath` | `false` | Disable host path volumes. |
| `velero.nodeAgent.podVolumePath` | `/var/lib/kubelet/pods` | Path to pod volumes. |
| `velero.nodeAgent.pluginVolumePath` | `/var/lib/kubelet/plugins` | Path to plugin volumes. |
| `velero.nodeAgent.priorityClassName` | `""` | Priority class name. |
| `velero.nodeAgent.runtimeClassName` | `""` | Runtime class name. |
| `velero.nodeAgent.resources` | `{}` | Resource requests/limits. |
| `velero.nodeAgent.resizePolicy` | `[]` | Container resize policy. |
| `velero.nodeAgent.tolerations` | `[]` | Tolerations. |
| `velero.nodeAgent.annotations` | `{}` | Annotations for the DaemonSet. |
| `velero.nodeAgent.labels` | `{}` | Labels for the DaemonSet. |
| `velero.nodeAgent.podLabels` | `{}` | Additional pod labels. |
| `velero.nodeAgent.useScratchEmptyDir` | `true` | Map `/scratch` to emptyDir. |
| `velero.nodeAgent.extraVolumes` | `[]` | Extra volumes. |
| `velero.nodeAgent.extraVolumeMounts` | `[]` | Extra volume mounts. |
| `velero.nodeAgent.extraEnvVars` | `[]` | Extra environment variables. |
| `velero.nodeAgent.extraArgs` | `[]` | Extra CLI arguments for node-agent. |
| `velero.nodeAgent.dnsPolicy` | `ClusterFirst` | DNS policy. |
| `velero.nodeAgent.hostAliases` | `[]` | Host aliases. |
| `velero.nodeAgent.podSecurityContext` | `{runAsUser: 0}` | Pod security context. |
| `velero.nodeAgent.containerSecurityContext` | `{}` | Container security context. |
| `velero.nodeAgent.lifecycle` | `{}` | Lifecycle hooks. |
| `velero.nodeAgent.nodeSelector` | `{}` | Node selector. |
| `velero.nodeAgent.affinity` | `{}` | Affinity rules. |
| `velero.nodeAgent.dnsConfig` | `{}` | DNS configuration. |
| `velero.nodeAgent.updateStrategy` | `{}` | DaemonSet update strategy. |

### Schedules & ConfigMaps

| Key | Default | Description |
|-----|---------|-------------|
| `velero.schedules` | `{}` | Backup schedule definitions. |
| `velero.configMaps` | `{}` | Velero ConfigMap definitions. |

---

## Dependency: `gateway-helm` (upstream chart v1.7.1)

### Global

| Key | Default | Description |
|-----|---------|-------------|
| `gateway-helm.global.imageRegistry` | `""` | Global override for image registry (highest precedence). |
| `gateway-helm.global.imagePullSecrets` | `[]` | Global override for image pull secrets. |
| `gateway-helm.global.images.envoyGateway.image` | `docker.io/envoyproxy/gateway:v1.7.1` | Full envoy-gateway image name. |
| `gateway-helm.global.images.envoyGateway.pullPolicy` | `IfNotPresent` | Pull policy for envoy-gateway image. |
| `gateway-helm.global.images.envoyGateway.pullSecrets` | `[]` | Pull secrets for envoy-gateway image. |
| `gateway-helm.global.images.ratelimit.image` | `docker.io/envoyproxy/ratelimit:c8765e89` | Full ratelimit image name. |
| `gateway-helm.global.images.ratelimit.pullPolicy` | `IfNotPresent` | Pull policy for ratelimit image. |
| `gateway-helm.global.images.ratelimit.pullSecrets` | `[]` | Pull secrets for ratelimit image. |

### Pod Disruption Budget

| Key | Default | Description |
|-----|---------|-------------|
| `gateway-helm.podDisruptionBudget.minAvailable` | `0` | Minimum available pods. |
| `gateway-helm.podDisruptionBudget.maxUnavailable` | _(unset)_ | Maximum unavailable pods. |

### Deployment

| Key | Default | Description |
|-----|---------|-------------|
| `gateway-helm.deployment.annotations` | `{}` | Annotations for the Deployment. |
| `gateway-helm.deployment.envoyGateway.image.repository` | `""` | Image repository (must include registry when global.imageRegistry is also set). |
| `gateway-helm.deployment.envoyGateway.image.tag` | `""` | Image tag. |
| `gateway-helm.deployment.envoyGateway.imagePullPolicy` | `""` | Image pull policy. |
| `gateway-helm.deployment.envoyGateway.imagePullSecrets` | `[]` | Image pull secrets. |
| `gateway-helm.deployment.envoyGateway.resources.limits.memory` | `1024Mi` | Memory limit. |
| `gateway-helm.deployment.envoyGateway.resources.requests.cpu` | `100m` | CPU request. |
| `gateway-helm.deployment.envoyGateway.resources.requests.memory` | `256Mi` | Memory request. |
| `gateway-helm.deployment.envoyGateway.securityContext` | _(see values)_ | Container security context. |
| `gateway-helm.deployment.ports` | _(see values)_ | List of container ports (grpc, ratelimit, wasm, metrics). |
| `gateway-helm.deployment.priorityClassName` | `null` | Priority class name. |
| `gateway-helm.deployment.replicas` | `1` | Number of replicas. |
| `gateway-helm.deployment.pod.affinity` | `{}` | Pod affinity rules. |
| `gateway-helm.deployment.pod.annotations` | _(prometheus scrape annotations)_ | Pod annotations. |
| `gateway-helm.deployment.pod.labels` | `{}` | Pod labels. |
| `gateway-helm.deployment.pod.topologySpreadConstraints` | `[]` | Topology spread constraints. |
| `gateway-helm.deployment.pod.tolerations` | `[]` | Pod tolerations. |
| `gateway-helm.deployment.pod.nodeSelector` | `{}` | Node selector. |

### Service

| Key | Default | Description |
|-----|---------|-------------|
| `gateway-helm.service.trafficDistribution` | `""` | Set to `PreferClose` to route traffic to topologically closest pods. |
| `gateway-helm.service.annotations` | `{}` | Service annotations. |
| `gateway-helm.service.type` | `ClusterIP` | Service type (e.g., `ClusterIP`, `LoadBalancer`). |

### HPA

| Key | Default | Description |
|-----|---------|-------------|
| `gateway-helm.hpa.enabled` | `false` | Enable Horizontal Pod Autoscaler. |
| `gateway-helm.hpa.minReplicas` | `1` | Minimum replicas. |
| `gateway-helm.hpa.maxReplicas` | `1` | Maximum replicas. |
| `gateway-helm.hpa.metrics` | `[]` | HPA metric definitions. |
| `gateway-helm.hpa.behavior` | `{}` | HPA scaling behavior. |

### Config (EnvoyGateway)

| Key | Default | Description |
|-----|---------|-------------|
| `gateway-helm.config.envoyGateway.gateway.controllerName` | `gateway.envoyproxy.io/gatewayclass-controller` | GatewayClass controller name. |
| `gateway-helm.config.envoyGateway.provider.type` | `Kubernetes` | Provider type. |
| `gateway-helm.config.envoyGateway.logging.level.default` | `info` | Default log level. |
| `gateway-helm.config.envoyGateway.extensionApis` | `{}` | Extension API configuration. |

### Misc

| Key | Default | Description |
|-----|---------|-------------|
| `gateway-helm.createNamespace` | `false` | Create the installation namespace. |
| `gateway-helm.kubernetesClusterDomain` | `cluster.local` | Kubernetes cluster domain. |

### Certgen

| Key | Default | Description |
|-----|---------|-------------|
| `gateway-helm.certgen.job.annotations` | `{}` | Annotations for the certgen Job. |
| `gateway-helm.certgen.job.args` | `[]` | Extra args for the certgen Job. |
| `gateway-helm.certgen.job.pod.annotations` | `{}` | Pod annotations for the certgen Job. |
| `gateway-helm.certgen.job.pod.labels` | `{}` | Pod labels for the certgen Job. |
| `gateway-helm.certgen.job.resources` | `{}` | Resource requests/limits for the certgen Job. |
| `gateway-helm.certgen.job.affinity` | `{}` | Affinity rules for the certgen Job. |
| `gateway-helm.certgen.job.tolerations` | `[]` | Tolerations for the certgen Job. |
| `gateway-helm.certgen.job.nodeSelector` | `{}` | Node selector for the certgen Job. |
| `gateway-helm.certgen.job.ttlSecondsAfterFinished` | `30` | TTL before job cleanup. |
| `gateway-helm.certgen.job.securityContext` | _(see values)_ | Container security context for the certgen Job. |
| `gateway-helm.certgen.rbac.annotations` | `{}` | RBAC resource annotations. |
| `gateway-helm.certgen.rbac.labels` | `{}` | RBAC resource labels. |

### Topology Injector

| Key | Default | Description |
|-----|---------|-------------|
| `gateway-helm.topologyInjector.enabled` | `true` | Enable the topology injector. |
| `gateway-helm.topologyInjector.annotations` | `{}` | Annotations for the topology injector. |
