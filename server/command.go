package main

import (
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

type moji struct {
	words []string
	ascii string
}

var textmojis = []moji{
	moji{words: []string{"acid"}, ascii: "⊂(◉‿◉)つ"},
	moji{words: []string{"afraid"}, ascii: "(ㆆ _ ㆆ)"},
	moji{words: []string{"alpha"}, ascii: "α"},
	moji{words: []string{"angel"}, ascii: "☜(⌒▽⌒)☞"},
	moji{words: []string{"angry"}, ascii: "•`_´•"},
	moji{words: []string{"arrowhead"}, ascii: "⤜(ⱺ ʖ̯ⱺ)⤏"},
	moji{words: []string{"apple"}, ascii: ""},
	moji{words: []string{"ass", "butt"}, ascii: "(‿|‿)"},
	moji{words: []string{"awkward"}, ascii: "•͡˘㇁•͡˘"},
	moji{words: []string{"bat"}, ascii: "/|\\ ^._.^ /|\\"},
	moji{words: []string{"bear", "koala"}, ascii: "ʕ·͡ᴥ·ʔ"},
	moji{words: []string{"bearflip"}, ascii: "ʕノ•ᴥ•ʔノ ︵ ┻━┻"},
	moji{words: []string{"bearhug"}, ascii: "ʕっ•ᴥ•ʔっ"},
	moji{words: []string{"because", "since"}, ascii: "∵"},
	moji{words: []string{"beta"}, ascii: "β"},
	moji{words: []string{"bigheart"}, ascii: "❤"},
	moji{words: []string{"blackeye"}, ascii: "0__#"},
	moji{words: []string{"blubby"}, ascii: "(      0    _   0    )"},
	moji{words: []string{"blush"}, ascii: "(˵ ͡° ͜ʖ ͡°˵)"},
	moji{words: []string{"bond", "007"}, ascii: "┌( ͝° ͜ʖ͡°)=ε/̵͇̿̿/’̿’̿ ̿"},
	moji{words: []string{"boobs"}, ascii: "( . Y . )"},
	moji{words: []string{"bored"}, ascii: "(-_-)"},
	moji{words: []string{"bribe"}, ascii: "( •͡˘ _•͡˘)ノð"},
	moji{words: []string{"bubbles"}, ascii: "( ˘ ³˘)ノ°ﾟº❍｡"},
	moji{words: []string{"butterfly"}, ascii: "ƸӜƷ"},
	moji{words: []string{"cat"}, ascii: "(= ФェФ=)"},
	moji{words: []string{"catlenny"}, ascii: "( ͡° ᴥ ͡°)"},
	moji{words: []string{"check"}, ascii: "✔"},
	moji{words: []string{"cheer"}, ascii: "※\\(^o^)/※"},
	moji{words: []string{"chubby"}, ascii: "╭(ʘ̆~◞౪◟~ʘ̆)╮"},
	moji{words: []string{"claro"}, ascii: "(͡ ° ͜ʖ ͡ °)"},
	moji{words: []string{"clique", "gang", "squad"}, ascii: "ヽ༼ ຈل͜ຈ༼ ▀̿̿Ĺ̯̿̿▀̿ ̿༽Ɵ͆ل͜Ɵ͆ ༽ﾉ"},
	moji{words: []string{"cloud"}, ascii: "☁"},
	moji{words: []string{"club"}, ascii: "♣"},
	moji{words: []string{"coffee", "cuppa"}, ascii: "c[_]"},
	moji{words: []string{"cmd", "command"}, ascii: "⌘"},
	moji{words: []string{"cool", "csi"}, ascii: "(•_•) ( •_•)>⌐■-■ (⌐■_■)"},
	moji{words: []string{"copy", "c"}, ascii: "©"},
	moji{words: []string{"creep"}, ascii: "ԅ(≖‿≖ԅ)"},
	moji{words: []string{"creepcute"}, ascii: "ƪ(ړײ)‎ƪ​​"},
	moji{words: []string{"crim3s"}, ascii: "( ✜︵✜ )"},
	moji{words: []string{"cross"}, ascii: "†"},
	moji{words: []string{"cry"}, ascii: "(╥﹏╥)"},
	moji{words: []string{"crywave"}, ascii: "( ╥﹏╥) ノシ"},
	moji{words: []string{"cute"}, ascii: "(｡◕‿‿◕｡)"},
	moji{words: []string{"d1"}, ascii: "⚀"},
	moji{words: []string{"d2"}, ascii: "⚁"},
	moji{words: []string{"d3"}, ascii: "⚂"},
	moji{words: []string{"d4"}, ascii: "⚃"},
	moji{words: []string{"d5"}, ascii: "⚄"},
	moji{words: []string{"d6"}, ascii: "⚅"},
	moji{words: []string{"dab"}, ascii: "ヽ( •_)ᕗ"},
	moji{words: []string{"damnyou"}, ascii: "(ᕗ ͠° ਊ ͠° )ᕗ"},
	moji{words: []string{"dance"}, ascii: "ᕕ(⌐■_■)ᕗ ♪♬"},
	moji{words: []string{"dead"}, ascii: "x⸑x"},
	moji{words: []string{"dealwithit", "dwi"}, ascii: "(⌐■_■)"},
	moji{words: []string{"delta"}, ascii: "Δ"},
	moji{words: []string{"depressed"}, ascii: "(︶︹︶)"},
	moji{words: []string{"derp"}, ascii: "☉ ‿ ⚆"},
	moji{words: []string{"diamond"}, ascii: "♦"},
	moji{words: []string{"dj"}, ascii: "d[-_-]b"},
	moji{words: []string{"dog"}, ascii: "(◕ᴥ◕ʋ)"},
	moji{words: []string{"dong"}, ascii: "(̿▀̿ ̿Ĺ̯̿̿▀̿ ̿)̄"},
	moji{words: []string{"donger"}, ascii: "ヽ༼ຈل͜ຈ༽ﾉ"},
	moji{words: []string{"dontcare", "idc"}, ascii: "(- ʖ̯-)"},
	moji{words: []string{"do not want", "dontwant"}, ascii: "ヽ(｀Д´)ﾉ"},
	moji{words: []string{"dope"}, ascii: "<(^_^)>"},
	moji{words: []string{"<<"}, ascii: "«"},
	moji{words: []string{">>"}, ascii: "»"},
	moji{words: []string{"doubleflat"}, ascii: "𝄫"},
	moji{words: []string{"doublesharp"}, ascii: "𝄪"},
	moji{words: []string{"doubletableflip"}, ascii: "┻━┻ ︵ヽ(`Д´)ﾉ︵ ┻━┻"},
	moji{words: []string{"down"}, ascii: "↓"},
	moji{words: []string{"duckface"}, ascii: "(・3・)"},
	moji{words: []string{"duel"}, ascii: "ᕕ(╭ರ╭ ͟ʖ╮•́)⊃¤=(————-"},
	moji{words: []string{"duh"}, ascii: "(≧︿≦)"},
	moji{words: []string{"dunno"}, ascii: "¯\\(°_o)/¯"},
	moji{words: []string{"ebola"}, ascii: "ᴇʙᴏʟᴀ"},
	moji{words: []string{"eeriemob"}, ascii: "(-(-_-(-_(-_(-_-)_-)-_-)_-)_-)-)"},
	moji{words: []string{"ellipsis", "..."}, ascii: "…"},
	moji{words: []string{"emdash", "--"}, ascii: "–"},
	moji{words: []string{"emptystar"}, ascii: "☆"},
	moji{words: []string{"emptytriangle", "t2"}, ascii: "△"},
	moji{words: []string{"endure"}, ascii: "(҂◡_◡) ᕤ"},
	moji{words: []string{"envelope", "letter"}, ascii: "✉︎"},
	moji{words: []string{"epsilon"}, ascii: "ɛ"},
	moji{words: []string{"euro"}, ascii: "€"},
	moji{words: []string{"evil"}, ascii: "ψ(｀∇´)ψ"},
	moji{words: []string{"evillenny"}, ascii: "(͠≖ ͜ʖ͠≖)"},
	moji{words: []string{"excited"}, ascii: "(ﾉ◕ヮ◕)ﾉ*:・ﾟ✧"},
	moji{words: []string{"execution"}, ascii: "(⌐■_■)︻╦╤─   (╥﹏╥)"},
	moji{words: []string{"facebook"}, ascii: "(╯°□°)╯︵ ʞooqǝɔɐɟ"},
	moji{words: []string{"facepalm"}, ascii: "(－‸ლ)"},
	moji{words: []string{"fart"}, ascii: "(ˆ⺫ˆ๑)<3"},
	moji{words: []string{"fight"}, ascii: "(ง •̀_•́)ง"},
	moji{words: []string{"finn"}, ascii: "| (• ◡•)|"},
	moji{words: []string{"fish"}, ascii: "<\"(((<3"},
	moji{words: []string{"5", "five"}, ascii: "卌"},
	moji{words: []string{"5/8"}, ascii: "⅝"},
	moji{words: []string{"flat", "bemolle"}, ascii: "♭"},
	moji{words: []string{"flexing"}, ascii: "ᕙ(`▽´)ᕗ"},
	moji{words: []string{"flipped", "heavytable"}, ascii: "┬─┬ ︵ /(.□. \\）"},
	moji{words: []string{"flower", "flor"}, ascii: "(✿◠‿◠)"},
	moji{words: []string{"f"}, ascii: "✿"},
	moji{words: []string{"fly"}, ascii: "─=≡Σ((( つ◕ل͜◕)つ"},
	moji{words: []string{"friendflip"}, ascii: "(╯°□°)╯︵ ┻━┻ ︵ ╯(°□° ╯)"},
	moji{words: []string{"frown"}, ascii: "(ღ˘⌣˘ღ)"},
	moji{words: []string{"fuckoff", "gtfo"}, ascii: "୧༼ಠ益ಠ╭∩╮༽"},
	moji{words: []string{"fuckyou", "fu"}, ascii: "┌П┐(ಠ_ಠ)"},
	moji{words: []string{"gentleman", "sir", "monocle"}, ascii: "ಠ_ರೃ"},
	moji{words: []string{"ghast"}, ascii: "= _ ="},
	moji{words: []string{"ghost"}, ascii: "༼ つ ╹ ╹ ༽つ"},
	moji{words: []string{"gift", "present"}, ascii: "(´・ω・)っ由"},
	moji{words: []string{"gimme"}, ascii: "༼ つ ◕_◕ ༽つ"},
	moji{words: []string{"givemeyourmoney"}, ascii: "(•-•)⌐"},
	moji{words: []string{"glitter"}, ascii: "(*・‿・)ノ⌒*:･ﾟ✧"},
	moji{words: []string{"glasses"}, ascii: "(⌐ ͡■ ͜ʖ ͡■)"},
	moji{words: []string{"glassesoff"}, ascii: "( ͡° ͜ʖ ͡°)ﾉ⌐■-■"},
	moji{words: []string{"glitterderp"}, ascii: "(ﾉ☉ヮ⚆)ﾉ ⌒*:･ﾟ✧"},
	moji{words: []string{"gloomy"}, ascii: "(_゜_゜_)"},
	moji{words: []string{"goatse"}, ascii: "(з๏ε)"},
	moji{words: []string{"gotit"}, ascii: "(☞ﾟ∀ﾟ)☞"},
	moji{words: []string{"greet", "greetings"}, ascii: "( ´◔ ω◔`) ノシ"},
	moji{words: []string{"gun", "mg"}, ascii: "︻╦╤─"},
	moji{words: []string{"hadouken"}, ascii: "༼つಠ益ಠ༽つ ─=≡ΣO))"},
	moji{words: []string{"hammerandsickle", "hs"}, ascii: "☭"},
	moji{words: []string{"handleft", "hl"}, ascii: "☜"},
	moji{words: []string{"handright", "hr"}, ascii: "☞"},
	moji{words: []string{"haha"}, ascii: "٩(^‿^)۶"},
	moji{words: []string{"happy"}, ascii: "٩( ๑╹ ꇴ╹)۶"},
	moji{words: []string{"happygarry"}, ascii: "ᕕ( ᐛ )ᕗ"},
	moji{words: []string{"h", "heart"}, ascii: "♥"},
	moji{words: []string{"hello", "ohai", "bye"}, ascii: "(ʘ‿ʘ)╯"},
	moji{words: []string{"help"}, ascii: "\\(°Ω°)/"},
	moji{words: []string{"highfive"}, ascii: "._.)/\\(._."},
	moji{words: []string{"hitting"}, ascii: "( ｀皿´)｡ﾐ/"},
	moji{words: []string{"hug", "hugs"}, ascii: "(づ｡◕‿‿◕｡)づ"},
	moji{words: []string{"iknowright", "ikr"}, ascii: "┐｜･ิω･ิ#｜┌"},
	moji{words: []string{"illuminati"}, ascii: "୧(▲ᴗ▲)ノ"},
	moji{words: []string{"infinity", "inf"}, ascii: "∞"},
	moji{words: []string{"inlove"}, ascii: "(っ´ω`c)♡"},
	moji{words: []string{"int"}, ascii: "∫"},
	moji{words: []string{"internet"}, ascii: "ଘ(੭*ˊᵕˋ)੭* ̀ˋ ɪɴᴛᴇʀɴᴇᴛ"},
	moji{words: []string{"interrobang"}, ascii: "‽"},
	moji{words: []string{"jake"}, ascii: "(❍ᴥ❍ʋ)"},
	moji{words: []string{"kappa"}, ascii: "(¬,‿,¬)"},
	moji{words: []string{"kawaii"}, ascii: "≧◡≦"},
	moji{words: []string{"keen"}, ascii: "┬┴┬┴┤Ɵ͆ل͜Ɵ͆ ༽ﾉ"},
	moji{words: []string{"kiahh"}, ascii: "~\\(≧▽≦)/~"},
	moji{words: []string{"kiss"}, ascii: "(づ ￣ ³￣)づ"},
	moji{words: []string{"kyubey"}, ascii: "／人◕ ‿‿ ◕人＼"},
	moji{words: []string{"lambda"}, ascii: "λ"},
	moji{words: []string{"lazy"}, ascii: "_(:3」∠)_"},
	moji{words: []string{"left", "<-"}, ascii: "←"},
	moji{words: []string{"lenny"}, ascii: "( ͡° ͜ʖ ͡°)"},
	moji{words: []string{"lennybill"}, ascii: "[̲̅$̲̅(̲̅ ͡° ͜ʖ ͡°̲̅)̲̅$̲̅]"},
	moji{words: []string{"lennyfight"}, ascii: "(ง ͠° ͟ʖ ͡°)ง"},
	moji{words: []string{"lennyflip"}, ascii: "(ノ ͡° ͜ʖ ͡°ノ)   ︵ ( ͜。 ͡ʖ ͜。)"},
	moji{words: []string{"lennygang"}, ascii: "( ͡°( ͡° ͜ʖ( ͡° ͜ʖ ͡°)ʖ ͡°) ͡°)"},
	moji{words: []string{"lennyshrug"}, ascii: "¯\\_( ͡° ͜ʖ ͡°)_/¯"},
	moji{words: []string{"lennysir"}, ascii: "( ಠ ͜ʖ ರೃ)"},
	moji{words: []string{"lennystalker"}, ascii: "┬┴┬┴┤( ͡° ͜ʖ├┬┴┬┴"},
	moji{words: []string{"lennystrong"}, ascii: "ᕦ( ͡° ͜ʖ ͡°)ᕤ"},
	moji{words: []string{"lennywizard"}, ascii: "╰( ͡° ͜ʖ ͡° )つ──☆*:・ﾟ"},
	moji{words: []string{"lol"}, ascii: "L(° O °L)"},
	moji{words: []string{"look"}, ascii: "(ಡ_ಡ)☞"},
	moji{words: []string{"loud", "noise"}, ascii: "ᕦ(⩾﹏⩽)ᕥ"},
	moji{words: []string{"love"}, ascii: "♥‿♥"},
	moji{words: []string{"lovebear"}, ascii: "ʕ♥ᴥ♥ʔ"},
	moji{words: []string{"lumpy"}, ascii: "꒰ ꒡⌓꒡꒱"},
	moji{words: []string{"luv"}, ascii: "-`ღ´-"},
	moji{words: []string{"magic"}, ascii: "ヽ(｀Д´)⊃━☆ﾟ. * ･ ｡ﾟ,"},
	moji{words: []string{"magicflip"}, ascii: "(/¯◡ ‿ ◡)/¯ ~ ┻━┻"},
	moji{words: []string{"meep"}, ascii: "\\(°^°)/"},
	moji{words: []string{"meh"}, ascii: "ಠ_ಠ"},
	moji{words: []string{"metal", "rock"}, ascii: "\\m/,(> . <)_\\m/"},
	moji{words: []string{"mistyeyes"}, ascii: "ಡ_ಡ"},
	moji{words: []string{"monster"}, ascii: "༼ ༎ຶ ෴ ༎ຶ༽"},
	moji{words: []string{"natural"}, ascii: "♮"},
	moji{words: []string{"needle", "inject"}, ascii: "┌(◉ ͜ʖ◉)つ┣▇▇▇═──"},
	moji{words: []string{"nerd"}, ascii: "(⌐⊙_⊙)"},
	moji{words: []string{"nice"}, ascii: "( ͡° ͜ °)"},
	moji{words: []string{"no"}, ascii: "→_←"},
	moji{words: []string{"noclue"}, ascii: "／人◕ __ ◕人＼"},
	moji{words: []string{"nom", "yummy", "delicious"}, ascii: "(っˆڡˆς)"},
	moji{words: []string{"note", "sing"}, ascii: "♫"},
	moji{words: []string{"nuclear", "radioactive", "nukular"}, ascii: "☢"},
	moji{words: []string{"nyan"}, ascii: "~=[,,_,,]:3"},
	moji{words: []string{"nyeh"}, ascii: "@^@"},
	moji{words: []string{"ohshit"}, ascii: "( º﹃º )"},
	moji{words: []string{"omega"}, ascii: "Ω"},
	moji{words: []string{"omg"}, ascii: "◕_◕"},
	moji{words: []string{"1/8"}, ascii: "⅛"},
	moji{words: []string{"1/4"}, ascii: "¼"},
	moji{words: []string{"1/2"}, ascii: "½"},
	moji{words: []string{"1/3"}, ascii: "⅓"},
	moji{words: []string{"opt", "option"}, ascii: "⌥"},
	moji{words: []string{"orly"}, ascii: "(눈_눈)"},
	moji{words: []string{"ohyou", "ou"}, ascii: "(◞థ౪థ)ᴖ"},
	moji{words: []string{"peace", "victory"}, ascii: "✌(-‿-)✌"},
	moji{words: []string{"pear"}, ascii: "(__>-"},
	moji{words: []string{"pi"}, ascii: "π"},
	moji{words: []string{"pingpong"}, ascii: "( •_•)O*¯`·.¸.·´¯`°Q(•_• )"},
	moji{words: []string{"plain"}, ascii: "._."},
	moji{words: []string{"pleased"}, ascii: "(˶‾᷄ ⁻̫ ‾᷅˵)"},
	moji{words: []string{"point"}, ascii: "(☞ﾟヮﾟ)☞"},
	moji{words: []string{"pooh"}, ascii: "ʕ •́؈•̀)"},
	moji{words: []string{"porcupine"}, ascii: "(•ᴥ• )́`́'́`́'́⻍"},
	moji{words: []string{"pound"}, ascii: "£"},
	moji{words: []string{"praise"}, ascii: "(☝ ՞ਊ ՞)☝"},
	moji{words: []string{"punch"}, ascii: "O=('-'Q)"},
	moji{words: []string{"rage", "mad"}, ascii: "t(ಠ益ಠt)"},
	moji{words: []string{"rageflip"}, ascii: "(ノಠ益ಠ)ノ彡┻━┻"},
	moji{words: []string{"rainbowcat"}, ascii: "(=^･ｪ･^=))ﾉ彡☆"},
	moji{words: []string{"really"}, ascii: "ò_ô"},
	moji{words: []string{"r"}, ascii: "®"},
	moji{words: []string{"right", "->"}, ascii: "→"},
	moji{words: []string{"riot"}, ascii: "୧༼ಠ益ಠ༽୨"},
	moji{words: []string{"rolleyes"}, ascii: "(◔_◔)"},
	moji{words: []string{"rose"}, ascii: "✿ڿڰۣ—"},
	moji{words: []string{"run"}, ascii: "(╯°□°)╯"},
	moji{words: []string{"sad"}, ascii: "ε(´סּ︵סּ`)з"},
	moji{words: []string{"saddonger"}, ascii: "ヽ༼ຈʖ̯ຈ༽ﾉ"},
	moji{words: []string{"sadlenny"}, ascii: "( ͡° ʖ̯ ͡°)"},
	moji{words: []string{"7/8"}, ascii: "⅞"},
	moji{words: []string{"sharp", "diesis"}, ascii: "♯"},
	moji{words: []string{"shout"}, ascii: "╚(•⌂•)╝"},
	moji{words: []string{"shrug"}, ascii: "¯\\_(ツ)_/¯"},
	moji{words: []string{"shy"}, ascii: "=^_^="},
	moji{words: []string{"sigma", "sum"}, ascii: "Σ"},
	moji{words: []string{"skull"}, ascii: "☠"},
	moji{words: []string{"smile"}, ascii: "ツ"},
	moji{words: []string{"smiley"}, ascii: "☺︎"},
	moji{words: []string{"smirk"}, ascii: "¬‿¬"},
	moji{words: []string{"snowman"}, ascii: "☃"},
	moji{words: []string{"sob"}, ascii: "(;´༎ຶД༎ຶ`)"},
	moji{words: []string{"soviettableflip"}, ascii: "ノ┬─┬ノ ︵ ( \\o°o)\\"},
	moji{words: []string{"spade"}, ascii: "♠"},
	moji{words: []string{"sqrt"}, ascii: "√"},
	moji{words: []string{"squid"}, ascii: "<コ:彡"},
	moji{words: []string{"star"}, ascii: "★"},
	moji{words: []string{"strong"}, ascii: "ᕙ(⇀‸↼‶)ᕗ"},
	moji{words: []string{"suicide"}, ascii: "ε/̵͇̿̿/’̿’̿ ̿(◡︵◡)"},
	moji{words: []string{"sum"}, ascii: "∑"},
	moji{words: []string{"sun"}, ascii: "☀"},
	moji{words: []string{"surprised"}, ascii: "(๑•́ ヮ •̀๑)"},
	moji{words: []string{"surrender"}, ascii: "\\_(-_-)_/"},
	moji{words: []string{"stalker"}, ascii: "┬┴┬┴┤(･_├┬┴┬┴"},
	moji{words: []string{"swag"}, ascii: "(̿▀̿‿ ̿▀̿ ̿)"},
	moji{words: []string{"sword"}, ascii: "o()xxxx[{::::::::::::::::::>"},
	moji{words: []string{"tabledown"}, ascii: "┬─┬ ノ( ゜-゜ノ)"},
	moji{words: []string{"tableflip"}, ascii: "(ノ ゜Д゜)ノ ︵ ┻━┻"},
	moji{words: []string{"tau"}, ascii: "τ"},
	moji{words: []string{"tears"}, ascii: "(ಥ﹏ಥ)"},
	moji{words: []string{"terrorist"}, ascii: "୧༼ಠ益ಠ༽︻╦╤─"},
	moji{words: []string{"thanks", "thankyou", "ty"}, ascii: "\\(^-^)/"},
	moji{words: []string{"therefore", "so"}, ascii: "⸫"},
	moji{words: []string{"this"}, ascii: "( ͡° ͜ʖ ͡°)_/¯"},
	moji{words: []string{"3/8"}, ascii: "⅜"},
	moji{words: []string{"tiefighter"}, ascii: "|=-(¤)-=|"},
	moji{words: []string{"tired"}, ascii: "(=____=)"},
	moji{words: []string{"toldyouso", "toldyou"}, ascii: "☜(꒡⌓꒡)"},
	moji{words: []string{"toogood"}, ascii: "ᕦ(òᴥó)ᕥ"},
	moji{words: []string{"tm"}, ascii: "™"},
	moji{words: []string{"triangle", "t"}, ascii: "▲"},
	moji{words: []string{"2/3"}, ascii: "⅔"},
	moji{words: []string{"unflip"}, ascii: "┬──┬ ノ(ò_óノ)"},
	moji{words: []string{"up"}, ascii: "↑"},
	moji{words: []string{"victory"}, ascii: "(๑•̀ㅂ•́)ง✧"},
	moji{words: []string{"wat"}, ascii: "(ÒДÓױ)"},
	moji{words: []string{"wave"}, ascii: "( * ^ *) ノシ"},
	moji{words: []string{"whaa"}, ascii: "Ö"},
	moji{words: []string{"whistle"}, ascii: "(っ^з^)♪♬"},
	moji{words: []string{"whoa"}, ascii: "(°o•)"},
	moji{words: []string{"why"}, ascii: "ლ(`◉◞౪◟◉‵ლ)"},
	moji{words: []string{"woo"}, ascii: "＼(＾O＾)／"},
	moji{words: []string{"wtf"}, ascii: "(⊙＿⊙')"},
	moji{words: []string{"wut"}, ascii: "⊙ω⊙"},
	moji{words: []string{"yay"}, ascii: "\\( ﾟヮﾟ)/"},
	moji{words: []string{"yeah", "yes"}, ascii: "(•̀ᴗ•́)و ̑̑"},
	moji{words: []string{"yen"}, ascii: "¥"},
	moji{words: []string{"yinyang", "yy"}, ascii: "☯"},
	moji{words: []string{"yolo"}, ascii: "Yᵒᵘ Oᶰˡʸ Lᶤᵛᵉ Oᶰᶜᵉ"},
	moji{words: []string{"youkids", "ukids"}, ascii: "ლ༼>╭ ͟ʖ╮<༽ლ"},
	moji{words: []string{"yuno"}, ascii: "(屮ﾟДﾟ)屮 Y U NO"},
	moji{words: []string{"zen", "meditation", "omm"}, ascii: "⊹╰(⌣ʟ⌣)╯⊹"},
	moji{words: []string{"zoidberg"}, ascii: "(V) (°,,,,°) (V)"},
	moji{words: []string{"zombie"}, ascii: "[¬º-°]¬"},
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

	for _, textmoji := range textmojis {
		for _, word := range textmoji.words {
			if action == word {
				p.API.CreatePost(&model.Post{
					Message:   textmoji.ascii + strings.TrimPrefix(args.Command, command+" "+action),
					UserId:    args.UserId,
					ChannelId: args.ChannelId,
					ParentId:  args.ParentId,
					RootId:    args.RootId,
				})
				return &model.CommandResponse{}, nil
			}
		}
	}
	return &model.CommandResponse{}, nil
}

func getAutocompleteData() *model.AutocompleteData {
	textmoji := model.NewAutocompleteData("textmoji", "[textmoji-name] [extra-text]", "Draw a text based emoji")

	for _, value := range textmojis {
		for _, word := range value.words {
			autocomp := model.NewAutocompleteData(word, "[extra-text]", value.ascii)
			textmoji.AddCommand(autocomp)
		}
	}
	return textmoji
}
