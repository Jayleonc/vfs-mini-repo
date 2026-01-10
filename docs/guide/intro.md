---
title: 向量数据库与 RAG 技术
description: 向量数据库是支持高效向量相似度检索的专用数据库，在RAG（检索增强生成）架构中起着至关重要的作用。本文档详细介绍向量数据库的工作原理、文本向量化过程、余弦相似度计算，以及如何通过RAG架构解决大语言模型幻觉问题并提高回答准确性。
tags:
  - vector-database
  - rag
  - retrieval
version: v1.0.0
status: stable
owner: ai-research@example.com
api_group: ai-architecture
---

# 向量数据库与 RAG 技术

向量数据库是支持高效向量相似度检索的专用数据库，在 RAG（检索增强生成）架构中起着至关重要的作用。它将文本转换为高维向量存储，当用户提问时，通过计算余弦相似度找到最相关的上下文，再交由大语言模型生成回答，有效解决了模型幻觉问题并提高了回答的准确性。

## 相关文档

- [Claude API 接口文档](/docs/api/claude-api.md)
- [MCP 服务器概述与集成指南](/docs/mcp/mcp-overview.md)
- [Claude Code 使用手册](/docs/guide/claude-code-guide.md)
