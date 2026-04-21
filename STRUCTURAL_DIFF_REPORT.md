# Helm Chart Structural Diff Report

## Summary

**BREAKING CHANGES DETECTED: YES** — The local `values.yaml` overrides keys in `gateway-helm` that do not exist in the upstream chart (wrong key names / wrong nesting), meaning those overrides will have no effect and the chart may behave unexpectedly.

---

## Version Info

| Dependency    | Local Version | Latest Upstream Version | Version Delta |
|---------------|--------------|------------------------|---------------|
| velero        | 11.0.0       | 12.0.0                 | ⚠️ Behind by 1 major version |
| gateway-helm  | v1.7.1       | v1.7.1                 | ✅ Up to date |

---

## Breaking Changes

| Key Path (local) | Change Type | Details |
|---|---|---|
| `gateway-helm.global.image` | **BREAKING — Key does not exist** | Upstream uses `global.images` (plural) with sub-keys `envoyGateway` and `ratelimit`. Local overrides `global.image` (singular) which is not a valid upstream key. This override is silently ignored. |
| `gateway-helm.global.image.repository` | **BREAKING — Wrong path** | Correct upstream paths: `global.images.envoyGateway.image` (full image string) or `deployment.envoyGateway.image.repository`. |
| `gateway-helm.global.image.tag` | **BREAKING — Wrong path** | No standalone `tag` field under `global.images.*`. Use `global.images.envoyGateway.image` (full image string including tag) or `deployment.envoyGateway.image.tag`. |
| `gateway-helm.global.image.pullPolicy` | **BREAKING — Wrong path** | Correct upstream path: `global.images.envoyGateway.pullPolicy`. |
| `gateway-helm.global.service` | **BREAKING — Key does not exist** | Upstream has no `global.service`. The service configuration lives at `service.type`, `service.annotations`, `service.trafficDistribution`. This override is silently ignored. |
| `gateway-helm.enabled` | **BREAKING — Key does not exist** | Upstream gateway-helm chart has no top-level `enabled` key. This override has no effect. |

---

## Missing Overrides (new upstream keys in overridden blocks)

### velero (11.0.0 → 12.0.0)

The local file overrides `velero.image.*`. The upstream image block is structurally compatible, but 12.0.0 adds new top-level keys you may want to review:

| New Upstream Key | Default | Notes |
|---|---|---|
| `velero.configuration.repositoryMaintenanceJob.repositoryConfigData.global.keepLatestMaintenanceJobs` | `3` | New in 12.x — controls how many maintenance jobs to keep. |
| `velero.configuration.repositoryMaintenanceJob.repositoryConfigData.repositories` | `{}` | Per-repository resource/maintenance config map. |
| `velero.upgradeCRDsJob.extraVolumes` | `[]` | New block for the upgrade CRDs job. |
| `velero.upgradeCRDsJob.extraVolumeMounts` | `[]` | New block for the upgrade CRDs job. |
| `velero.upgradeCRDsJob.extraEnvVars` | `[]` | New block for the upgrade CRDs job. |
| `velero.upgradeCRDsJob.automountServiceAccountToken` | `true` | New in upgrade CRDs job. |
| `velero.namespace.labels` | `{}` | New top-level namespace label support. |
| `velero.resizePolicy` | `[]` | Container resize policy for Velero deployment. |
| `velero.nodeAgent.disableHostPath` | `false` | New node-agent flag. |
| `velero.nodeAgent.pluginVolumePath` | `/var/lib/kubelet/plugins` | New node-agent volume path. |
| `velero.nodeAgent.resizePolicy` | `[]` | Container resize policy for node-agent. |
| `velero.metrics.service.externalTrafficPolicy` | `""` | New metrics service traffic policy fields. |
| `velero.metrics.service.internalTrafficPolicy` | `""` | New metrics service traffic policy fields. |
| `velero.metrics.service.ipFamilyPolicy` | `""` | Dual-stack IP family policy. |
| `velero.metrics.service.ipFamilies` | `[]` | Dual-stack IP families list. |
| `velero.metrics.nodeAgentPodMonitor` | (block) | New PodMonitor support for node-agent metrics. |

### gateway-helm (v1.7.1 — same version, but local overrides are broken)

| Upstream Key | Default | Notes |
|---|---|---|
| `global.imageRegistry` | `""` | Global registry override (not used in local). |
| `global.images.envoyGateway.image` | `docker.io/envoyproxy/gateway:v1.7.1` | **Use instead of** `global.image.repository`+`tag`. |
| `global.images.envoyGateway.pullPolicy` | `IfNotPresent` | **Use instead of** `global.image.pullPolicy`. |
| `global.images.ratelimit.image` | `docker.io/envoyproxy/ratelimit:c8765e89` | Ratelimit image (not configured locally). |
| `service.type` | `ClusterIP` | **Use instead of** `global.service.type`. |
| `deployment.envoyGateway.image.repository` | `""` | Alternative per-component image override. |

---

## Info (Non-Breaking Structural Differences)

| Key Path | Notes |
|---|---|
| `velero.image.*` | All four local overrides (`repository`, `tag`, `pullPolicy`, `imagePullSecrets`) are valid and structurally match upstream 12.0.0. Safe. |
| `gateway-helm` version | Local and upstream are both v1.7.1 — no version drift, but local override keys are incorrect (see Breaking Changes above). |

---

## Recommended Fixes

### Fix `gateway-helm` image overrides

Replace:
```yaml
gateway-helm:
  global:
    image:
      repository: docker.io/envoyproxy/envoy
      tag: v1.7.1
      pullPolicy: IfNotPresent
```

With:
```yaml
gateway-helm:
  global:
    images:
      envoyGateway:
        image: docker.io/envoyproxy/envoy:v1.7.1
        pullPolicy: IfNotPresent
```

### Fix `gateway-helm` service override

Replace:
```yaml
gateway-helm:
  global:
    service:
      type: ClusterIP
      port: 80
```

With:
```yaml
gateway-helm:
  service:
    type: ClusterIP
```

> Note: `port` is not a configurable field in upstream `service`; ports are defined under `deployment.ports`.

### Remove `gateway-helm.enabled`

This key does not exist in the upstream chart and has no effect. Remove it.
