package NotifierToOwner

import "fmt"

func EmailNotify(message string) error {
	// send email to owner
	return nil
}

func SendAlertToOwnerEmail() {
	message := "Hello, World!"
	message, err := translateMessageByOwnerLang(message)
	if err != nil {
		fmt.Println("翻譯失敗:", err)
		return
	}
	err = EmailNotify(message)
	if err != nil {
		fmt.Println("通知發送失敗:", err)
	} else {
		fmt.Println("通知發送成功")
	}
}

func SMSNotify(message string) error {
	// send sms to owner
	return nil
}

func SendAlertToOwnerSMS() {
	message := "Hello, World!"
	message, err := translateMessageByOwnerLang(message)
	if err != nil {
		fmt.Println("翻譯失敗:", err)
		return
	}
	err = SMSNotify(message)
	if err != nil {
		fmt.Println("通知發送失敗:", err)
	} else {
		fmt.Println("通知發送成功")
	}
}

func BotNotify(message string) error {
	// send bot to owner
	return nil
}

func SendAlertToOwnerBot() {
	message := "Hello, World!"
	message, err := translateMessageByOwnerLang(message)
	if err != nil {
		fmt.Println("翻譯失敗:", err)
		return
	}
	err = BotNotify(message)
	if err != nil {
		fmt.Println("通知發送失敗:", err)
	} else {
		fmt.Println("通知發送成功")
	}
}

func SlackNotify(message string) error {
	// send slack to owner
	return nil
}

func SendAlertToOwnerSlack() {
	message := "Hello, World!"
	message, err := translateMessageByOwnerLang(message)
	if err != nil {
		fmt.Println("翻譯失敗:", err)
		return
	}
	err = SlackNotify(message)
	if err != nil {
		fmt.Println("通知發送失敗:", err)
	} else {
		fmt.Println("通知發送成功")
	}
}

func TeamsNotify(message string) error {
	// send teams to owner
	return nil
}

func SendAlertToOwnerTeams() {
	message := "Hello, World!"
	message, err := translateMessageByOwnerLang(message)
	if err != nil {
		fmt.Println("翻譯失敗:", err)
		return
	}
	err = TeamsNotify(message)
	if err != nil {
		fmt.Println("通知發送失敗:", err)
	} else {
		fmt.Println("通知發送成功")
	}
}
