# f2b2f
Front-to-back-to-front

## Versions
Presently 3 versions:
* v1 or, the initial version, is a simple for-loop, which iterates over the slice, flipping it with each iteration
* v2 is a recursive version
* v3 is a single-pass version, which also uses a pre-allocated slice to circumvent the use of append    

## Benchmarks
Running the benchmarks as `go test -a -bench=. -run=XXX -benchmem -benchtime 2s -timeout 20m . > bench-v1.txt` on each respective version (with the same input) allows using `benchcmp` to compare them.

~~Details below, but the short conclusion: The recursive version (v2) scales better, is faster and uses less memory than the others.~~

The short conclusion is: When using only positive `int` values as input, the recursive (v2) version outperforms the other two.

The slightly longer conclusion is: When you start to test the `hustle([]string) []string` function on increasingly longer strings, above the range of possible `int` input values, the linear (v3) outperforms the other two in terms of speed, whilst the recursive (v2) version outperforms the others in terms of allocs and memory (0 and 0).

Even longer conclusion, pre-alloc

### Version 1 to Version 2
```
benchmark                                        old ns/op      new ns/op      delta
Benchmark_hustle/2-8                             3.42           3.54           +3.51%
Benchmark_hustle/4-8                             257            30.7           -88.05%
Benchmark_hustle/8-8                             448            99.5           -77.79%
Benchmark_hustle/16-8                            844            282            -66.59%
Benchmark_hustle/32-8                            1823           969            -46.85%
Benchmark_hustle/64-8                            4701           3270           -30.44%
Benchmark_hustle/128-8                           13495          11245          -16.67%
Benchmark_hustle/256-8                           43945          39662          -9.75%
Benchmark_hustle/512-8                           155347         148253         -4.57%
Benchmark_hustle/1024-8                          581838         571211         -1.83%
Benchmark_hustle/2048-8                          2308275        2334525        +1.14%
Benchmark_hustle/4096-8                          9577424        8968750        -6.36%
Benchmark_hustle/8192-8                          36272902       35261140       -2.79%
Benchmark_hustle/16384-8                         147704701      141311542      -4.33%
Benchmark_hustle/32768-8                         585325823      569298810      -2.74%
Benchmark_hustle/65536-8                         2325228042     2264542728     -2.61%
BenchmarkSolution/bench-3-8                      61.2           59.0           -3.59%
BenchmarkSolution/bench-71-8                     140            138            -1.43%
BenchmarkSolution/bench-613-8                    437            223            -48.97%
BenchmarkSolution/bench-2029-8                   432            223            -48.38%
BenchmarkSolution/bench-74035-8                  613            314            -48.78%
BenchmarkSolution/bench-261402-8                 616            312            -49.35%
BenchmarkSolution/bench-2629359-8                717            393            -45.19%
BenchmarkSolution/bench-24547482-8               681            392            -42.44%
BenchmarkSolution/bench-465093345-8              904            494            -45.35%
BenchmarkSolution/bench-8295678306-8             985            546            -44.57%
BenchmarkSolution/bench-85605547386-8            1015           605            -40.39%
BenchmarkSolution/bench-528789347646-8           1089           646            -40.68%
BenchmarkSolution/bench-6421242908142-8          1114           707            -36.54%
BenchmarkSolution/bench-71771243324629-8         1175           761            -35.23%
BenchmarkSolution/bench-704229078898971-8        1168           754            -35.45%
BenchmarkSolution/bench-4046455185423042-8       1227           821            -33.09%
BenchmarkSolution/bench-29461204712002348-8      1512           923            -38.96%
BenchmarkSolution/bench-954492294769786752-8     1539           1005           -34.70%

benchmark                                        old allocs     new allocs     delta
Benchmark_hustle/2-8                             0              0              +0.00%
Benchmark_hustle/4-8                             3              0              -100.00%
Benchmark_hustle/8-8                             4              0              -100.00%
Benchmark_hustle/16-8                            5              0              -100.00%
Benchmark_hustle/32-8                            6              0              -100.00%
Benchmark_hustle/64-8                            7              0              -100.00%
Benchmark_hustle/128-8                           8              0              -100.00%
Benchmark_hustle/256-8                           9              0              -100.00%
Benchmark_hustle/512-8                           10             0              -100.00%
Benchmark_hustle/1024-8                          11             0              -100.00%
Benchmark_hustle/2048-8                          14             0              -100.00%
Benchmark_hustle/4096-8                          16             0              -100.00%
Benchmark_hustle/8192-8                          19             0              -100.00%
Benchmark_hustle/16384-8                         22             0              -100.00%
Benchmark_hustle/32768-8                         25             0              -100.00%
Benchmark_hustle/65536-8                         28             0              -100.00%
BenchmarkSolution/bench-3-8                      1              1              +0.00%
BenchmarkSolution/bench-71-8                     2              2              +0.00%
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
Benchmark_hustle/2-8                             0             0             +0.00%
Benchmark_hustle/4-8                             112           0             -100.00%
Benchmark_hustle/8-8                             240           0             -100.00%
Benchmark_hustle/16-8                            496           0             -100.00%
Benchmark_hustle/32-8                            1008          0             -100.00%
Benchmark_hustle/64-8                            2032          0             -100.00%
Benchmark_hustle/128-8                           4080          0             -100.00%
Benchmark_hustle/256-8                           8176          0             -100.00%
Benchmark_hustle/512-8                           16368         0             -100.00%
Benchmark_hustle/1024-8                          32752         0             -100.00%
Benchmark_hustle/2048-8                          121456        0             -100.00%
Benchmark_hustle/4096-8                          252528        0             -100.00%
Benchmark_hustle/8192-8                          629360        0             -100.00%
Benchmark_hustle/16384-8                         1383024       0             -100.00%
Benchmark_hustle/32768-8                         2882160       0             -100.00%
Benchmark_hustle/65536-8                         5831280       0             -100.00%
BenchmarkSolution/bench-3-8                      16            16            +0.00%
BenchmarkSolution/bench-71-8                     34            34            +0.00%
BenchmarkSolution/bench-613-8                    166           54            -67.47%
BenchmarkSolution/bench-2029-8                   166           54            -67.47%
BenchmarkSolution/bench-74035-8                  330           90            -72.73%
BenchmarkSolution/bench-261402-8                 330           90            -72.73%
BenchmarkSolution/bench-2629359-8                368           128           -65.22%
BenchmarkSolution/bench-24547482-8               368           128           -65.22%
BenchmarkSolution/bench-465093345-8              672           176           -73.81%
BenchmarkSolution/bench-8295678306-8             688           192           -72.09%
BenchmarkSolution/bench-85605547386-8            704           208           -70.45%
BenchmarkSolution/bench-528789347646-8           720           224           -68.89%
BenchmarkSolution/bench-6421242908142-8          736           240           -67.39%
BenchmarkSolution/bench-71771243324629-8         752           256           -65.96%
BenchmarkSolution/bench-704229078898971-8        752           256           -65.96%
BenchmarkSolution/bench-4046455185423042-8       768           272           -64.58%
BenchmarkSolution/bench-29461204712002348-8      1360          352           -74.12%
BenchmarkSolution/bench-954492294769786752-8     1360          352           -74.12%
```

### Version 2 to Version 3
```
benchmark                                        old ns/op      new ns/op     delta
Benchmark_hustle/2-8                             3.54           3.42          -3.39%
Benchmark_hustle/4-8                             30.7           93.9          +205.86%
Benchmark_hustle/8-8                             99.5           132           +32.66%
Benchmark_hustle/16-8                            282            213           -24.47%
Benchmark_hustle/32-8                            969            362           -62.64%
Benchmark_hustle/64-8                            3270           719           -78.01%
Benchmark_hustle/128-8                           11245          1420          -87.37%
Benchmark_hustle/256-8                           39662          2816          -92.90%
Benchmark_hustle/512-8                           148253         5203          -96.49%
Benchmark_hustle/1024-8                          571211         9850          -98.28%
Benchmark_hustle/2048-8                          2334525        18861         -99.19%
Benchmark_hustle/4096-8                          8968750        42273         -99.53%
Benchmark_hustle/8192-8                          35261140       77438         -99.78%
Benchmark_hustle/16384-8                         141311542      168391        -99.88%
Benchmark_hustle/32768-8                         569298810      357075        -99.94%
Benchmark_hustle/65536-8                         2264542728     751634        -99.97%
BenchmarkSolution/bench-3-8                      59.0           60.9          +3.22%
BenchmarkSolution/bench-71-8                     138            142           +2.90%
BenchmarkSolution/bench-613-8                    223            283           +26.91%
BenchmarkSolution/bench-2029-8                   223            279           +25.11%
BenchmarkSolution/bench-74035-8                  314            342           +8.92%
BenchmarkSolution/bench-261402-8                 312            350           +12.18%
BenchmarkSolution/bench-2629359-8                393            407           +3.56%
BenchmarkSolution/bench-24547482-8               392            434           +10.71%
BenchmarkSolution/bench-465093345-8              494            503           +1.82%
BenchmarkSolution/bench-8295678306-8             546            503           -7.88%
BenchmarkSolution/bench-85605547386-8            605            541           -10.58%
BenchmarkSolution/bench-528789347646-8           646            575           -10.99%
BenchmarkSolution/bench-6421242908142-8          707            595           -15.84%
BenchmarkSolution/bench-71771243324629-8         761            629           -17.35%
BenchmarkSolution/bench-704229078898971-8        754            615           -18.44%
BenchmarkSolution/bench-4046455185423042-8       821            649           -20.95%
BenchmarkSolution/bench-29461204712002348-8      923            748           -18.96%
BenchmarkSolution/bench-954492294769786752-8     1005           757           -24.68%

benchmark                                        old allocs     new allocs     delta
Benchmark_hustle/2-8                             0              0              +0.00%
Benchmark_hustle/4-8                             0              1              +Inf%
Benchmark_hustle/8-8                             0              1              +Inf%
Benchmark_hustle/16-8                            0              1              +Inf%
Benchmark_hustle/32-8                            0              1              +Inf%
Benchmark_hustle/64-8                            0              1              +Inf%
Benchmark_hustle/128-8                           0              1              +Inf%
Benchmark_hustle/256-8                           0              1              +Inf%
Benchmark_hustle/512-8                           0              1              +Inf%
Benchmark_hustle/1024-8                          0              1              +Inf%
Benchmark_hustle/2048-8                          0              1              +Inf%
Benchmark_hustle/4096-8                          0              1              +Inf%
Benchmark_hustle/8192-8                          0              1              +Inf%
Benchmark_hustle/16384-8                         0              1              +Inf%
Benchmark_hustle/32768-8                         0              1              +Inf%
Benchmark_hustle/65536-8                         0              1              +Inf%
BenchmarkSolution/bench-3-8                      1              1              +0.00%
BenchmarkSolution/bench-71-8                     2              2              +0.00%
BenchmarkSolution/bench-613-8                    3              4              +33.33%
BenchmarkSolution/bench-2029-8                   3              4              +33.33%
BenchmarkSolution/bench-74035-8                  3              4              +33.33%
BenchmarkSolution/bench-261402-8                 3              4              +33.33%
BenchmarkSolution/bench-2629359-8                3              4              +33.33%
BenchmarkSolution/bench-24547482-8               3              4              +33.33%
BenchmarkSolution/bench-465093345-8              3              4              +33.33%
BenchmarkSolution/bench-8295678306-8             3              4              +33.33%
BenchmarkSolution/bench-85605547386-8            3              4              +33.33%
BenchmarkSolution/bench-528789347646-8           3              4              +33.33%
BenchmarkSolution/bench-6421242908142-8          3              4              +33.33%
BenchmarkSolution/bench-71771243324629-8         3              4              +33.33%
BenchmarkSolution/bench-704229078898971-8        3              4              +33.33%
BenchmarkSolution/bench-4046455185423042-8       3              4              +33.33%
BenchmarkSolution/bench-29461204712002348-8      3              4              +33.33%
BenchmarkSolution/bench-954492294769786752-8     3              4              +33.33%

benchmark                                        old bytes     new bytes     delta
Benchmark_hustle/2-8                             0             0             +0.00%
Benchmark_hustle/4-8                             0             64            +Inf%
Benchmark_hustle/8-8                             0             128           +Inf%
Benchmark_hustle/16-8                            0             256           +Inf%
Benchmark_hustle/32-8                            0             512           +Inf%
Benchmark_hustle/64-8                            0             1024          +Inf%
Benchmark_hustle/128-8                           0             2048          +Inf%
Benchmark_hustle/256-8                           0             4096          +Inf%
Benchmark_hustle/512-8                           0             8192          +Inf%
Benchmark_hustle/1024-8                          0             16384         +Inf%
Benchmark_hustle/2048-8                          0             32768         +Inf%
Benchmark_hustle/4096-8                          0             65536         +Inf%
Benchmark_hustle/8192-8                          0             131072        +Inf%
Benchmark_hustle/16384-8                         0             262147        +Inf%
Benchmark_hustle/32768-8                         0             524292        +Inf%
Benchmark_hustle/65536-8                         0             1048585       +Inf%
BenchmarkSolution/bench-3-8                      16            16            +0.00%
BenchmarkSolution/bench-71-8                     34            34            +0.00%
BenchmarkSolution/bench-613-8                    54            102           +88.89%
BenchmarkSolution/bench-2029-8                   54            102           +88.89%
BenchmarkSolution/bench-74035-8                  90            170           +88.89%
BenchmarkSolution/bench-261402-8                 90            170           +88.89%
BenchmarkSolution/bench-2629359-8                128           240           +87.50%
BenchmarkSolution/bench-24547482-8               128           240           +87.50%
BenchmarkSolution/bench-465093345-8              176           320           +81.82%
BenchmarkSolution/bench-8295678306-8             192           352           +83.33%
BenchmarkSolution/bench-85605547386-8            208           384           +84.62%
BenchmarkSolution/bench-528789347646-8           224           416           +85.71%
BenchmarkSolution/bench-6421242908142-8          240           448           +86.67%
BenchmarkSolution/bench-71771243324629-8         256           480           +87.50%
BenchmarkSolution/bench-704229078898971-8        256           480           +87.50%
BenchmarkSolution/bench-4046455185423042-8       272           512           +88.24%
BenchmarkSolution/bench-29461204712002348-8      352           640           +81.82%
BenchmarkSolution/bench-954492294769786752-8     352           640           +81.82%
```

### Version 1 to Version 3
```
benchmark                                        old ns/op      new ns/op     delta
Benchmark_hustle/2-8                             3.42           3.42          +0.00%
Benchmark_hustle/4-8                             257            93.9          -63.46%
Benchmark_hustle/8-8                             448            132           -70.54%
Benchmark_hustle/16-8                            844            213           -74.76%
Benchmark_hustle/32-8                            1823           362           -80.14%
Benchmark_hustle/64-8                            4701           719           -84.71%
Benchmark_hustle/128-8                           13495          1420          -89.48%
Benchmark_hustle/256-8                           43945          2816          -93.59%
Benchmark_hustle/512-8                           155347         5203          -96.65%
Benchmark_hustle/1024-8                          581838         9850          -98.31%
Benchmark_hustle/2048-8                          2308275        18861         -99.18%
Benchmark_hustle/4096-8                          9577424        42273         -99.56%
Benchmark_hustle/8192-8                          36272902       77438         -99.79%
Benchmark_hustle/16384-8                         147704701      168391        -99.89%
Benchmark_hustle/32768-8                         585325823      357075        -99.94%
Benchmark_hustle/65536-8                         2325228042     751634        -99.97%
BenchmarkSolution/bench-3-8                      61.2           60.9          -0.49%
BenchmarkSolution/bench-71-8                     140            142           +1.43%
BenchmarkSolution/bench-613-8                    437            283           -35.24%
BenchmarkSolution/bench-2029-8                   432            279           -35.42%
BenchmarkSolution/bench-74035-8                  613            342           -44.21%
BenchmarkSolution/bench-261402-8                 616            350           -43.18%
BenchmarkSolution/bench-2629359-8                717            407           -43.24%
BenchmarkSolution/bench-24547482-8               681            434           -36.27%
BenchmarkSolution/bench-465093345-8              904            503           -44.36%
BenchmarkSolution/bench-8295678306-8             985            503           -48.93%
BenchmarkSolution/bench-85605547386-8            1015           541           -46.70%
BenchmarkSolution/bench-528789347646-8           1089           575           -47.20%
BenchmarkSolution/bench-6421242908142-8          1114           595           -46.59%
BenchmarkSolution/bench-71771243324629-8         1175           629           -46.47%
BenchmarkSolution/bench-704229078898971-8        1168           615           -47.35%
BenchmarkSolution/bench-4046455185423042-8       1227           649           -47.11%
BenchmarkSolution/bench-29461204712002348-8      1512           748           -50.53%
BenchmarkSolution/bench-954492294769786752-8     1539           757           -50.81%

benchmark                                        old allocs     new allocs     delta
Benchmark_hustle/2-8                             0              0              +0.00%
Benchmark_hustle/4-8                             3              1              -66.67%
Benchmark_hustle/8-8                             4              1              -75.00%
Benchmark_hustle/16-8                            5              1              -80.00%
Benchmark_hustle/32-8                            6              1              -83.33%
Benchmark_hustle/64-8                            7              1              -85.71%
Benchmark_hustle/128-8                           8              1              -87.50%
Benchmark_hustle/256-8                           9              1              -88.89%
Benchmark_hustle/512-8                           10             1              -90.00%
Benchmark_hustle/1024-8                          11             1              -90.91%
Benchmark_hustle/2048-8                          14             1              -92.86%
Benchmark_hustle/4096-8                          16             1              -93.75%
Benchmark_hustle/8192-8                          19             1              -94.74%
Benchmark_hustle/16384-8                         22             1              -95.45%
Benchmark_hustle/32768-8                         25             1              -96.00%
Benchmark_hustle/65536-8                         28             1              -96.43%
BenchmarkSolution/bench-3-8                      1              1              +0.00%
BenchmarkSolution/bench-71-8                     2              2              +0.00%
BenchmarkSolution/bench-613-8                    6              4              -33.33%
BenchmarkSolution/bench-2029-8                   6              4              -33.33%
BenchmarkSolution/bench-74035-8                  7              4              -42.86%
BenchmarkSolution/bench-261402-8                 7              4              -42.86%
BenchmarkSolution/bench-2629359-8                7              4              -42.86%
BenchmarkSolution/bench-24547482-8               7              4              -42.86%
BenchmarkSolution/bench-465093345-8              8              4              -50.00%
BenchmarkSolution/bench-8295678306-8             8              4              -50.00%
BenchmarkSolution/bench-85605547386-8            8              4              -50.00%
BenchmarkSolution/bench-528789347646-8           8              4              -50.00%
BenchmarkSolution/bench-6421242908142-8          8              4              -50.00%
BenchmarkSolution/bench-71771243324629-8         8              4              -50.00%
BenchmarkSolution/bench-704229078898971-8        8              4              -50.00%
BenchmarkSolution/bench-4046455185423042-8       8              4              -50.00%
BenchmarkSolution/bench-29461204712002348-8      9              4              -55.56%
BenchmarkSolution/bench-954492294769786752-8     9              4              -55.56%

benchmark                                        old bytes     new bytes     delta
Benchmark_hustle/2-8                             0             0             +0.00%
Benchmark_hustle/4-8                             112           64            -42.86%
Benchmark_hustle/8-8                             240           128           -46.67%
Benchmark_hustle/16-8                            496           256           -48.39%
Benchmark_hustle/32-8                            1008          512           -49.21%
Benchmark_hustle/64-8                            2032          1024          -49.61%
Benchmark_hustle/128-8                           4080          2048          -49.80%
Benchmark_hustle/256-8                           8176          4096          -49.90%
Benchmark_hustle/512-8                           16368         8192          -49.95%
Benchmark_hustle/1024-8                          32752         16384         -49.98%
Benchmark_hustle/2048-8                          121456        32768         -73.02%
Benchmark_hustle/4096-8                          252528        65536         -74.05%
Benchmark_hustle/8192-8                          629360        131072        -79.17%
Benchmark_hustle/16384-8                         1383024       262147        -81.05%
Benchmark_hustle/32768-8                         2882160       524292        -81.81%
Benchmark_hustle/65536-8                         5831280       1048585       -82.02%
BenchmarkSolution/bench-3-8                      16            16            +0.00%
BenchmarkSolution/bench-71-8                     34            34            +0.00%
BenchmarkSolution/bench-613-8                    166           102           -38.55%
BenchmarkSolution/bench-2029-8                   166           102           -38.55%
BenchmarkSolution/bench-74035-8                  330           170           -48.48%
BenchmarkSolution/bench-261402-8                 330           170           -48.48%
BenchmarkSolution/bench-2629359-8                368           240           -34.78%
BenchmarkSolution/bench-24547482-8               368           240           -34.78%
BenchmarkSolution/bench-465093345-8              672           320           -52.38%
BenchmarkSolution/bench-8295678306-8             688           352           -48.84%
BenchmarkSolution/bench-85605547386-8            704           384           -45.45%
BenchmarkSolution/bench-528789347646-8           720           416           -42.22%
BenchmarkSolution/bench-6421242908142-8          736           448           -39.13%
BenchmarkSolution/bench-71771243324629-8         752           480           -36.17%
BenchmarkSolution/bench-704229078898971-8        752           480           -36.17%
BenchmarkSolution/bench-4046455185423042-8       768           512           -33.33%
BenchmarkSolution/bench-29461204712002348-8      1360          640           -52.94%
BenchmarkSolution/bench-954492294769786752-8     1360          640           -52.94%
```
