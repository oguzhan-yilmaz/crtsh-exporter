# Prometheus Exporter for [`crt.sh`](https://crt.sh)

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
ghcr.io/dazwilkin/crtsh-exporter:1234567890123456789012345678901234567890
```

> **NOTE** `cosign.pub` may be downloaded [here](./cosign.pub)

To install `cosign` e.g.:

```bash
go install github.com/sigstore/cosign/cmd/cosign@latest
```