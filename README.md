# Darban WebPanel

[فارسی](README.fa.md)

A web panel for [Darban OVPN](https://github.com/MohZeh/Darban-OVPN): run the whole server from a browser instead of an SSH menu. Clients, temporary links, every setting, backup, restore and updates all work from the panel. Only installing and removing stay in SSH.

**Status: not ready.** This repository holds the project's structure and CI only. Work starts after Darban OVPN `v1.1.0`, whose machine interface the panel is built on. Nothing here should be installed yet.

## Design

- **It runs beside Darban, not instead of it.** Darban OVPN stays the only part that changes the server. The panel asks it through Darban's machine interface, and never reads the CA, private keys or profiles.
- **It never runs as root.** It runs as a system user of its own. What it may do is granted in Darban's SSH menu, and the panel cannot widen it.
- **Risky actions need a fresh code, checked by Darban itself.** Revoking a client, a setting that restarts OpenVPN, backup, restore and updates each ask for a new two-factor code. The secret behind those codes stays with Darban, so even a taken-over panel cannot do them alone.
- **A risky setting undoes itself unless you confirm it.** You see exactly what will change first. After it is applied, you have two minutes to confirm that you can still reach the server; otherwise the server goes back to the old setting by itself.
- **No file paths and no web addresses.** A restore takes an uploaded file, and an update takes only a version that Darban offers from its official releases. The private key of your backups never reaches the server.
- **It is not on the internet.** By default it listens on the server itself, and you reach it through an SSH tunnel. Serving it publicly is not supported.
- **No default password.** Setup makes a random one, and a second factor.
- **SSH stays the way back.** Installing, removing, changing what the panel may do, and resetting a lost second factor are done over SSH only.
- **One static file, no runtime.** Written in Go, with pages rendered on the server, in English and Persian. Nothing is loaded from other sites.

## License

[MIT](LICENSE). Darban WebPanel is an independent project, not affiliated with or endorsed by OpenVPN Inc. OpenVPN is a registered trademark of OpenVPN Inc.
