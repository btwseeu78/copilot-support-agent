# Helm Chart Configuration Options

This document describes all configurable options available in the upstream Helm charts for your dependencies.

## velero (v12.0.0)

### Namespace Settings
- `namespace.labels` (map): Labels to add to the Velero namespace for Pod Security Standards

### Image Configuration
- `image.repository`: Container image repository (default: `docker.io/velero/velero`)
- `image.tag`: Container image tag (default: `v1.18.0`)
- `image.digest`: SHA256 digest (optional, overrides tag if set)
- `image.pullPolicy`: Image pull policy (default: `IfNotPresent`)
- `image.imagePullSecrets` (list): Secrets for private registry authentication

### Deployment Configuration
- `nameOverride`: Override release name
- `fullnameOverride`: Override full release name
- `annotations` (map): Annotations for Velero deployment
- `secretAnnotations` (map): Annotations for created secret
- `labels` (map): Labels for Velero deployment
- `podAnnotations` (map): Pod template annotations (e.g., for kube2iam)
- `podLabels` (map): Additional pod labels
- `revisionHistoryLimit`: Rollback history to retain
- `resources` (map): CPU/memory requests and limits
- `resizePolicy` (list): Container resize policies
- `hostAliases` (list): Pod host aliases
- `dnsPolicy`: DNS policy for pod (default: `ClusterFirst`)
- `priorityClassName`: Priority class for pod
- `runtimeClassName`: Runtime class for pod
- `terminationGracePeriodSeconds`: Graceful shutdown timeout (default: 3600)

### Pod Health & Probes
- `livenessProbe` (map): Liveness probe configuration
  - `httpGet.path`: Metrics path (default: `/metrics`)
  - `httpGet.port`: Port name (default: `http-monitoring`)
  - `httpGet.scheme`: HTTP or HTTPS (default: `HTTP`)
  - `initialDelaySeconds`: Initial delay (default: 10)
  - `periodSeconds`: Check interval (default: 30)
  - `timeoutSeconds`: Timeout per check (default: 5)
  - `failureThreshold`: Failures before restart (default: 5)

- `readinessProbe` (map): Readiness probe configuration (same structure as livenessProbe)

### Pod & Node Configuration
- `tolerations` (list): Tolerations for pod scheduling
- `affinity` (map): Pod affinity rules
- `nodeSelector` (map): Node labels for scheduling
- `dnsConfig` (map): DNS configuration for pod
- `initContainers` (list): Init containers to run before main container
- `podSecurityContext` (map): Pod-level security context
- `containerSecurityContext` (map): Container-level security context
- `lifecycle` (map): Container lifecycle hooks
- `extraVolumes` (list): Additional volumes for pod
- `extraVolumeMounts` (list): Additional volume mounts for container
- `extraObjects` (list): Extra Kubernetes manifests to deploy

### Metrics & Monitoring
- `metrics.enabled`: Enable Prometheus metrics (default: `true`)
- `metrics.scrapeInterval`: Scrape interval (default: `30s`)
- `metrics.scrapeTimeout`: Scrape timeout (default: `10s`)
- `metrics.service.type`: Service type (default: `ClusterIP`)
- `metrics.service.annotations` (map): Service annotations
- `metrics.service.labels` (map): Service labels
- `metrics.service.nodePort`: NodePort for service
- `metrics.service.externalTrafficPolicy`: Traffic policy (`Cluster` or `Local`)
- `metrics.service.internalTrafficPolicy`: Internal traffic policy
- `metrics.service.ipFamilyPolicy`: IP family policy for dual-stack
- `metrics.service.ipFamilies` (list): Supported IP families
- `metrics.podAnnotations` (map): Prometheus scrape annotations
- `metrics.serviceMonitor.enabled`: Create ServiceMonitor (default: `false`)
- `metrics.serviceMonitor.annotations` (map): ServiceMonitor annotations
- `metrics.serviceMonitor.additionalLabels` (map): ServiceMonitor labels
- `metrics.serviceMonitor.autodetect`: Auto-detect Prometheus Operator (default: `true`)
- `metrics.nodeAgentPodMonitor.*`: PodMonitor configuration (similar to serviceMonitor)
- `metrics.prometheusRule.enabled`: Create PrometheusRule (default: `false`)
- `metrics.prometheusRule.additionalLabels` (map): PrometheusRule labels
- `metrics.prometheusRule.spec` (list): Alert rules to deploy

### Kubectl Configuration (for CRD jobs)
- `kubectl.image.repository`: Kubectl image repository (default: `registry.k8s.io/kubectl`)
- `kubectl.image.tag`: Kubectl image tag (overrides cluster version if set)
- `kubectl.image.digest`: SHA256 digest (optional, overrides tag)
- `kubectl.containerSecurityContext` (map): Container security context
- `kubectl.resources` (map): Resource requests/limits for job pods
- `kubectl.annotations` (map): Annotations for jobs
- `kubectl.labels` (map): Labels for jobs
- `kubectl.extraVolumes` (list): Additional volumes for job pods
- `kubectl.extraVolumeMounts` (list): Additional volume mounts

### CRD Management
- `upgradeCRDs`: Run upgrade CRDs job (default: `true`)
- `cleanUpCRDs`: Run cleanup CRDs job (default: `false`)
- `upgradeJobResources` (map): Resource requests/limits for upgrade job
- `upgradeCRDsJob.extraVolumes` (list): Extra volumes for upgrade job
- `upgradeCRDsJob.extraVolumeMounts` (list): Extra volume mounts for upgrade job
- `upgradeCRDsJob.extraEnvVars` (list): Extra environment variables for upgrade job
- `upgradeCRDsJob.automountServiceAccountToken`: Auto-mount service account token (default: `true`)

### Backup Storage & Snapshots
- `configuration.backupStorageLocation` (list): Backup storage location(s)
  - `name`: Location name (default: `"default"` if omitted)
  - `provider`: Provider name (required, e.g., `aws`, `azure`, `gcp`)
  - `bucket`: Bucket name (required)
  - `caCert`: Base64-encoded CA certificate for TLS verification
  - `prefix`: Directory prefix within bucket
  - `default`: Mark as default location (default: `false`)
  - `validationFrequency`: How often to validate object storage
  - `accessMode`: `ReadWrite` or `ReadOnly` (default: `ReadWrite`)
  - `credential.name`: Secret name containing credentials
  - `credential.key`: Key within secret for credentials
  - `config` (map): Provider-specific configuration
  - `annotations` (map): Resource annotations

- `configuration.volumeSnapshotLocation` (list): Volume snapshot location(s)
  - `name`: Location name (default: `"default"` if omitted)
  - `provider`: Provider name
  - `credential.*`: Credential configuration (same as backupStorageLocation)
  - `config` (map): Provider-specific configuration
  - `annotations` (map): Resource annotations

### Server Settings
- `configuration.uploaderType`: File system uploader (default: `kopia`)
- `configuration.backupSyncPeriod`: Backup sync period (default: `1m`)
- `configuration.fsBackupTimeout`: File system backup timeout (default: `4h`)
- `configuration.clientBurst`: API client burst (default: `30`)
- `configuration.clientPageSize`: API page size (default: `500`)
- `configuration.clientQPS`: API QPS (default: `20.0`)
- `configuration.defaultBackupStorageLocation`: Default backup location name
- `configuration.defaultItemOperationTimeout`: Item operation timeout (default: `4h`)
- `configuration.defaultBackupTTL`: Backup retention time (default: `72h`)
- `configuration.defaultVolumeSnapshotLocations`: Default snapshot location name
- `configuration.disableControllers`: Comma-separated list of controllers to disable
- `configuration.disableInformerCache`: Disable informer cache (default: `false`)
- `configuration.garbageCollectionFrequency`: GC frequency (default: `1h`)
- `configuration.itemBlockWorkerCount`: Worker count for item operations (default: `1`)
- `configuration.logFormat`: Log format: `text` or `json` (default: `text`)
- `configuration.logLevel`: Log level (default: `info`)
- `configuration.metricsAddress`: Metrics endpoint address (default: `:8085`)
- `configuration.pluginDir`: Plugin directory (default: `/plugins`)
- `configuration.profilerAddress`: Profiler endpoint (default: `localhost:6060`)
- `configuration.restoreOnlyMode`: Read-only mode for restore (default: `false`)
- `configuration.restoreResourcePriorities`: Restore order for resource types
- `configuration.storeValidationFrequency`: Backup store validation frequency (default: `1m`)
- `configuration.terminatingResourceTimeout`: Timeout for resource termination (default: `10m`)
- `configuration.defaultSnapshotMoveData`: Enable data movement by default (default: `false`)
- `configuration.features`: Comma-separated feature flags (e.g., `EnableCSI`)
- `configuration.dataMoverPrepareTimeout`: CSI snapshot provisioning timeout (default: `30m`)
- `configuration.namespace`: Velero namespace (default: `velero`)
- `configuration.extraArgs` (list): Additional command-line arguments
- `configuration.extraEnvVars` (list): Additional environment variables
- `configuration.defaultVolumesToFsBackup`: Backup all pod volumes by default (default: `false`)
- `configuration.defaultRepoMaintainFrequency`: Repository maintenance frequency

### Repository Maintenance
- `configuration.repositoryMaintenanceJob.repositoryConfigData.name`: ConfigMap name (default: `velero-repo-maintenance`)
- `configuration.repositoryMaintenanceJob.repositoryConfigData.global.keepLatestMaintenanceJobs`: Keep latest N jobs (default: `3`)
- `configuration.repositoryMaintenanceJob.repositoryConfigData.repositories` (map): Per-repository settings

### RBAC
- `rbac.create`: Create RBAC roles/bindings (default: `true`)
- `rbac.clusterAdministrator`: Create cluster admin role binding (default: `true`)
- `rbac.clusterAdministratorName`: Name of cluster admin role (default: `cluster-admin`)

### Service Account
- `serviceAccount.server.create`: Create service account (default: `true`)
- `serviceAccount.server.name`: Service account name
- `serviceAccount.server.annotations` (map): Service account annotations
- `serviceAccount.server.labels` (map): Service account labels
- `serviceAccount.server.imagePullSecrets` (list): Image pull secrets
- `serviceAccount.server.automountServiceAccountToken`: Auto-mount token (default: `true`)

### Credentials
- `credentials.useSecret`: Create secret for credentials (default: `true`)
- `credentials.name`: Secret name to create
- `credentials.existingSecret`: Use existing secret name
- `credentials.secretContents` (map): Credentials data (key `cloud` required)
- `credentials.extraEnvVars` (map): Additional environment variables from secret
- `credentials.extraSecretRef`: Reference to external secret for env vars

### Feature Toggles
- `backupsEnabled`: Create backup storage location CRD (default: `true`)
- `snapshotsEnabled`: Create volume snapshot location CRD (default: `true`)
- `deployNodeAgent`: Deploy node-agent daemonset (default: `false`)

### Node Agent Configuration
- `nodeAgent.disableHostPath`: Disable host path volumes (default: `false`)
- `nodeAgent.podVolumePath`: Kubelet pods path (default: `/var/lib/kubelet/pods`)
- `nodeAgent.pluginVolumePath`: Plugin path (default: `/var/lib/kubelet/plugins`)
- `nodeAgent.priorityClassName`: Priority class
- `nodeAgent.runtimeClassName`: Runtime class
- `nodeAgent.resources` (map): Resource requests/limits
- `nodeAgent.resizePolicy` (list): Resize policies
- `nodeAgent.tolerations` (list): Tolerations
- `nodeAgent.annotations` (map): Daemonset annotations
- `nodeAgent.labels` (map): Daemonset labels
- `nodeAgent.podLabels` (map): Pod labels
- `nodeAgent.useScratchEmptyDir`: Use emptyDir for /scratch (default: `true`)
- `nodeAgent.extraVolumes` (list): Additional volumes
- `nodeAgent.extraVolumeMounts` (list): Additional volume mounts
- `nodeAgent.extraEnvVars` (list): Additional environment variables
- `nodeAgent.extraArgs` (list): Additional command-line arguments
- `nodeAgent.dnsPolicy`: DNS policy (default: `ClusterFirst`)
- `nodeAgent.hostAliases` (list): Host aliases
- `nodeAgent.podSecurityContext` (map): Pod security context
- `nodeAgent.containerSecurityContext` (map): Container security context
- `nodeAgent.lifecycle` (map): Lifecycle hooks
- `nodeAgent.nodeSelector` (map): Node selector
- `nodeAgent.affinity` (map): Affinity rules
- `nodeAgent.dnsConfig` (map): DNS configuration
- `nodeAgent.updateStrategy` (map): Daemonset update strategy

### Schedules
- `schedules` (map): Backup schedules to create
  - Each schedule contains: `disabled`, `labels`, `annotations`, `schedule` (cron), `template` (backup spec)

### ConfigMaps
- `configMaps` (map): Additional ConfigMaps to create (e.g., file system restore action config)

---

## gateway-helm (v1.7.1)

### Global Settings
- `global.imageRegistry`: Global image registry override
- `global.imagePullSecrets` (list): Global image pull secrets
- `global.images.envoyGateway.image`: Full image name including registry and tag
- `global.images.envoyGateway.pullPolicy`: Image pull policy (default: `IfNotPresent`)
- `global.images.envoyGateway.pullSecrets` (list): Image pull secrets
- `global.images.ratelimit.image`: Ratelimit image
- `global.images.ratelimit.pullPolicy`: Image pull policy (default: `IfNotPresent`)
- `global.images.ratelimit.pullSecrets` (list): Image pull secrets

### Pod Disruption Budget
- `podDisruptionBudget.minAvailable`: Minimum available pods (default: `0`)
- `podDisruptionBudget.maxUnavailable`: Maximum unavailable pods (optional)

### Deployment Configuration
- `deployment.annotations` (map): Deployment annotations
- `deployment.envoyGateway.image.repository`: Image repository override
- `deployment.envoyGateway.image.tag`: Image tag override
- `deployment.envoyGateway.imagePullPolicy`: Image pull policy override
- `deployment.envoyGateway.imagePullSecrets` (list): Image pull secrets override
- `deployment.envoyGateway.resources` (map): CPU/memory requests and limits
  - `resources.requests.cpu`: CPU request (default: `100m`)
  - `resources.requests.memory`: Memory request (default: `256Mi`)
  - `resources.limits.memory`: Memory limit (default: `1024Mi`)
- `deployment.envoyGateway.securityContext` (map): Container security context
  - `allowPrivilegeEscalation`: Allow privilege escalation (default: `false`)
  - `capabilities.drop`: Drop capabilities (default: `["ALL"]`)
  - `privileged`: Privileged container (default: `false`)
  - `runAsNonRoot`: Run as non-root (default: `true`)
  - `runAsUser`: User ID (default: `65532`)
  - `runAsGroup`: Group ID (default: `65532`)
  - `seccompProfile.type`: Seccomp profile (default: `RuntimeDefault`)

- `deployment.ports` (list): Container ports
  - `name`: Port name
  - `port`: Port number
  - `targetPort`: Container target port
  - Default ports: grpc (18000), ratelimit (18001), wasm (18002), metrics (19001)

- `deployment.priorityClassName`: Priority class name
- `deployment.replicas`: Number of replicas (default: `1`)

### Pod Configuration
- `deployment.pod.affinity` (map): Pod affinity rules
- `deployment.pod.annotations` (map): Pod annotations
- `deployment.pod.labels` (map): Pod labels
- `deployment.pod.topologySpreadConstraints` (list): Topology spread constraints
- `deployment.pod.tolerations` (list): Pod tolerations
- `deployment.pod.nodeSelector` (map): Node selector

### Service Configuration
- `service.trafficDistribution`: Traffic distribution policy (`PreferClose` or empty)
- `service.annotations` (map): Service annotations
- `service.type`: Service type (default: `ClusterIP`)

### Horizontal Pod Autoscaler
- `hpa.enabled`: Enable HPA (default: `false`)
- `hpa.minReplicas`: Minimum replicas (default: `1`)
- `hpa.maxReplicas`: Maximum replicas (default: `1`)
- `hpa.metrics` (list): Metrics for scaling
- `hpa.behavior` (map): Scaling behavior

### EnvoyGateway Configuration
- `config.envoyGateway` (map): EnvoyGateway API configuration
  - `gateway.controllerName`: Controller name (default: `gateway.envoyproxy.io/gatewayclass-controller`)
  - `provider.type`: Provider type (default: `Kubernetes`)
  - `logging.level.default`: Log level (default: `info`)
  - `extensionApis` (map): Extension APIs configuration

### Namespace & Domain
- `createNamespace`: Create namespace (default: `false`)
- `kubernetesClusterDomain`: Cluster domain (default: `cluster.local`)

### Certificate Generation
- `certgen.job.annotations` (map): Job annotations
- `certgen.job.args` (list): Job arguments
- `certgen.job.pod.annotations` (map): Pod annotations
- `certgen.job.pod.labels` (map): Pod labels
- `certgen.job.resources` (map): Resource requests/limits
- `certgen.job.affinity` (map): Pod affinity
- `certgen.job.tolerations` (list): Pod tolerations
- `certgen.job.nodeSelector` (map): Node selector
- `certgen.job.ttlSecondsAfterFinished`: TTL after completion (default: `30`)
- `certgen.job.securityContext` (map): Pod security context
- `certgen.rbac.annotations` (map): RBAC annotations
- `certgen.rbac.labels` (map): RBAC labels

### Topology Injector
- `topologyInjector.enabled`: Enable topology injector (default: `true`)
- `topologyInjector.annotations` (map): Deployment annotations
