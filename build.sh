#!/bin/bash

# アプリケーション名
APP_NAME="SystemMonitor.app"
CONTENTS_DIR="$APP_NAME/Contents"
MACOS_DIR="$CONTENTS_DIR/MacOS"
RESOURCES_DIR="$CONTENTS_DIR/Resources"

# 古いビルドを削除
rm -rf "$APP_NAME"

# ディレクトリ構造を作成
mkdir -p "$MACOS_DIR"
mkdir -p "$RESOURCES_DIR"

# アプリケーションをビルド
go build -o "$MACOS_DIR/mac_act_monitor"

# Info.plistをコピー
cp Info.plist "$CONTENTS_DIR/"

# 実行権限を設定
chmod +x "$MACOS_DIR/mac_act_monitor"

echo "ビルドが完了しました: $APP_NAME" 