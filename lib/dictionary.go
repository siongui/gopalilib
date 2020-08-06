// Package lib is common library for Pāli Dictionary and Pāli Tipiṭaka.
// This package consists of common data structures and methods to be shared and
// used in front-end (browser), back-end (server), and offline data processing.
package lib

import (
	"html/template"
)

// BookInfo stores information of a dictionary book.
//
// Lang: language of the dictionary book.
//
//   zh: Chinese
//   ja: Japanese
//   en: English
//   vi: Vietnamese
//   my: Burmese(Myanmar)
//
// Separator: used to get short explanation of the word.
//
// Name: short name of the dictionary
//
// Author: name and author of the dictionary
type BookInfo struct {
	Lang      string `json:"lang"`
	Separator string `json:"separator"`
	Name      string `json:"name"`
	Author    string `json:"author"`
}

// BookIdAndInfos = object of key-value pairs, where
//
// key: Id of the dictionary, which is a single letter
//
// value: See BookInfo type
//
// The JSON:
//
//   {
//     "A": {
//       "lang": "ja",
//       "separator": " -",
//       "name": "《パーリ語辞典》",
//       "author": "増補改訂パーリ語辞典  水野弘元著"
//     },
//     "B": {
//       "lang": "my",
//       "separator": "。",
//       "name": "Pali Myanmar Dictionary",
//       "author": "Pali Word Grammar from Pali Myanmar Dictionary"
//     },
//     "C": {
//       "lang": "en",
//       "separator": "\u003cbr\u003e",
//       "name": "Concise P-E Dictionary    ",
//       "author": "Concise Pali-English Dictionary by A.P. Buddhadatta Mahathera"
//     },
//     "D": {
//       "lang": "zh",
//       "separator": "~",
//       "name": "《巴漢詞典》",
//       "author": "《巴漢詞典》Mahāñāṇo Bhikkhu編著"
//     },
//     "E": {
//       "lang": "vi",
//       "separator": "。",
//       "name": "Pali Viet Abhi- Terms",
//       "author": "Pali Viet Abhidhamma Terms  Từ điển các thuật ngữ Vô Tỷ Pháp của ngài Tịnh Sự, được chép từ phần ghi chú thuật ngữ trong các bản dịch của ngài."
//     },
//     "F": {
//       "lang": "zh",
//       "separator": "。",
//       "name": "《巴漢詞典》",
//       "author": "《巴漢詞典》明法尊者增訂"
//     },
//     "G": {
//       "lang": "zh",
//       "separator": "。",
//       "name": "《巴利語字匯》",
//       "author": "四念住課程開示集要巴利語字匯（葛印卡）"
//     },
//     "H": {
//       "lang": "zh",
//       "separator": " -",
//       "name": "《漢譯パーリ語辭典》",
//       "author": "漢譯パーリ語辭典 黃秉榮譯"
//     },
//     "I": {
//       "lang": "en",
//       "separator": "。",
//       "name": "Pali-Dictonary from VRI",
//       "author": "Pali-Dictionary Vipassana Research Institute"
//     },
//     "J": {
//       "lang": "zh",
//       "separator": "。",
//       "name": "《パーリ語辭典-勘誤表》",
//       "author": "《水野弘元-巴利語辭典-勘誤表》 Bhikkhu Santagavesaka 覓寂尊者"
//     },
//     "K": {
//       "lang": "my",
//       "separator": "。",
//       "name": "Tipiṭaka Pāḷi-Myanmar Dictionary",
//       "author": "Tipiṭaka Pāḷi-Myanmar Dictionary တိပိဋက-ပါဠိျမန္မာ အဘိဓာန္"
//     },
//     "M": {
//       "lang": "zh",
//       "separator": "。",
//       "name": "《巴利語匯解》",
//       "author": "巴利語匯解\u0026巴利新音譯 瑪欣德尊者"
//     },
//     "N": {
//       "lang": "en",
//       "separator": "\u003cbr\u003e",
//       "name": "Buddhist Dictionary",
//       "author": "Buddhist Dictionary by NYANATILOKA MAHATHERA"
//     },
//     "O": {
//       "lang": "my",
//       "separator": "。",
//       "name": "Pali Roots Dictionary",
//       "author": "Pali Roots Dictionary ဓါတ္အဘိဓာန္"
//     },
//     "P": {
//       "lang": "en",
//       "separator": "\u003ci\u003e",
//       "name": "PTS P-E Dictionary",
//       "author": "PTS Pali-English dictionary The Pali Text Society's Pali-English dictionary"
//     },
//     "Q": {
//       "lang": "vi",
//       "separator": "。",
//       "name": "Pali Viet Vinaya Terms",
//       "author": "Pali Viet Vinaya Terms  Từ điển các thuật ngữ về luật do tỳ khưu Giác Nguyên sưu tầm."
//     },
//     "R": {
//       "lang": "my",
//       "separator": "。",
//       "name": "U Hau Sein’s Pāḷi-Myanmar Dictionary",
//       "author": "U Hau Sein’s Pāḷi-Myanmar Dictionary ပါဠိျမန္မာ အဘိဓာန္(ဦးဟုတ္စိန္)"
//     },
//     "S": {
//       "lang": "ja",
//       "separator": " -",
//       "name": "《パーリ語辞典》",
//       "author": "パーリ語辞典  水野弘元著"
//     },
//     "T": {
//       "lang": "zh",
//       "separator": " -",
//       "name": "《漢譯パーリ語辭典》",
//       "author": "漢譯パーリ語辭典 李瑩譯"
//     },
//     "U": {
//       "lang": "vi",
//       "separator": "。",
//       "name": "Pali Viet Dictionary",
//       "author": "Pali Viet Dictionary  Bản dịch của ngài Bửu Chơn."
//     },
//     "V": {
//       "lang": "en",
//       "separator": "。",
//       "name": "Pali Proper Names Dictionary",
//       "author": "Buddhist Dictionary of Pali Proper Names by G P Malalasekera"
//     },
//     "W": {
//       "lang": "zh",
//       "separator": "。",
//       "name": "《巴英術語匯編》",
//       "author": "巴英術語匯編 《法的醫療》附 溫宗堃"
//     },
//     "X": {
//       "lang": "zh",
//       "separator": "。",
//       "name": "《巴利語入門》",
//       "author": "《巴利語入門》釋性恩(Dhammajīvī)"
//     },
//     "Z": {
//       "lang": "zh",
//       "separator": "。",
//       "name": "《巴漢佛學辭匯》",
//       "author": "巴利文-漢文佛學名相辭匯 翻譯：張文明"
//     }
//   }
//
type BookIdAndInfos map[string]BookInfo

// Array of ["book id", "word explanation"]
//
// Example:
//
//   [
//      [
//         "F",
//         "dah\uff0c= =\u71d2\uff08burn\uff09\u3002cp\uff0e\uff08\u5df4dah\uff09\ufe50\u3010\u5b57\u6839I.\u3011\u6d41\u6d6a\u751f\u6d3b\u3001\u4e5e\u8a0e\uff08to bum\uff09\u3002"
//      ],
//      [
//         "F",
//         "dah\uff0c\ufe50\u3010\u5b57\u6839I.\u30111.\u5efa\u7acb\uff08to establish\uff09\u3001\u653e\u7f6e\uff08to place\uff09\u3002\u2192saddh\u0101\u2039\uff08sa\u1e41+dah\u5efa\u7acb\u3001\u653e\u7f6e\uff09\uff0c\u68b5sraddh\u0101\uff1bfaith \uff09\u30022.\u6d41\u6d6a\u751f\u6d3b\u3001\u4e5e\u8a0e\uff08to bum\uff09\u3002cp\uff0e\uff08\u68b5dah\uff09\uff0c\u71d2\uff08burn\uff09\u3002"
//      ],
//      [
//         "F",
//         "dah\uff0c\ufe50\u3010\u5b57\u6839I.\u3011\u6d41\u6d6a\u751f\u6d3b\u3001\u4e5e\u8a0e\uff08to bum\uff09\u3002cp\uff0e\uff08\u68b5dah\uff09\uff0c\u71d2\uff08burn\uff09\u3002"
//      ]
//   ]
type BookIdWordExps [][2]string

type BookNameWordExp struct {
	BookName    string
	Explanation template.HTML
}

func IdExps2BookNameWordExps(ies []IdExp, di BookIdAndInfos) []BookNameWordExp {
	var result []BookNameWordExp

	for _, ie := range ies {
		result = append(result, BookNameWordExp{
			BookName:    di[ie.Id].Author,
			Explanation: template.HTML(ie.Exp),
		})
	}

	return result
}
