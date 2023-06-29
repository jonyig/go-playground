package NotifierToStaff

type StaffInfo struct {
	Email string
	Phone string
	Line  string
	Slack string
}

func getStaffList(eids []string) (staffs []StaffInfo) {
	// get staffs by eids
	return
}

func EmailNotify(message string, contacts []string) error {
	// send email to owner
	return nil
}

func SendAlertToStaffsWithEmail(
	eids []string,
	message string,
) (err error) {
	staffs := getStaffList(eids)

	contacts := make([]string, len(staffs))
	for i, staff := range staffs {
		contacts[i] = staff.Email
	}
	err = EmailNotify(message, contacts)
	return
}

func SMSNotify(message string, contacts []string) error {
	// send sms to owner
	return nil
}

func SendAlertToStaffsWithSMS(
	eids []string,
	message string,
) (err error) {
	staffs := getStaffList(eids)

	contacts := make([]string, len(staffs))
	for i, staff := range staffs {
		contacts[i] = staff.Phone
	}
	err = SMSNotify(message, contacts)
	return
}

func LineNotify(message string, contacts []string) error {
	// send Line to owner
	return nil
}

func SendAlertToStaffsWithLine(
	eids []string,
	message string,
) (err error) {
	staffs := getStaffList(eids)

	contacts := make([]string, len(staffs))
	for i, staff := range staffs {
		contacts[i] = staff.Line
	}
	err = SMSNotify(message, contacts)
	return
}

func SlackNotify(message string, contacts []string) error {
	// send Slack to owner
	return nil
}

func SendAlertToStaffsWithSlack(
	eids []string,
	message string,
) (err error) {
	staffs := getStaffList(eids)

	contacts := make([]string, len(staffs))
	for i, staff := range staffs {
		contacts[i] = staff.Slack
	}
	err = SlackNotify(message, contacts)
	return
}
