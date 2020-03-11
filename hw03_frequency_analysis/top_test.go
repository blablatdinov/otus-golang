package hw03_frequency_analysis //nolint:golint

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/stretchr/testify/assert"
)

var russianText = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин. Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

var hindiText = `हमारी जाएन उदेशीत अधिकांश परिवहन आशाआपस प्राप्त बलवान लाभो किया यायेका बिन्दुओमे सीमित हमारी उद्योग प्रमान ऎसाजीस कारन स्थिति कार्यसिधान्तो मजबुत लेने होसके दर्शाता विचारशिलता विकेन्द्रित कीसे पुस्तक ज्यादा वेबजाल भाति मजबुत ७०है कुशलता आवश्यक निर्माता प्रति

आशाआपस अनुवाद विज्ञान लाभो बाजार प्राण अथवा गयेगया पहोचाना दस्तावेज संपादक भारतीय मुश्किले समजते बिना एसलिये प्राथमिक यन्त्रालय मुक्त सकते शीघ्र लगती मानसिक हमेहो। दिनांक कीसे मुख्य एसलिये पुष्टिकर्ता जागरुक प्राधिकरन नाकर पहोच। आवश्यक सभीकुछ विश्वास सकता प्रति

हीकम क्षमता विभाजन संस्थान सक्षम होभर चुनने सुविधा यायेका निर्देश विशेष सामूहिक कार्यलय एछित पढने सुस्पश्ट ढांचामात्रुभाषा समाजो ज्यादा सोफ़्टवेर करते समजते कर्य विभाग बेंगलूर समस्याए लक्षण होगा वातावरण निर्माता भेदनक्षमता सुविधा अनुवादक कम्प्युटर चिदंश जिसे प्रति पहोच बिन्दुओ आवश्यकत केन्द्रिय

अनुवादक स्वतंत्रता होभर चिदंश सहायता अनुवाद यायेका लेकिन पेदा पहोच। रहारुप बातसमय विकेन्द्रियकरण भारत लेने तरीके ज्यादा ढांचामात्रुभाषा संदेश द्वारा कम्प्युटर अमितकुमार सकता बनाकर भाति अनुवाद तकनिकल करता खरिदे एकत्रित बेंगलूर आवश्यक वास्तव प्रतिबध उदेश कराना सुचना होने निर्देश समस्याओ

मुख्य अंग्रेजी सार्वजनिक आपके भाति क्षमता सोफ़्टवेर मुख्य वास्तव प्रौध्योगिकी सहयोग जाने विकास हुएआदि निर्माण लाभान्वित सुनत नीचे मुख्य असक्षम सुचनाचलचित्र वर्णन हैं। हुएआदि यन्त्रालय मार्गदर्शन कराना गटकउसि अनुवाद विश्वास कारन सामूहिक संस्क्रुति मजबुत यधपि दर्शाता गोपनीयता भेदनक्षमता कारन बाटते ऎसाजीस दर्शाता अमितकुमार बनाकर मुखय गटकउसि हुआआदी मुख्यतह उन्हे दारी समाज और्४५० अंतर्गत बिन्दुओमे बनाति

आपके व्याख्यान करके(विशेष तरीके वर्णन मानव आवश्यक वर्णित स्थापित विशेष जिवन असरकारक स्वतंत्र अंग्रेजी लेने सिद्धांत स्वतंत्रता पहोच। लक्षण भारत जिम्मे प्रमान हमेहो। बहुत ढांचा वर्तमान अंग्रेजी एसेएवं जानकारी सक्षम उनको

जाता खयालात निर्देश माहितीवानीज्य बढाता स्थिति मार्गदर्शन बनाना आजपर पहोचने लाभान्वित उपयोगकर्ता भाषाओ जाता लक्षण शुरुआत अत्यंत सभिसमज २४भि होभर लक्ष्य गटकउसि पहोच। लेकिन लचकनहि अधिकांश केन्द्रिय मानव पहेला जैसी असरकारक विभाजनक्षमता वास्तव पुष्टिकर्ता विभाजनक्षमता विश्व समूह

विवरण प्रमान विश्वास बिन्दुओमे वहहर विषय बनाकर अत्यंत तरहथा। लेकिन केन्द्रित मार्गदर्शन विस्तरणक्षमता हीकम देते शीघ्र उसीएक् विस्तरणक्षमता वर्णित विकसित परिभाषित जाने मुख्य करते बलवान उद्योग उदेशीत स्वतंत्रता बनाना विकसित आंतरजाल बीसबतेबोध वर्ष आवश्यकत समस्याओ माहितीवानीज्य रखति जिसकी

प्राप्त भोगोलिक चुनने सके। संदेश हीकम सुविधा बातसमय सदस्य सभिसमज माहितीवानीज्य जिवन करेसाथ एकत्रित निर्माता करके विशेष संभव मजबुत उपेक्ष कम्प्युटर विश्व मुखय हुएआदि अंग्रेजी अर्थपुर्ण हुआआदी हार्डवेर नीचे सहयोग सहयोग एसेएवं विनिमय मार्गदर्शन भीयह पहोचाना दस्तावेज विचारशिलता खरिदे गटको

शुरुआत तरहथा। किया कार्य जिम्मे आधुनिक समस्याए पहेला साधन अपनि यायेका गोपनीयता लेकिन वर्णन उनका ज्यादा यन्त्रालय डाले। प्रसारन विश्वास पुष्टिकर्ता आजपर दिये दौरान दस्तावेज प्रव्रुति विकासक्षमता एछित समजते पुस्तक वर्णित अनुवादक हमारी विषय दिशामे वेबजाल मुख्यतह अधिकार प्राधिकरन अमितकुमार वर्णन आजपर औषधिक ब्रौशर प्राथमिक विकेन्द्रित जिसे`

type test struct {
	name string
	text string

	expected      []string
	isInsensitive bool
}

func TestTop10(t *testing.T) {
	t.Run("Empty string Case", func(t *testing.T) {
		assert.Len(t, Top10("", true), 0)
	})

	t.Run("Special characters Case", func(t *testing.T) {
		assert.Len(t, Top10("\n\n\n", true), 0)
	})

	for _, tst := range [...]test{
		{
			name:          "Russian Insensitive Case",
			text:          russianText,
			expected:      []string{"он", "а", "и", "что", "ты", "не", "если", "то", "его", "кристофер", "робин", "в"},
			isInsensitive: true,
		},
		{
			name:     "Russian Sensitive Case",
			text:     russianText,
			expected: []string{"он", "и", "а", "ты", "что", "то", "Робин", "его", "если", "не", "Кристофер"},
		},
		{
			name:          "Hindi Both Cases",
			text:          hindiText,
			expected:      []string{"मुख्य", "मुख्य", "विश्वास", "मार्गदर्शन", "विशेष", "यायेका", "मजबुत", "ज्यादा", "आवश्यक", "वर्णन", "अंग्रेजी", "लेकिन", "अनुवाद", "पहोच।"},
			isInsensitive: true,
		},
	} {
		t.Run(tst.name, func(t *testing.T) {
			require.Subset(t, tst.expected, Top10(tst.text, tst.isInsensitive))
		})
	}
}
