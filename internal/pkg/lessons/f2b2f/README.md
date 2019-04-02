# f2b2f
Front-to-back-to-front

## Versions
Presently 3 versions:
* v1 or, the initial version, is a simple for-loop, which iterates over the slice, flipping it with each iteration
* v2 is a recursive version, which also takes into account the special case of negative input
* v3 is a linear-time version, whose existence was pointed out to me and which I subsequently implemented 

## Benchmarks
Running the benchmarks as `go test -a -bench=. -run=XXX -benchmem -benchtime 2s -timeout 20m . > bench-v1.txt` on each respective version (with the same input) allows using `benchcmp` to compare them.

~~Details below, but the short conclusion: The recursive version (v2) scales better, is faster and uses less memory than the others.~~

The short conclusion is: When using only positive `int` values as input, the recursive (v2) version outperforms the other two.

The slightly longer conclusion is: When you start to test the `hustle([]string) []string` function on increasingly longer strings, above the range of possible `int` input values, the linear (v3) outperforms the other two in terms of speed, whilst the recursive (v2) version outperforms the others in terms of allocs and memory (0 and 0).

### Version 1 to Version 2
```
benchmark                                        old ns/op     new ns/op     delta
BenchmarkSolution/bench-3-8                      103           68.7          -33.30%
BenchmarkSolution/bench-71-8                     273           155           -43.22%
BenchmarkSolution/bench-613-8                    428           235           -45.09%
BenchmarkSolution/bench-2029-8                   474           271           -42.83%
BenchmarkSolution/bench-74035-8                  651           321           -50.69%
BenchmarkSolution/bench-261402-8                 660           353           -46.52%
BenchmarkSolution/bench-2629359-8                726           402           -44.63%
BenchmarkSolution/bench-24547482-8               747           438           -41.37%
BenchmarkSolution/bench-465093345-8              963           499           -48.18%
BenchmarkSolution/bench-8295678306-8             954           538           -43.61%
BenchmarkSolution/bench-85605547386-8            961           537           -44.12%
BenchmarkSolution/bench-528789347646-8           1035          634           -38.74%
BenchmarkSolution/bench-6421242908142-8          1089          685           -37.10%
BenchmarkSolution/bench-71771243324629-8         1147          734           -36.01%
BenchmarkSolution/bench-704229078898971-8        1163          797           -31.47%
BenchmarkSolution/bench-4046455185423042-8       1321          843           -36.18%
BenchmarkSolution/bench-29461204712002348-8      1475          917           -37.83%
BenchmarkSolution/bench-954492294769786752-8     1591          981           -38.34%

benchmark                                        old allocs     new allocs     delta
BenchmarkSolution/bench-3-8                      2              1              -50.00%
BenchmarkSolution/bench-71-8                     4              2              -50.00%
BenchmarkSolution/bench-613-8                    6              3              -50.00%
BenchmarkSolution/bench-2029-8                   6              3              -50.00%
BenchmarkSolution/bench-74035-8                  7              3              -57.14%
BenchmarkSolution/bench-261402-8                 7              3              -57.14%
BenchmarkSolution/bench-2629359-8                7              3              -57.14%
BenchmarkSolution/bench-24547482-8               7              3              -57.14%
BenchmarkSolution/bench-465093345-8              8              3              -62.50%
BenchmarkSolution/bench-8295678306-8             8              3              -62.50%
BenchmarkSolution/bench-85605547386-8            8              3              -62.50%
BenchmarkSolution/bench-528789347646-8           8              3              -62.50%
BenchmarkSolution/bench-6421242908142-8          8              3              -62.50%
BenchmarkSolution/bench-71771243324629-8         8              3              -62.50%
BenchmarkSolution/bench-704229078898971-8        8              3              -62.50%
BenchmarkSolution/bench-4046455185423042-8       8              3              -62.50%
BenchmarkSolution/bench-29461204712002348-8      9              3              -66.67%
BenchmarkSolution/bench-954492294769786752-8     9              3              -66.67%

benchmark                                        old bytes     new bytes     delta
BenchmarkSolution/bench-3-8                      32            16            -50.00%
BenchmarkSolution/bench-71-8                     82            34            -58.54%
BenchmarkSolution/bench-613-8                    166           54            -67.47%
BenchmarkSolution/bench-2029-8                   184           72            -60.87%
BenchmarkSolution/bench-74035-8                  330           90            -72.73%
BenchmarkSolution/bench-261402-8                 352           112           -68.18%
BenchmarkSolution/bench-2629359-8                368           128           -65.22%
BenchmarkSolution/bench-24547482-8               384           144           -62.50%
BenchmarkSolution/bench-465093345-8              672           176           -73.81%
BenchmarkSolution/bench-8295678306-8             688           192           -72.09%
BenchmarkSolution/bench-85605547386-8            688           192           -72.09%
BenchmarkSolution/bench-528789347646-8           720           224           -68.89%
BenchmarkSolution/bench-6421242908142-8          736           240           -67.39%
BenchmarkSolution/bench-71771243324629-8         752           256           -65.96%
BenchmarkSolution/bench-704229078898971-8        768           272           -64.58%
BenchmarkSolution/bench-4046455185423042-8       784           288           -63.27%
BenchmarkSolution/bench-29461204712002348-8      1360          352           -74.12%
BenchmarkSolution/bench-954492294769786752-8     1360          352           -74.12%
```

### Version 2 to Version 3
```
benchmark                                        old ns/op     new ns/op     delta
BenchmarkSolution/bench-3-8                      68.7          50.2          -26.93%
BenchmarkSolution/bench-71-8                     155           88.7          -42.77%
BenchmarkSolution/bench-613-8                    235           436           +85.53%
BenchmarkSolution/bench-2029-8                   271           471           +73.80%
BenchmarkSolution/bench-74035-8                  321           626           +95.02%
BenchmarkSolution/bench-261402-8                 353           619           +75.35%
BenchmarkSolution/bench-2629359-8                402           654           +62.69%
BenchmarkSolution/bench-24547482-8               438           706           +61.19%
BenchmarkSolution/bench-465093345-8              499           885           +77.35%
BenchmarkSolution/bench-8295678306-8             538           892           +65.80%
BenchmarkSolution/bench-85605547386-8            537           888           +65.36%
BenchmarkSolution/bench-528789347646-8           634           924           +45.74%
BenchmarkSolution/bench-6421242908142-8          685           966           +41.02%
BenchmarkSolution/bench-71771243324629-8         734           1012          +37.87%
BenchmarkSolution/bench-704229078898971-8        797           983           +23.34%
BenchmarkSolution/bench-4046455185423042-8       843           1016          +20.52%
BenchmarkSolution/bench-29461204712002348-8      917           1286          +40.24%
BenchmarkSolution/bench-954492294769786752-8     981           1308          +33.33%

benchmark                                        old allocs     new allocs     delta
BenchmarkSolution/bench-3-8                      1              1              +0.00%
BenchmarkSolution/bench-71-8                     2              1              -50.00%
BenchmarkSolution/bench-613-8                    3              6              +100.00%
BenchmarkSolution/bench-2029-8                   3              6              +100.00%
BenchmarkSolution/bench-74035-8                  3              7              +133.33%
BenchmarkSolution/bench-261402-8                 3              7              +133.33%
BenchmarkSolution/bench-2629359-8                3              7              +133.33%
BenchmarkSolution/bench-24547482-8               3              7              +133.33%
BenchmarkSolution/bench-465093345-8              3              8              +166.67%
BenchmarkSolution/bench-8295678306-8             3              8              +166.67%
BenchmarkSolution/bench-85605547386-8            3              8              +166.67%
BenchmarkSolution/bench-528789347646-8           3              8              +166.67%
BenchmarkSolution/bench-6421242908142-8          3              8              +166.67%
BenchmarkSolution/bench-71771243324629-8         3              8              +166.67%
BenchmarkSolution/bench-704229078898971-8        3              8              +166.67%
BenchmarkSolution/bench-4046455185423042-8       3              8              +166.67%
BenchmarkSolution/bench-29461204712002348-8      3              9              +200.00%
BenchmarkSolution/bench-954492294769786752-8     3              9              +200.00%

benchmark                                        old bytes     new bytes     delta
BenchmarkSolution/bench-3-8                      16            16            +0.00%
BenchmarkSolution/bench-71-8                     34            32            -5.88%
BenchmarkSolution/bench-613-8                    54            166           +207.41%
BenchmarkSolution/bench-2029-8                   72            184           +155.56%
BenchmarkSolution/bench-74035-8                  90            330           +266.67%
BenchmarkSolution/bench-261402-8                 112           352           +214.29%
BenchmarkSolution/bench-2629359-8                128           368           +187.50%
BenchmarkSolution/bench-24547482-8               144           384           +166.67%
BenchmarkSolution/bench-465093345-8              176           672           +281.82%
BenchmarkSolution/bench-8295678306-8             192           688           +258.33%
BenchmarkSolution/bench-85605547386-8            192           688           +258.33%
BenchmarkSolution/bench-528789347646-8           224           720           +221.43%
BenchmarkSolution/bench-6421242908142-8          240           736           +206.67%
BenchmarkSolution/bench-71771243324629-8         256           752           +193.75%
BenchmarkSolution/bench-704229078898971-8        272           768           +182.35%
BenchmarkSolution/bench-4046455185423042-8       288           784           +172.22%
BenchmarkSolution/bench-29461204712002348-8      352           1360          +286.36%
BenchmarkSolution/bench-954492294769786752-8     352           1360          +286.36%
```

### Version 1 to Version 3
```
benchmark                                        old ns/op     new ns/op     delta
BenchmarkSolution/bench-3-8                      103           50.2          -51.26%
BenchmarkSolution/bench-71-8                     273           88.7          -67.51%
BenchmarkSolution/bench-613-8                    428           436           +1.87%
BenchmarkSolution/bench-2029-8                   474           471           -0.63%
BenchmarkSolution/bench-74035-8                  651           626           -3.84%
BenchmarkSolution/bench-261402-8                 660           619           -6.21%
BenchmarkSolution/bench-2629359-8                726           654           -9.92%
BenchmarkSolution/bench-24547482-8               747           706           -5.49%
BenchmarkSolution/bench-465093345-8              963           885           -8.10%
BenchmarkSolution/bench-8295678306-8             954           892           -6.50%
BenchmarkSolution/bench-85605547386-8            961           888           -7.60%
BenchmarkSolution/bench-528789347646-8           1035          924           -10.72%
BenchmarkSolution/bench-6421242908142-8          1089          966           -11.29%
BenchmarkSolution/bench-71771243324629-8         1147          1012          -11.77%
BenchmarkSolution/bench-704229078898971-8        1163          983           -15.48%
BenchmarkSolution/bench-4046455185423042-8       1321          1016          -23.09%
BenchmarkSolution/bench-29461204712002348-8      1475          1286          -12.81%
BenchmarkSolution/bench-954492294769786752-8     1591          1308          -17.79%

benchmark                                        old allocs     new allocs     delta
BenchmarkSolution/bench-3-8                      2              1              -50.00%
BenchmarkSolution/bench-71-8                     4              1              -75.00%
BenchmarkSolution/bench-613-8                    6              6              +0.00%
BenchmarkSolution/bench-2029-8                   6              6              +0.00%
BenchmarkSolution/bench-74035-8                  7              7              +0.00%
BenchmarkSolution/bench-261402-8                 7              7              +0.00%
BenchmarkSolution/bench-2629359-8                7              7              +0.00%
BenchmarkSolution/bench-24547482-8               7              7              +0.00%
BenchmarkSolution/bench-465093345-8              8              8              +0.00%
BenchmarkSolution/bench-8295678306-8             8              8              +0.00%
BenchmarkSolution/bench-85605547386-8            8              8              +0.00%
BenchmarkSolution/bench-528789347646-8           8              8              +0.00%
BenchmarkSolution/bench-6421242908142-8          8              8              +0.00%
BenchmarkSolution/bench-71771243324629-8         8              8              +0.00%
BenchmarkSolution/bench-704229078898971-8        8              8              +0.00%
BenchmarkSolution/bench-4046455185423042-8       8              8              +0.00%
BenchmarkSolution/bench-29461204712002348-8      9              9              +0.00%
BenchmarkSolution/bench-954492294769786752-8     9              9              +0.00%

benchmark                                        old bytes     new bytes     delta
BenchmarkSolution/bench-3-8                      32            16            -50.00%
BenchmarkSolution/bench-71-8                     82            32            -60.98%
BenchmarkSolution/bench-613-8                    166           166           +0.00%
BenchmarkSolution/bench-2029-8                   184           184           +0.00%
BenchmarkSolution/bench-74035-8                  330           330           +0.00%
BenchmarkSolution/bench-261402-8                 352           352           +0.00%
BenchmarkSolution/bench-2629359-8                368           368           +0.00%
BenchmarkSolution/bench-24547482-8               384           384           +0.00%
BenchmarkSolution/bench-465093345-8              672           672           +0.00%
BenchmarkSolution/bench-8295678306-8             688           688           +0.00%
BenchmarkSolution/bench-85605547386-8            688           688           +0.00%
BenchmarkSolution/bench-528789347646-8           720           720           +0.00%
BenchmarkSolution/bench-6421242908142-8          736           736           +0.00%
BenchmarkSolution/bench-71771243324629-8         752           752           +0.00%
BenchmarkSolution/bench-704229078898971-8        768           768           +0.00%
BenchmarkSolution/bench-4046455185423042-8       784           784           +0.00%
BenchmarkSolution/bench-29461204712002348-8      1360          1360          +0.00%
BenchmarkSolution/bench-954492294769786752-8     1360          1360          +0.00%
```
