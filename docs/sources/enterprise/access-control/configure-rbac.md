---
aliases:
  - /docs/grafana/latest/enterprise/access-control/configure-rbac/
description: Learn how to configure RBAC.
menuTitle: Configure RBAC
title: Configure RBAC in Grafana
weight: 30
---

# Configure RBAC in Grafana

The table below describes all RBAC configuration options. Like any other Grafana configuration, you can apply these options as [environment variables]({{< relref "../../administration/configuration.md#configure-with-environment-variables" >}}).

| Setting            | Required | Description                                                                  | Default |
| ------------------ | -------- | ---------------------------------------------------------------------------- | ------- |
| `permission_cache` | No       | Enable to use in memory cache for loading and evaluating users' permissions. | `true`  |

## Example RBAC configuration

```bash
[rbac]

permission_cache = true
```
