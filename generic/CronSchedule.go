package generic

type CronSchedule struct {
	QuartzCronExpression string `json:"quartz_cron_expression,omitempty"`
	TimezoneID           string `json:"timezone_id,omitempty"`
}
