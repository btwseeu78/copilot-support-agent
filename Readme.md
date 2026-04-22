# Velero Helm Chart — Upstream Values Override Reference

This document lists all configurable values from the upstream `velero` chart **v12.0.0** (and sub-dependencies) that you can override in your local `values.yaml`.

---

## Dependency: velero (upstream v12.0.0)

> Repository: https://vmware-tanzu.github.io/helm-charts  
> Override prefix in local values.yaml: `velero:`

### Namespace

| Key | Type | Default | Description |
|---|---|---|---|
| `velero.namespace.labels` | map | `{}` | Labels applied to the Velero namespace (e.g., Pod Security Standards) |

### Image

| Key | Type | Default | Description |
|---|---|---|---|
| `velero.image.repository` | string | `docker.io/velero/velero` | Container image repository |
| `velero.image.tag` | string | `v1.18.0` | Container image tag |
| `velero.image.pullPolicy` | string | `IfNotPresent` | Image pull policy |
| `velero.image.imagePullSecrets` | list | `[]` | Image pull secrets |

### General Deployment

| Key | Type | Default | Description |
|---|---|---|---|
| `velero.nameOverride` | string | `""` | Override chart name |
| `velero.fullnameOverride` | string | `""` | Override full resource name |
| `velero.annotations` | map | `{}` | Annotations for the Velero Deployment |
| `velero.secretAnnotations` | map | `{}` | Annotations for the Velero Secret |
| `velero.labels` | map | `{}` | Labels for the Velero Deployment |
| `velero.podAnnotations` | map | `{}` | Annotations for the Velero pod template |
| `velero.podLabels` | map | `{}` | Additional labels for the Velero pod template |
| `velero.resources` | map | `{}` | Resource requests/limits for the Velero deployment |
| `velero.resizePolicy` | list | `[]` | Container resize policy (cpu/memory restart policy) |
| `velero.hostAliases` | list | `[]` | Custom `/etc/hosts` entries |
| `velero.dnsPolicy` | string | `ClusterFirst` | DNS policy for Velero pods |
| `velero.dnsConfig` | map | `{}` | DNS configuration for Velero pods |
| `velero.priorityClassName` | string | `""` | Pod priority class name |
| `velero.runtimeClassName` | string | `""` | Pod runtime class name |
| `velero.terminationGracePeriodSeconds` | int | `3600` | Graceful termination period |
| `velero.podSecurityContext` | map | `{}` | Pod-level security context |
| `velero.containerSecurityContext` | map | `{}` | Container-level security context |
| `velero.lifecycle` | map | `{}` | Container lifecycle hooks |
| `velero.tolerations` | list | `[]` | Tolerations for the Velero deployment |
| `velero.affinity` | map | `{}` | Affinity rules |
| `velero.nodeSelector` | map | `{}` | Node selector |
| `velero.extraVolumes` | list | `[]` | Extra volumes |
| `velero.extraVolumeMounts` | list | `[]` | Extra volume mounts |
| `velero.extraObjects` | list | `[]` | Extra Kubernetes manifests to deploy |

### Liveness / Readiness Probes

| Key | Type | Default | Description |
|---|---|---|---|
| `velero.livenessProbe` | map | See values.yaml | Liveness probe configuration |
| `velero.readinessProbe` | map | See values.yaml | Readiness probe configuration |

### Init Containers

| Key | Type | Default | Description |
|---|---|---|---|
| `velero.initContainers` | list | `[]` | Init containers — at least one plugin image required (e.g., velero-plugin-for-gcp) |

### CRD Upgrade Job

| Key | Type | Default | Description |
|---|---|---|---|
| `velero.upgradeCRDs` | bool | `true` | Run a job to upgrade CRDs on install/upgrade |
| `velero.cleanUpCRDs` | bool | `false` | Run a job to clean up CRDs (use only on CI, destructive on production) |
| `velero.upgradeJobResources` | map | `{}` | Resource requests/limits for the CRD upgrade job pod |
| `velero.upgradeCRDsJob.extraVolumes` | list | `[]` | Extra volumes for the CRD upgrade job |
| `velero.upgradeCRDsJob.extraVolumeMounts` | list | `[]` | Extra volume mounts for the CRD upgrade job |
| `velero.upgradeCRDsJob.extraEnvVars` | list | `[]` | Extra environment variables for the CRD upgrade job |
| `velero.upgradeCRDsJob.automountServiceAccountToken` | bool | `true` | Automount service account token in the CRD upgrade job |

### kubectl (upgrade/cleanup jobs)

| Key | Type | Default | Description |
|---|---|---|---|
| `velero.kubectl.image.repository` | string | `registry.k8s.io/kubectl` | kubectl image repository |
| `velero.kubectl.containerSecurityContext` | map | `{}` | Container security context for kubectl container |
| `velero.kubectl.resources` | map | `{}` | Resource requests/limits for kubectl job |
| `velero.kubectl.annotations` | map | `{}` | Annotations for the upgrade/cleanup job |
| `velero.kubectl.labels` | map | `{}` | Labels for the upgrade/cleanup job |
| `velero.kubectl.extraVolumes` | list | `[]` | Extra volumes for the upgrade/cleanup job |
| `velero.kubectl.extraVolumeMounts` | list | `[]` | Extra volume mounts for the upgrade/cleanup job |

### Metrics

| Key | Type | Default | Description |
|---|---|---|---|
| `velero.metrics.enabled` | bool | `true` | Enable Prometheus metrics |
| `velero.metrics.scrapeInterval` | string | `30s` | Prometheus scrape interval |
| `velero.metrics.scrapeTimeout` | string | `10s` | Prometheus scrape timeout |
| `velero.metrics.service.annotations` | map | `{}` | Annotations for the metrics Service |
| `velero.metrics.service.type` | string | `ClusterIP` | Metrics Service type |
| `velero.metrics.service.labels` | map | `{}` | Labels for the metrics Service |
| `velero.metrics.service.nodePort` | int | `null` | NodePort for the metrics Service |
| `velero.metrics.service.externalTrafficPolicy` | string | `""` | External traffic policy |
| `velero.metrics.service.internalTrafficPolicy` | string | `""` | Internal traffic policy |
| `velero.metrics.service.ipFamilyPolicy` | string | `""` | IP family policy (dual-stack) |
| `velero.metrics.service.ipFamilies` | list | `[]` | IP families list |
| `velero.metrics.podAnnotations` | map | Prometheus scrape annotations | Pod annotations for Prometheus scraping |
| `velero.metrics.serviceMonitor.autodetect` | bool | `true` | Auto-detect Prometheus Operator |
| `velero.metrics.serviceMonitor.enabled` | bool | `false` | Create a ServiceMonitor resource |
| `velero.metrics.serviceMonitor.annotations` | map | `{}` | Annotations for the ServiceMonitor |
| `velero.metrics.serviceMonitor.additionalLabels` | map | `{}` | Additional labels for the ServiceMonitor |
| `velero.metrics.nodeAgentPodMonitor.autodetect` | bool | `true` | Auto-detect Prometheus Operator for node-agent |
| `velero.metrics.nodeAgentPodMonitor.enabled` | bool | `false` | Create a PodMonitor for node-agent |
| `velero.metrics.nodeAgentPodMonitor.annotations` | map | `{}` | Annotations for the node-agent PodMonitor |
| `velero.metrics.nodeAgentPodMonitor.additionalLabels` | map | `{}` | Additional labels for the node-agent PodMonitor |
| `velero.metrics.prometheusRule.autodetect` | bool | `true` | Auto-detect Prometheus Operator for PrometheusRule |
| `velero.metrics.prometheusRule.enabled` | bool | `false` | Create a PrometheusRule resource |
| `velero.metrics.prometheusRule.additionalLabels` | map | `{}` | Additional labels for the PrometheusRule |
| `velero.metrics.prometheusRule.spec` | list | `[]` | Alerting rules spec |

### RBAC

| Key | Type | Default | Description |
|---|---|---|---|
| `velero.rbac.create` | bool | `true` | Create RBAC resources |
| `velero.rbac.clusterAdministrator` | bool | `true` | Create cluster-admin ClusterRoleBinding |
| `velero.rbac.clusterAdministratorName` | string | `cluster-admin` | Name of the ClusterRole |

### Service Account

| Key | Type | Default | Description |
|---|---|---|---|
| `velero.serviceAccount.server.create` | bool | `true` | Create the Velero ServiceAccount |
| `velero.serviceAccount.server.name` | string | _(release name)_ | Name of the ServiceAccount |
| `velero.serviceAccount.server.annotations` | map | `nil` | Annotations for the ServiceAccount |
| `velero.serviceAccount.server.labels` | map | `nil` | Labels for the ServiceAccount |
| `velero.serviceAccount.server.imagePullSecrets` | list | `[]` | Image pull secrets for the ServiceAccount |
| `velero.serviceAccount.server.automountServiceAccountToken` | bool | `true` | Automount service account token |

### Credentials

| Key | Type | Default | Description |
|---|---|---|---|
| `velero.credentials.useSecret` | bool | `true` | Use a Kubernetes Secret for provider credentials |
| `velero.credentials.name` | string | _(empty)_ | Name of secret to create |
| `velero.credentials.existingSecret` | string | _(empty)_ | Name of a pre-existing secret |
| `velero.credentials.secretContents` | map | `{}` | Secret data (key: `cloud`) |
| `velero.credentials.extraEnvVars` | map | `{}` | Additional env var key/value pairs stored in the secret |
| `velero.credentials.extraSecretRef` | string | `""` | Pre-existing secret to load env vars from |

### Backup/Snapshot Storage

| Key | Type | Default | Description |
|---|---|---|---|
| `velero.backupsEnabled` | bool | `true` | Create BackupStorageLocation CRD |
| `velero.snapshotsEnabled` | bool | `true` | Create VolumeSnapshotLocation CRD |

### Configuration (server flags)

| Key | Type | Default | Description |
|---|---|---|---|
| `velero.configuration.backupStorageLocation` | list | See values.yaml | BackupStorageLocation(s) — list of objects |
| `velero.configuration.backupStorageLocation[].name` | string | _(empty)_ | Name of the location |
| `velero.configuration.backupStorageLocation[].provider` | string | `""` | Provider name (e.g., `gcp`, `aws`, `azure`) |
| `velero.configuration.backupStorageLocation[].bucket` | string | `""` | Bucket name |
| `velero.configuration.backupStorageLocation[].prefix` | string | _(empty)_ | Path prefix in the bucket |
| `velero.configuration.backupStorageLocation[].default` | bool | `false` | Mark as default location |
| `velero.configuration.backupStorageLocation[].validationFrequency` | string | _(empty)_ | How often to validate |
| `velero.configuration.backupStorageLocation[].accessMode` | string | `ReadWrite` | Access mode |
| `velero.configuration.backupStorageLocation[].credential.name` | string | _(empty)_ | Secret name |
| `velero.configuration.backupStorageLocation[].credential.key` | string | _(empty)_ | Secret key |
| `velero.configuration.backupStorageLocation[].config` | map | `{}` | Provider-specific config |
| `velero.configuration.backupStorageLocation[].annotations` | map | `{}` | Annotations on the BSL resource |
| `velero.configuration.volumeSnapshotLocation` | list | See values.yaml | VolumeSnapshotLocation(s) |
| `velero.configuration.volumeSnapshotLocation[].name` | string | _(empty)_ | Name of the location |
| `velero.configuration.volumeSnapshotLocation[].provider` | string | `""` | Provider name |
| `velero.configuration.volumeSnapshotLocation[].credential.name` | string | _(empty)_ | Secret name |
| `velero.configuration.volumeSnapshotLocation[].credential.key` | string | _(empty)_ | Secret key |
| `velero.configuration.volumeSnapshotLocation[].config` | map | `{}` | Provider-specific config |
| `velero.configuration.volumeSnapshotLocation[].annotations` | map | `{}` | Annotations on the VSL resource |
| `velero.configuration.uploaderType` | string | `kopia` | Uploader type |
| `velero.configuration.backupSyncPeriod` | string | `1m` | How often to sync backups |
| `velero.configuration.fsBackupTimeout` | string | `4h` | File-system backup timeout |
| `velero.configuration.clientBurst` | int | `30` | API burst limit |
| `velero.configuration.clientPageSize` | int | `500` | API page size |
| `velero.configuration.clientQPS` | float | `20.0` | API QPS |
| `velero.configuration.defaultBackupStorageLocation` | string | `default` | Default BSL name |
| `velero.configuration.defaultItemOperationTimeout` | string | `4h` | Default item operation timeout |
| `velero.configuration.defaultBackupTTL` | string | `72h` | Default backup TTL |
| `velero.configuration.defaultVolumeSnapshotLocations` | string | _(empty)_ | Default VSL name |
| `velero.configuration.disableControllers` | string | _(empty)_ | Comma-separated list of controllers to disable |
| `velero.configuration.disableInformerCache` | bool | `false` | Disable informer cache |
| `velero.configuration.garbageCollectionFrequency` | string | `1h` | GC frequency |
| `velero.configuration.itemBlockWorkerCount` | int | `1` | Item block worker count |
| `velero.configuration.logFormat` | string | `text` | Log format (`text` or `json`) |
| `velero.configuration.logLevel` | string | `info` | Log level |
| `velero.configuration.metricsAddress` | string | `:8085` | Prometheus metrics address |
| `velero.configuration.pluginDir` | string | `/plugins` | Plugin directory |
| `velero.configuration.profilerAddress` | string | `localhost:6060` | pprof address |
| `velero.configuration.restoreOnlyMode` | bool | `false` | Restore-only mode |
| `velero.configuration.restoreResourcePriorities` | string | _(see upstream)_ | Restore resource order priorities |
| `velero.configuration.storeValidationFrequency` | string | `1m` | BSL validation frequency |
| `velero.configuration.terminatingResourceTimeout` | string | `10m` | Timeout waiting for PVs/namespaces on restore |
| `velero.configuration.defaultSnapshotMoveData` | bool | `false` | Move snapshot data by default |
| `velero.configuration.features` | string | _(empty)_ | Comma-separated feature flags (e.g., `EnableCSI`) |
| `velero.configuration.dataMoverPrepareTimeout` | string | `30m` | Timeout for CSI snapshot volume provisioning |
| `velero.configuration.repositoryMaintenanceJob.repositoryConfigData.name` | string | `velero-repo-maintenance` | ConfigMap name for repo maintenance config |
| `velero.configuration.repositoryMaintenanceJob.repositoryConfigData.global.keepLatestMaintenanceJobs` | int | `3` | Number of latest maintenance jobs to retain globally |
| `velero.configuration.repositoryMaintenanceJob.repositoryConfigData.repositories` | map | `{}` | Per-repository maintenance config |
| `velero.configuration.namespace` | string | `velero` | Velero server namespace |
| `velero.configuration.extraArgs` | list | `[]` | Extra CLI arguments for `velero server` |
| `velero.configuration.extraEnvVars` | list | `[]` | Extra environment variables |
| `velero.configuration.defaultVolumesToFsBackup` | bool | `false` | Back up all pod volumes via file-system backup by default |
| `velero.configuration.defaultRepoMaintainFrequency` | string | _(empty)_ | How often repository maintenance runs |

### node-agent DaemonSet

| Key | Type | Default | Description |
|---|---|---|---|
| `velero.deployNodeAgent` | bool | `false` | Deploy the node-agent DaemonSet |
| `velero.nodeAgent.disableHostPath` | bool | `false` | Disable host path volumes |
| `velero.nodeAgent.podVolumePath` | string | `/var/lib/kubelet/pods` | Pod volume path on host |
| `velero.nodeAgent.pluginVolumePath` | string | `/var/lib/kubelet/plugins` | Plugin volume path on host |
| `velero.nodeAgent.priorityClassName` | string | `""` | Priority class name |
| `velero.nodeAgent.runtimeClassName` | string | `""` | Runtime class name |
| `velero.nodeAgent.resources` | map | `{}` | Resource requests/limits |
| `velero.nodeAgent.resizePolicy` | list | `[]` | Container resize policy |
| `velero.nodeAgent.tolerations` | list | `[]` | Tolerations |
| `velero.nodeAgent.annotations` | map | `{}` | Annotations |
| `velero.nodeAgent.labels` | map | `{}` | Labels |
| `velero.nodeAgent.podLabels` | map | `{}` | Additional pod labels |
| `velero.nodeAgent.useScratchEmptyDir` | bool | `true` | Use emptyDir for `/scratch` |
| `velero.nodeAgent.extraVolumes` | list | `[]` | Extra volumes |
| `velero.nodeAgent.extraVolumeMounts` | list | `[]` | Extra volume mounts |
| `velero.nodeAgent.extraEnvVars` | list | `[]` | Extra environment variables |
| `velero.nodeAgent.extraArgs` | list | `[]` | Extra CLI arguments for node-agent |
| `velero.nodeAgent.dnsPolicy` | string | `ClusterFirst` | DNS policy |
| `velero.nodeAgent.hostAliases` | list | `[]` | Custom `/etc/hosts` entries |
| `velero.nodeAgent.podSecurityContext` | map | `{runAsUser: 0}` | Pod security context |
| `velero.nodeAgent.containerSecurityContext` | map | `{}` | Container security context |
| `velero.nodeAgent.lifecycle` | map | `{}` | Container lifecycle hooks |
| `velero.nodeAgent.nodeSelector` | map | `{}` | Node selector |
| `velero.nodeAgent.affinity` | map | `{}` | Affinity rules |
| `velero.nodeAgent.dnsConfig` | map | `{}` | DNS configuration |
| `velero.nodeAgent.updateStrategy` | map | `{}` | DaemonSet update strategy |

### Schedules & ConfigMaps

| Key | Type | Default | Description |
|---|---|---|---|
| `velero.schedules` | map | `{}` | Backup Schedule resources to create |
| `velero.configMaps` | map | `{}` | ConfigMap resources to create |

---

## Dependency: workload-identity (v0.4.3)

> Override prefix: `workloadidentity:`

| Key | Type | Default | Description |
|---|---|---|---|
| `workloadidentity.global.gcpProjectId` | string | `irn-70571-ope-92` | GCP project hosting the Kubernetes cluster |
| `workloadidentity.global.gsa.create` | bool | `true` | Create the Google Service Account |
| `workloadidentity.global.gsa.name` | string | `wi-k8s` | GSA name |
| `workloadidentity.global.gsa.project` | string | `""` | GCP project hosting the GSA |
| `workloadidentity.global.cnrmNamespace` | string | `""` | Override deployment namespace |
| `workloadidentity.global.abandon` | bool | `false` | Keep GCP resources after deleting K8s resources |
| `workloadidentity.global.ksa.namespace` | string | `""` | Kubernetes ServiceAccount namespace |
| `workloadidentity.global.ksa.name` | string | `default` | Kubernetes ServiceAccount name |
| `workloadidentity.annotations` | map | `{}` | Annotations for chart resources |

---

## Dependency: iam-policy-member (v1.2.0)

> Override prefix: `iampolicymember:`

| Key | Type | Default | Description |
|---|---|---|---|
| `iampolicymember.global.sia` | string | `run` | SIA of the project |
| `iampolicymember.global.irn` | string | `irn-12345` | IRN identifier |
| `iampolicymember.global.environment` | string | `dev` | Environment |
| `iampolicymember.global.context` | string | `dtf` | Context |
| `iampolicymember.global.cnrmDeletionPolicy` | string | `abandon` | CNRM deletion policy |
| `iampolicymember.global.cnrmNamespace` | string | `""` | Override deployment namespace |
| `iampolicymember.labels` | map | `{}` | Labels for chart resources |
| `iampolicymember.annotations` | map | `{}` | Annotations for chart resources |
| `iampolicymember.policyMembers` | list | `[]` | List of GCP IAM policy member bindings |
| `iampolicymember.policyMembers[].name` | string | — | Name of the binding |
| `iampolicymember.policyMembers[].member` | string | — | IAM member (e.g., `serviceAccount:sa@project.iam.gserviceaccount.com`) |
| `iampolicymember.policyMembers[].role` | string | — | IAM role to bind |
| `iampolicymember.policyMembers[].resourceRef` | map | — | Reference to the GCP resource |

---

## Dependency: iam-custom-role (v0.1.1)

> Override prefix: `iamcustomrole:`

| Key | Type | Default | Description |
|---|---|---|---|
| `iamcustomrole.global.gcpProjectId` | string | `irn-XXXXX-ope-YY` | GCP project for the custom role |
| `iamcustomrole.global.gcpOrganisationId` | string | `""` | GCP org ID (for org-level roles) |
| `iamcustomrole.global.skipUnspecifiedFields` | bool | `false` | Skip unspecified fields in K8s resource spec |
| `iamcustomrole.global.sia` | string | `gke` | SIA of the project |
| `iamcustomrole.global.env` | string | `lab` | Environment |
| `iamcustomrole.global.context` | string | `adm` | Context |
| `iamcustomrole.global.cnrmNamespace` | string | `""` | Override deployment namespace |
| `iamcustomrole.global.abandon` | bool | `true` | Keep GCP role after deleting K8s resources |
| `iamcustomrole.customRoleName` | string | `""` | Name of the custom IAM role |
| `iamcustomrole.description` | string | — | Human-readable description |
| `iamcustomrole.title` | string | — | Human-readable title |
| `iamcustomrole.permissions` | list | `[]` | List of IAM permissions to grant |

---

## Dependency: bucket (v0.1.2)

> Override prefix: `gcpbucket:`

| Key | Type | Default | Description |
|---|---|---|---|
| `gcpbucket.global.cnrmNamespace` | string | `""` | Override deployment namespace |
| `gcpbucket.global.location` | string | `EUROPE-WEST1` | Bucket region or multi-region |
| `gcpbucket.global.abandon` | bool | `false` | Keep bucket after deleting K8s resources |
| `gcpbucket.global.skipUnspecifiedFields` | bool | `false` | Skip list field ownership by Config Connector |
| `gcpbucket.global.sia` | string | `gke` | SIA of the project |
| `gcpbucket.global.env` | string | `lab` | Environment |
| `gcpbucket.global.context` | string | `adm` | Context |
| `gcpbucket.global.gcpProjectId` | string | `irn-XXXXX-ope-YY` | GCP project for the bucket |
| `gcpbucket.bucketName` | string | `myfirstbucket` | Unique bucket name |
| `gcpbucket.annotations` | map | `{}` | Annotations for chart resources |
| `gcpbucket.accessControl.uniformBucketLevelAccess` | bool | `false` | Enable uniform bucket-level access |
| `gcpbucket.accessControl.publicAccessPrevention` | string | `inherited` | Public access prevention (`inherited` or `enforced`) |
| `gcpbucket.accessControl.createCloudIamPolicy` | bool | `true` | Create IAM policy via KCC |
| `gcpbucket.accessControl.iamPolicy` | list | See values.yaml | List of IAM bindings for the bucket |
| `gcpbucket.cors` | list | `[]` | CORS configuration |
| `gcpbucket.defaultEventBasedHold` | bool | `false` | Apply event-based hold to new objects |
| `gcpbucket.encryption.enabled` | bool | `false` | Enable customer-managed encryption |
| `gcpbucket.encryption.kccControlled` | bool | `false` | KMS key managed by KCC |
| `gcpbucket.encryption.kmsKeyRefName` | string | `mykms` | Cloud KMS key name |
| `gcpbucket.encryption.kmsKeyRefNamespace` | string | _(empty)_ | KMS key namespace (if KCC-controlled) |
| `gcpbucket.lifecycleRule` | list | `[]` | Bucket lifecycle rules |
| `gcpbucket.logging.enabled` | bool | `false` | Enable bucket access logging |
| `gcpbucket.logging.logBucket` | string | `access_log_gcs_irn70740_lab_adm` | Destination bucket for logs |
| `gcpbucket.logging.logObjectPrefix` | string | _(empty)_ | Prefix for log object names |
| `gcpbucket.requestPays` | bool | `false` | Requester pays |
| `gcpbucket.retentionPolicy.enabled` | bool | `false` | Enable retention policy |
| `gcpbucket.retentionPolicy.isLocked` | bool | `true` | Lock the retention policy |
| `gcpbucket.retentionPolicy.retentionPeriod` | int | `8192` | Retention period in seconds |
| `gcpbucket.storageClass` | string | `STANDARD` | Storage class (`STANDARD`, `NEARLINE`, `COLDLINE`, `ARCHIVE`) |
| `gcpbucket.enableVersioning` | bool | `false` | Enable bucket versioning |
| `gcpbucket.website.enabled` | bool | `false` | Enable static website hosting |
| `gcpbucket.website.mainPageSuffix` | string | `index.html` | Main page filename |
| `gcpbucket.website.notFoundPage` | string | `404.html` | 404 error page filename |
