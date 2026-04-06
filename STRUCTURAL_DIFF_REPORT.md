# Helm Chart Structural Diff Analysis Report

**Analysis Date**: 2026-04-02  
**Chart**: velero-test v12.0.0  
**Dependencies Analyzed**:
- velero v12.0.0 (https://vmware-tanzu.github.io/helm-charts)
- gateway-helm v1.7.1 (oci://docker.io/envoyproxy)

---

## Executive Summary

| Chart | Status | Breaking Changes | Missing Overrides | Orphaned Keys |
|-------|--------|------------------|--------------------|---------------|
| **velero** | ✅ Safe | 0 | 100+ | 1 |
| **gateway-helm** | ⚠️ Caution | 0 | 40+ | 6 |

**Overall Verdict**: **No breaking changes detected**, but both dependencies have significant missing overrides that may need attention depending on your use case.

---

## Velero Dependency (v12.0.0)

### Summary
- **Local Overrides**: 5 keys (image.repository, image.tag, image.pullPolicy, image.imagePullSecrets, image)
- **Upstream Keys**: 150+ configuration options
- **Missing Overrides**: 100+ upstream keys not set locally
- **Type Conflicts**: None detected

### Breaking Changes

**None detected** ✅

All local keys match upstream types exactly. No upstream keys have been removed or restructured.

### Missing Overrides (Significant)

These upstream keys exist but are not configured in your local values.yaml. Review and set them if needed:

#### Critical Backup/Storage Configuration
| Key | Upstream Default | Recommendation |
|-----|------------------|-----------------|
| `configuration.backupStorageLocation` | `[]` (empty) | **MUST CONFIGURE** - Define backup storage provider (AWS, Azure, GCP) |
| `configuration.volumeSnapshotLocation` | `[]` (empty) | Configure snapshot provider if using snapshot-based backups |
| `backupsEnabled` | `true` | Keep enabled unless you only want snapshot capability |
| `snapshotsEnabled` | `true` | Consider disabling if not using snapshot backups |

#### RBAC & Service Account
| Key | Upstream Default | Recommendation |
|-----|------------------|-----------------|
| `rbac.create` | `true` | Recommended: Use upstream RBAC |
| `rbac.clusterAdministrator` | `true` | Review: May want to restrict permissions |
| `serviceAccount.server.create` | `true` | Recommended: Create service account |
| `credentials.useSecret` | `true` | Review: Set to false if using IAM roles/workload identity |

#### Metrics & Monitoring
| Key | Upstream Default | Recommendation |
|-----|------------------|-----------------|
| `metrics.enabled` | `true` | Recommended: Keep enabled for observability |
| `metrics.serviceMonitor.enabled` | `false` | Enable if using Prometheus Operator |
| `metrics.prometheusRule.enabled` | `false` | Enable to deploy alert rules |

#### Node Agent (File System Backups)
| Key | Upstream Default | Recommendation |
|-----|------------------|-----------------|
| `deployNodeAgent` | `false` | Set to `true` if you need file system backups |
| `nodeAgent.resources` | `{}` | Set resource requests/limits if deploying node agent |

#### Server Configuration
| Key | Upstream Default | Recommendation |
|-----|------------------|-----------------|
| `configuration.namespace` | `velero` | Use default unless multi-instance deployment |
| `configuration.logLevel` | `info` | Keep default or adjust to `debug` for troubleshooting |
| `configuration.features` | empty | Enable CSI: `features: EnableCSI` if needed |

### Non-Breaking Info

#### Orphaned/Extra Keys in Local

| Key | Local Value | Status |
|-----|-------------|--------|
| `image` (top-level container) | Present | Safe - This is the image block itself; no conflict |

**Explanation**: The "image" entry is not orphaned—it's the valid parent key for image configuration in Velero. This is expected.

#### New Upstream Options (Not in Local)

- **Kubernetes 1.16+**: livenessProbe, readinessProbe, resizePolicy, topologySpreadConstraints
- **CSI Support**: Configuration for CSI snapshots via `configuration.features: EnableCSI`
- **Repository Maintenance**: New `configuration.repositoryMaintenanceJob` for Kopia backend maintenance
- **Pod Scheduling**: `terminationGracePeriodSeconds: 3600`, multiple affinity/anti-affinity options
- **Extra Objects**: `extraObjects: []` for deploying custom K8s resources

### Recommendations

1. **URGENT**: Configure `configuration.backupStorageLocation` with your cloud provider credentials
2. **RECOMMENDED**: Enable `metrics.enabled: true` and `metrics.serviceMonitor.enabled: true` for monitoring
3. **OPTIONAL**: Deploy `nodeAgent` if you need file system backup capability
4. **OPTIONAL**: Set `configuration.logLevel: debug` if troubleshooting backup issues

---

## EnvoyGateway/gateway-helm Dependency (v1.7.1)

### Summary
- **Local Overrides**: 8 keys (enabled, image.*, service.type, service.port)
- **Upstream Keys**: 45+ configuration options
- **Missing Overrides**: 40+ upstream keys not set locally
- **Type Conflicts**: None detected

### Breaking Changes

**None detected** ✅

All local keys match upstream types. However, there are structural considerations:

#### Type Alignment Check

| Local Key | Local Type | Upstream Type | Status |
|-----------|-----------|---------------|--------|
| `envoy.image.repository` | string | ❌ Not in upstream | Extra key |
| `envoy.image.tag` | string | ❌ Not in upstream | Extra key |
| `envoy.image.pullPolicy` | string | ❌ Not in upstream | Extra key |
| `envoy.service.type` | string | ✅ string | Compatible |
| `envoy.service.port` | integer | ❌ Not in upstream | Extra key |

**Analysis**: Local values use flattened structure (`envoy.image.*`) while upstream uses nested structure (`global.images.envoyGateway.image`, `deployment.envoyGateway.image.*`). This may require template changes to apply correctly.

### Missing Overrides (Significant)

#### Global Image Configuration (Recommended Override Point)

| Key | Upstream Default | Recommendation |
|-----|------------------|-----------------|
| `global.imageRegistry` | empty string | Leave empty unless using private registry |
| `global.imagePullSecrets` | `[]` | Configure if pulling from private registry |
| `global.images.envoyGateway.image` | `docker.io/envoyproxy/gateway:v1.7.1` | Override for private registry |
| `global.images.ratelimit.image` | `docker.io/envoyproxy/ratelimit:c8765e89` | Override if using ratelimiting |

#### Deployment Configuration

| Key | Upstream Default | Recommendation |
|-----|------------------|-----------------|
| `deployment.envoyGateway.image.repository` | empty | Set for custom registry |
| `deployment.envoyGateway.image.tag` | empty | Set to override version |
| `deployment.envoyGateway.resources.requests.cpu` | `100m` | Adjust based on traffic |
| `deployment.envoyGateway.resources.limits.memory` | `1024Mi` | Increase for high throughput |
| `deployment.replicas` | `1` | Set to 3+ for HA |
| `deployment.pod.affinity` | `{}` | Recommend pod anti-affinity for HA |

#### Pod Disruption & Scheduling

| Key | Upstream Default | Recommendation |
|-----|------------------|-----------------|
| `podDisruptionBudget.minAvailable` | `0` | Set to `1` for HA deployments |
| `deployment.pod.topologySpreadConstraints` | `[]` | Set for even distribution across nodes |
| `hpa.enabled` | `false` | Enable for auto-scaling |
| `hpa.minReplicas` | `1` | Adjust based on needs |
| `hpa.maxReplicas` | `1` | Increase for HA |

#### Monitoring Configuration

| Key | Upstream Default | Value |
|-----|------------------|-------|
| `deployment.pod.annotations.prometheus.io/scrape` | `'true'` | Metrics enabled by default |
| `deployment.pod.annotations.prometheus.io/port` | `'19001'` | Metrics port |

#### Certificate Generation

| Key | Upstream Default | Note |
|-----|------------------|------|
| `certgen.job.ttlSecondsAfterFinished` | `30` | Job cleaned up after 30 seconds |
| `certgen.rbac.annotations` | `{}` | Add annotations as needed |

### Non-Breaking Info

#### Orphaned/Extra Keys in Local

| Key | Local Value | Analysis |
|-----|-------------|----------|
| `envoy.enabled` | `true` | **New Override** - Not in upstream; May be used in parent chart logic |
| `envoy.image.*` | multiple | **Structural Mismatch** - Upstream uses `global.images.envoyGateway.*` and `deployment.envoyGateway.image.*` |
| `envoy.service.port` | `80` | **New Override** - Not in upstream; Parent chart may use this |

**Analysis**: These keys are likely defined by the parent chart (velero-test) to customize the EnvoyGateway subchart. The parent chart template must reference these values explicitly.

#### Key Structural Differences

The local values use a simplified namespace:
```yaml
envoy:
  image:
    repository: docker.io/envoyproxy/envoy
    tag: v1.7.1
```

But upstream expects:
```yaml
global:
  images:
    envoyGateway:
      image: docker.io/envoyproxy/gateway:v1.7.1
      pullPolicy: IfNotPresent
OR
deployment:
  envoyGateway:
    image:
      repository: docker.io/envoyproxy/gateway
      tag: v1.7.1
```

**Note**: The parent chart (velero-test) must have template logic that transforms `envoy.*` into the upstream structure.

### Recommendations

1. **VERIFY TEMPLATE**: Ensure the parent chart's templates correctly map `envoy.*` keys to the upstream structure (`global.*` or `deployment.*`)
2. **HA DEPLOYMENT**: If this is production, configure:
   ```yaml
   deployment:
     replicas: 3
     pod:
       affinity:
         podAntiAffinity:
           preferredDuringSchedulingIgnoredDuringExecution: [...]
   hpa:
     enabled: true
     minReplicas: 2
     maxReplicas: 10
   podDisruptionBudget:
     minAvailable: 1
   ```
3. **RESOURCES**: Adjust CPU/memory limits based on your traffic patterns
4. **MONITORING**: Metrics are enabled by default; consider integrating with Prometheus

---

## Summary Table: All Missing Overrides

### Velero - High Priority
- `configuration.backupStorageLocation` - **CRITICAL**
- `credentials.*` - **CRITICAL**
- `rbac.*` - **IMPORTANT**
- `serviceAccount.*` - **IMPORTANT**
- `metrics.serviceMonitor.enabled` - Recommended for observability

### Velero - Optional
- `deployNodeAgent` - Only if file system backups needed
- `configuration.features` - For advanced features like CSI
- `configuration.logLevel` - For debugging
- `nodeAgent.*` - Only if deployNodeAgent is true

### EnvoyGateway - High Priority
- `deployment.replicas` - Set to 3+ for HA
- `deployment.pod.affinity` - For pod anti-affinity
- `hpa.enabled` - For auto-scaling
- `podDisruptionBudget.minAvailable` - For HA

### EnvoyGateway - Optional
- `global.imageRegistry` - For private registries
- `certgen.*` - Usually defaults are fine
- `config.envoyGateway` - Advanced configuration

---

## Type Safety Check: Passing ✅

| Aspect | Status | Details |
|--------|--------|---------|
| Image Keys | ✅ Safe | All image config keys use correct string/list types |
| Service Config | ✅ Safe | Service type is string (ClusterIP match) |
| Resources | ✅ Safe | No resource definitions, so no type conflicts |
| Lists | ✅ Safe | All list types (imagePullSecrets, tolerations) are correct |
| Booleans | ✅ Safe | All boolean configs match (enabled flags) |
| Maps | ✅ Safe | All dict/object configs match |

---

## Next Steps

1. **For Velero**: Configure backup storage location credentials
2. **For EnvoyGateway**: Verify parent chart templates correctly map custom keys to upstream structure
3. **For Both**: Review high-priority missing overrides based on your deployment requirements
4. **Test**: Deploy with updated values.yaml and verify both components start correctly

