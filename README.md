# Darban WebPanel

[فارسی](README.fa.md)

A web panel for [Darban OVPN](https://github.com/MohZeh/Darban-OVPN): manage clients and temporary links from a browser instead of an SSH menu.

**Status: not ready.** This repository holds the project's structure and CI only. Work starts after Darban OVPN `v1.1.0`, whose machine interface the panel is built on. Nothing here should be installed yet.

## Design

- **It runs beside Darban, not instead of it.** Darban OVPN stays the only part that changes the server. The panel asks it through Darban's machine interface, and never reads the CA, private keys or profiles.
- **It never runs as root.** It runs as a system user of its own, with a small set of allowed actions: list, create, share, revoke and end a session. Settings, backups, updates and removal stay in the SSH menu.
- **It is not on the internet.** By default it listens on the server itself, and you reach it through an SSH tunnel. Serving it publicly is not supported.
- **No default password.** Setup makes a random one, and a second factor.
- **One static file, no runtime.** Written in Go, with pages rendered on the server, in English and Persian. Nothing is loaded from other sites.

## License

[MIT](LICENSE). Darban WebPanel is an independent project, not affiliated with or endorsed by OpenVPN Inc. OpenVPN is a registered trademark of OpenVPN Inc.
