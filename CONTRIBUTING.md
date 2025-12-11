# 贡献指南

感谢您对 vtool 项目的兴趣！我们欢迎所有形式的贡献。

## 目录

- [行为准则](#行为准则)
- [如何贡献](#如何贡献)
- [开发环境设置](#开发环境设置)
- [代码规范](#代码规范)
- [测试](#测试)
- [文档](#文档)
- [提交规范](#提交规范)
- [问题反馈](#问题反馈)
- [Pull Request 流程](#pull-request-流程)

## 行为准则

请阅读并遵守我们的[行为准则](CODE_OF_CONDUCT.md)（如果存在）。

## 如何贡献

您可以通过以下方式为项目做贡献：

1. 🐛 报告 Bug
2. 💡 提出新功能建议
3. 📖 改进文档
4. 💻 提交代码修复或新功能
5. 🔧 帮助维护项目

## 开发环境设置

### 前置要求

- Go 1.23 或更高版本
- Git

### 克隆项目

```bash
git clone https://github.com/sword-demon/vtool.git
cd vtool
```

### 安装开发工具

```bash
# 安装所有开发工具
make devtools

# 或者单独安装
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install golang.org/x/tools/cmd/goimports@latest
```

### 验证环境

```bash
# 运行所有测试
make test

# 运行代码检查
make lint

# 构建项目
make build
```

## 代码规范

### 基本规范

- 遵循 [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- 使用 `gofmt` 格式化代码
- 遵循 `golint` 或 `revive` 的建议
- 使用中文注释（导出函数和类型必须有注释）
- 避免复杂的函数（圈复杂度 < 15）

### 代码风格

1. **命名规范**
   - 使用有意义的变量名
   - 导出函数使用驼峰命名
   - 私有变量使用驼峰命名
   - 常量使用全大写加下划线

2. **注释规范**
   ```go
   // Add 在指定位置添加元素
   // 返回新的切片和可能的错误
   func Add[T any](slice []T, element T, index int) ([]T, error) {
       // ...
   }
   ```

3. **导入规范**
   ```go
   import (
       "fmt"
       "strings"

       "github.com/example/package"
   )
   ```

### 使用 linters

```bash
# 运行所有检查
make check

# 自动修复可修复的问题
make lint-fix

# 只运行特定的 linter
golangci-lint run --enable=govet,revive
```

## 测试

### 编写测试

- 所有公共 API 必须有测试
- 测试文件命名为 `*_test.go`
- 使用 `table-driven` 测试模式
- 测试函数名使用 `TestFunctionName`

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        slice    []int
        element  int
        index    int
        want     []int
        wantErr  bool
    }{
        {
            name:     "在中间位置添加",
            slice:    []int{1, 2, 4, 5},
            element:  3,
            index:    2,
            want:     []int{1, 2, 3, 4, 5},
            wantErr:  false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Add(tt.slice, tt.element, tt.index)
            if (err != nil) != tt.wantErr {
                t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("Add() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

### 运行测试

```bash
# 运行所有测试
make test

# 运行测试并显示覆盖率
make test-coverage

# 运行特定包的测试
make test-package PKG=internal/bean

# 运行基准测试
make bench
```

### 测试覆盖率

- 公共 API 测试覆盖率应达到 90% 以上
- 使用 `make test-coverage` 生成覆盖率报告

## 文档

### 编写文档

- 所有导出的函数、类型、常量必须有注释
- 注释应清楚说明用途、参数和返回值
- 复杂算法需要说明实现思路

```go
// HashMap 基于 Go 内置 map 实现的增强版映射
// 提供可预测的迭代顺序
type HashMap[K comparable, V any] struct {
    // ...
}
```

### 生成文档

```bash
# 生成 API 文档（如果使用 swag）
make docs

# 查看文档
go doc ./internal/bean

# 启动本地文档服务器
godoc -http=:6060
```

## 提交规范

我们使用 [Conventional Commits](https://www.conventionalcommits.org/) 规范：

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

### 类型 (type)

- `feat`: 新功能
- `fix`: Bug 修复
- `docs`: 文档更新
- `style`: 代码格式调整
- `refactor`: 重构
- `perf`: 性能优化
- `test`: 测试相关
- `chore`: 构建或工具相关

### 示例

```bash
feat(bean): 添加结构体深度复制功能

实现 deepCopyValue 函数，支持：
- 指针字段的递归复制
- 切片的深度复制
- Map 的深度复制

Closes #123
```

## 问题反馈

### Bug 报告

使用 [Bug 报告模板](.github/ISSUE_TEMPLATE/bug_report.md)（如果存在）：

```markdown
**Bug 描述**
简要描述 Bug

**复现步骤**
1. 运行 '...'
2. 使用参数 '...'
3. 看到错误 '...'

**期望行为**
清楚描述期望的行为

**实际行为**
描述实际发生的情况

**环境信息**
- OS: [e.g. macOS]
- Go 版本: [e.g. 1.23]
```

### 功能请求

使用 [功能请求模板](.github/ISSUE_TEMPLATE/feature_request.md)（如果存在）：

```markdown
**功能描述**
简要描述所需功能

**动机**
为什么需要这个功能？

**详细描述**
详细描述功能需求

**替代方案**
是否有其他解决方案？
```

## Pull Request 流程

### 1. Fork 项目

```bash
git clone https://github.com/sword-demon/vtool.git
```

### 2. 创建分支

```bash
git checkout -b feature/your-feature-name
# 或者
git checkout -b fix/bug-description
```

### 3. 编写代码

- 遵循代码规范
- 编写测试
- 确保所有测试通过

```bash
make check
make test
```

### 4. 提交代码

```bash
git add .
git commit -m "feat(bean): 添加深度复制功能"
```

### 5. 推送分支

```bash
git push origin feature/your-feature-name
```

### 6. 创建 PR

在 GitHub 上创建 Pull Request，包含：

1. 清晰的标题和描述
2. 关联的 Issue（如果有）
3. 测试结果
4. 截图或示例（如果适用）

### PR 检查清单

- [ ] 代码遵循项目规范
- [ ] 添加了必要的测试
- [ ] 所有测试通过
- [ ] 运行 `make check` 通过
- [ ] 更新了文档（如需要）
- [ ] PR 描述清晰完整

### PR 模板

```markdown
## 描述
简要描述 PR 的变更

## 变更类型
- [ ] Bug 修复
- [ ] 新功能
- [ ] 破坏性变更
- [ ] 文档更新

## 测试
- [ ] 添加了单元测试
- [ ] 所有测试通过
- [ ] 手动测试通过

## 清单
- [ ] 代码遵循代码规范
- [ ] 自审查代码
- [ ] 添加了适当的注释
- [ ] 更新了相关文档

## 相关 Issue
Closes #123
```

### 代码审查

审查者会检查：

1. 代码质量和规范
2. 测试覆盖率
3. 文档完整性
4. 性能影响
5. 安全性

### 合并

PR 符合以下条件才会被合并：

- 所有 CI 检查通过
- 至少一名审查者批准
- 没有阻塞性问题

## 开发提示

### 常用命令

```bash
# 开发
make fmt          # 格式化代码
make vet          # 代码检查
make lint         # 运行 linters
make test         # 运行测试

# 构建
make build        # 构建
make clean        # 清理

# 高级
make test-coverage # 测试覆盖率
make cyclo        # 圈复杂度检查
make security     # 安全检查
```

### 有用的工具

- `go doc` - 查看文档
- `go test -v` - 详细测试输出
- `go test -race` - 竞态条件检测
- `go build -race` - 构建带竞态检测的版本

## 常见问题

### Q: 如何运行特定测试？
A: 使用 `go test -v ./path/to/package`

### Q: 如何调试测试？
A: 使用 `go test -v -run TestName`

### Q: 如何检查测试覆盖率？
A: 运行 `make test-coverage`

### Q: linter 报错怎么办？
A: 查看 [LINTER.md](LINTER.md) 了解如何处理

## 感谢

感谢所有为 vtool 项目做出贡献的开发者！您的贡献使这个项目变得更好。

## 许可证

通过贡献代码，您同意您的贡献将在 MIT 许可证下许可。