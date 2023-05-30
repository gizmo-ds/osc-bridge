# OSC Bridge

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gizmo-ds/osc-bridge?style=flat-square)
[![CI](https://img.shields.io/github/actions/workflow/status/gizmo-ds/osc-bridge/release.yml?branch=main&label=CI&style=flat-square)](https://github.com/gizmo-ds/osc-bridge/actions/workflows/release.yml)
[![Release](https://img.shields.io/github/v/release/gizmo-ds/osc-bridge.svg?include_prereleases&style=flat-square)](https://github.com/gizmo-ds/osc-bridge/releases/latest)
[![License](https://img.shields.io/github/license/gizmo-ds/osc-bridge?style=flat-square)](./LICENSE)

OSC Bridge solves the problem of VRChat's inability to listen to multiple OSC ports simultaneously. This project revolutionizes the way users interact with VRChat by providing a seamless communication bridge for OSC protocols.

With OSC Bridge, developers can create custom interactions, synchronize audiovisual elements, and integrate external devices into VRChat environments. Experience the freedom to create, connect, and interact like never before in VRChat with OSC Bridge.

## Usage

You can download the precompiled files from [Release](https://github.com/gizmo-ds/osc-bridge/releases/latest), unzip them, and run the application.

## Configuration

| Key              | Description                                                     |
| ---------------- | --------------------------------------------------------------- |
| addr             | The address to listen for OSC messages.                         |
| bridges.enable   | Indicates whether the bridges are enabled or not. Default: true |
| bridges.name     | The name of the bridge.                                         |
| bridges.addr     | The address of the bridge.                                      |
| bridges.patterns | The OSC patterns to be matched by the bridge.                   |

## Thanks

Thanks to [JetBrains](https://jb.gg/OpenSourceSupport) for the open source license(s).

![JetBrains Logo (Main) logo](https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.svg)
