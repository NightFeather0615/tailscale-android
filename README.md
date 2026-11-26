# Tailscale for Android with WebSocket fallbacks

Builds of the open source
[Tailscale Android app](https://github.com/tailscale/tailscale-android) that
can reach a control server through an HTTP proxy that only lets WebSocket
upgrades through. The case it was made for is a
[Headscale](https://headscale.net) server behind a Cloudflare Tunnel.

This is an unofficial fork for personal use. It is not made, endorsed, or
supported by Tailscale Inc. If you don't need the change below, use the
official app.

## What is different

The app falls back to WebSockets for the control and DERP connections when a
proxy refuses Tailscale's own HTTP upgrade. The change itself is in
[NightFeather0615/tailscale](https://github.com/NightFeather0615/tailscale),
and is proposed upstream in
[tailscale/tailscale#22643](https://github.com/tailscale/tailscale/pull/22643).

Everything else is the upstream release of the same version number.

## Install

The APK is on the
[releases page](https://github.com/NightFeather0615/tailscale-android/releases).

*   **The official app has to go first.** This build has the same package
    name (`com.tailscale.ipn`) and another signing key, so Android refuses to
    install it over the official app, or over a debug build. Uninstalling
    removes the login, so the device has to sign in again.
*   **Updates:** install a newer APK over the old one, or add
    `https://github.com/NightFeather0615/tailscale-android` to
    [Obtainium](https://github.com/ImranR98/Obtainium), which then offers each
    new release.

Releases are signed with a key whose certificate has this SHA-256
fingerprint:

```text
51:92:05:6A:57:53:A6:DC:14:AF:33:D6:44:64:66:5A:18:58:BB:87:1A:47:50:33:87:DF:19:AC:3E:82:5C:5A
```

`apksigner verify --print-certs` and the release notes show it without the
colons, in lower case.

### First login

Before signing in, open the account screen, and choose "Use an alternate
server" in its menu. Enter the address of the control server, such as
`https://headscale.example.com`.

## How releases are made

The [Release workflow](.github/workflows/release.yml) runs every day. When
upstream has a new stable release, it:

1.  Puts this repository's patches on the upstream release tag.
1.  Takes the tailscale.com version that tag is built from, and puts the
    patches of NightFeather0615/tailscale on it.
1.  Builds the app against that tree, and signs the APK.
1.  Publishes a release named after the upstream version plus a build
    number, such as `1.106.1-nf.1`.

A release holds the APK, a `SHA256SUMS` file, and `.patch` files with the
changes made to both upstream sources. Its tag points at this branch, which
holds the tooling.

## Branches

| Branch | Holds |
| --- | --- |
| `fork` | This README and the release workflow. |
| `stack` | The patches, as commits on top of an upstream release tag. |
| `main` | A copy of upstream's `main`. |

## Maintaining

The workflow fails when a set of patches no longer applies to a new release,
and says which one. For this repository's:

```shell
git fetch upstream --tags
git switch stack
git rebase --onto <new tag> <the tag stack was on>
git push --force-with-lease origin stack
gh workflow run release.yml
```

For tailscale.com's, see the README of NightFeather0615/tailscale.

To try a build without publishing it, run the workflow with `publish` turned
off; the files are attached to the workflow run.

The signing key is in the repository secrets `ANDROID_KEYSTORE_BASE64` (a
PKCS #12 keystore with one key, base64 encoded) and
`ANDROID_KEYSTORE_PASSWORD`. Keep a copy of the keystore: an APK signed with
another key can't update an installed one.

## License

BSD 3-Clause, the same as upstream; see [LICENSE](LICENSE). Tailscale is a
registered trademark of Tailscale Inc. WireGuard is a registered trademark of
Jason A. Donenfeld.
