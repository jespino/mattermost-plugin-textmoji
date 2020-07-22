package main

import (
	"fmt"
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

const textmojiCommand = "textmoji"

func createTextmojiCommand() *model.Command {
	return &model.Command{
		Trigger:          textmojiCommand,
		AutoComplete:     true,
		AutoCompleteDesc: "Draw a textmoji",
		AutoCompleteHint: "[textmoji]",
		AutocompleteData: getAutocompleteData(),
	}
}

var textmojis = map[string]string{
	"4chan_emoticon":             "( ͡° ͜ʖ ͡°)",
	"angry_birds":                "( ఠൠఠ )ﾉ",
	"angry_face":                 "(╬ ಠ益ಠ)",
	"angry_troll":                "ヽ༼ ಠ益ಠ ༽ﾉ",
	"at_what_cost":               "ლ(ಠ益ಠლ)",
	"barf":                       "(´ж｀ς)",
	"basking_in_glory":           "ヽ(´ー｀)ノ",
	"boxing":                     "ლ(•́•́ლ)",
	"breakdown":                  "ಥ﹏ಥ",
	"careless":                   "◔_◔",
	"cheers":                     "（ ^_^）o自自o（^_^ ）",
	"chicken":                    "ʚ(•｀",
	"confused":                   "¿ⓧ_ⓧﮌ",
	"confused_scratch":           "(⊙.☉)7",
	"crazy":                      "ミ●﹏☉ミ",
	"creeper":                    "ƪ(ړײ)‎ƪ​​",
	"cry_face":                   "｡ﾟ( ﾟஇ‸இﾟ)ﾟ｡",
	"crying_face":                "ಥ_ಥ",
	"cry_troll":                  "༼ ༎ຶ ෴ ༎ຶ༽",
	"cute_bear":                  "ʕ•ᴥ•ʔ",
	"cute_face_with_big_eyes":    "(｡◕‿◕｡)",
	"dab":                        "ヽ( •_)ᕗ",
	"dance":                      "♪♪ ヽ(ˇ∀ˇ )ゞ",
	"dancing":                    "┌(ㆆ㉨ㆆ)ʃ",
	"dear_god_why":               "щ（ﾟДﾟщ）",
	"devious_smile":              "ಠ‿ಠ",
	"disagree":                   "٩◔̯◔۶",
	"discombobulated":            "⊙﹏⊙",
	"dislike":                    "( ಠ ʖ̯ ಠ)",
	"double_flip":                "┻━┻ ︵ヽ(`Д´)ﾉ︵ ┻━┻",
	"do_you_even_lift_bro?":      " ᕦ(ò_óˇ)ᕤ",
	"emo_dance":                  "ヾ(-_- )ゞ",
	"excited":                    "☜(⌒▽⌒)☞",
	"exorcism":                   "ح(•̀ж•́)ง †",
	"eye_roll":                   "⥀.⥀",
	"feel_perky":                 "(`･ω･´)",
	"fido":                       "V•ᴥ•V",
	"fight":                      "(ง'̀-'́)ง",
	"fisticuffs":                 "ლ(｀ー´ლ)",
	"flexing":                    "ᕙ(⇀‸↼‶)ᕗ",
	"flip_friend":                "(ノಠ ∩ಠ)ノ彡( \\o°o)\\",
	"fly_away":                   "⁽⁽ଘ( ˊᵕˋ )ଓ⁾⁾",
	"flying":                     "ح˚௰˚づ",
	"fuck_it":                    "t(-_-t)",
	"fuck_off":                   "(° ͜ʖ͡°)╭∩╮",
	"gtfo_bear":                  "ʕ •`ᴥ•´ʔ",
	"happy_face":                 "ヽ(´▽`)/",
	"happy_hug":                  "\\(ᵔᵕᵔ)/",
	"hitchhiking":                "(งツ)ว",
	"hugger":                     "(づ￣ ³￣)づ",
	"im_a_hugger":                "(⊃｡•́‿•̀｡)⊃",
	"injured":                    "(҂◡_◡)",
	"innocent_face":              "ʘ‿ʘ",
	"japanese_lion_face":         "°‿‿°",
	"judgemental":                "{ಠʖಠ}",
	"judging":                    "( ఠ ͟ʖ ఠ)",
	"kirby":                      "⊂(◉‿◉)つ",
	"kissing":                    "( ˘ ³˘)♥",
	"kitty_emote":                "ᵒᴥᵒ#",
	"listening_to_headphones":    "◖ᵔᴥᵔ◗ ♪ ♫",
	"looking_down":               "(._.)",
	"love":                       "♥‿♥",
	"meh":                        "¯\\(°_o)/¯",
	"meow":                       "ฅ^•ﻌ•^ฅ",
	"no_support":                 "乁( ◔ ౪◔)「      ┑(￣Д ￣)┍",
	"opera":                      "ヾ(´〇`)ﾉ♪♪♪",
	"peepers":                    "ಠಠ",
	"pointing":                   "(☞ﾟヮﾟ)☞",
	"pretty_eyes":                "ఠ_ఠ",
	"put_the_table_back":         "┬─┬ ノ( ゜-゜ノ)",
	"questionable":               "(Ծ‸ Ծ)",
	"reddit_disapproval_face":    "ಠ_ಠ",
	"resting_my_eyes":            "ᴖ̮ ̮ᴖ",
	"robot":                      "{•̃_•̃}",
	"running":                    "ε=ε=ε=┌(;*´Д`)ﾉ",
	"sad_and_confused":           "¯\\_(⊙︿⊙)_/¯",
	"sad_and_crying":             "(ᵟຶ︵ ᵟຶ)",
	"sad_face":                   "(ಥ⌣ಥ)",
	"satisfied":                  "(◠﹏◠)",
	"seal":                       "(ᵔᴥᵔ)",
	"shark_face":                 "( ˇ෴ˇ )",
	"shrug_face":                 "¯\\_(ツ)_/¯",
	"shy":                        "(๑•́ ₃ •̀๑)",
	"sleepy":                     "눈_눈",
	"smiley_toast":               "ʕʘ̅͜ʘ̅ʔ",
	"squinting_bear":             "ʕᵔᴥᵔʔ",
	"staring":                    "٩(๏_๏)۶",
	"stranger_danger":            "(づ｡◕‿‿◕｡)づ",
	"strut":                      "ᕕ( ᐛ )ᕗ",
	"stunna_shades":              "(っ▀¯▀)つ",
	"surprised":                  "（　ﾟДﾟ）",
	"table_flip":                 "(╯°□°）╯︵ ┻━┻",
	"taking_a_dump":              "(⩾﹏⩽)",
	"tgif":                       "“ヽ(´▽｀)ノ”",
	"things_that_cant_be_unseen": "♨_♨",
	"tidy_up":                    "┬─┬⃰͡ (ᵔᵕᵔ͜ )",
	"tired":                      "( ͡ಠ ʖ̯ ͡ಠ)",
	"touchy_feely":               "ԅ(≖‿≖ԅ)",
	"tripping_out":               "q(❂‿❂)p",
	"trolling":                   "༼∵༽ ༼⍨༽ ༼⍢༽ ༼⍤༽",
	"wave_dance":                 "~(^-^)~",
	"whistling":                  "(っ•́｡•́)♪♬",
	"winnie_the_pooh":            "ʕ •́؈•̀)",
	"winning":                    "(•̀ᴗ•́)و ̑̑",
	"wizard":                     "(∩｀-´)⊃━☆ﾟ.*･｡ﾟ",
	"worried":                    "(´･_･`)",
	"yum":                        "(っ˘ڡ˘ς)",
	"zombie":                     "[¬º-°]¬",
	"zoned":                      "(⊙_◎)",
}

func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	split := strings.Fields(args.Command)
	command := split[0]
	action := ""
	if len(split) > 1 {
		action = split[1]
	}

	if command != "/"+textmojiCommand {
		return &model.CommandResponse{}, nil
	}

	if textmoji, ok := textmojis[action]; ok {
		p.API.CreatePost(&model.Post{
			Message:   textmoji + strings.TrimPrefix(args.Command, command+" "+action),
			UserId:    args.UserId,
			ChannelId: args.ChannelId,
		})
	}
	return &model.CommandResponse{}, nil
}

func getAutocompleteData() *model.AutocompleteData {
	textmoji := model.NewAutocompleteData("textmoji", "[textmoji-name] [extra-text]", "Draw a text based emoji")

	for key, value := range textmojis {
		if strings.ToLower(key) != key {
			fmt.Printf("\"%s\" != \"%s\"\n", strings.ToLower(key), key)
		}
		autocomp := model.NewAutocompleteData(key, "[extra-text]", value)
		textmoji.AddCommand(autocomp)
	}
	return textmoji
}
