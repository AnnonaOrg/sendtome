package about

import (
	"github.com/umfaka/sendtome/internal/constvar"
	"github.com/umfaka/sendtome/internal/features"

	tele "gopkg.in/telebot.v3"
)

func init() {
	features.RegisterFeature("/about", OnProcess)
	features.RegisterFeature("/version", OnVersion)
}

func OnProcess(c tele.Context) error {
	text := constvar.APPAbout()
	return c.Reply(text)
}
func OnVersion(c tele.Context) error {
	text := constvar.APPAbout()
	return c.Reply(text)
}
