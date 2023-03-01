# char to binary
コマンドライン引数で受け取った文字列を、１文字ずつ `UTF-8` のバイナリに変換する

## デモ
input
```
ちぃかわ
```

output
```
1文字目 ち
[0]:11100011
[1]:10000001
[2]:10100001

2文字目 ぃ
[0]:11100011
[1]:10000001
[2]:10000011

3文字目 か
[0]:11100011
[1]:10000001
[2]:10001011

4文字目 わ
[0]:11100011
[1]:10000010
[2]:10001111

```