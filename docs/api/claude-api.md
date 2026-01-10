---
title: Claude API 接口文档
description: Claude API 提供了强大的大语言模型能力，支持消息对话、工具调用、流式响应等功能。本文档详细介绍 Claude API 的核心接口、请求参数、响应格式以及错误处理机制。
tags:
  - api
  - claude
  - llm
version: v1.0.0
status: stable
owner: ai-team@example.com
api_group: claude
---

# Claude API 接口文档

## 概览

Claude API 是 Anthropic 提供的大语言模型 API 服务，基于 Claude 系列模型。它支持多种功能，包括但不限于：

- 多轮对话消息处理
- 工具调用（Tool Use）
- 流式响应（Streaming）
- JSON 输出格式控制
- 上下文窗口管理

本 API 适用于各种应用场景，如智能客服、内容生成、代码辅助、数据分析等。

## 快速开始

要开始使用 Claude API，您需要：

1. 获取 API 密钥（从 Anthropic 控制台）
2. 设置正确的请求头
3. 构造符合规范的请求体

基础请求示例：
```bash
curl https://api.anthropic.com/v1/messages \
  -H "x-api-key: YOUR_API_KEY" \
  -H "anthropic-version: 2023-06-01" \
  -H "content-type: application/json" \
  -d '{
    "model": "claude-3-opus-20240229",
    "max_tokens": 1024,
    "messages": [
      {"role": "user", "content": "你好，Claude！"}
    ]
  }'
```

## 接口列表

### POST /v1/messages

创建新的消息对话或继续现有对话。

#### 请求参数

- **model** (string, required): 使用的模型名称，如 `claude-3-opus-20240229`
- **max_tokens** (integer, required): 最大生成 token 数量
- **messages** (array, required): 消息历史数组
- **system** (string, optional): 系统提示词
- **tools** (array, optional): 可用工具定义
- **tool_choice** (object, optional): 工具选择策略
- **stream** (boolean, optional): 是否启用流式响应

#### 响应字段

- **id** (string): 消息唯一标识符
- **type** (string): 响应类型 (`message` 或 `error`)
- **role** (string): 角色 (`assistant`)
- **content** (array): 内容块数组
- **model** (string): 使用的模型
- **stop_reason** (string): 停止原因
- **usage** (object): Token 使用统计

#### 示例

**基本对话请求：**
```json
{
  "model": "claude-3-haiku-20240307",
  "max_tokens": 1000,
  "messages": [
    {
      "role": "user",
      "content": "请解释什么是 RAG 技术？"
    }
  ]
}
```

**工具调用请求：**
```json
{
  "model": "claude-3-opus-20240229",
  "max_tokens": 1000,
  "tools": [
    {
      "name": "get_weather",
      "description": "获取指定城市的天气信息",
      "input_schema": {
        "type": "object",
        "properties": {
          "city": {"type": "string"}
        },
        "required": ["city"]
      }
    }
  ],
  "messages": [
    {
      "role": "user",
      "content": "北京今天天气怎么样？"
    }
  ]
}
```

### GET /v1/models

获取可用模型列表。

#### 请求参数

无特殊参数，标准认证即可。

#### 响应字段

- **models** (array): 模型信息数组
  - **name** (string): 模型名称
  - **type** (string): 模型类型
  - **context_window** (integer): 上下文窗口大小

## 错误码

### 认证相关错误码

- **401 Unauthorized**: API 密钥无效或缺失
- **403 Forbidden**: API 密钥无访问权限

### 请求相关错误码

- **400 Bad Request**: 请求参数格式错误
- **429 Too Many Requests**: 超出速率限制
- **500 Internal Server Error**: 服务器内部错误

### 模型相关错误码

- **404 Not Found**: 指定的模型不存在
- **422 Unprocessable Entity**: 模型不支持请求的功能

## 常见问题

### 如何处理长上下文？

Claude 模型支持长上下文窗口（最高 200K tokens）。对于超长输入，建议：
1. 合理分段处理
2. 使用摘要技术提取关键信息
3. 利用 Claude 的上下文理解能力

### 如何优化工具调用？

- 提供清晰的工具描述和参数说明
- 限制工具数量避免混淆
- 处理工具调用失败的回退策略

### 流式响应如何处理？

流式响应通过 Server-Sent Events (SSE) 实现，客户端需要：
1. 设置 `Accept: text/event-stream` 头
2. 正确解析 SSE 格式的事件流
3. 处理不同类型的事件（message_start, content_block_delta, message_stop 等）

## 相关文档

- [MCP 服务器集成指南](../mcp-rename/mcp-overview.md)
- [Claude Code 使用手册](../guide/claude-code-guide.md)
- [向量数据库与 RAG 技术](../guide/intro.md)