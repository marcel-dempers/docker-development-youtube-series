# Traefik

[Traefik](https://traefik.io/) is a modern HTTP reverse proxy and load balancer made to deploy microservices with ease.

## Introduction

### Philosophy

The Traefik Helm chart is focused on Traefik deployment configuration.

To keep this Helm chart as generic as possible, we avoid integrating third-party solutions or targeting specific use cases.

If you want to customize the chart for your needs, you can:

1. Override the default Traefik configuration values (see [yaml file or CLI](https://helm.sh/docs/chart_template_guide/values_files/)).
2. Append your own configurations (for example, by running `kubectl apply -f myconf.yaml`).

[Examples](https://github.com/traefik/traefik-helm-chart/blob/master/EXAMPLES.md) of common usage are provided.

If you need to include additional Kubernetes objects or extend functionality, use [`extraObjects`](./traefik/tests/values/extra.yaml) or add this chart as a [subchart](https://helm.sh/docs/chart_template_guide/subcharts_and_globals/).

### Major Changes

Starting with v28.x, this chart bootstraps Traefik Proxy version 3 as a Kubernetes ingress controller, using the [`IngressRoute`](https://doc.traefik.io/traefik/v3.0/routing/providers/kubernetes-crd/) Custom Resource.

To upgrade from chart versions prior to v28.x (which use Traefik Proxy version 2), see:

- [Migration guide from v2 to v3](https://doc.traefik.io/traefik/v3.0/migration/v2-to-v3/)
- Upgrade notes in the [`README` on the v27 branch](https://github.com/traefik/traefik-helm-chart/tree/v27)

Starting with v34.x, to work around [Helm caveats](https://helm.sh/docs/chart_best_practices/custom_resource_definitions/#some-caveats-and-explanations), you can use an additional chart dedicated to CRDs: **traefik-crds**.

To deploy with this setup, see the [instructions below](#with-additional-crds-chart).

### Support for Traefik Proxy v2

If you need to use this chart with Traefik Proxy v2, use chart version v27.x.

This chart's support policy aligns with the [upstream support policy](https://doc.traefik.io/traefik/deprecation/releases/) of Traefik Proxy.

For compatibility details, installation instructions, or previous upgrade notes, check the [`README` on the v27 branch](https://github.com/traefik/traefik-helm-chart/tree/v27).

## Installing

### Prerequisites

1. Kubernetes (server) version **v1.22.0 or higher**: `kubectl version`
1. Helm **v3.9.0 or higher** [installed](https://helm.sh/docs/using_helm/#installing-helm): `helm version`
1. Traefik's chart repository: `helm repo add traefik https://traefik.github.io/charts`

### Deploying

#### Standard Installation

To install the chart with default values:

```bash
helm install traefik traefik/traefik
```

or, to install from the OCI registry:

```bash
helm install traefik oci://ghcr.io/traefik/helm/traefik
```

To customize the installation, provide a custom `values` file:

```bash
helm install -f myvalues.yaml traefik traefik/traefik
```

To see example values files, refer to the provided [EXAMPLES](./EXAMPLES.md).

For complete documentation on all available parameters, check the [default values file](./traefik/values.yaml).

#### With Additional CRDs Chart

To manage CRDs separately, use the optional CRDs chart. When using it, the CRDs from the regular Traefik chart are not required.
For more details, see [here](./CONTRIBUTING.md#about-crds).

To install with the CRDs chart:

```bash
helm install traefik-crds traefik/traefik-crds
helm install traefik traefik/traefik --skip-crds
helm list # should display two charts installed
```

## Verification

Starting with v37.0.0, chart releases are signed using *provenance files*.

To verify the chart, follow these steps:

### 1. Download the Public Signing Key

To download the official Traefik Helm chart signing key, run:

```shell
gpg --receive-keys --keyserver hkps://keys.openpgp.org 'B0FBA7678F685E9B7024B79FFD92BB57C5A71A99'
```

Example output:

```shell
gpg: key FD92BB57C5A71A99: public key "TraefikLabs Chart Signing Key <noreply@traefik.io>" imported
gpg: Total number processed: 1
gpg:               imported: 1
```

### 2. Export the Signing Key

By default, GnuPG v2 stores keyrings in a format that is not compatible with Helm chart provenance verification. Before you can verify a Helm chart, you must convert your keyrings to the legacy format:

```shell
gpg --export --output $HOME/.gnupg/pubring.gpg 'B0FBA7678F685E9B7024B79FFD92BB57C5A71A99'
```

### 3. Verify the Chart

To verify the chart, use the appropriate command for your registry:

- OCI Registry

  ```shell
  helm fetch --verify --keyring $HOME/.gnupg/traefik.pubring.gpg oci://ghcr.io/traefik/helm/traefik:<VERSION>
  ```

- Helm Registry (GitHub Pages)

  ```shell
  helm fetch --verify --keyring $HOME/.gnupg/traefik.pubring.gpg traefik/traefik --version <VERSION>
  ```

## Upgrading

To see what has changed in each release, check the [Changelog](./traefik/Changelog.md).

A new major version indicates that there is an incompatible breaking change.

> [!WARNING]
> To avoid issues, **always read the release notes for this chart before upgrading**.

### Upgrade the Standalone Traefik Chart

If you use Helm's native CRD management, you **MUST** upgrade CRDs before running `helm upgrade`, since Helm does **not** update CRDs automatically. See [HIP-0011](https://github.com/helm/community/blob/main/hips/hip-0011.md) for details.

To upgrade the Traefik chart and its CRDs:

```bash
# Update the chart repository
helm repo update
# Check current chart & Traefik version
helm search repo traefik/traefik
# Update CRDs
helm show crds traefik/traefik | kubectl apply --server-side --force-conflicts -f -
# Upgrade Traefik release
helm upgrade traefik traefik/traefik
```

### Upgrade from the Standard Traefik Chart to Traefik + Opt-In CRDs Chart

> [!WARNING]
> To avoid conflicts, **you must change the ownership of CRDs before installing the CRDs chart**.

To migrate to the setup with the additional CRDs chart:

```bash
# Update the chart repository
helm repo update
# Update CRD ownership
kubectl get customresourcedefinitions.apiextensions.k8s.io -o name | grep traefik.io | \
  xargs kubectl patch --type='json' -p='[{"op": "add", "path": "/metadata/labels", "value": {"app.kubernetes.io/managed-by":"Helm"}},{"op": "add", "path": "/metadata/annotations/meta.helm.sh~1release-name", "value":"traefik-crds"},{"op": "add", "path": "/metadata/annotations/meta.helm.sh~1release-namespace", "value":"default"}]'
# If you use gateway API, also change Gateway API ownership
kubectl get customresourcedefinitions.apiextensions.k8s.io -o name | grep gateway.networking.k8s.io | \
  xargs kubectl patch --type='json' -p='[{"op": "add", "path": "/metadata/labels", "value": {"app.kubernetes.io/managed-by":"Helm"}},{"op": "add", "path": "/metadata/annotations/meta.helm.sh~1release-name", "value":"traefik-crds"},{"op": "add", "path": "/metadata/annotations/meta.helm.sh~1release-namespace", "value":"default"}]'
# Deploy the optional CRDs chart
helm install traefik-crds traefik/traefik-crds
# Upgrade Traefik release
helm upgrade traefik traefik/traefik
```

### Upgrade When Using Both Traefik and Opt-In CRDs Chart

To upgrade both Traefik and CRDs charts:

```bash
# Update the chart repository
helm repo update
# Check the current chart & Traefik version
helm search repo traefik/traefik
# Upgrade CRDs (Traefik Proxy v3 CRDs)
helm upgrade traefik-crds traefik/traefik
# Upgrade Traefik release
helm upgrade traefik traefik/traefik
```

## Contributing

To contribute to this chart, please read the [Contributing Guide](./CONTRIBUTING.md).

Thank you to everyone who has already contributed!

<a href="https://github.com/traefik/traefik-helm-chart/graphs/contributors">
  <img src="https://contributors-img.web.app/image?repo=traefik/traefik-helm-chart" alt="Contributors"/>
</a>
