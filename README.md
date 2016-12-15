## side by side gorm

Some tests to check gorm differences after refactoring

Example (test run on 15th of December 2016) of result produced:

| Test name | Allocs | Bytes | Duration  |
| :-------: | -----: | ----: | --------: 
| *OpenTestConnection* | 128 | 12976 | 1ms |
| OpenTestConnection | 99 | 10896 | nothing. |
| Diffs |  :snail: 29 | :snail: 2080 | :snail: 1ms |
| *RunNewMigration* | 9178 | 3069208 | 5.9183385s |
| RunMigration | 17590 | 10690824 | 7.674439s |
| Diffs |  :zap: 8412 | :zap: 7621616 | :zap: 1.7561005s |
| *StringPrimaryKey* | 588 | 23560 | 342.0195ms |
| StringPrimaryKey | 983 | 434672 | 544.0311ms |
| Diffs |  :zap: 395 | :zap: 411112 | :zap: 202.0116ms |
| *SetTable* | 19776 | 995000 | 1.4240815s |
| SetTable | 19593 | 1935368 | 841.0481ms |
| Diffs |  :snail: 183 | :zap: 940368 | :snail: 583.0334ms |
| *ExceptionsWithInvalidSql* | 1361 | 77008 | 1ms |
| ExceptionsWithInvalidSql | 2318 | 1101432 | 1.0001ms |
| Diffs |  :zap: 957 | :zap: 1024424 | :zap: 100ns |
| *HasTable* | 338 | 12600 | 403.0231ms |
| HasTable | 696 | 418256 | 2.0351164s |
| Diffs |  :zap: 358 | :zap: 405656 | :zap: 1.6320933s |
| *TableName* | 187 | 12752 | nothing. |
| TableName | 161 | 22432 | nothing. |
| Diffs |  :snail: 26 | :zap: 9680 | :zzz: |
| *NullValues* | 1475 | 58272 | 467.0267ms |
| NullValues | 2226 | 881832 | 259.0148ms |
| Diffs |  :zap: 751 | :zap: 823560 | :snail: 208.0119ms |
| *NullValuesWithFirstOrCreate* | 1144 | 57384 | 115.0066ms |
| NullValuesWithFirstOrCreate | 932 | 55704 | 58.0033ms |
| Diffs |  :snail: 212 | :snail: 1680 | :snail: 57.0033ms |
| *Transaction* | 4131 | 215248 | 92.0053ms |
| Transaction | 4134 | 629216 | 267.0153ms |
| Diffs |  :zap: 3 | :zap: 413968 | :zap: 175.01ms |
| *Row* | 2424 | 125920 | 100.0057ms |
| Row | 2344 | 145200 | 100.0057ms |
| Diffs |  :snail: 80 | :zap: 19280 | :zzz: |
| *Rows* | 2437 | 126560 | 100.0057ms |
| Rows | 2356 | 145632 | 91.0052ms |
| Diffs |  :snail: 81 | :zap: 19072 | :snail: 9.0005ms |
| *ScanRows* | 2564 | 132408 | 83.0047ms |
| ScanRows | 2473 | 153048 | 83.0048ms |
| Diffs |  :snail: 91 | :zap: 20640 | :zap: 100ns |
| *Scan* | 2766 | 143152 | 100.0057ms |
| Scan | 2858 | 182112 | 83.0047ms |
| Diffs |  :zap: 92 | :zap: 38960 | :snail: 17.001ms |
| *Raw* | 2940 | 154584 | 108.0062ms |
| Raw | 3053 | 193248 | 127.0073ms |
| Diffs |  :zap: 113 | :zap: 38664 | :zap: 19.0011ms |
| *Group* | 127 | 5200 | nothing. |
| Group | 118 | 5584 | 1ms |
| Diffs |  :snail: 9 | :zap: 384 | :zap: 1ms |
| *Joins* | 3776 | 221296 | 48.0028ms |
| Joins | 4038 | 274632 | 50.0028ms |
| Diffs |  :zap: 262 | :zap: 53336 | :zap: 2ms |
| *JoinsWithSelect* | 1213 | 56376 | 32.0018ms |
| JoinsWithSelect | 1322 | 84896 | 25.0015ms |
| Diffs |  :zap: 109 | :zap: 28520 | :snail: 7.0003ms |
| *Having* | 113 | 5712 | 1ms |
| Having | 195 | 13288 | nothing. |
| Diffs |  :zap: 82 | :zap: 7576 | :snail: 1ms |
| *TimeWithZone* | 3908 | 269008 | 66.0038ms |
| TimeWithZone | 3708 | 282432 | 58.0033ms |
| Diffs |  :snail: 200 | :zap: 13424 | :snail: 8.0005ms |
| *Hstore* | 28 | 1136 | nothing. |
| Hstore | 31 | 1232 | nothing. |
| Diffs |  :zap: 3 | :zap: 96 | :zzz: |
| *SetAndGet* | 23 | 1184 | nothing. |
| SetAndGet | 28 | 1808 | nothing. |
| Diffs |  :zap: 5 | :zap: 624 | :zzz: |
| *CompatibilityMode* | 747 | 52632 | nothing. |
| CompatibilityMode | 527 | 35320 | nothing. |
| Diffs |  :snail: 220 | :snail: 17312 | :zzz: |
| *OpenExistingDB* | 1149 | 61200 | 89.0051ms |
| OpenExistingDB | 1061 | 66792 | 51.003ms |
| Diffs |  :snail: 88 | :zap: 5592 | :snail: 38.0021ms |
| *DdlErrors* | 315 | 22136 | 1ms |
| DdlErrors | 603 | 415560 | 1.0001ms |
| Diffs |  :zap: 288 | :zap: 393424 | :zap: 100ns |
| *OpenWithOneParameter* | 20 | 864 | nothing. |
| OpenWithOneParameter | 23 | 976 | nothing. |
| Diffs |  :zap: 3 | :zap: 112 | :zzz: |
| *BelongsTo* | 10340 | 568152 | 404.0231ms |
| BelongsTo | 11939 | 1527752 | 418.0239ms |
| Diffs |  :zap: 1599 | :zap: 959600 | :zap: 14.0008ms |
| *BelongsToOverrideForeignKey1* | 349 | 16864 | nothing. |
| BelongsToOverrideForeignKey1 | 341 | 20120 | nothing. |
| Diffs |  :snail: 8 | :zap: 3256 | :zzz: |
| *BelongsToOverrideForeignKey2* | 279 | 13624 | nothing. |
| BelongsToOverrideForeignKey2 | 247 | 17528 | nothing. |
| Diffs |  :snail: 32 | :zap: 3904 | :zzz: |
| *HasOne* | 15143 | 830160 | 330.0188ms |
| HasOne | 15245 | 945632 | 346.0198ms |
| Diffs |  :zap: 102 | :zap: 115472 | :zap: 16.001ms |
| *HasOneOverrideForeignKey1* | 307 | 20216 | nothing. |
| HasOneOverrideForeignKey1 | 273 | 18248 | nothing. |
| Diffs |  :snail: 34 | :snail: 1968 | :zzz: |
| *HasOneOverrideForeignKey2* | 273 | 13576 | nothing. |
| HasOneOverrideForeignKey2 | 246 | 17464 | nothing. |
| Diffs |  :snail: 27 | :zap: 3888 | :zzz: |
| *HasMany* | 11913 | 672392 | 404.0231ms |
| Many | 12337 | 848800 | 542.031ms |
| Diffs |  :zap: 424 | :zap: 176408 | :zap: 138.0079ms |
| *HasManyOverrideForeignKey1* | 302 | 15136 | 1ms |
| HasManyOverrideForeignKey1 | 269 | 17808 | nothing. |
| Diffs |  :snail: 33 | :zap: 2672 | :snail: 1ms |
| *HasManyOverrideForeignKey2* | 269 | 14656 | nothing. |
| HasManyOverrideForeignKey2 | 250 | 20504 | 1.0001ms |
| Diffs |  :snail: 19 | :zap: 5848 | :zap: 1.0001ms |
| *ManyToMany* | 25211 | 1355856 | 1.5550889s |
| ManyToMany | 27399 | 1727680 | 1.3910795s |
| Diffs |  :zap: 2188 | :zap: 371824 | :snail: 164.0094ms |
| *Related* | 7506 | 400200 | 127.0073ms |
| Related | 7104 | 433768 | 66.0038ms |
| Diffs |  :snail: 402 | :zap: 33568 | :snail: 61.0035ms |
| *ForeignKey* | 53 | 4672 | nothing. |
| ForeignKey | 60 | 6896 | nothing. |
| Diffs |  :zap: 7 | :zap: 2224 | :zzz: |
| *LongForeignKey* | 304 | 52808 | 1.4300818s |
| LongForeignKey | 366 | 102768 | 566.0324ms |
| Diffs |  :zap: 62 | :zap: 49960 | :snail: 864.0494ms |
| *LongForeignKeyWithShortDest* | 268 | 11312 | 524.0299ms |
| LongForeignKeyWithShortDest | 342 | 101728 | 624.0357ms |
| Diffs |  :zap: 74 | :zap: 90416 | :zap: 100.0058ms |
| *HasManyChildrenWithOneStruct* | 686 | 29152 | 50.0028ms |
| HasManyChildrenWithOneStruct | 640 | 43000 | 41.0023ms |
| Diffs |  :snail: 46 | :zap: 13848 | :snail: 9.0005ms |
| *RunCallbacks* | 2703 | 133984 | 109.0063ms |
| RunCallbacks | 2643 | 148376 | 191.0109ms |
| Diffs |  :snail: 60 | :zap: 14392 | :zap: 82.0046ms |
| *CallbacksWithErrors* | 5174 | 242608 | 208.0119ms |
| CallbacksWithErrors | 8577 | 4306136 | 166.0095ms |
| Diffs |  :zap: 3403 | :zap: 4063528 | :snail: 42.0024ms |
| *Create* | 2537 | 137192 | 79.0045ms |
| Create | 2016 | 110560 | 163.0094ms |
| Diffs |  :snail: 521 | :snail: 26632 | :zap: 84.0049ms |
| *CreateWithAutoIncrement* | 31 | 1728 | 1ms |
| CreateWithAutoIncrement | 34 | 1824 | nothing. |
| Diffs |  :zap: 3 | :zap: 96 | :snail: 1ms |
| *CreateWithNoGORMPrimayKey* | 256 | 11432 | 83.0048ms |
| CreateWithNoGORMPrimayKey | 267 | 18376 | 86.0049ms |
| Diffs |  :zap: 11 | :zap: 6944 | :zap: 3.0001ms |
| *CreateWithNoStdPrimaryKeyAndDefaultValues* | 1019 | 46552 | 172.0099ms |
| CreateWithNoStdPrimaryKeyAndDefaultValues | 1123 | 74184 | 197.0112ms |
| Diffs |  :zap: 104 | :zap: 27632 | :zap: 25.0013ms |
| *AnonymousScanner* | 1129 | 58912 | 31.0018ms |
| AnonymousScanner | 1048 | 62376 | 78.0045ms |
| Diffs |  :snail: 81 | :zap: 3464 | :zap: 47.0027ms |
| *AnonymousField* | 1622 | 84296 | 459.0262ms |
| AnonymousField | 1562 | 95784 | 499.0286ms |
| Diffs |  :snail: 60 | :zap: 11488 | :zap: 40.0024ms |
| *SelectWithCreate* | 3108 | 151104 | 233.0134ms |
| SelectWithCreate | 3250 | 214576 | 209.0119ms |
| Diffs |  :zap: 142 | :zap: 63472 | :snail: 24.0015ms |
| *OmitWithCreate* | 3197 | 164544 | 58.0033ms |
| OmitWithCreate | 3303 | 215992 | 100.0057ms |
| Diffs |  :zap: 106 | :zap: 51448 | :zap: 42.0024ms |
| *CustomizeColumn* | 872 | 37392 | 340.0195ms |
| CustomizeColumn | 1227 | 461048 | 275.0157ms |
| Diffs |  :zap: 355 | :zap: 423656 | :snail: 65.0038ms |
| *CustomColumnAndIgnoredFieldClash* | 193 | 15352 | 199.0114ms |
| CustomColumnAndIgnoredFieldClash | 542 | 409768 | 242.0138ms |
| Diffs |  :zap: 349 | :zap: 394416 | :zap: 43.0024ms |
| *ManyToManyWithCustomizedColumn* | 1588 | 74224 | 842.0482ms |
| ManyToManyWithCustomizedColumn | 2580 | 549488 | 2.150123s |
| Diffs |  :zap: 992 | :zap: 475264 | :zap: 1.3080748s |
| *OneToOneWithCustomizedColumn* | 1591 | 74256 | 875.05ms |
| OneToOneWithCustomizedColumn | 1920 | 497032 | 954.0546ms |
| Diffs |  :zap: 329 | :zap: 422776 | :zap: 79.0046ms |
| *OneToManyWithCustomizedColumn* | 3340 | 165152 | 946.0541ms |
| OneToManyWithCustomizedColumn | 3817 | 614560 | 1.0170582s |
| Diffs |  :zap: 477 | :zap: 449408 | :zap: 71.0041ms |
| *HasOneWithPartialCustomizedColumn* | 2295 | 111648 | 1.2830733s |
| HasOneWithPartialCustomizedColumn | 2760 | 546040 | 842.0482ms |
| Diffs |  :zap: 465 | :zap: 434392 | :snail: 441.0251ms |
| *BelongsToWithPartialCustomizedColumn* | 2521 | 124248 | 1.2660724s |
| BelongsToWithPartialCustomizedColumn | 3001 | 563968 | 1.2070691s |
| Diffs |  :zap: 480 | :zap: 439720 | :snail: 59.0033ms |
| *Delete* | 2268 | 119312 | 176.01ms |
| Delete | 2101 | 125888 | 185.0106ms |
| Diffs |  :snail: 167 | :zap: 6576 | :zap: 9.0006ms |
| *InlineDelete* | 2309 | 121664 | 704.0402ms |
| InlineDelete | 2242 | 137064 | 288.0165ms |
| Diffs |  :snail: 67 | :zap: 15400 | :snail: 416.0237ms |
| *SoftDelete* | 985 | 41272 | 1.3960799s |
| SoftDelete | 1237 | 78336 | 492.0281ms |
| Diffs |  :zap: 252 | :zap: 37064 | :snail: 904.0518ms |
| *PrefixColumnNameForEmbeddedStruct* | 458 | 20400 | 6.0003ms |
| PrefixColumnNameForEmbeddedStruct | 448 | 31672 | 5.0003ms |
| Diffs |  :snail: 10 | :zap: 11272 | :snail: 1ms |
| *SaveAndQueryEmbeddedStruct* | 1217 | 48624 | 239.0137ms |
| SaveAndQueryEmbeddedStruct | 1296 | 70672 | 183.0105ms |
| Diffs |  :zap: 79 | :zap: 22048 | :snail: 56.0032ms |
| *CalculateField* | 495 | 28504 | 1ms |
| CalculateField | 451 | 34872 | nothing. |
| Diffs |  :snail: 44 | :zap: 6368 | :snail: 1ms |
| *JoinTable* | 2951 | 160248 | 459.0263ms |
| JoinTable | 4602 | 594792 | 1.4640837s |
| Diffs |  :zap: 1651 | :zap: 434544 | :zap: 1.0050574s |
| *Indexes* | 8994 | 426960 | 1.7010973s |
| Indexes | 10780 | 3156624 | 2.0491172s |
| Diffs |  :zap: 1786 | :zap: 2729664 | :zap: 348.0199ms |
| *AutoMigration* | 1444 | 62600 | 2.2101264s |
| AutoMigration | 1434 | 62800 | 1.2330706s |
| Diffs |  :snail: 10 | :zap: 200 | :snail: 977.0558ms |
| *MultipleIndexes* | 2521 | 102816 | 1.644094s |
| MultipleIndexes | 3138 | 927128 | 2.7291561s |
| Diffs |  :zap: 617 | :zap: 824312 | :zap: 1.0850621s |
| *ManyToManyWithMultiPrimaryKeys* | 8870 | 441600 | 1.6590949s |
| ManyToManyWithMultiPrimaryKeys | 13149 | 1622904 | 1.783102s |
| Diffs |  :zap: 4279 | :zap: 1181304 | :zap: 124.0071ms |
| *ManyToManyWithCustomizedForeignKeys* | 10364 | 536704 | 1.3750786s |
| ManyToManyWithCustomizedForeignKeys | 14262 | 960176 | 2.5861479s |
| Diffs |  :zap: 3898 | :zap: 423472 | :zap: 1.2110693s |
| *ManyToManyWithCustomizedForeignKeys2* | 13998 | 734688 | 1.5560891s |
| ManyToManyWithCustomizedForeignKeys2 | 19582 | 1314992 | 2.3661353s |
| Diffs |  :zap: 5584 | :zap: 580304 | :zap: 810.0462ms |
| *PointerFields* | 1947 | 77656 | 483.0276ms |
| PointerFields | 2854 | 554888 | 409.0234ms |
| Diffs |  :zap: 907 | :zap: 477232 | :snail: 74.0042ms |
| *Polymorphic* | 18023 | 913672 | 857.049ms |
| Polymorphic | 24242 | 1695944 | 699.04ms |
| Diffs |  :zap: 6219 | :zap: 782272 | :snail: 158.009ms |
| *NamedPolymorphic* | 11136 | 614288 | 649.0372ms |
| NamedPolymorphic | 15834 | 1123960 | 609.0348ms |
| Diffs |  :zap: 4698 | :zap: 509672 | :snail: 40.0024ms |
| *Preload* | 21755 | 1053280 | 1.4310818s |
| Preload | 21985 | 1314776 | 367.021ms |
| Diffs |  :zap: 230 | :zap: 261496 | :snail: 1.0640608s |
| *NestedPreload1* | 1714 | 113112 | 694.0397ms |
| NestedPreload1 | 1911 | 139352 | 717.041ms |
| Diffs |  :zap: 197 | :zap: 26240 | :zap: 23.0013ms |
| *NestedPreload2* | 2218 | 91688 | 986.0564ms |
| NestedPreload2 | 2391 | 154264 | 1.1450655s |
| Diffs |  :zap: 173 | :zap: 62576 | :zap: 159.0091ms |
| *NestedPreload3* | 1991 | 83352 | 2.2861308s |
| NestedPreload3 | 2152 | 133360 | 1.2790732s |
| Diffs |  :zap: 161 | :zap: 50008 | :snail: 1.0070576s |
| *NestedPreload4* | 1779 | 76064 | 1.3180753s |
| NestedPreload4 | 1926 | 121680 | 1.1340649s |
| Diffs |  :zap: 147 | :zap: 45616 | :snail: 184.0104ms |
| *NestedPreload5* | 2200 | 91640 | 1.9621122s |
| NestedPreload5 | 2371 | 144232 | 1.188068s |
| Diffs |  :zap: 171 | :zap: 52592 | :snail: 774.0442ms |
| *NestedPreload6* | 3420 | 145176 | 942.0539ms |
| NestedPreload6 | 3643 | 216696 | 1.1340648s |
| Diffs |  :zap: 223 | :zap: 71520 | :zap: 192.0109ms |
| *NestedPreload7* | 3047 | 129000 | 1.1330649s |
| NestedPreload7 | 3262 | 192072 | 1.2180696s |
| Diffs |  :zap: 215 | :zap: 63072 | :zap: 85.0047ms |
| *NestedPreload8* | 2596 | 107144 | 1.7571005s |
| NestedPreload8 | 2789 | 168040 | 941.0538ms |
| Diffs |  :zap: 193 | :zap: 60896 | :snail: 816.0467ms |
| *NestedPreload9* | 5912 | 266176 | 1.4090806s |
| NestedPreload9 | 6189 | 361904 | 2.1541232s |
| Diffs |  :zap: 277 | :zap: 95728 | :zap: 745.0426ms |
| *NestedPreload10* | 2185 | 90336 | 1.2740728s |
| NestedPreload10 | 2336 | 137736 | 663.038ms |
| Diffs |  :zap: 151 | :zap: 47400 | :snail: 611.0348ms |
| *NestedPreload11* | 1718 | 73608 | 798.0456ms |
| NestedPreload11 | 2197 | 143928 | 743.0425ms |
| Diffs |  :zap: 479 | :zap: 70320 | :snail: 55.0031ms |
| *NestedPreload12* | 3072 | 142064 | 625.0357ms |
| NestedPreload12 | 3256 | 204416 | 811.0464ms |
| Diffs |  :zap: 184 | :zap: 62352 | :zap: 186.0107ms |
| *ManyToManyPreloadWithMultiPrimaryKeys* | 5813 | 282152 | 994.0569ms |
| ManyToManyPreloadWithMultiPrimaryKeys | 7862 | 511400 | 1.2100692s |
| Diffs |  :zap: 2049 | :zap: 229248 | :zap: 216.0123ms |
| *ManyToManyPreloadForNestedPointer* | 6210 | 279192 | 1.4490829s |
| ManyToManyPreloadForNestedPointer | 8488 | 543960 | 2.658152s |
| Diffs |  :zap: 2278 | :zap: 264768 | :zap: 1.2090691s |
| *NestedManyToManyPreload* | 3775 | 170304 | 1.6140923s |
| NestedManyToManyPreload | 5379 | 358376 | 1.5660896s |
| Diffs |  :zap: 1604 | :zap: 188072 | :snail: 48.0027ms |
| *NestedManyToManyPreload2* | 2500 | 109920 | 1.7340992s |
| NestedManyToManyPreload2 | 3358 | 218776 | 3.3151896s |
| Diffs |  :zap: 858 | :zap: 108856 | :zap: 1.5810904s |
| *NestedManyToManyPreload3* | 4831 | 212016 | 1.7501001s |
| NestedManyToManyPreload3 | 5987 | 397272 | 2.3501344s |
| Diffs |  :zap: 1156 | :zap: 185256 | :zap: 600.0343ms |
| *NestedManyToManyPreload3ForStruct* | 4719 | 206944 | 1.5840906s |
| NestedManyToManyPreload3ForStruct | 5894 | 384672 | 1.6910967s |
| Diffs |  :zap: 1175 | :zap: 177728 | :zap: 107.0061ms |
| *NestedManyToManyPreload4* | 3161 | 152032 | 2.0421168s |
| NestedManyToManyPreload4 | 4287 | 293608 | 2.3081321s |
| Diffs |  :zap: 1126 | :zap: 141576 | :zap: 266.0153ms |
| *ManyToManyPreloadForPointer* | 4561 | 206912 | 2.4931426s |
| ManyToManyPreloadForPointer | 6537 | 431208 | 1.3170753s |
| Diffs |  :zap: 1976 | :zap: 224296 | :snail: 1.1760673s |
| *NilPointerSlice* | 1858 | 77912 | 2.2311276s |
| NilPointerSlice | 1999 | 121800 | 1.0690612s |
| Diffs |  :zap: 141 | :zap: 43888 | :snail: 1.1620664s |
| *NilPointerSlice2* | 1747 | 75792 | 1.9161096s |
| NilPointerSlice2 | 2138 | 142192 | 1.7260987s |
| Diffs |  :zap: 391 | :zap: 66400 | :snail: 190.0109ms |
| *PrefixedPreloadDuplication* | 4995 | 207928 | 2.6861537s |
| PrefixedPreloadDuplication | 5254 | 323136 | 2.3921368s |
| Diffs |  :zap: 259 | :zap: 115208 | :snail: 294.0169ms |
| *FirstAndLast* | 4485 | 242488 | 110.0063ms |
| FirstAndLast | 3735 | 217176 | 83.0047ms |
| Diffs |  :snail: 750 | :snail: 25312 | :snail: 27.0016ms |
| *FirstAndLastWithNoStdPrimaryKey* | 1425 | 67168 | 73.0042ms |
| FirstAndLastWithNoStdPrimaryKey | 1466 | 93632 | 143.0082ms |
| Diffs |  :zap: 41 | :zap: 26464 | :zap: 70.004ms |
| *UIntPrimaryKey* | 511 | 25136 | 1ms |
| UIntPrimaryKey | 440 | 27480 | nothing. |
| Diffs |  :snail: 71 | :zap: 2344 | :snail: 1ms |
| *StringPrimaryKeyForNumericValueStartingWithZero* | 413 | 16816 | 356.0204ms |
| StringPrimaryKeyForNumericValueStartingWithZero | 522 | 32496 | 207.0118ms |
| Diffs |  :zap: 109 | :zap: 15680 | :snail: 149.0086ms |
| *FindAsSliceOfPointers* | 20617 | 1290800 | 28.0016ms |
| FindAsSliceOfPointers | 15636 | 895152 | 17.001ms |
| Diffs |  :snail: 4981 | :snail: 395648 | :snail: 11.0006ms |
| *SearchWithPlainSQL* | 9875 | 650728 | 77.0044ms |
| SearchWithPlainSQL | 9503 | 646488 | 108.0062ms |
| Diffs |  :snail: 372 | :snail: 4240 | :zap: 31.0018ms |
| *SearchWithStruct* | 7322 | 426512 | 122.0069ms |
| SearchWithStruct | 5939 | 347192 | 76.0044ms |
| Diffs |  :snail: 1383 | :snail: 79320 | :snail: 46.0025ms |
| *SearchWithMap* | 5983 | 337272 | 90.0051ms |
| SearchWithMap | 5008 | 306808 | 92.0053ms |
| Diffs |  :snail: 975 | :snail: 30464 | :zap: 2.0002ms |
| *SearchWithEmptyChain* | 4111 | 223584 | 198.0113ms |
| SearchWithEmptyChain | 3842 | 233328 | 67.0038ms |
| Diffs |  :snail: 269 | :zap: 9744 | :snail: 131.0075ms |
| *Select* | 1062 | 56040 | 25.0014ms |
| Select | 982 | 57568 | 24.0014ms |
| Diffs |  :snail: 80 | :zap: 1528 | :snail: 1ms |
| *OrderAndPluck* | 15340 | 956944 | 72.0041ms |
| OrderAndPluck | 11864 | 699896 | 90.0052ms |
| Diffs |  :snail: 3476 | :snail: 257048 | :zap: 18.0011ms |
| *Limit* | 20312 | 1357696 | 125.0071ms |
| Limit | 15773 | 1058840 | 114.0065ms |
| Diffs |  :snail: 4539 | :snail: 298856 | :snail: 11.0006ms |
| *Offset* | 88737 | 5802616 | 520.0297ms |
| Offset | 68809 | 4351768 | 489.028ms |
| Diffs |  :snail: 19928 | :snail: 1450848 | :snail: 31.0017ms |
| *Or* | 2508 | 136496 | 55.0031ms |
| Or | 2342 | 147208 | 83.0047ms |
| Diffs |  :snail: 166 | :zap: 10712 | :zap: 28.0016ms |
| *Count* | 3231 | 177528 | 76.0044ms |
| Count | 3306 | 205992 | 74.0042ms |
| Diffs |  :zap: 75 | :zap: 28464 | :snail: 2.0002ms |
| *Not* | 21839 | 1189088 | 561.0321ms |
| Not | 20980 | 1544408 | 435.0249ms |
| Diffs |  :snail: 859 | :zap: 355320 | :snail: 126.0072ms |
| *FillSmallerStruct* | 894 | 42792 | 46.0026ms |
| FillSmallerStruct | 909 | 55568 | 57.0033ms |
| Diffs |  :zap: 15 | :zap: 12776 | :zap: 11.0007ms |
| *FindOrInitialize* | 7060 | 401904 | 259.0149ms |
| FindOrInitialize | 5115 | 279344 | 83.0047ms |
| Diffs |  :snail: 1945 | :snail: 122560 | :snail: 176.0102ms |
| *FindOrCreate* | 11807 | 642936 | 241.0138ms |
| FindOrCreate | 10107 | 1332360 | 183.0105ms |
| Diffs |  :snail: 1700 | :zap: 689424 | :snail: 58.0033ms |
| *SelectWithEscapedFieldName* | 2299 | 118592 | 67.0038ms |
| SelectWithEscapedFieldName | 1990 | 120408 | 75.0043ms |
| Diffs |  :snail: 309 | :zap: 1816 | :zap: 8.0005ms |
| *SelectWithVariables* | 691 | 34528 | 25.0014ms |
| SelectWithVariables | 630 | 38792 | 25.0014ms |
| Diffs |  :snail: 61 | :zap: 4264 | :zzz: |
| *FirstAndLastWithRaw* | 2609 | 134272 | 42.0024ms |
| FirstAndLastWithRaw | 2433 | 146128 | 50.0029ms |
| Diffs |  :snail: 176 | :zap: 11856 | :zap: 8.0005ms |
| *ScannableSlices* | 510 | 22256 | 299.0171ms |
| ScannableSlices | 496 | 26728 | 232.0132ms |
| Diffs |  :snail: 14 | :zap: 4472 | :snail: 67.0039ms |
| *Scopes* | 3565 | 205264 | 84.0048ms |
| Scopes | 3286 | 206592 | 133.0076ms |
| Diffs |  :snail: 279 | :zap: 1328 | :zap: 49.0028ms |
| *Update* | 6528 | 313360 | 416.0238ms |
| Update | 6190 | 346064 | 449.0257ms |
| Diffs |  :snail: 338 | :zap: 32704 | :zap: 33.0019ms |
| *UpdateWithNoStdPrimaryKeyAndDefaultValues* | 2896 | 129536 | 937.0536ms |
| UpdateWithNoStdPrimaryKeyAndDefaultValues | 2783 | 163448 | 963.0551ms |
| Diffs |  :snail: 113 | :zap: 33912 | :zap: 26.0015ms |
| *Updates* | 4533 | 210640 | 183.0105ms |
| Updates | 4333 | 239976 | 242.0138ms |
| Diffs |  :snail: 200 | :zap: 29336 | :zap: 59.0033ms |
| *UpdateColumn* | 3124 | 147000 | 175.01ms |
| UpdateColumn | 2667 | 143272 | 217.0124ms |
| Diffs |  :snail: 457 | :snail: 3728 | :zap: 42.0024ms |
| *SelectWithUpdate* | 7033 | 336408 | 201.0115ms |
| SelectWithUpdate | 7038 | 443592 | 250.0143ms |
| Diffs |  :zap: 5 | :zap: 107184 | :zap: 49.0028ms |
| *SelectWithUpdateWithMap* | 7063 | 338880 | 192.011ms |
| SelectWithUpdateWithMap | 7088 | 448704 | 377.0216ms |
| Diffs |  :zap: 25 | :zap: 109824 | :zap: 185.0106ms |
| *OmitWithUpdate* | 5987 | 292824 | 131.0075ms |
| OmitWithUpdate | 5990 | 378352 | 150.0086ms |
| Diffs |  :zap: 3 | :zap: 85528 | :zap: 19.0011ms |
| *OmitWithUpdateWithMap* | 5809 | 285112 | 117.0067ms |
| OmitWithUpdateWithMap | 5848 | 372328 | 310.0177ms |
| Diffs |  :zap: 39 | :zap: 87216 | :zap: 193.011ms |
| *SelectWithUpdateColumn* | 4417 | 214168 | 375.0215ms |
| SelectWithUpdateColumn | 4054 | 249312 | 125.0071ms |
| Diffs |  :snail: 363 | :zap: 35144 | :snail: 250.0144ms |
| *OmitWithUpdateColumn* | 4421 | 214536 | 229.0131ms |
| OmitWithUpdateColumn | 4046 | 248224 | 198.0114ms |
| Diffs |  :snail: 375 | :zap: 33688 | :snail: 31.0017ms |
| *UpdateColumnsSkipsAssociations* | 4357 | 208456 | 135.0077ms |
| UpdateColumnsSkipsAssociations | 4028 | 245888 | 959.0548ms |
| Diffs |  :snail: 329 | :zap: 37432 | :zap: 824.0471ms |
| *UpdatesWithBlankValues* | 1254 | 61424 | 99.0057ms |
| UpdatesWithBlankValues | 1067 | 57896 | 75.0043ms |
| Diffs |  :snail: 187 | :snail: 3528 | :snail: 24.0014ms |
| *UpdatesTableWithIgnoredValues* | 400 | 15848 | 117.0067ms |
| UpdatesTableWithIgnoredValues | 498 | 27280 | 100.0057ms |
| Diffs |  :zap: 98 | :zap: 11432 | :snail: 17.001ms |
| *UpdateDecodeVirtualAttributes* | 1042 | 53832 | 83.0048ms |
| UpdateDecodeVirtualAttributes | 900 | 50304 | 75.0042ms |
| Diffs |  :snail: 142 | :snail: 3528 | :snail: 8.0006ms |
| *ToDBNameGenerateFriendlyName* | 120 | 5056 | nothing. |
| ToDBNameGenerateFriendlyName | 123 | 5152 | nothing. |
| Diffs |  :zap: 3 | :zap: 96 | :zzz: |
| *SkipSaveAssociation* | 1408 | 57552 | 2.1251216s |
| SkipSaveAssociation | 1452 | 77928 | 1.6670953s |
| Diffs |  :zap: 44 | :zap: 20376 | :snail: 458.0263ms |
| TOTAL (original) | 674659 | 69142816 | 1m25.9979189s |
| TOTAL (new) | 652590 | 37905616 | 1m22.794735s |
| TOTAL (diffs) |  :zap: 22069 |  :zap: 31237200 |  :zap: 3.2031839s |