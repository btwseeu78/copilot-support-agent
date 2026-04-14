# Helm Chart Structural Diff Report

## Summary

**BREAKING CHANGES DETECTED: YES**

The `gateway-helm` dependency contains two overrides targeting key paths that do not exist in upstream v1.7.1. These overrides are silently dropped at render time. The `velero` dependency overrides are structurally compatible with upstream v12.0.0, but a major version bump (11.0.0 → 12.0.0) introduces new configurable keys worth reviewing.

---

## Version Info

| Dependency    | Local Version | Latest Upstream Version | Status         |
|---------------|---------------|-------------------------|----------------|
| `velero`      | 11.0.0        | 12.0.0                  | ⚠️ Out of date  |
| `gateway-helm`| v1.7.1        | v1.7.1                  | ✅ Up to date   |

---

## Breaking Changes

| # | Key Path (local)                        | Change Type              | Details |
|---|-----------------------------------------|--------------------------|---------|
| 1 | `gateway-helm.global.image`             | **Key does not exist**   | Local overrides `global.image.repository`, `global.image.tag`, and `global.image.pullPolicy`, but upstream v1.7.1 has **no `global.image` key**. Upstream uses `global.imageRegistry` (string) for registry override and `global.images.envoyGateway` / `global.images.ratelimit` for per-image settings. These overrides are silently ignored. |
| 2 | `gateway-helm.global.service`           | **Key does not exist**   | Local overrides `global.service.type` and `global.service.port`, but upstream has **no `global.service` key**. The `service` block exists at **top level** (`service.type`, `service.annotations`). The local path is wrong and overrides are silently dropped. Additionally, `service.port` does not exist in upstream at all (top-level `service` has no `port` field). |

---

## Missing Overrides

New keys in blocks the local file already overrides, or notable additions in the upstream version bump.

### gateway-helm (v1.7.1 — same version, structural mismatches only)

The correct paths to override image and service settings:

| Correct Upstream Key Path                          | Type   | Default                                   | Notes |
|----------------------------------------------------|--------|-------------------------------------------|-------|
| `gateway-helm.global.imageRegistry`                | string | `""`                                      | Global registry override (replaces registry portion of all images) |
| `gateway-helm.global.imagePullSecrets`             | list   | `[]`                                      | Global pull secrets for all images |
| `gateway-helm.global.images.envoyGateway.image`    | string | `docker.io/envoyproxy/gateway:v1.7.1`     | Full image reference (registry+repo+tag) for EnvoyGateway |
| `gateway-helm.global.images.envoyGateway.pullPolicy` | string | `IfNotPresent`                          | Pull policy for EnvoyGateway image |
| `gateway-helm.global.images.ratelimit.image`       | string | `docker.io/envoyproxy/ratelimit:c8765e89` | Full image reference for ratelimit sidecar |
| `gateway-helm.service.type`                        | string | `ClusterIP`                               | Correct path for service type (not under `global`) |
| `gateway-helm.deployment.envoyGateway.image.repository` | string | `""`                                | Per-deployment image repository override |
| `gateway-helm.deployment.envoyGateway.image.tag`   | string | `""`                                      | Per-deployment image tag override |

### velero (v11.0.0 → v12.0.0 — notable new keys)

The local file overrides `velero.image.*`, which is structurally unchanged and safe. However, the following keys are new in v12.0.0 and may require attention:

| New Upstream Key Path                                                        | Type   | Default | Notes |
|------------------------------------------------------------------------------|--------|---------|-------|
| `velero.namespace.labels`                                                    | map    | `{}`    | Labels applied to the Velero install namespace (e.g. Pod Security Standards) |
| `velero.resizePolicy`                                                        | list   | `[]`    | Container resize policy for the Velero deployment |
| `velero.upgradeCRDsJob.extraVolumes`                                         | list   | `[]`    | Extra volumes for the upgrade CRDs job |
| `velero.upgradeCRDsJob.extraVolumeMounts`                                    | list   | `[]`    | Extra volume mounts for the upgrade CRDs job |
| `velero.upgradeCRDsJob.extraEnvVars`                                         | list   | `[]`    | Extra env vars for the upgrade CRDs job |
| `velero.upgradeCRDsJob.automountServiceAccountToken`                         | bool   | `true`  | Controls SA token automount in the upgrade job |
| `velero.configuration.repositoryMaintenanceJob.repositoryConfigData.name`   | string | `velero-repo-maintenance` | ConfigMap name for per-repository maintenance settings |
| `velero.configuration.repositoryMaintenanceJob.repositoryConfigData.global.keepLatestMaintenanceJobs` | int | `3` | Global retention for maintenance jobs |
| `velero.configuration.repositoryMaintenanceJob.repositoryConfigData.repositories` | map | `{}` | Per-repository resource/job settings |
| `velero.nodeAgent.disableHostPath`                                           | bool   | `false` | Disables host path volumes for node-agent |
| `velero.nodeAgent.podVolumePath`                                             | string | `/var/lib/kubelet/pods` | Kubelet pods path for node-agent |
| `velero.nodeAgent.pluginVolumePath`                                          | string | `/var/lib/kubelet/plugins` | Kubelet plugins path for node-agent |
| `velero.nodeAgent.resizePolicy`                                              | list   | `[]`    | Container resize policy for node-agent daemonset |
| `velero.nodeAgent.lifecycle`                                                 | map    | `{}`    | Lifecycle hooks for node-agent containers |
| `velero.metrics.service.externalTrafficPolicy`                               | string | `""`    | External traffic policy for metrics service |
| `velero.metrics.service.internalTrafficPolicy`                               | string | `""`    | Internal traffic policy for metrics service |
| `velero.metrics.service.ipFamilyPolicy`                                      | string | `""`    | IP family policy for dual-stack metrics service |
| `velero.metrics.service.ipFamilies`                                          | list   | `[]`    | IP families for metrics service |
| `velero.metrics.nodeAgentPodMonitor`                                         | map    | (block) | New PodMonitor block for node-agent metrics |
| `velero.configuration.itemBlockWorkerCount`                                  | int    | `1`     | Worker count for item block processing |

---

## Info (Non-Breaking Structural Differences)

| Key Path                              | Note |
|---------------------------------------|------|
| `gateway-helm.enabled`                | Helm subchart convention key; not present in upstream `values.yaml` but valid — controls whether the subchart is deployed. No issue. |
| `velero.image.*`                      | All four local overrides (`repository`, `tag`, `pullPolicy`, `imagePullSecrets`) match upstream v12.0.0 structure exactly. Safe. |
| `velero` version bump 11→12           | Major version bump. Review full changelog before upgrading. The locally-overridden keys are structurally compatible, but the new keys above may alter default behavior (especially `repositoryMaintenanceJob` and node-agent path defaults). |
