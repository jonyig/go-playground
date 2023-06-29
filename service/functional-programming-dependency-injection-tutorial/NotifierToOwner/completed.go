package NotifierToOwner

import "fmt"

type Notifier func(message string) error

func EmailNotifier(message string) error {
	// send email to owner
	return nil
}

func SendAlertToOwner(message string, Notifier func(message string) error) (err error) {
	message, err = translateMessageByOwnerLang(message)
	if err != nil {
		fmt.Println("翻譯失敗:", err)
		return
	}
	err = Notifier(message)
	return
}

func translateMessageByOwnerLang(message string) (string, error) {
	// do translate
	return message, nil
}
func main() {
	// 創建依賴組件函數
	emailNotifier := EmailNotifier

	// 使用依賴注入函數創建具有注入依賴的通知函數

	// 使用注入依賴的通知函數進行通知操作
	err := SendAlertToOwner("Hello, World!", emailNotifier)
	if err != nil {
		fmt.Println("通知發送失敗:", err)
	} else {
		fmt.Println("通知發送成功")
	}
}
