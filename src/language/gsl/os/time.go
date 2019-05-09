package os

import "time"

func Now() time.Time { return time.Now() }
func NowDay() int { return time.Now().Day() }
func NowMinute() int { return time.Now().Minute() }
func NowYear() int { return time.Now().Year() }
func NowFormat(str string) string { return time.Now().Format(str) }
func Sleep(d time.Duration) { time.Sleep(d) }

