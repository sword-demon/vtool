# golangci-lint 配置说明

本项目已配置了完整的 `golangci-lint` 规则，用于保证代码质量和风格一致性。

## 安装 golangci-lint

### 方法一：使用安装脚本
```bash
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.64.8
```

### 方法二：使用 go install
```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

### 方法三：使用包管理器
```bash
# macOS
brew install golangci-lint

# Arch Linux
yay -S golangci-lint

# Docker
docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:latest golangci-lint run
```

## 运行检查

### 基本用法
```bash
# 检查所有文件
golangci-lint run

# 检查特定文件
golangci-lint run ./...

# 检查单个包
golangci-lint run ./internal/bean/

# 使用自定义配置文件
golangci-lint run --config .golangci.yml

# 只运行指定的 linters
golangci-lint run --enable=govet,revive,gofmt

# 跳过某些 linters
golangci-lint run --disable=gomnd,lll

# 输出到文件
golangci-lint run -o results.txt --out-format=json

# 检查时修复可修复的问题
golangci-lint run --fix
```

### 常见命令
```bash
# 格式化代码
golangci-lint run --fix

# 检查并输出详细报告
golangci-lint run -v

# 仅检查修改的文件（需要 git）
golangci-lint run --new-from-rev=HEAD~1

# 检查特定目录
golangci-lint run ./internal/bean/

# 跳过某些目录
golangci-lint run --skip-dirs=vendor,examples

# 跳过某些文件
golangci-lint run --skip-files=*_test.go,*.pb.go
```

## 配置说明

### 启用的主要 Linters

| Linter | 描述 | 用途 |
|--------|------|------|
| **govet** | Go 静态分析器 | 检查常见错误 |
| **revive** | 现代的 linter | 代码风格和最佳实践 |
| **gocritic** | 代码诊断工具 | 性能和问题检查 |
| **gocyclo** | 圈复杂度检查 | 控制代码复杂度 |
| **gofmt** | 格式化检查 | 确保代码格式一致 |
| **goimports** | 导入排序和格式化 | 管理 import 语句 |
| **errcheck** | 错误处理检查 | 确保错误被处理 |
| **staticcheck** | 静态分析 | 性能和安全检查 |
| **stylecheck** | 风格检查 | 代码风格一致性 |
| **misspell** | 拼写检查 | 检查拼写错误 |
| **ineffassign** | 无效赋值检查 | 检查未使用的赋值 |
| **unused** | 未使用代码检查 | 删除未使用的代码 |
| **whitespace** | 空白空间检查 | 确保代码整洁 |
| **nolintlint** | nolint 指令检查 | 确保 nolint 正确使用 |
| **depguard** | 依赖检查 | 控制依赖使用 |
| **dupl** | 重复代码检查 | 发现重复代码 |
| **funlen** | 函数长度检查 | 控制函数长度 |
| **gomnd** | 魔数检查 | 避免硬编码数字 |

### 特殊配置

#### 测试文件放宽规则
对于 `_test.go` 文件，以下 linters 会被禁用：
- `errcheck` - 忽略错误处理
- `funlen` - 允许长函数
- `gocyclo` - 允许高复杂度
- `gomnd` - 允许魔数
- `lll` - 允许长行
- `misspell` - 允许拼写错误

#### Internal 包放宽规则
对于 `internal/` 目录中的文件：
- `revive` - 允许未导出的导出符号
- `stylecheck` - 放宽风格检查
- `gocritic` - 允许某些诊断

#### 排除的文件
- `vendor/` - 第三方依赖
- `.git/` - Git 目录
- `.idea/`, `.vscode/` - IDE 配置
- `node_modules/` - Node.js 依赖
- `dist/`, `build/` - 构建产物

## 常见问题解决

### 1. 版本兼容问题
如果遇到版本兼容错误，尝试：
```bash
# 更新到最新版本
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

### 2. 忽略特定问题
使用 `nolint` 指令：
```go
// nolint: gomnd // 忽略魔数检查
const Timeout = 30

// nolint: revive // 忽略导出符号检查
func Helper() {}

// 多 linter 忽略
// nolint: errcheck,gocyclo,lll
func test() {
    // ...
}
```

### 3. 配置文件修改
根据项目需求调整 `.golangci.yml`：
```yaml
# 调整圈复杂度
gocyclo:
  min-complexity: 20  # 原来是 15

# 调整字符串最小长度
goconst:
  min-len: 3  # 原来是 2

# 添加排除规则
exclude-rules:
  - path: internal/test/
    linters:
      - gomnd
      - lll
```

### 4. 性能优化
```bash
# 增加并发数
golangci-lint run --concurrency=8

# 跳过某些检查
golangci-lint run --disable=gosec,gomnd

# 只检查修改的文件
golangci-lint run --new-from-rev=HEAD~5
```

## 集成到 CI/CD

### GitHub Actions 示例
```yaml
name: Lint

on: [push, pull_request]

jobs:
  lint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: "1.23"
      - name: Run golangci-lint
        uses: golangci/golangci-lint-action@v6
        with:
          version: latest
          args: --timeout=5m
```

### GitLab CI 示例
```yaml
lint:
  image: golang:1.23
  script:
    - go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
    - golangci-lint run --timeout=5m
  only:
    changes:
      - "**/*.go"
```

### pre-commit 钩子示例
```yaml
repos:
  - repo: https://github.com/golangci/golangci-lint-action
    rev: v6.0.1
    hooks:
      - id: golangci-lint
        args: [--timeout=5m]
```

## 扩展配置

### 添加自定义规则
```yaml
linters-settings:
  gocritic:
    enabled-checks:
      - hugeParam
      - rangeValCopy
```

### 调整排除规则
```yaml
issues:
  exclude-rules:
    # 忽略特定错误
    - linters:
        - gocritic
      text: "unslice"

    # 忽略特定文件
    - path: "mocks/"
      linters:
        - gomnd
        - funlen
```

## 参考资料

- [golangci-lint 官方文档](https://golangci-lint.run/)
- [Linters 列表](https://golangci-lint.run/usage/linters/)
- [配置选项](https://golangci-lint.run/usage/configuration/)
- [最佳实践](https://github.com/golangci/golangci-lint/wiki/Best-practices)

---

💡 **提示**：定期运行 `golangci-lint run --fix` 可以自动修复大部分格式问题！