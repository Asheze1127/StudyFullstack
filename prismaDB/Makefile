# ==============================
# Prisma 用 Makefile
# ==============================
# 使い方:
#   make help
#   make fmt
#   make mig-dev NAME=add_users
# ==============================

SHELL := /bin/bash

# ------------------------------
# 設定
# ------------------------------

# Prisma CLI（npx 経由で実行）
PRISMA_BIN := npx prisma

# schema.prisma のパス
SCHEMA := schema.prisma

# 環境変数ファイル（Prisma が自動で読むが、明示用）
ENV ?= .env

# ------------------------------
# ヘルプ表示
# ------------------------------
.PHONY: help
help: ## 利用可能なコマンド一覧を表示
	@awk 'BEGIN{FS=":.*##"} /^[a-zA-Z0-9_.-]+:.*##/{printf "\033[36m%-20s\033[0m %s\n",$$1,$$2}' $(MAKEFILE_LIST)

# ------------------------------
# Prisma 基本情報
# ------------------------------
.PHONY: prisma
prisma: ## Prisma のバージョンを表示
	@$(PRISMA_BIN) -v

# ------------------------------
# schema.prisma の品質チェック
# ------------------------------
.PHONY: fmt
fmt: ## schema.prisma をフォーマット
	@$(PRISMA_BIN) format --schema $(SCHEMA)

.PHONY: validate
validate: ## schema.prisma を検証（構文・整合性チェック）
	@$(PRISMA_BIN) validate --schema $(SCHEMA)

.PHONY: check
check: validate ## validate のエイリアス

# ------------------------------
# Prisma Client 生成
# ------------------------------
.PHONY: gen
gen: ## Prisma Client を生成
	@$(PRISMA_BIN) generate --schema $(SCHEMA)

.PHONY: gen-watch
gen-watch: ## schema.prisma 変更時に自動 generate（watchexec 必須）
	@command -v watchexec >/dev/null 2>&1 || { echo "watchexec をインストールしてください"; exit 1; }
	@watchexec -w prisma -e prisma -- $(PRISMA_BIN) generate --schema $(SCHEMA)

# ------------------------------
# データベース同期
# ------------------------------
.PHONY: db-push
db-push: ## schema を DB に直接反映（開発専用）
	@$(PRISMA_BIN) db push --schema $(SCHEMA)

.PHONY: db-pull
db-pull: ## DB 定義を schema.prisma に取り込む
	@$(PRISMA_BIN) db pull --schema $(SCHEMA)

.PHONY: studio
studio: ## Prisma Studio を起動
	@$(PRISMA_BIN) studio --schema $(SCHEMA)

# ------------------------------
# マイグレーション
# ------------------------------
NAME ?= change

.PHONY: mig-dev
mig-dev: ## マイグレーション作成＋適用（開発環境） NAME=xxx
	@$(PRISMA_BIN) migrate dev --schema $(SCHEMA) --name $(NAME)

.PHONY: mig-deploy
mig-deploy: ## マイグレーションを適用（本番・CI 用）
	@$(PRISMA_BIN) migrate deploy --schema $(SCHEMA)

.PHONY: mig-reset
mig-reset: ## DB を初期化（⚠ 開発環境専用）
	@echo "⚠ データベースを全削除します。続行しますか？ (y/N)"
	@read ans; \
	if [[ $$ans == "y" ]]; then \
		$(PRISMA_BIN) migrate reset --schema $(SCHEMA); \
	else \
		echo "キャンセルしました"; \
	fi

# ------------------------------
# シードデータ投入
# ------------------------------
.PHONY: seed
seed: ## prisma.seed を実行
	@$(PRISMA_BIN) db seed --schema $(SCHEMA)

# ------------------------------
# 開発環境 初期セットアップ
# ------------------------------
.PHONY: dev-setup
dev-setup: fmt validate gen db-push seed ## 開発環境を一括セットアップ

# ------------------------------
# 生成物の削除
# ------------------------------
.PHONY: clean
clean: ## 生成された Prisma Client を削除
	@rm -rf node_modules/.prisma
	@echo "node_modules/.prisma を削除しました"
