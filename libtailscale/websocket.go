// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package libtailscale

// Fall back to WebSockets for the control and DERP connections when an HTTP
// proxy, such as Cloudflare's, refuses their own HTTP upgrades.
import (
	_ "tailscale.com/feature/controlwebsocket"
	_ "tailscale.com/feature/derpwebsocket"
)
