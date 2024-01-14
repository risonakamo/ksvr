# sentence table
stores sentences. keyed by md5 hash of each sentence

## key
`sentence:{md5 hash}`

md5 hash of the entire sentence, maybe shortened to some length.

## value
- sentence str: the whole sentence string
- kanjis []str: (derived) list of kanjis extracted from the sentence

## add rules
- each sentence added to this table causes multiple additions to kanji table. each kanji in the sentence modifies the kanji table.
    - kanji table can be derived from this table, so regeneration is possible

# kanji table
derived table. stores singular kanjis and their example sentences. can be fully reconstructed from the sentence table.

used to quickly retrieve information about single kanjis.

## key
`kanji:{kanji rune}`

## value
- sentences []str: list of sentence hashes.

## add rules
this table should only be added to indirectly from adding to sentence table.