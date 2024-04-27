#!/bin/zsh

# need system audio default to be set to something that knows japanese...
# e.g., add the japanese siri voice and set the default to that

# INP should be one word per line

INP='../merging/merged-kanji-cards.txt'
OUTP='../cards-for-import'

split -l 2000 ${INP}
files=(x*)
for item in ${files[*]}
do
  mv $item "${OUTP}/cards-800-${item}.txt"
done
