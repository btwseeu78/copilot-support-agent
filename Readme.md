# Helm Chart Values.yaml Configuration Reference

This document provides a comprehensive list of all configurable options from the upstream Helm chart values for Velero (v12.0.0) and EnvoyGateway (v1.7.1) that you can override in your custom values.yaml.

## Table of Contents

- [Velero Chart Configuration](#velero-chart-configuration)
- [EnvoyGateway Chart Configuration](#envoygateway-chart-configuration)

---

## Velero Chart Configuration

The Velero chart configures a backup and disaster recovery solution for Kubernetes.

### Namespace Configuration

```yaml
namespace:
  labels: {}
    # Enforce Pod Security Standards with Namespace Labels
    # pod-security.kubernetes.io/enforce: privileged
    # pod-security.kubernetes.io/enforce-version: latest
    # pod-security.kubernetes.io/audit: privileged
    # pod-security.kubernetes.io/audit-version: latest
    # pod-security.kubernetes.io/warn: privileged
    # pod-security.kubernetes.io/warn-version: latest
```

### Image Configuration

```yaml
image:
  repository: docker.io/velero/velero
  tag: v1.18.0
  # Digest value example: sha256:73266e6bd7e2fe3f65fb20677d36c4a8df76d417d9ba76bd9932e0d983a776ec
  digest: ""  # If used, takes precedence over tag
  pullPolicy: IfNotPresent
  imagePullSecrets: []
    # - registrySecretName
```

### Helm Release Configuration

```yaml
nameOverride: ""
fullnameOverride: ""
```

### Annotations and Labels

```yaml
# Annotations on Velero deployment (reloader example: secret.reloader.stakater.com/reload: "<VELERO_SECRET_NAME>")
annotations: {}

# Annotations on Velero secret
secretAnnotations: {}

# Labels on Velero deployment
labels: {}

# Pod template annotations (for kube2iam/kiam: iam.amazonaws.com/role: "arn:aws:iam::<AWS_ACCOUNT_ID>:role/<VELERO_ROLE_NAME>")
podAnnotations: {}

# Pod template labels
podLabels: {}
```

### Deployment Configuration

```yaml
# Revision history limit for rollback
revisionHistoryLimit: 10  # Default Kubernetes value

# Resource requests and limits
resources: {}
  # requests:
  #   cpu: 500m
  #   memory: 128Mi
  # limits:
  #   cpu: 1000m
  #   memory: 512Mi

# Container resize policy
resizePolicy: []
  # - resourceName: cpu
  #   restartPolicy: NotRequired
  # - resourceName: memory
  #   restartPolicy: RestartContainer

# Configure hostAliases for custom DNS
hostAliases: []
  # - ip: "127.0.0.1"
  #   hostnames:
  #     - "foo.local"
  #     - "bar.local"

# DNS policy
dnsPolicy: ClusterFirst

# DNS configuration
dnsConfig: {}

# Init containers (at least one plugin provider image recommended)
initContainers: []
  # - name: velero-plugin-for-aws
  #   image: velero/velero-plugin-for-aws:v1.13.1
  #   imagePullPolicy: IfNotPresent
  #   volumeMounts:
  #     - mountPath: /target
  #       name: plugins

# Pod security context (fsGroup recommended for AWS IAM Roles)
podSecurityContext: {}
  # fsGroup: 1337

# Container security context
containerSecurityContext: {}
  # allowPrivilegeEscalation: false
  # capabilities:
  #   drop: ["ALL"]
  # readOnlyRootFilesystem: true

# Container lifecycle hooks
lifecycle: {}

# Extra volumes and volume mounts
extraVolumes: []
extraVolumeMounts: []

# Pod scheduling
priorityClassName: ""
runtimeClassName: ""
terminationGracePeriodSeconds: 3600
tolerations: []
affinity: {}
nodeSelector: {}

# Liveness and readiness probes
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

### CRD Management

```yaml
upgradeCRDs: true  # Upgrade CRDs on Helm install/upgrade
cleanUpCRDs: false  # Cleanup CRDs (destructive - use cautiously)

upgradeJobResources: {}
  # requests:
  #   cpu: 50m
  #   memory: 128Mi
  # limits:
  #   cpu: 100m
  #   memory: 256Mi

upgradeCRDsJob:
  extraVolumes: []
  extraVolumeMounts: []
  extraEnvVars: []
  automountServiceAccountToken: true
  # shellCmd: /tmp/sh
  # updateCmd: /velero install --crds-only --dry-run -o yaml | /tmp/kubectl apply -f -

kubectl:
  image:
    repository: registry.k8s.io/kubectl
    # tag: v1.34.5  # Overrides cluster Kubernetes version
    # digest: ""
  containerSecurityContext: {}
  resources: {}
  annotations: {}
  labels: {}
  extraVolumes: []
  extraVolumeMounts: []
```

### Metrics Configuration

```yaml
metrics:
  enabled: true
  scrapeInterval: 30s
  scrapeTimeout: 10s

  service:
    annotations: {}
    type: ClusterIP
    labels: {}
    nodePort: null
    externalTrafficPolicy: ""  # Cluster or Local
    internalTrafficPolicy: ""
    ipFamilyPolicy: ""  # For dual-stack support
    ipFamilies: []      # IPv4 and/or IPv6

  podAnnotations:
    prometheus.io/scrape: "true"
    prometheus.io/port: "8085"
    prometheus.io/path: "/metrics"

  serviceMonitor:
    autodetect: true
    enabled: false
    annotations: {}
    additionalLabels: {}
    # metricRelabelings: []
    # relabelings: []
    # namespace: ""
    # scheme: ""
    # tlsConfig: {}

  nodeAgentPodMonitor:
    autodetect: true
    enabled: false
    annotations: {}
    additionalLabels: {}
    # metricRelabelings: []
    # relabelings: []
    # namespace: ""
    # scheme: ""
    # tlsConfig: {}

  prometheusRule:
    autodetect: true
    enabled: false
    additionalLabels: {}
    # namespace: ""
    spec: []
    # - alert: VeleroBackupFailed
    #   annotations:
    #     message: Velero backup {{ $labels.schedule }} has failed
    #   expr: |-
    #     velero_backup_last_status{schedule!=""} != 1
    #   for: 15m
    #   labels:
    #     severity: warning
```

### Backup and Snapshot Configuration

```yaml
backupsEnabled: true   # Create BackupStorageLocation CRD
snapshotsEnabled: true # Create VolumeSnapshotLocation CRD

configuration:
  # BackupStorageLocation(s) - configure multiple by adding more elements
  backupStorageLocation:
  - name: ""             # Defaults to "default"
    provider: ""         # e.g., aws, azure, gcp
    bucket: ""           # Storage bucket name
    caCert: ""           # Base64 encoded CA bundle
    prefix: ""           # Directory prefix
    default: false       # Mark as default location
    validationFrequency: ""
    accessMode: ReadWrite  # ReadWrite or ReadOnly
    credential:
      name: ""           # Secret name
      key: ""            # Secret key
    config: {}
    #  region: ""
    #  s3ForcePathStyle: ""
    #  s3Url: ""
    #  kmsKeyId: ""
    #  resourceGroup: ""
    #  subscriptionId: ""
    #  storageAccount: ""
    #  publicUrl: ""
    #  serviceAccount: ""
    #  insecureSkipTLSVerify: ""
    annotations: {}

  # VolumeSnapshotLocation(s)
  volumeSnapshotLocation:
  - name: ""             # Defaults to "default"
    provider: ""         # Snapshot provider
    credential:
      name: ""
      key: ""
    config: {}
    #  region: ""
    #  apiTimeout: ""
    #  resourceGroup: ""
    #  subscriptionId: ""
    #  incremental: ""
    #  snapshotLocation: ""
    #  project: ""
    annotations: {}

  # Server-level settings (velero server CLI flags)
  uploaderType: ""         # Default: kopia
  backupSyncPeriod: ""     # Default: 1m
  fsBackupTimeout: ""      # Default: 4h
  clientBurst: ""          # Default: 30
  clientPageSize: ""       # Default: 500
  clientQPS: ""            # Default: 20.0
  defaultBackupStorageLocation: ""  # Default: "default"
  defaultItemOperationTimeout: ""   # Default: 4h
  defaultBackupTTL: ""     # Default: 72h
  defaultVolumeSnapshotLocations: ""
  disableControllers: ""   # Comma-separated list
  disableInformerCache: false
  garbageCollectionFrequency: ""  # Default: 1h
  itemBlockWorkerCount: ""        # Default: 1
  logFormat: ""            # text or json
  logLevel: ""             # info, debug, warning, error, fatal, panic
  metricsAddress: ""       # Default: :8085
  pluginDir: ""            # Default: /plugins
  profilerAddress: ""      # Default: localhost:6060
  restoreOnlyMode: ""      # Default: false
  restoreResourcePriorities: ""
  storeValidationFrequency: ""  # Default: 1m
  terminatingResourceTimeout: ""  # Default: 10m
  defaultSnapshotMoveData: ""     # Default: false
  dataMoverPrepareTimeout: ""     # Default: 30m
  features: ""             # Comma-separated feature flags (e.g., EnableCSI)
  defaultVolumesToFsBackup: ""    # Default: false
  defaultRepoMaintainFrequency: ""

  namespace: ""     # Default: velero

  repositoryMaintenanceJob:
    repositoryConfigData:
      name: "velero-repo-maintenance"
      global:
        keepLatestMaintenanceJobs: 3
      # repositories:
      #   "kibishii-default-kopia":
      #     podResources:
      #       cpuRequest: "200m"
      #       cpuLimit: "400m"
      #       memoryRequest: "200Mi"
      #       memoryLimit: "400Mi"
      #     keepLatestMaintenanceJobs: 2
      repositories: {}

  extraArgs: []
    # - "--foo=bar"

  extraEnvVars: []
    # - name: SIMPLE_VAR
    #   value: "simple-value"
    # - name: MY_POD_LABEL
    #   valueFrom:
    #     fieldRef:
    #       fieldPath: metadata.labels['my_label']
```

### RBAC Configuration

```yaml
rbac:
  create: true               # Create Velero role and role binding
  clusterAdministrator: true # Create cluster role binding
  clusterAdministratorName: cluster-admin
```

### Service Account Configuration

```yaml
serviceAccount:
  server:
    create: true
    name: ""
    annotations: {}
    labels: {}
    imagePullSecrets: []
    automountServiceAccountToken: true
```

### Credentials Configuration

```yaml
credentials:
  useSecret: true            # Set false if using kube2iam, kiam, or workload identity
  name: ""                   # Secret name to create
  existingSecret: ""         # Pre-existing secret name
  secretContents: {}         # Cloud provider credentials
    # cloud: |
    #   [default]
    #   aws_access_key_id=<REDACTED>
    #   aws_secret_access_key=<REDACTED>
  extraEnvVars: {}           # Environment variables stored in secret
  extraSecretRef: ""         # Pre-existing secret for environment variables
```

### Node Agent Configuration

```yaml
deployNodeAgent: false

nodeAgent:
  disableHostPath: false
  podVolumePath: /var/lib/kubelet/pods
  pluginVolumePath: /var/lib/kubelet/plugins
  
  priorityClassName: ""
  runtimeClassName: ""
  
  resources: {}
    # requests:
    #   cpu: 500m
    #   memory: 512Mi
    # limits:
    #   cpu: 1000m
    #   memory: 1024Mi
  
  resizePolicy: []
    # - resourceName: cpu
    #   restartPolicy: NotRequired
  
  tolerations: []
  annotations: {}
  labels: {}
  podLabels: {}
  
  useScratchEmptyDir: true
  extraVolumes: []
  extraVolumeMounts: []
  
  extraEnvVars: []
  extraArgs: []
  
  dnsPolicy: ClusterFirst
  hostAliases: []
  
  podSecurityContext:
    runAsUser: 0
    # fsGroup: 1337
  
  containerSecurityContext: {}
  lifecycle: {}
  nodeSelector: {}
  affinity: {}
  dnsConfig: {}
  updateStrategy: {}
```

### Backup Schedules

```yaml
schedules: {}
# Example:
# mybackup:
#   disabled: false
#   labels:
#     myenv: foo
#   annotations:
#     myenv: foo
#   schedule: "0 0 * * *"
#   useOwnerReferencesInBackup: false
#   paused: false
#   skipImmediately: false
#   template:
#     ttl: "240h"
#     storageLocation: default
#     includedNamespaces:
#     - foo
#     excludedNamespaceScopedResources:
#     - persistentVolumeClaims
#     excludedClusterScopedResources:
#     - persistentVolumes
```

### ConfigMaps

```yaml
configMaps: {}
# Example:
# fs-restore-action-config:
#   labels:
#     velero.io/plugin-config: ""
#     velero.io/pod-volume-restore: RestoreItemAction
#   data:
#     image: velero/velero:v1.17.1
#     cpuRequest: 200m
#     memRequest: 128Mi
#     cpuLimit: 200m
#     memLimit: 128Mi
```

### Extra Objects

```yaml
extraObjects: []
# Example: Deploy custom SecretProviderClass
# - apiVersion: secrets-store.csi.x-k8s.io/v1
#   kind: SecretProviderClass
#   metadata:
#     name: velero-secrets-store
#   spec:
#     provider: aws
#     parameters:
#       objects: |
#         - objectName: "velero"
#           objectType: "secretsmanager"
#           jmesPath:
#               - path: "access_key"
#                 objectAlias: "access_key"
#               - path: "secret_key"
#                 objectAlias: "secret_key"
#     secretObjects:
#       - data:
#         - key: access_key
#           objectName: client-id
#         - key: client-secret
#           objectName: client-secret
#         secretName: velero-secrets-store
#         type: Opaque
```

---

## EnvoyGateway Chart Configuration

The EnvoyGateway chart deploys the Envoy Gateway API implementation.

### Global Configuration

```yaml
global:
  imageRegistry: ""          # Global image registry override
  imagePullSecrets: []       # Global image pull secrets
  
  images:
    envoyGateway:
      image: docker.io/envoyproxy/gateway:v1.7.1
      pullPolicy: IfNotPresent
      pullSecrets: []
    ratelimit:
      image: docker.io/envoyproxy/ratelimit:c8765e89
      pullPolicy: IfNotPresent
      pullSecrets: []
```

### Pod Disruption Budget

```yaml
podDisruptionBudget:
  minAvailable: 0
  # maxUnavailable: 1
```

### Deployment Configuration

```yaml
deployment:
  annotations: {}
  
  envoyGateway:
    image:
      repository: ""  # Override with full image if using custom registry
      tag: ""         # Override version
    imagePullPolicy: ""
    imagePullSecrets: []
    
    resources:
      limits:
        memory: 1024Mi
      requests:
        cpu: 100m
        memory: 256Mi
    
    securityContext:
      allowPrivilegeEscalation: false
      capabilities:
        drop:
        - ALL
      privileged: false
      runAsNonRoot: true
      runAsGroup: 65532
      runAsUser: 65532
      seccompProfile:
        type: RuntimeDefault
  
  ports:
    - name: grpc
      port: 18000
      targetPort: 18000
    - name: ratelimit
      port: 18001
      targetPort: 18001
    - name: wasm
      port: 18002
      targetPort: 18002
    - name: metrics
      port: 19001
      targetPort: 19001
  
  priorityClassName: null
  replicas: 1
  
  pod:
    affinity: {}
    annotations:
      prometheus.io/scrape: 'true'
      prometheus.io/port: '19001'
    labels: {}
    topologySpreadConstraints: []
    tolerations: []
    nodeSelector: {}
```

### Service Configuration

```yaml
service:
  trafficDistribution: ""    # PreferClose for topology-aware routing
  annotations: {}
  type: "ClusterIP"
  # loadBalancerIP: 10.236.90.20
```

### Horizontal Pod Autoscaler

```yaml
hpa:
  enabled: false
  minReplicas: 1
  maxReplicas: 1
  metrics: []
  behavior: {}
```

### EnvoyGateway Configuration

```yaml
config:
  envoyGateway:
    gateway:
      controllerName: gateway.envoyproxy.io/gatewayclass-controller
    provider:
      type: Kubernetes
    logging:
      level:
        default: info
    extensionApis: {}
```

### Namespace Management

```yaml
createNamespace: false
kubernetesClusterDomain: cluster.local
```

### Certificate Generation

```yaml
# Certificate generation for EnvoyGateway (OIDC, OAuth2, etc.)
# Do not disable certgen; it may cause issues with OIDC, OAuth2, etc.
certgen:
  job:
    annotations: {}
    args: []
    
    pod:
      annotations: {}
      labels: {}
    
    resources: {}
    affinity: {}
    tolerations: []
    nodeSelector: {}
    ttlSecondsAfterFinished: 30
    
    securityContext:
      allowPrivilegeEscalation: false
      capabilities:
        drop:
        - ALL
      privileged: false
      readOnlyRootFilesystem: true
      runAsNonRoot: true
      runAsGroup: 65532
      runAsUser: 65532
      seccompProfile:
        type: RuntimeDefault
  
  rbac:
    annotations: {}
    labels: {}
```

### Topology Injector

```yaml
topologyInjector:
  enabled: true
  annotations: {}
```

---

## Common Customization Examples

### Backup to AWS S3

```yaml
velero:
  image:
    repository: docker.io/velero/velero
    tag: v1.18.0
  configuration:
    backupStorageLocation:
    - name: default
      provider: aws
      bucket: my-backup-bucket
      prefix: velero
      default: true
      config:
        region: us-east-1
        s3ForcePathStyle: false
    credentials:
      useSecret: true
      name: velero-aws-credentials
      secretContents:
        cloud: |
          [default]
          aws_access_key_id=AKIAIOSFODNN7EXAMPLE
          aws_secret_access_key=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
```

### Enable Prometheus Monitoring

```yaml
velero:
  metrics:
    enabled: true
    serviceMonitor:
      enabled: true
      additionalLabels:
        release: prometheus
```

### Deploy Node Agent with Backup Schedules

```yaml
velero:
  deployNodeAgent: true
  nodeAgent:
    resources:
      requests:
        cpu: 500m
        memory: 512Mi
      limits:
        cpu: 1000m
        memory: 1024Mi
  schedules:
    daily-backup:
      schedule: "0 2 * * *"
      template:
        ttl: "240h"
        storageLocation: default
        includedNamespaces:
        - "*"
```

### Custom EnvoyGateway Deployment

```yaml
envoy:
  deployment:
    replicas: 3
    envoyGateway:
      resources:
        requests:
          cpu: 200m
          memory: 512Mi
        limits:
          cpu: 500m
          memory: 1024Mi
    pod:
      affinity:
        podAntiAffinity:
          preferredDuringSchedulingIgnoredDuringExecution:
          - weight: 100
            podAffinityTerm:
              labelSelector:
                matchExpressions:
                - key: app
                  operator: In
                  values:
                  - envoy-gateway
              topologyKey: kubernetes.io/hostname
  hpa:
    enabled: true
    minReplicas: 2
    maxReplicas: 10
    metrics:
    - type: Resource
      resource:
        name: cpu
        target:
          type: Utilization
          averageUtilization: 70
```
