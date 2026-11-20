# Security policy

## Reporting a vulnerability

Please do not open a public issue for a security problem.

Report it through [GitHub private vulnerability reporting](https://github.com/MohZeh/Darban-WebPanel/security/advisories/new). Include the affected version, the steps to reproduce it, and its impact. Do not include real passwords, two-factor secrets, profiles or link addresses.

This is a project maintained in spare time, so responses are best effort. I will acknowledge a report as soon as I can, and publish the fix and advisory together.

## Supported versions

There is no release yet. Once there is, only the latest release receives security fixes.

## The design this policy holds the panel to

The panel can run everything except installing and removing the server: clients, settings, backup, restore and updates. Whoever controls it controls the server, as with SSH. The design keeps it no weaker than SSH:

- The panel never runs as root, and never reads the CA, private keys or profiles. Everything it knows comes from Darban's machine interface.
- What the panel may do is granted in Darban's SSH menu. The panel cannot change its own access.
- Revoking a client, a setting that restarts OpenVPN, and every backup, restore and update need a fresh two-factor code. Darban itself checks it, with a secret the panel never sees, and accepts each code once only.
- A setting that restarts OpenVPN is rolled back by the server itself unless it is confirmed within two minutes.
- Nothing takes a file path or a web address from the panel. Uploads are handled by ID, only encrypted backups can be downloaded, and an update installs only a version Darban offers from its official releases.
- The private key of the backups never reaches the server.
- By default the panel listens on the server itself only, and is reached through an SSH tunnel. Serving it on the internet is not supported.
- There is no default password, and no header from the client, such as `X-Forwarded-For`, is trusted.
- Installing, removing, changing the panel's access and resetting a lost second factor are possible over SSH only.

## In scope

- Signing in without the password or the second factor
- A risky action done without a fresh code, or with a code used before
- A setting that should have been rolled back staying in place without a confirmation
- Reading or changing anything from another origin, or from an unauthenticated request
- Doing through the panel anything its granted access does not allow, or widening that access
- A file path or web address chosen through the panel reaching Darban
- Passwords, codes, cookies or link addresses reaching logs
- A tampered or damaged release being installed by accident

## Out of scope

- An attacker who already has root on the server
- Vulnerabilities in Darban OVPN itself: report those [to that project](https://github.com/MohZeh/Darban-OVPN/security/advisories/new)
