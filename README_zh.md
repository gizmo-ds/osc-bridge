# OSC Bridge

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gizmo-ds/osc-bridge?style=flat-square)
[![CI](https://img.shields.io/github/actions/workflow/status/gizmo-ds/osc-bridge/release.yml?branch=main&label=CI&style=flat-square)](https://github.com/gizmo-ds/osc-bridge/actions/workflows/release.yml)
[![Release](https://img.shields.io/github/v/release/gizmo-ds/osc-bridge.svg?include_prereleases&style=flat-square)](https://github.com/gizmo-ds/osc-bridge/releases/latest)
[![License](https://img.shields.io/github/license/gizmo-ds/osc-bridge?style=flat-square)](./LICENSE)

[Engine](./README.md) | 简体中文

OSC Bridge 解决了 VRChat 无法同时监听多个 OSC 端口的问题。这个项目通过提供无缝的 OSC 协议通信桥梁，彻底改变了用户与 VRChat 互动的方式。

借助 OSC Bridge，开发者可以创建自定义互动、同步音频视觉元素，并将外部设备整合到 VRChat 环境中。在 VRChat 中，通过 OSC Bridge，您将体验到前所未有的创作、连接和互动的自由。

## 使用

你可以通过 [Release](https://github.com/gizmo-ds/osc-bridge/releases/latest) 下载预编译的文件，解压缩并运行。

## 配置选项

| 键               | 描述                         |
| ---------------- | ---------------------------- |
| addr             | 监听 OSC 消息的地址。        |
| bridges.enable   | 指示是否启用桥接。默认为启用 |
| bridges.name     | 桥接的名称。                 |
| bridges.addr     | 桥接的地址。                 |
| bridges.patterns | 桥接要匹配的 OSC 模式。      |

> `bridges.patterns` 配置项仅

## 感谢

感谢 [JetBrains](https://jb.gg/OpenSourceSupport) 提供的开源许可证。

![JetBrains Logo (Main) logo](https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.svg)
