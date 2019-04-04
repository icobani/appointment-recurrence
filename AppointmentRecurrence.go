/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   Kemal CAN SİLER
   4.04.2019
   09:39
*/
package appointment_recurrence

type AppointmentRecurrence struct {
	ETimeModel *ExecutionTimeModel
}

type ErrorStruct struct {
	Error string
}

func (this AppointmentRecurrence) New(etm *ExecutionTimeModel) (*AppointmentRecurrence, *ErrorStruct) {
	return &AppointmentRecurrence{
		ETimeModel: etm,
	}, nil
}
