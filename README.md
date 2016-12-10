## side by side gorm

Some tests to check gorm differences after refactoring

Example of result produced:
        | Test name | Allocs | Bytes | Duration  | Dif Allocs | Dif Bytes | Dif Duration |
        |:---------:|-------:|------:|----------:|-----------:|----------:|-------------:|
	|      0 OpenTestConnection    |   0    |   0   |     nothing.    |            |           |              |
	| original (same test) |   63    |   6400   |     nothing.    |            |           |              |
	| differences |       |      |         |3 less allocs | 256 less bytes |same time |
	|      1 RunNewMigration    |   60    |   6144   |     nothing.    |            |           |              |
	| original (same test) |   8321    |   1274424   |     4.5912627s    |            |           |              |
	| differences |       |      |         |487 MORE allocs |1782776 MORE bytes |took MORE with 508.0289ms |
	|      2 StringPrimaryKey    |   8808    |   3057200   |     5.0992916s    |            |           |              |
	| original (same test) |   643    |   36352   |     223.0127ms    |            |           |              |
	| differences |       |      |         |33 less allocs | 12136 less bytes |took MORE with 28.0017ms |
	|      3 SetTable    |   610    |   24216   |     251.0144ms    |            |           |              |
	| original (same test) |   19155    |   1579216   |     820.047ms    |            |           |              |
	| differences |       |      |         |308 MORE allocs |595664 less bytes |took less with 81.0047ms |
	|      4 ExceptionsWithInvalidSql    |   19463    |   983552   |     739.0423ms    |            |           |              |
	| original (same test) |   2332    |   1101352   |     1.0001ms    |            |           |              |
	| differences |       |      |         |940 less allocs | 1023496 less bytes |took less with 100ns |
	|      5 HasTable    |   1392    |   77856   |     1ms    |            |           |              |
	| original (same test) |   286    |   18480   |     124.0071ms    |            |           |              |
	| differences |       |      |         |7 less allocs | 7640 less bytes |took MORE with 7.0004ms |
	|      6 TableName    |   279    |   10840   |     131.0075ms    |            |           |              |
	| original (same test) |   162    |   22512   |     nothing.    |            |           |              |
	| differences |       |      |         |26 MORE allocs |9680 less bytes |took MORE with 1ms |
	|      7 NullValues    |   188    |   12832   |     1ms    |            |           |              |
	| original (same test) |   1883    |   481440   |     279.0159ms    |            |           |              |
	| differences |       |      |         |410 less allocs | 420512 less bytes |took less with 1ms |
	|      8 NullValuesWithFirstOrCreate    |   1473    |   60928   |     278.0159ms    |            |           |              |
	| original (same test) |   968    |   55784   |     132.0075ms    |            |           |              |
	| differences |       |      |         |225 MORE allocs |4384 MORE bytes |took MORE with 2.0002ms |
	|      9 Transaction    |   1193    |   60168   |     134.0077ms    |            |           |              |
	| original (same test) |   4258    |   629680   |     88.0051ms    |            |           |              |
	| differences |       |      |         |46 less allocs | 412032 less bytes |took less with 34.002ms |
	|      10 Row    |   4212    |   217648   |     54.0031ms    |            |           |              |
	| original (same test) |   2417    |   149608   |     230.0131ms    |            |           |              |
	| differences |       |      |         |31 less allocs | 25032 less bytes |took less with 2ms |
	|      11 Rows    |   2386    |   124576   |     228.0131ms    |            |           |              |
	| original (same test) |   2420    |   147232   |     238.0137ms    |            |           |              |
	| differences |       |      |         |16 less allocs | 22144 less bytes |took less with 26.0016ms |
	|      12 ScanRows    |   2404    |   125088   |     212.0121ms    |            |           |              |
	| original (same test) |   2535    |   154632   |     195.0112ms    |            |           |              |
	| differences |       |      |         |5 less allocs | 23552 less bytes |took MORE with 7.0003ms |
	|      13 Scan    |   2530    |   131080   |     202.0115ms    |            |           |              |
	| original (same test) |   2933    |   183968   |     213.0122ms    |            |           |              |
	| differences |       |      |         |185 less allocs | 41312 less bytes |took less with 23.0014ms |
	|      14 Raw    |   2748    |   142656   |     190.0108ms    |            |           |              |
	| original (same test) |   3144    |   194776   |     279.016ms    |            |           |              |
	| differences |       |      |         |205 less allocs | 40176 less bytes |took less with 21.0012ms |
	|      15 Group    |   2939    |   154600   |     258.0148ms    |            |           |              |
	| original (same test) |   163    |   6528   |     nothing.    |            |           |              |
	| differences |       |      |         |8 MORE allocs |576 less bytes |same time |
	|      16 Joins    |   171    |   5952   |     nothing.    |            |           |              |
	| original (same test) |   4194    |   274984   |     79.0045ms    |            |           |              |
	| differences |       |      |         |267 less allocs | 47176 less bytes |took MORE with 11.0006ms |
	|      17 JoinsWithSelect    |   3927    |   227808   |     90.0051ms    |            |           |              |
	| original (same test) |   1370    |   86368   |     88.0051ms    |            |           |              |
	| differences |       |      |         |144 less allocs | 29064 less bytes |took less with 18.0011ms |
	|      18 Having    |   1226    |   57304   |     70.004ms    |            |           |              |
	| original (same test) |   201    |   13432   |     nothing.    |            |           |              |
	| differences |       |      |         |82 less allocs | 7576 less bytes |took MORE with 1ms |
	|      19 TimeWithZone    |   119    |   5856   |     1ms    |            |           |              |
	| original (same test) |   3884    |   283424   |     146.0084ms    |            |           |              |
	| differences |       |      |         |166 MORE allocs |9392 less bytes |took less with 35.002ms |
	|      20 Hstore    |   4050    |   274032   |     111.0064ms    |            |           |              |
	| original (same test) |   31    |   1280   |     nothing.    |            |           |              |
	| differences |       |      |         |3 less allocs | 96 less bytes |same time |
	|      21 SetAndGet    |   28    |   1184   |     nothing.    |            |           |              |
	| original (same test) |   28    |   1680   |     nothing.    |            |           |              |
	| differences |       |      |         |4 less allocs | 416 less bytes |same time |
	|      22 CompatibilityMode    |   24    |   1264   |     nothing.    |            |           |              |
	| original (same test) |   529    |   35816   |     nothing.    |            |           |              |
	| differences |       |      |         |219 MORE allocs |16896 MORE bytes |same time |
	|      23 OpenExistingDB    |   748    |   52712   |     nothing.    |            |           |              |
	| original (same test) |   1111    |   67544   |     72.0041ms    |            |           |              |
	| differences |       |      |         |59 MORE allocs |5800 less bytes |took MORE with 5.0003ms |
	|      24 DdlErrors    |   1170    |   61744   |     77.0044ms    |            |           |              |
	| original (same test) |   570    |   411384   |     1.0001ms    |            |           |              |
	| differences |       |      |         |300 less allocs | 395232 less bytes |took less with 1.0001ms |
	|      25 OpenWithOneParameter    |   270    |   16152   |     nothing.    |            |           |              |
	| original (same test) |   29    |   1616   |     nothing.    |            |           |              |
	| differences |       |      |         |8 less allocs | 672 less bytes |same time |
	|      26 BelongsTo    |   21    |   944   |     nothing.    |            |           |              |
	| original (same test) |   11525    |   733720   |     785.0449ms    |            |           |              |
	| differences |       |      |         |924 less allocs | 162264 less bytes |took less with 211.0121ms |
	|      27 BelongsToOverrideForeignKey1    |   10601    |   571456   |     574.0328ms    |            |           |             |
	| original (same test) |   342    |   20200   |     nothing.    |            |           |              |
	| differences |       |      |         |8 MORE allocs |2984 less bytes |same time |
	|      28 BelongsToOverrideForeignKey2    |   350    |   17216   |     nothing.    |            |           |              |
	| original (same test) |   248    |   17608   |     nothing.    |            |           |              |
	| differences |       |      |         |30 MORE allocs |3936 less bytes |same time |
	|      29 HasOne    |   278    |   13672   |     nothing.    |            |           |              |
	| original (same test) |   15690    |   952832   |     774.0443ms    |            |           |              |
	| differences |       |      |         |157 less allocs | 110312 less bytes |took less with 177.0102ms |
	|      30 HasOneOverrideForeignKey1    |   15533    |   842520   |     597.0341ms    |            |           |              |
	| original (same test) |   276    |   18824   |     nothing.    |            |           |              |
	| differences |       |      |         |31 MORE allocs |1456 MORE bytes |took MORE with 1.0001ms |
	|      31 HasOneOverrideForeignKey2    |   307    |   20280   |     1.0001ms    |            |           |              |
	| original (same test) |   247    |   17544   |     nothing.    |            |           |              |
	| differences |       |      |         |24 MORE allocs |4128 less bytes |same time |
	|      32 HasMany    |   271    |   13416   |     nothing.    |            |           |              |
	| original (same test) |   12089    |   811696   |     951.0544ms    |            |           |              |
	| differences |       |      |         |521 less allocs | 163608 less bytes |took less with 226.013ms |
	|      33 HasManyOverrideForeignKey1    |   11568    |   648088   |     725.0414ms    |            |           |              |
	| original (same test) |   269    |   17680   |     nothing.    |            |           |              |
	| differences |       |      |         |31 MORE allocs |2800 less bytes |same time |
	|      34 HasManyOverrideForeignKey2    |   300    |   14880   |     nothing.    |            |           |              |
	| original (same test) |   245    |   18912   |     1.0001ms    |            |           |              |
	| differences |       |      |         |24 MORE allocs |4000 less bytes |took less with 1.0001ms |
	|      35 ManyToMany    |   269    |   14912   |     nothing.    |            |           |              |
	| original (same test) |   27575    |   1716168   |     1.7601006s    |            |           |              |
	| differences |       |      |         |2359 less allocs | 366456 less bytes |took less with 8.0004ms |
	|      36 Related    |   25216    |   1349712   |     1.7521002s    |            |           |              |
	| original (same test) |   7414    |   439384   |     163.0093ms    |            |           |              |
	| differences |       |      |         |352 MORE allocs |32960 less bytes |took less with 74.0042ms |
	|      37 ForeignKey    |   7766    |   406424   |     89.0051ms    |            |           |              |
	| original (same test) |   61    |   6976   |     nothing.    |            |           |              |
	| differences |       |      |         |7 less allocs | 2224 less bytes |same time |
	|      38 LongForeignKey    |   54    |   4752   |     nothing.    |            |           |              |
	| original (same test) |   27    |   1136   |     nothing.    |            |           |              |
	| differences |       |      |         |3 less allocs | 64 less bytes |same time |
	|      39 LongForeignKeyWithShortDest    |   24    |   1072   |     nothing.    |            |           |              |
	| original (same test) |   27    |   1152   |     nothing.    |            |           |              |
	| differences |       |      |         |3 less allocs | 64 less bytes |same time |
	|      40 HasManyChildrenWithOneStruct    |   24    |   1088   |     nothing.    |            |           |              |
	| original (same test) |   668    |   43896   |     81.0046ms    |            |           |              |
	| differences |       |      |         |34 MORE allocs |13192 less bytes |took less with 14.0008ms |
	|      41 RunCallbacks    |   702    |   30704   |     67.0038ms    |            |           |              |
	| original (same test) |   2778    |   149832   |     190.0109ms    |            |           |              |
	| differences |       |      |         |23 MORE allocs |16056 less bytes |took less with 17.001ms |
	|      42 CallbacksWithErrors    |   2801    |   133776   |     173.0099ms    |            |           |              |
	| original (same test) |   8806    |   4309160   |     205.0118ms    |            |           |              |
	| differences |       |      |         |3475 less allocs | 4064096 less bytes |took less with 100ns |
	|      43 Create    |   5331    |   245064   |     205.0117ms    |            |           |              |
	| original (same test) |   2115    |   111600   |     132.0076ms    |            |           |              |
	| differences |       |      |         |508 MORE allocs |28968 MORE bytes |took MORE with 9.0004ms |
	|      44 CreateWithAutoIncrement    |   2623    |   140568   |     141.008ms    |            |           |              |
	| original (same test) |   34    |   1872   |     nothing.    |            |           |              |
	| differences |       |      |         |3 less allocs | 96 less bytes |same time |
	|      45 CreateWithNoGORMPrimayKey    |   31    |   1776   |     nothing.    |            |           |              |
	| original (same test) |   280    |   18664   |     77.0044ms    |            |           |              |
	| differences |       |      |         |5 less allocs | 6624 less bytes |took less with 18.001ms |
	|      46 CreateWithNoStdPrimaryKeyAndDefaultValues    |   275    |   12040   |     59.0034ms    |            |           |              |
		| original (same test) |   1188    |   75960   |     155.0089ms    |            |           |              |
		| differences |       |      |         |94 less allocs | 25856 less bytes |took less with 21.0013ms |
		|      47 AnonymousScanner    |   1094    |   50104   |     134.0076ms    |            |           |              |
		| original (same test) |   1098    |   63080   |     74.0042ms    |            |           |              |
		| differences |       |      |         |54 MORE allocs |3368 less bytes |took less with 7.0004ms |
		|      48 AnonymousField    |   1152    |   59712   |     67.0038ms    |            |           |              |
		| original (same test) |   1622    |   96760   |     78.0045ms    |            |           |              |
		| differences |       |      |         |32 MORE allocs |12000 less bytes |took MORE with 135.0077ms |
		|      49 SelectWithCreate    |   1654    |   84760   |     213.0122ms    |            |           |              |
		| original (same test) |   3252    |   205480   |     188.0108ms    |            |           |              |
		| differences |       |      |         |150 less allocs | 54912 less bytes |took less with 16.001ms |
		|      50 OmitWithCreate    |   3102    |   150568   |     172.0098ms    |            |           |              |
		| original (same test) |   3425    |   217512   |     205.0118ms    |            |           |              |
		| differences |       |      |         |138 less allocs | 49592 less bytes |took less with 60.0036ms |
		|      51 CustomizeColumn    |   3287    |   167920   |     145.0082ms    |            |           |              |
		| original (same test) |   862    |   59408   |     365.0209ms    |            |           |              |
		| differences |       |      |         |42 MORE allocs |17296 less bytes |took MORE with 4.0002ms |
		|      52 CustomColumnAndIgnoredFieldClash    |   904    |   42112   |     369.0211ms    |            |           |              |
		| original (same test) |   161    |   10568   |     152.0087ms    |            |           |              |
		| differences |       |      |         |1 MORE allocs |3408 MORE bytes |took MORE with 4.0002ms |
		|      53 ManyToManyWithCustomizedColumn    |   162    |   13976   |     156.0089ms    |            |           |              |
		| original (same test) |   2081    |   138648   |     593.0339ms    |            |           |              |
		| differences |       |      |         |392 less allocs | 61416 less bytes |took MORE with 65.0037ms |
		|      54 OneToOneWithCustomizedColumn    |   1689    |   77232   |     658.0376ms    |            |           |              |
		| original (same test) |   1564    |   98328   |     593.0339ms    |            |           |              |
		| differences |       |      |         |13 MORE allocs |23176 less bytes |took MORE with 55.0032ms |
		|      55 OneToManyWithCustomizedColumn    |   1577    |   75152   |     648.0371ms    |            |           |              |
		| original (same test) |   3510    |   217376   |     533.0305ms    |            |           |              |
		| differences |       |      |         |143 less allocs | 49808 less bytes |took less with 67.0038ms |
		|      56 HasOneWithPartialCustomizedColumn    |   3367    |   167568   |     466.0267ms    |            |           |              |
		| original (same test) |   2447    |   148440   |     592.0339ms    |            |           |              |
		| differences |       |      |         |129 less allocs | 34984 less bytes |took less with 110.0064ms |
		|      57 BelongsToWithPartialCustomizedColumn    |   2318    |   113456   |     482.0275ms    |            |           |              |
		| original (same test) |   2692    |   166704   |     575.0328ms    |            |           |              |
		| differences |       |      |         |141 less allocs | 39416 less bytes |took less with 116.0065ms |
		|      58 Delete    |   2551    |   127288   |     459.0263ms    |            |           |              |
		| original (same test) |   2181    |   126480   |     235.0134ms    |            |           |              |
		| differences |       |      |         |113 MORE allocs |6432 less bytes |took less with 2ms |
		|      59 InlineDelete    |   2294    |   120048   |     233.0134ms    |            |           |              |
		| original (same test) |   2304    |   136888   |     271.0155ms    |            |           |              |
		| differences |       |      |         |13 MORE allocs |15224 less bytes |took MORE with 19.0011ms |
		|      60 SoftDelete    |   2317    |   121664   |     290.0166ms    |            |           |              |
		| original (same test) |   1276    |   76400   |     278.0159ms    |            |           |              |
		| differences |       |      |         |233 less allocs | 33736 less bytes |took less with 47.0027ms |
		|      61 PrefixColumnNameForEmbeddedStruct    |   1043    |   42664   |     231.0132ms    |            |           |              |
		| original (same test) |   426    |   31240   |     1ms    |            |           |              |
		| differences |       |      |         |10 MORE allocs |11240 less bytes |took MORE with 100ns |
		|      62 SaveAndQueryEmbeddedStruct    |   436    |   20000   |     1.0001ms    |            |           |              |
		| original (same test) |   1369    |   71392   |     208.0119ms    |            |           |              |
		| differences |       |      |         |81 less allocs | 19120 less bytes |took less with 7.0004ms |
		|      63 CalculateField    |   1288    |   52272   |     201.0115ms    |            |           |              |
		| original (same test) |   451    |   31624   |     nothing.    |            |           |              |
		| differences |       |      |         |44 MORE allocs |6368 less bytes |same time |
		|      64 JoinTable    |   495    |   25256   |     nothing.    |            |           |              |
		| original (same test) |   4325    |   283712   |     501.0286ms    |            |           |              |
		| differences |       |      |         |258 less allocs | 79192 less bytes |took less with 26.0014ms |
		|      65 Indexes    |   4067    |   204520   |     475.0272ms    |            |           |              |
		| original (same test) |   9820    |   3086928   |     957.0548ms    |            |           |              |
		| differences |       |      |         |1588 less allocs | 2691840 less bytes |took less with 100ns |
		|      66 AutoMigration    |   8232    |   395088   |     957.0547ms    |            |           |              |
		| original (same test) |   1346    |   59760   |     388.0222ms    |            |           |              |
		| differences |       |      |         |3 MORE allocs |6648 less bytes |took MORE with 50.0028ms |
		|      67 MultipleIndexes    |   1349    |   53112   |     438.025ms    |            |           |              |
		| original (same test) |   3076    |   936888   |     799.0457ms    |            |           |              |
		| differences |       |      |         |637 less allocs | 835064 less bytes |took MORE with 131.0075ms |
		|      68 ManyToManyWithMultiPrimaryKeys    |   2439    |   101824   |     930.0532ms    |            |           |              |
		| original (same test) |   26    |   1216   |     nothing.    |            |           |              |
		| differences |       |      |         |3 less allocs | 96 less bytes |same time |
		|      69 ManyToManyWithCustomizedForeignKeys    |   23    |   1120   |     nothing.    |            |           |              |
		| original (same test) |   26    |   1232   |     nothing.    |            |           |              |
		| differences |       |      |         |3 less allocs | 96 less bytes |same time |
		|      70 ManyToManyWithCustomizedForeignKeys2    |   23    |   1136   |     nothing.    |            |           |              |
		| original (same test) |   26    |   1232   |     nothing.    |            |           |              |
		| differences |       |      |         |3 less allocs | 96 less bytes |same time |
		|      71 PointerFields    |   23    |   1136   |     nothing.    |            |           |              |
		| original (same test) |   2564    |   156776   |     351.0201ms    |            |           |              |
		| differences |       |      |         |517 less allocs | 73840 less bytes |took MORE with 35.002ms |
		|      72 Polymorphic    |   2047    |   82936   |     386.0221ms    |            |           |              |
		| original (same test) |   23326    |   1611144   |     1.4120807s    |            |           |              |
		| differences |       |      |         |6284 less allocs | 740816 less bytes |took less with 332.0189ms |
		|      73 NamedPolymorphic    |   17042    |   870328   |     1.0800618s    |            |           |              |
		| original (same test) |   16186    |   1130120   |     1.1120636s    |            |           |              |
		| differences |       |      |         |4686 less allocs | 503672 less bytes |took less with 216.0123ms |
		|      74 Preload    |   11500    |   626448   |     896.0513ms    |            |           |              |
		| original (same test) |   22746    |   1310672   |     403.023ms    |            |           |              |
		| differences |       |      |         |434 less allocs | 239928 less bytes |took less with 4.0002ms |
		|      75 NestedPreload1    |   22312    |   1070744   |     399.0228ms    |            |           |              |
		| original (same test) |   1960    |   124664   |     648.037ms    |            |           |              |
		| differences |       |      |         |198 less allocs | 8672 less bytes |took less with 73.0041ms |
		|      76 NestedPreload2    |   1762    |   115992   |     575.0329ms    |            |           |              |
		| original (same test) |   2390    |   144616   |     590.0337ms    |            |           |              |
		| differences |       |      |         |172 less allocs | 50848 less bytes |took less with 72.004ms |
		|      77 NestedPreload3    |   2218    |   93768   |     518.0297ms    |            |           |              |
		| original (same test) |   2147    |   132240   |     533.0305ms    |            |           |              |
		| differences |       |      |         |167 less allocs | 41704 less bytes |took MORE with 15.0008ms |
		|      78 NestedPreload4    |   1980    |   90536   |     548.0313ms    |            |           |              |
		| original (same test) |   1917    |   120704   |     490.028ms    |            |           |              |
		| differences |       |      |         |156 less allocs | 45936 less bytes |took MORE with 38.0022ms |
		|      79 NestedPreload5    |   1761    |   74768   |     528.0302ms    |            |           |              |
		| original (same test) |   2376    |   143128   |     612.035ms    |            |           |              |
		| differences |       |      |         |180 less allocs | 52560 less bytes |took less with 33.0018ms |
		|      80 NestedPreload6    |   2196    |   90568   |     579.0332ms    |            |           |              |
		| original (same test) |   3688    |   220760   |     583.0333ms    |            |           |              |
		| differences |       |      |         |215 less allocs | 80800 less bytes |took less with 38.0021ms |
		|      81 NestedPreload7    |   3473    |   139960   |     545.0312ms    |            |           |              |
		| original (same test) |   3292    |   191512   |     671.0384ms    |            |           |              |
		| differences |       |      |         |214 less allocs | 63456 less bytes |took MORE with 29.0016ms |
		|      82 NestedPreload8    |   3078    |   128056   |     700.04ms    |            |           |              |
		| original (same test) |   2808    |   166952   |     654.0374ms    |            |           |              |
		| differences |       |      |         |194 less allocs | 60896 less bytes |took less with 17.0009ms |
		|      83 NestedPreload9    |   2614    |   106056   |     637.0365ms    |            |           |              |
		| original (same test) |   6276    |   382864   |     998.0571ms    |            |           |              |
		| differences |       |      |         |290 less allocs | 131936 less bytes |took less with 55.0032ms |
		|      84 NestedPreload10    |   5986    |   250928   |     943.0539ms    |            |           |              |
		| original (same test) |   2265    |   133520   |     785.0449ms    |            |           |              |
		| differences |       |      |         |155 less allocs | 28728 less bytes |took less with 40.0023ms |
		|      85 NestedPreload11    |   2110    |   104792   |     745.0426ms    |            |           |              |
		| original (same test) |   2017    |   122456   |     598.0342ms    |            |           |              |
		| differences |       |      |         |207 less allocs | 44464 less bytes |took MORE with 62.0035ms |
		|      86 NestedPreload12    |   1810    |   77992   |     660.0377ms    |            |           |              |
		| original (same test) |   2691    |   159072   |     658.0376ms    |            |           |              |
		| differences |       |      |         |202 less allocs | 42064 less bytes |took less with 26.0014ms |
		|      87 ManyToManyPreloadWithMultiPrimaryKeys    |   2489    |   117008   |     632.0362ms    |            |           |              |
		| original (same test) |   26    |   1232   |     nothing.    |            |           |              |
		| differences |       |      |         |2 less allocs | 13472 MORE bytes |same time |
		|      88 ManyToManyPreloadForNestedPointer    |   24    |   14704   |     nothing.    |            |           |              |
		| original (same test) |   8606    |   544960   |     718.0411ms    |            |           |              |
		| differences |       |      |         |2115 less allocs | 259384 less bytes |took less with 42.0025ms |
		|      89 NestedManyToManyPreload    |   6491    |   285576   |     676.0386ms    |            |           |              |
		| original (same test) |   5376    |   365592   |     982.0562ms    |            |           |              |
		| differences |       |      |         |1216 less allocs | 182936 less bytes |took less with 71.0041ms |
		|      90 NestedManyToManyPreload2    |   4160    |   182656   |     911.0521ms    |            |           |              |
		| original (same test) |   3320    |   216312   |     732.0419ms    |            |           |              |
		| differences |       |      |         |673 less allocs | 95784 less bytes |took MORE with 39.0022ms |
		|      91 NestedManyToManyPreload3    |   2647    |   120528   |     771.0441ms    |            |           |              |
		| original (same test) |   5402    |   350264   |     946.0541ms    |            |           |              |
		| differences |       |      |         |963 less allocs | 159640 less bytes |took MORE with 13.0007ms |
		|      92 NestedManyToManyPreload3ForStruct    |   4439    |   190624   |     959.0548ms    |            |           |              |
		| original (same test) |   5631    |   360416   |     935.0535ms    |            |           |              |
		| differences |       |      |         |978 less allocs | 161664 less bytes |took less with 33.0019ms |
		|      93 NestedManyToManyPreload4    |   4653    |   198752   |     902.0516ms    |            |           |              |
		| original (same test) |   4211    |   293880   |     915.0523ms    |            |           |              |
		| differences |       |      |         |740 less allocs | 142728 less bytes |took MORE with 132.0076ms |
		|      94 ManyToManyPreloadForPointer    |   3471    |   151152   |     1.0470599s    |            |           |              |
		| original (same test) |   6615    |   430360   |     612.035ms    |            |           |              |
		| differences |       |      |         |1779 less allocs | 205496 less bytes |took less with 24.0014ms |
		|      95 NilPointerSlice    |   4836    |   224864   |     588.0336ms    |            |           |              |
		| original (same test) |   1998    |   120584   |     654.0374ms    |            |           |              |
		| differences |       |      |         |150 less allocs | 44000 less bytes |took less with 48.0027ms |
		|      96 NilPointerSlice2    |   1848    |   76584   |     606.0347ms    |            |           |              |
		| original (same test) |   1848    |   125312   |     946.0541ms    |            |           |              |
		| differences |       |      |         |132 less allocs | 50272 less bytes |took MORE with 49.0028ms |
		|      97 PrefixedPreloadDuplication    |   1716    |   75040   |     995.0569ms    |            |           |              |
		| original (same test) |   4305    |   253552   |     1.258072s    |            |           |              |
		| differences |       |      |         |279 less allocs | 88808 less bytes |took less with 10.0006ms |
		|      98 FirstAndLast    |   4026    |   164744   |     1.2480714s    |            |           |              |
		| original (same test) |   3908    |   216664   |     191.0109ms    |            |           |              |
		| differences |       |      |         |703 MORE allocs |27488 MORE bytes |took MORE with 104.006ms |
		|      99 FirstAndLastWithNoStdPrimaryKey    |   4611    |   244152   |     295.0169ms    |            |           |              |
		| original (same test) |   1583    |   96368   |     141.0081ms    |            |           |              |
		| differences |       |      |         |33 less allocs | 24144 less bytes |took less with 1.0001ms |
		|      100 UIntPrimaryKey    |   1550    |   72224   |     140.008ms    |            |           |              |
		| original (same test) |   487    |   28888   |     nothing.    |            |           |              |
		| differences |       |      |         |79 MORE allocs |696 less bytes |took MORE with 1.0001ms |
		|      101 StringPrimaryKeyForNumericValueStartingWithZero    |   566    |   28192   |     1.0001ms    |            |           |              |
		| original (same test) |   923    |   431584   |     1.0001ms    |            |           |              |
		| differences |       |      |         |427 less allocs | 410368 less bytes |took less with 100ns |
		|      102 FindAsSliceOfPointers    |   496    |   21216   |     1ms    |            |           |              |
		| original (same test) |   15633    |   894960   |     74.0042ms    |            |           |              |
		| differences |       |      |         |4947 MORE allocs |393664 MORE bytes |took MORE with 13.0008ms |
		|      103 SearchWithPlainSQL    |   20580    |   1288624   |     87.005ms    |            |           |              |
		| original (same test) |   10108    |   654392   |     216.0123ms    |            |           |              |
		| differences |       |      |         |281 MORE allocs |3120 MORE bytes |took MORE with 7.0005ms |
		|      104 SearchWithStruct    |   10389    |   657512   |     223.0128ms    |            |           |              |
		| original (same test) |   6314    |   353128   |     245.014ms    |            |           |              |
		| differences |       |      |         |1311 MORE allocs |82008 MORE bytes |took less with 21.0012ms |
		|      105 SearchWithMap    |   7625    |   435136   |     224.0128ms    |            |           |              |
		| original (same test) |   5295    |   309864   |     267.0152ms    |            |           |              |
		| differences |       |      |         |854 MORE allocs |28192 MORE bytes |took MORE with 104.0061ms |
		|      106 SearchWithEmptyChain    |   6149    |   338056   |     371.0213ms    |            |           |              |
		| original (same test) |   3985    |   234496   |     248.0142ms    |            |           |              |
		| differences |       |      |         |193 MORE allocs |8496 less bytes |took less with 13.0007ms |
		|      107 Select    |   4178    |   226000   |     235.0135ms    |            |           |              |
		| original (same test) |   1007    |   58240   |     82.0047ms    |            |           |              |
		| differences |       |      |         |48 MORE allocs |2504 less bytes |took MORE with 1ms |
		|      108 OrderAndPluck    |   1055    |   55736   |     83.0047ms    |            |           |              |
		| original (same test) |   12062    |   702584   |     224.0128ms    |            |           |              |
		| differences |       |      |         |3384 MORE allocs |255016 MORE bytes |took less with 21.0011ms |
		|      109 Limit    |   15446    |   957600   |     203.0117ms    |            |           |              |
		| original (same test) |   15889    |   1043304   |     524.03ms    |            |           |              |
		| differences |       |      |         |4397 MORE allocs |314984 MORE bytes |took less with 181.0104ms |
		|      110 Offset    |   20286    |   1358288   |     343.0196ms    |            |           |              |
		| original (same test) |   68919    |   4359736   |     1.5900909s    |            |           |              |
		| differences |       |      |         |19334 MORE allocs |1437280 MORE bytes |took less with 198.0113ms |
		|      111 Or    |   88253    |   5797016   |     1.3920796s    |            |           |              |
		| original (same test) |   2437    |   148728   |     278.0159ms    |            |           |              |
		| differences |       |      |         |74 MORE allocs |4648 MORE bytes |took less with 12.0007ms |
		|      112 Count    |   2511    |   153376   |     266.0152ms    |            |           |              |
		| original (same test) |   3422    |   208616   |     258.0148ms    |            |           |              |
		| differences |       |      |         |171 less allocs | 31616 less bytes |took MORE with 25.0014ms |
		|      113 Not    |   3251    |   177000   |     283.0162ms    |            |           |              |
		| original (same test) |   21568    |   1530112   |     539.0308ms    |            |           |              |
		| differences |       |      |         |542 MORE allocs |340104 less bytes |took MORE with 141.0081ms |
		|      114 FillSmallerStruct    |   22110    |   1190008   |     680.0389ms    |            |           |              |
		| original (same test) |   959    |   56192   |     81.0046ms    |            |           |              |
		| differences |       |      |         |46 less allocs | 13528 less bytes |took MORE with 2.0002ms |
		|      115 FindOrInitialize    |   913    |   42664   |     83.0048ms    |            |           |              |
		| original (same test) |   5245    |   276592   |     77.0044ms    |            |           |              |
		| differences |       |      |         |1953 MORE allocs |129936 MORE bytes |took less with 15.0009ms |
		|      116 FindOrCreate    |   7198    |   406528   |     62.0035ms    |            |           |              |
		| original (same test) |   10433    |   1332712   |     497.0284ms    |            |           |              |
		| differences |       |      |         |1572 MORE allocs |683856 less bytes |took less with 35.002ms |
		|      117 SelectWithEscapedFieldName    |   12005    |   648856   |     462.0264ms    |            |           |              |
		| original (same test) |   2054    |   122344   |     222.0127ms    |            |           |              |
		| differences |       |      |         |210 MORE allocs |4904 less bytes |took less with 36.002ms |
		|      118 SelectWithVariables    |   2264    |   117440   |     186.0107ms    |            |           |              |
		| original (same test) |   656    |   39496   |     80.0046ms    |            |           |              |
		| differences |       |      |         |29 MORE allocs |5256 less bytes |took MORE with 6.0003ms |
		|      119  fix #1214 : FirstAndLastWithRaw    |   685    |   34240   |     86.0049ms    |            |           |              |
		| original (same test) |   2551    |   148336   |     186.0107ms    |            |           |              |
		| differences |       |      |         |113 MORE allocs |12624 less bytes |took less with 5.0004ms |
		|      120 ScannableSlices    |   2664    |   135712   |     181.0103ms    |            |           |              |
		| original (same test) |   1282    |   67464   |     68.0039ms    |            |           |              |
		| differences |       |      |         |470 MORE allocs |16216 MORE bytes |took less with 2.0001ms |
		|      121 Scopes    |   1752    |   83680   |     66.0038ms    |            |           |              |
		| original (same test) |   3464    |   209360   |     222.0127ms    |            |           |              |
		| differences |       |      |         |183 MORE allocs |4256 less bytes |took less with 45.0026ms |
		|      122 Update    |   3647    |   205104   |     177.0101ms    |            |           |              |
		| original (same test) |   6498    |   346080   |     574.0329ms    |            |           |              |
		| differences |       |      |         |357 MORE allocs |22336 less bytes |took less with 75.0044ms |
		|      123 UpdateWithNoStdPrimaryKeyAndDefaultValues    |   6855    |   323744   |     499.0285ms    |            |           |              |
		| original (same test) |   3013    |   169672   |     437.025ms    |            |           |              |
		| differences |       |      |         |20 less allocs | 35176 less bytes |took MORE with 67.0038ms |
		|      124 Updates    |   2993    |   134496   |     504.0288ms    |            |           |              |
		| original (same test) |   4602    |   241576   |     307.0176ms    |            |           |              |
		| differences |       |      |         |185 MORE allocs |23896 less bytes |took less with 44.0026ms |
		|      125 UpdateColumn    |   4787    |   217680   |     263.015ms    |            |           |              |
		| original (same test) |   2850    |   145944   |     306.0175ms    |            |           |              |
		| differences |       |      |         |414 MORE allocs |3760 MORE bytes |took less with 8.0005ms |
		|      126 SelectWithUpdate    |   3264    |   149704   |     298.017ms    |            |           |              |
		| original (same test) |   7252    |   440944   |     305.0174ms    |            |           |              |
		| differences |       |      |         |76 less allocs | 98992 less bytes |took less with 49.0027ms |
		|      127 SelectWithUpdateWithMap    |   7176    |   341952   |     256.0147ms    |            |           |              |
		| original (same test) |   7289    |   443360   |     270.0155ms    |            |           |              |
		| differences |       |      |         |68 less allocs | 96176 less bytes |took MORE with 35.0019ms |
		|      128 OmitWithUpdate    |   7221    |   347184   |     305.0174ms    |            |           |              |
		| original (same test) |   6125    |   372728   |     263.015ms    |            |           |              |
		| differences |       |      |         |85 less allocs | 80184 less bytes |took less with 10.0005ms |
		|      129 OmitWithUpdateWithMap    |   6040    |   292544   |     253.0145ms    |            |           |              |
		| original (same test) |   5982    |   367328   |     219.0125ms    |            |           |              |
		| differences |       |      |         |104 less allocs | 79824 less bytes |took less with 14.0008ms |
		|      130 SelectWithUpdateColumn    |   5878    |   287504   |     205.0117ms    |            |           |              |
		| original (same test) |   4074    |   240760   |     185.0106ms    |            |           |              |
		| differences |       |      |         |320 MORE allocs |26888 less bytes |took MORE with 23.0013ms |
		|      131 OmitWithUpdateColumn    |   4394    |   213872   |     208.0119ms    |            |           |              |
		| original (same test) |   4068    |   239720   |     210.012ms    |            |           |              |
		| differences |       |      |         |327 MORE allocs |25912 less bytes |took MORE with 30.0017ms |
		|      132 UpdateColumnsSkipsAssociations    |   4395    |   213808   |     240.0137ms    |            |           |              |
		| original (same test) |   4045    |   238168   |     251.0143ms    |            |           |              |
		| differences |       |      |         |283 MORE allocs |30136 less bytes |took MORE with 39.0023ms |
		|      133 UpdatesWithBlankValues    |   4328    |   208032   |     290.0166ms    |            |           |              |
		| original (same test) |   1126    |   58936   |     155.0089ms    |            |           |              |
		| differences |       |      |         |164 MORE allocs |3048 MORE bytes |took MORE with 53.003ms |
		|      134 UpdatesTableWithIgnoredValues    |   1290    |   61984   |     208.0119ms    |            |           |              |
		| original (same test) |   529    |   27968   |     154.0088ms    |            |           |              |
		| differences |       |      |         |93 less allocs | 11160 less bytes |took less with 4.0002ms |
		|      135 UpdateDecodeVirtualAttributes    |   436    |   16808   |     150.0086ms    |            |           |              |
		| original (same test) |   930    |   51088   |     166.0094ms    |            |           |              |
		| differences |       |      |         |117 MORE allocs |3304 MORE bytes |took less with 25.0013ms |
		|      136 ToDBNameGenerateFriendlyName    |   1047    |   54392   |     141.0081ms    |            |           |              |
		| original (same test) |   124    |   5232   |     nothing.    |            |           |              |
		| differences |       |      |         |3 less allocs | 96 less bytes |same time |
		|      137 SkipSaveAssociation    |   121    |   5136   |     nothing.    |            |           |              |
		| original (same test) |   1366    |   74760   |     468.0268ms    |            |           |              |
		| differences |       |      |         |42 less allocs | 18728 less bytes |took less with 3.0002ms |
		|     TOTAL (original)    |   610944    |   49274112   |     51.0389194s    |            |           |              |
		|      TOTAL (new)    |   618071    |   35972776   |     49.7268438s    |            |           |              |
