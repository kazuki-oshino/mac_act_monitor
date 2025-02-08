# Mac Activity Monitor

macOS用のメニューバーアプリケーションで、システムリソースの使用状況をリアルタイムで表示します。

## 機能

- メニューバーに常駐して動作
- CPU使用率のリアルタイム表示
- メモリ使用率のリアルタイム表示
- 2秒ごとの自動更新
- Dockアイコンなしで動作（メニューバー専用アプリケーション）

## 必要要件

- macOS 10.10以上
- Go 1.16以上

## インストール方法

### ソースコードからビルド

1. リポジトリをクローン
```bash
git clone https://github.com/[your-username]/mac_act_monitor.git
cd mac_act_monitor
```

2. 依存パッケージのインストール
```bash
go mod tidy
```

3. アプリケーションのビルド
```bash
chmod +x build.sh
./build.sh
```

4. アプリケーションのインストール
```bash
cp -r SystemMonitor.app /Applications/
```

### 直接インストール

1. [Releases](https://github.com/[your-username]/mac_act_monitor/releases)ページから最新のバージョンをダウンロード
2. `SystemMonitor.app`を`/Applications`フォルダにドラッグ＆ドロップ

## 使用方法

1. Finderから`SystemMonitor.app`を実行
2. メニューバーにCPUとメモリの使用率が表示されます
3. アプリケーションを終了するには、メニューバーのアイコンを右クリックして終了を選択

### システム起動時に自動起動する方法

1. システム環境設定を開く
2. 「ユーザーとグループ」を選択
3. 「ログイン項目」タブを選択
4. 「+」ボタンをクリック
5. `SystemMonitor.app`を選択して追加

## 開発

### 使用している主なパッケージ

- [github.com/getlantern/systray](https://github.com/getlantern/systray) - メニューバーアプリケーション用フレームワーク
- [github.com/shirou/gopsutil](https://github.com/shirou/gopsutil) - システムリソース情報取得用ライブラリ

### プロジェクト構造

```
.
├── main.go              # メインアプリケーションコード
├── build.sh            # ビルドスクリプト
├── Info.plist          # macOSアプリケーション設定ファイル
├── go.mod              # Goモジュール定義
├── go.sum              # Goモジュールチェックサム
└── README.md           # このファイル
```

## ライセンス

MIT License - 詳細は[LICENSE](LICENSE)ファイルを参照してください。

## 貢献

1. このリポジトリをフォーク
2. 新しいブランチを作成 (`git checkout -b feature/amazing-feature`)
3. 変更をコミット (`git commit -m 'Add some amazing feature'`)
4. ブランチをプッシュ (`git push origin feature/amazing-feature`)
5. プルリクエストを作成

## 作者

[Your Name]

## 謝辞

このプロジェクトは以下のオープンソースプロジェクトを使用しています：

- [systray](https://github.com/getlantern/systray)
- [gopsutil](https://github.com/shirou/gopsutil) 