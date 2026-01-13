---
description: 根据当前暂存区的代码变更，生成一条符合Conventional Commits规范的Commit Message。
allowed-tools: Bash(git add:*)
---

你是一位Git专家。请根据以下代码变更的diff信息，为我生成一条符合Conventional Commits规范的、高质量的`git commit`消息

**当前分支:**
!`git branch --show-current`

**暂存区变更 (Staged Changes):**
!`git diff --staged`

请确保输出为单行文本，避免换行或项目符号。
不需要包含创建者/作者信息（例如 author、Created by、Co-authored-by 等）。
请只输出commit message本身，不要有任何额外的解释。
