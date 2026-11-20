# Security policy

## Reporting a vulnerability

Please do not open a public issue for a security problem.

Report it through [GitHub private vulnerability reporting](https://github.com/MohZeh/Darban-WebPanel/security/advisories/new). Include the affected version, the steps to reproduce it, and its impact. Do not include real passwords, two-factor secrets, profiles or link addresses.

This is a project maintained in spare time, so responses are best effort. I will acknowledge a report as soon as I can, and publish the fix and advisory together.

## Supported versions

There is no release yet. Once there is, only the latest release receives security fixes.

## The design this policy holds the panel to

- The panel never runs as root, and never reads the CA, private keys or profiles. Everything it knows comes from Darban's machine interface.
- That interface gives the panel no way to change settings, restore a backup, update or remove the server, or touch the CA.
- By default the panel listens on the server itself only, and is reached through an SSH tunnel. Serving it on the internet is not supported.
- There is no default password, and no header from the client, such as `X-Forwarded-For`, is trusted.

## In scope

- Signing in without the password or the second factor
- Reading or changing anything from another origin, or from an unauthenticated request
- Doing through the panel anything its role does not allow
- Passwords, codes, cookies or link addresses reaching logs
- A tampered or damaged release being installed by accident

## Out of scope

- An attacker who already has root on the server
- Vulnerabilities in Darban OVPN itself: report those [to that project](https://github.com/MohZeh/Darban-OVPN/security/advisories/new)
