# CLAUDE.md

本文档为 Claude Code (claude.ai/code) 在本仓库中处理代码时提供指导。

## 仓库概述

这是一个用于 Claude 开发者工具的文档仓库，遵循 VFS 2.0 知识工程标准。仓库包含：
- 最小化的源代码（`src/main.go` - 一个简单的 Go "hello world" 示例）
- `/docs/` 目录下的完整文档，涵盖 Claude API、MCP 协议、Claude Code CLI 和技术指南
- 文档文件按主题组织，具有标准化的 YAML front matter 和内部链接

## 开发命令

由于这主要是一个文档仓库，没有传统的构建/测试命令。但常见操作包括：

### 文档验证
- 验证内部链接是否遵循 VFS 2.0 标准（以 `/` 开头的绝对路径，内部链接不使用 HTTP URL）
- 确保所有文档文件都具有适当的 YAML front matter，包含 `title`、`description`、`tags`、`version`、`status` 和 `owner`
- 检查标题结构：每个文档一个 H1 标题，有意义的 H2 章节，对于 API 文档，H3 端点格式为 `METHOD /path`

### 文件管理
- 新文档应放置在 `/docs/` 下的适当子目录中
- 内部链接必须使用绝对路径（例如 `/docs/api/claude-api.md`）
- 锚点链接应与目标标题完全匹配（例如 `/docs/api/claude-api.md#错误码`）

## 代码架构

仓库具有最小化的代码结构：
- `src/main.go`：打印 "VFS Mini Repo" 的简单 Go 程序
- `docs/`：主文档目录，包含不同主题的子目录
  - `api/`：API 参考文档
  - `guide/`：用户指南和教程
  - `mcp-rename/`：MCP（Model Context Protocol）文档
  - 编号目录如 `01-机器人相关/`：中文文档部分

## VFS 2.0 标准合规性

所有文档必须遵守 `standards_checklist.md` 中概述的标准：
- **内部链接**：使用绝对路径（`/docs/guide/install.md`），不使用 HTTP URL
- **YAML Front Matter**：所有文档必需包含完整的元数据
- **标题结构**：清晰的 H1-H3 层次结构，标题唯一
- **锚点**：精确的章节链接，标题完全匹配
- **文件组织**：反映内容关系的逻辑目录结构

在添加或修改文档时，始终根据 VFS 2.0 检查清单进行验证，以确保与知识工程系统的兼容性。