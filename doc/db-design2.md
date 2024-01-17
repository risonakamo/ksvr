# summary
v1

- sentence table: `sentence:{md5 hash} -> SentenceInfo`
    - stores sentence info, keyed by hash of sentence

v2

- sentence table: `sentence:{md5 hash} -> str`
    - stores sentences, keyed by hash of sentence

# other types
```typescript
interface SentenceInfo
{
    // the whole sentence string
    sentence: string
    // list of kanjis extracted from the string
    kanjis: string[]
}
```