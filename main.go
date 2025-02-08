package main

import (
	"fmt"
	"time"

	"github.com/getlantern/systray"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTitle("🖥️")
	updateStats()

	// 定期的な更新
	go func() {
		ticker := time.NewTicker(2 * time.Second)
		for {
			select {
			case <-ticker.C:
				updateStats()
			}
		}
	}()
}

func onExit() {
	// クリーンアップ処理
}

func updateStats() {
	// CPU使用率の取得
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		fmt.Println("CPU情報の取得に失敗:", err)
		return
	}

	// メモリ使用率の取得
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("メモリ情報の取得に失敗:", err)
		return
	}

	// 表示を更新
	statsText := fmt.Sprintf("CPU: %.1f%% | MEM: %.1f%%",
		cpuPercent[0],
		memInfo.UsedPercent,
	)
	systray.SetTitle(statsText)
}
