# Traefik Helm Chart Guidelines

This document outlines the guidelines for developing, managing and extending the Traefik helm chart.

This Helm Chart is documented using field description from comments with [helm-docs](https://github.com/norwoodj/helm-docs).

It comes with a JSON schema generated from values with [helm schema](https://github.com/losisin/helm-values-schema-json) plugin.

## Feature Example

```yaml
logs:
  general:
    # -- Set [logs format](https://doc.traefik.io/traefik/observability/logs/#format)
    format:  # @schema enum:["common", "json", null]; type:[string, null]; default: "common"
```

Documention is on the first comment, starting with `# --`
Specific instructions for schema, when needed, are done with the inline comment starting with `# @schema`.

## Whitespace

Extra whitespace is to be avoided in templating. Conditionals should chomp whitespace:

```yaml
{{- if .Values }}
{{- end }}
```

There should be an empty commented line between each primary key in the values.yaml file to separate features from each other.

## Values YAML Design

The values.yaml file is designed to be user-friendly. It does not have to resemble the templated configuration if it is not conducive. Similarly, value names do not have to correspond to fields in the template if it is not conducive.
