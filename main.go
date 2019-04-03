package appointment_recurrence

/*
func SetExecutionTimes() {
	log.Println("SetExecution Times")
	var broadcastJobTasks []BroadcastJobTask

	this.DB.Find(&broadcastJobTasks,
		"active = ? AND "+
			"(next_executing_time = ? OR next_executing_time is null) AND "+
			"(end_date >= ? OR no_end_date = ?)",
		true,
		"0001-01-01",
		time.Now().Format("2006-01-02 15:04:00"),
		true)

	for _, job := range broadcastJobTasks {
		job.Calc()
		this.DB.Save(&job)

	}
}


