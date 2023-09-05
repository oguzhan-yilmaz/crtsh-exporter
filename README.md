# Prometheus Exporter for [`crt.sh`](https://crt.sh)

[![build](https://github.com/DazWilkin/crtsh-exporter/actions/workflows/build.yml/badge.svg)](https://github.com/DazWilkin/crtsh-exporter/actions/workflows/build.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/DazWilkin/crtsh-exporter.svg)](https://pkg.go.dev/github.com/DazWilkin/crtsh-exporter)
[![Go Report Card](https://goreportcard.com/badge/github.com/DazWilkin/crtsh-exporter)](https://goreportcard.com/report/github.com/DazWilkin/crtsh-exporter)

+ `ghcr.io/dazwilkin/crtsh-exporter:ec4dbe7c7abdb58f57cb679f7f7db678f123c4a2`

## Example

```bash
HOST="..."

curl \
--silent \
--get \
--data-urlencode "q=${HOST}" \
--data-urlencode "output=json" \
https://crt.sh
```
Returns
```JSON
[
    {
        "issuer_ca_id": 123456,
        "issuer_name": "C=US, O=Let's Encrypt, CN=R3",
        "common_name": "{HOST}",
        "name_value": "{HOST}",
        "id": 10123456789,
        "entry_timestamp": "2023-01-01T23:59:59.000",
        "not_before": "2023-01-01T23:59:59",
        "not_after": "2023-01-01T23:59:59",
        "serial_number": "123456789abcdef0123456789abcdef0"
    }
]
```

## Metrics

|Name|Type|Description|
|----|----|-----------|
|`crtsh_exporter_build_info`|Counter|A metric with a constant '1' value|
|`crtsh_exporter_certificate_records`|Gauge|Number of Certificate records, labeled by most recent record's metadata|
|`crtsh_exporter_start_time`|Gauge|Exporter start time in UNIX epoch|

## [Sigstore](https://www.sigstore.dev/)

`crtsh-exporter` container images are signed by [Sigstore](https://www.sigstore.dev/) and may be verified:

```bash
cosign verify \
--key=./cosign.pub \
ghcr.io/dazwilkin/crtsh-exporter:ec4dbe7c7abdb58f57cb679f7f7db678f123c4a2
```

> **NOTE** `cosign.pub` may be downloaded [here](./cosign.pub)

To install `cosign` e.g.:

```bash
go install github.com/sigstore/cosign/cmd/cosign@latest
```

## Similar Exporters

+ [Prometheus Exporter for Azure](https://github.com/DazWilkin/azure-exporter)
+ [Prometheus Exporter for Fly.io](https://github.com/DazWilkin/fly-exporter)
+ [Prometheus Exporter for GCP](https://github.com/DazWilkin/gcp-exporter)
+ [Prometheus Exporter for GoatCounter](https://github.com/DazWilkin/goatcounter-exporter)
+ [Prometheus Exporter for Koyeb](https://github.com/DazWilkin/koyeb-exporter)
+ [Prometheus Exporter for Linode](https://github.com/DazWilkin/linode-exporter)
+ [Prometheus Exporter for Porkbun](https://github.com/DazWilkin/porkbun-exporter)
+ [Prometheus Exporter for Vultr](https://github.com/DazWilkin/vultr-exporter)

<hr/>
<br/>
<a href="https://www.buymeacoffee.com/dazwilkin" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/default-orange.png" alt="Buy Me A Coffee" height="41" width="174"></a>