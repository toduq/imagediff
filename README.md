# imagediff

2枚の画像のPSNRとSSIMを計算します

PSNR : [ピーク信号対雑音比](https://ja.wikipedia.org/wiki/%E3%83%94%E3%83%BC%E3%82%AF%E4%BF%A1%E5%8F%B7%E5%AF%BE%E9%9B%91%E9%9F%B3%E6%AF%94)
SSIM : [Structural similarity](https://en.wikipedia.org/wiki/Structural_similarity)

# インストール

```
go get github.com/toduq/imagediff
```

# 実行

```
imagediff -m METHOD file1 file2
# METHODにはpsnrまたはssimを指定できます(デフォルトはpsnrです)
```
