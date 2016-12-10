## side by side gorm

Some tests to check gorm differences after refactoring

Example (test run on 10th of December 2016) of result produced:

| Test name | Allocs | Bytes | Duration  | Dif Allocs | Dif Bytes | Dif Duration |
| :-------: | -----: | ----: | --------: | ---------: | --------: | -----------: |
| *OpenTestConnection* | 59 | 6064 | 1ms | | | |
| OpenTestConnection | 59 | 4624 | nothing. | | | |
| diffs | | | | :zzz: | :snail: 1440 | :snail: 1ms |
| *RunNewMigration* | 8807 | 3056816 | 4.5782619s | | | |
| RunMigration | 8321 | 1274568 | 4.825276s | | | |
| diffs | | | | :snail: 486 | :snail: 1782248 | :zap: 247.0141ms |
| *StringPrimaryKey* | 609 | 24136 | 293.0167ms | | | |
| StringPrimaryKey | 642 | 36272 | 323.0185ms | | | |
| diffs | | | | :zap: 33 | :zap: 12136 | :zap: 30.0018ms |
| *SetTable* | 19450 | 983888 | 847.0485ms | | | |
| SetTable | 19151 | 1579696 | 964.0551ms | | | |
| diffs | | | | :snail: 299 | :zap: 595808 | :zap: 117.0066ms |
| *ExceptionsWithInvalidSql* | 1385 | 78320 | 1.0001ms | | | |
| ExceptionsWithInvalidSql | 2340 | 1101976 | 1ms | | | |
| diffs | | | | :zap: 955 | :zap: 1023656 | :snail: 100ns |
| *HasTable* | 278 | 10760 | 142.0081ms | | | |
| HasTable | 285 | 18400 | 162.0093ms | | | |
| diffs | | | | :zap: 7 | :zap: 7640 | :zap: 20.0012ms |
| *TableName* | 187 | 12752 | nothing. | | | |
| TableName | 161 | 22432 | nothing. | | | |
| diffs | | | | :snail: 26 | :zap: 9680 | :zzz: |
| *NullValues* | 1478 | 61088 | 321.0184ms | | | |
| NullValues | 1877 | 480784 | 372.0212ms | | | |
| diffs | | | | :zap: 399 | :zap: 419696 | :zap: 51.0028ms |
| *NullValuesWithFirstOrCreate* | 1184 | 59800 | 132.0075ms | | | |
| NullValuesWithFirstOrCreate | 967 | 55704 | 165.0095ms | | | |
| diffs | | | | :snail: 217 | :snail: 4096 | :zap: 33.002ms |
| *Transaction* | 4214 | 218816 | 68.0039ms | | | |
| Transaction | 4262 | 630336 | 82.0047ms | | | |
| diffs | | | | :zap: 48 | :zap: 411520 | :zap: 14.0008ms |
| *Row* | 2394 | 126728 | 201.0115ms | | | |
| Row | 2404 | 147360 | 250.0143ms | | | |
| diffs | | | | :zap: 10 | :zap: 20632 | :zap: 49.0028ms |
| *Rows* | 2403 | 125008 | 185.0106ms | | | |
| Rows | 2419 | 147216 | 248.0142ms | | | |
| diffs | | | | :zap: 16 | :zap: 22208 | :zap: 63.0036ms |
| *ScanRows* | 2529 | 131000 | 233.0133ms | | | |
| ScanRows | 2534 | 154488 | 284.0163ms | | | |
| diffs | | | | :zap: 5 | :zap: 23488 | :zap: 51.003ms |
| *Scan* | 2747 | 142576 | 235.0134ms | | | |
| Scan | 2932 | 183824 | 257.0147ms | | | |
| diffs | | | | :zap: 185 | :zap: 41248 | :zap: 22.0013ms |
| *Raw* | 2939 | 154728 | 283.0162ms | | | |
| Raw | 3138 | 194304 | 348.0199ms | | | |
| diffs | | | | :zap: 199 | :zap: 39576 | :zap: 65.0037ms |
| *Group* | 170 | 5872 | nothing. | | | |
| Group | 161 | 6240 | nothing. | | | |
| diffs | | | | :snail: 9 | :zap: 368 | :zzz: |
| *Joins* | 3925 | 227904 | 80.0045ms | | | |
| Joins | 4195 | 275192 | 89.0051ms | | | |
| diffs | | | | :zap: 270 | :zap: 47288 | :zap: 9.0006ms |
| *JoinsWithSelect* | 1225 | 57224 | 85.0049ms | | | |
| JoinsWithSelect | 1369 | 86096 | 106.006ms | | | |
| diffs | | | | :zap: 144 | :zap: 28872 | :zap: 21.0011ms |
| *Having* | 118 | 5776 | nothing. | | | |
| Having | 200 | 13352 | nothing. | | | |
| diffs | | | | :zap: 82 | :zap: 7576 | :zzz: |
| *TimeWithZone* | 4051 | 274224 | 137.0078ms | | | |
| TimeWithZone | 3880 | 282912 | 150.0086ms | | | |
| diffs | | | | :snail: 171 | :zap: 8688 | :zap: 13.0008ms |
| *Hstore* | 27 | 1104 | nothing. | | | |
| Hstore | 30 | 1200 | nothing. | | | |
| diffs | | | | :zap: 3 | :zap: 96 | :zzz: |
| *SetAndGet* | 23 | 1184 | nothing. | | | |
| SetAndGet | 27 | 1600 | nothing. | | | |
| diffs | | | | :zap: 4 | :zap: 416 | :zzz: |
| *CompatibilityMode* | 747 | 52968 | 1ms | | | |
| CompatibilityMode | 528 | 35736 | nothing. | | | |
| diffs | | | | :snail: 219 | :snail: 17232 | :snail: 1ms |
| *OpenExistingDB* | 1169 | 61664 | 60.0035ms | | | |
| OpenExistingDB | 1109 | 67320 | 81.0046ms | | | |
| diffs | | | | :snail: 60 | :zap: 5656 | :zap: 21.0011ms |
| *DdlErrors* | 268 | 16008 | nothing. | | | |
| DdlErrors | 567 | 410280 | 1.0001ms | | | |
| diffs | | | | :zap: 299 | :zap: 394272 | :zap: 1.0001ms |
| *OpenWithOneParameter* | 20 | 864 | nothing. | | | |
| OpenWithOneParameter | 23 | 976 | nothing. | | | |
| diffs | | | | :zap: 3 | :zap: 112 | :zzz: |
| *BelongsTo* | 10610 | 572352 | 590.0337ms | | | |
| BelongsTo | 11526 | 733992 | 657.0376ms | | | |
| diffs | | | | :zap: 916 | :zap: 161640 | :zap: 67.0039ms |
| *BelongsToOverrideForeignKey1* | 347 | 16640 | 1.0001ms | | | |
| BelongsToOverrideForeignKey1 | 341 | 20120 | nothing. | | | |
| diffs | | | | :snail: 6 | :zap: 3480 | :snail: 1.0001ms |
| *BelongsToOverrideForeignKey2* | 278 | 13800 | nothing. | | | |
| BelongsToOverrideForeignKey2 | 247 | 17528 | nothing. | | | |
| diffs | | | | :snail: 31 | :zap: 3728 | :zzz: |
| *HasOne* | 15530 | 842712 | 675.0386ms | | | |
| HasOne | 15690 | 952848 | 774.0442ms | | | |
| diffs | | | | :zap: 160 | :zap: 110136 | :zap: 99.0056ms |
| *HasOneOverrideForeignKey1* | 305 | 19992 | 1.0001ms | | | |
| HasOneOverrideForeignKey1 | 273 | 18248 | nothing. | | | |
| diffs | | | | :snail: 32 | :snail: 1744 | :snail: 1.0001ms |
| *HasOneOverrideForeignKey2* | 270 | 13336 | nothing. | | | |
| HasOneOverrideForeignKey2 | 246 | 17464 | nothing. | | | |
| diffs | | | | :snail: 24 | :zap: 4128 | :zzz: |
| *HasMany* | 11567 | 647880 | 767.0439ms | | | |
| Many | 12088 | 811232 | 850.0486ms | | | |
| diffs | | | | :zap: 521 | :zap: 163352 | :zap: 83.0047ms |
| *HasManyOverrideForeignKey1* | 300 | 15008 | nothing. | | | |
| HasManyOverrideForeignKey1 | 268 | 17600 | 1.0001ms | | | |
| diffs | | | | :snail: 32 | :zap: 2592 | :zap: 1.0001ms |
| *HasManyOverrideForeignKey2* | 268 | 14736 | nothing. | | | |
| HasManyOverrideForeignKey2 | 243 | 18688 | nothing. | | | |
| diffs | | | | :snail: 25 | :zap: 3952 | :zzz: |
| *ManyToMany* | 25215 | 1349472 | 1.9121094s | | | |
| ManyToMany | 27580 | 1716280 | 1.9721128s | | | |
| diffs | | | | :zap: 2365 | :zap: 366808 | :zap: 60.0034ms |
| *Related* | 7763 | 405992 | 104.0059ms | | | |
| Related | 7410 | 438744 | 97.0056ms | | | |
| diffs | | | | :snail: 353 | :zap: 32752 | :snail: 7.0003ms |
| *ForeignKey* | 53 | 4672 | nothing. | | | |
| ForeignKey | 61 | 7104 | 1ms | | | |
| diffs | | | | :zap: 8 | :zap: 2432 | :zap: 1ms |
| *LongForeignKey* | 23 | 992 | nothing. | | | |
| LongForeignKey | 26 | 1056 | nothing. | | | |
| diffs | | | | :zap: 3 | :zap: 64 | :zzz: |
| *LongForeignKeyWithShortDest* | 23 | 1008 | nothing. | | | |
| LongForeignKeyWithShortDest | 26 | 1072 | nothing. | | | |
| diffs | | | | :zap: 3 | :zap: 64 | :zzz: |
| *HasManyChildrenWithOneStruct* | 701 | 30624 | 83.0048ms | | | |
| HasManyChildrenWithOneStruct | 666 | 43608 | 105.006ms | | | |
| diffs | | | | :snail: 35 | :zap: 12984 | :zap: 22.0012ms |
| *RunCallbacks* | 2796 | 132816 | 208.0119ms | | | |
| RunCallbacks | 2777 | 150360 | 224.0128ms | | | |
| diffs | | | | :snail: 19 | :zap: 17544 | :zap: 16.0009ms |
| *CallbacksWithErrors* | 5319 | 242072 | 235.0134ms | | | |
| CallbacksWithErrors | 8793 | 4305736 | 276.0158ms | | | |
| diffs | | | | :zap: 3474 | :zap: 4063664 | :zap: 41.0024ms |
| *Create* | 2622 | 140408 | 141.0081ms | | | |
| Create | 2115 | 111552 | 173.0099ms | | | |
| diffs | | | | :snail: 507 | :snail: 28856 | :zap: 32.0018ms |
| *CreateWithAutoIncrement* | 31 | 1744 | nothing. | | | |
| CreateWithAutoIncrement | 33 | 1792 | nothing. | | | |
| diffs | | | | :zap: 2 | :zap: 48 | :zzz: |
| *CreateWithNoGORMPrimayKey* | 274 | 11960 | 72.0042ms | | | |
| CreateWithNoGORMPrimayKey | 279 | 18584 | 140.008ms | | | |
| diffs | | | | :zap: 5 | :zap: 6624 | :zap: 68.0038ms |
| *CreateWithNoStdPrimaryKeyAndDefaultValues* | 1093 | 49992 | 153.0087ms | | | |
| CreateWithNoStdPrimaryKeyAndDefaultValues | 1187 | 75880 | 182.0104ms | | | |
| diffs | | | | :zap: 94 | :zap: 25888 | :zap: 29.0017ms |
| *AnonymousScanner* | 1152 | 59840 | 70.004ms | | | |
| AnonymousScanner | 1098 | 63144 | 78.0045ms | | | |
| diffs | | | | :snail: 54 | :zap: 3304 | :zap: 8.0005ms |
| *AnonymousField* | 1653 | 84680 | 70.004ms | | | |
| AnonymousField | 1620 | 96472 | 87.005ms | | | |
| diffs | | | | :snail: 33 | :zap: 11792 | :zap: 17.001ms |
| *SelectWithCreate* | 3102 | 150312 | 147.0084ms | | | |
| SelectWithCreate | 3253 | 205336 | 147.0084ms | | | |
| diffs | | | | :zap: 151 | :zap: 55024 | :zzz: |
| *OmitWithCreate* | 3285 | 167696 | 194.0111ms | | | |
| OmitWithCreate | 3426 | 217784 | 155.0089ms | | | |
| diffs | | | | :zap: 141 | :zap: 50088 | :snail: 39.0022ms |
| *CustomizeColumn* | 903 | 42032 | 334.0191ms | | | |
| CustomizeColumn | 861 | 59328 | 363.0208ms | | | |
| diffs | | | | :snail: 42 | :zap: 17296 | :zap: 29.0017ms |
| *CustomColumnAndIgnoredFieldClash* | 161 | 13896 | 142.0081ms | | | |
| CustomColumnAndIgnoredFieldClash | 160 | 10488 | 191.0109ms | | | |
| diffs | | | | :snail: 1 | :snail: 3408 | :zap: 49.0028ms |
| *ManyToManyWithCustomizedColumn* | 1688 | 77184 | 680.0389ms | | | |
| ManyToManyWithCustomizedColumn | 2080 | 138568 | 695.0398ms | | | |
| diffs | | | | :zap: 392 | :zap: 61384 | :zap: 15.0009ms |
| *OneToOneWithCustomizedColumn* | 1574 | 74752 | 776.0444ms | | | |
| OneToOneWithCustomizedColumn | 1563 | 98248 | 729.0417ms | | | |
| diffs | | | | :snail: 11 | :zap: 23496 | :snail: 47.0027ms |
| *OneToManyWithCustomizedColumn* | 3369 | 167840 | 621.0356ms | | | |
| OneToManyWithCustomizedColumn | 3509 | 217264 | 650.0371ms | | | |
| diffs | | | | :zap: 140 | :zap: 49424 | :zap: 29.0015ms |
| *HasOneWithPartialCustomizedColumn* | 2317 | 113376 | 538.0308ms | | | |
| HasOneWithPartialCustomizedColumn | 2446 | 148360 | 595.034ms | | | |
| diffs | | | | :zap: 129 | :zap: 34984 | :zap: 57.0032ms |
| *BelongsToWithPartialCustomizedColumn* | 2550 | 127208 | 576.0329ms | | | |
| BelongsToWithPartialCustomizedColumn | 2691 | 166624 | 665.0381ms | | | |
| diffs | | | | :zap: 141 | :zap: 39416 | :zap: 89.0052ms |
| *Delete* | 2292 | 119856 | 216.0123ms | | | |
| Delete | 2179 | 126256 | 239.0137ms | | | |
| diffs | | | | :snail: 113 | :zap: 6400 | :zap: 23.0014ms |
| *InlineDelete* | 2317 | 121840 | 310.0177ms | | | |
| InlineDelete | 2312 | 138624 | 424.0243ms | | | |
| diffs | | | | :snail: 5 | :zap: 16784 | :zap: 114.0066ms |
| *SoftDelete* | 1052 | 43704 | 309.0177ms | | | |
| SoftDelete | 1276 | 76400 | 314.0179ms | | | |
| diffs | | | | :zap: 224 | :zap: 32696 | :zap: 5.0002ms |
| *PrefixColumnNameForEmbeddedStruct* | 434 | 19712 | 1.0001ms | | | |
| PrefixColumnNameForEmbeddedStruct | 425 | 31160 | 1.0001ms | | | |
| diffs | | | | :snail: 9 | :zap: 11448 | :zzz: |
| *SaveAndQueryEmbeddedStruct* | 1287 | 52192 | 236.0135ms | | | |
| SaveAndQueryEmbeddedStruct | 1368 | 71312 | 264.0151ms | | | |
| diffs | | | | :zap: 81 | :zap: 19120 | :zap: 28.0016ms |
| *CalculateField* | 494 | 25016 | 1ms | | | |
| CalculateField | 452 | 31864 | nothing. | | | |
| diffs | | | | :snail: 42 | :zap: 6848 | :snail: 1ms |
| *JoinTable* | 4056 | 202656 | 491.0281ms | | | |
| JoinTable | 4324 | 283632 | 540.0309ms | | | |
| diffs | | | | :zap: 268 | :zap: 80976 | :zap: 49.0028ms |
| *Indexes* | 8227 | 394608 | 846.0484ms | | | |
| Indexes | 9818 | 3086720 | 1.0040574s | | | |
| diffs | | | | :zap: 1591 | :zap: 2692112 | :zap: 158.009ms |
| *AutoMigration* | 1347 | 52680 | 414.0237ms | | | |
| AutoMigration | 1345 | 59680 | 437.025ms | | | |
| diffs | | | | :snail: 2 | :zap: 7000 | :zap: 23.0013ms |
| *MultipleIndexes* | 2440 | 102112 | 916.0524ms | | | |
| MultipleIndexes | 3082 | 938464 | 969.0554ms | | | |
| diffs | | | | :zap: 642 | :zap: 836352 | :zap: 53.003ms |
| *ManyToManyWithMultiPrimaryKeys* | 22 | 1040 | nothing. | | | |
| ManyToManyWithMultiPrimaryKeys | 25 | 1136 | nothing. | | | |
| diffs | | | | :zap: 3 | :zap: 96 | :zzz: |
| *ManyToManyWithCustomizedForeignKeys* | 22 | 1056 | nothing. | | | |
| ManyToManyWithCustomizedForeignKeys | 25 | 1152 | nothing. | | | |
| diffs | | | | :zap: 3 | :zap: 96 | :zzz: |
| *ManyToManyWithCustomizedForeignKeys2* | 22 | 1056 | nothing. | | | |
| ManyToManyWithCustomizedForeignKeys2 | 25 | 1152 | nothing. | | | |
| diffs | | | | :zap: 3 | :zap: 96 | :zzz: |
| *PointerFields* | 2046 | 82856 | 472.027ms | | | |
| PointerFields | 2562 | 156488 | 512.0292ms | | | |
| diffs | | | | :zap: 516 | :zap: 73632 | :zap: 40.0022ms |
| *Polymorphic* | 17033 | 868176 | 1.2030689s | | | |
| Polymorphic | 23324 | 1610984 | 1.5410881s | | | |
| diffs | | | | :zap: 6291 | :zap: 742808 | :zap: 338.0192ms |
| *NamedPolymorphic* | 11500 | 626720 | 1.1990685s | | | |
| NamedPolymorphic | 16185 | 1130040 | 1.1340649s | | | |
| diffs | | | | :zap: 4685 | :zap: 503320 | :snail: 65.0036ms |
| *Preload* | 22326 | 1073664 | 396.0227ms | | | |
| Preload | 22741 | 1309768 | 503.0287ms | | | |
| diffs | | | | :zap: 415 | :zap: 236104 | :zap: 107.006ms |
| *NestedPreload1* | 1759 | 115832 | 524.03ms | | | |
| NestedPreload1 | 1959 | 124680 | 655.0375ms | | | |
| diffs | | | | :zap: 200 | :zap: 8848 | :zap: 131.0075ms |
| *NestedPreload2* | 2217 | 93832 | 638.0365ms | | | |
| NestedPreload2 | 2391 | 144888 | 657.0376ms | | | |
| diffs | | | | :zap: 174 | :zap: 51056 | :zap: 19.0011ms |
| *NestedPreload3* | 1978 | 90344 | 658.0377ms | | | |
| NestedPreload3 | 2145 | 132048 | 651.0372ms | | | |
| diffs | | | | :zap: 167 | :zap: 41704 | :snail: 7.0005ms |
| *NestedPreload4* | 1757 | 74256 | 591.0338ms | | | |
| NestedPreload4 | 1916 | 120624 | 580.0332ms | | | |
| diffs | | | | :zap: 159 | :zap: 46368 | :snail: 11.0006ms |
| *NestedPreload5* | 2196 | 90632 | 670.0383ms | | | |
| NestedPreload5 | 2375 | 143048 | 708.0405ms | | | |
| diffs | | | | :zap: 179 | :zap: 52416 | :zap: 38.0022ms |
| *NestedPreload6* | 3475 | 140520 | 733.0419ms | | | |
| NestedPreload6 | 3688 | 220824 | 742.0425ms | | | |
| diffs | | | | :zap: 213 | :zap: 80304 | :zap: 9.0006ms |
| *NestedPreload7* | 3079 | 128376 | 796.0456ms | | | |
| NestedPreload7 | 3290 | 191320 | 770.044ms | | | |
| diffs | | | | :zap: 211 | :zap: 62944 | :snail: 26.0016ms |
| *NestedPreload8* | 2614 | 106088 | 680.0389ms | | | |
| NestedPreload8 | 2806 | 166664 | 811.0464ms | | | |
| diffs | | | | :zap: 192 | :zap: 60576 | :zap: 131.0075ms |
| *NestedPreload9* | 5989 | 251808 | 909.052ms | | | |
| NestedPreload9 | 6275 | 383376 | 1.1170639s | | | |
| diffs | | | | :zap: 286 | :zap: 131568 | :zap: 208.0119ms |
| *NestedPreload10* | 2109 | 104712 | 815.0466ms | | | |
| NestedPreload10 | 2261 | 133040 | 874.05ms | | | |
| diffs | | | | :zap: 152 | :zap: 28328 | :zap: 59.0034ms |
| *NestedPreload11* | 1811 | 77960 | 658.0376ms | | | |
| NestedPreload11 | 2017 | 122568 | 656.0375ms | | | |
| diffs | | | | :zap: 206 | :zap: 44608 | :snail: 2.0001ms |
| *NestedPreload12* | 2490 | 117248 | 742.0425ms | | | |
| NestedPreload12 | 2689 | 158800 | 820.0469ms | | | |
| diffs | | | | :zap: 199 | :zap: 41552 | :zap: 78.0044ms |
| *ManyToManyPreloadWithMultiPrimaryKeys* | 24 | 14832 | nothing. | | | |
| ManyToManyPreloadWithMultiPrimaryKeys | 25 | 1152 | nothing. | | | |
| diffs | | | | :zap: 1 | :snail: 13680 | :zzz: |
| *ManyToManyPreloadForNestedPointer* | 6492 | 285816 | 900.0515ms | | | |
| ManyToManyPreloadForNestedPointer | 8601 | 544544 | 934.0535ms | | | |
| diffs | | | | :zap: 2109 | :zap: 258728 | :zap: 34.002ms |
| *NestedManyToManyPreload* | 4157 | 182256 | 1.0290588s | | | |
| NestedManyToManyPreload | 5377 | 365928 | 961.055ms | | | |
| diffs | | | | :zap: 1220 | :zap: 183672 | :snail: 68.0038ms |
| *NestedManyToManyPreload2* | 2644 | 120192 | 774.0443ms | | | |
| NestedManyToManyPreload2 | 3318 | 215816 | 758.0434ms | | | |
| diffs | | | | :zap: 674 | :zap: 95624 | :snail: 16.0009ms |
| *NestedManyToManyPreload3* | 4438 | 190368 | 922.0527ms | | | |
| NestedManyToManyPreload3 | 5402 | 350248 | 1.0530602s | | | |
| diffs | | | | :zap: 964 | :zap: 159880 | :zap: 131.0075ms |
| *NestedManyToManyPreload3ForStruct* | 4652 | 198912 | 914.0523ms | | | |
| NestedManyToManyPreload3ForStruct | 5630 | 360448 | 994.0569ms | | | |
| diffs | | | | :zap: 978 | :zap: 161536 | :zap: 80.0046ms |
| *NestedManyToManyPreload4* | 3468 | 150752 | 1.0530602s | | | |
| NestedManyToManyPreload4 | 4211 | 293912 | 1.1970685s | | | |
| diffs | | | | :zap: 743 | :zap: 143160 | :zap: 144.0083ms |
| *ManyToManyPreloadForPointer* | 4836 | 224928 | 681.0389ms | | | |
| ManyToManyPreloadForPointer | 6613 | 430072 | 820.0469ms | | | |
| diffs | | | | :zap: 1777 | :zap: 205144 | :zap: 139.008ms |
| *NilPointerSlice* | 1847 | 76504 | 632.0362ms | | | |
| NilPointerSlice | 1996 | 120296 | 791.0452ms | | | |
| diffs | | | | :zap: 149 | :zap: 43792 | :zap: 159.009ms |
| *NilPointerSlice2* | 1714 | 74848 | 978.056ms | | | |
| NilPointerSlice2 | 1846 | 125120 | 947.0541ms | | | |
| diffs | | | | :zap: 132 | :zap: 50272 | :snail: 31.0019ms |
| *PrefixedPreloadDuplication* | 4026 | 165016 | 1.2130694s | | | |
| PrefixedPreloadDuplication | 4304 | 253392 | 1.3870793s | | | |
| diffs | | | | :zap: 278 | :zap: 88376 | :zap: 174.0099ms |
| *FirstAndLast* | 4607 | 243768 | 205.0117ms | | | |
| FirstAndLast | 3908 | 216792 | 228.0131ms | | | |
| diffs | | | | :snail: 699 | :snail: 26976 | :zap: 23.0014ms |
| *FirstAndLastWithNoStdPrimaryKey* | 1549 | 72208 | 136.0078ms | | | |
| FirstAndLastWithNoStdPrimaryKey | 1583 | 96368 | 189.0108ms | | | |
| diffs | | | | :zap: 34 | :zap: 24160 | :zap: 53.003ms |
| *UIntPrimaryKey* | 565 | 28144 | nothing. | | | |
| UIntPrimaryKey | 485 | 28600 | nothing. | | | |
| diffs | | | | :snail: 80 | :zap: 456 | :zzz: |
| *StringPrimaryKeyForNumericValueStartingWithZero* | 496 | 21008 | 1ms | | | |
| StringPrimaryKeyForNumericValueStartingWithZero | 921 | 431088 | nothing. | | | |
| diffs | | | | :zap: 425 | :zap: 410080 | :snail: 1ms |
| *FindAsSliceOfPointers* | 20579 | 1288576 | 84.0048ms | | | |
| FindAsSliceOfPointers | 15623 | 892560 | 93.0054ms | | | |
| diffs | | | | :snail: 4956 | :snail: 396016 | :zap: 9.0006ms |
| *SearchWithPlainSQL* | 10386 | 656952 | 229.0131ms | | | |
| SearchWithPlainSQL | 10102 | 653432 | 303.0173ms | | | |
| diffs | | | | :snail: 284 | :snail: 3520 | :zap: 74.0042ms |
| *SearchWithStruct* | 7621 | 434432 | 254.0145ms | | | |
| SearchWithStruct | 6311 | 352856 | 290.0166ms | | | |
| diffs | | | | :snail: 1310 | :snail: 81576 | :zap: 36.0021ms |
| *SearchWithMap* | 6152 | 338584 | 321.0184ms | | | |
| SearchWithMap | 5295 | 309928 | 350.02ms | | | |
| diffs | | | | :snail: 857 | :snail: 28656 | :zap: 29.0016ms |
| *SearchWithEmptyChain* | 4182 | 226248 | 197.0112ms | | | |
| SearchWithEmptyChain | 3982 | 234000 | 236.0135ms | | | |
| diffs | | | | :snail: 200 | :zap: 7752 | :zap: 39.0023ms |
| *Select* | 1054 | 55656 | 66.0038ms | | | |
| Select | 1006 | 58160 | 83.0048ms | | | |
| diffs | | | | :snail: 48 | :zap: 2504 | :zap: 17.001ms |
| *OrderAndPluck* | 15445 | 957472 | 204.0117ms | | | |
| OrderAndPluck | 12059 | 702152 | 231.0132ms | | | |
| diffs | | | | :snail: 3386 | :snail: 255320 | :zap: 27.0015ms |
| *Limit* | 20278 | 1356640 | 381.0218ms | | | |
| Limit | 15890 | 1043320 | 403.0231ms | | | |
| diffs | | | | :snail: 4388 | :snail: 313320 | :zap: 22.0013ms |
| *Offset* | 88244 | 5794952 | 1.6370936s | | | |
| Offset | 68915 | 4357416 | 1.7991029s | | | |
| diffs | | | | :snail: 19329 | :snail: 1437536 | :zap: 162.0093ms |
| *Or* | 2510 | 153360 | 227.013ms | | | |
| Or | 2439 | 149208 | 226.0129ms | | | |
| diffs | | | | :snail: 71 | :snail: 4152 | :snail: 1.0001ms |
| *Count* | 3249 | 176872 | 249.0142ms | | | |
| Count | 3421 | 208536 | 259.0149ms | | | |
| diffs | | | | :zap: 172 | :zap: 31664 | :zap: 10.0007ms |
| *Not* | 22104 | 1188136 | 598.0342ms | | | |
| Not | 21568 | 1530912 | 530.0303ms | | | |
| diffs | | | | :snail: 536 | :zap: 342776 | :snail: 68.0039ms |
| *FillSmallerStruct* | 912 | 42584 | 64.0036ms | | | |
| FillSmallerStruct | 958 | 56112 | 91.0053ms | | | |
| diffs | | | | :zap: 46 | :zap: 13528 | :zap: 27.0017ms |
| *FindOrInitialize* | 7202 | 407864 | 63.0036ms | | | |
| FindOrInitialize | 5244 | 276768 | 71.004ms | | | |
| diffs | | | | :snail: 1958 | :snail: 131096 | :zap: 8.0004ms |
| *FindOrCreate* | 11999 | 646960 | 435.0249ms | | | |
| FindOrCreate | 10429 | 1332024 | 538.0307ms | | | |
| diffs | | | | :snail: 1570 | :zap: 685064 | :zap: 103.0058ms |
| *SelectWithEscapedFieldName* | 2255 | 117600 | 216.0124ms | | | |
| SelectWithEscapedFieldName | 2053 | 122264 | 276.0158ms | | | |
| diffs | | | | :snail: 202 | :zap: 4664 | :zap: 60.0034ms |
| *SelectWithVariables* | 685 | 34368 | 73.0042ms | | | |
| SelectWithVariables | 655 | 39416 | 92.0053ms | | | |
| diffs | | | | :snail: 30 | :zap: 5048 | :zap: 19.0011ms |
| *FirstAndLastWithRaw* | 2662 | 135696 | 222.0127ms | | | |
| FirstAndLastWithRaw | 2547 | 147728 | 212.0121ms | | | |
| diffs | | | | :snail: 115 | :zap: 12032 | :snail: 10.0006ms |
| *ScannableSlices* | 2570 | 125312 | 74.0043ms | | | |
| ScannableSlices | 655 | 32936 | 91.0052ms | | | |
| diffs | | | | :snail: 1915 | :snail: 92376 | :zap: 17.0009ms |
| *Scopes* | 3644 | 204480 | 225.0129ms | | | |
| Scopes | 3462 | 208816 | 336.0192ms | | | |
| diffs | | | | :snail: 182 | :zap: 4336 | :zap: 111.0063ms |
| *Update* | 6852 | 323168 | 521.0298ms | | | |
| Update | 6501 | 346704 | 624.0357ms | | | |
| diffs | | | | :snail: 351 | :zap: 23536 | :zap: 103.0059ms |
| *UpdateWithNoStdPrimaryKeyAndDefaultValues* | 2992 | 134448 | 518.0297ms | | | |
| UpdateWithNoStdPrimaryKeyAndDefaultValues | 3012 | 169384 | 548.0313ms | | | |
| diffs | | | | :zap: 20 | :zap: 34936 | :zap: 30.0016ms |
| *Updates* | 4788 | 218080 | 320.0183ms | | | |
| Updates | 4601 | 241544 | 431.0247ms | | | |
| diffs | | | | :snail: 187 | :zap: 23464 | :zap: 111.0064ms |
| *UpdateColumn* | 3262 | 149480 | 379.0216ms | | | |
| UpdateColumn | 2848 | 145720 | 404.0231ms | | | |
| diffs | | | | :snail: 414 | :snail: 3760 | :zap: 25.0015ms |
| *SelectWithUpdate* | 7178 | 342592 | 330.0189ms | | | |
| SelectWithUpdate | 7269 | 443608 | 339.0194ms | | | |
| diffs | | | | :zap: 91 | :zap: 101016 | :zap: 9.0005ms |
| *SelectWithUpdateWithMap* | 7212 | 345272 | 329.0188ms | | | |
| SelectWithUpdateWithMap | 7287 | 442960 | 299.0171ms | | | |
| diffs | | | | :zap: 75 | :zap: 97688 | :snail: 30.0017ms |
| *OmitWithUpdate* | 6040 | 292976 | 283.0162ms | | | |
| OmitWithUpdate | 6123 | 372504 | 272.0156ms | | | |
| diffs | | | | :zap: 83 | :zap: 79528 | :snail: 11.0006ms |
| *OmitWithUpdateWithMap* | 5879 | 287776 | 203.0116ms | | | |
| OmitWithUpdateWithMap | 5980 | 366624 | 189.0108ms | | | |
| diffs | | | | :zap: 101 | :zap: 78848 | :snail: 14.0008ms |
| *SelectWithUpdateColumn* | 4393 | 213280 | 181.0103ms | | | |
| SelectWithUpdateColumn | 4072 | 240536 | 189.0108ms | | | |
| diffs | | | | :snail: 321 | :zap: 27256 | :zap: 8.0005ms |
| *OmitWithUpdateColumn* | 4395 | 213936 | 159.0091ms | | | |
| OmitWithUpdateColumn | 4069 | 240504 | 176.0101ms | | | |
| diffs | | | | :snail: 326 | :zap: 26568 | :zap: 17.001ms |
| *UpdateColumnsSkipsAssociations* | 4328 | 208096 | 232.0133ms | | | |
| UpdateColumnsSkipsAssociations | 4044 | 238024 | 258.0147ms | | | |
| diffs | | | | :snail: 284 | :zap: 29928 | :zap: 26.0014ms |
| *UpdatesWithBlankValues* | 1288 | 61696 | 129.0074ms | | | |
| UpdatesWithBlankValues | 1124 | 58648 | 154.0088ms | | | |
| diffs | | | | :snail: 164 | :snail: 3048 | :zap: 25.0014ms |
| *UpdatesTableWithIgnoredValues* | 435 | 16728 | 117.0067ms | | | |
| UpdatesTableWithIgnoredValues | 527 | 27680 | 207.0118ms | | | |
| diffs | | | | :zap: 92 | :zap: 10952 | :zap: 90.0051ms |
| *UpdateDecodeVirtualAttributes* | 1046 | 54312 | 135.0078ms | | | |
| UpdateDecodeVirtualAttributes | 929 | 51008 | 141.008ms | | | |
| diffs | | | | :snail: 117 | :snail: 3304 | :zap: 6.0002ms |
| *ToDBNameGenerateFriendlyName* | 121 | 5344 | nothing. | | | |
| ToDBNameGenerateFriendlyName | 123 | 5152 | nothing. | | | |
| diffs | | | | :zap: 2 | :snail: 192 | :zzz: |
| *SkipSaveAssociation* | 1323 | 55952 | 481.0276ms | | | |
| SkipSaveAssociation | 1365 | 74616 | 553.0316ms | | | |
| diffs | | | | :zap: 42 | :zap: 18664 | :zap: 72.004ms |
| TOTAL (original) | 610151 | 49218240 | 56.9942599s | | | |
| TOTAL (new) | 618719 | 36000280 | 51.9739731s | | | |