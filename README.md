# AutoDNS Module for Caddy

This package provides a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It allows you to manage DNS records with AutoDNS.

## Caddy Module Name

```
dns.providers.AutoDNS
```

## Installation

To install this module, add it to your Caddy configuration.

## Configuration Examples

### JSON Configuration

To use this module for the ACME DNS challenge, configure the ACME issuer in your Caddy JSON like so:

```json
{
  "apps": {
    "tls": {
      "automation": {
        "policies": [
          {
            "issuers": [
              {
                "module": "acme",
                "challenges": {
                  "dns": {
                    "provider": {
                      "name": "AutoDNS",
                      "Username": "AUTODNS_USERNAME",
                      "Password": "AUTODNS_PASSWORD",
					  "Endpoint": "https://api.autodns.com/v1", # Optional
					  "Context": 4 # Optional
                    }
                  }
                }
              }
            ]
          }
        ]
      }
    }
  }
}
```

### Caddyfile Configuration

You can also configure the module using the Caddyfile.

#### Global Configuration

```
{
  acme_dns AutoDNS {
    username "<username>"
    password "<password>"
  }
}
```

#### Per-Site Configuration

```
tls {
  dns AutoDNS {
    username "<username>"
    password "<password>"
  }
}
```

## Environment Variables

You can also set the following environment variables to configure the module:

- `AUTODNS_USERNAME`
- `AUTODNS_PASSWORD`
- `AUTODNS_ENDPOINT`
- `AUTODNS_CONTEXT`

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub.

## Acknowledgements

This module is based on the [libdns](https://github.com/libdns/libdns) and [Caddy](https://github.com/caddyserver/caddy) projects.