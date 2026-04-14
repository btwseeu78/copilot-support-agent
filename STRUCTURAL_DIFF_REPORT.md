# Helm Chart Structural Diff Report

## Summary
**BREAKING CHANGES DETECTED: YES**

Incompatibility found in `gateway-helm` values structure between local overrides and upstream chart.

---

## Version Information

| Chart | Local Version | Upstream Version | Repository |
|-------|---------------|------------------|------------|
| velero | 11.0.0 | 12.0.0 | https://vmware-tanzu.github.io/helm-charts |
| gateway-helm | v1.7.1 | 1.7.1 | oci://docker.io/envoyproxy |

---

## Breaking Changes

### gateway-helm: Structural Incompatibility

| Key Path | Change Type | Details |
|----------|-------------|---------|
| `gateway-helm.image` | **MISSING IN UPSTREAM** | Local defines flat `image` block (repository, tag, pullPolicy). Upstream uses `global.images.envoyGateway.image` and/or `deployment.envoyGateway.image` instead. |
| `gateway-helm.service.port` | **NOT IN UPSTREAM** | Local sets `service.port: 80`, but upstream `service` block only contains `type`, `annotations`, `trafficDistribution`. Port config should use `deployment.ports` instead. |

#### Impact
- Local override `gateway-helm.image.*` will be ignored by Helm chart
- Local `gateway-helm.service.port: 80` has no effect; requires restructuring to use proper upstream paths
- These must be updated to match upstream structure for deployment to work correctly

---

## Missing Overrides (New Upstream Keys)

### velero (11.0.0 → 12.0.0)
No new keys in overridden blocks. The `image.*` keys used locally are stable across versions.

### gateway-helm (v1.7.1)
While local overrides exist, they target non-existent keys. User should add:
- `gateway-helm.global.images.envoyGateway.image` (or use `deployment.envoyGateway.image.repository` + `.tag`)
- If port customization needed: use `gateway-helm.deployment.ports` (list of port objects)

---

## Info (Non-Breaking Differences)

### velero
- Local pins `image.tag: v1.18.0` matches upstream `v1.18.0` ✓
- All local overrides target stable keys present in upstream 12.0.0
- Version bump from 11.0.0 → 12.0.0 introduces many new configuration options but no breaking changes to existing overridden keys

### gateway-helm
- Local version matches upstream version (1.7.1)
- Gateway chart structure: upstream uses hierarchical `global.images`, `deployment.envoyGateway`, and `deployment.pod` for configuration
- Local overrides use incompatible flat structure

---

## Recommendations

1. **Velero**: No action required. Local overrides are compatible with upstream 12.0.0.

2. **Gateway-helm**: Update local `values.yaml` to use correct upstream structure:
   ```yaml
   gateway-helm:
     enabled: true
     global:
       images:
         envoyGateway:
           image: docker.io/envoyproxy/gateway:v1.7.1
           pullPolicy: IfNotPresent
     deployment:
       envoyGateway:
         imagePullSecrets: []
     service:
       type: ClusterIP
   ```
   Note: `service.port` does not exist in upstream; port configuration is in `deployment.ports` list.
