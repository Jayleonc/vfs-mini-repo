---
title: Claude 开发者文档中心
description: 本知识库包含 Claude 相关的完整开发者文档，涵盖 Claude API、MCP 服务器协议、Claude Code 工具以及相关技术指南。所有文档都遵循 VFS 2.0 知识工程标准，支持高效的链接寻址和内容检索。
tags:
  - documentation
  - claude
  - developer
version: v1.0.0
status: stable
owner: docs-team@example.com
---

# Claude 开发者文档中心

## 概览

欢迎使用 Claude 开发者文档中心！本知识库提供了 Claude 生态系统的完整文档，包括：

- **Claude API**：大语言模型的核心接口
- **MCP 协议**：标准化的工具调用协议
- **Claude Code**：智能开发命令行工具
- **技术指南**：最佳实践和深度技术解析

所有文档都遵循 [VFS 2.0 知识工程标准](/standards_checklist.md)，确保最佳的可读性和可维护性。

## 快速导航

### API 文档

- [Claude API 接口文档](api/claude-api.md) - 核心 API 参考
- [错误码参考](api/claude-api.md#错误码) - 完整的错误处理指南

### 协议与集成

- [MCP 服务器概述与集成指南](mcp/mcp-overview.md) - Model Context Protocol 详解
- [MCP 工具设计最佳实践](mcp/mcp-overview.md#最佳实践) - 工具开发指南

### 开发工具

- [Claude Code 使用手册](guide/claude-code-guide.md) - CLI 工具完整指南
- [高级功能配置](guide/claude-code-guide.md#高级功能) - 自定义和扩展

### 技术指南

- [向量数据库与 RAG 技术](guide/intro.md) - 检索增强生成技术
- [深度学习基础知识](readme.md) - AI 基础概念

## 文档标准

本知识库严格遵循以下标准：

- **内部链接**：使用绝对路径格式，如 `/docs/api/claude-api.md`
- **YAML Front Matter**：每个文档都有完整的元数据信息
- **标题结构**：清晰的 H1-H3 层次结构
- **锚点导航**：所有章节都有精确的锚点链接

## 贡献指南

如果您希望为本文档做出贡献，请遵循以下步骤：

1. 克隆仓库并创建新分支
2. 按照 [VFS 2.0 文档自检说明](/standards_checklist.md) 编写文档
3. 确保所有内部链接使用正确的绝对路径格式
4. 添加完整的 YAML Front Matter
5. 提交 Pull Request 进行审查

## 相关资源

- [VFS 2.0 文档自检说明](/standards_checklist.md) - 文档编写标准
- [Claude 官方文档](https://docs.anthropic.com) - 官方 API 文档
- [MCP 协议规范](https://github.com/anthropic/mcp) - 开源协议实现

## 版本信息

- **文档版本**：v1.0.0
- **最后更新**：2026-01-10
- **维护团队**：Developer Experience Team