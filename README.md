# 800kanji

## overview

1. acquire a sorted list of 400 *more* kanji by frequency: #401-800
2. get 10 sample words for each character
3. produce 3 versions of the word list
- english, kanji+hiragana
- kanji, english+hiragana
- audio, kanji+hiragana+english
4. ask chatgpt for a sample sentence for each word on the word list:
   `i am going to provide a list of japanese words. for each word on the list, please generate a short sample sentence in japanese that uses this word. prepare the sample sentence in two saparate versions. the first is a kanji version. the second is a hiragana version. next produce an english translation of this japanese sentence. format your output as follows: the word in brackets, a tab, the kanji sentence, a tab, the hiragana version of the kanji sentence, a tab, the english translation. for example, if the word is '同日', then the output would look like '(同日) 	同日に会議があります。	どうじつにかいぎがあります。	There is a meeting on the same day.' generate an output for each of the words on the list provided. continue generating sentences until you have performed this task for every item on the list. the list follows. 期待, 招待,...`
5. produce multiple versions of the sentence list
- kanji, english+hiragana
- audio, kanji+hiragana+english
6. use shell scripts to generate the audio files via macOS `say` + `ffmpeg`:
- `shell-scripts/kanji-audio_conversion.sh`
- `shell-scripts/sent-audio_conversion.sh`
7. move audio files into `collection.media`
8. use `800kanji-builder` application to interleave the lists from above:
- `go build`
- `mv 800kanji-builder ./merging`
- `cd ./merging`
- `./800kanji-builder`
9. use `split` to split the merged product into chunks: `shell-scripts/splitter.sh`
10. import each chunk into `anki`
11. export the whole back out from `anki` to an `.apkg`

## caveats

1. irregularities in the `chatgpt` output have not been smoothed over
2. `chatgpt` output; so, buyer beware
3. the number of sentences is out of sync with the number of kanji (*fixme...*)