package about

import (
	"github.com/dbidib/sendtome/features"
	"github.com/dbidib/sendtome/internal/constvar"

	tele "gopkg.in/telebot.v3"
)

func OnProcess(c tele.Context) error {
	text := constvar.About()
	return c.Reply(text)
}
func OnVersion(c tele.Context) error {
	text := constvar.Version()
	return c.Reply(text)
}

func init() {
	features.RegisterFeature("/about", OnProcess)
	features.RegisterFeature("/version", OnVersion)
}
