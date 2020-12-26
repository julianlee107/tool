package cmd

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/julianlee107/tool/internal/timer"
	"github.com/spf13/cobra"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var nowtimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {

		now := timer.GetNowTime()
		log.Printf("输出结果：%s,%d", now.Format(time.UnixDate), now.Unix())
	},
}

var calculateTime string
var duration string

var calculateTimeCmd = &cobra.Command{
	Use:   "cal",
	Short: "计算时间",
	Long:  "计算时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTime time.Time
		var layout = "2006-01-02 15:04:05"
		if calculateTime == "" {
			currentTime = timer.GetNowTime()
		} else {
			var err error
			space := strings.Count(calculateTime, " ")
			if space == 0 {
				layout = "2006-01-02"
			}
			if space == 1 {
				layout = "2006-01-02 15:04"
			}
			currentTime, err = time.Parse(layout, calculateTime)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTime = time.Unix(int64(t), 0)
			}
		}
		t, err := timer.GetCalculateTime(currentTime, duration)

		if err != nil {
			log.Fatalf("timer.GetCalculateTime err:%v", err)
		}
		log.Printf("输出结果:%s,%d", t.Format(layout), t.Unix())

	},
}

func init() {
	timeCmd.AddCommand(nowtimeCmd, calculateTimeCmd)

	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", "持续时间")
	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", "需要计算的时间，有效单位为时间戳或者格式化后的时间")

}
