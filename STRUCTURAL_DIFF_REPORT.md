# Structural Diff Report

## Summary

**BREAKING CHANGES DETECTED: YES** — The `gateway-helm` dependency contains overrides for keys that do not exist in upstream, causing them to have no effect.

---

## Version Info

| Dependency   | Local Version | Latest Upstream Version | Status    |
|--------------|---------------|------------------------|-----------|
| velero       | 11.0.0        | 12.0.0                 | Outdated  |
| gateway-helm | v1.7.1        | v1.7.1                 | Current   |

---

## Breaking Changes

| Key Path                       | Change Type               | Details |
|--------------------------------|---------------------------|---------|
| `gateway-helm.image`           | **Key does not exist**    | Top-level `image` block is not a valid upstream key. Image configuration must use `deployment.envoyGateway.image.{repository,tag}`, `deployment.envoyGateway.imagePullPolicy`, or `global.images.envoyGateway.*`. The local override (`image.repository`, `image.tag`, `image.pullPolicy`) has **no effect**. |
| `gateway-helm.service.port`    | **Key does not exist**    | Upstream `service` block has only `trafficDistribution`, `annotations`, and `type`. There is no `port` field. The `port: 80` override is silently ignored. |

---

## Missing Overrides

New upstream keys within blocks that the local file already partially overrides:

### gateway-helm — `service` block (local overrides `service.type`)

| Key Path                           | Upstream Default | Notes |
|------------------------------------|-----------------|-------|
| `gateway-helm.service.trafficDistribution` | `""` | New field; set to `"PreferClose"` to route Envoy fleet traffic to topologically closest pods. |
| `gateway-helm.service.annotations`         | `{}`  | Available for service-level annotations; local does not set this. |

---

## Info (Non-Breaking Differences)

### velero — local version behind upstream (11.0.0 vs 12.0.0)

The local `velero` overrides (`image.repository`, `image.tag`, `image.pullPolicy`, `image.imagePullSecrets`) all map to valid keys in upstream 12.0.0. No breaking changes. However, upgrading from 11.0.0 to 12.0.0 may introduce new upstream keys you might want to review:

| Key Path | Upstream Default | Notes |
|---|---|---|
| `velero.namespace.labels` | `{}` | New top-level namespace label configuration. |
| `velero.upgradeCRDsJob` | (block) | New block with `extraVolumes`, `extraVolumeMounts`, `extraEnvVars`, `automountServiceAccountToken`. |
| `velero.resizePolicy` | `[]` | Container resize policy for the Velero deployment. |
| `velero.hostAliases` | `[]` | Host aliases for Velero pods. |
| `velero.configuration.repositoryMaintenanceJob.repositoryConfigData` | (block) | Per-repository and global maintenance job configuration. |
| `velero.configuration.itemBlockWorkerCount` | (not set) | New server-level flag. |
| `velero.configuration.dataMoverPrepareTimeout` | (not set) | Timeout for CSI snapshot volume provisioning. |
| `velero.nodeAgent.disableHostPath` | `false` | Option to disable host path volumes for node-agent. |
| `velero.nodeAgent.pluginVolumePath` | `/var/lib/kubelet/plugins` | New node-agent plugin volume path. |
| `velero.nodeAgent.resizePolicy` | `[]` | Container resize policy for node-agent. |
| `velero.metrics.service.externalTrafficPolicy` | `""` | New traffic policy fields on metrics service. |
| `velero.metrics.service.internalTrafficPolicy` | `""` | New traffic policy fields on metrics service. |
| `velero.metrics.service.ipFamilyPolicy` | `""` | IP family policy for metrics service. |
| `velero.metrics.service.ipFamilies` | `[]` | IP families for metrics service. |
| `velero.metrics.nodeAgentPodMonitor` | (block) | New PodMonitor for node-agent metrics. |

### gateway-helm — `enabled` key

| Key Path | Notes |
|---|---|
| `gateway-helm.enabled` | Not a valid upstream key. Helm ignores unknown keys; this has no effect. Remove or replace with the correct pattern if conditional deployment is needed. |
