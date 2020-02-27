package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ikawaha/kagome/tokenizer"
)

//形態素解析をします
func Separate(word string) []string {
	words := []string{}
	t := tokenizer.New()
	morphs := t.Tokenize(word)
	for _, m := range morphs {
		words = append(words, m.Surface)
	}
	return words
}

//単語ずつのブロックに分けていきます
func Gather(words []string, words2 []string) [][]string {
	result := [][]string{}
	for i := 0; i < 2; i++ {
		if i == 0 {
			if len(words) < 3 {
				return result
			}
			for i := 0; i < len(words)-2; i++ {
				Block := []string{words[i], words[i+1], words[i+2]} //ここの要素の数でブロックの単語の数を決めます
				result = append(result, Block)
			}
		} else if i == 1 {
			if len(words2) < 3 {
				return result
			}
			for i := 0; i < len(words2)-2; i++ {
				Block := []string{words2[i], words2[i+1], words2[i+2]}//こっちも上と同じ要素数にしておいてください
				result = append(result, Block)
			}
		}
	}
	return result
}

//先頭の単語を見つけます
func Find(array [][]string, target string) [][]string {
	block := [][]string{}
	for _, s := range array {
		if s[0] == target {
			block = append(block, s)
		}
	}
	return block
}

//先頭以外の言葉をくっつけます
func Connect(array [][]string, list []string) []string {
	rand.Seed((time.Now().UnixNano()))
	i := 0
	for _, word := range array[rand.Intn(len(array))] {
		if i != 0 {
			list = append(list, word)
		}
		i++
	}
	return list
}

//文章を生成します
func Markov(array [][]string) string {
	result := ""
	words := []string{}
	block := [][]string{}
	count := 0

	block = Find(array, "BOS")
	words = Connect(block, words)
	for words[len(words)-1] != "EOS" {
		block = Find(array, words[len(words)-1])
		if len(block) == 0 {
			break
		}
		words = Connect(block, words)
		count++
		if count == 150 {
			break
		}
	}
	for _, s := range words {
		result += s
	}
	return result
}

func main() {
	s1 := "私はカードを拾います" //この二つに文章をいれます
	s2 := "彼は神を束ねます"    //上と同じ文章でもそうでなくても起動はします
	s3 := Gather(Separate(s1), Separate(s2))
	fmt.Println(Markov(s3))
}
