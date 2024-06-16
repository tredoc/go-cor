package notification_test

import (
	"github.com/tredoc/go-cor/notification"
	"github.com/tredoc/go-cor/utils"
	"testing"
)

func TestNotification_Send(test *testing.T) {
	t, err := utils.GetTransaction(100)
	if err != nil {
		test.Fatal(err)
	}

	n, err := notification.NewNotification(t)
	if err != nil {
		test.Fatal(err)
	}

	err = n.Send()
	if err != nil {
		test.Fatal(err)
	}
}
