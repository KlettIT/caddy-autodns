# AutoDNS Module for Caddy

This package provides a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It allows you to manage DNS records with AutoDNS.

## Caddy Module Name

```
dns.providers.autodns
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
                      "name": "autodns",
                      "Username": "<AUTODNS_USERNAME>",
                      "Password": "<AUTODNS_PASSWORD>",
					  "Endpoint": "https://api.autodns.com/v1", //Optional
					  "Context": 4 //Optional
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
  acme_dns autodns {
    username "<username>"
    password "<password>"
  }
}
```

#### Per-Site Configuration

```
tls {
  dns autodns {
    username "<username>"
    password "<password>"
  }
}
```

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub.

## Acknowledgements

This module is based on the [libdns](https://github.com/libdns/libdns) and [Caddy](https://github.com/caddyserver/caddy) projects.