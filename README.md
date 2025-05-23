# gob-serialization

#### !! This repo is fork version of [gob-serialization by RSheremeta](https://github.com/RSheremeta/gob-serialization) to benchmark [bytedance/sonic](https://github.com/bytedance/sonic) !!
#### Benchmarking Gob vs other data formats like JSON, XML and YAML.

This repo is made just out of curiosity and has an experimental reasoning.

### Description
There are 4 structs: `Tiny`, `Medium`, `Big`, `Huge` which are being compared against all the defined formats above. 
And there is a *separate* struct to compare **Gob vs JSON only** - a struct called `ComplexAndHuge` which holds a slice of `type complexMap map[string]Huge`.

Each of them has a respective size and complexity.
And each of them builds the following targets to run encoding & decoding benchmarks against: 
- `Single` means a single struct - eg - `Tiny{}`
- `Slice` means a slice of given struct - eg - `[]Tiny`
- `Pointer Slice` means a slice of pointers to given struct - eg - `[]*Tiny`

### Usage
You can just look at and analyze the results I got while running on my machine by observing csv samples in **results_run/sample/** dir.

Also, feel free to play on your own by doing the following:

Run a corresponding **Make** command in terminal and observe the results in **results_run/** dir.
Otherwise run `make-all` in order to run everything. Warning: beware as it's a very CPU/RAM consuming operation.

### Results
**Software and Hardware:**
The experiment is run on:
- MacBook Pro 2021 15" with macOS Ventura 13.2.1 
- 32gb RAM
- Apple M1 Pro CPU
- Go 1.19.5

Here is the full analysis of mine about all the results run: [Analysis.md](https://github.com/RSheremeta/gob-serialization/blob/master/Analysis.md)

[**map_ptr_slice_decode.csv:**](https://github.com/RSheremeta/gob-serialization/blob/master/results_run/sample/map_ptr_slice_decode.csv)

```bash
goos: darwin
goarch: arm64
pkg: github.com/RSheremeta/gob-serialization
BenchmarkDecodePtrSliceComplexMap/type=GOB_struct_size=huge_complex_map-10         	      10	 646447171 ns/op	630836108 B/op	11072058 allocs/op
BenchmarkDecodePtrSliceComplexMap/type=JSON_struct_size=huge_complex_map-10        	      10	6356559012 ns/op	1610172119 B/op	18605410 allocs/op
PASS
ok  	github.com/RSheremeta/gob-serialization	85.338s
```

[**all.csv:**](https://github.com/RSheremeta/gob-serialization/blob/master/results_run/sample/all.csv)

```bash
goos: darwin
goarch: arm64
pkg: github.com/RSheremeta/gob-serialization
BenchmarkEncodeSingle/type=GOB_struct_size=tiny-10         	 1146312	      1037 ns/op	    1136 B/op	      20 allocs/op
BenchmarkEncodeSingle/type=JSON_struct_size=tiny-10        	11647294	       101.6 ns/op	      48 B/op	       1 allocs/op
BenchmarkEncodeSingle/type=XML_struct_size=tiny-10         	 1247809	       960.3 ns/op	    4560 B/op	       9 allocs/op
BenchmarkEncodeSingle/type=YAML_struct_size=tiny-10        	  404588	      2924 ns/op	    6776 B/op	      23 allocs/op
BenchmarkEncodeSingle/type=GOB_struct_size=medium-10       	  420466	      2786 ns/op	    1720 B/op	      38 allocs/op
BenchmarkEncodeSingle/type=JSON_struct_size=medium-10      	 1995039	       611.5 ns/op	     240 B/op	       2 allocs/op
BenchmarkEncodeSingle/type=XML_struct_size=medium-10       	  409347	      2731 ns/op	    5053 B/op	      14 allocs/op
BenchmarkEncodeSingle/type=YAML_struct_size=medium-10      	  132667	      8870 ns/op	   17112 B/op	      56 allocs/op
BenchmarkEncodeSingle/type=GOB_struct_size=big-10          	  227510	      5232 ns/op	    3072 B/op	      65 allocs/op
BenchmarkEncodeSingle/type=JSON_struct_size=big-10         	  461298	      2522 ns/op	     944 B/op	       6 allocs/op
BenchmarkEncodeSingle/type=XML_struct_size=big-10          	  127340	      9153 ns/op	    6034 B/op	      21 allocs/op
BenchmarkEncodeSingle/type=YAML_struct_size=big-10         	   33507	     34729 ns/op	   80288 B/op	     176 allocs/op
BenchmarkEncodeSingle/type=GOB_struct_size=huge-10         	     224	   5244342 ns/op	 7035920 B/op	   44617 allocs/op
BenchmarkEncodeSingle/type=JSON_struct_size=huge-10        	      62	  17295433 ns/op	 6592181 B/op	   39871 allocs/op
BenchmarkEncodeSingle/type=XML_struct_size=huge-10         	      14	  78271318 ns/op	19909926 B/op	   97568 allocs/op
BenchmarkEncodeSingle/type=YAML_struct_size=huge-10        	       4	 298569250 ns/op	845118380 B/op	 1215044 allocs/op
BenchmarkDecodeSingle/type=GOB_struct_size=tiny-10         	  132409	      9390 ns/op	    6760 B/op	     179 allocs/op
BenchmarkDecodeSingle/type=JSON_struct_size=tiny-10        	 2591954	       461.0 ns/op	     248 B/op	       6 allocs/op
BenchmarkDecodeSingle/type=XML_struct_size=tiny-10         	  627777	      1931 ns/op	    1480 B/op	      33 allocs/op
BenchmarkDecodeSingle/type=YAML_struct_size=tiny-10        	  263422	      4338 ns/op	    7280 B/op	      52 allocs/op
BenchmarkDecodeSingle/type=GOB_struct_size=medium-10       	   92902	     12639 ns/op	    8604 B/op	     235 allocs/op
BenchmarkDecodeSingle/type=JSON_struct_size=medium-10      	  646796	      1797 ns/op	     440 B/op	      14 allocs/op
BenchmarkDecodeSingle/type=XML_struct_size=medium-10       	  135865	      9034 ns/op	    3608 B/op	      92 allocs/op
BenchmarkDecodeSingle/type=YAML_struct_size=medium-10      	  101642	     11995 ns/op	   10724 B/op	     126 allocs/op
BenchmarkDecodeSingle/type=GOB_struct_size=big-10          	   69265	     17420 ns/op	   11557 B/op	     324 allocs/op
BenchmarkDecodeSingle/type=JSON_struct_size=big-10         	  194416	      6186 ns/op	    1080 B/op	      32 allocs/op
BenchmarkDecodeSingle/type=XML_struct_size=big-10          	   37994	     31857 ns/op	   10376 B/op	     280 allocs/op
BenchmarkDecodeSingle/type=YAML_struct_size=huge-10        	   29328	     40508 ns/op	   23081 B/op	     386 allocs/op
BenchmarkDecodeSingle/type=GOB_struct_size=huge-10         	     277	   4408610 ns/op	 4108664 B/op	   71653 allocs/op
BenchmarkDecodeSingle/type=JSON_struct_size=huge-10        	      27	  42904093 ns/op	12024463 B/op	  120170 allocs/op
BenchmarkDecodeSingle/type=XML_struct_size=huge-10         	       5	 224652383 ns/op	74270280 B/op	 1836139 allocs/op
BenchmarkDecodeSingle/type=YAML_struct_size=huge#01-10     	       4	 280739250 ns/op	121876510 B/op	 2425871 allocs/op
BenchmarkEncodePtrSlice/type=GOB_struct_size=tiny-10       	  766519	      1588 ns/op	    1392 B/op	      22 allocs/op
BenchmarkEncodePtrSlice/type=JSON_struct_size=tiny-10      	 3311300	       367.0 ns/op	     176 B/op	       1 allocs/op
BenchmarkEncodePtrSlice/type=XML_struct_size=tiny-10       	  523192	      2330 ns/op	    4816 B/op	       9 allocs/op
BenchmarkEncodePtrSlice/type=YAML_struct_size=tiny-10      	  122904	     11151 ns/op	   36392 B/op	      78 allocs/op
BenchmarkEncodePtrSlice/type=GOB_struct_size=medium-10     	  238899	      5326 ns/op	    4600 B/op	      62 allocs/op
BenchmarkEncodePtrSlice/type=JSON_struct_size=medium-10    	  213715	      5767 ns/op	    2272 B/op	      11 allocs/op
BenchmarkEncodePtrSlice/type=XML_struct_size=medium-10     	   46567	     26118 ns/op	    8773 B/op	      51 allocs/op
BenchmarkEncodePtrSlice/type=YAML_struct_size=medium-10    	   15430	     75262 ns/op	  171416 B/op	     494 allocs/op
BenchmarkEncodePtrSlice/type=GOB_struct_size=big-10        	   23326	     49705 ns/op	   68808 B/op	     571 allocs/op
BenchmarkEncodePtrSlice/type=JSON_struct_size=big-10       	   10000	    123179 ns/op	   44774 B/op	     251 allocs/op
BenchmarkEncodePtrSlice/type=XML_struct_size=big-10        	    2264	    538901 ns/op	  152597 B/op	     715 allocs/op
BenchmarkEncodePtrSlice/type=YAML_struct_size=big-10       	     799	   1538428 ns/op	 4682072 B/op	    8842 allocs/op
BenchmarkEncodePtrSlice/type=GOB_struct_size=huge-10       	       5	 249447842 ns/op	341030953 B/op	 2226117 allocs/op
BenchmarkEncodePtrSlice/type=JSON_struct_size=huge-10      	       2	 868415104 ns/op	329189260 B/op	 1993403 allocs/op
BenchmarkEncodePtrSlice/type=XML_struct_size=huge-10       	       1	3949217375 ns/op	1230326816 B/op	 4877230 allocs/op
BenchmarkEncodePtrSlice/type=YAML_struct_size=huge-10      	       1	17784455125 ns/op	38385133288 B/op	60748369 allocs/op
BenchmarkDecodePtrSlice/type=GOB_struct_size=tiny-10       	  105658	     11314 ns/op	    7724 B/op	     208 allocs/op
BenchmarkDecodePtrSlice/type=JSON_struct_size=tiny-10      	  521919	      2337 ns/op	     544 B/op	      20 allocs/op
BenchmarkDecodePtrSlice/type=XML_struct_size=tiny-10       	  540841	      2240 ns/op	    1536 B/op	      36 allocs/op
BenchmarkDecodePtrSlice/type=YAML_struct_size=tiny-10      	   69958	     17167 ns/op	   12592 B/op	     177 allocs/op
BenchmarkDecodePtrSlice/type=GOB_struct_size=medium-10     	   67569	     17878 ns/op	   11221 B/op	     310 allocs/op
BenchmarkDecodePtrSlice/type=JSON_struct_size=medium-10    	   69162	     17254 ns/op	    2608 B/op	      96 allocs/op
BenchmarkDecodePtrSlice/type=XML_struct_size=medium-10     	  128006	      9378 ns/op	    3664 B/op	      95 allocs/op
BenchmarkDecodePtrSlice/type=YAML_struct_size=medium-10    	   10000	    112053 ns/op	   52418 B/op	    1030 allocs/op
BenchmarkDecodePtrSlice/type=GOB_struct_size=big-10        	   16177	     74428 ns/op	   46335 B/op	    1179 allocs/op
BenchmarkDecodePtrSlice/type=JSON_struct_size=big-10       	    3897	    299794 ns/op	   38088 B/op	    1225 allocs/op
BenchmarkDecodePtrSlice/type=XML_struct_size=big-10        	   36422	     32854 ns/op	   10432 B/op	     283 allocs/op
BenchmarkDecodePtrSlice/type=YAML_struct_size=big-10       	     561	   2155220 ns/op	  846919 B/op	   17894 allocs/op
BenchmarkDecodePtrSlice/type=GOB_struct_size=huge-10       	      18	  65971963 ns/op	44990502 B/op	     802 allocs/op
BenchmarkDecodePtrSlice/type=JSON_struct_size=huge-10      	       1	1302005041 ns/op	   14008 B/op	      72 allocs/op
BenchmarkDecodePtrSlice/type=XML_struct_size=huge-10       	       8	 128212328 ns/op	48727516 B/op	 1374540 allocs/op
BenchmarkDecodePtrSlice/type=YAML_struct_size=huge-10      	       1	12199034416 ns/op	5480352464 B/op	96667450 allocs/op
BenchmarkEncodeSlice/type=GOB_struct_size=tiny-10          	  747290	      1678 ns/op	    1392 B/op	      22 allocs/op
BenchmarkEncodeSlice/type=JSON_struct_size=tiny-10         	 2159914	       558.1 ns/op	     320 B/op	       1 allocs/op
BenchmarkEncodeSlice/type=XML_struct_size=tiny-10          	  321912	      3843 ns/op	    5072 B/op	       9 allocs/op
BenchmarkEncodeSlice/type=YAML_struct_size=tiny-10         	   73617	     17534 ns/op	   37584 B/op	     129 allocs/op
BenchmarkEncodeSlice/type=GOB_struct_size=medium-10        	  197379	      6180 ns/op	    4600 B/op	      62 allocs/op
BenchmarkEncodeSlice/type=JSON_struct_size=medium-10       	  116884	     10030 ns/op	    4416 B/op	      21 allocs/op
BenchmarkEncodeSlice/type=XML_struct_size=medium-10        	   25900	     46562 ns/op	   18709 B/op	      72 allocs/op
BenchmarkEncodeSlice/type=YAML_struct_size=medium-10       	    9133	    143216 ns/op	  340952 B/op	     877 allocs/op
BenchmarkEncodeSlice/type=GOB_struct_size=big-10           	   18789	     64469 ns/op	   69400 B/op	     571 allocs/op
BenchmarkEncodeSlice/type=JSON_struct_size=big-10          	    6282	    197420 ns/op	   78944 B/op	     451 allocs/op
BenchmarkEncodeSlice/type=XML_struct_size=big-10           	    1371	    889912 ns/op	  298069 B/op	    1116 allocs/op
BenchmarkEncodeSlice/type=YAML_struct_size=big-10          	     487	   2359149 ns/op	 6520168 B/op	   13745 allocs/op
BenchmarkEncodeSlice/type=GOB_struct_size=huge-10          	       5	 246363042 ns/op	341030620 B/op	 2226116 allocs/op
BenchmarkEncodeSlice/type=JSON_struct_size=huge-10         	       2	 869657292 ns/op	597634568 B/op	 1993422 allocs/op
BenchmarkEncodeSlice/type=XML_struct_size=huge-10          	       1	3874216916 ns/op	1230326832 B/op	 4877230 allocs/op
BenchmarkEncodeSlice/type=YAML_struct_size=huge-10         	       1	16234461667 ns/op	38385152888 B/op	60749219 allocs/op
BenchmarkDecodeSlice/type=GOB_struct_size=tiny-10          	  101436	     11605 ns/op	    7812 B/op	     203 allocs/op
BenchmarkDecodeSlice/type=JSON_struct_size=tiny-10         	  282842	      4248 ns/op	    1176 B/op	      19 allocs/op
BenchmarkDecodeSlice/type=XML_struct_size=tiny-10          	  596376	      2057 ns/op	    1488 B/op	      32 allocs/op
BenchmarkDecodeSlice/type=YAML_struct_size=tiny-10         	   40621	     29642 ns/op	   18280 B/op	     278 allocs/op
BenchmarkDecodeSlice/type=GOB_struct_size=medium-10        	   66992	     17821 ns/op	   11973 B/op	     300 allocs/op
BenchmarkDecodeSlice/type=JSON_struct_size=medium-10       	   35984	     33182 ns/op	    9120 B/op	     100 allocs/op
BenchmarkDecodeSlice/type=XML_struct_size=medium-10        	  172459	      6964 ns/op	    2872 B/op	      66 allocs/op
BenchmarkDecodeSlice/type=YAML_struct_size=medium-10       	    5738	    208889 ns/op	   95046 B/op	    1791 allocs/op
BenchmarkDecodeSlice/type=GOB_struct_size=big-10           	   15175	     78775 ns/op	   58495 B/op	    1129 allocs/op
BenchmarkDecodeSlice/type=JSON_struct_size=big-10          	    2388	    497792 ns/op	  141000 B/op	    1379 allocs/op
BenchmarkDecodeSlice/type=XML_struct_size=big-10           	   62374	     19358 ns/op	    6200 B/op	     152 allocs/op
BenchmarkDecodeSlice/type=YAML_struct_size=big-10          	     349	   3426646 ns/op	 1381654 B/op	   27396 allocs/op
BenchmarkDecodeSlice/type=GOB_struct_size=huge-10          	       5	 209612733 ns/op	204711942 B/op	 3565079 allocs/op
BenchmarkDecodeSlice/type=JSON_struct_size=huge-10         	       1	2101034833 ns/op	601242440 B/op	 6008030 allocs/op
BenchmarkDecodeSlice/type=XML_struct_size=huge-10          	 1652984	       722.0 ns/op	     920 B/op	      15 allocs/op
BenchmarkDecodeSlice/type=YAML_struct_size=huge-10         	       1	14744372417 ns/op	6093520992 B/op	121293601 allocs/op
PASS
ok  	github.com/RSheremeta/gob-serialization	285.131s
```