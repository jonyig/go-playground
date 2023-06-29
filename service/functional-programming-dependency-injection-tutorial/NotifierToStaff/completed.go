package NotifierToStaff

import "fmt"

type Staff struct {
	Email string
	Phone string
	Line  string
}
type Notifier func(message string, contacts []string) error
type GetContacts func(eids []Staff) []string

func EmailNotifier(message string, contacts []string) error {
	// send email to owner
	return nil
}

func GetEmailContacts(eids []Staff) (emails []string) {
	for _, eid := range eids {
		emails = append(emails, eid.Email)
	}
	return
}

func getStaffs(eids []string) (staffs []Staff) {
	// get staffs by eids
	return
}
func SendAlertToStaffs(
	eids []string,
	message string,
	notifier func(message string, contacts []string) error,
	GetContacts func(eids []Staff) []string,
) (err error) {
	staffs := getStaffs(eids)
	contacts := GetContacts(staffs)
	err = notifier(message, contacts)
	return
}

func main() {
	// 創建依賴組件函數
	emailNotifier := EmailNotifier
	eids := []string{"E001", "E002", "E003"}
	// 使用依賴注入函數創建具有注入依賴的通知函數

	// 使用注入依賴的通知函數進行通知操作
	err := SendAlertToStaffs(eids, "Hello, World!", emailNotifier, GetEmailContacts)
	if err != nil {
		fmt.Println("通知發送失敗:", err)
	} else {
		fmt.Println("通知發送成功")
	}
}
