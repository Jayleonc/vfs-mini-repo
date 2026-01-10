---
title: Python 并发编程指南
description: Python 并发编程通过多线程、多进程和异步IO三种主要方式实现。本文档详细分析全局解释器锁（GIL）对多线程的影响，介绍 multiprocessing 模块在CPU密集型任务中的应用，以及 asyncio 在高并发网络服务中的优势和最佳实践。
tags:
  - python
  - concurrency
  - asyncio
version: v1.0.0
status: stable
owner: python-team@example.com
api_group: programming
---

# Python 并发编程指南

Python 中的并发编程主要通过多线程、多进程和异步 IO 实现。由于全局解释器锁（GIL）的存在，多线程在处理计算密集型任务时受到限制，但在 IO 密集型场景下表现良好。对于 CPU 密集型任务，建议使用 multiprocessing 模块利用多核性能。而 asyncio 则是现代异步编程的核心，能显著提升高并发网络服务的效率。

## 相关文档

- [Claude API 接口文档](api/claude-api.md)
- [MCP 服务器概述与集成指南](mcp-rename/mcp-overview.md)
- [Claude Code 使用手册](guide/claude-code-guide.md)
